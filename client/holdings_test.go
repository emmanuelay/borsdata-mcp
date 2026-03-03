package client

import (
	"testing"

	"github.com/emmanuelay/borsdata-mcp/models"
)

func TestGetInsiderHoldings(t *testing.T) {
	expected := models.InsiderHoldingsResponse{
		List: []models.InsiderRespV1{
			{
				InsID: 1,
				Values: []models.InsiderRowV1{
					{OwnerName: "John Doe", OwnerPosition: "CEO", Shares: 1000, Price: 50.0, Amount: 50000, Currency: "SEK", TransactionType: 1, TransactionDate: "2025-01-10"},
				},
			},
		},
	}
	log, c := newTestServer(t, map[string]any{
		"/v1/holdings/insider": expected,
	})

	result, err := c.GetInsiderHoldings("1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.List) != 1 {
		t.Fatalf("expected 1 insider resp, got %d", len(result.List))
	}
	if result.List[0].Values[0].OwnerName != "John Doe" {
		t.Errorf("expected John Doe, got %s", result.List[0].Values[0].OwnerName)
	}
	log.assertParam(t, "/v1/holdings/insider", "authKey", "test-key")
	log.assertParam(t, "/v1/holdings/insider", "instList", "1")
}

func TestGetShortHoldings(t *testing.T) {
	expected := models.ShortHoldingsResponse{
		List: []models.ShortsRespV1{
			{
				InsID: 1,
				Values: []models.ShortsRowV1{
					{ReportID: 100, PositionHolderName: "HedgeFund AB", ISIN: "SE0001234567", ShortingPercent: 2.5, NotifiedCapital: 1000000, Currency: "SEK", ReportDate: "2025-01-15"},
				},
			},
		},
	}
	log, c := newTestServer(t, map[string]any{
		"/v1/holdings/shorts": expected,
	})

	result, err := c.GetShortHoldings()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.List) != 1 {
		t.Fatalf("expected 1 shorts resp, got %d", len(result.List))
	}
	if result.List[0].Values[0].PositionHolderName != "HedgeFund AB" {
		t.Errorf("expected HedgeFund AB, got %s", result.List[0].Values[0].PositionHolderName)
	}
	log.assertParam(t, "/v1/holdings/shorts", "authKey", "test-key")
}

func TestGetBuybackHoldings(t *testing.T) {
	expected := models.BuybackHoldingsResponse{
		List: []models.BuybackRespV1{
			{
				InsID: 1,
				Values: []models.BuybackRowV1{
					{Change: 500, ChangeProc: 0.1, Price: 100.0, Currency: "SEK", Shares: 5000, SharesProc: 0.5, Date: "2025-01-12"},
				},
			},
		},
	}
	log, c := newTestServer(t, map[string]any{
		"/v1/holdings/buyback": expected,
	})

	result, err := c.GetBuybackHoldings("1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.List) != 1 {
		t.Fatalf("expected 1 buyback resp, got %d", len(result.List))
	}
	if result.List[0].Values[0].Change != 500 {
		t.Errorf("expected change 500, got %d", result.List[0].Values[0].Change)
	}
	log.assertParam(t, "/v1/holdings/buyback", "authKey", "test-key")
	log.assertParam(t, "/v1/holdings/buyback", "instList", "1")
}
