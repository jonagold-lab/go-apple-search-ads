package searchads

type TargetingKeyword struct {
	ID               int       `json:"id,omitempty"`
	AdGroupID        int       `json:"adGroupId,omitempty"`
	Text             string    `json:"text,omitempty"`
	Status           Status    `json:"status,omitempty"`
	MatchType        MatchType `json:"matchType,omitempty"`
	BidAmount        Amount    `json:"bidAmount,omitempty"`
	ModificationTime string    `json:"modificationTime,omitempty"`
	Deleted          bool      `json:"deleted,omitempty"`
}

/*

type TargetingKeywordService service

// List function to get Adgroups from campaign
func (s *TargetingKeywordService) List(ctx context.Context, campaignID int, opt *ListOptions) ([]*AdGroup, *Response, error) {
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
*/
