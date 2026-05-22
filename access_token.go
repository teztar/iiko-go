package iiko

import (
	"github.com/google/uuid"
)

type AccessTokenRequest struct {
	// API login. It is set in iikoWeb [required]
	ApiLogin string `json:"apiLogin"`

	// AppId and ClientSecret are part of the new iiko authorization scheme
	// (mandatory from 2026-06-01), obtained when registering the integration
	// on the iiko developer portal. The existing ApiLogin stays the same.
	// Both are optional for backward compatibility; when empty they are
	// omitted from the request body.
	AppId        string `json:"appId,omitempty"`
	ClientSecret string `json:"clientSecret,omitempty"`
}

type AccessTokenResponse struct {
	// Operation ID. [required]
	CorrelationID uuid.UUID `json:"correlationId"`

	// Authentication token. The standard token lifetime is 1 hour. [required]
	Token string `json:"token"`
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
