package models

// ReportCalendarDateV1 represents a single report release date entry with the date and report type (e.g. Q1, Q2, annual).
type ReportCalendarDateV1 struct {
	ReleaseDate string `json:"releaseDate"`
	ReportType  string `json:"reportType"`
}

// ReportCalendarRespV1 represents the report calendar for a single instrument, containing upcoming report release dates.
type ReportCalendarRespV1 struct {
	InsID  int64                  `json:"insId"`
	Values []ReportCalendarDateV1 `json:"values"`
	Error  string                 `json:"error"`
}

// ReportCalendarResponse is the top-level response from the report calendar endpoint.
type ReportCalendarResponse struct {
	List []ReportCalendarRespV1 `json:"list"`
}

// DividendDateV1 represents a single dividend payment entry with amount, currency, ex-date, and distribution frequency.
type DividendDateV1 struct {
	AmountPaid            float64 `json:"amountPaid"`
	CurrencyShortName     string  `json:"currencyShortName"`
	DistributionFrequency int     `json:"distributionFrequency"`
	ExcludingDate         string  `json:"excludingDate"`
	DividendType          int     `json:"dividendType"`
}

// DividendCalendarRespV1 represents the dividend calendar for a single instrument, containing upcoming dividend dates.
type DividendCalendarRespV1 struct {
	InsID  int64            `json:"insId"`
	Values []DividendDateV1 `json:"values"`
	Error  string           `json:"error"`
}

// DividendCalendarResponse is the top-level response from the dividend calendar endpoint.
type DividendCalendarResponse struct {
	List []DividendCalendarRespV1 `json:"list"`
}
