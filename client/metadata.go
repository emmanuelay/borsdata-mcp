package client

import (
	"github.com/emmanuelay/borsdata-mcp/models"
)

// GetCountries returns all countries.
func (c *Client) GetCountries() (*models.CountriesResponse, error) {
	var resp models.CountriesResponse
	if err := c.getJSON("/v1/countries", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetMarkets returns all markets.
func (c *Client) GetMarkets() (*models.MarketsResponse, error) {
	var resp models.MarketsResponse
	if err := c.getJSON("/v1/markets", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSectors returns all sectors.
func (c *Client) GetSectors() (*models.SectorsResponse, error) {
	var resp models.SectorsResponse
	if err := c.getJSON("/v1/sectors", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetBranches returns all branches.
func (c *Client) GetBranches() (*models.BranchesResponse, error) {
	var resp models.BranchesResponse
	if err := c.getJSON("/v1/branches", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetKpiMetadata returns KPI metadata (names, abbreviations, etc).
func (c *Client) GetKpiMetadata() (*models.KpisMetadataResponse, error) {
	var resp models.KpisMetadataResponse
	if err := c.getJSON("/v1/instruments/kpis/metadata", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetReportMetadata returns report field metadata.
func (c *Client) GetReportMetadata() (*models.ReportMetadataResponse, error) {
	var resp models.ReportMetadataResponse
	if err := c.getJSON("/v1/instruments/reports/metadata", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTranslationMetadata returns translation metadata.
func (c *Client) GetTranslationMetadata() (*models.TranslationMetadataResponse, error) {
	var resp models.TranslationMetadataResponse
	if err := c.getJSON("/v1/translationmetadata", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
