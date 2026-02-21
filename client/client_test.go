package client

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/emmanuelay/borsdata-mcp/models"
)

// newTestServer creates a mock HTTP server that returns the given response for matching paths.
func newTestServer(t *testing.T, handlers map[string]any) (*httptest.Server, *Client) {
	t.Helper()
	mux := http.NewServeMux()
	for path, resp := range handlers {
		body, err := json.Marshal(resp)
		if err != nil {
			t.Fatalf("marshaling mock response: %v", err)
		}
		mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write(body)
		})
	}
	server := httptest.NewServer(mux)
	client := NewWithBaseURL("test-key", server.URL, "test")
	t.Cleanup(server.Close)
	return server, client
}

func TestGetInstruments(t *testing.T) {
	expected := models.InstrumentsResponse{
		Instruments: []models.InstrumentV1{
			{InsID: 1, Name: "TestCompany", Ticker: "TEST", MarketID: 1, SectorID: 1, BranchID: 1, CountryID: 1},
			{InsID: 2, Name: "AnotherCompany", Ticker: "AC", MarketID: 2, SectorID: 2, BranchID: 2, CountryID: 1},
		},
	}
	_, c := newTestServer(t, map[string]any{
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
}

func TestGetInstrumentsUpdated(t *testing.T) {
	expected := models.InstrumentsUpdatedResponse{
		Instruments: []models.InstrumentUpdatedV1{
			{InsID: 1, UpdatedDate: "2025-01-15"},
		},
	}
	_, c := newTestServer(t, map[string]any{
		"/v1/instruments/updated": expected,
	})

	result, err := c.GetInstrumentsUpdated()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Instruments) != 1 {
		t.Fatalf("expected 1 instrument, got %d", len(result.Instruments))
	}
}

func TestGetStockSplits(t *testing.T) {
	expected := models.StockSplitsResponse{
		StockSplits: []models.StockSplitV1{
			{InsID: 1, SplitDate: "2024-06-01", SplitType: "split", Ratio: 2.0},
		},
	}
	_, c := newTestServer(t, map[string]any{
		"/v1/instruments/stocksplits": expected,
	})

	result, err := c.GetStockSplits()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.StockSplits) != 1 {
		t.Fatalf("expected 1 stock split, got %d", len(result.StockSplits))
	}
}

func TestGetStockPrices(t *testing.T) {
	expected := models.StockPricesResponse{
		StockPricesList: []models.StockPriceV1{
			{D: "2025-01-15", C: 100.5, H: 102.0, L: 99.0, O: 100.0, V: 50000},
			{D: "2025-01-14", C: 99.5, H: 101.0, L: 98.0, O: 99.0, V: 45000},
		},
	}
	_, c := newTestServer(t, map[string]any{
		"/v1/instruments/1/stockprices": expected,
	})

	result, err := c.GetStockPrices(1, 10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.StockPricesList) != 2 {
		t.Fatalf("expected 2 stock prices, got %d", len(result.StockPricesList))
	}
}

func TestGetStockPricesDateRange(t *testing.T) {
	expected := models.StockPricesResponse{
		StockPricesList: []models.StockPriceV1{
			{D: "2025-01-15", C: 100.5, H: 102.0, L: 99.0, O: 100.0, V: 50000},
		},
	}
	_, c := newTestServer(t, map[string]any{
		"/v1/instruments/1/stockprices": expected,
	})

	result, err := c.GetStockPricesDateRange(1, "2025-01-01", "2025-01-31")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.StockPricesList) != 1 {
		t.Fatalf("expected 1 stock price, got %d", len(result.StockPricesList))
	}
}

func TestGetLatestStockPrices(t *testing.T) {
	expected := models.StockPricesLastResponse{
		StockPricesList: []models.StockPriceLastV1{
			{I: 1, D: "2025-01-15", C: 100.5, H: 102.0, L: 99.0, O: 100.0, V: 50000},
		},
	}
	_, c := newTestServer(t, map[string]any{
		"/v1/instruments/stockprices/last": expected,
	})

	result, err := c.GetLatestStockPrices()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.StockPricesList) != 1 {
		t.Fatalf("expected 1 stock price, got %d", len(result.StockPricesList))
	}
}

func TestGetStockPricesByDate(t *testing.T) {
	expected := models.StockPricesDateResponse{
		StockPricesList: []models.StockPriceDateV1{
			{I: 1, D: "2025-01-15", C: 100.5, H: 102.0, L: 99.0, O: 100.0, V: 50000},
		},
	}
	_, c := newTestServer(t, map[string]any{
		"/v1/instruments/stockprices/date": expected,
	})

	result, err := c.GetStockPricesByDate("2025-01-15")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.StockPricesList) != 1 {
		t.Fatalf("expected 1 stock price, got %d", len(result.StockPricesList))
	}
}

func TestGetYearlyReports(t *testing.T) {
	expected := models.ReportsResponse{
		Reports: []models.ReportV1{
			{Year: 2024, Period: 0, Revenues: 1000000},
		},
	}
	_, c := newTestServer(t, map[string]any{
		"/v1/instruments/1/reports/year": expected,
	})

	result, err := c.GetYearlyReports(1, 5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Reports) != 1 {
		t.Fatalf("expected 1 report, got %d", len(result.Reports))
	}
}

func TestGetR12Reports(t *testing.T) {
	expected := models.ReportsResponse{
		Reports: []models.ReportV1{
			{Year: 2024, Period: 4, Revenues: 500000},
		},
	}
	_, c := newTestServer(t, map[string]any{
		"/v1/instruments/1/reports/r12": expected,
	})

	result, err := c.GetR12Reports(1, 5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Reports) != 1 {
		t.Fatalf("expected 1 report, got %d", len(result.Reports))
	}
}

func TestGetQuarterlyReports(t *testing.T) {
	expected := models.ReportsResponse{
		Reports: []models.ReportV1{
			{Year: 2024, Period: 3, Revenues: 250000},
		},
	}
	_, c := newTestServer(t, map[string]any{
		"/v1/instruments/1/reports/quarter": expected,
	})

	result, err := c.GetQuarterlyReports(1, 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Reports) != 1 {
		t.Fatalf("expected 1 report, got %d", len(result.Reports))
	}
}

func TestGetAllReports(t *testing.T) {
	expected := models.ReportsCompoundResponse{
		ReportsYear:    []models.ReportV1{{Year: 2024, Period: 0}},
		ReportsR12:     []models.ReportV1{{Year: 2024, Period: 4}},
		ReportsQuarter: []models.ReportV1{{Year: 2024, Period: 3}},
	}
	_, c := newTestServer(t, map[string]any{
		"/v1/instruments/1/reports": expected,
	})

	result, err := c.GetAllReports(1, 5, 5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.ReportsYear) != 1 {
		t.Fatalf("expected 1 yearly report, got %d", len(result.ReportsYear))
	}
}

func TestGetKpiHistory(t *testing.T) {
	expected := models.KpisHistoryResponse{
		Values: []models.KpiHistoryV1{
			{Y: 2024, P: 1, V: 15.5},
			{Y: 2024, P: 2, V: 16.0},
		},
	}
	_, c := newTestServer(t, map[string]any{
		"/v1/instruments/1/kpis/2/quarter/mean/history": expected,
	})

	result, err := c.GetKpiHistory(1, 2, "quarter", "mean", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Values) != 2 {
		t.Fatalf("expected 2 KPI history values, got %d", len(result.Values))
	}
}

func TestGetKpiSummary(t *testing.T) {
	expected := models.KpisSummaryResponse{
		Instrument: 1,
		ReportType: "year",
		Kpis: []models.KpisSummaryGroupV1{
			{KpiID: 1, Values: []models.KpisSummaryValueV1{
				{Y: 2025, P: 4, V: 2.22},
			}},
		},
	}
	_, c := newTestServer(t, map[string]any{
		"/v1/instruments/1/kpis/year/summary": expected,
	})

	result, err := c.GetKpiSummary(1, "year", 5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Kpis) != 1 {
		t.Fatalf("expected 1 KPI summary, got %d", len(result.Kpis))
	}
}

func TestGetCountries(t *testing.T) {
	expected := models.CountriesResponse{
		Countries: []models.CountryV1{
			{ID: 1, Name: "Sweden"},
		},
	}
	_, c := newTestServer(t, map[string]any{
		"/v1/countries": expected,
	})

	result, err := c.GetCountries()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Countries) != 1 {
		t.Fatalf("expected 1 country, got %d", len(result.Countries))
	}
	if result.Countries[0].Name != "Sweden" {
		t.Errorf("expected Sweden, got %s", result.Countries[0].Name)
	}
}

func TestGetMarkets(t *testing.T) {
	expected := models.MarketsResponse{
		Markets: []models.MarketV1{
			{ID: 1, Name: "Large Cap", CountryID: 1},
		},
	}
	_, c := newTestServer(t, map[string]any{
		"/v1/markets": expected,
	})

	result, err := c.GetMarkets()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Markets) != 1 {
		t.Fatalf("expected 1 market, got %d", len(result.Markets))
	}
}

func TestGetSectors(t *testing.T) {
	expected := models.SectorsResponse{
		Sectors: []models.SectorV1{
			{ID: 1, Name: "Technology"},
		},
	}
	_, c := newTestServer(t, map[string]any{
		"/v1/sectors": expected,
	})

	result, err := c.GetSectors()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Sectors) != 1 {
		t.Fatalf("expected 1 sector, got %d", len(result.Sectors))
	}
}

func TestGetBranches(t *testing.T) {
	expected := models.BranchesResponse{
		Branches: []models.BranchV1{
			{ID: 1, Name: "Software", SectorID: 1},
		},
	}
	_, c := newTestServer(t, map[string]any{
		"/v1/branches": expected,
	})

	result, err := c.GetBranches()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Branches) != 1 {
		t.Fatalf("expected 1 branch, got %d", len(result.Branches))
	}
}

func TestGetKpiMetadata(t *testing.T) {
	expected := models.KpisMetadataResponse{
		KpiHistoryMetadatas: []models.KpiMetadataV1{
			{KpiID: 1, NameSv: "P/E-tal", NameEn: "P/E ratio"},
		},
	}
	_, c := newTestServer(t, map[string]any{
		"/v1/instruments/kpis/metadata": expected,
	})

	result, err := c.GetKpiMetadata()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.KpiHistoryMetadatas) != 1 {
		t.Fatalf("expected 1 KPI metadata, got %d", len(result.KpiHistoryMetadatas))
	}
}

func TestAPIErrorHandling(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error": "unauthorized"}`))
	}))
	defer server.Close()

	c := NewWithBaseURL("bad-key", server.URL, "test")
	_, err := c.GetInstruments()
	if err == nil {
		t.Fatal("expected error for unauthorized request")
	}
}

// reportsMaxCountHandler returns a handler that respects the maxCount query parameter.
// It returns up to maxCount reports from the full set, simulating the real API behavior.
func reportsMaxCountHandler(t *testing.T, allReports []models.ReportV1) http.HandlerFunc {
	t.Helper()
	return func(w http.ResponseWriter, r *http.Request) {
		reports := allReports
		if mc := r.URL.Query().Get("maxCount"); mc != "" {
			n, err := strconv.Atoi(mc)
			if err == nil && n < len(reports) {
				reports = reports[:n]
			}
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(models.ReportsResponse{Reports: reports})
	}
}

func TestGetYearlyReportsMaxCount(t *testing.T) {
	allReports := []models.ReportV1{
		{Year: 2024, Period: 0, Revenues: 400},
		{Year: 2023, Period: 0, Revenues: 300},
		{Year: 2022, Period: 0, Revenues: 200},
		{Year: 2021, Period: 0, Revenues: 100},
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/instruments/1/reports/year", reportsMaxCountHandler(t, allReports))
	server := httptest.NewServer(mux)
	defer server.Close()
	c := NewWithBaseURL("test-key", server.URL, "test")

	result, err := c.GetYearlyReports(1, 2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Reports) != 2 {
		t.Fatalf("expected 2 reports with maxCount=2, got %d", len(result.Reports))
	}

	result, err = c.GetYearlyReports(1, 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Reports) != 4 {
		t.Fatalf("expected all 4 reports with maxCount=0, got %d", len(result.Reports))
	}
}

func TestGetR12ReportsMaxCount(t *testing.T) {
	allReports := []models.ReportV1{
		{Year: 2024, Period: 4, Revenues: 400},
		{Year: 2023, Period: 4, Revenues: 300},
		{Year: 2022, Period: 4, Revenues: 200},
		{Year: 2021, Period: 4, Revenues: 100},
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/instruments/1/reports/r12", reportsMaxCountHandler(t, allReports))
	server := httptest.NewServer(mux)
	defer server.Close()
	c := NewWithBaseURL("test-key", server.URL, "test")

	result, err := c.GetR12Reports(1, 3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Reports) != 3 {
		t.Fatalf("expected 3 reports with maxCount=3, got %d", len(result.Reports))
	}
}

func TestGetAllReportsMaxCount(t *testing.T) {
	allYear := []models.ReportV1{
		{Year: 2024, Period: 0}, {Year: 2023, Period: 0}, {Year: 2022, Period: 0},
	}
	allR12 := []models.ReportV1{
		{Year: 2024, Period: 4}, {Year: 2023, Period: 4}, {Year: 2022, Period: 4},
	}
	allQuarter := []models.ReportV1{
		{Year: 2024, Period: 3}, {Year: 2024, Period: 2}, {Year: 2024, Period: 1},
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/instruments/1/reports", func(w http.ResponseWriter, r *http.Request) {
		year := allYear
		r12 := allR12
		if mc := r.URL.Query().Get("maxYearCount"); mc != "" {
			n, err := strconv.Atoi(mc)
			if err == nil && n < len(year) {
				year = year[:n]
			}
		}
		if mc := r.URL.Query().Get("maxR12QCount"); mc != "" {
			n, err := strconv.Atoi(mc)
			if err == nil && n < len(r12) {
				r12 = r12[:n]
			}
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(models.ReportsCompoundResponse{
			ReportsYear:    year,
			ReportsR12:     r12,
			ReportsQuarter: allQuarter,
		})
	})
	server := httptest.NewServer(mux)
	defer server.Close()
	c := NewWithBaseURL("test-key", server.URL, "test")

	result, err := c.GetAllReports(1, 1, 2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.ReportsYear) != 1 {
		t.Fatalf("expected 1 yearly report with maxYearCount=1, got %d", len(result.ReportsYear))
	}
	if len(result.ReportsR12) != 2 {
		t.Fatalf("expected 2 R12 reports with maxR12QCount=2, got %d", len(result.ReportsR12))
	}
	if len(result.ReportsQuarter) != 3 {
		t.Fatalf("expected all 3 quarterly reports (no limit param), got %d", len(result.ReportsQuarter))
	}
}

func TestGetKpiSummaryMaxCount(t *testing.T) {
	allKpis := []models.KpisSummaryGroupV1{
		{KpiID: 1, Values: []models.KpisSummaryValueV1{{Y: 2025, P: 4, V: 10.0}}},
		{KpiID: 2, Values: []models.KpisSummaryValueV1{{Y: 2025, P: 4, V: 20.0}}},
		{KpiID: 3, Values: []models.KpisSummaryValueV1{{Y: 2025, P: 4, V: 30.0}}},
		{KpiID: 4, Values: []models.KpisSummaryValueV1{{Y: 2025, P: 4, V: 40.0}}},
		{KpiID: 5, Values: []models.KpisSummaryValueV1{{Y: 2025, P: 4, V: 50.0}}},
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/instruments/1/kpis/year/summary", func(w http.ResponseWriter, r *http.Request) {
		kpis := allKpis
		if mc := r.URL.Query().Get("maxCount"); mc != "" {
			n, err := strconv.Atoi(mc)
			if err == nil && n < len(kpis) {
				kpis = kpis[:n]
			}
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(models.KpisSummaryResponse{Instrument: 1, ReportType: "year", Kpis: kpis})
	})
	server := httptest.NewServer(mux)
	defer server.Close()
	c := NewWithBaseURL("test-key", server.URL, "test")

	result, err := c.GetKpiSummary(1, "year", 3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Kpis) != 3 {
		t.Fatalf("expected 3 kpis with maxCount=3, got %d", len(result.Kpis))
	}

	result, err = c.GetKpiSummary(1, "year", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Kpis) != 5 {
		t.Fatalf("expected all 5 kpis with maxCount=0, got %d", len(result.Kpis))
	}
}

func TestGetGlobalInstruments(t *testing.T) {
	expected := models.InstrumentsResponse{
		Instruments: []models.InstrumentV1{
			{InsID: 100, Name: "Apple Inc", Ticker: "AAPL", MarketID: 10, SectorID: 5, BranchID: 5, CountryID: 10},
			{InsID: 101, Name: "Microsoft Corp", Ticker: "MSFT", MarketID: 10, SectorID: 5, BranchID: 5, CountryID: 10},
		},
	}
	_, c := newTestServer(t, map[string]any{
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
}

func TestGetGlobalLatestStockPrices(t *testing.T) {
	expected := models.StockPricesLastResponse{
		StockPricesList: []models.StockPriceLastV1{
			{I: 100, D: "2025-01-15", C: 230.5, H: 232.0, L: 229.0, O: 230.0, V: 1000000},
		},
	}
	_, c := newTestServer(t, map[string]any{
		"/v1/instruments/stockprices/global/last": expected,
	})

	result, err := c.GetGlobalLatestStockPrices()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.StockPricesList) != 1 {
		t.Fatalf("expected 1 stock price, got %d", len(result.StockPricesList))
	}
}

func TestGetGlobalStockPricesByDate(t *testing.T) {
	expected := models.StockPricesDateResponse{
		StockPricesList: []models.StockPriceDateV1{
			{I: 100, D: "2025-01-15", C: 230.5, H: 232.0, L: 229.0, O: 230.0, V: 1000000},
		},
	}
	_, c := newTestServer(t, map[string]any{
		"/v1/instruments/stockprices/global/date": expected,
	})

	result, err := c.GetGlobalStockPricesByDate("2025-01-15")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.StockPricesList) != 1 {
		t.Fatalf("expected 1 stock price, got %d", len(result.StockPricesList))
	}
}

func TestGetGlobalKpiListScreener(t *testing.T) {
	expected := models.KpisAllCompResponse{
		Values: []models.KpiAllCompV1{
			{I: 100, N: 25.5},
			{I: 101, N: 30.0},
		},
	}
	_, c := newTestServer(t, map[string]any{
		"/v1/instruments/global/kpis/2/last/mean": expected,
	})

	result, err := c.GetGlobalKpiListScreener(2, "last", "mean")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Values) != 2 {
		t.Fatalf("expected 2 KPI values, got %d", len(result.Values))
	}
}

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
	_, c := newTestServer(t, map[string]any{
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
	_, c := newTestServer(t, map[string]any{
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
	_, c := newTestServer(t, map[string]any{
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
}

func TestGetReportCalendar(t *testing.T) {
	expected := models.ReportCalendarResponse{
		List: []models.ReportCalendarRespV1{
			{
				InsID: 1,
				Values: []models.ReportCalendarDateV1{
					{ReleaseDate: "2025-02-15", ReportType: "quarter"},
				},
			},
		},
	}
	_, c := newTestServer(t, map[string]any{
		"/v1/instruments/report/calendar": expected,
	})

	result, err := c.GetReportCalendar("1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.List) != 1 {
		t.Fatalf("expected 1 calendar resp, got %d", len(result.List))
	}
	if result.List[0].Values[0].ReportType != "quarter" {
		t.Errorf("expected quarter, got %s", result.List[0].Values[0].ReportType)
	}
}

func TestGetDividendCalendar(t *testing.T) {
	expected := models.DividendCalendarResponse{
		List: []models.DividendCalendarRespV1{
			{
				InsID: 1,
				Values: []models.DividendDateV1{
					{AmountPaid: 5.0, CurrencyShortName: "SEK", DistributionFrequency: 1, ExcludingDate: "2025-04-01", DividendType: 0},
				},
			},
		},
	}
	_, c := newTestServer(t, map[string]any{
		"/v1/instruments/dividend/calendar": expected,
	})

	result, err := c.GetDividendCalendar("1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.List) != 1 {
		t.Fatalf("expected 1 dividend calendar resp, got %d", len(result.List))
	}
	if result.List[0].Values[0].AmountPaid != 5.0 {
		t.Errorf("expected amount paid 5.0, got %f", result.List[0].Values[0].AmountPaid)
	}
}

func TestGetStockPricesMaxCount(t *testing.T) {
	// The API does not support server-side limiting for stock prices,
	// so the client truncates the response. The mock always returns all prices.
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
