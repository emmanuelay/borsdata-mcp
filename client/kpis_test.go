package client

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/emmanuelay/borsdata-mcp/models"
)

func TestGetKpiHistory(t *testing.T) {
	expected := models.KpisHistoryResponse{
		Values: []models.KpiHistoryV1{
			{Y: 2024, P: 1, V: 15.5},
			{Y: 2024, P: 2, V: 16.0},
		},
	}
	log, c := newTestServer(t, map[string]any{
		"/v1/instruments/1/kpis/2/quarter/mean/history": expected,
	})

	result, err := c.GetKpiHistory(1, 2, "quarter", "mean", 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Values) != 2 {
		t.Fatalf("expected 2 KPI history values, got %d", len(result.Values))
	}
	log.assertParam(t, "/v1/instruments/1/kpis/2/quarter/mean/history", "authKey", "test-key")
	log.assertNoParam(t, "/v1/instruments/1/kpis/2/quarter/mean/history", "maxCount")
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
	log, c := newTestServer(t, map[string]any{
		"/v1/instruments/1/kpis/year/summary": expected,
	})

	result, err := c.GetKpiSummary(1, "year", 5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Kpis) != 1 {
		t.Fatalf("expected 1 KPI summary, got %d", len(result.Kpis))
	}
	log.assertParam(t, "/v1/instruments/1/kpis/year/summary", "authKey", "test-key")
	log.assertParam(t, "/v1/instruments/1/kpis/year/summary", "maxCount", "5")
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

func TestGetGlobalKpiListScreener(t *testing.T) {
	expected := models.KpisAllCompResponse{
		Values: []models.KpiAllCompV1{
			{I: 100, N: 25.5},
			{I: 101, N: 30.0},
		},
	}
	log, c := newTestServer(t, map[string]any{
		"/v1/instruments/global/kpis/2/last/mean": expected,
	})

	result, err := c.GetGlobalKpiListScreener(2, "last", "mean")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Values) != 2 {
		t.Fatalf("expected 2 KPI values, got %d", len(result.Values))
	}
	log.assertParam(t, "/v1/instruments/global/kpis/2/last/mean", "authKey", "test-key")
}
