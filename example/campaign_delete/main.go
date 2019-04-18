package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/jonagold-lab/go-apple-search-ads/searchads"
)

func main() {
	orgID := 1405310
	client, err := searchads.NewClient(nil, "../cert.pem", "../cert.key", &orgID)
	if err != nil {
		log.Fatalf("Client error: %s", err)
		panic(err)
	}

	resp, err := client.Campaign.Delete(context.Background(), 259428355)
	if err != nil {
		log.Fatalf("Campaign Delete error: %s", err)
		panic(err)
	}
	res, _ := json.Marshal(resp)
	fmt.Println(string(res))
}
