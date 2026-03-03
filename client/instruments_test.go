package client

import (
	"testing"

	"github.com/emmanuelay/borsdata-mcp/models"
)

func TestGetInstruments(t *testing.T) {
	expected := models.InstrumentsResponse{
		Instruments: []models.InstrumentV1{
			{InsID: 1, Name: "TestCompany", Ticker: "TEST", MarketID: 1, SectorID: 1, BranchID: 1, CountryID: 1},
			{InsID: 2, Name: "AnotherCompany", Ticker: "AC", MarketID: 2, SectorID: 2, BranchID: 2, CountryID: 1},
		},
	}
	log, c := newTestServer(t, map[string]any{
		"/v1/instruments": expected,
	})

	result, err := c.GetInstruments()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Instruments) != 2 {
		t.Fatalf("expected 2 instruments, got %d", len(result.Instruments))
	}
	if result.Instruments[0].Name != "TestCompany" {
		t.Errorf("expected TestCompany, got %s", result.Instruments[0].Name)
	}
	log.assertParam(t, "/v1/instruments", "authKey", "test-key")
}

func TestGetInstrumentsUpdated(t *testing.T) {
	expected := models.InstrumentsUpdatedResponse{
		Instruments: []models.InstrumentUpdatedV1{
			{InsID: 1, UpdatedDate: "2025-01-15"},
		},
	}
	log, c := newTestServer(t, map[string]any{
		"/v1/instruments/updated": expected,
	})

	result, err := c.GetInstrumentsUpdated()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Instruments) != 1 {
		t.Fatalf("expected 1 instrument, got %d", len(result.Instruments))
	}
	log.assertParam(t, "/v1/instruments/updated", "authKey", "test-key")
}

func TestGetStockSplits(t *testing.T) {
	expected := models.StockSplitsResponse{
		StockSplits: []models.StockSplitV1{
			{InsID: 1, SplitDate: "2024-06-01", SplitType: "split", Ratio: 2.0},
		},
	}
	log, c := newTestServer(t, map[string]any{
		"/v1/instruments/stocksplits": expected,
	})

	result, err := c.GetStockSplits()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.StockSplits) != 1 {
		t.Fatalf("expected 1 stock split, got %d", len(result.StockSplits))
	}
	log.assertParam(t, "/v1/instruments/stocksplits", "authKey", "test-key")
}

func TestGetGlobalInstruments(t *testing.T) {
	expected := models.InstrumentsResponse{
		Instruments: []models.InstrumentV1{
			{InsID: 100, Name: "Apple Inc", Ticker: "AAPL", MarketID: 10, SectorID: 5, BranchID: 5, CountryID: 10},
			{InsID: 101, Name: "Microsoft Corp", Ticker: "MSFT", MarketID: 10, SectorID: 5, BranchID: 5, CountryID: 10},
		},
	}
	log, c := newTestServer(t, map[string]any{
		"/v1/instruments/global": expected,
	})

	result, err := c.GetGlobalInstruments()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Instruments) != 2 {
		t.Fatalf("expected 2 instruments, got %d", len(result.Instruments))
	}
	if result.Instruments[0].Name != "Apple Inc" {
		t.Errorf("expected Apple Inc, got %s", result.Instruments[0].Name)
	}
	log.assertParam(t, "/v1/instruments/global", "authKey", "test-key")
}
