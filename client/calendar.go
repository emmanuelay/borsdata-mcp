package client

import (
	"net/url"

	"github.com/emmanuelay/borsdata-mcp/models"
)

// GetReportCalendar returns report release dates for the given instruments.
func (c *Client) GetReportCalendar(instList string) (*models.ReportCalendarResponse, error) {
	params := url.Values{}
	params.Set("instList", instList)
	var resp models.ReportCalendarResponse
	if err := c.getJSON("/v1/instruments/report/calendar", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetDividendCalendar returns dividend dates for the given instruments.
func (c *Client) GetDividendCalendar(instList string) (*models.DividendCalendarResponse, error) {
	params := url.Values{}
	params.Set("instList", instList)
	var resp models.DividendCalendarResponse
	if err := c.getJSON("/v1/instruments/dividend/calendar", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
