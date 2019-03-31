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
