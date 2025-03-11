package fmpcloud

import (
	"fmt"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/zhoub/fmpcloud-go/objects"
)

// Url const for request
const (
	UrlAPITechnicalIndicatorSymbol = "/v3/technical_indicator/%s/%s"
)

// TechnicalIndicator client
type TechnicalIndicator struct {
	Client *HTTPClient
}

// Indicators - Daily Indicators. Types: SMA - EMA - WMA - DEMA - TEMA - williams - RSI - ADX - standardDeviation
func (t *TechnicalIndicator) Indicators(req objects.RequestIndicators) (iList []objects.ResponseIndicators, err error) {
	params := map[string]string{
		"type":   req.Indicator.String(),
		"period": fmt.Sprint(req.Timeperiod),
	}
	if req.From != nil {
		params["from"] = req.From.Format(time.DateOnly)
	}
	if req.To != nil {
		params["to"] = req.To.Format(time.DateOnly)
	}

	data, err := t.Client.Get(
		fmt.Sprintf(UrlAPITechnicalIndicatorSymbol, req.Resolution.String(), req.Symbol),
		params)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &iList)
	if err != nil {
		return nil, err
	}

	return iList, nil
}
