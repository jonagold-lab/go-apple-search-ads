package searchads

// NegativeKeyword to define negative Keyword and connection to other
type NegativeKeyword struct {
	ID               int       `json:"id,omitempty"`
	CampaignID       int       `json:"campaignId,omitempty"`
	AdGroupID        int       `json:"adGroupId,omitempty"`
	ModificationTime string    `json:"modificationTime,omitempty"`
	Text             string    `json:"text,omitempty"`
	MatchType        MatchType `json:"matchType,omitempty"`
	Status           Status    `json:"status,omitempty"`
	Deleted          bool      `json:"deleted,omitempty"`
}
