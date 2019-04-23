package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/jonagold-lab/go-apple-search-ads/searchads"
)

func main() {
	orgID := 1461850
	adamID := 1222530780
	client, err := searchads.NewClient(nil, "../api-keys.pem", "../api-keys.key", &orgID)
	if err != nil {
		log.Fatalf("Client error: %s", err)
		panic(err)
	}

	data := searchads.Campaign{
		Name:         "US_BRAND_EXACT_2",
		OrgID:        orgID,
		AdamID:       adamID,
		PaymentModel: searchads.PAYG,
		BudgetAmount: searchads.Amount{
			Amount:   "10000",
			Currency: "EUR",
		},
		DailyBudgetAmount: searchads.Amount{
			Amount:   "50",
			Currency: "EUR",
		},
		CountriesOrRegions: []searchads.CountryCode{
			searchads.US,
		},
	}

	createdCamapaign, _, err := client.Campaign.Create(context.Background(), &data)
	if err != nil {
		datab, _ := json.Marshal(&data)
		log.Fatalf("Campaign Create error: %s, Data: %v", err, string(datab))
		panic(err)
	}
	res, _ := json.Marshal(&createdCamapaign)
	fmt.Println(string(res))
}
