package models

// InstrumentV1 represents a financial instrument (stock) with identifiers, ticker, ISIN, and classification IDs.
type InstrumentV1 struct {
	InsID        int64  `json:"insId"`
	Name         string `json:"name"`
	URLName      string `json:"urlName"`
	Instrument   int    `json:"instrument"`
	ISIN         string `json:"isin"`
	Ticker       string `json:"ticker"`
	Yahoo        string `json:"yahoo"`
	SectorID     int    `json:"sectorId"`
	MarketID     int    `json:"marketId"`
	BranchID     int    `json:"branchId"`
	CountryID    int    `json:"countryId"`
	ListingDate  string `json:"listingDate"`
	StockPriceID int64  `json:"stockPriceId"`
}

// InstrumentsResponse is the top-level response from the instruments listing endpoint.
type InstrumentsResponse struct {
	Instruments []InstrumentV1 `json:"instruments"`
}

// InstrumentUpdatedV1 represents the last update timestamp for an instrument.
type InstrumentUpdatedV1 struct {
	InsID       int64  `json:"insId"`
	UpdatedDate string `json:"updatedDate"`
}

// InstrumentsUpdatedResponse is the top-level response from the instruments updated endpoint.
type InstrumentsUpdatedResponse struct {
	Instruments []InstrumentUpdatedV1 `json:"instruments"`
}

// StockSplitV1 represents a stock split event with the split date, type, and ratio.
type StockSplitV1 struct {
	InsID     int64   `json:"insId"`
	SplitDate string  `json:"splitDate"`
	SplitType string  `json:"splitType"`
	Ratio     float64 `json:"ratio"`
}

// StockSplitsResponse is the top-level response from the stock splits endpoint.
type StockSplitsResponse struct {
	StockSplits []StockSplitV1 `json:"stockSplitList"`
}
