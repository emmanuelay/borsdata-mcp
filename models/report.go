package models

type ReportV1 struct {
	Year                int     `json:"year"`
	Period              int     `json:"period"`
	Revenues            float64 `json:"revenues"`
	Gross_Income        float64 `json:"gross_Income"`
	Operating_Income    float64 `json:"operating_Income"`
	Profit_Before_Tax   float64 `json:"profit_Before_Tax"`
	Profit_To_Equity_Holders float64 `json:"profit_To_Equity_Holders"`
	Earnings_Per_Share  float64 `json:"earnings_Per_Share"`
	NumberOfShares      float64 `json:"numberOfShares"`
	Dividend            float64 `json:"dividend"`
	IntangibleAssets    float64 `json:"intangibleAssets"`
	TangibleAssets      float64 `json:"tangibleAssets"`
	FinancialAssets     float64 `json:"financialAssets"`
	NonCurrentAssets    float64 `json:"nonCurrentAssets"`
	CashAndEquivalents  float64 `json:"cashAndEquivalents"`
	CurrentAssets       float64 `json:"currentAssets"`
	TotalAssets         float64 `json:"totalAssets"`
	TotalEquity         float64 `json:"totalEquity"`
	NonCurrentLiabilities float64 `json:"nonCurrentLiabilities"`
	CurrentLiabilities  float64 `json:"currentLiabilities"`
	TotalLiabilitiesAndEquity float64 `json:"totalLiabilitiesAndEquity"`
	Net_Debt            float64 `json:"net_Debt"`
	ReportStartDate     string  `json:"report_Start_Date"`
	ReportEndDate       string  `json:"report_End_Date"`
	Currency            string  `json:"currency"`
}

type ReportsResponse struct {
	Reports []ReportV1 `json:"reports"`
}

type ReportsCompoundResponse struct {
	ReportsYear    []ReportV1 `json:"reportsYear"`
	ReportsR12     []ReportV1 `json:"reportsR12"`
	ReportsQuarter []ReportV1 `json:"reportsQuarter"`
}
