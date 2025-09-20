package iiko

import (
	"github.com/google/uuid"
)

type DeliveriesByIDRequest struct {
	// Organization ID for which an order search will be performed.
	// Can be obtained by /api/1/organizations operation.
	OrganizationID uuid.UUID `json:"organizationId"`
	// IDs of orders information on which is required.
	// Required if "posOrderIds" is null. Must be null if "posOrderIds" is not null.
	// Maximum allowed "orderIds" to request - 200.
	// The guaranteed order availability period is the last 7 days.
	// To access earlier orders, use the /api/1/deliveries/history/by_delivery_date_and_phone method.
	OrderIDs []uuid.UUID `json:"orderIds,omitempty"`
	// Source keys.
	SourceKeys []string `json:"sourceKeys,omitempty"`
	// POS order IDs information on which is required.
	// Required if "orderIds" is null. Must be null if "orderIds" is not null.
	// Maximum allowed "posOrderIds" to request - 200.
	// The guaranteed order availability period is the last 7 days.
	// To access earlier orders, use the /api/1/deliveries/history/by_delivery_date_and_phone method.
	PosOrderIDs []uuid.UUID `json:"posOrderIds,omitempty"`
	// Keys for return external data information.
	ReturnExternalDataKeys []string `json:"returnExternalDataKeys,omitempty"`
	// Whether to check and return LockedByUser property (see FullOrderUpdateRequest.EmployeeId).
	ReturnLockedByUser bool `json:"returnLockedByUser"`
}

type DeliveriesByIDResponse struct {
	// Orders information.
	Orders []DeliveryOrderInfo `json:"orders"`
}

// DeliveryOrderInfo represents delivery order information
type DeliveryOrderInfo struct {
	// Order ID.
	ID uuid.UUID `json:"id"`
	// POS order ID.
	PosID uuid.UUID `json:"posId"`
	// Order external number.
	ExternalNumber string `json:"externalNumber"`
	// Organization ID.
	OrganizationID uuid.UUID `json:"organizationId"`
	// Timestamp of most recent order change that took place on iikoTransport server.
	Timestamp int `json:"timestamp"`
	// Order creation status.
	CreationStatus OrderCreationStatus `json:"creationStatus"`
	// Order creation error details.
	ErrorInfo ErrorInfo `json:"errorInfo"`
	// Delivery order details.
	Order DeliveryOrder `json:"order"`
}

// DeliveryOrder represents delivery order details
type DeliveryOrder struct {
	// Parent delivery ID.
	ParentDeliveryID *uuid.UUID `json:"parentDeliveryId,omitempty"`
	// Customer information.
	Customer DeliveryCustomer `json:"customer"`
	// Guest phone.
	Phone string `json:"phone"`
	// Phone extension.
	PhoneExtension string `json:"phoneExtension"`
	// Delivery point.
	DeliveryPoint DeliveryOrderPoint `json:"deliveryPoint"`
	// Delivery status.
	Status DeliveryStatus `json:"status"`
	// Cancel information.
	CancelInfo *DeliveryCancelInfo `json:"cancelInfo,omitempty"`
	// Courier information.
	CourierInfo *DeliveryCourierInfo `json:"courierInfo,omitempty"`
	// Complete before time.
	CompleteBefore *EventTime `json:"completeBefore,omitempty"`
	// Order creation date.
	WhenCreated EventTime `json:"whenCreated"`
	// When confirmed.
	WhenConfirmed *EventTime `json:"whenConfirmed,omitempty"`
	// When printed.
	WhenPrinted *EventTime `json:"whenPrinted,omitempty"`
	// When cooking completed.
	WhenCookingCompleted *EventTime `json:"whenCookingCompleted,omitempty"`
	// When sent.
	WhenSended *EventTime `json:"whenSended,omitempty"`
	// When delivered.
	WhenDelivered *EventTime `json:"whenDelivered,omitempty"`
	// Comment.
	Comment string `json:"comment"`
	// Problem information.
	Problem *DeliveryProblem `json:"problem,omitempty"`
	// Operator information.
	Operator *DeliveryOperator `json:"operator,omitempty"`
	// Marketing source.
	MarketingSource *DeliveryMarketingSource `json:"marketingSource,omitempty"`
	// Delivery duration in minutes.
	DeliveryDuration int `json:"deliveryDuration"`
	// Index in courier route.
	IndexInCourierRoute int `json:"indexInCourierRoute"`
	// Cooking start time.
	CookingStartTime *EventTime `json:"cookingStartTime,omitempty"`
	// Is deleted flag.
	IsDeleted bool `json:"isDeleted"`
	// When received by API.
	WhenReceivedByApi *EventTime `json:"whenReceivedByApi,omitempty"`
	// When received from front.
	WhenReceivedFromFront *EventTime `json:"whenReceivedFromFront,omitempty"`
	// Moved from delivery ID.
	MovedFromDeliveryID *uuid.UUID `json:"movedFromDeliveryId,omitempty"`
	// Moved from terminal group ID.
	MovedFromTerminalGroupID *uuid.UUID `json:"movedFromTerminalGroupId,omitempty"`
	// Moved from organization ID.
	MovedFromOrganizationID *uuid.UUID `json:"movedFromOrganizationId,omitempty"`
	// External courier service.
	ExternalCourierService *DeliveryExternalCourierService `json:"externalCourierService,omitempty"`
	// Moved to delivery ID.
	MovedToDeliveryID *uuid.UUID `json:"movedToDeliveryId,omitempty"`
	// Moved to terminal group ID.
	MovedToTerminalGroupID *uuid.UUID `json:"movedToTerminalGroupId,omitempty"`
	// Moved to organization ID.
	MovedToOrganizationID *uuid.UUID `json:"movedToOrganizationId,omitempty"`
	// Menu ID.
	MenuID string `json:"menuId"`
	// Delivery zone.
	DeliveryZone string `json:"deliveryZone"`
	// Locked at time.
	LockedAt *EventTime `json:"lockedAt,omitempty"`
	// Estimated time.
	EstimatedTime *EventTime `json:"estimatedTime,omitempty"`
	// Is ASAP flag.
	IsAsap bool `json:"isAsap"`
	// When packed.
	WhenPacked *EventTime `json:"whenPacked,omitempty"`
	// Price category.
	PriceCategory *DeliveryPriceCategory `json:"priceCategory,omitempty"`
	// Order sum.
	Sum float64 `json:"sum"`
	// Order number.
	Number int `json:"number"`
	// Source key.
	SourceKey string `json:"sourceKey"`
	// Invoice printing time.
	WhenBillPrinted *EventTime `json:"whenBillPrinted,omitempty"`
	// Delivery closing time.
	WhenClosed *EventTime `json:"whenClosed,omitempty"`
	// Concept.
	Conception Conception `json:"conception"`
	// Guests information.
	GuestsInfo Guests `json:"guestsInfo"`
	// Order items.
	Items []DeliveryItem `json:"items"`
	// Combos included in order.
	Combos []DeliveryCombo `json:"combos"`
	// Order payment components.
	Payments []DeliveryPayment `json:"payments"`
	// Order tips components.
	Tips []DeliveryTip `json:"tips"`
	// Discounts/surcharges.
	Discounts []DeliveryDiscount `json:"discounts"`
	// Order type.
	OrderType DeliveryOrderType `json:"orderType"`
	// Terminal group ID.
	TerminalGroupID uuid.UUID `json:"terminalGroupId"`
	// Processed payments sum.
	ProcessedPaymentsSum float64 `json:"processedPaymentsSum"`
	// Loyalty information.
	LoyaltyInfo *DeliveryLoyaltyInfo `json:"loyaltyInfo,omitempty"`
	// External data.
	ExternalData []DeliveryExternalData `json:"externalData"`
}

// Types specific to deliveries_by_id that are not shared with delivery_create

// DeliveryStatus represents delivery status
type DeliveryStatus string

const (
	DeliveryStatusUnconfirmed      DeliveryStatus = "Unconfirmed"
	DeliveryStatusWaitCooking      DeliveryStatus = "WaitCooking"
	DeliveryStatusReadyForCooking  DeliveryStatus = "ReadyForCooking"
	DeliveryStatusCookingStarted   DeliveryStatus = "CookingStarted"
	DeliveryStatusCookingCompleted DeliveryStatus = "CookingCompleted"
	DeliveryStatusWaiting          DeliveryStatus = "Waiting"
	DeliveryStatusOnWay            DeliveryStatus = "OnWay"
	DeliveryStatusDelivered        DeliveryStatus = "Delivered"
	DeliveryStatusClosed           DeliveryStatus = "Closed"
	DeliveryStatusCancelled        DeliveryStatus = "Cancelled"
)

// DeliveryCustomer represents delivery customer
type DeliveryCustomer struct {
	Type string `json:"type"`
}

// DeliveryCancelInfo represents cancel information
type DeliveryCancelInfo struct {
	// When cancelled
	WhenCancelled EventTime `json:"whenCancelled"`
	// Cancel cause
	Cause *DeliveryCancelCause `json:"cause,omitempty"`
	// Comment
	Comment string `json:"comment"`
}

// DeliveryCancelCause represents cancel cause
type DeliveryCancelCause struct {
	// ID
	ID uuid.UUID `json:"id"`
	// Name
	Name string `json:"name"`
}

// DeliveryCourierInfo represents courier information
type DeliveryCourierInfo struct {
	// Courier
	Courier *DeliveryCourier `json:"courier,omitempty"`
	// Is courier selected manually
	IsCourierSelectedManually bool `json:"isCourierSelectedManually"`
}

// DeliveryCourier represents courier
type DeliveryCourier struct {
	// ID
	ID uuid.UUID `json:"id"`
	// Name
	Name string `json:"name"`
	// Phone
	Phone string `json:"phone"`
}

// DeliveryProblem represents delivery problem
type DeliveryProblem struct {
	// Has problem flag
	HasProblem bool `json:"hasProblem"`
	// Description
	Description string `json:"description"`
}

// DeliveryOperator represents operator
type DeliveryOperator struct {
	// ID
	ID uuid.UUID `json:"id"`
	// Name
	Name string `json:"name"`
	// Phone
	Phone string `json:"phone"`
}

// DeliveryMarketingSource represents marketing source
type DeliveryMarketingSource struct {
	// ID
	ID uuid.UUID `json:"id"`
	// Name
	Name string `json:"name"`
}

// DeliveryExternalCourierService represents external courier service
type DeliveryExternalCourierService struct {
	// ID
	ID uuid.UUID `json:"id"`
	// Name
	Name string `json:"name"`
}

// DeliveryPriceCategory represents price category
type DeliveryPriceCategory struct {
	// ID
	ID uuid.UUID `json:"id"`
	// Name
	Name string `json:"name"`
}

// DeliveryItem represents delivery item
type DeliveryItem struct {
	// Type
	Type string `json:"type"`
	// Status
	Status DeliveryItemStatus `json:"status"`
	// Deleted information
	Deleted *DeliveryItemDeleted `json:"deleted,omitempty"`
	// Amount
	Amount float32 `json:"amount"`
	// Comment
	Comment string `json:"comment"`
	// When printed
	WhenPrinted *EventTime `json:"whenPrinted,omitempty"`
	// Size
	Size *DeliverySize `json:"size,omitempty"`
	// Combo information
	ComboInformation *DeliveryComboInformation `json:"comboInformation,omitempty"`
}

// DeliveryItemStatus represents delivery item status
type DeliveryItemStatus string

const (
	DeliveryItemStatusAdded           DeliveryItemStatus = "Added"
	DeliveryItemStatusCookingStarted  DeliveryItemStatus = "CookingStarted"
	DeliveryItemStatusCookingComplete DeliveryItemStatus = "CookingComplete"
	DeliveryItemStatusServed          DeliveryItemStatus = "Served"
)

// DeliveryItemDeleted represents deleted item information
type DeliveryItemDeleted struct {
	// Deletion method
	DeletionMethod *DeliveryDeletionMethod `json:"deletionMethod,omitempty"`
}

// DeliveryDeletionMethod represents deletion method
type DeliveryDeletionMethod struct {
	// ID
	ID string `json:"id"`
	// Comment
	Comment string `json:"comment"`
	// Removal type
	RemovalType *DeliveryRemovalType `json:"removalType,omitempty"`
}

// DeliveryRemovalType represents removal type
type DeliveryRemovalType struct {
	// ID
	ID uuid.UUID `json:"id"`
	// Name
	Name string `json:"name"`
}

// DeliverySize represents size
type DeliverySize struct {
	// ID
	ID uuid.UUID `json:"id"`
	// Name
	Name string `json:"name"`
}

// DeliveryComboInformation represents combo information
type DeliveryComboInformation struct {
	// Combo ID
	ComboID uuid.UUID `json:"comboId"`
	// Combo source ID
	ComboSourceID uuid.UUID `json:"comboSourceId"`
	// Group ID
	GroupID uuid.UUID `json:"groupId"`
	// Group name
	GroupName string `json:"groupName"`
}

// DeliveryCombo represents delivery combo
type DeliveryCombo struct {
	// ID
	ID uuid.UUID `json:"id"`
	// Name
	Name string `json:"name"`
	// Amount
	Amount int `json:"amount"`
	// Price
	Price float64 `json:"price"`
	// Source ID
	SourceID uuid.UUID `json:"sourceId"`
	// Size
	Size *DeliverySize `json:"size,omitempty"`
}

// DeliveryPayment represents delivery payment
type DeliveryPayment struct {
	// Payment type
	PaymentType *DeliveryPaymentType `json:"paymentType,omitempty"`
	// Sum
	Sum float64 `json:"sum"`
	// Is preliminary
	IsPreliminary bool `json:"isPreliminary"`
	// Is external
	IsExternal bool `json:"isExternal"`
	// Is processed externally
	IsProcessedExternally bool `json:"isProcessedExternally"`
	// Is fiscalized externally
	IsFiscalizedExternally bool `json:"isFiscalizedExternally"`
	// Is prepay
	IsPrepay bool `json:"isPrepay"`
}

// DeliveryPaymentType represents delivery payment type
type DeliveryPaymentType struct {
	// ID
	ID uuid.UUID `json:"id"`
	// Name
	Name string `json:"name"`
	// Kind
	Kind PaymentTypeKind `json:"kind"`
}

// DeliveryTip represents delivery tip
type DeliveryTip struct {
	// Tips type
	TipsType *DeliveryTipsType `json:"tipsType,omitempty"`
	// Payment type
	PaymentType *DeliveryPaymentType `json:"paymentType,omitempty"`
	// Sum
	Sum float64 `json:"sum"`
	// Is preliminary
	IsPreliminary bool `json:"isPreliminary"`
	// Is external
	IsExternal bool `json:"isExternal"`
	// Is processed externally
	IsProcessedExternally bool `json:"isProcessedExternally"`
	// Is fiscalized externally
	IsFiscalizedExternally bool `json:"isFiscalizedExternally"`
	// Is prepay
	IsPrepay bool `json:"isPrepay"`
}

// DeliveryTipsType represents delivery tips type
type DeliveryTipsType struct {
	// ID
	ID uuid.UUID `json:"id"`
	// Name
	Name string `json:"name"`
}

// DeliveryOrderType represents delivery order type
type DeliveryOrderType struct {
	// ID
	ID uuid.UUID `json:"id"`
	// Name
	Name string `json:"name"`
	// Order service type
	OrderServiceType OrderServiceType `json:"orderServiceType"`
}

// DeliveriesByID Get delivery orders by IDs
//
// iiko API: /api/1/deliveries/by_id
func (c *Client) DeliveriesByID(req *DeliveriesByIDRequest, opts ...Option) (*DeliveriesByIDResponse, error) {
	var response DeliveriesByIDResponse

	if err := c.post(true, "/api/1/deliveries/by_id", req, &response, opts...); err != nil {
		return nil, err
	}

	return &response, nil
}
