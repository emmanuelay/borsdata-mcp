# Börsdata MCP

En MCP-server (Model Context Protocol) som ger tillgång till Börsdata API. Byggd med [go-sdk](https://github.com/modelcontextprotocol/go-sdk).

> [!IMPORTANT]
> Börsdata API kräver en API-nyckel. Skaffa en via [Börsdata](https://borsdata.se). API-nyckeln anges via miljövariabeln `BORSDATA_API_KEY`.
>
> Vissa verktyg stödjer `scope: global` för icke-nordiska instrument, vilket kräver Pro+-åtkomst.

## Exempelfrågor

- "Screena fram alla nordiska bolag med P/E under 10 och positiv vinsttillväxt senaste 5 åren"
- "Jämför rörelsemarginalen för Hexagon, Atlas Copco och Alfa Laval kvartalsvis de senaste 3 åren"
- "Vilka bolag har störst blankningspositioner just nu och hur har deras kurser utvecklats senaste månaden?"
- "Visa insiderköp över 1 MSEK i Investor och Kinnevik det senaste halvåret"
- "Analysera Evolution Gamings R12-rapporter och identifiera trender i omsättningstillväxt och marginalutveckling"
- "Vilka bolag i tekniksektorn har kommande rapportdatum inom de närmaste två veckorna?"
- "Sammanfatta nyckeltalen för Sandvik — P/E, EV/EBITDA, direktavkastning och skuldsättningsgrad — och jämför med branschsnittet"
- "Hämta utdelningshistoriken för Handelsbanken och beräkna den genomsnittliga utdelningstillväxten per år"

## Verktyg

| Verktyg | Beskrivning |
|---------|-------------|
| `list_instruments` | Lista alla tillgängliga instrument (bolag) med ID, namn, ticker, sektor, marknad och land |
| `search_instruments` | Sök instrument via namn eller ticker |
| `get_instruments_updated` | Visa uppdateringstidpunkter för alla instrument |
| `get_stock_splits` | Hämta alla aktiesplittar med datum, typ och förhållande |
| `get_stock_prices` | Historiska aktiekurser för ett specifikt instrument (öppning, hög, låg, stängning, volym) |
| `get_latest_stock_prices` | Senaste aktiekurser för alla instrument |
| `get_stock_prices_by_date` | Aktiekurser för alla instrument på ett specifikt datum |
| `get_reports` | Finansiella rapporter: årsredovisning, R12, kvartalsrapport |
| `get_kpis_updated` | Uppdateringstidpunkter för nyckeltal |
| `get_kpi_history` | Nyckeltalshistorik för ett instrument |
| `get_kpi_screener` | Enskilt nyckeltalsvärde för ett specifikt instrument |
| `get_kpi_list_screener` | Nyckeltalsvärde för alla instrument (screening/filtrering) |
| `get_kpi_summary` | Nyckeltalssammanfattning för ett instrument |
| `get_metadata` | Referensdata: länder, marknader, sektorer, branscher, nyckeltal, rapporter, översättningar |
| `get_insider_holdings` | Insidertransaktioner med ägare, position, antal aktier, pris och belopp |
| `get_short_holdings` | Blankningspositioner för alla instrument |
| `get_buyback_holdings` | Aktieåterköp för angivna instrument |
| `get_report_calendar` | Rapportdatum för angivna instrument |
| `get_dividend_calendar` | Utdelningsdatum med belopp, valuta och frekvens |

## Installera

### Installera som Desktop Extension (.mcpb) för t.ex. Claude Desktop

Det enklaste sättet att installera Börsdata MCP'n är genom att ladda ner `.mcpb`-filen för din plattform från [senaste releasen](https://github.com/emmanuelay/borsdata-mcp/releases/latest) och öppna den i Claude Desktop. Du kommer att bli ombedd att ange din Börsdata API-nyckel vid installation.


### Installera via Homebrew

```sh
brew tap emmanuelay/tap
brew install borsdata-mcp
```

### Från källkod

```sh
git clone https://github.com/emmanuelay/borsdata-mcp.git
cd borsdata-mcp
make build
```

## Aktivera i Anthropics Claude Desktop / Claude Code

### Claude Desktop

Lägg till servern i din `claude_desktop_config.json`:

```json
{
  "mcpServers": {
    "borsdata": {
      "command": "borsdata-mcp",
      "env": {
        "BORSDATA_API_KEY": "din-api-nyckel"
      }
    }
  }
}
```

Konfigurationsfilen finns på:

- macOS: `~/Library/Application Support/Claude/claude_desktop_config.json`
- Windows: `%APPDATA%\Claude\claude_desktop_config.json`

Starta om Claude Desktop efter att du sparat konfigurationen.

### Claude Code

```sh
claude mcp add borsdata -e BORSDATA_API_KEY=din-api-nyckel -- borsdata-mcp
```

## Utveckling

```sh
make help       # Visa alla tillgängliga targets
make deps       # Ladda ner beroenden
make build      # Bygg binären
make test       # Kör tester
make lint       # Statisk analys med staticcheck
make scan       # Säkerhetsskanning med govulncheck
make install    # Installera till $GOPATH/bin
make tags       # Lista alla git-taggar
make mcpb       # Bygg .mcpb-paket från goreleaser dist/
```

## Licens

MIT
