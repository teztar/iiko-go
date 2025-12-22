package iiko

import (
	"github.com/google/uuid"
)

type UpdateOrderDeliveryStatusRequest struct {
	// Organization ID
	// Can be obtained by /api/1/organizations operation. [required]
	OrganizationId uuid.UUID `json:"organizationId"`
	// Order ID [required]
	OrderId uuid.UUID `json:"orderId"`
	// Delivery date [required]
	DeliveryDate string `json:"deliveryDate"`
	// Delivery status [required]
	Status DeliveryStatus `json:"status"`
}

type UpdateOrderDeliveryStatusResponse struct {
	// Operation ID [required]
	CorrelationId uuid.UUID `json:"correlationId"`
}

// UpdateOrderDeliveryStatus Update order delivery status
//
// iiko API: /api/1/deliveries/update_order_delivery_status
func (c *Client) UpdateOrderDeliveryStatus(req *UpdateOrderDeliveryStatusRequest, opts ...Option) (*UpdateOrderDeliveryStatusResponse, error) {
	var response UpdateOrderDeliveryStatusResponse

	if err := c.post(true, "/api/1/deliveries/update_order_delivery_status", req, &response, opts...); err != nil {
		return nil, err
	}

	return &response, nil
}
