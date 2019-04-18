package searchads

import (
	"context"
	"fmt"
)

// ReportService a service to do campaign based reports
type ReportService service

// CampaignReport to hold Reports of campaign especially Metadata
type CampaignReport struct {
	ReportingDataResponse CampaignReportingDataResponse `json:"reportingDataResponse"`
}

type CampaignReportingDataResponse struct {
	Row         []CampaignReportRow `json:"row"`
	GrandTotals GrandTotals         `json:"grandTotals"`
}

type CampaignReportRow struct {
	Other       bool             `json:"other"`
	Metadata    CampaignMetadata `json:"metadata"`
	Granularity []Statistics     `json:"granularity,omitemtpy"`
	Total       Statistics       `json:"total"`
}
type CampaignMetadata struct {
	CampaignID                         int               `json:"campaignId"`
	CampaignName                       string            `json:"campaignName"`
	Deleted                            bool              `json:"deleted"`
	CampaignStatus                     Status            `json:"campaignStatus"`
	App                                App               `json:"app"`
	ServingStatus                      ServingStatus     `json:"servingStatus"`
	ServingStateReasons                *string           `json:"servingStateReasons"`
	CountriesOrRegions                 []CountryCode     `json:"countriesOrRegions"`
	ModificationTime                   string            `json:"modificationTime"`
	TotalBudget                        Amount            `json:"totalBudget"`
	DailyBudget                        Amount            `json:"dailyBudget"`
	DisplayStatus                      DisplayStatus     `json:"displayStatus"`
	OrgID                              int               `json:"orgId"`
	CountryOrRegionServingStateReasons map[string]string `json:"countryOrRegionServingStateReasons"`
	CountryOrRegion                    CountryCode       `json:"countryOrRegion"`
}

// Campaigns to get reports from all campaigns with filter
func (s *ReportService) Campaigns(ctx context.Context, filter *ReportFilter) (*CampaignReport, *Response, error) {
	u := fmt.Sprintf("reports/campaigns")
	req, err := s.client.NewRequest("POST", u, filter)
	if err != nil {
		return nil, nil, err
	}
	report := new(CampaignReport)
	resp, err := s.client.Do(ctx, req, &report)
	if err != nil {
		return nil, resp, err
	}
	return report, resp, nil
}

// AdGroupReport to hold Reports of campaign especially Metadata
type AdGroupReport struct {
	ReportingDataResponse AdGroupReportingDataResponse `json:"reportingDataResponse"`
}

type AdGroupReportingDataResponse struct {
	Row         []AdGroupReportRow `json:"row"`
	GrandTotals GrandTotals        `json:"grandTotals"`
}

type AdGroupReportRow struct {
	Other       bool            `json:"other"`
	Metadata    AdGroupMetadata `json:"metadata"`
	Granularity []Statistics    `json:"granularity,omitemtpy"`
	Total       Statistics      `json:"total"`
}
type AdGroupMetadata struct {
	AdGroupID                  int                `json:"adGroupId"`
	AdGroupName                string             `json:"adGroupName"`
	StartTime                  string             `json:"startTime"`
	EndTime                    *string            `json:"endTime,omitempty"`
	CpaGoal                    Amount             `json:"cpaGoal"`
	DefaultCpcBid              Amount             `json:"defaultCpcBid"`
	Deleted                    bool               `json:"deleted"`
	AdGroupStatus              Status             `json:"adGroupStatus"`
	AdGroupServingStatus       ServingStatus      `json:"adGroupServingStatus"`
	AdGroupServingStateReasons *map[string]string `json:"servingStateReasons"`
	ModificationTime           string             `json:"modificationTime"`
	AdGroupDisplayStatus       DisplayStatus      `json:"adGroupDisplayStatus"`
	AutomatedKeywordsOptIn     bool               `json:"automatedKeywordsOptIn"`
}

// AdGroups to return reports of Adgroups
func (s *ReportService) AdGroups(ctx context.Context, campaignID int, filter *ReportFilter) (*AdGroupReport, *Response, error) {
	if campaignID == 0 {
		return nil, nil, fmt.Errorf("campaignID can not be 0")
	}
	u := fmt.Sprintf("reports/campaigns/%d/adgroups", campaignID)
	req, err := s.client.NewRequest("POST", u, filter)
	if err != nil {
		return nil, nil, err
	}
	report := new(AdGroupReport)
	resp, err := s.client.Do(ctx, req, &report)
	if err != nil {
		return nil, resp, err
	}
	return report, resp, nil
}

// KeywordReport to hold Reports of campaign especially Metadata
type KeywordReport struct {
	ReportingDataResponse KeywordReportingDataResponse `json:"reportingDataResponse"`
}

type KeywordReportingDataResponse struct {
	Row         []KeywordReportRow `json:"row"`
	GrandTotals GrandTotals        `json:"grandTotals"`
}

type KeywordReportRow struct {
	Other       bool            `json:"other"`
	Metadata    KeywordMetadata `json:"metadata"`
	Granularity []Statistics    `json:"granularity,omitemtpy"`
	Total       Statistics      `json:"total"`
}
type KeywordMetadata struct {
	KeywordID            int                  `json:"keywordId"`
	Keyword              string               `json:"keyword"`
	KeywordStatus        KeywordStatus        `json:"keywordStatus"`
	MatchType            MatchType            `json:"matchType"`
	BidAmount            Amount               `json:"bidAmount"`
	Deleted              bool                 `json:"deleted"`
	KeywordDisplayStatus KeywordDisplayStatus `json:"keywordDisplayStatus"`
	AdGroupID            int                  `json:"adGroupId"`
	AdGroupName          string               `json:"adGroupName"`
	AdGroupDeleted       bool                 `json:"adGroupDeleted"`
	ModificationTime     string               `json:"modificationTime"`
}

// Keywords to return reports of Kewords
func (s *ReportService) Keywords(ctx context.Context, campaignID int, filter *ReportFilter) (*KeywordReport, *Response, error) {
	if campaignID == 0 {
		return nil, nil, fmt.Errorf("campaignID can not be 0")
	}
	u := fmt.Sprintf("reports/campaigns/%d/keywords", campaignID)
	req, err := s.client.NewRequest("POST", u, filter)
	if err != nil {
		return nil, nil, err
	}
	report := new(KeywordReport)
	resp, err := s.client.Do(ctx, req, &report)
	if err != nil {
		return nil, resp, err
	}
	return report, resp, nil
}

// SearchTermsReport to hold Reports of campaign especially Metadata
type SearchTermsReport struct {
	ReportingDataResponse SearchTermReportingDataResponse `json:"reportingDataResponse"`
}

type SearchTermReportingDataResponse struct {
	Row         []SearchTermsReportRow `json:"row"`
	GrandTotals GrandTotals            `json:"grandTotals"`
}

type SearchTermsReportRow struct {
	Other       bool               `json:"other"`
	Metadata    SearchTermMetadata `json:"metadata"`
	Granularity []Statistics       `json:"granularity,omitemtpy"`
	Total       Statistics         `json:"total"`
}
type SearchTermMetadata struct {
	KeywordID            int                  `json:"keywordId"`
	Keyword              string               `json:"keyword"`
	MatchType            MatchType            `json:"matchType"`
	BidAmount            Amount               `json:"bidAmount"`
	KeywordDisplayStatus KeywordDisplayStatus `json:"keywordDisplayStatus"`
	Deleted              bool                 `json:"deleted"`
	AdGroupID            int                  `json:"adGroupId"`
	AdGroupName          string               `json:"adGroupName"`
	AdGroupDeleted       bool                 `json:"adGroupDeleted"`
	SearchTermText       *string              `json:"searchTermText"`
	SearchTermSource     SearchTermSource     `json:"searchTermSource"`
}

// SearchTerms to return reports of SearchTerms
func (s *ReportService) SearchTerms(ctx context.Context, campaignID int, filter *ReportFilter) (*SearchTermsReport, *Response, error) {
	if campaignID == 0 {
		return nil, nil, fmt.Errorf("campaignID can not be 0")
	}
	u := fmt.Sprintf("reports/campaigns/%d/searchterms", campaignID)
	req, err := s.client.NewRequest("POST", u, filter)
	if err != nil {
		return nil, nil, err
	}
	report := new(SearchTermsReport)
	resp, err := s.client.Do(ctx, req, &report)
	if err != nil {
		return nil, resp, err
	}
	return report, resp, nil
}

// General Report data model stuff

type GrandTotals struct {
	Other bool       `json:"other"`
	Total Statistics `json:"total"`
}

type Statistics struct {
	Impressions    int     `json:"impressions"`
	Taps           int     `json:"taps"`
	Installs       int     `json:"installs"`
	NewDownloads   int     `json:"newDownloads"`
	Redownloads    int     `json:"redownloads"`
	LatOnInstalls  int     `json:"latOnInstalls"`
	LatOffInstalls int     `json:"latOffInstalls"`
	TTR            float64 `json:"ttr"`
	AvgCPA         Amount  `json:"avgCPA"`
	AvgCPT         Amount  `json:"avgCPT"`
	LocalSpend     Amount  `json:"localSpent"`
	ConversionRate float64 `json:"conversionRate"`
	Date           string  `json:"date,omitemtpy"`
}

type App struct {
	AppName string `json:"appName"`
	AdamID  int    `json:"adamId"`
}

type ReportFilter struct {
	StartTime                  string      `json:"startTime"`
	EndTime                    string      `json:"endTime"`
	TimeZone                   TimeZone    `json:"timeZone"`
	Granularity                Granularity `json:"granularity"`
	Selector                   Selector    `json:"selector"`
	GroupBy                    []string    `json:"groupBy"`
	ReturnRowTotals            bool        `json:"returnRowTotals"`
	ReturnGrandTotals          bool        `json:"returnGrandTotals"`
	ReturnRecordsWithNoMetrics bool        `json:"returnRecordsWithNoMetrics"`
}

type Selector struct {
	Conditions []Condition      `json:"conditions"`
	Fields     []string         `json:"fields"`
	OrderBy    []OrderBy        `json:"orderBy"`
	Pagination FilterPagination `json:"pagination"`
}

type FilterPagination struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type OrderBy struct {
	Field     string    `json:"field"`
	SortOrder SortOrder `json:"sortOrder"`
}

type Condition struct {
	Field    string   `json:"field"`
	Operator Operator `json:"operator"`
	Values   []string `json:"values"`
}
