package models

type CountryV1 struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CountriesResponse struct {
	Countries []CountryV1 `json:"countries"`
}

type MarketV1 struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CountryID int    `json:"countryId"`
	IsIndex   bool   `json:"isIndex"`
	ExchangeName string `json:"exchangeName"`
}

type MarketsResponse struct {
	Markets []MarketV1 `json:"markets"`
}

type SectorV1 struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type SectorsResponse struct {
	Sectors []SectorV1 `json:"sectors"`
}

type BranchV1 struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	SectorID int    `json:"sectorId"`
}

type BranchesResponse struct {
	Branches []BranchV1 `json:"branches"`
}

type ReportMetadataV1 struct {
	ReportID     int    `json:"reportId"`
	NameSv       string `json:"nameSv"`
	NameEn       string `json:"nameEn"`
}

type ReportMetadataResponse struct {
	ReportMetadatas []ReportMetadataV1 `json:"reportMetadatas"`
}

type TranslationMetadataV1 struct {
	TranslationKey string `json:"translationKey"`
	NameSv         string `json:"nameSv"`
	NameEn         string `json:"nameEn"`
}

type TranslationMetadataResponse struct {
	TranslationMetadatas []TranslationMetadataV1 `json:"translationMetadatas"`
}
