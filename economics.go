package fmpcloud

import (
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/zhoub/fmpcloud-go/objects"
)

// Url const for request
const (
	UrlAPIEconomicsMarketRiskPremium = "/v4/market_risk_premium"
	UrlAPIEconomicsTreasury          = "/v4/treasury"
	UrlAPIEconomicsIndicator         = "/v4/economic"
)

// Economics client
type Economics struct {
	Client *HTTPClient
}

// MarketRiskPremium - Market Risk Premium for each countr
func (e *Economics) MarketRiskPremium() (mList []objects.EconomicsMarketRisk, err error) {
	data, err := e.Client.Get(UrlAPIEconomicsMarketRiskPremium, nil)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &mList)
	if err != nil {
		return nil, err
	}

	return mList, nil
}

// TreasuryRates - Historical treasury rates (between the "from" and "to" parameters the maximum time interval can be 3 months)
func (e *Economics) TreasuryRates(from time.Time, to time.Time) (tList []objects.EconomicsTreasuryRates, err error) {
	data, err := e.Client.Get(
		UrlAPIEconomicsTreasury,
		map[string]string{
			"from": from.Format("2006-01-02"),
			"to":   to.Format("2006-01-02"),
		})
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &tList)
	if err != nil {
		return nil, err
	}

	return tList, nil
}

// Indicator - https://site.financialmodelingprep.com/developer/docs/economic-indicator-api
func (e *Economics) Indicator(indicator string, from *time.Time, to *time.Time) (iList []objects.EconomicsIndicator, err error) {
	req := map[string]string{"indicator": indicator}
	if from != nil {
		req["from"] = from.Format("2006-01-02")
	}

	if to != nil {
		req["to"] = to.Format("2006-01-02")
	}

	data, err := e.Client.Get(UrlAPIEconomicsIndicator, req)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &iList)
	if err != nil {
		return nil, err
	}

	return iList, nil
}
