package fmpcloud

import (
	"testing"

	"github.com/zhoub/fmpcloud-go/objects"
)

func TestInsiderTradingList(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.InsiderTrading.List(objects.RequestInsiderTrading{
		Symbol: "AAPL",
		Limit:  25,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestInsiderTradingRSSFeed(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.InsiderTrading.RSSFeed(25)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestInsiderTradingMapperCikCompany(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.InsiderTrading.MapperCikCompany("AAPL")
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestInsiderTradingMapperCikName(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.InsiderTrading.MapperCikName(nil)
	if err != nil {
		t.Fatal(err.Error())
	}
}
