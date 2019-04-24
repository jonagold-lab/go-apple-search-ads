package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/jonagold-lab/go-apple-search-ads/searchads"
)

func main() {
	orgID := int64(1405310)
	client, err := searchads.NewClient(nil, "../cert.pem", "../cert.key", &orgID)
	if err != nil {
		log.Fatalf("Client error: %s", err)
		panic(err)
	}

	data := searchads.Campaign{
		Name: "New Name",
	}

	editedCamapaign, _, err := client.Campaign.Edit(context.Background(), int64(259433851), &data)
	if err != nil {
		log.Fatalf("Campaign Edit error: %s", err)
		panic(err)
	}
	res, _ := json.Marshal(&editedCamapaign)
	fmt.Println(string(res))
}
