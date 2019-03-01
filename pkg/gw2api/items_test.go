package gw2api

import "testing"

func TestItems(t *testing.T) {
	var err error
	api := NewGW2Api()

	var testItems []int
	if testItems, err = api.Items(); err != nil {
		t.Error("Failed to fetch items")
	}

	if _, err = api.ItemDetails(-1, 0, "en"); err == nil {
		t.Error("Failed to fetch error for impossible arguments")
	}
	if _, err = api.ItemPages(-1, 0, "en"); err == nil {
		t.Error("Failed to fetch error for impossible arguments")
	}

	var items []Item
	if items, err = api.ItemDetails(0, 2, "en"); err != nil {
		t.Error("Failed to parse the item data: ", err)
	} else if len(items) < 1 {
		t.Error("Failed to fetch existing items")
	}
	if items, err = api.ItemDetails(0, 0, "en", testItems[0:2]...); err != nil {
		t.Error("Failed to parse the item data: ", err)
	} else if len(items) != 2 {
		t.Error("Failed to fetch existing items")
	}
}
