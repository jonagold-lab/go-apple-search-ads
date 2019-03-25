package searchads

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestACLService_List(t *testing.T) {
	client, mux, _, teardown := setup()
	t.Log("Setup Done")
	defer teardown()

	mux.HandleFunc("/acls", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		w.WriteHeader(http.StatusOK)
		w.Write(loadFixture("acls.json"))
	})
	opt := ListOptions{}
	got, _, err := client.ACL.List(context.Background(), &opt)
	if err != nil {
		t.Errorf("ACL.List returned error: %v", err)
	}

	want := []*ACL{}
	responseToInterface(loadFixture("acls.json"), &want)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ACL.List = %+v, want %+v", got, want)
	}
}
