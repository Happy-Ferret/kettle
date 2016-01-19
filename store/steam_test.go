package store_test

import (
	"net/url"
	"testing"

	"kettle/store"
)

var storeApi *store.Store

func init() {
	storeApi = steamstoreapi.NewSteamStoreApi()
}

func Test_AppDetails(t *testing.T) {
	details, err := storeApi.GetAppDetails(49520, url.Values{})
	if err != nil {
		t.Errorf("Getting app data failed: %s", err.Error())
	}

	if !details["49520"].Success {
		t.Fatal("Expected a successful hit")
	}

	if details["49520"].Data.Name != "Borderlands 2" {
		t.Fatalf("Expected title Borderlands 2 got, %s", details["49520"].Data.Name)
	}
}