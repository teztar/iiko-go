package iiko

import "github.com/google/uuid"

// WebhookSettingsRequest represents the request structure for getting webhook settings
type WebhookSettingsRequest struct {
	// Organization ID
	OrganizationId uuid.UUID `json:"organizationId"`
}

// DeliveryOrderFilter represents delivery order filter configuration
type DeliveryOrderFilter struct {
	// Order statuses to filter
	OrderStatuses []string `json:"orderStatuses"`
	// Item statuses to filter
	ItemStatuses []string `json:"itemStatuses"`
	// Include errors flag
	Errors bool `json:"errors"`
	// Returned external data keys
	ReturnedExternalDataKeys []string `json:"returnedExternalDataKeys"`
}

// TableOrderFilter represents table order filter configuration
type TableOrderFilter struct {
	// Order statuses to filter
	OrderStatuses []string `json:"orderStatuses"`
	// Item statuses to filter
	ItemStatuses []string `json:"itemStatuses"`
	// Include errors flag
	Errors bool `json:"errors"`
}

// ReserveFilter represents reserve filter configuration
type ReserveFilter struct {
	// Include updates flag
	Updates bool `json:"updates"`
	// Include errors flag
	Errors bool `json:"errors"`
}

// StopListUpdateFilter represents stop list update filter configuration
type StopListUpdateFilter struct {
	// Include updates flag
	Updates bool `json:"updates"`
}

// PersonalShiftFilter represents personal shift filter configuration
type PersonalShiftFilter struct {
	// Include updates flag
	Updates bool `json:"updates"`
}

// NomenclatureUpdateFilter represents nomenclature update filter configuration
type NomenclatureUpdateFilter struct {
	// Include updates flag
	Updates bool `json:"updates"`
}

// BusinessHoursAndMappingUpdateFilter represents business hours and mapping update filter configuration
type BusinessHoursAndMappingUpdateFilter struct {
	// Include updates flag
	Updates bool `json:"updates"`
}

// WebHooksFilter represents webhook filter configuration
type WebHooksFilter struct {
	// Delivery order filter
	DeliveryOrderFilter *DeliveryOrderFilter `json:"deliveryOrderFilter,omitempty"`
	// Table order filter
	TableOrderFilter *TableOrderFilter `json:"tableOrderFilter,omitempty"`
	// Reserve filter
	ReserveFilter *ReserveFilter `json:"reserveFilter,omitempty"`
	// Stop list update filter
	StopListUpdateFilter *StopListUpdateFilter `json:"stopListUpdateFilter,omitempty"`
	// Personal shift filter
	PersonalShiftFilter *PersonalShiftFilter `json:"personalShiftFilter,omitempty"`
	// Nomenclature update filter
	NomenclatureUpdateFilter *NomenclatureUpdateFilter `json:"nomenclatureUpdateFilter,omitempty"`
	// Business hours and mapping update filter
	BusinessHoursAndMappingUpdateFilter *BusinessHoursAndMappingUpdateFilter `json:"businessHoursAndMappingUpdateFilter,omitempty"`
}

// WebhookSettingsResponse represents the response structure for webhook settings
type WebhookSettingsResponse struct {
	// Correlation ID
	CorrelationId uuid.UUID `json:"correlationId"`
	// API login name
	ApiLoginName string `json:"apiLoginName"`
	// Webhook URI
	WebHooksUri string `json:"webHooksUri"`
	// Auth token
	AuthToken string `json:"authToken"`
	// Webhook filter configuration
	WebHooksFilter WebHooksFilter `json:"webHooksFilter"`
}

// WebhookUpdateSettingsRequest represents the request structure for updating webhook settings
type WebhookUpdateSettingsRequest struct {
	// Organization ID
	OrganizationId uuid.UUID `json:"organizationId"`
	// Webhook URI
	WebHooksUri string `json:"webHooksUri"`
	// Auth token
	AuthToken string `json:"authToken"`
	// Webhook filter configuration
	WebHooksFilter WebHooksFilter `json:"webHooksFilter"`
}

// WebhookUpdateSettingsResponse represents the response structure for updating webhook settings
type WebhookUpdateSettingsResponse struct {
	// Correlation ID
	CorrelationId uuid.UUID `json:"correlationId"`
}

// WebhookSettings Retrieve webhook settings for organization
//
// iiko API: /api/1/webhooks/settings
func (c *Client) WebhookSettings(req *WebhookSettingsRequest, opts ...Option) (*WebhookSettingsResponse, error) {
	var response WebhookSettingsResponse

	if err := c.post(true, "/api/1/webhooks/settings", req, &response, opts...); err != nil {
		return nil, err
	}

	return &response, nil
}

// WebhookUpdateSettings Update webhook settings for organization
//
// iiko API: /api/1/webhooks/update_settings
func (c *Client) WebhookUpdateSettings(req *WebhookUpdateSettingsRequest, opts ...Option) (*WebhookUpdateSettingsResponse, error) {
	var response WebhookUpdateSettingsResponse

	if err := c.post(true, "/api/1/webhooks/update_settings", req, &response, opts...); err != nil {
		return nil, err
	}

	return &response, nil
}
