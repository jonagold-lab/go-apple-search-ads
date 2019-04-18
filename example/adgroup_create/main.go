package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/jonagold-lab/go-apple-search-ads/searchads"
)

func main() {
	orgID := 1405310
	campaignID := 12312323
	client, err := searchads.NewClient(nil, "../cert.pem", "../cert.key", &orgID)
	if err != nil {
		log.Fatalf("Client error: %s", err)
		panic(err)
	}
	now := time.Now().UTC()
	startTime := fmt.Sprintf("%4d-%02d-%02dT%02d:%02d:%02d.000",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())
	data := searchads.AdGroup{
		CampaignID:             campaignID,
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
	}

	createdCamapaign, _, err := client.AdGroup.Create(context.Background(), campaignID, &data)
	if err != nil {
		log.Fatalf("Campaign Create error: %s", err)
		panic(err)
	}
	res, _ := json.Marshal(&createdCamapaign)
	fmt.Println(string(res))
}
