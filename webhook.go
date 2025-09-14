package iiko

import (
	"encoding/json"
	"fmt"
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

type WebhookHandlerFunc func(event *WebhookEvent) error

// WebhookHandler represents a webhook event handler function
type WebhookHandler struct {
	name  string
	handle WebhookHandlerFunc
}

// WebhookServer represents a server for handling webhooks
type WebhookServer struct {
	handlers map[WebhookEventType][]WebhookHandler
}

// NewWebhookServer creates a new webhook server
func NewWebhookServer() *WebhookServer {
	return &WebhookServer{
		handlers: make(map[WebhookEventType][]WebhookHandler),
	}
}

// RegisterHandler registers a webhook event handler
func (s *WebhookServer) RegisterHandler(eventType WebhookEventType, handlerName string, handler WebhookHandlerFunc) {
	s.handlers[eventType] = append(s.handlers[eventType], WebhookHandler{
		name:  handlerName,
		handle: handler,
	})
}

// HandleEvent processes a webhook event using registered handlers
func (s *WebhookServer) HandleEvent(event *WebhookEvent) error {
	handlers, exists := s.handlers[event.EventType]
	if !exists {
		return fmt.Errorf("no handlers registered for event type: %s", event.EventType)
	}

	for _, handler := range handlers {
		if err := handler.handle(event); err != nil {
			return fmt.Errorf("handler %s returned error for event type %s: %w", handler.name, event.EventType, err)
		}
	}

	return nil
}
