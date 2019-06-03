package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/jonagold-lab/go-apple-search-ads/searchads"
)

func main() {
	campaignID := int64(262922238)
	adGroupID := int64(262982513)
	targetingKeywordID := int64(306708156)
	client, err := searchads.NewClient(nil, "../cert.pem", "../cert.key", nil)
	if err != nil {
		log.Fatalf("Client error: %s", err)
		panic(err)
	}
	k, rs, err := client.AdGroupTargetingKeyword.Get(context.Background(), campaignID, adGroupID, targetingKeywordID)
	if err != nil {
		log.Fatalf("TargetingKeyword error: %s", err)
		panic(err)
	}
	res, _ := json.Marshal(&k)
	fmt.Println(string(res))
	fmt.Println("----------------")
	fmt.Println(rs.Pagination.ItemsPerPage)
	fmt.Println(rs.Pagination.StartIndex)
	fmt.Println(rs.Pagination.TotalResults)
}
