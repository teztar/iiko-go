package iiko

import "github.com/google/uuid"

// DeleteCustomersRequest represents the request structure for deleting customers
type DeleteCustomersRequest struct {
	// Array of customer IDs to delete
	CustomerIds []uuid.UUID `json:"customerIds"`
	// Organization ID where the customers should be deleted
	OrganizationId uuid.UUID `json:"organizationId"`
}

// DeleteCustomersResponse represents the response structure for the delete customers operation
type DeleteCustomersResponse struct {
	// Total number of customers processed
	Total int `json:"total"`
	// Number of customers successfully deleted
	Deleted int `json:"deleted"`
	// Number of customers not found
	NotFound int `json:"notFound"`
}

// DeleteCustomers Delete customers by their IDs
//
// iiko API: /api/1/loyalty/iiko/delete_customers
func (c *Client) DeleteCustomers(req *DeleteCustomersRequest, opts ...Option) (*DeleteCustomersResponse, error) {
	var response DeleteCustomersResponse

	if err := c.post(true, "/api/1/loyalty/iiko/delete_customers", req, &response, opts...); err != nil {
		return nil, err
	}

	return &response, nil
}
