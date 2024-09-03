package fmpcloud

import jsoniter "github.com/json-iterator/go"

const (
	urlAPIExchangesList = "/v3/exchanges-list"
)

type Exchange struct {
	Client *HTTPClient
}

func (f *Forex) AvailableExchanges() (eList []string, err error) {
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
