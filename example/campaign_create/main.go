package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/maier-stefan/go-apple-search-ads/searchads"
)

func main() {
	orgID := 1405310
	adamID := 1112839581
	client, err := searchads.NewClient(nil, "../cert.pem", "../cert.key", &orgID)
	if err != nil {
		log.Fatalf("Client error: %s", err)
		panic(err)
	}

	data := searchads.Campaign{
		Name:   "US_BRAND_EXACT_2",
		OrgID:  orgID,
		AdamID: adamID,
		BudgetAmount: searchads.Amount{
			Amount:   "10000",
			Currency: "USD",
		},
		DailyBudgetAmount: searchads.Amount{
			Amount:   "50",
			Currency: "USD",
		},
		CountriesOrRegions: []searchads.CountryCode{
			searchads.US,
		},
	}

	createdCamapaign, _, err := client.Campaign.Create(context.Background(), &data)
	if err != nil {
		log.Fatalf("Campaign Create error: %s", err)
		panic(err)
	}
	res, _ := json.Marshal(&createdCamapaign)
	fmt.Println(string(res))
}
