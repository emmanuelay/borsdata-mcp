package models

// StockPriceV1 represents a daily stock price entry with OHLCV data for a single instrument.
type StockPriceV1 struct {
	D string  `json:"d"` // Date
	C float64 `json:"c"` // Close
	H float64 `json:"h"` // High
	L float64 `json:"l"` // Low
	O float64 `json:"o"` // Open
	V int64   `json:"v"` // Volume
}

// StockPricesResponse is the top-level response from the stock prices endpoint for a single instrument.
type StockPricesResponse struct {
	StockPricesList []StockPriceV1 `json:"stockPricesList"`
}

// StockPriceLastV1 represents the latest stock price for an instrument, including the instrument ID and OHLCV data.
type StockPriceLastV1 struct {
	I int64   `json:"i"` // Instrument ID
	D string  `json:"d"` // Date
	C float64 `json:"c"` // Close
	H float64 `json:"h"` // High
	L float64 `json:"l"` // Low
	O float64 `json:"o"` // Open
	V int64   `json:"v"` // Volume
}

// StockPricesLastResponse is the top-level response from the latest stock prices endpoint.
type StockPricesLastResponse struct {
	StockPricesList []StockPriceLastV1 `json:"stockPricesList"`
}

// StockPriceDateV1 represents a stock price for an instrument on a specific date, including the instrument ID and OHLCV data.
type StockPriceDateV1 struct {
	I int64   `json:"i"` // Instrument ID
	D string  `json:"d"` // Date
	C float64 `json:"c"` // Close
	H float64 `json:"h"` // High
	L float64 `json:"l"` // Low
	O float64 `json:"o"` // Open
	V int64   `json:"v"` // Volume
}

// StockPricesDateResponse is the top-level response from the stock prices by date endpoint.
type StockPricesDateResponse struct {
	StockPricesList []StockPriceDateV1 `json:"stockPricesList"`
}
