package iiko

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// WebhookEventType represents the type of webhook event
type WebhookEventType string

const (
	StopListUpdateWebhookEvent WebhookEventType = "StopListUpdate"

	DeliveryOrderUpdateWebhookEvent WebhookEventType = "DeliveryOrderUpdate"
	DeliveryOrderErrorWebhookEvent  WebhookEventType = "DeliveryOrderError"
)

// WebhookEvent represents a generic webhook event
type WebhookEvent struct {
	// Event Type of the webhook event
	EventType      WebhookEventType `json:"eventType"`
	EventTime      time.Time        `json:"eventTime"`
	OrganizationID uuid.UUID        `json:"organizationId"`
	CorrelationID  uuid.UUID        `json:"correlationId"`
	EventInfo      json.RawMessage  `json:"eventInfo"`
}

// WebhookHandler represents a webhook event handler function
type WebhookHandler func(event *WebhookEvent) error

// WebhookHandlerRegistry manages webhook event handlers
type WebhookHandlerRegistry struct {
	handlers map[WebhookEventType][]WebhookHandler
}

// NewWebhookHandlerRegistry creates a new webhook handler registry
func NewWebhookHandlerRegistry() *WebhookHandlerRegistry {
	return &WebhookHandlerRegistry{
		handlers: make(map[WebhookEventType][]WebhookHandler),
	}
}

// RegisterHandler registers a handler for a specific event type
func (r *WebhookHandlerRegistry) RegisterHandler(eventType WebhookEventType, handler WebhookHandler) {
	r.handlers[eventType] = append(r.handlers[eventType], handler)
}

// HandleEvent processes a webhook event using registered handlers
func (r *WebhookHandlerRegistry) HandleEvent(event *WebhookEvent) error {
	handlers, exists := r.handlers[event.EventType]
	if !exists {
		return fmt.Errorf("no handlers registered for event type: %s", event.EventType)
	}

	for _, handler := range handlers {
		if err := handler(event); err != nil {
			return fmt.Errorf("handler error for event type %s: %w", event.EventType, err)
		}
	}

	return nil
}

// WebhookServer represents an HTTP server for handling webhooks
type WebhookServer struct {
	registry *WebhookHandlerRegistry
	secret   string // Optional webhook secret for validation
}

// NewWebhookServer creates a new webhook server
func NewWebhookServer(secret string) *WebhookServer {
	return &WebhookServer{
		registry: NewWebhookHandlerRegistry(),
		secret:   secret,
	}
}

// RegisterHandler registers a webhook event handler
func (s *WebhookServer) RegisterHandler(eventType WebhookEventType, handler WebhookHandler) {
	s.registry.RegisterHandler(eventType, handler)
}

// ServeHTTP handles incoming webhook requests
func (s *WebhookServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Only accept POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Validate webhook signature if secret is provided
	if s.secret != "" {
		if err := s.validateSignature(r, body); err != nil {
			http.Error(w, "Invalid signature", http.StatusUnauthorized)
			return
		}
	}

	// Parse webhook event
	var event WebhookEvent
	if err := json.Unmarshal(body, &event); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// Handle the event
	if err := s.registry.HandleEvent(&event); err != nil {
		http.Error(w, fmt.Sprintf("Failed to handle event: %v", err), http.StatusInternalServerError)
		return
	}

	// Return success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": "Webhook processed successfully",
	})
}

// validateSignature validates the webhook signature (placeholder implementation)
func (s *WebhookServer) validateSignature(r *http.Request, body []byte) error {
	// This is a placeholder implementation
	// In a real implementation, you would validate the signature using HMAC-SHA256
	// with the webhook secret and the request body
	signature := r.Header.Get("X-Iiko-Signature")
	if signature == "" {
		return fmt.Errorf("missing signature header")
	}

	// TODO: Implement actual signature validation
	// expectedSignature := generateHMACSignature(s.secret, body)
	// if !hmac.Equal([]byte(signature), []byte(expectedSignature)) {
	//     return fmt.Errorf("signature mismatch")
	// }

	return nil
}

// WebhookMiddleware provides middleware functionality for webhook handling
type WebhookMiddleware struct {
	next http.Handler
}

// NewWebhookMiddleware creates a new webhook middleware
func NewWebhookMiddleware(next http.Handler) *WebhookMiddleware {
	return &WebhookMiddleware{next: next}
}

// ServeHTTP implements the middleware pattern
func (m *WebhookMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Add CORS headers for webhook requests
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-Iiko-Signature")

	// Handle preflight requests
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Add request ID for tracking
	requestID := uuid.New().String()
	w.Header().Set("X-Request-ID", requestID)

	// Call next handler
	m.next.ServeHTTP(w, r)
}
