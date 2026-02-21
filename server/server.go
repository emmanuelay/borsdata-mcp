package server

import (
	"github.com/emmanuelay/borsdata-mcp/client"
	"github.com/google/jsonschema-go/jsonschema"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// New creates a new MCP server with all Börsdata tools registered.
func New(apiClient *client.Client, version string) *mcp.Server {
	s := mcp.NewServer(&mcp.Implementation{
		Name:    "borsdata-mcp",
		Version: version,
	}, &mcp.ServerOptions{
		Capabilities: &mcp.ServerCapabilities{
			Tools: &mcp.ToolCapabilities{ListChanged: false},
		},
	})

	h := &handlers{client: apiClient}
	registerTools(s, h)
	return s
}

var emptySchema = &jsonschema.Schema{Type: "object"}

func registerTools(s *mcp.Server, h *handlers) {
	mcp.AddTool(s, &mcp.Tool{
		Name:        "list_instruments",
		Description: "List all instruments (companies) available in Börsdata. Returns instrument ID, name, ticker, sector, market, branch, and country. Use scope 'global' for non-Nordic instruments (requires Pro+).",
	}, h.listInstruments)

	mcp.AddTool(s, &mcp.Tool{
		Name:        "search_instruments",
		Description: "Search instruments by name or ticker. Performs a case-insensitive substring match. Use scope 'global' for non-Nordic instruments (requires Pro+).",
	}, h.searchInstruments)

	s.AddTool(&mcp.Tool{
		Name:        "get_instruments_updated",
		Description: "Get update timestamps for all instruments. Shows when each instrument was last updated.",
		InputSchema: emptySchema,
	}, h.getInstrumentsUpdated)

	s.AddTool(&mcp.Tool{
		Name:        "get_stock_splits",
		Description: "Get all stock splits. Returns instrument ID, split date, split type, and ratio.",
		InputSchema: emptySchema,
	}, h.getStockSplits)

	mcp.AddTool(s, &mcp.Tool{
		Name:        "get_stock_prices",
		Description: "Get historical stock prices for a specific instrument. Returns date, open, high, low, close, and volume.",
	}, h.getStockPrices)

	mcp.AddTool(s, &mcp.Tool{
		Name:        "get_latest_stock_prices",
		Description: "Get the latest stock prices for all instruments. Use scope 'global' for non-Nordic instruments (requires Pro+).",
	}, h.getLatestStockPrices)

	mcp.AddTool(s, &mcp.Tool{
		Name:        "get_stock_prices_by_date",
		Description: "Get stock prices for all instruments on a specific date. Use scope 'global' for non-Nordic instruments (requires Pro+).",
	}, h.getStockPricesByDate)

	mcp.AddTool(s, &mcp.Tool{
		Name:        "get_reports",
		Description: "Get financial reports for an instrument. Supports yearly, R12 (rolling 12 months), and quarterly reports.",
	}, h.getReports)

	s.AddTool(&mcp.Tool{
		Name:        "get_kpis_updated",
		Description: "Get update timestamps for KPI calculations. Shows when KPIs were last recalculated for each instrument.",
		InputSchema: emptySchema,
	}, h.getKpisUpdated)

	mcp.AddTool(s, &mcp.Tool{
		Name:        "get_kpi_history",
		Description: "Get KPI history for an instrument. Supports configurable report type (year/r12/quarter) and price type (mean/high/low/latest).",
	}, h.getKpiHistory)

	mcp.AddTool(s, &mcp.Tool{
		Name:        "get_kpi_screener",
		Description: "Get a single KPI value for a specific instrument. calc_group must be one of: 1year, 3year, 5year, 7year, 10year, 15year. Other values like 'last' or 'quarter' are not supported and will return 400.",
	}, h.getKpiScreener)

	mcp.AddTool(s, &mcp.Tool{
		Name:        "get_kpi_list_screener",
		Description: "Get a specific KPI value for all instruments. Useful for screening/filtering instruments by a KPI. calc_group must be one of: 1year, 3year, 5year, 7year, 10year, 15year. Other values like 'last' or 'quarter' are not supported and will return 400. Use scope 'global' for non-Nordic instruments (requires Pro+).",
	}, h.getKpiListScreener)

	mcp.AddTool(s, &mcp.Tool{
		Name:        "get_kpi_summary",
		Description: "Get KPI summary for an instrument. Supports report type (year/r12/quarter).",
	}, h.getKpiSummary)

	mcp.AddTool(s, &mcp.Tool{
		Name:        "get_metadata",
		Description: "Get reference/metadata. Use this to look up names for IDs returned by other tools.",
	}, h.getMetadata)

	mcp.AddTool(s, &mcp.Tool{
		Name:        "get_insider_holdings",
		Description: "Get insider transactions for specified instruments. Returns owner name, position, shares, price, amount, currency, transaction type, and dates.",
	}, h.getInsiderHoldings)

	s.AddTool(&mcp.Tool{
		Name:        "get_short_holdings",
		Description: "Get short positions for all instruments. Returns position holder, ISIN, shorting percent, notified capital, and report date. Only Nordic instruments.",
		InputSchema: emptySchema,
	}, h.getShortHoldings)

	mcp.AddTool(s, &mcp.Tool{
		Name:        "get_buyback_holdings",
		Description: "Get share buybacks for specified instruments. Returns change, price, currency, shares, and date. Use max_count to limit entries per instrument (default: 20).",
	}, h.getBuybackHoldings)

	mcp.AddTool(s, &mcp.Tool{
		Name:        "get_report_calendar",
		Description: "Get report release dates for specified instruments. Returns release date and report type. Use max_count to limit entries per instrument (default: 20).",
	}, h.getReportCalendar)

	mcp.AddTool(s, &mcp.Tool{
		Name:        "get_dividend_calendar",
		Description: "Get dividend dates for specified instruments. Returns amount paid, currency, distribution frequency, excluding date, and dividend type. Use max_count to limit entries per instrument (default: 20).",
	}, h.getDividendCalendar)
}
