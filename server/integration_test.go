//go:build integration

package server

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/emmanuelay/borsdata-mcp/client"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func integrationHandlers(t *testing.T) *handlers {
	t.Helper()
	apiKey := os.Getenv("BORSDATA_API_KEY")
	if apiKey == "" {
		t.Skip("BORSDATA_API_KEY not set")
	}
	c := client.New(apiKey, "test-integration")
	return &handlers{client: c}
}

func parsePaginatedResults(t *testing.T, result *mcp.CallToolResult) []paginatedInstrumentResult {
	t.Helper()
	if result.IsError {
		for _, c := range result.Content {
			if tc, ok := c.(*mcp.TextContent); ok {
				t.Fatalf("tool error: %s", tc.Text)
			}
		}
	}
	var text string
	for _, c := range result.Content {
		if tc, ok := c.(*mcp.TextContent); ok {
			text = tc.Text
			break
		}
	}
	var results []paginatedInstrumentResult
	if err := json.Unmarshal([]byte(text), &results); err != nil {
		t.Fatalf("unmarshal paginated results: %v (text: %.200s)", err, text)
	}
	return results
}

func countValues(t *testing.T, v any) int {
	t.Helper()
	if v == nil {
		return 0
	}
	data, _ := json.Marshal(v)
	var arr []json.RawMessage
	if err := json.Unmarshal(data, &arr); err != nil {
		return 0
	}
	return len(arr)
}

// --- Insider Holdings ---

func TestIntegrationInsiderHoldingsDefaultPagination(t *testing.T) {
	h := integrationHandlers(t)
	ctx := context.Background()

	result, _, err := h.getInsiderHoldings(ctx, nil, insiderHoldingsInput{
		InstrumentIDs: "3",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	results := parsePaginatedResults(t, result)
	if len(results) == 0 {
		t.Fatal("expected at least one instrument result")
	}
	r := results[0]
	t.Logf("InsID=%d, Total=%d, Offset=%d, Limit=%d, Returned=%d",
		r.InsID, r.Total, r.Offset, r.Limit, countValues(t, r.Values))

	if r.Limit != 50 {
		t.Errorf("expected default limit=50, got %d", r.Limit)
	}
	if r.Offset != 0 {
		t.Errorf("expected default offset=0, got %d", r.Offset)
	}
	returned := countValues(t, r.Values)
	if r.Total > 50 && returned != 50 {
		t.Errorf("expected 50 returned when total=%d, got %d", r.Total, returned)
	}
}

func TestIntegrationInsiderHoldingsWithLimit(t *testing.T) {
	h := integrationHandlers(t)
	ctx := context.Background()

	result, _, err := h.getInsiderHoldings(ctx, nil, insiderHoldingsInput{
		InstrumentIDs: "3",
		Limit:         5,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	results := parsePaginatedResults(t, result)
	if len(results) == 0 {
		t.Fatal("expected at least one instrument result")
	}
	r := results[0]
	returned := countValues(t, r.Values)
	t.Logf("Total=%d, Returned=%d (limit=5)", r.Total, returned)

	if returned > 5 {
		t.Errorf("expected at most 5 returned, got %d", returned)
	}
	if r.Limit != 5 {
		t.Errorf("expected limit=5, got %d", r.Limit)
	}
}

func TestIntegrationInsiderHoldingsOffset(t *testing.T) {
	h := integrationHandlers(t)
	ctx := context.Background()

	result1, _, _ := h.getInsiderHoldings(ctx, nil, insiderHoldingsInput{
		InstrumentIDs: "3",
		Limit:         2,
	})
	result2, _, _ := h.getInsiderHoldings(ctx, nil, insiderHoldingsInput{
		InstrumentIDs: "3",
		Limit:         2,
		Offset:        2,
	})

	page1 := parsePaginatedResults(t, result1)
	page2 := parsePaginatedResults(t, result2)

	if len(page1) == 0 || len(page2) == 0 {
		t.Fatal("expected results on both pages")
	}

	t.Logf("Page1: Total=%d, Offset=%d", page1[0].Total, page1[0].Offset)
	t.Logf("Page2: Total=%d, Offset=%d", page2[0].Total, page2[0].Offset)

	if page1[0].Total != page2[0].Total {
		t.Errorf("total should be consistent across pages: %d vs %d", page1[0].Total, page2[0].Total)
	}
	if page2[0].Offset != 2 {
		t.Errorf("expected offset=2 on page2, got %d", page2[0].Offset)
	}
}

func TestIntegrationInsiderHoldingsFilterByTransactionType(t *testing.T) {
	h := integrationHandlers(t)
	ctx := context.Background()

	buyResult, _, _ := h.getInsiderHoldings(ctx, nil, insiderHoldingsInput{
		InstrumentIDs:   "3",
		TransactionType: "buy",
		Limit:           100,
	})
	allResult, _, _ := h.getInsiderHoldings(ctx, nil, insiderHoldingsInput{
		InstrumentIDs: "3",
		Limit:         100,
	})

	buyResults := parsePaginatedResults(t, buyResult)
	allResults := parsePaginatedResults(t, allResult)

	if len(buyResults) == 0 || len(allResults) == 0 {
		t.Fatal("expected results")
	}

	t.Logf("Buy total=%d, All total=%d", buyResults[0].Total, allResults[0].Total)

	if buyResults[0].Total > allResults[0].Total {
		t.Errorf("buy-only total (%d) should be <= all total (%d)", buyResults[0].Total, allResults[0].Total)
	}
}

func TestIntegrationInsiderHoldingsFilterByDateRange(t *testing.T) {
	h := integrationHandlers(t)
	ctx := context.Background()

	result, _, err := h.getInsiderHoldings(ctx, nil, insiderHoldingsInput{
		InstrumentIDs: "3",
		From:          "2024-01-01",
		To:            "2024-12-31",
		Limit:         100,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	results := parsePaginatedResults(t, result)
	if len(results) == 0 {
		t.Fatal("expected results")
	}

	t.Logf("Date-filtered total=%d (2024 only)", results[0].Total)
}

func TestIntegrationInsiderHoldingsFilterByMinAmount(t *testing.T) {
	h := integrationHandlers(t)
	ctx := context.Background()

	filteredResult, _, _ := h.getInsiderHoldings(ctx, nil, insiderHoldingsInput{
		InstrumentIDs: "3",
		MinAmount:     1000000,
		Limit:         100,
	})
	allResult, _, _ := h.getInsiderHoldings(ctx, nil, insiderHoldingsInput{
		InstrumentIDs: "3",
		Limit:         100,
	})

	filtered := parsePaginatedResults(t, filteredResult)
	all := parsePaginatedResults(t, allResult)

	if len(filtered) == 0 || len(all) == 0 {
		t.Fatal("expected results")
	}

	t.Logf("MinAmount>=1M total=%d, All total=%d", filtered[0].Total, all[0].Total)

	if filtered[0].Total > all[0].Total {
		t.Errorf("filtered total (%d) should be <= all total (%d)", filtered[0].Total, all[0].Total)
	}
}

func TestIntegrationInsiderHoldingsCombinedFilters(t *testing.T) {
	h := integrationHandlers(t)
	ctx := context.Background()

	result, _, err := h.getInsiderHoldings(ctx, nil, insiderHoldingsInput{
		InstrumentIDs:   "3",
		TransactionType: "buy",
		MinAmount:       100000,
		From:            "2023-01-01",
		Limit:           10,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	results := parsePaginatedResults(t, result)
	if len(results) == 0 {
		t.Fatal("expected results")
	}

	t.Logf("Combined filter total=%d, returned=%d", results[0].Total, countValues(t, results[0].Values))
}

// --- Short Holdings ---

func parsePaginatedSingleResult(t *testing.T, result *mcp.CallToolResult) paginatedResult {
	t.Helper()
	if result.IsError {
		for _, c := range result.Content {
			if tc, ok := c.(*mcp.TextContent); ok {
				t.Fatalf("tool error: %s", tc.Text)
			}
		}
	}
	var text string
	for _, c := range result.Content {
		if tc, ok := c.(*mcp.TextContent); ok {
			text = tc.Text
			break
		}
	}
	var r paginatedResult
	if err := json.Unmarshal([]byte(text), &r); err != nil {
		t.Fatalf("unmarshal paginated result: %v (text: %.200s)", err, text)
	}
	return r
}

func TestIntegrationShortHoldingsDefaultPagination(t *testing.T) {
	h := integrationHandlers(t)
	ctx := context.Background()

	result, _, err := h.getShortHoldings(ctx, nil, shortHoldingsInput{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	r := parsePaginatedSingleResult(t, result)

	t.Logf("Total=%d, Offset=%d, Limit=%d, Returned=%d",
		r.Total, r.Offset, r.Limit, countValues(t, r.Values))

	if r.Total == 0 {
		t.Fatal("expected at least one instrument with short positions")
	}
	if r.Limit != 50 {
		t.Errorf("expected default limit=50, got %d", r.Limit)
	}
	returned := countValues(t, r.Values)
	if r.Total > 50 && returned != 50 {
		t.Errorf("expected 50 returned when total=%d, got %d", r.Total, returned)
	}
}

func TestIntegrationShortHoldingsFilterByMinPercent(t *testing.T) {
	h := integrationHandlers(t)
	ctx := context.Background()

	filteredResult, _, _ := h.getShortHoldings(ctx, nil, shortHoldingsInput{
		MinShortingPercent: 1.0,
	})
	allResult, _, _ := h.getShortHoldings(ctx, nil, shortHoldingsInput{})

	filtered := parsePaginatedSingleResult(t, filteredResult)
	all := parsePaginatedSingleResult(t, allResult)

	t.Logf("Instruments with >= 1.0%% shorts: %d, All: %d", filtered.Total, all.Total)

	if filtered.Total > all.Total {
		t.Errorf("filtered total (%d) should be <= all total (%d)", filtered.Total, all.Total)
	}
}

func TestIntegrationShortHoldingsOffset(t *testing.T) {
	h := integrationHandlers(t)
	ctx := context.Background()

	result1, _, _ := h.getShortHoldings(ctx, nil, shortHoldingsInput{
		Limit: 5,
	})
	result2, _, _ := h.getShortHoldings(ctx, nil, shortHoldingsInput{
		Limit:  5,
		Offset: 5,
	})

	page1 := parsePaginatedSingleResult(t, result1)
	page2 := parsePaginatedSingleResult(t, result2)

	t.Logf("Page1: Total=%d, Returned=%d", page1.Total, countValues(t, page1.Values))
	t.Logf("Page2: Total=%d, Returned=%d", page2.Total, countValues(t, page2.Values))

	if page1.Total != page2.Total {
		t.Errorf("total should be consistent: %d vs %d", page1.Total, page2.Total)
	}
	if page2.Offset != 5 {
		t.Errorf("expected offset=5 on page2, got %d", page2.Offset)
	}
}
