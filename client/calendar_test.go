package client

import (
	"testing"

	"github.com/emmanuelay/borsdata-mcp/models"
)

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
	log, c := newTestServer(t, map[string]any{
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
	log.assertParam(t, "/v1/instruments/report/calendar", "authKey", "test-key")
	log.assertParam(t, "/v1/instruments/report/calendar", "instList", "1")
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
	log, c := newTestServer(t, map[string]any{
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
	log.assertParam(t, "/v1/instruments/dividend/calendar", "authKey", "test-key")
	log.assertParam(t, "/v1/instruments/dividend/calendar", "instList", "1")
}
