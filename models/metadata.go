package models

// CountryV1 represents a country with its ID and name.
type CountryV1 struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// CountriesResponse is the top-level response from the countries metadata endpoint.
type CountriesResponse struct {
	Countries []CountryV1 `json:"countries"`
}

// MarketV1 represents a stock market/exchange with its country, name, and whether it is an index.
type MarketV1 struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	CountryID    int    `json:"countryId"`
	IsIndex      bool   `json:"isIndex"`
	ExchangeName string `json:"exchangeName"`
}

// MarketsResponse is the top-level response from the markets metadata endpoint.
type MarketsResponse struct {
	Markets []MarketV1 `json:"markets"`
}

// SectorV1 represents an industry sector with its ID and name.
type SectorV1 struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// SectorsResponse is the top-level response from the sectors metadata endpoint.
type SectorsResponse struct {
	Sectors []SectorV1 `json:"sectors"`
}

// BranchV1 represents an industry branch (sub-sector) with its ID, name, and parent sector ID.
type BranchV1 struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	SectorID int    `json:"sectorId"`
}

// BranchesResponse is the top-level response from the branches metadata endpoint.
type BranchesResponse struct {
	Branches []BranchV1 `json:"branches"`
}

// ReportMetadataV1 represents metadata for a financial report field, including its Swedish and English names.
type ReportMetadataV1 struct {
	ReportID int    `json:"reportId"`
	NameSv   string `json:"nameSv"`
	NameEn   string `json:"nameEn"`
}

// ReportMetadataResponse is the top-level response from the report metadata endpoint.
type ReportMetadataResponse struct {
	ReportMetadatas []ReportMetadataV1 `json:"reportMetadatas"`
}

// TranslationMetadataV1 represents a translation entry with a key and its Swedish and English values.
type TranslationMetadataV1 struct {
	TranslationKey string `json:"translationKey"`
	NameSv         string `json:"nameSv"`
	NameEn         string `json:"nameEn"`
}

// TranslationMetadataResponse is the top-level response from the translation metadata endpoint.
type TranslationMetadataResponse struct {
	TranslationMetadatas []TranslationMetadataV1 `json:"translationMetadatas"`
}
