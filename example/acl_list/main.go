package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/maier-stefan/go-apple-search-ads/searchads"
)

func main() {
	client, err := searchads.NewClient(nil, "../cert.pem", "../cert.key", nil)
	if err != nil {
		log.Fatalf("Client error: %s", err)
		panic(err)
	}
	opt := searchads.ListOptions{Limit: 1000, Offset: 0}
	list, rs, err := client.ACL.List(context.Background(), &opt)
	if err != nil {
		log.Fatalf("ACL List error: %s", err)
		panic(err)
	}
	res, _ := json.Marshal(&list)
	fmt.Println(string(res))
	fmt.Println("----------------")
	fmt.Println(len(list))
	fmt.Println(rs.Pagination.ItemsPerPage)
	fmt.Println(rs.Pagination.StartIndex)
	fmt.Println(rs.Pagination.TotalResults)
}
