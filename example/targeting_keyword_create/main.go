package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/jonagold-lab/go-apple-search-ads/searchads"
)

func main() {
	campaignID := int64(284979913)
	adGroupID := int64(285227483)
	client, err := searchads.NewClient(nil, "../cert.pem", "../cert.key", nil)
	if err != nil {
		log.Fatalf("Client error: %s", err)
		panic(err)
	}
	input := []*searchads.TargetingKeyword{
		&searchads.TargetingKeyword{
			AdGroupID: adGroupID,
			MatchType: searchads.MatchTypeExact,
			Status:    searchads.KEYWORD_PAUSED,
			Text:      "test 2",
		},
	}
	createdKeyword, rs, err := client.AdGroupTargetingKeyword.CreateBulk(context.Background(), campaignID, adGroupID, input)
	if err != nil {
		log.Fatalf("TargetingKeyword error: %s", err)
		panic(err)
	}
	res, _ := json.Marshal(&createdKeyword)
	fmt.Println(string(res))
	fmt.Println("----------------")
	fmt.Println(rs.Pagination.ItemsPerPage)
	fmt.Println(rs.Pagination.StartIndex)
	fmt.Println(rs.Pagination.TotalResults)
}
