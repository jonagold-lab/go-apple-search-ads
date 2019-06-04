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
	orgID := int64(1405310)
	pemdat, _ := ioutil.ReadFile("../cert.pem")
	keydat, _ := ioutil.ReadFile("../cert.key")
	client, err := searchads.NewClient(nil, pemdat, keydat, &orgID)
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
