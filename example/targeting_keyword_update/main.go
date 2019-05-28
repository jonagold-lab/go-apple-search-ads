package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/jonagold-lab/go-apple-search-ads/searchads"
)

func main() {
	campaignID := int64(262931515)
	adGroupID := int64(262987552)
	targetingKeywordID := int64(263022253)
	client, err := searchads.NewClient(nil, "../cert.pem", "../cert.key", nil)
	if err != nil {
		log.Fatalf("Client error: %s", err)
		panic(err)
	}
	input := []*searchads.TargetingKeyword{
		&searchads.TargetingKeyword{
			ID:        targetingKeywordID,
			AdGroupID: adGroupID,
			Status:    "ACTIVE",
		},
	}
	updatedKeywords, rs, err := client.AdGroupTargetingKeyword.UpdateBulk(context.Background(), campaignID, adGroupID, input)
	if err != nil {
		log.Fatalf("TargetingKeyword error: %s", err)
		panic(err)
	}
	res, _ := json.Marshal(&updatedKeywords)
	fmt.Println(string(res))
	fmt.Println("----------------")
	fmt.Println(rs.Pagination.ItemsPerPage)
	fmt.Println(rs.Pagination.StartIndex)
	fmt.Println(rs.Pagination.TotalResults)
}
