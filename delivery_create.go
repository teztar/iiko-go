package iiko

import (
	"github.com/google/uuid"
)

type DeliveryCreateRequest struct {
	// Organization ID
	OrganizationId uuid.UUID `json:"organizationId"`
	// Terminal group ID
	TerminalGroupId uuid.UUID `json:"terminalGroupId"`
	// Create order settings
	CreateOrderSettings CreateOrderSettings `json:"createOrderSettings"`
	// Order information
	Order CreateDeliveryOrderSettings `json:"order"`
}

type CreateDeliveryOrderSettings struct {
	// Menu ID (nullable)
	MenuId *string `json:"menuId"`
	// Price category ID
	PriceCategoryId *string `json:"priceCategoryId"`
	// Order ID
	Id *uuid.UUID `json:"id"`
	// External number
	ExternalNumber *string `json:"externalNumber"`
	// Complete before time
	CompleteBefore *string `json:"completeBefore"`
	// Phone number
	Phone string `json:"phone"`
	// Phone extension
	PhoneExtension *string `json:"phoneExtension"`
	// Order type ID
	OrderTypeId *uuid.UUID `json:"orderTypeId"`
	// Order service type
	OrderServiceType *string `json:"orderServiceType"`
	// Delivery point
	DeliveryPoint *DeliveryOrderPoint `json:"deliveryPoint"`
	// Order comment
	Comment *string `json:"comment"`
	// Customer information
	Customer *Customer `json:"customer"`
	// Guests information
	Guests *Guests `json:"guests"`
	// Marketing source ID
	MarketingSourceId *uuid.UUID `json:"marketingSourceId"`
	// Operator ID
	OperatorId *uuid.UUID `json:"operatorId"`
	// Delivery duration in minutes
	DeliveryDuration *int `json:"deliveryDuration"`
	// Delivery zone
	DeliveryZone *string `json:"deliveryZone"`
	// Order items
	Items []DeliveryOrderItem `json:"items"`
	// Order combos
	Combos []DeliveryOrderCombo `json:"combos"`
	// Order payments
	Payments []DeliveryOrderPayment `json:"payments"`
	// Order tips
	Tips []DeliveryOrderTip `json:"tips"`
	// Source key
	SourceKey *string `json:"sourceKey"`
	// Discounts information
	DiscountsInfo *DeliveryDiscountsInfo `json:"discountsInfo"`
	// Loyalty information
	LoyaltyInfo *DeliveryLoyaltyInfo `json:"loyaltyInfo"`
	// Cheque additional information
	ChequeAdditionalInfo *DeliveryChequeAdditionalInfo `json:"chequeAdditionalInfo"`
	// External data
	ExternalData []DeliveryExternalData `json:"externalData"`
}

type DeliveryCreateResponse struct {
	// Correlation ID
	CorrelationId uuid.UUID `json:"correlationId"`
	// Order information
	OrderInfo DeliveryOrderInfo `json:"orderInfo"`
}

// Common delivery types

// DeliveryOrderPoint represents delivery point for order
type DeliveryOrderPoint struct {
	// Coordinates
	Coordinates *DeliveryCoordinates `json:"coordinates"`
	// Address
	Address *DeliveryAddress `json:"address"`
	// External cartography ID
	ExternalCartographyId *string `json:"externalCartographyId"`
	// Comment
	Comment *string `json:"comment"`
}

// DeliveryCoordinates represents geographical coordinates
type DeliveryCoordinates struct {
	// Latitude
	Latitude float64 `json:"latitude"`
	// Longitude
	Longitude float64 `json:"longitude"`
}

// DeliveryAddress represents delivery address
type DeliveryAddress struct {
	// Address type
	Type string `json:"type"`
	Line1 string `json:"line1"`
	Street *DeliveryStreet `json:"street"`
	House string `json:"house"`
	Building *string `json:"building"`
	Flat *string `json:"flat"`
	Entrance *string `json:"entrance"`
	Floor *string `json:"floor"`
	Doorphone *string `json:"doorphone"`
	RegionID *string `json:"regionId"`
}

// DeliveryStreet represents street information
type DeliveryStreet struct {
	// ID
	ID *uuid.UUID `json:"id"`
	// Name
	Name *string `json:"name"`
	// City
	City interface{} `json:"city"`
}

// DeliveryCity represents city information
type DeliveryCity struct {
	// ID
	ID uuid.UUID `json:"id"`
	// Name
	Name string `json:"name"`
}

// DeliveryOrderItem represents an order item for delivery
type DeliveryOrderItem struct {
	ProductID string `json:"productId"`
	Modifiers []*Modifier `json:"modifiers"`
	Price float64 `json:"price"`
	PositionID *string `json:"positionId"`
	// Item type
	Type string `json:"type"`
	// Item amount
	Amount float64 `json:"amount"`
	// Product size ID
	ProductSizeId *uuid.UUID `json:"productSizeId"`
	// Combo information
	ComboInformation *ComboInformation `json:"comboInformation"`
	// Item comment
	Comment *string `json:"comment"`
}

// DeliveryOrderCombo represents order combo for delivery
type DeliveryOrderCombo struct {
	// Combo ID
	Id uuid.UUID `json:"id"`
	// Combo name
	Name string `json:"name"`
	// Combo amount
	Amount int `json:"amount"`
	// Combo price
	Price float64 `json:"price"`
	// Source ID
	SourceId uuid.UUID `json:"sourceId"`
	// Program ID
	ProgramId uuid.UUID `json:"programId"`
	// Size ID
	SizeId uuid.UUID `json:"sizeId"`
}

// DeliveryOrderPayment represents order payment for delivery
type DeliveryOrderPayment struct {
	// Payment type kind
	PaymentTypeKind string `json:"paymentTypeKind"`
	// Payment sum
	Sum float64 `json:"sum"`
	// Payment type ID
	PaymentTypeId uuid.UUID `json:"paymentTypeId"`
	// Is processed externally flag
	IsProcessedExternally bool `json:"isProcessedExternally"`
	// Payment additional data
	PaymentAdditionalData PaymentAdditionalData `json:"paymentAdditionalData"`
	// Is fiscalized externally flag
	IsFiscalizedExternally bool `json:"isFiscalizedExternally"`
	// Is prepay flag
	IsPrepay bool `json:"isPrepay"`
}

// DeliveryOrderTip represents order tip for delivery
type DeliveryOrderTip struct {
	// Payment type kind
	PaymentTypeKind string `json:"paymentTypeKind"`
	// Tips type ID
	TipsTypeId uuid.UUID `json:"tipsTypeId"`
	// Tip sum
	Sum float64 `json:"sum"`
	// Payment type ID
	PaymentTypeId uuid.UUID `json:"paymentTypeId"`
	// Is processed externally flag
	IsProcessedExternally bool `json:"isProcessedExternally"`
	// Payment additional data
	PaymentAdditionalData PaymentAdditionalData `json:"paymentAdditionalData"`
	// Is fiscalized externally flag
	IsFiscalizedExternally bool `json:"isFiscalizedExternally"`
	// Is prepay flag
	IsPrepay bool `json:"isPrepay"`
}

// DeliveryLoyaltyCard represents loyalty card for delivery
type DeliveryLoyaltyCard struct {
	// Card track
	Track string `json:"track"`
}

// DeliveryDiscount represents delivery discount
type DeliveryDiscount struct {
	// Discount type
	DiscountType *DeliveryDiscountType `json:"discountType,omitempty"`
	// Sum
	Sum float64 `json:"sum"`
	// Selective positions
	SelectivePositions []uuid.UUID `json:"selectivePositions"`
	// Selective positions with sum
	SelectivePositionsWithSum []DeliverySelectivePositionWithSum `json:"selectivePositionsWithSum"`
}

// DeliveryDiscountType represents delivery discount type
type DeliveryDiscountType struct {
	// ID
	ID uuid.UUID `json:"id"`
	// Name
	Name string `json:"name"`
}

// DeliverySelectivePositionWithSum represents selective position with sum
type DeliverySelectivePositionWithSum struct {
	// Position ID
	PositionID uuid.UUID `json:"positionId"`
	// Sum
	Sum float64 `json:"sum"`
}

// DeliveryDiscountsInfo represents discounts information for delivery
type DeliveryDiscountsInfo struct {
	// Loyalty card
	Card DeliveryLoyaltyCard `json:"card"`
	// Applied discounts
	Discounts []DeliveryDiscount `json:"discounts"`
	// Fixed loyalty discounts flag
	FixedLoyaltyDiscounts bool `json:"fixedLoyaltyDiscounts"`
}

// DeliveryLoyaltyInfo represents loyalty information for delivery
type DeliveryLoyaltyInfo struct {
	// Coupon code
	Coupon string `json:"coupon"`
	// Applicable manual conditions
	ApplicableManualConditions []uuid.UUID `json:"applicableManualConditions"`
}

// DeliveryChequeAdditionalInfo represents additional cheque information
type DeliveryChequeAdditionalInfo struct {
	// Need receipt flag
	NeedReceipt bool `json:"needReceipt"`
	// Email for receipt
	Email string `json:"email"`
	// Settlement place
	SettlementPlace string `json:"settlementPlace"`
	// Phone number
	Phone string `json:"phone"`
}

// DeliveryExternalData represents external data for delivery
type DeliveryExternalData struct {
	// Data key
	Key string `json:"key"`
	// Data value
	Value string `json:"value"`
	// Is public flag
	IsPublic bool `json:"isPublic"`
}

// DeliveryCreate Create a new delivery order
//
// iiko API: /api/1/deliveries/create
func (c *Client) DeliveryCreate(req *DeliveryCreateRequest, opts ...Option) (*DeliveryCreateResponse, error) {
	var response DeliveryCreateResponse

	if err := c.post(true, "/api/1/deliveries/create", req, &response, opts...); err != nil {
		return nil, err
	}

	return &response, nil
}
