package iiko

import "github.com/google/uuid"

// ExternalMenu represents external menu information
type ExternalMenu struct {
	// External menu ID
	Id string `json:"id"`
	// External menu name
	Name string `json:"name"`
}

// PriceCategory represents price category information
type PriceCategory struct {
	// Price category ID
	Id string `json:"id"`
	// Price category name
	Name string `json:"name"`
}

// MenuResponse represents the response structure for the menu endpoint
type MenuResponse struct {
	// Operation correlation ID
	CorrelationId uuid.UUID `json:"correlationId"`
	// List of external menus
	ExternalMenus []ExternalMenu `json:"externalMenus"`
	// List of price categories
	PriceCategories []PriceCategory `json:"priceCategories"`
}

// Menu Retrieve external menus and price categories
//
// iiko API: /api/2/menu
func (c *Client) Menu(opts ...Option) (*MenuResponse, error) {
	var response MenuResponse

	if err := c.post(true, "/api/2/menu", nil, &response, opts...); err != nil {
		return nil, err
	}

	return &response, nil
}
