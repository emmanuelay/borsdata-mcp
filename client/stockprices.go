package client

import (
	"fmt"
	"net/url"

	"github.com/emmanuelay/borsdata-mcp/models"
)

// GetStockPrices returns stock prices for an instrument.
// The API does not support server-side limiting, so results are truncated client-side.
func (c *Client) GetStockPrices(insID int64, maxCount int) (*models.StockPricesResponse, error) {
	var resp models.StockPricesResponse
	path := fmt.Sprintf("/v1/instruments/%d/stockprices", insID)
	if err := c.getJSON(path, nil, &resp); err != nil {
		return nil, err
	}
	if maxCount > 0 && len(resp.StockPricesList) > maxCount {
		resp.StockPricesList = resp.StockPricesList[len(resp.StockPricesList)-maxCount:]
	}
	return &resp, nil
}

// GetStockPricesDateRange returns stock prices for an instrument within a date range.
func (c *Client) GetStockPricesDateRange(insID int64, from, to string) (*models.StockPricesResponse, error) {
	params := url.Values{}
	params.Set("from", from)
	params.Set("to", to)
	var resp models.StockPricesResponse
	path := fmt.Sprintf("/v1/instruments/%d/stockprices", insID)
	if err := c.getJSON(path, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetLatestStockPrices returns the latest stock prices for all instruments.
func (c *Client) GetLatestStockPrices() (*models.StockPricesLastResponse, error) {
	var resp models.StockPricesLastResponse
	if err := c.getJSON("/v1/instruments/stockprices/last", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetGlobalLatestStockPrices returns the latest stock prices for all global instruments.
// Requires Pro+ API access.
func (c *Client) GetGlobalLatestStockPrices() (*models.StockPricesLastResponse, error) {
	var resp models.StockPricesLastResponse
	if err := c.getJSON("/v1/instruments/stockprices/global/last", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetGlobalStockPricesByDate returns stock prices for all global instruments on a specific date.
// Requires Pro+ API access.
func (c *Client) GetGlobalStockPricesByDate(date string) (*models.StockPricesDateResponse, error) {
	params := url.Values{}
	params.Set("date", date)
	var resp models.StockPricesDateResponse
	if err := c.getJSON("/v1/instruments/stockprices/global/date", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetStockPricesByDate returns stock prices for all instruments on a specific date.
func (c *Client) GetStockPricesByDate(date string) (*models.StockPricesDateResponse, error) {
	params := url.Values{}
	params.Set("date", date)
	var resp models.StockPricesDateResponse
	if err := c.getJSON("/v1/instruments/stockprices/date", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
