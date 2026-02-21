package models

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

type InsiderRespV1 struct {
	InsID  int64          `json:"insId"`
	Values []InsiderRowV1 `json:"values"`
	Error  string         `json:"error"`
}

type InsiderHoldingsResponse struct {
	List []InsiderRespV1 `json:"list"`
}

type ShortsRowV1 struct {
	ReportID           int     `json:"reportId"`
	PositionHolderName string  `json:"positionHolderName"`
	ISIN               string  `json:"isin"`
	ShortingPercent    float64 `json:"shortingPercent"`
	NotifiedCapital    float64 `json:"notifiedCapital"`
	Currency           string  `json:"currency"`
	ReportDate         string  `json:"reportDate"`
}

type ShortsRespV1 struct {
	InsID  int64         `json:"insId"`
	Values []ShortsRowV1 `json:"values"`
	Error  string        `json:"error"`
}

type ShortHoldingsResponse struct {
	List []ShortsRespV1 `json:"list"`
}

type BuybackRowV1 struct {
	Change     int64   `json:"change"`
	ChangeProc float64 `json:"changeProc"`
	Price      float64 `json:"price"`
	Currency   string  `json:"currency"`
	Shares     int64   `json:"shares"`
	SharesProc float64 `json:"sharesProc"`
	Date       string  `json:"date"`
}

type BuybackRespV1 struct {
	InsID  int64          `json:"insId"`
	Values []BuybackRowV1 `json:"values"`
	Error  string         `json:"error"`
}

type BuybackHoldingsResponse struct {
	List []BuybackRespV1 `json:"list"`
}
