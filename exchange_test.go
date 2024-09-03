package fmpcloud

import "testing"

func TestAvailableExchanges(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Exchange.AvailableExchanges()
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestIsTheMarketOpenAll(t *testing.T) {
	APIClient, err := NewAPIClient(testCaseAPIConfig)
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = APIClient.Exchange.IsTheMarketOpenAll()
	if err != nil {
		t.Fatal(err.Error())
	}
}
