package models

// KpiV1 represents a single KPI value for an instrument, used in screener and list screener responses.
type KpiV1 struct {
	I int64   `json:"i"` // Instrument ID
	N int     `json:"n"` // KPI ID
	V float64 `json:"v"` // Value
}

// KpisResponse is the top-level response for KPI list screener queries.
type KpisResponse struct {
	Values []KpiV1 `json:"values"`
}

// KpiHistoryV1 represents a historical KPI data point with year, period, and value.
type KpiHistoryV1 struct {
	Y int     `json:"y"` // Year
	P int     `json:"p"` // Period
	V float64 `json:"v"` // Value
}

// KpisHistoryResponse is the top-level response from the KPI history endpoint.
type KpisHistoryResponse struct {
	Values []KpiHistoryV1 `json:"values"`
}

// KpiScreenerValue represents a single KPI screener value for one instrument.
// API returns: {"kpiId":2,"group":"1year","calculation":"mean","value":{"i":2,"n":20.95,"s":null}}
// N holds numeric KPIs; S holds the value for string KPIs (isString=true in KPI metadata).
type KpiScreenerValue struct {
	I int64    `json:"i"`
	N *float64 `json:"n"`
	S *string  `json:"s"`
}

// KpiScreenerResponse is the top-level response from the KPI screener endpoint, containing the KPI ID, group, calculation method, and value.
type KpiScreenerResponse struct {
	KpiID       int              `json:"kpiId"`
	Group       string           `json:"group"`
	Calculation string           `json:"calculation"`
	Value       KpiScreenerValue `json:"value"`
}

// KpisSummaryValueV1 represents a single KPI summary data point.
type KpisSummaryValueV1 struct {
	Y int     `json:"y"`
	P int     `json:"p"`
	V float64 `json:"v"`
}

// KpisSummaryGroupV1 represents a group of KPI summary values for a specific KPI ID.
type KpisSummaryGroupV1 struct {
	KpiID  int                  `json:"KpiId"`
	Values []KpisSummaryValueV1 `json:"values"`
}

// KpisSummaryResponse is the top-level response from the KPI summary endpoint for a single instrument.
type KpisSummaryResponse struct {
	Instrument int                  `json:"instrument"`
	ReportType string               `json:"reportType"`
	Kpis       []KpisSummaryGroupV1 `json:"kpis"`
}

// KpiMetadataV1 represents metadata for a KPI, including its ID, Swedish/English names, and abbreviations.
type KpiMetadataV1 struct {
	KpiID          int    `json:"kpiId"`
	NameSv         string `json:"nameSv"`
	NameEn         string `json:"nameEn"`
	AbbreviationSv string `json:"abbreviationSv"`
	AbbreviationEn string `json:"abbreviationEn"`
	IsString       bool   `json:"isString"`
}

// KpisMetadataResponse is the top-level response from the KPI metadata endpoint.
type KpisMetadataResponse struct {
	KpiHistoryMetadatas []KpiMetadataV1 `json:"kpiHistoryMetadatas"`
}

// KpisCalcUpdatedV1 represents the last KPI calculation update timestamp for an instrument.
type KpisCalcUpdatedV1 struct {
	InsID       int64  `json:"insId"`
	UpdatedDate string `json:"updatedDate"`
}

// KpisCalcUpdatedResponse is the top-level response from the KPI calculation updated endpoint.
type KpisCalcUpdatedResponse struct {
	Instruments []KpisCalcUpdatedV1 `json:"instruments"`
}

// KpiAllCompV1 represents a KPI value for a single instrument in an all-companies screener response.
type KpiAllCompV1 struct {
	I int64   `json:"i"` // Instrument ID
	N float64 `json:"n"` // Value
}

// KpisAllCompResponse is the top-level response from the KPI all-companies screener endpoint.
type KpisAllCompResponse struct {
	Values []KpiAllCompV1 `json:"values"`
}
