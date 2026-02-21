package client

import (
	"net/url"

	"github.com/emmanuelay/borsdata-mcp/models"
)

// GetInsiderHoldings returns insider transactions for the given instruments.
func (c *Client) GetInsiderHoldings(instList string) (*models.InsiderHoldingsResponse, error) {
	params := url.Values{}
	params.Set("instList", instList)
	var resp models.InsiderHoldingsResponse
	if err := c.getJSON("/v1/holdings/insider", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetShortHoldings returns short positions for all instruments.
func (c *Client) GetShortHoldings() (*models.ShortHoldingsResponse, error) {
	var resp models.ShortHoldingsResponse
	if err := c.getJSON("/v1/holdings/shorts", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetBuybackHoldings returns share buybacks for the given instruments.
func (c *Client) GetBuybackHoldings(instList string) (*models.BuybackHoldingsResponse, error) {
	params := url.Values{}
	params.Set("instList", instList)
	var resp models.BuybackHoldingsResponse
	if err := c.getJSON("/v1/holdings/buyback", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
