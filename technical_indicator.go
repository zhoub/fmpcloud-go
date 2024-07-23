package fmpcloud

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
	"github.com/zhoub/fmpcloud-go/objects"
)

// Url const for request
const (
	urlAPITechnicalIndicatorSymbol = "/v3/technical_indicator/%s/%s"
)

// TechnicalIndicator client
type TechnicalIndicator struct {
	Client *HTTPClient
}

// Indicators - Daily Indicators. Types: SMA - EMA - WMA - DEMA - TEMA - williams - RSI - ADX - standardDeviation
func (t *TechnicalIndicator) Indicators(req objects.RequestIndicators) (iList []objects.ResponseIndicators, err error) {
	data, err := t.Client.Get(
		fmt.Sprintf(urlAPITechnicalIndicatorSymbol, req.Resolution.String(), req.Symbol),
		map[string]string{
			"type":   req.Indicator.String(),
			"period": fmt.Sprint(req.Timeperiod),
		})
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &iList)
	if err != nil {
		return nil, err
	}

	return iList, nil
}
