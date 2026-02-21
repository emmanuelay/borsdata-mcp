package models

type ReportCalendarDateV1 struct {
	ReleaseDate string `json:"releaseDate"`
	ReportType  string `json:"reportType"`
}

type ReportCalendarRespV1 struct {
	InsID  int64                  `json:"insId"`
	Values []ReportCalendarDateV1 `json:"values"`
	Error  string                 `json:"error"`
}

type ReportCalendarResponse struct {
	List []ReportCalendarRespV1 `json:"list"`
}

type DividendDateV1 struct {
	AmountPaid            float64 `json:"amountPaid"`
	CurrencyShortName     string  `json:"currencyShortName"`
	DistributionFrequency int     `json:"distributionFrequency"`
	ExcludingDate         string  `json:"excludingDate"`
	DividendType          int     `json:"dividendType"`
}

type DividendCalendarRespV1 struct {
	InsID  int64            `json:"insId"`
	Values []DividendDateV1 `json:"values"`
	Error  string           `json:"error"`
}

type DividendCalendarResponse struct {
	List []DividendCalendarRespV1 `json:"list"`
}
