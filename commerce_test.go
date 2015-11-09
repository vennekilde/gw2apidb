package gw2api

import (
	"os"
	"testing"
)

func TestCommerceListings(t *testing.T) {
	var err error
	api := NewGW2Api()

	var testCommerceListings []int
	if testCommerceListings, err = api.CommerceListings(); err != nil {
		t.Error("Failed to fetch listings")
	}

	var listings []ArticleListing
	if listings, err = api.CommerceListingIds(testCommerceListings[0], testCommerceListings[1]); err != nil {
		t.Error("Failed to parse the listing data: ", err)
	} else if len(listings) != 2 {
		t.Error("Failed to fetch existing listings")
	}
	if listings, err = api.CommerceListingPages(0, 2); err != nil {
		t.Error("Failed to parse the listing data: ", err)
	} else if len(listings) != 2 {
		t.Error("Failed to fetch existing listings")
	}
}

func TestCommerceExchange(t *testing.T) {
	var err error
	api := NewGW2Api()

	var ex Exchange
	if ex, err = api.CommerceExchangeGems(100); err != nil || int(ex.CoinsPerGem) > 0 {
		t.Error("Failed to fetch gem exchange rate")
	}

	if ex, err = api.CommerceExchangeCoins(10000); err != nil || int(ex.CoinsPerGem) > 0 {
		t.Error("Failed to fetch coin exchange rate")
	}
}

func TestCommercePrices(t *testing.T) {
	var err error
	api := NewGW2Api()

	var testCommercePrices []int
	if testCommercePrices, err = api.CommercePrices(); err != nil {
		t.Error("Failed to fetch prices")
	}

	var prices []ArticlePrice
	if prices, err = api.CommercePriceIds(testCommercePrices[0], testCommercePrices[1]); err != nil {
		t.Error("Failed to parse the listing data: ", err)
	} else if len(prices) != 2 {
		t.Error("Failed to fetch existing prices")
	}
}

func TestCommerceTransactions(t *testing.T) {
	var apikey string
	if apikey = os.Getenv("APIKEY"); len(apikey) < 1 {
		t.Skip("Cannot test without APIKEY")
	}
	var api *GW2Api
	var err error
	if api, err = NewAuthenticatedGW2Api(apikey); err != nil {
		t.Error(err)
	}
	if !api.HasPermission(PermTradingpost) {
		t.Skip("API-Key does not have required permission for the test")
	}

	if _, err = api.CommerceTransactionsCurrentBuys(); err != nil {
		t.Error("Failed parsing current buys")
	}

	if _, err = api.CommerceTransactionsCurrentSells(); err != nil {
		t.Error("Failed parsing current sells")
	}

	if _, err = api.CommerceTransactionsHistoryBuys(); err != nil {
		t.Error("Failed parsing history buys")
	}

	if _, err = api.CommerceTransactionsHistorySells(); err != nil {
		t.Error("Failed parsing history sells")
	}
}
