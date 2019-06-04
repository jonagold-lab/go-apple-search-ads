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
	opt := searchads.ListOptions{Limit: 1000, Offset: 0}
	k, rs, err := client.AdGroupTargetingKeyword.List(context.Background(), campaignID, adGroupID, &opt)
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
