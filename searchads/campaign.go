package searchads

import (
	"context"
	"fmt"
)

// CampaignService struct to hold individual service
type CampaignService service

// Campaign to hold information about a campaign
type Campaign struct {
	ID                                 int64              `json:"id,omitempty"`
	OrgID                              int64              `json:"orgId,omitempty"`
	Name                               string             `json:"name,omitempty"`
	BudgetAmount                       Amount             `json:"budgetAmount,omitempty"`
	DailyBudgetAmount                  Amount             `json:"dailyBudgetAmount,omitempty"`
	AdamID                             int64              `json:"adamId,omitempty"`
	PaymentModel                       PaymentModel       `json:"paymentModel,omitempty"`
	BudgetOrders                       []int              `json:"budgetOrders,omitempty"`
	Status                             Status             `json:"status,omitempty"`
	ServingStatus                      ServingStatus      `json:"servingStatus,omitempty"`
	ServingStateReasons                interface{}        `json:"servingStateReasons,omitempty"`
	ModificationTime                   string             `json:"modificationTime,omitempty"`
	StartTime                          string             `json:"startTime,omitempty"`
	EndTime                            string             `json:"endTime,omitempty"`
	Deleted                            bool               `json:"deleted,omitempty"`
	CountriesOrRegions                 []CountryCode      `json:"countriesOrRegions,omitempty"`
	CountryOrRegionServingStateReasons interface{}        `json:"CountryOrRegionServingStateReasons,omitempty"`
	SapinLawResponse                   SapinLawResponse   `json:"sapinLawResponse,omitempty"`
	LocInvoiceDetails                  *LocInvoiceDetails `json:"locInvoiceDetails,omitempty"`
}

// LocInvoiceDetails needs to be used for LOC Payment
type LocInvoiceDetails struct {
	ClientName  string `json:"clientName,omitempty"`
	OrderNumber string `json:"orderNumber,omitempty"`
	BuyerName   string `json:"buyerName,omitempty"`
	BuyerEmail  string `json:"buyerEmail,omitempty"`
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
func (s *CampaignService) Edit(ctx context.Context, id int64, data *Campaign) (*Campaign, *Response, error) {
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
func (s *CampaignService) Delete(ctx context.Context, id int64) (*Response, error) {
	u := fmt.Sprintf("campaigns/%v", id)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}
