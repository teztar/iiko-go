package iiko

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

// fakeIiko отвечает как iikoCloud после 02.11.2026: ключ принимается только
// в /api/v2/access_token, старый /api/1/access_token отвечает 401.
type fakeIiko struct {
	mu        sync.Mutex
	tokenReqs []string
	bodies    []map[string]any
	expired   bool
}

func (f *fakeIiko) handler(t *testing.T) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		raw, _ := io.ReadAll(r.Body)
		f.mu.Lock()
		defer f.mu.Unlock()

		switch r.URL.Path {
		case "/api/1/access_token", "/api/v2/access_token":
			var body map[string]any
			if err := json.Unmarshal(raw, &body); err != nil {
				t.Errorf("тело запроса токена не JSON: %v", err)
			}
			f.tokenReqs = append(f.tokenReqs, r.URL.Path)
			f.bodies = append(f.bodies, body)
			if r.URL.Path == "/api/1/access_token" {
				w.WriteHeader(http.StatusUnauthorized)
				_, _ = w.Write([]byte(`{"errorDescription":"This API key does not support /api/1/access_token. Please use /api/v2/access_token instead."}`))
				return
			}
			f.expired = false
			_, _ = w.Write([]byte(`{"correlationId":"5b1b0e36-4e9c-4e51-8d3a-9a2b1c0d0000","token":"jwt-v2"}`))
		case "/api/1/organizations":
			if f.expired || r.Header.Get("Authorization") != "Bearer jwt-v2" {
				w.WriteHeader(http.StatusUnauthorized)
				_, _ = w.Write([]byte(`{"errorDescription":"token expired"}`))
				return
			}
			_, _ = w.Write([]byte(`{"correlationId":"5b1b0e36-4e9c-4e51-8d3a-9a2b1c0d0000","organizations":[]}`))
		default:
			t.Errorf("неожиданный запрос %s", r.URL.Path)
			w.WriteHeader(http.StatusNotFound)
		}
	})
}

func newTestClient(srv *httptest.Server, opts ...ClientOption) *Client {
	c := &Client{
		baseURL:              srv.URL,
		httpClient:           srv.Client(),
		apiLogin:             "api-key-32",
		timeout:              DefaultTimeout,
		refreshTokenInterval: DefaultRefreshTokenInterval,
		quit:                 make(chan struct{}),
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func TestRefreshTokenWithAppUsesV2(t *testing.T) {
	fake := &fakeIiko{}
	srv := httptest.NewServer(fake.handler(t))
	defer srv.Close()

	c := newTestClient(srv, WithApp("4038e1a2-bca1-4654-924f-d17fee404436", "secret"))
	if err := c.refreshToken(); err != nil {
		t.Fatalf("refreshToken: %v", err)
	}
	if c.getToken() != "jwt-v2" {
		t.Fatalf("token = %q, want jwt-v2", c.getToken())
	}
	if len(fake.tokenReqs) != 1 || fake.tokenReqs[0] != "/api/v2/access_token" {
		t.Fatalf("запросы токена %v, want только /api/v2/access_token", fake.tokenReqs)
	}

	// iiko не принимает лишние поля (additionalProperties: false), apiLogin слать нельзя.
	want := map[string]any{
		"apiKey":       "api-key-32",
		"appId":        "4038e1a2-bca1-4654-924f-d17fee404436",
		"clientSecret": "secret",
	}
	got := fake.bodies[0]
	if len(got) != len(want) {
		t.Fatalf("тело %v, want ровно %v", got, want)
	}
	for k, v := range want {
		if got[k] != v {
			t.Fatalf("тело[%s] = %v, want %v", k, got[k], v)
		}
	}
}

func TestRetryOn401RefreshesThroughV2(t *testing.T) {
	fake := &fakeIiko{}
	srv := httptest.NewServer(fake.handler(t))
	defer srv.Close()

	c := newTestClient(srv, WithApp("app", "secret"))
	if err := c.refreshToken(); err != nil {
		t.Fatalf("refreshToken: %v", err)
	}

	fake.mu.Lock()
	fake.expired = true
	fake.mu.Unlock()

	var resp map[string]any
	if err := c.post(true, "/api/1/organizations", map[string]any{}, &resp); err != nil {
		t.Fatalf("запрос после истечения токена: %v", err)
	}
	if len(fake.tokenReqs) != 2 || fake.tokenReqs[1] != "/api/v2/access_token" {
		t.Fatalf("запросы токена %v, want повтор через /api/v2/access_token", fake.tokenReqs)
	}
}

func TestRefreshTokenWithoutAppUsesV1(t *testing.T) {
	fake := &fakeIiko{}
	srv := httptest.NewServer(fake.handler(t))
	defer srv.Close()

	c := newTestClient(srv)
	err := c.refreshToken()
	if err == nil {
		t.Fatal("ждали ошибку старого метода")
	}
	if len(fake.tokenReqs) != 1 || fake.tokenReqs[0] != "/api/1/access_token" {
		t.Fatalf("запросы токена %v, want /api/1/access_token", fake.tokenReqs)
	}
	if got := fake.bodies[0]; len(got) != 1 || got["apiLogin"] != "api-key-32" {
		t.Fatalf("тело v1 %v, want только apiLogin", got)
	}
}
