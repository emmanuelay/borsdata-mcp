package models

type StockPriceV1 struct {
	D string  `json:"d"` // Date
	C float64 `json:"c"` // Close
	H float64 `json:"h"` // High
	L float64 `json:"l"` // Low
	O float64 `json:"o"` // Open
	V int64   `json:"v"` // Volume
}

type StockPricesResponse struct {
	StockPricesList []StockPriceV1 `json:"stockPricesList"`
}

type StockPriceLastV1 struct {
	I int64   `json:"i"` // Instrument ID
	D string  `json:"d"` // Date
	C float64 `json:"c"` // Close
	H float64 `json:"h"` // High
	L float64 `json:"l"` // Low
	O float64 `json:"o"` // Open
	V int64   `json:"v"` // Volume
}

type StockPricesLastResponse struct {
	StockPricesList []StockPriceLastV1 `json:"stockPricesList"`
}

type StockPriceDateV1 struct {
	I int64   `json:"i"` // Instrument ID
	D string  `json:"d"` // Date
	C float64 `json:"c"` // Close
	H float64 `json:"h"` // High
	L float64 `json:"l"` // Low
	O float64 `json:"o"` // Open
	V int64   `json:"v"` // Volume
}

type StockPricesDateResponse struct {
	StockPricesList []StockPriceDateV1 `json:"stockPricesList"`
}
