package iiko

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

var ErrMissingToken = errors.New("missing API token")

// doRequest performs a single POST to iikoCloud. When requiresAuth is set and
// iiko answers 401 (token expired/revoked), it refreshes the token once and
// retries the request a single time before giving up.
func (c *Client) doRequest(
	requiresAuth bool,
	endpoint string,
	body interface{},
	response interface{},
	opts ...Option,
) error {
	if requiresAuth && c.getToken() == "" {
		return ErrMissingToken
	}

	// Marshal json body for request once; reused on retry.
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return err
	}

	send := func() (*http.Response, error) {
		req, reqErr := http.NewRequest(http.MethodPost, c.baseURL+endpoint, bytes.NewBuffer(jsonBody))
		if reqErr != nil {
			return nil, reqErr
		}

		req.Header.Set("Content-Type", "application/json; charset=UTF-8")
		if requiresAuth {
			req.Header.Set("Authorization", "Bearer "+c.getToken())
		}
		req.Header.Set("Timeout", strconv.Itoa(int(c.timeout.Seconds())))

		for _, opt := range opts {
			opt.Apply(req)
		}

		return c.httpClient.Do(req)
	}

	resp, err := send()
	if err != nil {
		return err
	}

	// Token expired/revoked — refresh once and retry.
	if requiresAuth && resp.StatusCode == http.StatusUnauthorized {
		resp.Body.Close()
		if refreshErr := c.refreshToken(); refreshErr != nil {
			return refreshErr
		}
		resp, err = send()
		if err != nil {
			return err
		}
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errorResponse ErrorResponse
		if err = json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {
			return err
		}
		errorResponse.StatusCode = resp.StatusCode
		return &errorResponse
	}

	if err = json.NewDecoder(resp.Body).Decode(response); err != nil {
		return err
	}

	return nil
}

func (c *Client) post(requiresAuth bool, endpoint string, body interface{}, response interface{}, opts ...Option) error {
	return c.doRequest(requiresAuth, endpoint, body, response, opts...)
}

func (c *Client) Post(requiresAuth bool, endpoint string, body interface{}, response interface{}, opts ...Option) error {
	return c.doRequest(requiresAuth, endpoint, body, response, opts...)
}
