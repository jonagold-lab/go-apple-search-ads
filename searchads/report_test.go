package searchads

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestReportServive_Campaigns(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	filter := &ReportFilter{
		StartTime:   "2019-04-01",
		EndTime:     "2019-04-01",
		Granularity: HOURLY,
		Selector: Selector{
			OrderBy: []OrderBy{
				OrderBy{
					Field:     "countryOrRegion",
					SortOrder: ASCENDING,
				},
			},
			Conditions: []Condition{
				Condition{
					Field:    "countryOrRegion",
					Operator: CONTAINS_ANY,
					Values:   []string{"US", "GB"},
				},
			},
			Pagination: FilterPagination{
				Offset: 0,
				Limit:  1000,
			},
		},
		GroupBy:                    []string{"countryOrRegion"},
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
