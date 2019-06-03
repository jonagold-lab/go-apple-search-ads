package searchads

import (
	"context"
	"fmt"
)

type TargetingKeyword struct {
	ID               int64          `json:"id,omitempty"`
	AdGroupID        int64          `json:"adGroupId,omitempty"`
	Text             string         `json:"text,omitempty"`
	Status           KeywordStatus  `json:"status"`
	MatchType        MatchType      `json:"matchType"`
	BidAmount        Amount         `json:"bidAmount"`
	ModificationTime string         `json:"modificationTime,omitempty"`
	Deleted          bool           `json:"deleted,omitempty"`
}

// AdGroupTargetingKeywordService to handle Targeting Keywords of
type AdGroupTargetingKeywordService service

// List function to get Adgroups from campaign
func (s *AdGroupTargetingKeywordService) List(ctx context.Context, campaignID, adGroupID int64, opt *ListOptions) ([]*TargetingKeyword, *Response, error) {
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
	targetingkeywords := []*TargetingKeyword{}
	resp, err := s.client.Do(ctx, req, &targetingkeywords)
	if err != nil {
		return nil, resp, err
	}

	return targetingkeywords, resp, nil
}

// Get will get a specific targeting Keyword
func (s *AdGroupTargetingKeywordService) Get(ctx context.Context, campaignID, adGroupID, targetingkeywordID int64) (*TargetingKeyword, *Response, error) {
	if campaignID == 0 {
		return nil, nil, fmt.Errorf("campaignID can not be 0")
	}
	if adGroupID == 0 {
		return nil, nil, fmt.Errorf("adGroupID can not be 0")
	}
	if targetingkeywordID == 0 {
		return nil, nil, fmt.Errorf("targetingkeywordID can not be 0")
	}
	u := fmt.Sprintf("campaigns/%d/adgroups/%d/targetingkeywords/%d", campaignID, adGroupID, targetingkeywordID)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	targetingkeyword := TargetingKeyword{}
	resp, err := s.client.Do(ctx, req, &targetingkeyword)
	if err != nil {
		return nil, resp, err
	}

	return &targetingkeyword, resp, nil
}

// CreateBulk will create multiple Targeting Keywords for a campaign
func (s *AdGroupTargetingKeywordService) CreateBulk(ctx context.Context, campaignID, adGroupID int64, data []*TargetingKeyword) ([]*TargetingKeyword, *Response, error) {
	if campaignID == 0 {
		return nil, nil, fmt.Errorf("campaignID can not be 0")
	}
	if adGroupID == 0 {
		return nil, nil, fmt.Errorf("adGroupID can not be 0")
	}
	fmt.Println(data)
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

// UpdateBulk will update an existing Targeting on a Adgroup
func (s *AdGroupTargetingKeywordService) UpdateBulk(ctx context.Context, campaignID, adGroupID int64, data []*TargetingKeyword) ([]*TargetingKeyword, *Response, error) {
	if campaignID == 0 {
		return nil, nil, fmt.Errorf("campaignID can not be 0")
	}
	if adGroupID == 0 {
		return nil, nil, fmt.Errorf("adGroupID can not be 0")
	}
	fmt.Println(data)
	u := fmt.Sprintf("campaigns/%d/adgroups/%d/targetingkeywords/bulk", campaignID, adGroupID)
	req, err := s.client.NewRequest("PUT", u, data)
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
