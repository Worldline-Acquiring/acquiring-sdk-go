package communicator

import (
	"testing"

	"github.com/Worldline-Acquiring/acquiring-sdk-go/communicator/communication"
)

type Parameter struct {
	name    string
	allowed bool
}

func TestWithAdditionalRequestHeader(t *testing.T) {
	params := make(chan Parameter)
	go func(params chan Parameter) {
		for _, prohibitedHeaders := range prohibitedHeaders {
			params <- Parameter{prohibitedHeaders, false}
		}
		params <- Parameter{"Dummy", true}
		params <- Parameter{"Accept", true}
		params <- Parameter{"If-None-Match", true}
		params <- Parameter{"If-Modified-Since", true}
		close(params)
	}(params)
	for param := range params {
		additionalRequestHeader, _ := communication.NewHeader(param.name, "value")
		builder, _ := NewMetadataProviderBuilder("Worldline")
		builder.AdditionalRequestHeaders = []communication.Header{*additionalRequestHeader}
		metadataProvider, err := builder.Build()
		if err != nil && param.allowed {
			t.Error("Cannot build MetadataProvider")
			continue
		}
		if err == nil && !param.allowed {
			t.Error("Can build invalid MetadataProvider")
		} else {
			continue
		}
		headers := metadataProvider.metadataHeaders
		if len(headers) != 2 {
			t.Error("Wrong amount of headers stored")
		}
		header := headers[0]
		if "X-WL-ServerMetaInfo" != header.Name() {
			t.Error("X-WL-ServerMetaInfo should always be set")
		}
		header = headers[1]
		if param.name != header.Name() {
			t.Error("No other headers should be set")
		}
	}
}
