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

// MenuByIdRequest represents the request structure for menu by id endpoint
type MenuByIdRequest struct {
	// External menu ID
	ExternalMenuId string `json:"externalMenuId"`
	// Organization IDs
	OrganizationIds []uuid.UUID `json:"organizationIds"`
	// Price category ID
	PriceCategoryId string `json:"priceCategoryId"`
	// API version
	Version int `json:"version"`
	// Language code
	Language string `json:"language"`
	// Async mode flag
	AsyncMode bool `json:"asyncMode"`
	// Start revision number
	StartRevision int `json:"startRevision"`
}

// MenuPrice represents price information for organization
type MenuPrice struct {
	// Organization ID
	OrganizationId string `json:"organizationId"`
	// Price value
	Price float64 `json:"price"`
}

// AllergenGroup represents allergen group information
type AllergenGroup struct {
	// Allergen group ID
	Id uuid.UUID `json:"id"`
	// Allergen group code
	Code string `json:"code"`
	// Allergen group name
	Name string `json:"name"`
}

// Tag represents tag information
type Tag struct {
	// Tag ID
	Id string `json:"id"`
	// Tag name
	Name string `json:"name"`
}

// Restrictions represents quantity restrictions
type Restrictions struct {
	// Minimum quantity
	MinQuantity int `json:"minQuantity"`
	// Maximum quantity
	MaxQuantity int `json:"maxQuantity"`
	// Free quantity
	FreeQuantity int `json:"freeQuantity"`
	// Default quantity
	ByDefault int `json:"byDefault"`
}

// TaxCategory represents tax category information
type TaxCategory struct {
	// Tax category ID
	Id string `json:"id"`
	// Tax category name
	Name string `json:"name"`
	// Tax percentage
	Percentage float64 `json:"percentage"`
}

// ModifierItem represents modifier item
type ModifierItem struct {
	// Prices for different organizations
	Prices []MenuPrice `json:"prices"`
	// SKU
	Sku string `json:"sku"`
	// Name
	Name string `json:"name"`
	// Description
	Description string `json:"description"`
	// Button image URL
	ButtonImage string `json:"buttonImage"`
	// Quantity restrictions
	Restrictions Restrictions `json:"restrictions"`
	// Allergen groups
	AllergenGroups []AllergenGroup `json:"allergenGroups"`
	// Nutrition per hundred grams (keeping as interface{} for flexibility)
	NutritionPerHundredGrams interface{} `json:"nutritionPerHundredGrams"`
	// Portion weight in grams
	PortionWeightGrams int `json:"portionWeightGrams"`
	// Tags
	Tags []Tag `json:"tags"`
	// Item ID
	ItemId uuid.UUID `json:"itemId"`
}

// ItemModifierGroup represents item modifier group
type ItemModifierGroup struct {
	// Modifier items
	Items []ModifierItem `json:"items"`
	// Group name
	Name string `json:"name"`
	// Group description
	Description string `json:"description"`
	// Quantity restrictions
	Restrictions Restrictions `json:"restrictions"`
	// Can be divided flag
	CanBeDivided bool `json:"canBeDivided"`
	// Item group ID
	ItemGroupId uuid.UUID `json:"itemGroupId"`
	// Child modifiers have min max restrictions flag
	ChildModifiersHaveMinMaxRestrictions bool `json:"childModifiersHaveMinMaxRestrictions"`
	// SKU
	Sku string `json:"sku"`
}

// ItemSize represents item size information
type ItemSize struct {
	// Prices for different organizations
	Prices []MenuPrice `json:"prices"`
	// Item modifier groups
	ItemModifierGroups []ItemModifierGroup `json:"itemModifierGroups"`
	// SKU
	Sku string `json:"sku"`
	// Size code
	SizeCode string `json:"sizeCode"`
	// Size name
	SizeName string `json:"sizeName"`
	// Is default size flag
	IsDefault bool `json:"isDefault"`
	// Portion weight in grams
	PortionWeightGrams int `json:"portionWeightGrams"`
	// Size ID
	SizeId uuid.UUID `json:"sizeId"`
	// Nutrition per hundred grams (keeping as interface{} for flexibility)
	NutritionPerHundredGrams interface{} `json:"nutritionPerHundredGrams"`
	// Button image URL
	ButtonImageUrl string `json:"buttonImageUrl"`
	// Button image cropped URLs
	ButtonImageCroppedUrl []string `json:"buttonImageCroppedUrl"`
}

// MenuItem represents menu item
type MenuItem struct {
	// Item sizes
	ItemSizes []ItemSize `json:"itemSizes"`
	// SKU
	Sku string `json:"sku"`
	// Name
	Name string `json:"name"`
	// Description
	Description string `json:"description"`
	// Allergen groups
	AllergenGroups []AllergenGroup `json:"allergenGroups"`
	// Item ID
	ItemId uuid.UUID `json:"itemId"`
	// Modifier schema ID
	ModifierSchemaId uuid.UUID `json:"modifierSchemaId"`
	// Tax category
	TaxCategory TaxCategory `json:"taxCategory"`
	// Order item type
	OrderItemType string `json:"orderItemType"`
}

// ItemCategory represents item category
type ItemCategory struct {
	// Menu items
	Items []MenuItem `json:"items"`
	// Category ID
	Id uuid.UUID `json:"id"`
	// Category name
	Name string `json:"name"`
	// Category description
	Description string `json:"description"`
	// Button image URL
	ButtonImageUrl string `json:"buttonImageUrl"`
	// Header image URL
	HeaderImageUrl string `json:"headerImageUrl"`
}

// MenuByIdResponse represents the response structure for menu by id endpoint
type MenuByIdResponse struct {
	// Menu ID
	Id int `json:"id"`
	// Menu name
	Name string `json:"name"`
	// Menu description
	Description string `json:"description"`
	// Item categories
	ItemCategories []ItemCategory `json:"itemCategories"`
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

// MenuById Retrieve menu by external menu ID
//
// iiko API: /api/2/menu/by_id
func (c *Client) MenuById(req *MenuByIdRequest, opts ...Option) (*MenuByIdResponse, error) {
	var response MenuByIdResponse

	if err := c.post(true, "/api/2/menu/by_id", req, &response, opts...); err != nil {
		return nil, err
	}

	return &response, nil
}
