package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/jonagold-lab/go-apple-search-ads/searchads"
)

func main() {
	campaignID := int64(262922238)
	adGroupID := int64(262982513)
	pemdat, _ := ioutil.ReadFile("../cert.pem")
	keydat, _ := ioutil.ReadFile("../cert.key")
	client, err := searchads.NewClient(nil, pemdat, keydat, nil)
	if err != nil {
		log.Fatalf("Client error: %s", err)
		panic(err)
	}
	status, _ := searchads.ParseKeywordStatus("ACTIVE")
	matchType, _ := searchads.ParseMatchType("EXACT")
	input := []*searchads.TargetingKeyword{
		&searchads.TargetingKeyword{
			AdGroupID: adGroupID,
			Text:      "withings health mate",
			Status:    status,
			BidAmount: searchads.Amount{
				Amount:   fmt.Sprintf("%f", 3.0),
				Currency: "EUR",
			},
			MatchType: matchType,
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
