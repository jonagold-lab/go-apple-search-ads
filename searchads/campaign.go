package searchads

import (
	"context"
	"fmt"
)

// CampaignService struct to hold individual service
type CampaignService service

// Campaign to hold information about a campaign
type Campaign struct {
	ID                  int               `json:"id,omitempty"`
	OrgID               int               `json:"orgId,omitempty"`
	AdamID              int               `json:"adamId,omitempty"`
	Name                string            `json:"name,omitempty"`
	BudgetAmount        Amount            `json:"budgetAmount,omitempty"`
	DailyBudgetAmount   Amount            `json:"dailyBudgetAmount,omitempty"`
	ReferenceLabel      string            `json:"referenceLabel,omitempty"`
	NegativeKeywords    []NegativeKeyword `json:"negativeKeywords,omitempty"`
	PaymentModel        PaymentModel      `json:"paymentModel,omitempty"`
	LocInvoiceDetails   string            `json:"locInvoiceDetails,omitempty"`
	BudgetOrders        []string          `json:"budgetOrders,omitempty"`
	DisplayStatus       string            `json:"displayStatus,omitempty"`
	StartTime           string            `json:"startTime,omitempty"`
	EndTime             string            `json:"endTime,omitempty"`
	Status              Status            `json:"status,omitempty"`
	ServingStatus       ServingStatus     `json:"servingStatus,omitempty"`
	AdditionalDetail    string            `json:"additionalDetail,omitempty"`
	ServingStateReasons []string          `json:"servingStateReasons,omitempty"`
	Storefront          CountryCode       `json:"storefront,omitempty"`
	ModificationTime    string            `json:"modificationTime,omitempty"`
	Deleted             bool              `json:"deleted,omitempty"`
	SapinLawResponse    SapinLawResponse  `json:"sapinLawResponse,omitempty"`
	AdGroups            []AdGroup         `json:"adGroups,omitempty"`
}

// NegativeKeyword to define negative Keyword and connection to other
type NegativeKeyword struct {
	ID               int    `json:"id"`
	CampaignID       int    `json:"campaignId"`
	AdGroupID        int    `json:"adGroupId"`
	ModificationTime string `json:"modificationTime"`
	Text             string `json:"text"`
	MatchType        string `json:"matchType"`
	Status           string `json:"status"`
	Deleted          bool   `json:"deleted"`
}

// Amount  to hold amount and currency for various fields
type Amount struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

// List function to get Campaigns
func (s *CampaignService) List(ctx context.Context, opt *ListOptions) ([]*Campaign, *Response, error) {
	u, err := addOptions("campaigns", opt)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	campaigns := []*Campaign{}

	resp, err := s.client.Do(ctx, req, &campaigns)
	if err != nil {
		return nil, resp, err
	}

	return campaigns, resp, nil
}

// Get function to get specific campaign by id
func (s *CampaignService) Get(ctx context.Context, id int) (*Campaign, *Response, error) {
	u := fmt.Sprintf("campaigns/%d", id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	campaign := new(Campaign)
	resp, err := s.client.Do(ctx, req, campaign)
	if err != nil {
		return nil, resp, err
	}

	return campaign, resp, nil
}

// Create will create a new Campaign
func (s *CampaignService) Create(ctx context.Context, data *Campaign) (*Campaign, *Response, error) {
	req, err := s.client.NewRequest("POST", "campaigns", data)
	if err != nil {
		return nil, nil, err
	}

	c := new(Campaign)
	resp, err := s.client.Do(ctx, req, c)
	if err != nil {
		return nil, resp, err
	}

	return c, resp, nil
}

// Edit will update an existing Campaign
func (s *CampaignService) Edit(ctx context.Context, id int, data *Campaign) (*Campaign, *Response, error) {
	u := fmt.Sprintf("campaigns/%v", id)
	req, err := s.client.NewRequest("PUT", u, data)
	if err != nil {
		return nil, nil, err
	}

	c := new(Campaign)
	resp, err := s.client.Do(ctx, req, c)
	if err != nil {
		return nil, resp, err
	}

	return c, resp, nil
}

// Delete will remove an existing Campaign
func (s *CampaignService) Delete(ctx context.Context, id int) (*Response, error) {
	u := fmt.Sprintf("campaigns/%v", id)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}
