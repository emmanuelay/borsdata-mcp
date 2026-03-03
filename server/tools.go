package server

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/emmanuelay/borsdata-mcp/client"
	"github.com/emmanuelay/borsdata-mcp/models"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type handlers struct {
	client *client.Client
}

func textResult(v any) *mcp.CallToolResult {
	data, _ := json.MarshalIndent(v, "", "  ")
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: string(data)},
		},
	}
}

func errorResult(msg string) *mcp.CallToolResult {
	r := &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: msg},
		},
		IsError: true,
	}
	return r
}

type emptyInput struct{}

type scopeInput struct {
	Scope string `json:"scope,omitempty" jsonschema:"Market scope: nordic or global (default: nordic)"`
}

func isGlobal(scope string) bool {
	return strings.EqualFold(scope, "global")
}

func (h *handlers) listInstruments(ctx context.Context, request *mcp.CallToolRequest, input scopeInput) (*mcp.CallToolResult, any, error) {
	var resp *models.InstrumentsResponse
	var err error
	if isGlobal(input.Scope) {
		resp, err = h.client.GetGlobalInstruments()
	} else {
		resp, err = h.client.GetInstruments()
	}
	if err != nil {
		return errorResult(fmt.Sprintf("Failed to get instruments: %v", err)), nil, nil
	}
	return textResult(resp.Instruments), nil, nil
}

func (h *handlers) getInstrumentsUpdated(ctx context.Context, request *mcp.CallToolRequest, _ emptyInput) (*mcp.CallToolResult, any, error) {
	resp, err := h.client.GetInstrumentsUpdated()
	if err != nil {
		return errorResult(fmt.Sprintf("Failed to get instruments updated: %v", err)), nil, nil
	}
	return textResult(resp.Instruments), nil, nil
}

func (h *handlers) getStockSplits(ctx context.Context, request *mcp.CallToolRequest, _ emptyInput) (*mcp.CallToolResult, any, error) {
	resp, err := h.client.GetStockSplits()
	if err != nil {
		return errorResult(fmt.Sprintf("Failed to get stock splits: %v", err)), nil, nil
	}
	return textResult(resp.StockSplits), nil, nil
}

type searchInput struct {
	Query string `json:"query" jsonschema:"Search query to match against instrument name or ticker"`
	Scope string `json:"scope,omitempty" jsonschema:"Market scope: nordic or global (default: nordic)"`
}

func (h *handlers) searchInstruments(ctx context.Context, request *mcp.CallToolRequest, input searchInput) (*mcp.CallToolResult, any, error) {
	query := strings.ToLower(input.Query)
	if query == "" {
		return errorResult("query parameter is required"), nil, nil
	}

	var resp *models.InstrumentsResponse
	var err error
	if isGlobal(input.Scope) {
		resp, err = h.client.GetGlobalInstruments()
	} else {
		resp, err = h.client.GetInstruments()
	}
	if err != nil {
		return errorResult(fmt.Sprintf("Failed to get instruments: %v", err)), nil, nil
	}

	var matches []any
	for _, ins := range resp.Instruments {
		nameLower := strings.ToLower(ins.Name)
		tickerLower := strings.ToLower(ins.Ticker)
		if strings.Contains(nameLower, query) || strings.Contains(tickerLower, query) {
			matches = append(matches, ins)
		}
	}

	if len(matches) == 0 {
		return textResult(map[string]string{"message": fmt.Sprintf("No instruments found matching '%s'", query)}), nil, nil
	}
	return textResult(matches), nil, nil
}

type stockPricesInput struct {
	InstrumentID int64  `json:"instrument_id" jsonschema:"The instrument ID"`
	MaxCount     int    `json:"max_count,omitempty" jsonschema:"Maximum number of price entries to return (default: 20)"`
	From         string `json:"from,omitempty" jsonschema:"Start date in YYYY-MM-DD format (optional)"`
	To           string `json:"to,omitempty" jsonschema:"End date in YYYY-MM-DD format (optional)"`
}

func (h *handlers) getStockPrices(ctx context.Context, request *mcp.CallToolRequest, input stockPricesInput) (*mcp.CallToolResult, any, error) {
	if input.InstrumentID == 0 {
		return errorResult("instrument_id is required"), nil, nil
	}

	if input.From != "" && input.To != "" {
		resp, err := h.client.GetStockPricesDateRange(input.InstrumentID, input.From, input.To)
		if err != nil {
			return errorResult(fmt.Sprintf("Failed to get stock prices: %v", err)), nil, nil
		}
		return textResult(resp.StockPricesList), nil, nil
	}

	maxCount := input.MaxCount
	if maxCount == 0 {
		maxCount = 20
	}
	resp, err := h.client.GetStockPrices(input.InstrumentID, maxCount)
	if err != nil {
		return errorResult(fmt.Sprintf("Failed to get stock prices: %v", err)), nil, nil
	}
	return textResult(resp.StockPricesList), nil, nil
}

func (h *handlers) getLatestStockPrices(ctx context.Context, request *mcp.CallToolRequest, input scopeInput) (*mcp.CallToolResult, any, error) {
	var resp *models.StockPricesLastResponse
	var err error
	if isGlobal(input.Scope) {
		resp, err = h.client.GetGlobalLatestStockPrices()
	} else {
		resp, err = h.client.GetLatestStockPrices()
	}
	if err != nil {
		return errorResult(fmt.Sprintf("Failed to get latest stock prices: %v", err)), nil, nil
	}
	return textResult(resp.StockPricesList), nil, nil
}

type stockPricesByDateInput struct {
	Date  string `json:"date" jsonschema:"Date in YYYY-MM-DD format"`
	Scope string `json:"scope,omitempty" jsonschema:"Market scope: nordic or global (default: nordic)"`
}

func (h *handlers) getStockPricesByDate(ctx context.Context, request *mcp.CallToolRequest, input stockPricesByDateInput) (*mcp.CallToolResult, any, error) {
	if input.Date == "" {
		return errorResult("date is required"), nil, nil
	}
	var resp *models.StockPricesDateResponse
	var err error
	if isGlobal(input.Scope) {
		resp, err = h.client.GetGlobalStockPricesByDate(input.Date)
	} else {
		resp, err = h.client.GetStockPricesByDate(input.Date)
	}
	if err != nil {
		return errorResult(fmt.Sprintf("Failed to get stock prices by date: %v", err)), nil, nil
	}
	return textResult(resp.StockPricesList), nil, nil
}

type reportsInput struct {
	InstrumentID int64  `json:"instrument_id" jsonschema:"The instrument ID"`
	ReportType   string `json:"report_type,omitempty" jsonschema:"Type of report: year or r12 or quarter or all (default: all)"`
	MaxCount     int    `json:"max_count,omitempty" jsonschema:"Maximum number of reports to return (default: 10)"`
}

func (h *handlers) getReports(ctx context.Context, request *mcp.CallToolRequest, input reportsInput) (*mcp.CallToolResult, any, error) {
	if input.InstrumentID == 0 {
		return errorResult("instrument_id is required"), nil, nil
	}

	reportType := input.ReportType
	if reportType == "" {
		reportType = "all"
	}

	maxCount := input.MaxCount
	if maxCount == 0 {
		maxCount = 10
	}

	switch reportType {
	case "year":
		resp, err := h.client.GetYearlyReports(input.InstrumentID, maxCount)
		if err != nil {
			return errorResult(fmt.Sprintf("Failed to get yearly reports: %v", err)), nil, nil
		}
		return textResult(resp.Reports), nil, nil
	case "r12":
		resp, err := h.client.GetR12Reports(input.InstrumentID, maxCount)
		if err != nil {
			return errorResult(fmt.Sprintf("Failed to get R12 reports: %v", err)), nil, nil
		}
		return textResult(resp.Reports), nil, nil
	case "quarter":
		resp, err := h.client.GetQuarterlyReports(input.InstrumentID, maxCount)
		if err != nil {
			return errorResult(fmt.Sprintf("Failed to get quarterly reports: %v", err)), nil, nil
		}
		return textResult(resp.Reports), nil, nil
	case "all":
		resp, err := h.client.GetAllReports(input.InstrumentID, maxCount, maxCount)
		if err != nil {
			return errorResult(fmt.Sprintf("Failed to get reports: %v", err)), nil, nil
		}
		return textResult(resp), nil, nil
	default:
		return errorResult(fmt.Sprintf("Unknown report type: %s (use 'year', 'r12', 'quarter', or 'all')", reportType)), nil, nil
	}
}

func (h *handlers) getKpisUpdated(ctx context.Context, request *mcp.CallToolRequest, _ emptyInput) (*mcp.CallToolResult, any, error) {
	resp, err := h.client.GetKpisUpdated()
	if err != nil {
		return errorResult(fmt.Sprintf("Failed to get KPIs updated: %v", err)), nil, nil
	}
	return textResult(resp.Instruments), nil, nil
}

type kpiHistoryInput struct {
	InstrumentID int64  `json:"instrument_id" jsonschema:"The instrument ID"`
	KpiID        int    `json:"kpi_id" jsonschema:"The KPI ID (use get_metadata with type kpis to list available KPIs)"`
	ReportType   string `json:"report_type,omitempty" jsonschema:"Report type: year, r12, or quarter (default: quarter)"`
	PriceType    string `json:"price_type,omitempty" jsonschema:"Price type: mean, high, low, or latest (default: mean)"`
	MaxCount     int    `json:"max_count,omitempty" jsonschema:"Maximum number of history entries to return (default: API default)"`
}

func (h *handlers) getKpiHistory(ctx context.Context, request *mcp.CallToolRequest, input kpiHistoryInput) (*mcp.CallToolResult, any, error) {
	if input.InstrumentID == 0 || input.KpiID == 0 {
		return errorResult("instrument_id and kpi_id are required"), nil, nil
	}

	reportType := input.ReportType
	if reportType == "" {
		reportType = "quarter"
	}
	priceType := input.PriceType
	if priceType == "" {
		priceType = "mean"
	}

	resp, err := h.client.GetKpiHistory(input.InstrumentID, input.KpiID, reportType, priceType, input.MaxCount)
	if err != nil {
		return errorResult(fmt.Sprintf("Failed to get KPI history: %v", err)), nil, nil
	}
	return textResult(resp.Values), nil, nil
}

type kpiScreenerInput struct {
	InstrumentID int64  `json:"instrument_id" jsonschema:"The instrument ID"`
	KpiID        int    `json:"kpi_id" jsonschema:"The KPI ID"`
	CalcGroup    string `json:"calc_group,omitempty" jsonschema:"Calculation group: 1year, 3year, 5year, 7year, 10year, or 15year (default: 1year)"`
	CalcType     string `json:"calc_type,omitempty" jsonschema:"Calculation type: mean, high, low, sum, latest, or count (default: mean)"`
}

func (h *handlers) getKpiScreener(ctx context.Context, request *mcp.CallToolRequest, input kpiScreenerInput) (*mcp.CallToolResult, any, error) {
	if input.InstrumentID == 0 || input.KpiID == 0 {
		return errorResult("instrument_id and kpi_id are required"), nil, nil
	}
	calcGroup := input.CalcGroup
	if calcGroup == "" {
		calcGroup = "1year"
	}
	calcType := input.CalcType
	if calcType == "" {
		calcType = "mean"
	}
	resp, err := h.client.GetKpiScreener(input.InstrumentID, input.KpiID, calcGroup, calcType)
	if err != nil {
		return errorResult(fmt.Sprintf("Failed to get KPI screener: %v", err)), nil, nil
	}
	return textResult(resp.Value), nil, nil
}

type kpiListScreenerInput struct {
	KpiID     int    `json:"kpi_id" jsonschema:"The KPI ID"`
	CalcGroup string `json:"calc_group,omitempty" jsonschema:"Calculation group: 1year, 3year, 5year, 7year, 10year, or 15year (default: 1year)"`
	CalcType  string `json:"calc_type,omitempty" jsonschema:"Calculation type: mean, high, low, sum, latest, or count (default: mean)"`
	Scope     string `json:"scope,omitempty" jsonschema:"Market scope: nordic or global (default: nordic)"`
}

func (h *handlers) getKpiListScreener(ctx context.Context, request *mcp.CallToolRequest, input kpiListScreenerInput) (*mcp.CallToolResult, any, error) {
	if input.KpiID == 0 {
		return errorResult("kpi_id is required"), nil, nil
	}
	calcGroup := input.CalcGroup
	if calcGroup == "" {
		calcGroup = "1year"
	}
	calcType := input.CalcType
	if calcType == "" {
		calcType = "mean"
	}
	var resp *models.KpisAllCompResponse
	var err error
	if isGlobal(input.Scope) {
		resp, err = h.client.GetGlobalKpiListScreener(input.KpiID, calcGroup, calcType)
	} else {
		resp, err = h.client.GetKpiListScreener(input.KpiID, calcGroup, calcType)
	}
	if err != nil {
		return errorResult(fmt.Sprintf("Failed to get KPI list screener: %v", err)), nil, nil
	}
	return textResult(resp.Values), nil, nil
}

type kpiSummaryInput struct {
	InstrumentID int64  `json:"instrument_id" jsonschema:"The instrument ID"`
	ReportType   string `json:"report_type,omitempty" jsonschema:"Report type: year, r12, or quarter (default: year)"`
	MaxCount     int    `json:"max_count,omitempty" jsonschema:"Maximum number of periods to return (default: 10)"`
}

func (h *handlers) getKpiSummary(ctx context.Context, request *mcp.CallToolRequest, input kpiSummaryInput) (*mcp.CallToolResult, any, error) {
	if input.InstrumentID == 0 {
		return errorResult("instrument_id is required"), nil, nil
	}

	reportType := input.ReportType
	if reportType == "" {
		reportType = "year"
	}

	maxCount := input.MaxCount
	if maxCount == 0 {
		maxCount = 10
	}

	resp, err := h.client.GetKpiSummary(input.InstrumentID, reportType, maxCount)
	if err != nil {
		return errorResult(fmt.Sprintf("Failed to get KPI summary: %v", err)), nil, nil
	}
	return textResult(resp.Kpis), nil, nil
}

type instrumentListInput struct {
	InstrumentIDs string `json:"instrument_ids" jsonschema:"Comma-separated instrument IDs (max 50)"`
}

func (h *handlers) getInsiderHoldings(ctx context.Context, request *mcp.CallToolRequest, input instrumentListInput) (*mcp.CallToolResult, any, error) {
	if input.InstrumentIDs == "" {
		return errorResult("instrument_ids is required"), nil, nil
	}
	resp, err := h.client.GetInsiderHoldings(input.InstrumentIDs)
	if err != nil {
		return errorResult(fmt.Sprintf("Failed to get insider holdings: %v", err)), nil, nil
	}
	return textResult(resp.List), nil, nil
}

func (h *handlers) getShortHoldings(ctx context.Context, request *mcp.CallToolRequest, _ emptyInput) (*mcp.CallToolResult, any, error) {
	resp, err := h.client.GetShortHoldings()
	if err != nil {
		return errorResult(fmt.Sprintf("Failed to get short holdings: %v", err)), nil, nil
	}
	// Filter out entries where Values is nil/empty
	var filtered []models.ShortsRespV1
	for _, item := range resp.List {
		if len(item.Values) > 0 {
			filtered = append(filtered, item)
		}
	}
	return textResult(filtered), nil, nil
}

type instrumentListWithMaxCountInput struct {
	InstrumentIDs string `json:"instrument_ids" jsonschema:"Comma-separated instrument IDs (max 50)"`
	MaxCount      int    `json:"max_count,omitempty" jsonschema:"Maximum number of entries per instrument (default: 20)"`
}

func (h *handlers) getBuybackHoldings(ctx context.Context, request *mcp.CallToolRequest, input instrumentListWithMaxCountInput) (*mcp.CallToolResult, any, error) {
	if input.InstrumentIDs == "" {
		return errorResult("instrument_ids is required"), nil, nil
	}
	resp, err := h.client.GetBuybackHoldings(input.InstrumentIDs)
	if err != nil {
		return errorResult(fmt.Sprintf("Failed to get buyback holdings: %v", err)), nil, nil
	}
	maxCount := input.MaxCount
	if maxCount == 0 {
		maxCount = 20
	}
	for i := range resp.List {
		if len(resp.List[i].Values) > maxCount {
			resp.List[i].Values = resp.List[i].Values[:maxCount]
		}
	}
	return textResult(resp.List), nil, nil
}

func (h *handlers) getReportCalendar(ctx context.Context, request *mcp.CallToolRequest, input instrumentListWithMaxCountInput) (*mcp.CallToolResult, any, error) {
	if input.InstrumentIDs == "" {
		return errorResult("instrument_ids is required"), nil, nil
	}
	resp, err := h.client.GetReportCalendar(input.InstrumentIDs)
	if err != nil {
		return errorResult(fmt.Sprintf("Failed to get report calendar: %v", err)), nil, nil
	}
	maxCount := input.MaxCount
	if maxCount == 0 {
		maxCount = 20
	}
	for i := range resp.List {
		if len(resp.List[i].Values) > maxCount {
			resp.List[i].Values = resp.List[i].Values[:maxCount]
		}
	}
	return textResult(resp.List), nil, nil
}

func (h *handlers) getDividendCalendar(ctx context.Context, request *mcp.CallToolRequest, input instrumentListWithMaxCountInput) (*mcp.CallToolResult, any, error) {
	if input.InstrumentIDs == "" {
		return errorResult("instrument_ids is required"), nil, nil
	}
	resp, err := h.client.GetDividendCalendar(input.InstrumentIDs)
	if err != nil {
		return errorResult(fmt.Sprintf("Failed to get dividend calendar: %v", err)), nil, nil
	}
	maxCount := input.MaxCount
	if maxCount == 0 {
		maxCount = 20
	}
	for i := range resp.List {
		if len(resp.List[i].Values) > maxCount {
			resp.List[i].Values = resp.List[i].Values[:maxCount]
		}
	}
	return textResult(resp.List), nil, nil
}

type metadataInput struct {
	Type string `json:"type" jsonschema:"Type of metadata: countries or markets or sectors or branches or kpis or reports or translations"`
}

func (h *handlers) getMetadata(ctx context.Context, request *mcp.CallToolRequest, input metadataInput) (*mcp.CallToolResult, any, error) {
	switch input.Type {
	case "countries":
		resp, err := h.client.GetCountries()
		if err != nil {
			return errorResult(fmt.Sprintf("Failed to get countries: %v", err)), nil, nil
		}
		return textResult(resp.Countries), nil, nil
	case "markets":
		resp, err := h.client.GetMarkets()
		if err != nil {
			return errorResult(fmt.Sprintf("Failed to get markets: %v", err)), nil, nil
		}
		return textResult(resp.Markets), nil, nil
	case "sectors":
		resp, err := h.client.GetSectors()
		if err != nil {
			return errorResult(fmt.Sprintf("Failed to get sectors: %v", err)), nil, nil
		}
		return textResult(resp.Sectors), nil, nil
	case "branches":
		resp, err := h.client.GetBranches()
		if err != nil {
			return errorResult(fmt.Sprintf("Failed to get branches: %v", err)), nil, nil
		}
		return textResult(resp.Branches), nil, nil
	case "kpis":
		resp, err := h.client.GetKpiMetadata()
		if err != nil {
			return errorResult(fmt.Sprintf("Failed to get KPI metadata: %v", err)), nil, nil
		}
		return textResult(resp.KpiHistoryMetadatas), nil, nil
	case "reports":
		resp, err := h.client.GetReportMetadata()
		if err != nil {
			return errorResult(fmt.Sprintf("Failed to get report metadata: %v", err)), nil, nil
		}
		return textResult(resp.ReportMetadatas), nil, nil
	case "translations":
		resp, err := h.client.GetTranslationMetadata()
		if err != nil {
			return errorResult(fmt.Sprintf("Failed to get translation metadata: %v", err)), nil, nil
		}
		return textResult(resp.TranslationMetadatas), nil, nil
	default:
		return errorResult(fmt.Sprintf("Unknown metadata type: %s (use 'countries', 'markets', 'sectors', 'branches', 'kpis', 'reports', or 'translations')", input.Type)), nil, nil
	}
}
