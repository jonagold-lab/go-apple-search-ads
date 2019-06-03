package searchads

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestAdGroupTargetingKeywordService_List(t *testing.T) {
	client, mux, _, teardown := setup()
	t.Log("Setup Done")
	defer teardown()

	mux.HandleFunc("/campaigns/1234/adgroups/1234/targetingkeywords", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.WriteHeader(http.StatusOK)
		w.Write(loadFixture("adgroup_targeting_keywords.json"))
	})
	opt := ListOptions{}
	got, _, err := client.AdGroupTargetingKeyword.List(context.Background(), 1234, 1234, &opt)
	if err != nil {
		t.Errorf("AdGroupTargetingKeyword.List returned error: %v", err)
	}

	want := []*TargetingKeyword{}
	responseToInterface(loadFixture("adgroup_targeting_keywords.json"), &want)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("AdGroupTargetingKeyword.List = %+v, want %+v", got, want)
	}
}
func TestAdGroupTargetingKeywordService_CreateBulk(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	nk := TargetingKeyword{
		AdGroupID: 1234,
		Text:      "i do negative keywords",
		MatchType: MatchTypeExact,
		Status:    KEYWORD_ACTIVE,
		BidAmount: Amount{
			Amount:   "1.50",
			Currency: "EUR",
		},
	}

	input := []*TargetingKeyword{&nk}

	wantAcceptHeaders := []string{"application/json"}
	mux.HandleFunc("/campaigns/1234/adgroups/1234/targetingkeywords/bulk", func(w http.ResponseWriter, r *http.Request) {
		v := []*TargetingKeyword{}
		json.NewDecoder(r.Body).Decode(&v)
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", strings.Join(wantAcceptHeaders, ", "))
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write(loadFixture("adgroup_targeting_keyword_bulk.json"))
	})

	got, _, err := client.AdGroupTargetingKeyword.CreateBulk(context.Background(), 1234, 1234, input)

	if err != nil {
		t.Errorf("AdGroupTargetingKeyword.CreateBulk returned error: %v", err)
	}
	want := []*TargetingKeyword{
		&TargetingKeyword{
			ID:        1,
			AdGroupID: 1234,
			Text:      "i do negative keywords",
			MatchType: MatchTypeExact,
			Status:    KEYWORD_ACTIVE,
			BidAmount: Amount{
				Amount:   "1.50",
				Currency: "EUR",
			},
			ModificationTime: "2019-02-22T15:25:46.851",
			Deleted:          false,
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("AdGroupTargetingKeyword.CreateBulk returned %+v, want %+v", got, want)
	}
}
