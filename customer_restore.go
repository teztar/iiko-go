package iiko

import "github.com/google/uuid"

// RestoreCustomersRequest represents the request structure for restoring customers
type RestoreCustomersRequest struct {
	// Array of customer IDs to restore
	CustomerIds []uuid.UUID `json:"customerIds"`
	// Organization ID where the customers should be restored
	OrganizationId uuid.UUID `json:"organizationId"`
}

// RestoreCustomersResponse represents the response structure for the restore customers operation
type RestoreCustomersResponse struct {
	// Total number of customers processed
	Total int `json:"total"`
	// Number of customers successfully restored
	Restored int `json:"restored"`
	// Number of customers not found
	NotFound int `json:"notFound"`
}

// RestoreCustomers Restore customers by their IDs
//
// iiko API: /api/1/loyalty/iiko/restore_customers
func (c *Client) RestoreCustomers(req *RestoreCustomersRequest, opts ...Option) (*RestoreCustomersResponse, error) {
	var response RestoreCustomersResponse

	if err := c.post(true, "/api/1/loyalty/iiko/restore_customers", req, &response, opts...); err != nil {
		return nil, err
	}

	return &response, nil
}
