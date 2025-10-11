package iiko

// CustomerCategoriesRequest represents request body for getting customer categories
type CustomerCategoriesRequest struct {
	OrganizationId string `json:"organizationId"`
}

// CustomerCategoriesResponse represents response for getting customer categories
type CustomerCategoriesResponse struct {
	GuestCategories []Category `json:"guestCategories"`
}

// CustomerCategoryAddRequest represents request body for adding category to customer
type CustomerCategoryAddRequest struct {
	CustomerId     string `json:"customerId"`
	CategoryId     string `json:"categoryId"`
	OrganizationId string `json:"organizationId"`
}

// CustomerCategoryRemoveRequest represents request body for removing category from customer
type CustomerCategoryRemoveRequest struct {
	CustomerId     string `json:"customerId"`
	CategoryId     string `json:"categoryId"`
	OrganizationId string `json:"organizationId"`
}

// CustomerCategories gets customer categories for organization
//
// iiko API: POST /api/1/loyalty/iiko/customer_category
func (c *Client) CustomerCategories(req *CustomerCategoriesRequest, opts ...Option) (*CustomerCategoriesResponse, error) {
	var response CustomerCategoriesResponse

	if err := c.post(true, "/api/1/loyalty/iiko/customer_category", req, &response, opts...); err != nil {
		return nil, err
	}

	return &response, nil
}

// CustomerCategoryAdd adds category to customer
//
// iiko API: POST /api/1/loyalty/iiko/customer_category/add
func (c *Client) CustomerCategoryAdd(req *CustomerCategoryAddRequest, opts ...Option) error {
	// Create empty response since API returns no response body
	var response interface{}

	if err := c.post(true, "/api/1/loyalty/iiko/customer_category/add", req, &response, opts...); err != nil {
		return err
	}

	return nil
}

// CustomerCategoryRemove removes category from customer
//
// iiko API: POST /api/1/loyalty/iiko/customer_category/remove
func (c *Client) CustomerCategoryRemove(req *CustomerCategoryRemoveRequest, opts ...Option) error {
	// Create empty response since API returns no response body
	var response interface{}

	if err := c.post(true, "/api/1/loyalty/iiko/customer_category/remove", req, &response, opts...); err != nil {
		return err
	}

	return nil
}
