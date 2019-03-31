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

func TestAdGroupService_List(t *testing.T) {
	client, mux, _, teardown := setup()
	t.Log("Setup Done")
	defer teardown()

	mux.HandleFunc("/campaigns/1234/adgroups", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		w.WriteHeader(http.StatusOK)
		w.Write(loadFixture("adgroups.json"))
	})
	opt := ListOptions{}
	got, _, err := client.AdGroup.List(context.Background(), 1234, &opt)
	if err != nil {
		t.Errorf("AdGroup.List returned error: %v", err)
	}

	want := []*AdGroup{}
	responseToInterface(loadFixture("adgroups.json"), &want)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("AdGroup.List = %+v, want %+v", got, want)
	}
}

func TestAdGroupService_Get(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	wantAcceptHeaders := []string{"application/json"}
	mux.HandleFunc("/campaigns/1234/adgroups/1234", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", strings.Join(wantAcceptHeaders, ", "))
		w.Write(loadFixture("adgroup.json"))
	})

	got, _, err := client.AdGroup.Get(context.Background(), 1234, 1234)
	if err != nil {
		t.Errorf("AdGroup.Get returned error: %v", err)
	}

	want := &AdGroup{}
	responseToInterface(loadFixture("adgroup.json"), want)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("AdGroup.Get returned %+v, want %+v", got, want)
	}
}

func TestAdGroupService_Create(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	input := &AdGroup{Name: "t"}

	wantAcceptHeaders := []string{"application/json"}
	mux.HandleFunc("/campaigns/1234/adgroups", func(w http.ResponseWriter, r *http.Request) {
		v := new(AdGroup)
		json.NewDecoder(r.Body).Decode(v)
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", strings.Join(wantAcceptHeaders, ", "))
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		fmt.Fprint(w, `{ "data": {"id":1} }`)
	})

	got, _, err := client.AdGroup.Create(context.Background(), 1234, input)
	if err != nil {
		t.Errorf("AdGroup.Create returned error: %v", err)
	}

	want := &AdGroup{ID: 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("AdGroup.Create returned %+v, want %+v", got, want)
	}
}

func TestAdGroupService_Edit(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	input := &AdGroup{Name: "New Name"}

	wantAcceptHeaders := []string{"application/json"}
	mux.HandleFunc("/campaigns/1234/adgroups/1234", func(w http.ResponseWriter, r *http.Request) {
		v := new(AdGroup)
		json.NewDecoder(r.Body).Decode(v)
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", strings.Join(wantAcceptHeaders, ", "))
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Request body = %+v, want %+v", v, input)
		}
		w.Write(loadFixture("adgroup_update.json"))
	})

	got, _, err := client.AdGroup.Edit(context.Background(), 1234, 1234, input)
	if err != nil {
		t.Errorf("AdGroup.Edit returned error: %v", err)
	}

	want := &AdGroup{}
	responseToInterface(loadFixture("adgroup_update.json"), want)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("AdGroup.Edit returned %+v, want %+v", got, want)
	}
}

func TestAdGroupService_Delete(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	wantAcceptHeaders := []string{"application/json"}
	mux.HandleFunc("/campaigns/1234/adgroups/1234", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		testHeader(t, r, "Accept", strings.Join(wantAcceptHeaders, ", "))
		w.WriteHeader(http.StatusOK)
		w.Write(loadFixture("adgroup_delete.json"))
	})

	resp, err := client.AdGroup.Delete(context.Background(), 1234, 1234)
	if err != nil {
		t.Errorf("AdGroup.Delete returned error: %v", err)
	}
	want := http.StatusOK
	got := resp.StatusCode
	if !reflect.DeepEqual(got, want) {
		t.Errorf("AdGroup.Delete returned %+v, want %+v", got, want)
	}
}
