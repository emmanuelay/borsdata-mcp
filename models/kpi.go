package models

type KpiV1 struct {
	I int64       `json:"i"` // Instrument ID
	N int         `json:"n"` // KPI ID
	V float64     `json:"v"` // Value
}

type KpisResponse struct {
	Values []KpiV1 `json:"values"`
}

type KpiHistoryV1 struct {
	Y int     `json:"y"` // Year
	P int     `json:"p"` // Period
	V float64 `json:"v"` // Value
}

type KpisHistoryResponse struct {
	Values []KpiHistoryV1 `json:"values"`
}

// KpiScreenerValue represents a single KPI screener value for one instrument.
// API returns: {"kpiId":2,"group":"1year","calculation":"mean","value":{"i":2,"n":20.95,"s":null}}
type KpiScreenerValue struct {
	I int64    `json:"i"`
	N float64  `json:"n"`
	S *float64 `json:"s"`
}

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

type KpisSummaryGroupV1 struct {
	KpiID  int                  `json:"KpiId"`
	Values []KpisSummaryValueV1 `json:"values"`
}

type KpisSummaryResponse struct {
	Instrument int                  `json:"instrument"`
	ReportType string               `json:"reportType"`
	Kpis       []KpisSummaryGroupV1 `json:"kpis"`
}

type KpiMetadataV1 struct {
	KpiID          int    `json:"kpiId"`
	NameSv         string `json:"nameSv"`
	NameEn         string `json:"nameEn"`
	AbbreviationSv string `json:"abbreviationSv"`
	AbbreviationEn string `json:"abbreviationEn"`
	IsString       bool   `json:"isString"`
}

type KpisMetadataResponse struct {
	KpiHistoryMetadatas []KpiMetadataV1 `json:"kpiHistoryMetadatas"`
}

type KpisCalcUpdatedV1 struct {
	InsID       int64  `json:"insId"`
	UpdatedDate string `json:"updatedDate"`
}

type KpisCalcUpdatedResponse struct {
	Instruments []KpisCalcUpdatedV1 `json:"instruments"`
}

type KpiAllCompV1 struct {
	I int64   `json:"i"` // Instrument ID
	N float64 `json:"n"` // Value
}

type KpisAllCompResponse struct {
	Values []KpiAllCompV1 `json:"values"`
}
