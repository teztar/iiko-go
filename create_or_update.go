package iiko

type CreateOrUpdateRequest struct {
	// Customer uuid
	Id *string `json:"id"`
	// Customer phone
	Phone *string `json:"phone"`
	// Card track
	CardTrack *string `json:"cardTrack"`
	// Card number
	CardNumber *string `json:"cardNumber"`
	// Customer name
	Name *string `json:"name"`
	// Customer middle name
	MiddleName *string `json:"middleName"`
	// Customer surname
	SurName *string `json:"surName"`
	// Customer birthday <yyyy-MM-dd HH:mm:ss.fff>
	Birthday *string `json:"birthday"`
	// Customer email
	Email         *string       `json:"email"`
	Sex           SexType       `json:"sex"`
	ConsentStatus ConsentStatus `json:"consentStatus"`
	// Customer get promo messages (email, sms). If null - unknown
	ShouldReceivePromoActionsInfo *bool `json:"shouldReceivePromoActionsInfo"`
	// Id for referrer guest. Null for old integrations, Guid.Empty - for referrer deletion
	ReferrerId *string `json:"referrerId"`
	// Customer user data
	UserData *string `json:"userData"`
	// Customer organization id
	OrganizationId string `json:"organizationId"`
}

type CreateOrUpdateResponse struct {
	// Customer uuid
	Id string `json:"id"`
}

// CreateOrUpdate Create or update customer info by id or phone or card track
//
// iiko API: /api/1/loyalty/iiko/customer/create_or_update
func (c *Client) CreateOrUpdate(req *CreateOrUpdateRequest, opts ...Option) (*CreateOrUpdateResponse, error) {
	var response CreateOrUpdateResponse

	if err := c.post(true, "/api/1/loyalty/iiko/customer/create_or_update", req, &response, opts...); err != nil {
		return nil, err
	}

	return &response, nil
}
