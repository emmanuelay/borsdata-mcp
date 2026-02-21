package client

import (
	"fmt"
	"net/url"

	"github.com/emmanuelay/borsdata-mcp/models"
)

// GetKpisUpdated returns update timestamps for KPI calculations.
func (c *Client) GetKpisUpdated() (*models.KpisCalcUpdatedResponse, error) {
	var resp models.KpisCalcUpdatedResponse
	if err := c.getJSON("/v1/instruments/kpis/updated", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetKpiHistory returns KPI history for an instrument and KPI.
// reportType: year, r12, or quarter
// priceType: mean, high, low, latest
func (c *Client) GetKpiHistory(insID int64, kpiID int, reportType, priceType string, maxCount int) (*models.KpisHistoryResponse, error) {
	params := url.Values{}
	if maxCount > 0 {
		params.Set("maxCount", fmt.Sprintf("%d", maxCount))
	}
	var resp models.KpisHistoryResponse
	path := fmt.Sprintf("/v1/instruments/%d/kpis/%d/%s/%s/history", insID, kpiID, reportType, priceType)
	if err := c.getJSON(path, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetKpiScreener returns a single KPI value for an instrument.
// calcGroup: 1month, 3month, 1year, 3year, 5year, 7year, 10year, 15year, last, quarter
// calcType: mean, high, low, sum, latest, count
func (c *Client) GetKpiScreener(insID int64, kpiID int, calcGroup, calcType string) (*models.KpiScreenerResponse, error) {
	var resp models.KpiScreenerResponse
	path := fmt.Sprintf("/v1/instruments/%d/kpis/%d/%s/%s", insID, kpiID, calcGroup, calcType)
	if err := c.getJSON(path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetKpiListScreener returns a KPI value for all instruments.
// calcGroup: 1month, 3month, 1year, 3year, 5year, 7year, 10year, 15year, last, quarter
// calcType: mean, high, low, sum, latest, count
func (c *Client) GetKpiListScreener(kpiID int, calcGroup, calcType string) (*models.KpisAllCompResponse, error) {
	var resp models.KpisAllCompResponse
	path := fmt.Sprintf("/v1/instruments/kpis/%d/%s/%s", kpiID, calcGroup, calcType)
	if err := c.getJSON(path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetGlobalKpiListScreener returns a KPI value for all global instruments.
// Requires Pro+ API access.
func (c *Client) GetGlobalKpiListScreener(kpiID int, calcGroup, calcType string) (*models.KpisAllCompResponse, error) {
	var resp models.KpisAllCompResponse
	path := fmt.Sprintf("/v1/instruments/global/kpis/%d/%s/%s", kpiID, calcGroup, calcType)
	if err := c.getJSON(path, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetKpiSummary returns KPI summary for an instrument.
// reportType: year, r12, or quarter
func (c *Client) GetKpiSummary(insID int64, reportType string, maxCount int) (*models.KpisSummaryResponse, error) {
	params := url.Values{}
	if maxCount > 0 {
		params.Set("maxCount", fmt.Sprintf("%d", maxCount))
	}
	var resp models.KpisSummaryResponse
	path := fmt.Sprintf("/v1/instruments/%d/kpis/%s/summary", insID, reportType)
	if err := c.getJSON(path, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
