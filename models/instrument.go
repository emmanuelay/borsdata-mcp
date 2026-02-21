package models

type InstrumentV1 struct {
	InsID       int64   `json:"insId"`
	Name        string  `json:"name"`
	URLName     string  `json:"urlName"`
	Instrument  int     `json:"instrument"`
	ISIN        string  `json:"isin"`
	Ticker      string  `json:"ticker"`
	Yahoo       string  `json:"yahoo"`
	SectorID    int     `json:"sectorId"`
	MarketID    int     `json:"marketId"`
	BranchID    int     `json:"branchId"`
	CountryID   int     `json:"countryId"`
	ListingDate string  `json:"listingDate"`
	StockPriceID int64  `json:"stockPriceId"`
}

type InstrumentsResponse struct {
	Instruments []InstrumentV1 `json:"instruments"`
}

type InstrumentUpdatedV1 struct {
	InsID       int64  `json:"insId"`
	UpdatedDate string `json:"updatedDate"`
}

type InstrumentsUpdatedResponse struct {
	Instruments []InstrumentUpdatedV1 `json:"instruments"`
}

type StockSplitV1 struct {
	InsID       int64   `json:"insId"`
	SplitDate   string  `json:"splitDate"`
	SplitType   string  `json:"splitType"`
	Ratio       float64 `json:"ratio"`
}

type StockSplitsResponse struct {
	StockSplits []StockSplitV1 `json:"stockSplitList"`
}
