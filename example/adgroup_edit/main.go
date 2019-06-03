package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/jonagold-lab/go-apple-search-ads/searchads"
)

func main() {
	orgID := int64(378940)
	adGroupID := int64(285227483)
	campaignID := int64(284979913)
	client, err := searchads.NewClient(nil, "../cert.pem", "../cert.key", &orgID)
	if err != nil {
		log.Fatalf("Client error: %s", err)
		panic(err)
	}
	status := searchads.PAUSED
	data := searchads.AdGroup{
		Status: &status,
	}

	uag, _, err := client.AdGroup.Edit(context.Background(), campaignID, adGroupID, &data)
	if err != nil {
		log.Fatalf("Campaign Create error: %s", err)
		panic(err)
	}
	res, _ := json.Marshal(&uag)
	fmt.Println(string(res))
}
