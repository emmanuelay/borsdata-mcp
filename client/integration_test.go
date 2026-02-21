//go:build integration

package client

import (
	"os"
	"testing"
)

func integrationClient(t *testing.T) *Client {
	t.Helper()
	apiKey := os.Getenv("BORSDATA_API_KEY")
	if apiKey == "" {
		t.Skip("BORSDATA_API_KEY not set")
	}
	return New(apiKey, "test-integration")
}

func TestIntegrationGetInstruments(t *testing.T) {
	c := integrationClient(t)
	result, err := c.GetInstruments()
	if err != nil {
		t.Fatalf("GetInstruments: %v", err)
	}
	if len(result.Instruments) == 0 {
		t.Fatal("expected at least one instrument")
	}
	t.Logf("Got %d instruments", len(result.Instruments))
}

func TestIntegrationGetCountries(t *testing.T) {
	c := integrationClient(t)
	result, err := c.GetCountries()
	if err != nil {
		t.Fatalf("GetCountries: %v", err)
	}
	if len(result.Countries) == 0 {
		t.Fatal("expected at least one country")
	}
	t.Logf("Got %d countries", len(result.Countries))
}

func TestIntegrationGetMarkets(t *testing.T) {
	c := integrationClient(t)
	result, err := c.GetMarkets()
	if err != nil {
		t.Fatalf("GetMarkets: %v", err)
	}
	if len(result.Markets) == 0 {
		t.Fatal("expected at least one market")
	}
	t.Logf("Got %d markets", len(result.Markets))
}

func TestIntegrationGetSectors(t *testing.T) {
	c := integrationClient(t)
	result, err := c.GetSectors()
	if err != nil {
		t.Fatalf("GetSectors: %v", err)
	}
	if len(result.Sectors) == 0 {
		t.Fatal("expected at least one sector")
	}
	t.Logf("Got %d sectors", len(result.Sectors))
}

func TestIntegrationGetBranches(t *testing.T) {
	c := integrationClient(t)
	result, err := c.GetBranches()
	if err != nil {
		t.Fatalf("GetBranches: %v", err)
	}
	if len(result.Branches) == 0 {
		t.Fatal("expected at least one branch")
	}
	t.Logf("Got %d branches", len(result.Branches))
}

func TestIntegrationGetLatestStockPrices(t *testing.T) {
	c := integrationClient(t)
	result, err := c.GetLatestStockPrices()
	if err != nil {
		t.Fatalf("GetLatestStockPrices: %v", err)
	}
	if len(result.StockPricesList) == 0 {
		t.Fatal("expected at least one stock price")
	}
	t.Logf("Got %d latest stock prices", len(result.StockPricesList))
}

func TestIntegrationGetKpiMetadata(t *testing.T) {
	c := integrationClient(t)
	result, err := c.GetKpiMetadata()
	if err != nil {
		t.Fatalf("GetKpiMetadata: %v", err)
	}
	if len(result.KpiHistoryMetadatas) == 0 {
		t.Fatal("expected at least one KPI metadata entry")
	}
	t.Logf("Got %d KPI metadata entries", len(result.KpiHistoryMetadatas))
}

// firstInstrumentID returns the ID of the first instrument from the API.
func firstInstrumentID(t *testing.T, c *Client) int64 {
	t.Helper()
	result, err := c.GetInstruments()
	if err != nil {
		t.Fatalf("GetInstruments: %v", err)
	}
	if len(result.Instruments) == 0 {
		t.Fatal("expected at least one instrument")
	}
	return result.Instruments[0].InsID
}

func TestIntegrationGetYearlyReportsMaxCount(t *testing.T) {
	c := integrationClient(t)
	insID := firstInstrumentID(t, c)

	result, err := c.GetYearlyReports(insID, 2)
	if err != nil {
		t.Fatalf("GetYearlyReports: %v", err)
	}
	if len(result.Reports) > 2 {
		t.Fatalf("expected at most 2 reports with maxCount=2, got %d", len(result.Reports))
	}
	t.Logf("Got %d yearly reports with maxCount=2", len(result.Reports))
}

func TestIntegrationGetR12ReportsMaxCount(t *testing.T) {
	c := integrationClient(t)
	insID := firstInstrumentID(t, c)

	result, err := c.GetR12Reports(insID, 3)
	if err != nil {
		t.Fatalf("GetR12Reports: %v", err)
	}
	if len(result.Reports) > 3 {
		t.Fatalf("expected at most 3 reports with maxCount=3, got %d", len(result.Reports))
	}
	t.Logf("Got %d R12 reports with maxCount=3", len(result.Reports))
}

func TestIntegrationGetAllReportsMaxCount(t *testing.T) {
	c := integrationClient(t)
	insID := firstInstrumentID(t, c)

	result, err := c.GetAllReports(insID, 2, 2)
	if err != nil {
		t.Fatalf("GetAllReports: %v", err)
	}
	if len(result.ReportsYear) > 2 {
		t.Fatalf("expected at most 2 yearly reports with maxYearCount=2, got %d", len(result.ReportsYear))
	}
	if len(result.ReportsR12) > 2 {
		t.Fatalf("expected at most 2 R12 reports with maxR12QCount=2, got %d", len(result.ReportsR12))
	}
	t.Logf("Got %d yearly, %d R12, %d quarterly reports with maxCount=2",
		len(result.ReportsYear), len(result.ReportsR12), len(result.ReportsQuarter))
}

func TestIntegrationGetKpiSummaryMaxCount(t *testing.T) {
	c := integrationClient(t)
	insID := firstInstrumentID(t, c)

	result, err := c.GetKpiSummary(insID, "year", 3)
	if err != nil {
		t.Fatalf("GetKpiSummary: %v", err)
	}
	if len(result.Kpis) == 0 {
		t.Fatal("expected at least 1 KPI group")
	}
	// maxCount limits the number of values per KPI group, not the number of groups
	for _, kpi := range result.Kpis {
		if len(kpi.Values) > 3 {
			t.Fatalf("expected at most 3 values per KPI with maxCount=3, got %d for KpiID=%d", len(kpi.Values), kpi.KpiID)
		}
	}
	t.Logf("Got %d KPI groups with maxCount=3, each with <= 3 values", len(result.Kpis))
}

func TestIntegrationGetKpiScreener(t *testing.T) {
	c := integrationClient(t)
	insID := firstInstrumentID(t, c)

	result, err := c.GetKpiScreener(insID, 2, "1year", "mean")
	if err != nil {
		t.Fatalf("GetKpiScreener: %v", err)
	}
	if result.KpiID != 2 {
		t.Fatalf("expected kpiId=2, got %d", result.KpiID)
	}
	if result.Group != "1year" {
		t.Fatalf("expected group=1year, got %s", result.Group)
	}
	if result.Value.I != insID {
		t.Fatalf("expected value.i=%d, got %d", insID, result.Value.I)
	}
	t.Logf("KPI screener: insID=%d, kpiID=%d, value=%.4f", result.Value.I, result.KpiID, result.Value.N)
}

func TestIntegrationGetKpiListScreener(t *testing.T) {
	c := integrationClient(t)

	result, err := c.GetKpiListScreener(2, "1year", "mean")
	if err != nil {
		t.Fatalf("GetKpiListScreener: %v", err)
	}
	if len(result.Values) == 0 {
		t.Fatal("expected at least one value")
	}
	// Verify the values have non-zero instrument IDs
	for i, v := range result.Values {
		if v.I == 0 {
			t.Fatalf("value[%d] has zero instrument ID", i)
		}
	}
	t.Logf("Got %d KPI list screener values", len(result.Values))
}

func TestIntegrationGetStockPricesMaxCount(t *testing.T) {
	c := integrationClient(t)
	insID := firstInstrumentID(t, c)

	result, err := c.GetStockPrices(insID, 5)
	if err != nil {
		t.Fatalf("GetStockPrices: %v", err)
	}
	if len(result.StockPricesList) > 5 {
		t.Fatalf("expected at most 5 prices with maxcount=5, got %d", len(result.StockPricesList))
	}
	t.Logf("Got %d stock prices with maxcount=5", len(result.StockPricesList))
}
