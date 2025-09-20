package iiko

import (
	"encoding/json"
	"strings"
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
	EventTime      EventTime        `json:"eventTime"`
	OrganizationID uuid.UUID        `json:"organizationId"`
	CorrelationID  uuid.UUID        `json:"correlationId"`
	EventInfo      json.RawMessage  `json:"eventInfo"`
}

// EventTime кастомный тип для поддержки разных форматов времени
type EventTime struct {
	time.Time
}

func (et *EventTime) UnmarshalJSON(data []byte) error {
	s := string(data)
	s = strings.Trim(s, `"`) // remove 
	if s == "null" || s == "" {
		et.Time = time.Time{}
		return nil
	}
	// parse RFC3339 format first for compatibility
	t, err := time.Parse(time.RFC3339, s)
	if err == nil {
		et.Time = t
		return nil
	}
	// iiko uses "2006-01-02 15:04:05.999" format that is not RFC3339
	t, err = time.Parse("2006-01-02 15:04:05.999", s)
	if err == nil {
		et.Time = t
		return nil
	}

	return fmt.Errorf("cannot parse eventTime: %s", s)
}

type WebhookHandlerFunc func(event *WebhookEvent) error

// WebhookHandler represents a webhook event handler function
type WebhookHandler struct {
	name   string
	handle WebhookHandlerFunc
}

// WebhookServer represents a server for handling webhooks
type WebhookServer struct {
	handlers map[WebhookEventType][]WebhookHandler
	secret   string
}

// NewWebhookServer creates a new webhook server
func NewWebhookServer(secret string) *WebhookServer {
	return &WebhookServer{
		handlers: make(map[WebhookEventType][]WebhookHandler),
		secret:   secret,
	}
}

// RegisterHandler registers a webhook event handler
func (s *WebhookServer) RegisterHandler(eventType WebhookEventType, handlerName string, handler WebhookHandlerFunc) {
	s.handlers[eventType] = append(s.handlers[eventType], WebhookHandler{
		name:   handlerName,
		handle: handler,
	})
}

// HandleEvent processes a webhook event using registered handlers
func (s *WebhookServer) HandleEvent(event *WebhookEvent, secret string) error {
	if s.secret != secret {
		return fmt.Errorf("invalid secret")
	}

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
