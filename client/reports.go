package client

import (
	"fmt"
	"net/url"

	"github.com/emmanuelay/borsdata-mcp/models"
)

// GetYearlyReports returns yearly reports for an instrument.
func (c *Client) GetYearlyReports(insID int64, maxCount int) (*models.ReportsResponse, error) {
	params := url.Values{}
	if maxCount > 0 {
		params.Set("maxCount", fmt.Sprintf("%d", maxCount))
	}
	var resp models.ReportsResponse
	path := fmt.Sprintf("/v1/instruments/%d/reports/year", insID)
	if err := c.getJSON(path, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetR12Reports returns rolling 12-month reports for an instrument.
func (c *Client) GetR12Reports(insID int64, maxCount int) (*models.ReportsResponse, error) {
	params := url.Values{}
	if maxCount > 0 {
		params.Set("maxCount", fmt.Sprintf("%d", maxCount))
	}
	var resp models.ReportsResponse
	path := fmt.Sprintf("/v1/instruments/%d/reports/r12", insID)
	if err := c.getJSON(path, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetQuarterlyReports returns quarterly reports for an instrument.
func (c *Client) GetQuarterlyReports(insID int64, maxCount int) (*models.ReportsResponse, error) {
	params := url.Values{}
	if maxCount > 0 {
		params.Set("maxCount", fmt.Sprintf("%d", maxCount))
	}
	var resp models.ReportsResponse
	path := fmt.Sprintf("/v1/instruments/%d/reports/quarter", insID)
	if err := c.getJSON(path, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetAllReports returns all report types (year, r12, quarter) for an instrument.
func (c *Client) GetAllReports(insID int64, maxYearCount, maxR12QCount int) (*models.ReportsCompoundResponse, error) {
	params := url.Values{}
	if maxYearCount > 0 {
		params.Set("maxYearCount", fmt.Sprintf("%d", maxYearCount))
	}
	if maxR12QCount > 0 {
		params.Set("maxR12QCount", fmt.Sprintf("%d", maxR12QCount))
	}
	var resp models.ReportsCompoundResponse
	path := fmt.Sprintf("/v1/instruments/%d/reports", insID)
	if err := c.getJSON(path, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
