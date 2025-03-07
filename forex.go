package fmpcloud

import (
	"fmt"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/zhoub/fmpcloud-go/objects"
)

// Url const for request
const (
	UrlAPIForexListAndQuotes = "/v3/fx"
	UrlAPIForexSymbols       = "/v3/symbol/available-forex-currency-pairs"
	UrlAPIForexQuotes        = "/v3/quotes/forex"
	UrlAPIForexCandles       = "/v3/historical-chart/%s/%s"
	UrlAPIForexDaily         = "/v3/historical-price-full/%s"
)

// Forex client
type Forex struct {
	Client *HTTPClient
}

// AvalibleSymbols - available symbol list
func (f *Forex) AvalibleSymbols() (sList []objects.ForexSymbol, err error) {
	data, err := f.Client.Get(UrlAPIForexSymbols, nil)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &sList)
	if err != nil {
		return nil, err
	}

	return sList, nil
}

// Quotes - all real-time prices
func (f *Forex) Quotes() (qList []objects.ForexQuote, err error) {
	data, err := f.Client.Get(UrlAPIForexQuotes, nil)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &qList)
	if err != nil {
		return nil, err
	}

	return qList, nil
}

// ListSymbolsAndQuotes - Forex List And Price (Get last bid/ask data)
func (f *Forex) ListSymbolsAndQuotes() (bList []objects.ForexBindAsk, err error) {
	data, err := f.Client.Get(UrlAPIForexListAndQuotes, nil)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &bList)
	if err != nil {
		return nil, err
	}

	return bList, nil
}

// Candles - historical candles
func (f *Forex) Candles(req objects.RequestForexCandleList) (cList []objects.ForexCandle, err error) {
	reqParam := make(map[string]string)
	if req.From != nil {
		reqParam["from"] = req.From.Format("2006-01-02")
	}

	if req.To != nil {
		reqParam["to"] = req.To.Format("2006-01-02")
	}

	data, err := f.Client.Get(fmt.Sprintf(UrlAPIForexCandles, req.Period, req.Symbol), reqParam)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &cList)
	if err != nil {
		return nil, err
	}

	return cList, nil
}

// DailyLine - faily line
func (f *Forex) DailyLine(symbol string, serieType objects.ForexSerieType) (cList *objects.ForexDailyLineList, err error) {
	data, err := f.Client.Get(fmt.Sprintf(UrlAPIForexDaily, symbol), map[string]string{"serietype": string(serieType)})
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &cList)
	if err != nil {
		return nil, err
	}

	return cList, nil
}

// DailyChangeAndVolume - daily candle change and volume
func (f *Forex) DailyChangeAndVolume(symbol string) (cList *objects.ForexDailyCandleList, err error) {
	data, err := f.Client.Get(fmt.Sprintf(UrlAPIForexDaily, symbol), nil)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &cList)
	if err != nil {
		return nil, err
	}

	return cList, nil
}

// DailySpecificPeriod - daily candle list by specific period
func (f *Forex) DailySpecificPeriod(symbol string, from time.Time, to time.Time) (cList *objects.ForexDailyCandleList, err error) {
	data, err := f.Client.Get(
		fmt.Sprintf(UrlAPIForexDaily, symbol),
		map[string]string{
			"from": from.Format("2006-01-02"),
			"to":   to.Format("2006-01-02"),
		})
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &cList)
	if err != nil {
		return nil, err
	}

	return cList, nil
}

// DailyLastNDays - daily candle list last N days
func (f *Forex) DailyLastNDays(symbol string, days int) (cList *objects.ForexDailyCandleList, err error) {
	data, err := f.Client.Get(fmt.Sprintf(UrlAPIForexDaily, symbol), map[string]string{"timeseries": fmt.Sprint(days)})
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &cList)
	if err != nil {
		return nil, err
	}

	return cList, nil
}
