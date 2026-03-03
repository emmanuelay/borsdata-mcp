package client

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/emmanuelay/borsdata-mcp/models"
)

func TestGetYearlyReports(t *testing.T) {
	expected := models.ReportsResponse{
		Reports: []models.ReportV1{
			{Year: 2024, Period: 0, Revenues: 1000000},
		},
	}
	log, c := newTestServer(t, map[string]any{
		"/v1/instruments/1/reports/year": expected,
	})

	result, err := c.GetYearlyReports(1, 5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Reports) != 1 {
		t.Fatalf("expected 1 report, got %d", len(result.Reports))
	}
	log.assertParam(t, "/v1/instruments/1/reports/year", "authKey", "test-key")
	log.assertParam(t, "/v1/instruments/1/reports/year", "maxCount", "5")
}

func TestGetR12Reports(t *testing.T) {
	expected := models.ReportsResponse{
		Reports: []models.ReportV1{
			{Year: 2024, Period: 4, Revenues: 500000},
		},
	}
	log, c := newTestServer(t, map[string]any{
		"/v1/instruments/1/reports/r12": expected,
	})

	result, err := c.GetR12Reports(1, 5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Reports) != 1 {
		t.Fatalf("expected 1 report, got %d", len(result.Reports))
	}
	log.assertParam(t, "/v1/instruments/1/reports/r12", "authKey", "test-key")
	log.assertParam(t, "/v1/instruments/1/reports/r12", "maxCount", "5")
}

func TestGetQuarterlyReports(t *testing.T) {
	expected := models.ReportsResponse{
		Reports: []models.ReportV1{
			{Year: 2024, Period: 3, Revenues: 250000},
		},
	}
	log, c := newTestServer(t, map[string]any{
		"/v1/instruments/1/reports/quarter": expected,
	})

	result, err := c.GetQuarterlyReports(1, 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Reports) != 1 {
		t.Fatalf("expected 1 report, got %d", len(result.Reports))
	}
	log.assertParam(t, "/v1/instruments/1/reports/quarter", "authKey", "test-key")
	log.assertNoParam(t, "/v1/instruments/1/reports/quarter", "maxCount")
}

func TestGetAllReports(t *testing.T) {
	expected := models.ReportsCompoundResponse{
		ReportsYear:    []models.ReportV1{{Year: 2024, Period: 0}},
		ReportsR12:     []models.ReportV1{{Year: 2024, Period: 4}},
		ReportsQuarter: []models.ReportV1{{Year: 2024, Period: 3}},
	}
	log, c := newTestServer(t, map[string]any{
		"/v1/instruments/1/reports": expected,
	})

	result, err := c.GetAllReports(1, 5, 5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.ReportsYear) != 1 {
		t.Fatalf("expected 1 yearly report, got %d", len(result.ReportsYear))
	}
	log.assertParam(t, "/v1/instruments/1/reports", "authKey", "test-key")
	log.assertParam(t, "/v1/instruments/1/reports", "maxYearCount", "5")
	log.assertParam(t, "/v1/instruments/1/reports", "maxR12QCount", "5")
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
