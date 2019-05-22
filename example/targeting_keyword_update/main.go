package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/jonagold-lab/go-apple-search-ads/searchads"
)

func main() {
	campaignID := int64(262727835)
	adGroupID := int64(262765077)
	targetingKeywordID := int64(262765077)
	client, err := searchads.NewClient(nil, "../cert.pem", "../cert.key", nil)
	if err != nil {
		log.Fatalf("Client error: %s", err)
		panic(err)
	}
	input := &searchads.TargetingKeyword{
		ID:        targetingKeywordID,
		AdGroupID: adGroupID,
		Status:    searchads.KEYWORD_PAUSED,
	}
	updatedKeyword, rs, err := client.AdGroupTargetingKeyword.Update(context.Background(), campaignID, adGroupID, targetingKeywordID, input)
	if err != nil {
		log.Fatalf("TargetingKeyword error: %s", err)
		panic(err)
	}
	res, _ := json.Marshal(&updatedKeyword)
	fmt.Println(string(res))
	fmt.Println("----------------")
	fmt.Println(rs.Pagination.ItemsPerPage)
	fmt.Println(rs.Pagination.StartIndex)
	fmt.Println(rs.Pagination.TotalResults)
}
