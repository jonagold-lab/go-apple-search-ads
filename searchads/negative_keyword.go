package searchads

import (
	"context"
	"fmt"
)

// NegativeKeyword to define negative Keyword and connection to other
type NegativeKeyword struct {
	ID               int64         `json:"id,omitempty"`
	CampaignID       int64         `json:"campaignId,omitempty"`
	AdGroupID        int64         `json:"adGroupId,omitempty"`
	ModificationTime string        `json:"modificationTime,omitempty"`
	Text             string        `json:"text,omitempty"`
	MatchType        MatchType     `json:"matchType,omitempty"`
	Status           KeywordStatus `json:"status,omitempty"`
	Deleted          bool          `json:"deleted,omitempty"`
}

// CampaignNegativeKeywordServive to handle Negative Keywords of
type CampaignNegativeKeywordServive service

// List function to get Adgroups from campaign
func (s *CampaignNegativeKeywordServive) List(ctx context.Context, campaignID int64, opt *ListOptions) ([]*NegativeKeyword, *Response, error) {
	if campaignID == 0 {
		return nil, nil, fmt.Errorf("campaignID can not be 0")
	}
	u, err := addOptions(fmt.Sprintf("campaigns/%d/negativekeywords", campaignID), opt)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	negativekeywords := []*NegativeKeyword{}
	resp, err := s.client.Do(ctx, req, &negativekeywords)
	if err != nil {
		return nil, resp, err
	}

	return negativekeywords, resp, nil
}

// CreateBulk will create multiple Negative Keywords for a campaign
func (s *CampaignNegativeKeywordServive) CreateBulk(ctx context.Context, campaignID int64, data []*NegativeKeyword) ([]*NegativeKeyword, *Response, error) {
	if campaignID == 0 {
		return nil, nil, fmt.Errorf("campaignID can not be 0")
	}
	u := fmt.Sprintf("campaigns/%d/negativekeywords/bulk", campaignID)
	req, err := s.client.NewRequest("POST", u, data)
	if err != nil {
		return nil, nil, err
	}
	negativekeywords := []*NegativeKeyword{}
	resp, err := s.client.Do(ctx, req, &negativekeywords)
	if err != nil {
		return nil, resp, err
	}
	return negativekeywords, resp, nil
}

// Delete will remove an existing Negative Keywords on a Adgroup
func (s *CampaignNegativeKeywordServive) Delete(ctx context.Context, campaignID, id int64) (*Response, error) {
	if campaignID == 0 {
		return nil, fmt.Errorf("campaignID can not be 0")
	}
	if id == 0 {
		return nil, fmt.Errorf("id can not be 0")
	}
	u := fmt.Sprintf("campaigns/%d/negativekeywords/%d", campaignID, id)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// AdGroupNegativeKeywordServive to handle Negative Keywords of
type AdGroupNegativeKeywordServive service

// List function to get Adgroups from campaign
func (s *AdGroupNegativeKeywordServive) List(ctx context.Context, campaignID int64, adGroupID int64, opt *ListOptions) ([]*NegativeKeyword, *Response, error) {
	if campaignID == 0 {
		return nil, nil, fmt.Errorf("campaignID can not be 0")
	}
	if adGroupID == 0 {
		return nil, nil, fmt.Errorf("adGroupID can not be 0")
	}
	u, err := addOptions(fmt.Sprintf("campaigns/%d/adgroups/%d/negativekeywords", campaignID, adGroupID), opt)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	negativekeywords := []*NegativeKeyword{}
	resp, err := s.client.Do(ctx, req, &negativekeywords)
	if err != nil {
		return nil, resp, err
	}

	return negativekeywords, resp, nil
}

// CreateBulk will create multiple Negative Keywords for a campaign
func (s *AdGroupNegativeKeywordServive) CreateBulk(ctx context.Context, campaignID int64, adGroupID int64, data []*NegativeKeyword) ([]*NegativeKeyword, *Response, error) {
	if campaignID == 0 {
		return nil, nil, fmt.Errorf("campaignID can not be 0")
	}
	if adGroupID == 0 {
		return nil, nil, fmt.Errorf("adGroupID can not be 0")
	}
	u := fmt.Sprintf("campaigns/%d/adgroups/%d/negativekeywords/bulk", campaignID, adGroupID)
	req, err := s.client.NewRequest("POST", u, data)
	if err != nil {
		return nil, nil, err
	}
	negativekeywords := []*NegativeKeyword{}
	resp, err := s.client.Do(ctx, req, &negativekeywords)
	if err != nil {
		return nil, resp, err
	}

	return negativekeywords, resp, nil
}

// Delete will remove an existing Negative Keywords on a Adgroup
func (s *AdGroupNegativeKeywordServive) Delete(ctx context.Context, campaignID, adGroupID, id int64) (*Response, error) {
	if campaignID == 0 {
		return nil, fmt.Errorf("campaignID can not be 0")
	}
	if adGroupID == 0 {
		return nil, fmt.Errorf("adGroupID can not be 0")
	}
	if id == 0 {
		return nil, fmt.Errorf("id can not be 0")
	}
	u := fmt.Sprintf("campaigns/%d/adgroups/%d/negativekeywords/%d", campaignID, adGroupID, id)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}
