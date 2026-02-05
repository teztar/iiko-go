package iiko

import (
	"github.com/google/uuid"
)

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
	Version *int `json:"version"`
	// Language code
	Language *string `json:"language"`
	// Start revision number
	StartRevision *int `json:"startRevision"`
}

// MenuPrice represents price information for organization
type MenuPrice struct {
	// Organization ID
	OrganizationId string `json:"organizationId"`
	// Price value
	Price float64 `json:"price"`
}

// MenuPriceWithTax represents price information with organizations and tax category for menu
type MenuPriceWithTax struct {
	// Organization IDs
	Organizations []string `json:"organizations"`
	// Price value
	Price string `json:"price"`
	// Tax category ID
	TaxCategoryId uuid.UUID `json:"taxCategoryId"`
}

// AllergenGroup represents allergen group information
type AllergenGroup struct {
	// Allergen group ID
	Id uuid.UUID `json:"id"`
	// Allergen group code
	Code string `json:"code"`
	// Allergen group name
	Name string `json:"name"`
	// Is deleted flag
	IsDeleted interface{} `json:"isDeleted"` // Can be bool or string "false"
}

// Tag represents tag information
type Tag struct {
	// Tag ID
	Id string `json:"id"`
	// Tag name
	Name string `json:"name"`
}

// Allergen represents allergen information
type Allergen struct {
	// Allergen ID
	Id uuid.UUID `json:"id"`
	// Allergen code
	Code string `json:"code"`
	// Allergen name
	Name string `json:"name"`
	// Is deleted flag
	IsDeleted interface{} `json:"isDeleted"` // Can be bool or string "false"
}

// Label represents label information
type Label struct {
	// Label code
	Code string `json:"code"`
	// Label name
	Name string `json:"name"`
}

// MenuProductCategory represents product category information for menu
type MenuProductCategory struct {
	// Category ID
	Id string `json:"id"`
	// Category name
	Name string `json:"name"`
	// Is deleted flag
	IsDeleted bool `json:"isDeleted"`
}

// CustomerTag represents customer tag information
type CustomerTag struct {
	// Tag ID
	Id string `json:"id"`
	// Tag name
	Name string `json:"name"`
}

// CustomerTagGroup represents customer tag group information
type CustomerTagGroup struct {
	// Group ID
	Id string `json:"id"`
	// Group name
	Name string `json:"name"`
	// Select several tags flag
	SelectSeveralTags bool `json:"selectSeveralTags"`
	// Tag items
	Items []CustomerTag `json:"items"`
}

// CustomerTagGroupSelection represents customer tag group selection
type CustomerTagGroupSelection struct {
	// Customer tag group ID
	CustomerTagGroupId uuid.UUID `json:"customerTagGroupId"`
	// Selected tag IDs
	SelectedTagIds []*string `json:"selectedTagIds"` // Can contain null
}

// Interval represents time interval for organization
type Interval struct {
	// Organization ID
	OrganizationId uuid.UUID `json:"organizationId"`
	// From time
	FromTime string `json:"fromTime"`
	// To time
	ToTime string `json:"toTime"`
}

// Schedule represents schedule information
type Schedule struct {
	// Begin time
	Begin string `json:"begin"`
	// End time
	End string `json:"end"`
	// Days of week
	DaysOfWeek []string `json:"daysOfWeek"`
}

// Nutrition represents nutrition information
type Nutrition struct {
	// Fats
	Fats float64 `json:"fats"`
	// Proteins
	Proteins float64 `json:"proteins"`
	// Carbs
	Carbs float64 `json:"carbs"`
	// Energy
	Energy float64 `json:"energy"`
	// Organizations
	Organizations []string `json:"organizations"`
	// Saturated fatty acid
	SaturatedFattyAcid float64 `json:"saturatedFattyAcid"`
	// Salt
	Salt float64 `json:"salt"`
	// Sugar
	Sugar float64 `json:"sugar"`
}

// Barcode represents barcode information
type Barcode struct {
	// Barcode value
	Barcode string `json:"barcode"`
	// Container
	Container string `json:"container"`
}

// ComboImage represents combo image information
type ComboImage struct {
	// Image URL
	Url string `json:"url"`
	// Image hash
	Hash string `json:"hash"`
}

// ComboSize represents combo size information
type ComboSize struct {
	// Size name
	Name string `json:"name"`
	// Size ID
	Id string `json:"id"`
	// Button image
	ButtonImage ComboImage `json:"buttonImage"`
	// Short name
	ShortName string `json:"shortName"`
}

// ComboItemSize represents combo item size information
type ComboItemSize struct {
	// Combo size ID
	ComboSizeId uuid.UUID `json:"comboSizeId"`
	// Size ID
	SizeId uuid.UUID `json:"sizeId"`
	// Name
	Name string `json:"name"`
	// Short name
	ShortName string `json:"shortName"`
	// Prices
	Prices []MenuPriceWithTax `json:"prices"`
}

// ComboGroupItem represents combo group item
type ComboGroupItem struct {
	// Item ID
	ItemId string `json:"itemId"`
	// Forbidden modifiers
	ForbiddenModifiers []*string `json:"forbiddenModifiers"` // Can contain null
	// Size ID
	SizeId string `json:"sizeId"`
	// Price modification amount
	PriceModificationAmount float64 `json:"priceModificationAmount"`
	// Sizes
	Sizes []ComboItemSize `json:"sizes"`
}

// ComboGroup represents combo group
type ComboGroup struct {
	// Group ID
	Id *string `json:"id"` // Can be null
	// Group name
	Name string `json:"name"`
	// Is main group flag
	IsMainGroup bool `json:"isMainGroup"`
	// Items
	Items []ComboGroupItem `json:"items"`
	// Skip step flag
	SkipStep bool `json:"skipStep"`
}

// MenuCombo represents combo information for menu
type MenuCombo struct {
	// Combo name
	Name string `json:"name"`
	// Combo price
	Price string `json:"price"`
	// Groups
	Groups []ComboGroup `json:"groups"`
	// Images
	Image []ComboImage `json:"image"`
	// Description
	Description string `json:"description"`
	// Sizes
	Sizes []ComboSize `json:"sizes"`
	// Price strategy
	PriceStrategy string `json:"priceStrategy"`
	// Start date
	StartDate string `json:"startDate"`
	// Expiration date
	ExpirationDate string `json:"expirationDate"`
	// Combo ID
	Id uuid.UUID `json:"id"`
}

// MenuComboCategory represents combo category for menu
type MenuComboCategory struct {
	// Category ID
	Id *string `json:"id"` // Can be null
	// Category name
	Name string `json:"name"`
	// Combos
	Combos []MenuCombo `json:"combos"`
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
	ByDefault interface{} `json:"byDefault"` // Can be int or string "0"
	// Hide if default quantity flag
	HideIfDefaultQuantity bool `json:"hideIfDefaultQuantity"`
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
	// SKU
	Sku string `json:"sku"`
	// Name
	Name string `json:"name"`
	// Description
	Description string `json:"description"`
	// Quantity restrictions (array in JSON)
	Restrictions []Restrictions `json:"restrictions"`
	// Allergen groups
	AllergenGroups []AllergenGroup `json:"allergenGroups"`
	// Nutrition per hundred grams
	NutritionPerHundredGrams Nutrition `json:"nutritionPerHundredGrams"`
	// Portion weight in grams (string in JSON)
	PortionWeightGrams float64 `json:"portionWeightGrams"`
	// Tags
	Tags []Tag `json:"tags"`
	// Labels
	Labels []Label `json:"labels"`
	// Item ID (string in JSON)
	ItemId string `json:"itemId"`
	// Is hidden flag
	IsHidden bool `json:"isHidden"`
	// Prices
	Prices []MenuPriceWithTax `json:"prices"`
	// Position
	Position int `json:"position"`
	// Independent quantity flag
	IndependentQuantity bool `json:"independentQuantity"`
	// Product category ID
	ProductCategoryId string `json:"productCategoryId"`
	// Customer tag groups
	CustomerTagGroups []CustomerTagGroupSelection `json:"customerTagGroups"`
	// Payment subject
	PaymentSubject string `json:"paymentSubject"`
	// Outer EAN code
	OuterEanCode string `json:"outerEanCode"`
	// Is marked flag
	IsMarked bool `json:"isMarked"`
	// Measure unit type
	MeasureUnitType string `json:"measureUnitType"`
	// Payment subject code
	PaymentSubjectCode string `json:"paymentSubjectCode"`
	// Barcodes
	Barcodes []Barcode `json:"barcodes"`
	// Button image URL
	ButtonImageUrl string `json:"buttonImageUrl"`
}

// ItemModifierGroup represents item modifier group
type ItemModifierGroup struct {
	// Group name
	Name string `json:"name"`
	// Group description
	Description string `json:"description"`
	// Quantity restrictions
	Restrictions Restrictions `json:"restrictions"`
	// Modifier items
	Items []ModifierItem `json:"items"`
	// Can be divided flag (string in JSON)
	CanBeDivided interface{} `json:"canBeDivided"` // Can be bool or string "false"
	// Item group ID
	ItemGroupId uuid.UUID `json:"itemGroupId"`
	// Is hidden flag
	IsHidden bool `json:"isHidden"`
	// Child modifiers have min max restrictions flag
	ChildModifiersHaveMinMaxRestrictions bool `json:"childModifiersHaveMinMaxRestrictions"`
	// SKU
	Sku string `json:"sku"`
}

// ItemSize represents item size information
type ItemSize struct {
	// SKU
	Sku string `json:"sku"`
	// Size code
	SizeCode string `json:"sizeCode"`
	// Size name
	SizeName string `json:"sizeName"`
	// Is default size flag
	IsDefault bool `json:"isDefault"`
	// Portion weight in grams (string in JSON)
	PortionWeightGrams float64 `json:"portionWeightGrams"`
	// Item modifier groups
	ItemModifierGroups []ItemModifierGroup `json:"itemModifierGroups"`
	// Size ID
	SizeId uuid.UUID `json:"sizeId"`
	// Nutrition per hundred grams
	NutritionPerHundredGrams Nutrition `json:"nutritionPerHundredGrams"`
	// Prices
	Prices []MenuPriceWithTax `json:"prices"`
	// Nutritions
	Nutritions []Nutrition `json:"nutritions"`
	// Is hidden flag
	IsHidden bool `json:"isHidden"`
	// Measure unit type
	MeasureUnitType string `json:"measureUnitType"`
	// Button image URL
	ButtonImageUrl string `json:"buttonImageUrl"`
}

// MenuItem represents menu item
type MenuItem struct {
	// SKU
	Sku string `json:"sku"`
	// Name
	Name string `json:"name"`
	// Description
	Description string `json:"description"`
	// Allergens
	Allergens []Allergen `json:"allergens"`
	// Tags
	Tags []Tag `json:"tags"`
	// Labels
	Labels []Label `json:"labels"`
	// Item sizes
	ItemSizes []ItemSize `json:"itemSizes"`
	// Item ID
	ItemId uuid.UUID `json:"itemId"`
	// Modifier schema ID
	ModifierSchemaId uuid.UUID `json:"modifierSchemaId"`
	// Tax category (array in JSON)
	TaxCategory []TaxCategory `json:"taxCategory"`
	// Modifier schema name
	ModifierSchemaName string `json:"modifierSchemaName"`
	// Type
	Type string `json:"type"`
	// Can be divided flag (string in JSON)
	CanBeDivided interface{} `json:"canBeDivided"` // Can be bool or string "false"
	// Can set open price flag
	CanSetOpenPrice bool `json:"canSetOpenPrice"`
	// Use balance for sell flag
	UseBalanceForSell bool `json:"useBalanceForSell"`
	// Measure unit
	MeasureUnit string `json:"measureUnit"`
	// Product category ID
	ProductCategoryId uuid.UUID `json:"productCategoryId"`
	// Customer tag groups
	CustomerTagGroups []CustomerTagGroupSelection `json:"customerTagGroups"`
	// Payment subject
	PaymentSubject string `json:"paymentSubject"`
	// Payment subject code
	PaymentSubjectCode string `json:"paymentSubjectCode"`
	// Outer EAN code
	OuterEanCode string `json:"outerEanCode"`
	// Is marked flag
	IsMarked bool `json:"isMarked"`
	// Is hidden flag
	IsHidden bool `json:"isHidden"`
	// Barcodes
	Barcodes []Barcode `json:"barcodes"`
	// Order item type
	OrderItemType string `json:"orderItemType"`
}

// ItemCategory represents item category
type ItemCategory struct {
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
	// iiko group ID
	IikoGroupId uuid.UUID `json:"iikoGroupId"`
	// Menu items
	Items []MenuItem `json:"items"`
	// Schedule ID
	ScheduleId uuid.UUID `json:"scheduleId"`
	// Schedule name
	ScheduleName string `json:"scheduleName"`
	// Schedules
	Schedules []Schedule `json:"schedules"`
	// Is hidden flag
	IsHidden bool `json:"isHidden"`
	// Tags
	Tags []Tag `json:"tags"`
	// Labels
	Labels []Label `json:"labels"`
}

// MenuByIdResponse represents the response structure for menu by id endpoint
type MenuByIdResponse struct {
	// Product categories
	ProductCategories []MenuProductCategory `json:"productCategories"`
	// Customer tag groups
	CustomerTagGroups []CustomerTagGroup `json:"customerTagGroups"`
	// Revision
	Revision int `json:"revision"`
	// Format version
	FormatVersion int `json:"formatVersion"`
	// Menu ID
	Id int `json:"id"`
	// Menu name
	Name string `json:"name"`
	// Menu description
	Description string `json:"description"`
	// Button image URL
	ButtonImageUrl string `json:"buttonImageUrl"`
	// Intervals
	Intervals []Interval `json:"intervals"`
	// Item categories
	ItemCategories []ItemCategory `json:"itemCategories"`
	// Combo categories
	ComboCategories []MenuComboCategory `json:"comboCategories"`
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
