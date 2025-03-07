package fmpcloud

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/zhoub/fmpcloud-go/objects"
)

// Url const for request
const (
	UrlAPIRevenueProductSegmentation = "/v4/revenue-product-segmentation"
	UrlAPIRevenueGeoSegmentation     = "/v4/revenue-geographic-segmentation"
)

// Revenue client
type Revenue struct {
	Client *HTTPClient
}

func (r *Revenue) ProductSegmentation(req objects.RequestRevenueProductSegmentation) ([]objects.RevenueProductSegmentation, error) {
	reqParam := map[string]string{
		"symbol":    req.Symbol,
		"structure": req.Structure,
		"period":    req.Period,
	}
	data, err := r.Client.Get(UrlAPIRevenueProductSegmentation, reqParam)
	if err != nil {
		return nil, err
	}

	var psList []objects.RevenueProductSegmentation
	err = jsoniter.Unmarshal(data.Body(), &psList)
	if err != nil {
		return nil, err
	}

	return psList, nil
}

func (r *Revenue) GeoSegmentation(req objects.RequestRevenueGeoSegmentation) ([]objects.RevenueGeoSegmentation, error) {
	reqParam := map[string]string{
		"symbol":    req.Symbol,
		"structure": req.Structure,
	}
	data, err := r.Client.Get(UrlAPIRevenueGeoSegmentation, reqParam)
	if err != nil {
		return nil, err
	}

	var gsList []objects.RevenueGeoSegmentation
	err = jsoniter.Unmarshal(data.Body(), &gsList)
	if err != nil {
		return nil, err
	}

	return gsList, nil
}
