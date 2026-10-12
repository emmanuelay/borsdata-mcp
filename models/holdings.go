package models

// InsiderRowV1 represents a single insider transaction with owner details, share count, price, and transaction date.
type InsiderRowV1 struct {
	Misc             bool    `json:"misc"`
	OwnerName        string  `json:"ownerName"`
	OwnerPosition    string  `json:"ownerPosition"`
	EquityProgram    bool    `json:"equityProgram"`
	Shares           int64   `json:"shares"`
	Price            float64 `json:"price"`
	Amount           float64 `json:"amount"`
	Currency         string  `json:"currency"`
	TransactionType  int     `json:"transactionType"`
	VerificationDate string  `json:"verificationDate"`
	TransactionDate  string  `json:"transactionDate"`
}

// InsiderRespV1 represents insider transactions for a single instrument.
type InsiderRespV1 struct {
	InsID  int64          `json:"insId"`
	Values []InsiderRowV1 `json:"values"`
	Error  string         `json:"error"`
}

// InsiderHoldingsResponse is the top-level response from the insider holdings endpoint.
type InsiderHoldingsResponse struct {
	List []InsiderRespV1 `json:"list"`
}

// ShortsRespV1 represents short position summary for a single instrument.
type ShortsRespV1 struct {
	InsID               int64    `json:"insId"`
	ShortsProc          *float64 `json:"shortsProc"`
	ShortsHolders       *float64 `json:"shortsHolders"`
	ShortsAvgProc       *float64 `json:"shortsAvgProc"`
	ShortsMilj          *float64 `json:"shortsMilj"`
	ShortsAvgMilj       *float64 `json:"shortsAvgMilj"`
	LastTransactionDate *string  `json:"lastTransactionDate"`
	DtcSum              *float64 `json:"dtcSum"`
	DtcAvg              *float64 `json:"dtcAvg"`
	Trend1w             float64  `json:"trend1w"`
	Trend1m             float64  `json:"trend1m"`
	Trend3m             float64  `json:"trend3m"`
	Trend6m             float64  `json:"trend6m"`
}

// ShortHoldingsResponse is the top-level response from the short holdings endpoint.
type ShortHoldingsResponse struct {
	List []ShortsRespV1 `json:"list"`
}

// BuybackRowV1 represents a single share buyback entry with share change, price, and date.
type BuybackRowV1 struct {
	Change     int64   `json:"change"`
	ChangeProc float64 `json:"changeProc"`
	Price      float64 `json:"price"`
	Currency   string  `json:"currency"`
	Shares     int64   `json:"shares"`
	SharesProc float64 `json:"sharesProc"`
	Date       string  `json:"date"`
}

// BuybackRespV1 represents share buyback data for a single instrument.
type BuybackRespV1 struct {
	InsID  int64          `json:"insId"`
	Values []BuybackRowV1 `json:"values"`
	Error  string         `json:"error"`
}

// BuybackHoldingsResponse is the top-level response from the buyback holdings endpoint.
type BuybackHoldingsResponse struct {
	List []BuybackRespV1 `json:"list"`
}
