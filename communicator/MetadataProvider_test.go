package communicator

import (
	"encoding/base64"
	"encoding/json"
	"testing"

	"github.com/Worldline-Acquiring/acquiring-sdk-go/communicator/communication"
	"github.com/Worldline-Acquiring/acquiring-sdk-go/domain"
)

func TestGetServerMetadataHeadersNoAdditionalHeaders(t *testing.T) {
	metadataProvider, err := NewMetadataProvider("Worldline")
	if err != nil {
		t.Fatalf("TestGetServerMetadataHeadersNoAdditionalHeaders : %v", err)
	}

	requestHeaders := metadataProvider.MetadataHeaders()
	if requestHeaders == nil {
		t.Fatal("TestGetServerMetadataHeadersNoAdditionalHeaders : nil headers")
	}
	if len(requestHeaders) != 1 {
		t.Fatal("TestGetServerMetadataHeadersNoAdditionalHeaders : len != 1")
	}

	AssertServerMetaInfo("TestGetServerMetadataHeadersNoAdditionalHeaders",
		t, metadataProvider, "Worldline", requestHeaders[0])
}

func TestServerMetadataHeadersFull(t *testing.T) {
	shoppingCartExtension, err := domain.NewShoppingCartExtensionWithExtensionID("Worldline.creator", "Extension", "1.0", "ExtensionID")
	if err != nil {
		t.Fatalf("TestServerMetadataHeadersFull : %v", err)
	}

	integrator := "Worldline"
	metadataProviderBuilder, _ := NewMetadataProviderBuilder(integrator)
	metadataProviderBuilder.ShoppingCartExtension = shoppingCartExtension
	metadataProvider, err := metadataProviderBuilder.Build()
	if err != nil {
		t.Fatalf("TestServerMetadataHeadersFull : %v", err)
	}

	requestHeaders := metadataProvider.MetadataHeaders()
	if len(requestHeaders) != 1 {
		t.Fatal("TestServerMetadataHeadersFull : len != 1")
	}

	AssertServerMetaInfo("TestServerMetadataHeadersFull",
		t, metadataProvider, integrator, requestHeaders[0])
	AssertShoppingCard("TestServerMetadataHeadersFull",
		t, metadataProvider, shoppingCartExtension)
}

func TestServerMetadataHeadersFullNewShoppingCartExtensionWithExtensionID(t *testing.T) {
	shoppingCartExtension, err := domain.NewShoppingCartExtension("Worldline.creator", "Extension", "1.0")
	if err != nil {
		t.Fatalf("TestServerMetadataHeadersFullNewShoppingCartExtensionWithExtensionID : %v", err)
	}

	integrator := "Worldline"
	metadataProviderBuilder, _ := NewMetadataProviderBuilder(integrator)
	metadataProviderBuilder.ShoppingCartExtension = shoppingCartExtension
	metadataProvider, err := metadataProviderBuilder.Build()
	if err != nil {
		t.Fatalf("TestServerMetadataHeadersFull : %v", err)
	}

	requestHeaders := metadataProvider.MetadataHeaders()
	if len(requestHeaders) != 1 {
		t.Fatal("TestServerMetadataHeadersFull : len != 1")
	}

	AssertServerMetaInfo("TestServerMetadataHeadersFull",
		t, metadataProvider, integrator, requestHeaders[0])
	AssertShoppingCard("TestServerMetadataHeadersFull",
		t, metadataProvider, shoppingCartExtension)
}

func TestServerMetadataHeadersAdditionalHeaders(t *testing.T) {
	header1, _ := communication.NewHeader("Header1", "&=$%")
	header2, _ := communication.NewHeader("Header2", "blah blah")
	header3, _ := communication.NewHeader("Header3", "foo")

	additionalHeaders := []communication.Header{*header1, *header2, *header3}
	metadataProviderBuilder, _ := NewMetadataProviderBuilder("Worldline")
	metadataProviderBuilder.AdditionalRequestHeaders = additionalHeaders
	metadataProvider, err := metadataProviderBuilder.Build()
	if err != nil {
		t.Fatalf("TestServerMetadataHeadersAdditionalHeaders : %v", err)
	}

	requestHeaders := metadataProvider.MetadataHeaders()
	if len(requestHeaders) != 4 {
		t.Fatal("TestServerMetadataHeadersAdditionalHeaders : len != 4")
	}

	for i := 1; i < len(requestHeaders); i++ {
		if additionalHeaders[i-1] != requestHeaders[i] {
			t.Fatalf("TestServerMetadataHeadersAdditionalHeaders : [%d]%v != [%d]%v",
				i-1, additionalHeaders[i-1],
				i, requestHeaders[i])
		}
	}
}

func TestConstructorWithProhibitedHeaders(t *testing.T) {
	header1, _ := communication.NewHeader("Header1", "&=$%")
	header2, _ := communication.NewHeader("Header2", "blah blah")

	for _, v := range prohibitedHeaders {
		header3, _ := communication.NewHeader(v, "prohibited header")

		additionalHeaders := []communication.Header{*header1, *header2, *header3}
		builder, _ := NewMetadataProviderBuilder("builder")
		builder.AdditionalRequestHeaders = additionalHeaders
		metadataProvider, err := builder.Build()
		if err == nil {
			t.Fatal("TestConstructorWithProhibitedHeaders : err == nil")
		}

		if metadataProvider != nil {
			t.Fatal("TestConstructorWithProhibitedHeaders : metadataProvider != nil")
		}
	}
}

func AssertServerMetaInfo(prefix string, t *testing.T, metadataProvider *MetadataProvider, integrator string, requestHeader communication.Header) {
	if requestHeader.Name() != "X-WL-ServerMetaInfo" {
		t.Fatalf("[%s]AssertServerMetaInfo : %s != %s", prefix, requestHeader.Name(), "X-WL-ServerMetaInfo")
	}
	if len(requestHeader.Value()) < 1 {
		t.Fatalf("[%s]AssertServerMetaInfo : requestHeader value is empty", prefix)
	}

	serverMetaInfo := metadataProvider.serverMetaInfo
	requestHeaderBytes, err := base64.StdEncoding.DecodeString(requestHeader.Value())
	if err != nil {
		t.Fatalf("[%s]AssertServerMetaInfo : %v", prefix, err)
	}

	err = json.Unmarshal(requestHeaderBytes, &serverMetaInfo)
	if err != nil {
		t.Fatalf("[%s]AssertServerMetaInfo : %v", prefix, err)
	}

	if serverMetaInfo.PlatformIdentifier != metadataProvider.serverMetaInfo.PlatformIdentifier {
		t.Fatalf("[%s]AssertServerMetaInfo : platformIdentifier '%s' != '%s'", prefix,
			serverMetaInfo.PlatformIdentifier, metadataProvider.serverMetaInfo.PlatformIdentifier)
	}
	if serverMetaInfo.SDKIdentifier != metadataProvider.serverMetaInfo.SDKIdentifier {
		t.Fatalf("[%s]AssertServerMetaInfo : SDKIdentifier '%s' != '%s'", prefix,
			serverMetaInfo.SDKIdentifier, metadataProvider.serverMetaInfo.SDKIdentifier)
	}
	if serverMetaInfo.Integrator != integrator {
		t.Fatalf("[%s]AssertServerMetaInfo : Integrator '%s' != '%s'", prefix,
			serverMetaInfo.Integrator, integrator)
	}
	if serverMetaInfo.SDKCreator != "Worldline" {
		t.Fatalf("[%s]AssertServerMetaInfo : SDKCreator '%s' != '%s'", prefix,
			serverMetaInfo.SDKCreator, "Worldline")
	}
}

func AssertShoppingCard(prefix string, t *testing.T, metadataProvider *MetadataProvider, shoppingCard *domain.ShoppingCartExtension) {
	serverMetaInfo := metadataProvider.serverMetaInfo

	if serverMetaInfo.ShoppingCartExtension.Creator != shoppingCard.Creator {
		t.Fatalf("[%s]AssertShoppingCard : Creator '%s' != '%s'", prefix,
			serverMetaInfo.ShoppingCartExtension.Creator, shoppingCard.Creator)
	}
	if serverMetaInfo.ShoppingCartExtension.Name != shoppingCard.Name {
		t.Fatalf("[%s]AssertShoppingCard : Name '%s' != '%s'", prefix,
			serverMetaInfo.ShoppingCartExtension.Name, shoppingCard.Name)
	}
	if serverMetaInfo.ShoppingCartExtension.Version != shoppingCard.Version {
		t.Fatalf("[%s]AssertShoppingCard : Version '%s' != '%s'", prefix,
			serverMetaInfo.ShoppingCartExtension.Version, shoppingCard.Version)
	}
	if serverMetaInfo.ShoppingCartExtension.ExtensionID != shoppingCard.ExtensionID {
		t.Fatalf("[%s]AssertShoppingCard : ExtensionID '%s' != '%s'", prefix,
			serverMetaInfo.ShoppingCartExtension.ExtensionID, shoppingCard.ExtensionID)
	}
}
