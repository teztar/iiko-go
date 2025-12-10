package iiko

import (
	"github.com/google/uuid"
)

type GetProgramsRequest struct {
	// Organization ID.
	//
	// Can be obtained by /api/1/organizations operation. [required]
	OrganizationID uuid.UUID `json:"organizationId"`

	// Exclude marketing campaigns from response. [optional]
	WithoutMarketingCampaigns bool `json:"withoutMarketingCampaigns"`
}

type GetProgramsResponse struct {
	// List of loyalty programs. [required]
	Programs []Program `json:"Programs"`
}

type Program struct {
	// Program ID. [required]
	ID uuid.UUID `json:"id"`

	// Program name. [required]
	Name string `json:"name"`

	// Program description. [optional]
	Description *string `json:"description"`

	// Service period start date. [required]
	ServiceFrom string `json:"serviceFrom"`

	// Service period end date. [optional]
	ServiceTo *string `json:"serviceTo"`

	// Notify about balance changes. [required]
	NotifyAboutBalanceChanges bool `json:"notifyAboutBalanceChanges"`

	// Program type. [required]
	ProgramType int `json:"programType"`

	// Is program active. [required]
	IsActive bool `json:"isActive"`

	// Wallet ID. [optional]
	WalletID *uuid.UUID `json:"walletId"`

	// Marketing campaigns. [required]
	MarketingCampaigns []MarketingCampaign `json:"marketingCampaigns"`

	// Applied organizations. [required]
	AppliedOrganizations []uuid.UUID `json:"appliedOrganizations"`

	// Template type. [required]
	TemplateType int `json:"templateType"`

	// Is exchange rate enabled. [required]
	IsExchangeRateEnabled bool `json:"isExchangeRateEnabled"`

	// Refill type. [required]
	RefillType int `json:"refillType"`
}

type MarketingCampaign struct {
	// Campaign ID. [required]
	ID uuid.UUID `json:"id"`

	// Program ID. [required]
	ProgramID uuid.UUID `json:"programId"`

	// Campaign name. [required]
	Name string `json:"name"`

	// Campaign description. [optional]
	Description *string `json:"description"`

	// Is campaign active. [required]
	IsActive bool `json:"isActive"`

	// Period start date. [required]
	PeriodFrom string `json:"periodFrom"`

	// Period end date. [optional]
	PeriodTo *string `json:"periodTo"`

	// Order action condition bindings. [required]
	OrderActionConditionBindings []ActionConditionBinding `json:"orderActionConditionBindings"`

	// Periodic action condition bindings. [required]
	PeriodicActionConditionBindings []ActionConditionBinding `json:"periodicActionConditionBindings"`

	// Overdraft action condition bindings. [required]
	OverdraftActionConditionBindings []ActionConditionBinding `json:"overdraftActionConditionBindings"`

	// Guest registration action condition bindings. [required]
	GuestRegistrationActionConditionBindings []ActionConditionBinding `json:"guestRegistrationActionConditionBindings"`
}

type ActionConditionBinding struct {
	// Binding ID. [required]
	ID uuid.UUID `json:"id"`

	// Stop further execution. [required]
	StopFurtherExecution bool `json:"stopFurtherExecution"`

	// Actions. [required]
	Actions []Action `json:"actions"`

	// Conditions. [required]
	Conditions []Condition `json:"conditions"`
}

type Action struct {
	// Action ID. [required]
	ID uuid.UUID `json:"id"`

	// Action settings (JSON string). [required]
	Settings string `json:"settings"`

	// Action type name. [required]
	TypeName string `json:"typeName"`

	// Checksum. [required]
	CheckSum string `json:"checkSum"`
}

type Condition struct {
	// Condition ID. [required]
	ID uuid.UUID `json:"id"`

	// Condition settings (JSON string). [required]
	Settings string `json:"settings"`

	// Condition type name. [required]
	TypeName string `json:"typeName"`

	// Checksum. [required]
	CheckSum string `json:"checkSum"`
}

// GetPrograms returns loyalty programs.
//
// iiko API: /api/1/loyalty/iiko/program
func (c *Client) GetPrograms(req *GetProgramsRequest, opts ...Option) (*GetProgramsResponse, error) {
	var response GetProgramsResponse

	if err := c.post(true, "/api/1/loyalty/iiko/program", req, &response, opts...); err != nil {
		return nil, err
	}

	return &response, nil
}
