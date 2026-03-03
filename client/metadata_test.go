package client

import (
	"testing"

	"github.com/emmanuelay/borsdata-mcp/models"
)

func TestGetCountries(t *testing.T) {
	expected := models.CountriesResponse{
		Countries: []models.CountryV1{
			{ID: 1, Name: "Sweden"},
		},
	}
	log, c := newTestServer(t, map[string]any{
		"/v1/countries": expected,
	})

	result, err := c.GetCountries()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Countries) != 1 {
		t.Fatalf("expected 1 country, got %d", len(result.Countries))
	}
	if result.Countries[0].Name != "Sweden" {
		t.Errorf("expected Sweden, got %s", result.Countries[0].Name)
	}
	log.assertParam(t, "/v1/countries", "authKey", "test-key")
}

func TestGetMarkets(t *testing.T) {
	expected := models.MarketsResponse{
		Markets: []models.MarketV1{
			{ID: 1, Name: "Large Cap", CountryID: 1},
		},
	}
	log, c := newTestServer(t, map[string]any{
		"/v1/markets": expected,
	})

	result, err := c.GetMarkets()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Markets) != 1 {
		t.Fatalf("expected 1 market, got %d", len(result.Markets))
	}
	log.assertParam(t, "/v1/markets", "authKey", "test-key")
}

func TestGetSectors(t *testing.T) {
	expected := models.SectorsResponse{
		Sectors: []models.SectorV1{
			{ID: 1, Name: "Technology"},
		},
	}
	log, c := newTestServer(t, map[string]any{
		"/v1/sectors": expected,
	})

	result, err := c.GetSectors()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Sectors) != 1 {
		t.Fatalf("expected 1 sector, got %d", len(result.Sectors))
	}
	log.assertParam(t, "/v1/sectors", "authKey", "test-key")
}

func TestGetBranches(t *testing.T) {
	expected := models.BranchesResponse{
		Branches: []models.BranchV1{
			{ID: 1, Name: "Software", SectorID: 1},
		},
	}
	log, c := newTestServer(t, map[string]any{
		"/v1/branches": expected,
	})

	result, err := c.GetBranches()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.Branches) != 1 {
		t.Fatalf("expected 1 branch, got %d", len(result.Branches))
	}
	log.assertParam(t, "/v1/branches", "authKey", "test-key")
}

func TestGetKpiMetadata(t *testing.T) {
	expected := models.KpisMetadataResponse{
		KpiHistoryMetadatas: []models.KpiMetadataV1{
			{KpiID: 1, NameSv: "P/E-tal", NameEn: "P/E ratio"},
		},
	}
	log, c := newTestServer(t, map[string]any{
		"/v1/instruments/kpis/metadata": expected,
	})

	result, err := c.GetKpiMetadata()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result.KpiHistoryMetadatas) != 1 {
		t.Fatalf("expected 1 KPI metadata, got %d", len(result.KpiHistoryMetadatas))
	}
	log.assertParam(t, "/v1/instruments/kpis/metadata", "authKey", "test-key")
}
