package client

import (
	"github.com/emmanuelay/borsdata-mcp/models"
)

// GetInstruments returns all instruments (companies).
func (c *Client) GetInstruments() (*models.InstrumentsResponse, error) {
	var resp models.InstrumentsResponse
	if err := c.getJSON("/v1/instruments", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetInstrumentsUpdated returns update timestamps for all instruments.
func (c *Client) GetInstrumentsUpdated() (*models.InstrumentsUpdatedResponse, error) {
	var resp models.InstrumentsUpdatedResponse
	if err := c.getJSON("/v1/instruments/updated", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetGlobalInstruments returns all global (non-Nordic) instruments.
// Requires Pro+ API access.
func (c *Client) GetGlobalInstruments() (*models.InstrumentsResponse, error) {
	var resp models.InstrumentsResponse
	if err := c.getJSON("/v1/instruments/global", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetStockSplits returns all stock splits.
func (c *Client) GetStockSplits() (*models.StockSplitsResponse, error) {
	var resp models.StockSplitsResponse
	if err := c.getJSON("/v1/instruments/stocksplits", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
