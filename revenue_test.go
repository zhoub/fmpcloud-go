package fmpcloud

import (
	"testing"

	"github.com/zhoub/fmpcloud-go/objects"
)

func TestRevenueProductSegementation(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, symbol := range testCaseSymbolList {
		_, err = APIClient.Revenue.ProductSegmentation(objects.RequestRevenueProductSegmentation{
			Symbol:    symbol,
			Structure: objects.RevenuStructureFlat,
			Period:    objects.RevenuePeriodAnnual,
		})
		if err != nil {
			t.Fatal(err.Error())
		}
	}
}

func TestRevenueGeoSegementation(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, symbol := range testCaseSymbolList {
		_, err = APIClient.Revenue.GeoSegmentation(objects.RequestRevenueGeoSegmentation{
			Symbol:    symbol,
			Structure: objects.RevenuStructureFlat,
		})
		if err != nil {
			t.Fatal(err.Error())
		}
	}
}
