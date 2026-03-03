package client

import (
	"testing"

	"github.com/emmanuelay/borsdata-mcp/models"
)

func TestGetStockPrices(t *testing.T) {
	expected := models.StockPricesResponse{
		StockPricesList: []models.StockPriceV1{
			{D: "2025-01-15", C: 100.5, H: 102.0, L: 99.0, O: 100.0, V: 50000},
			{D: "2025-01-14", C: 99.5, H: 101.0, L: 98.0, O: 99.0, V: 45000},
		},
	}
	log, c := newTestServer(t, map[string]any{
		"/v1/instruments/1/stockprices": expected,
	})

	result, err := c.GetStockPrices(1, 10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.StockPricesList) != 2 {
		t.Fatalf("expected 2 stock prices, got %d", len(result.StockPricesList))
	}
	log.assertParam(t, "/v1/instruments/1/stockprices", "authKey", "test-key")
}

func TestGetStockPricesDateRange(t *testing.T) {
	expected := models.StockPricesResponse{
		StockPricesList: []models.StockPriceV1{
			{D: "2025-01-15", C: 100.5, H: 102.0, L: 99.0, O: 100.0, V: 50000},
		},
	}
	log, c := newTestServer(t, map[string]any{
		"/v1/instruments/1/stockprices": expected,
	})

	result, err := c.GetStockPricesDateRange(1, "2025-01-01", "2025-01-31")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.StockPricesList) != 1 {
		t.Fatalf("expected 1 stock price, got %d", len(result.StockPricesList))
	}
	log.assertParam(t, "/v1/instruments/1/stockprices", "authKey", "test-key")
	log.assertParam(t, "/v1/instruments/1/stockprices", "from", "2025-01-01")
	log.assertParam(t, "/v1/instruments/1/stockprices", "to", "2025-01-31")
}

func TestGetLatestStockPrices(t *testing.T) {
	expected := models.StockPricesLastResponse{
		StockPricesList: []models.StockPriceLastV1{
			{I: 1, D: "2025-01-15", C: 100.5, H: 102.0, L: 99.0, O: 100.0, V: 50000},
		},
	}
	log, c := newTestServer(t, map[string]any{
		"/v1/instruments/stockprices/last": expected,
	})

	result, err := c.GetLatestStockPrices()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.StockPricesList) != 1 {
		t.Fatalf("expected 1 stock price, got %d", len(result.StockPricesList))
	}
	log.assertParam(t, "/v1/instruments/stockprices/last", "authKey", "test-key")
}

func TestGetStockPricesByDate(t *testing.T) {
	expected := models.StockPricesDateResponse{
		StockPricesList: []models.StockPriceDateV1{
			{I: 1, D: "2025-01-15", C: 100.5, H: 102.0, L: 99.0, O: 100.0, V: 50000},
		},
	}
	log, c := newTestServer(t, map[string]any{
		"/v1/instruments/stockprices/date": expected,
	})

	result, err := c.GetStockPricesByDate("2025-01-15")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.StockPricesList) != 1 {
		t.Fatalf("expected 1 stock price, got %d", len(result.StockPricesList))
	}
	log.assertParam(t, "/v1/instruments/stockprices/date", "authKey", "test-key")
	log.assertParam(t, "/v1/instruments/stockprices/date", "date", "2025-01-15")
}

func TestGetGlobalLatestStockPrices(t *testing.T) {
	expected := models.StockPricesLastResponse{
		StockPricesList: []models.StockPriceLastV1{
			{I: 100, D: "2025-01-15", C: 230.5, H: 232.0, L: 229.0, O: 230.0, V: 1000000},
		},
	}
	log, c := newTestServer(t, map[string]any{
		"/v1/instruments/stockprices/global/last": expected,
	})

	result, err := c.GetGlobalLatestStockPrices()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.StockPricesList) != 1 {
		t.Fatalf("expected 1 stock price, got %d", len(result.StockPricesList))
	}
	log.assertParam(t, "/v1/instruments/stockprices/global/last", "authKey", "test-key")
}

func TestGetGlobalStockPricesByDate(t *testing.T) {
	expected := models.StockPricesDateResponse{
		StockPricesList: []models.StockPriceDateV1{
			{I: 100, D: "2025-01-15", C: 230.5, H: 232.0, L: 229.0, O: 230.0, V: 1000000},
		},
	}
	log, c := newTestServer(t, map[string]any{
		"/v1/instruments/stockprices/global/date": expected,
	})

	result, err := c.GetGlobalStockPricesByDate("2025-01-15")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.StockPricesList) != 1 {
		t.Fatalf("expected 1 stock price, got %d", len(result.StockPricesList))
	}
	log.assertParam(t, "/v1/instruments/stockprices/global/date", "authKey", "test-key")
	log.assertParam(t, "/v1/instruments/stockprices/global/date", "date", "2025-01-15")
}

func TestGetStockPricesMaxCount(t *testing.T) {
	allPrices := []models.StockPriceV1{
		{D: "2025-01-15", C: 105.0, H: 106.0, L: 104.0, O: 104.5, V: 50000},
		{D: "2025-01-14", C: 104.0, H: 105.0, L: 103.0, O: 103.5, V: 45000},
		{D: "2025-01-13", C: 103.0, H: 104.0, L: 102.0, O: 102.5, V: 40000},
		{D: "2025-01-10", C: 102.0, H: 103.0, L: 101.0, O: 101.5, V: 35000},
		{D: "2025-01-09", C: 101.0, H: 102.0, L: 100.0, O: 100.5, V: 30000},
	}

	_, c := newTestServer(t, map[string]any{
		"/v1/instruments/1/stockprices": models.StockPricesResponse{StockPricesList: allPrices},
	})

	result, err := c.GetStockPrices(1, 3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.StockPricesList) != 3 {
		t.Fatalf("expected 3 prices with maxCount=3, got %d", len(result.StockPricesList))
	}
	if result.StockPricesList[0].D != "2025-01-13" {
		t.Errorf("expected most recent 3 prices starting at 2025-01-13, got %s", result.StockPricesList[0].D)
	}

	result, err = c.GetStockPrices(1, 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.StockPricesList) != 5 {
		t.Fatalf("expected all 5 prices with maxCount=0, got %d", len(result.StockPricesList))
	}
}
