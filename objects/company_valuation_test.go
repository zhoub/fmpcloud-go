package objects

import (
	"encoding/json"
	"testing"
)

func TestUnmarshalStockCompanyProfile(t *testing.T) {
	s := `
[
  {
    "symbol": "EQV",
    "price": 9.94,
    "beta": 0,
    "volAvg": 224545.2,
    "mktCap": 354485250,
    "lastDiv": 0,
    "range": "9.93-9.94",
    "changes": -0.01,
    "companyName": "EQV Ventures Acquisition Corp.",
    "currency": null,
    "cik": null,
    "isin": "KYG3106N1253",
    "cusip": null,
    "exchange": "New York Stock Exchange",
    "exchangeShortName": "NYSE",
    "industry": null,
    "website": null,
    "description": null,
    "ceo": null,
    "sector": null,
    "country": null,
    "fullTimeEmployees": null,
    "phone": null,
    "address": null,
    "city": null,
    "state": null,
    "zip": null,
    "dcfDiff": null,
    "dcf": 0,
    "image": "https://financialmodelingprep.com/image-stock/EQV.png",
    "ipoDate": "",
    "defaultImage": true,
    "isEtf": false,
    "isActivelyTrading": true,
    "isAdr": false,
    "isFund": false
  }
]
	`
	var o []StockCompanyProfile
	if err := json.Unmarshal([]byte(s), &o); err != nil {
		t.Fatal(err.Error())
	}
}
