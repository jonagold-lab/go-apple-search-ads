package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/jonagold-lab/go-apple-search-ads/searchads"
)

func main() {
	orgID := int64(378940)
	client, err := searchads.NewClient(nil, "../cert.pem", "../cert.key", &orgID)
	if err != nil {
		log.Fatalf("Client error: %s", err)
		panic(err)
	}
	filter := searchads.ReportFilter{
		StartTime:   "2019-04-01",
		EndTime:     "2019-04-05",
		Granularity: searchads.HOURLY,
		Selector: searchads.Selector{
			OrderBy: []searchads.OrderBySelector{
				searchads.OrderBySelector{
					Field:     searchads.OrderByCountryOrRegion,
					SortOrder: searchads.ASCENDING,
				},
			},
			Conditions: []searchads.Condition{
				searchads.Condition{
					Field:    "countriesOrRegions",
					Operator: searchads.OperatorContainsAny,
					Values:   []string{"US", "GB"},
				},
			},
			Pagination: searchads.PaginationSelector{
				Offset: 0,
				Limit:  1000,
			},
		},
		GroupBy: []searchads.GroupBy{
			searchads.GroupByCountryOrRegion,
		},
		TimeZone:                   searchads.UTC,
		ReturnRecordsWithNoMetrics: true,
		ReturnRowTotals:            true,
		ReturnGrandTotals:          true,
	}
	report, rs, err := client.Report.Campaigns(context.Background(), &filter)
	if err != nil {
		log.Fatalf("Campaign Reports error: %s", err)
		panic(err)
	}
	res, _ := json.Marshal(&report)
	fmt.Println(string(res))
	fmt.Println("----------------")
	fmt.Println(report.ReportingDataResponse.Row)
	fmt.Println(rs.Pagination.ItemsPerPage)
	fmt.Println(rs.Pagination.StartIndex)
	fmt.Println(rs.Pagination.TotalResults)
}
