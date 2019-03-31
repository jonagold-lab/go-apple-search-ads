package searchads

import (
	"context"
	"fmt"
)

type TargetingKeyword struct {
	ID               int           `json:"id,omitempty"`
	AdGroupID        int           `json:"adGroupId,omitempty"`
	Text             string        `json:"text,omitempty"`
	Status           KeywordStatus `json:"status,omitempty"`
	MatchType        MatchType     `json:"matchType,omitempty"`
	BidAmount        Amount        `json:"bidAmount,omitempty"`
	ModificationTime string        `json:"modificationTime,omitempty"`
	Deleted          bool          `json:"deleted,omitempty"`
}

// AdGroupTargetingKeywordServive to handle Targeting Keywords of
type AdGroupTargetingKeywordServive service

// List function to get Adgroups from campaign
func (s *AdGroupTargetingKeywordServive) List(ctx context.Context, campaignID int, adGroupID int, opt *ListOptions) ([]*TargetingKeyword, *Response, error) {
	if campaignID == 0 {
		return nil, nil, fmt.Errorf("campaignID can not be 0")
	}
	if adGroupID == 0 {
		return nil, nil, fmt.Errorf("adGroupID can not be 0")
	}
	u, err := addOptions(fmt.Sprintf("campaigns/%d/adgroups/%d/targetingkeywords", campaignID, adGroupID), opt)
	if err != nil {
		return nil, nil, err
	}
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	Targetingkeywords := []*TargetingKeyword{}
	resp, err := s.client.Do(ctx, req, &Targetingkeywords)
	if err != nil {
		return nil, resp, err
	}

	return Targetingkeywords, resp, nil
}

// CreateBulk will create multiple Targeting Keywords for a campaign
func (s *AdGroupTargetingKeywordServive) CreateBulk(ctx context.Context, campaignID int, adGroupID int, data []*TargetingKeyword) ([]*TargetingKeyword, *Response, error) {
	if campaignID == 0 {
		return nil, nil, fmt.Errorf("campaignID can not be 0")
	}
	if adGroupID == 0 {
		return nil, nil, fmt.Errorf("adGroupID can not be 0")
	}
	u := fmt.Sprintf("campaigns/%d/adgroups/%d/targetingkeywords/bulk", campaignID, adGroupID)
	req, err := s.client.NewRequest("POST", u, data)
	if err != nil {
		return nil, nil, err
	}

	targetingkeywords := []*TargetingKeyword{}
	resp, err := s.client.Do(ctx, req, &targetingkeywords)
	if err != nil {
		return nil, resp, err
	}

	return targetingkeywords, resp, nil
}

// Delete will remove an existing Targeting on a Adgroup
func (s *AdGroupTargetingKeywordServive) Delete(ctx context.Context, campaignID, adGroupID, id int) (*Response, error) {
	if campaignID == 0 {
		return nil, fmt.Errorf("campaignID can not be 0")
	}
	if adGroupID == 0 {
		return nil, fmt.Errorf("adGroupID can not be 0")
	}
	if id == 0 {
		return nil, fmt.Errorf("id can not be 0")
	}
	u := fmt.Sprintf("campaigns/%d/adgroups/%d/targetingkeywords/%d", campaignID, adGroupID, id)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}
