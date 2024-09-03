package fmpcloud

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/zhoub/fmpcloud-go/objects"
)

const (
	urlAPIExchangesList      = "/v3/exchanges-list"
	urlAPIIsTheMarketOpenAll = "/v3/is-the-market-open-all"
)

type Exchange struct {
	Client *HTTPClient
}

func (f *Exchange) AvailableExchanges() (eList []string, err error) {
	data, err := f.Client.Get(urlAPIExchangesList, nil)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &eList)
	if err != nil {
		return nil, err
	}

	return eList, nil
}

func (f *Exchange) IsTheMarketOpenAll() (emList []objects.ExchangeMarket, err error) {
	data, err := f.Client.Get(urlAPIIsTheMarketOpenAll, nil)
	if err != nil {
		return nil, err
	}

	err = jsoniter.Unmarshal(data.Body(), &emList)
	if err != nil {
		return nil, err
	}

	return emList, nil
}
