package searchads

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestReportService_Campaigns(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	filter := &ReportFilter{
		StartTime:   "2019-04-01",
		EndTime:     "2019-04-01",
		Granularity: HOURLY,
		Selector: Selector{
			OrderBy: []OrderBySelector{
				OrderBySelector{
					Field:     OrderByCountryOrRegion,
					SortOrder: ASCENDING,
				},
			},
			Conditions: []Condition{
				Condition{
					Field:    "countryOrRegion",
					Operator: OperatorContainsAny,
					Values:   []string{"US", "GB"},
				},
			},
			Pagination: PaginationSelector{
				Offset: 0,
				Limit:  1000,
			},
		},
		GroupBy:                    []GroupBy{GroupByCountryOrRegion},
		TimeZone:                   UTC,
		ReturnRowTotals:            true,
		ReturnGrandTotals:          true,
		ReturnRecordsWithNoMetrics: true,
	}

	wantAcceptHeaders := []string{"application/json"}
	mux.HandleFunc("/reports/campaigns", func(w http.ResponseWriter, r *http.Request) {
		v := new(ReportFilter)
		json.NewDecoder(r.Body).Decode(v)
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", strings.Join(wantAcceptHeaders, ", "))
		if !reflect.DeepEqual(v, filter) {
			t.Errorf("Request body = %+v, want %+v", v, filter)
		}
		w.Write(loadFixture("report_campaigns_hourly.json"))
	})

	got, _, err := client.Report.Campaigns(context.Background(), filter)

	if err != nil {
		t.Errorf("Report.Campaigns returned error: %v", err)
	}
	metadata := CampaignMetadata{
		CampaignID:     262806613,
		CampaignName:   "US_HighValue_JG",
		Deleted:        false,
		CampaignStatus: ENABLED,
		App: App{
			AppName: "Back Pain Relief - Kaia",
			AdamID:  1100673977,
		},
		ServingStatus:       RUNNING,
		ServingStateReasons: nil,
		CountriesOrRegions:  []CountryCode{US},
		ModificationTime:    "2019-03-20T23:58:51.455",
		TotalBudget: Amount{
			Amount:   "50000",
			Currency: "EUR",
		},
		DailyBudget: Amount{
			Amount:   "500",
			Currency: "EUR",
		},
		DisplayStatus:                      DS_RUNNING,
		OrgID:                              378940,
		CountryOrRegionServingStateReasons: map[string]string{},
		CountryOrRegion:                    US,
	}
	if !reflect.DeepEqual(got.ReportingDataResponse.Row[0].Metadata.App, metadata.App) {
		t.Errorf("Report.Campaigns First Row Metadata App returned %+v, want %+v", got.ReportingDataResponse.Row[0].Metadata.App, metadata.App)
	}
	if !reflect.DeepEqual(got.ReportingDataResponse.Row[0].Metadata.DailyBudget, metadata.DailyBudget) {
		t.Errorf("Report.Campaigns First Row Metadata DailyBudget returned %+v, want %+v", got.ReportingDataResponse.Row[0].Metadata.DailyBudget, metadata.DailyBudget)
	}
	if !reflect.DeepEqual(got.ReportingDataResponse.Row[0].Metadata.TotalBudget, metadata.TotalBudget) {
		t.Errorf("Report.Campaigns First Row Metadata TotalBudget returned %+v, want %+v", got.ReportingDataResponse.Row[0].Metadata.TotalBudget, metadata.TotalBudget)
	}
	if !reflect.DeepEqual(got.ReportingDataResponse.Row[0].Metadata, metadata) {
		t.Errorf("Report.Campaigns First Row Metadata returned %+v, want %+v", got.ReportingDataResponse.Row[0].Metadata, metadata)
	}
}

func TestReportService_AdGroups(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	filter := &ReportFilter{
		StartTime:   "2019-04-01",
		EndTime:     "2019-04-01",
		Granularity: HOURLY,
		Selector: Selector{
			OrderBy: []OrderBySelector{
				OrderBySelector{
					Field:     OrderByAdGroupID,
					SortOrder: ASCENDING,
				},
			},
			Conditions: []Condition{},
			Pagination: PaginationSelector{
				Offset: 0,
				Limit:  1000,
			},
		},
		GroupBy:                    []GroupBy{},
		TimeZone:                   UTC,
		ReturnRowTotals:            true,
		ReturnGrandTotals:          true,
		ReturnRecordsWithNoMetrics: true,
	}

	wantAcceptHeaders := []string{"application/json"}
	mux.HandleFunc("/reports/campaigns/262806613/adgroups", func(w http.ResponseWriter, r *http.Request) {
		v := new(ReportFilter)
		json.NewDecoder(r.Body).Decode(v)
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", strings.Join(wantAcceptHeaders, ", "))
		if !reflect.DeepEqual(v, filter) {
			t.Errorf("Request body = %+v, want %+v", v, filter)
		}
		w.Write(loadFixture("report_adgroups_hourly.json"))
	})

	got, _, err := client.Report.AdGroups(context.Background(), 262806613, filter)

	if err != nil {
		t.Errorf("Report.AdGroups returned error: %v", err)
	}
	metadata := AdGroupMetadata{
		AdGroupID:   262770380,
		AdGroupName: "Exact Match",
		StartTime:   "2019-03-21T00:00:00.000",
		EndTime:     nil,
		CpaGoal: Amount{
			Amount:   "10",
			Currency: "EUR",
		},
		DefaultCpcBid: Amount{
			Amount:   "12",
			Currency: "EUR",
		},
		Deleted:                    false,
		AdGroupStatus:              ENABLED,
		AdGroupServingStatus:       RUNNING,
		AdGroupServingStateReasons: nil,
		ModificationTime:           "2019-03-20T23:58:51.437",
		AdGroupDisplayStatus:       DS_RUNNING,
		AutomatedKeywordsOptIn:     false,
	}
	if !reflect.DeepEqual(got.ReportingDataResponse.Row[0].Metadata.CpaGoal, metadata.CpaGoal) {
		t.Errorf("Report.AdGroups First Row Metadata CpaGoal returned %+v, want %+v", got.ReportingDataResponse.Row[0].Metadata.CpaGoal, metadata.CpaGoal)
	}
	if !reflect.DeepEqual(got.ReportingDataResponse.Row[0].Metadata.DefaultCpcBid, metadata.DefaultCpcBid) {
		t.Errorf("Report.AdGroups First Row Metadata DefaultCpcBid returned %+v, want %+v", got.ReportingDataResponse.Row[0].Metadata.DefaultCpcBid, metadata.DefaultCpcBid)
	}
	if !reflect.DeepEqual(got.ReportingDataResponse.Row[0].Metadata, metadata) {
		t.Errorf("Report.AdGroups First Row Metadata returned %+v, want %+v", got.ReportingDataResponse.Row[0].Metadata, metadata)
	}
}

func TestReportService_SearchTerms(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	filter := &ReportFilter{
		StartTime:   "2019-03-01",
		EndTime:     "2019-04-01",
		Granularity: DAILY,
		Selector: Selector{
			OrderBy: []OrderBySelector{
				OrderBySelector{
					Field:     OrderByAdGroupID,
					SortOrder: ASCENDING,
				},
			},
			Conditions: []Condition{},
			Pagination: PaginationSelector{
				Offset: 0,
				Limit:  1000,
			},
		},
		GroupBy:                    []GroupBy{},
		TimeZone:                   UTC,
		ReturnRowTotals:            true,
		ReturnGrandTotals:          true,
		ReturnRecordsWithNoMetrics: true,
	}

	wantAcceptHeaders := []string{"application/json"}
	mux.HandleFunc("/reports/campaigns/262806613/searchterms", func(w http.ResponseWriter, r *http.Request) {
		v := new(ReportFilter)
		json.NewDecoder(r.Body).Decode(v)
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", strings.Join(wantAcceptHeaders, ", "))
		if !reflect.DeepEqual(v, filter) {
			t.Errorf("Request body = %+v, want %+v", v, filter)
		}
		w.Write(loadFixture("report_searchterms_daily.json"))
	})

	got, _, err := client.Report.SearchTerms(context.Background(), 262806613, filter)

	if err != nil {
		t.Errorf("Report.SearchTerms returned error: %v", err)
	}
	metadata := SearchTermMetadata{
		KeywordID: 262823190,
		Keyword:   "back pain",
		MatchType: MatchTypeExact,
		BidAmount: Amount{
			Amount:   "12",
			Currency: "EUR",
		},
		KeywordDisplayStatus: KeywordDisplayStatusRunning,
		Deleted:              false,
		AdGroupID:            262770380,
		AdGroupName:          "Exact Match",
		AdGroupDeleted:       false,
		SearchTermText:       nil,
		SearchTermSource:     SearchTermSourceTargeted,
	}

	if !reflect.DeepEqual(got.ReportingDataResponse.Row[0].Metadata.BidAmount, metadata.BidAmount) {
		t.Errorf("Report.SearchTerms First Row Metadata BidAmount returned %+v, want %+v", got.ReportingDataResponse.Row[0].Metadata.BidAmount, metadata.BidAmount)
	}
	if !reflect.DeepEqual(got.ReportingDataResponse.Row[0].Metadata, metadata) {
		t.Errorf("Report.SearchTerms First Row Metadata returned %+v, want %+v", got.ReportingDataResponse.Row[0].Metadata, metadata)
	}
}

func TestReportService_Keywords(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	filter := &ReportFilter{
		StartTime:   "2019-03-01",
		EndTime:     "2019-04-01",
		Granularity: DAILY,
		Selector: Selector{
			OrderBy: []OrderBySelector{
				OrderBySelector{
					Field:     OrderByAdGroupID,
					SortOrder: ASCENDING,
				},
			},
			Conditions: []Condition{},
			Pagination: PaginationSelector{
				Offset: 0,
				Limit:  1000,
			},
		},
		GroupBy:                    []GroupBy{},
		TimeZone:                   UTC,
		ReturnRowTotals:            true,
		ReturnGrandTotals:          true,
		ReturnRecordsWithNoMetrics: true,
	}

	wantAcceptHeaders := []string{"application/json"}
	mux.HandleFunc("/reports/campaigns/262806613/keywords", func(w http.ResponseWriter, r *http.Request) {
		v := new(ReportFilter)
		json.NewDecoder(r.Body).Decode(v)
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", strings.Join(wantAcceptHeaders, ", "))
		if !reflect.DeepEqual(v, filter) {
			t.Errorf("Request body = %+v, want %+v", v, filter)
		}
		w.Write(loadFixture("report_keywords_daily.json"))
	})

	got, _, err := client.Report.Keywords(context.Background(), 262806613, filter)

	if err != nil {
		t.Errorf("Report.Keywords returned error: %v", err)
	}
	metadata := KeywordMetadata{
		KeywordID:     262823190,
		Keyword:       "back pain",
		KeywordStatus: KEYWORD_ACTIVE,
		MatchType:     MatchTypeExact,
		BidAmount: Amount{
			Amount:   "12",
			Currency: "EUR",
		},
		Deleted:              false,
		KeywordDisplayStatus: KeywordDisplayStatusRunning,
		AdGroupID:            262770380,
		AdGroupName:          "Exact Match",
		AdGroupDeleted:       false,
		ModificationTime:     "2019-03-20T23:06:51.606",
	}

	if !reflect.DeepEqual(got.ReportingDataResponse.Row[0].Metadata.BidAmount, metadata.BidAmount) {
		t.Errorf("Report.Keywords First Row Metadata BidAmount returned %+v, want %+v", got.ReportingDataResponse.Row[0].Metadata.BidAmount, metadata.BidAmount)
	}
	if !reflect.DeepEqual(got.ReportingDataResponse.Row[0].Metadata, metadata) {
		t.Errorf("Report.Keywords First Row Metadata returned %+v, want %+v", got.ReportingDataResponse.Row[0].Metadata, metadata)
	}
}
