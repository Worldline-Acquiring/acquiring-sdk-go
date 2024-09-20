package communicator

import (
	"net/url"
	"testing"

	"github.com/Worldline-Acquiring/acquiring-sdk-go/communicator/communication"
)

var baseURL = url.URL{
	Scheme: "https",
	Host:   "api.preprod.acquiring.worldline-solutions.com",
}

func TestToURIWithoutRequestParams(t *testing.T) {
	communicator := &Communicator{apiEndpoint: &baseURL}

	expectedURL, err := url.Parse("https://api.preprod.acquiring.worldline-solutions.com/services/v1/100812/520000214/dcc-rates")
	if err != nil {
		t.Fatalf("TestToURIWithoutRequestParams: %v", err)
	}

	uri1, err := communicator.toAbsoluteURI("services/v1/100812/520000214/dcc-rates", nil)
	if err != nil {
		t.Fatalf("TestToURIWithoutRequestParams: %v", err)
	}
	if uri1 != *expectedURL {
		t.Fatalf("TestToURIWithoutRequestParams: url mismatch '%v' '%v'", uri1, *expectedURL)
	}

	uri2, err := communicator.toAbsoluteURI("/services/v1/100812/520000214/dcc-rates", nil)
	if err != nil {
		t.Fatalf("TestToURIWithoutRequestParams: %v", err)
	}
	if uri2 != *expectedURL {
		t.Fatalf("TestToURIWithoutRequestParams: url mismatch '%v' '%v'", uri2, *expectedURL)
	}
}

func TestToURIWithRequestParams(t *testing.T) {
	amountParam, _ := communication.NewRequestParam("amount", "123")
	sourceParam, _ := communication.NewRequestParam("source", "USD")
	targetParam, _ := communication.NewRequestParam("target", "EUR")
	dummyParam, _ := communication.NewRequestParam("dummy", "é&%=")
	params := append(communication.RequestParams{}, *amountParam, *sourceParam, *targetParam, *dummyParam)

	communicator := &Communicator{apiEndpoint: &baseURL}

	expectedURL, err := url.Parse("https://api.preprod.acquiring.worldline-solutions.com/services/v1/100812/520000214/dcc-rates?amount=123&source=USD&target=EUR&dummy=%C3%A9%26%25%3D")
	if err != nil {
		t.Fatalf("TestToURIWithoutRequestParams: %v", err)
	}

	uri1, err := communicator.toAbsoluteURI("services/v1/100812/520000214/dcc-rates", params)
	if err != nil {
		t.Fatalf("TestToURIWithoutRequestParams: %v", err)
	}
	if uri1 != *expectedURL {
		t.Fatalf("TestToURIWithoutRequestParams: url mismatch '%v' '%v'", uri1, *expectedURL)
	}

	uri2, err := communicator.toAbsoluteURI("/services/v1/100812/520000214/dcc-rates", params)
	if err != nil {
		t.Fatalf("TestToURIWithoutRequestParams: %v", err)
	}
	if uri2 != *expectedURL {
		t.Fatalf("TestToURIWithoutRequestParams: url mismatch '%v' '%v'", uri2, *expectedURL)
	}
}
