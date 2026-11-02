package iiko

import (
	"github.com/google/uuid"
)

type AccessTokenRequest struct {
	// API login. It is set in iikoWeb [required]
	ApiLogin string `json:"apiLogin"`
}

type AccessTokenResponse struct {
	// Operation ID. [required]
	CorrelationID uuid.UUID `json:"correlationId"`

	// Authentication token. The standard token lifetime is 1 hour. [required]
	Token string `json:"token"`
}

// AccessTokenV2Request is the new iiko authorization scheme: the API key from
// iikoWeb plus appId and clientSecret of the application registered on the
// iiko developer portal. All three are required and iiko rejects unknown fields,
// so apiLogin is not sent here.
type AccessTokenV2Request struct {
	// API key generated in iikoWeb under "Integrations → API Keys". [required]
	ApiKey string `json:"apiKey"`

	// Application identifier from the iiko developer portal. [required]
	AppId string `json:"appId"`

	// Application secret from the iiko developer portal. [required]
	ClientSecret string `json:"clientSecret"`
}

// Retrieve session key for API user.
//
// iiko API: /api/1/access_token
func (c *Client) accessToken(req *AccessTokenRequest, opts ...Option) (*AccessTokenResponse, error) {
	var response AccessTokenResponse

	if err := c.post(false, "/api/1/access_token", req, &response, opts...); err != nil {
		return nil, err
	}

	return &response, nil
}

// Retrieve session key for API access (v2). Keys issued after 2026-06-01 work
// only here: /api/1/access_token answers "This API key does not support
// /api/1/access_token. Please use /api/v2/access_token instead."
//
// iiko API: /api/v2/access_token
func (c *Client) accessTokenV2(req *AccessTokenV2Request, opts ...Option) (*AccessTokenResponse, error) {
	var response AccessTokenResponse

	if err := c.post(false, "/api/v2/access_token", req, &response, opts...); err != nil {
		return nil, err
	}

	return &response, nil
}
