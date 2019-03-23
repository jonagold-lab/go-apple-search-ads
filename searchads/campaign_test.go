package searchads

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestCampaignService_List(t *testing.T) {
	client, mux, _, teardown := setup()
	t.Log("Setup Done")
	defer teardown()

	mux.HandleFunc("/campaigns", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		w.WriteHeader(http.StatusOK)
		w.Write(loadFixture("campaigns.json"))
	})
	opt := ListOptions{}
	got, _, err := client.Campaign.List(context.Background(), &opt)
	if err != nil {
		t.Errorf("Campaign.List returned error: %v", err)
	}

	want := []*Campaign{}
	responseToInterface(loadFixture("campaigns.json"), &want)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Campaign.List = %+v, want %+v", got, want)
	}
}

func TestCampaignService_Get(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	wantAcceptHeaders := []string{"application/json"}
	mux.HandleFunc("/campaigns/235557343", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(wantAcceptHeaders, ", "))
		w.Write(loadFixture("campaign.json"))
	})

	got, _, err := client.Campaign.Get(context.Background(), 235557343)
	if err != nil {
		t.Errorf("Campaign.Get returned error: %v", err)
	}

	want := &Campaign{}
	responseToInterface(loadFixture("campaign.json"), want)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Campaign.Get returned %+v, want %+v", got, want)
	}
}

func TestCampaignService_Create(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	input := &NewCampaign{Name: "t"}

	wantAcceptHeaders := []string{"application/json"}
	mux.HandleFunc("/campaigns", func(w http.ResponseWriter, r *http.Request) {
		v := new(NewCampaign)
		json.NewDecoder(r.Body).Decode(v)
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", strings.Join(wantAcceptHeaders, ", "))
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		fmt.Fprint(w, `{ "data": {"id":1} }`)
	})

	got, _, err := client.Campaign.Create(context.Background(), input)
	if err != nil {
		t.Errorf("Campaign.Create returned error: %v", err)
	}

	want := &Campaign{ID: 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Campaign.Create returned %+v, want %+v", got, want)
	}
}

func TestCampaignService_Edit(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	input := &UpdateCampaign{Name: "New Name"}

	wantAcceptHeaders := []string{"application/json"}
	mux.HandleFunc("/campaigns/235557343", func(w http.ResponseWriter, r *http.Request) {
		v := new(UpdateCampaign)
		json.NewDecoder(r.Body).Decode(v)
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", strings.Join(wantAcceptHeaders, ", "))
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write(loadFixture("campaign_update.json"))
	})

	got, _, err := client.Campaign.Edit(context.Background(), 235557343, input)
	if err != nil {
		t.Errorf("Campaign.Edit returned error: %v", err)
	}

	want := &Campaign{}
	responseToInterface(loadFixture("campaign_update.json"), want)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Campaign.Edit returned %+v, want %+v", got, want)
	}
}

func TestCampaignService_Delete(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	wantAcceptHeaders := []string{"application/json"}
	mux.HandleFunc("/campaigns/235557343", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		testHeader(t, r, "Accept", strings.Join(wantAcceptHeaders, ", "))
		w.WriteHeader(http.StatusOK)
	})

	resp, err := client.Campaign.Delete(context.Background(), 235557343)
	if err != nil {
		t.Errorf("Campaign.Edit returned error: %v", err)
	}
	want := http.StatusOK
	got := resp.StatusCode
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Campaign.Edit returned %+v, want %+v", got, want)
	}
}
