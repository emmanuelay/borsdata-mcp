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
	shortsProc := -2.5
	shortsHolders := 3.0
	expected := models.ShortHoldingsResponse{
		List: []models.ShortsRespV1{
			{
				InsID:         1,
				ShortsProc:    &shortsProc,
				ShortsHolders: &shortsHolders,
				Trend1w:       5.0,
				Trend1m:       10.0,
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
	if result.List[0].ShortsProc == nil || *result.List[0].ShortsProc != -2.5 {
		t.Errorf("expected shortsProc=-2.5, got %v", result.List[0].ShortsProc)
	}
	if result.List[0].ShortsHolders == nil || *result.List[0].ShortsHolders != 3.0 {
		t.Errorf("expected shortsHolders=3.0, got %v", result.List[0].ShortsHolders)
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
