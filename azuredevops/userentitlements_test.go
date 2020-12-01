package azuredevops_test

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/mcdafydd/go-azuredevops/azuredevops"
)

func Test_UserEntitlementsGet(t *testing.T) {
	c, mux, _, teardown := setup()
	defer teardown()
	u, _ := url.Parse("")
	c.VsaexBaseURL = *u
	mux.HandleFunc("/orgName/_apis/userentitlements", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{
				"members": [
					{
						"id": "6416203b-98bb-4910-8f8a-b12aa19a399f"
					}
				],
				"continuationToken": null,
				"totalCount": 0,
				"items": [
					{
						"id": "6416203b-98bb-4910-8f8a-b12aa19a399f"
					}
				]
		}`)
	})

	got, _, err := c.UserEntitlements.Get(context.Background(), "userName", "orgName")
	if err != nil {
		t.Fatalf("returned error: %v", err)
	}

	want := &azuredevops.UserEntitlements{}
	item := azuredevops.Item{
		ID: "6416203b-98bb-4910-8f8a-b12aa19a399f",
	}
	items := []azuredevops.Item{item}
	want.Members = items
	want.Items = items
	want.TotalCount = 0

	if !cmp.Equal(got, want) {
		diff := cmp.Diff(got, want)
		fmt.Printf(diff)
		t.Errorf("UserEntitlements.Get returned %+v, want %+v", got, want)
	}
}

func Test_UserEntitlementsGetUserId(t *testing.T) {
	c, mux, _, teardown := setup()
	defer teardown()
	u, _ := url.Parse("")
	c.VsaexBaseURL = *u
	mux.HandleFunc("/orgName/_apis/userentitlements", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{
				"members": [
					{
						"id": "6416203b-98bb-4910-8f8a-b12aa19a399f"
					}
				],
				"continuationToken": null,
				"totalCount": 0,
				"items": [
					{
						"id": "6416203b-98bb-4910-8f8a-b12aa19a399f"
					}
				]
		}`)
	})

	got, err := c.UserEntitlements.GetUserID(context.Background(), "userName", "orgName")
	if err != nil {
		t.Fatalf("returned error: %v", err)
	}

	want := "6416203b-98bb-4910-8f8a-b12aa19a399f"

	if !cmp.Equal(got, &want) {
		diff := cmp.Diff(got, &want)
		fmt.Printf(diff)
		t.Errorf("UserEntitlements.GetUserId returned %+v, want %+v", *got, want)
	}
}
