package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/maier-stefan/go-apple-search-ads/searchads"
)

func main() {
	orgID := 1405310
	client, err := searchads.NewClient(nil, "../cert.pem", "../cert.key", &orgID)
	if err != nil {
		log.Fatalf("Client error: %s", err)
		panic(err)
	}
	now := time.Now().UTC()
	startTime := fmt.Sprintf("%4d-%02d-%02dT%02d:%02d:%02d.000",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())
	adGroups := []searchads.AdGroup{
		searchads.AdGroup{
			Name:                   "US_BRAND_EXACT_2",
			StartTime:              startTime,
			AutomatedKeywordsOptIn: false,
			CpaGoal: searchads.Amount{
				Amount:   "5",
				Currency: "USD",
			},
			DefaultCpcBid: searchads.Amount{
				Amount:   "3",
				Currency: "USD",
			},
		},
	}
	us := searchads.US
	data := searchads.Campaign{
		Name:   "US_BRAND_EXACT_2",
		OrgID:  orgID,
		AdamID: 1112839581,
		BudgetAmount: searchads.Amount{
			Amount:   "10000",
			Currency: "USD",
		},
		DailyBudgetAmount: searchads.Amount{
			Amount:   "50",
			Currency: "USD",
		},
		Storefront: us,
		AdGroups:   adGroups,
	}

	createdCamapaign, _, err := client.Campaign.Create(context.Background(), &data)
	if err != nil {
		log.Fatalf("Campaign Create error: %s", err)
		panic(err)
	}
	res, _ := json.Marshal(&createdCamapaign)
	fmt.Println(string(res))
}
