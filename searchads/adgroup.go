package searchads

import (
	"context"
	"fmt"
)

// AdGroupService struct to hold individual service
type AdGroupService service

type AdGroup struct {
	ID                     int                 `json:"id,omitempty"`
	CampaignID             int                 `json:"campaignId,omitempty"`
	Name                   string              `json:"name,omitempty"`
	CpaGoal                Amount              `json:"cpaGoal,omitempty"`
	Storefronts            []string            `json:"storefronts,omitempty"`
	StartTime              string              `json:"startTime,omitempty"`
	EndTime                string              `json:"endTime,omitempty"`
	AutomatedKeywordsOptIn bool                `json:"automatedKeywordsOptIn,omitempty"`
	DefaultCpcBid          Amount              `json:"defaultCpcBid,omitempty"`
	Keywords               []Keywords          `json:"keywords,omitempty"`
	NegativeKeywords       []NegativeKeyword   `json:"negativeKeywords,omitempty"`
	TargetingDimensions    TargetingDimensions `json:"targetingDimensions,omitempty"`
	OrgID                  int                 `json:"orgId,omitempty"`
	ModificationTime       string              `json:"modificationTime,omitempty"`
	Status                 string              `json:"status,omitempty"`
	ServingStatus          string              `json:"servingStatus,omitempty"`
	ServingStateReasons    []string            `json:"servingStateReasons,omitempty"`
	DisplayStatus          string              `json:"displayStatus,omitempty"`
	Deleted                bool                `json:"deleted,omitempty"`
}

type Keywords struct {
	ID               int    `json:"id"`
	AdGroupID        int    `json:"adGroupId"`
	Text             string `json:"text"`
	Status           string `json:"status"`
	MatchType        string `json:"matchType"`
	BidAmount        Amount `json:"bidAmount"`
	ModificationTime string `json:"modificationTime"`
	Deleted          bool   `json:"deleted"`
}
type DeviceClass struct {
	Included []string `json:"included"`
}
type AppDownloaders struct {
	Included []string `json:"included"`
	Excluded []string `json:"excluded"`
}
type TargetingDimensions struct {
	Age            *Age           `json:"age"`
	Gender         *Gender        `json:"gender"`
	Country        *Country       `json:"country"`
	AdminArea      *AdminArea     `json:"adminArea"`
	Locality       *Locality      `json:"locality"`
	DeviceClass    DeviceClass    `json:"deviceClass"`
	Daypart        *Daypart       `json:"daypart"`
	AppDownloaders AppDownloaders `json:"appDownloaders"`
}

type Country struct {
	Included []string `json:"included"`
}

type Age struct {
	Included []AgeObj `json:"included"`
}

type AgeObj struct {
	MinAge string `json:"minAge"`
	MaxAge string `json:"maxAge"`
}

type Gender struct {
	Included []string `json:"included"`
}

type AdminArea struct {
	Included []string `json:"included"`
}

type Locality struct {
	Included []string `json:"included"`
}

type Daypart struct {
	UserTime UserTime `json:"userTime"`
}

type UserTime struct {
	Included []string `json:"included"`
}

// List function to get Adgroups from campaign
func (s *AdGroupService) List(ctx context.Context, campaignID int, opt *ListOptions) ([]*AdGroup, *Response, error) {
	if campaignID == 0 {
		return nil, nil, fmt.Errorf("campaignID can not be 0")
	}
	u, err := addOptions(fmt.Sprintf("campaigns/%d/adgroups", campaignID), opt)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	adGroups := []*AdGroup{}

	resp, err := s.client.Do(ctx, req, &adGroups)
	if err != nil {
		return nil, resp, err
	}

	return adGroups, resp, nil
}

// Get function to get specific AdGroup by id and campaignID
func (s *AdGroupService) Get(ctx context.Context, campaignID int, id int) (*AdGroup, *Response, error) {
	if campaignID == 0 {
		return nil, nil, fmt.Errorf("campaignID can not be 0")
	}
	u := fmt.Sprintf("campaigns/%d/adgroups/%d", campaignID, id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	adGroup := new(AdGroup)
	resp, err := s.client.Do(ctx, req, adGroup)
	if err != nil {
		return nil, resp, err
	}

	return adGroup, resp, nil
}

// Create will create a new AdGroup on a specific campaign
func (s *AdGroupService) Create(ctx context.Context, campaignID int, data *AdGroup) (*AdGroup, *Response, error) {
	u := fmt.Sprintf("campaigns/%d/adgroups", campaignID)
	req, err := s.client.NewRequest("POST", u, data)
	if err != nil {
		return nil, nil, err
	}

	adG := new(AdGroup)
	resp, err := s.client.Do(ctx, req, adG)
	if err != nil {
		return nil, resp, err
	}

	return adG, resp, nil
}

// Edit will update an existing AdGroup on a Campaign
func (s *AdGroupService) Edit(ctx context.Context, campaignID int, id int, data *AdGroup) (*AdGroup, *Response, error) {
	u := fmt.Sprintf("campaigns/%d/adgroups/%d", campaignID, id)
	req, err := s.client.NewRequest("PUT", u, data)
	if err != nil {
		return nil, nil, err
	}

	adG := new(AdGroup)
	resp, err := s.client.Do(ctx, req, adG)
	if err != nil {
		return nil, resp, err
	}

	return adG, resp, nil
}

// Delete will remove an existing AdGroup on a Campaign
func (s *AdGroupService) Delete(ctx context.Context, campaignID int, id int) (*Response, error) {
	u := fmt.Sprintf("campaigns/%d/adgroups/%d", campaignID, id)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}
