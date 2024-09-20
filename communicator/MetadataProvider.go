package communicator

import (
	"encoding/base64"
	"errors"
	"runtime"
	"strings"

	"github.com/Worldline-Acquiring/acquiring-sdk-go/communicator/communication"
	"github.com/Worldline-Acquiring/acquiring-sdk-go/domain"
	"github.com/Worldline-Acquiring/acquiring-sdk-go/json"
)

type serverMetaInfo struct {
	PlatformIdentifier    string                        `json:"platformIdentifier,omitempty"`
	SDKIdentifier         string                        `json:"sdkIdentifier,omitempty"`
	SDKCreator            string                        `json:"sdkCreator,omitempty"`
	Integrator            string                        `json:"integrator,omitempty"`
	ShoppingCartExtension *domain.ShoppingCartExtension `json:"shoppingCartExtension,omitempty"`
}

// MetadataProvider provides meta info about the server. Thread-safe.
type MetadataProvider struct {
	serverMetaInfo  serverMetaInfo
	metadataHeaders []communication.Header
}

// prohibitedHeaders are the headers that can't be included in the server requests
var prohibitedHeaders = []string{serverMetaInfoHeader, "Date", "Content-Type", "Authorization"}

func getPlatformIdentifier() string {
	//I believe there is no standard function thus we need to do something like
	//https://github.com/matishsiao/goInfo
	return runtime.GOOS + " " + runtime.Version() + "(" + runtime.GOARCH + ")"
}

const sdkVersion = "0.1.0"
const sdkIdentifier = "GoServerSDK/v" + sdkVersion
const serverMetaInfoHeader = "X-WL-ServerMetaInfo"

// MetadataHeaders returns the server related headers containing the metadata to be associated with the request (if any).
// This will always contain at least an automatically generated header X-WL-ServerMetaInfo.
func (m MetadataProvider) MetadataHeaders() []communication.Header {
	// Return a clone instead of the original slice - immutability insurance
	return append([]communication.Header{}, m.metadataHeaders...)
}

// NewMetadataProvider creates a MetadataProvider
func NewMetadataProvider(integrator string) (*MetadataProvider, error) {
	return newMetadataProvider(integrator, nil, nil)
}

func newMetadataProvider(integrator string, shoppingCartExtension *domain.ShoppingCartExtension, additionalRequestHeaders []communication.Header) (*MetadataProvider, error) {
	if strings.TrimSpace(integrator) == "" {
		return nil, errors.New("integrator is required")
	}

	err := validateAdditionalRequestHeaders(additionalRequestHeaders)
	if err != nil {
		return nil, err
	}

	serverMetaInfo := serverMetaInfo{
		getPlatformIdentifier(),
		sdkIdentifier,
		"Worldline",
		integrator,
		shoppingCartExtension,
	}
	marshaller := json.DefaultMarshaller()

	serverMetaInfoString, err := marshaller.Marshal(&serverMetaInfo)
	if err != nil {
		return nil, err
	}

	serverMetaInfoBase64 := base64.StdEncoding.EncodeToString([]byte(serverMetaInfoString))
	header, _ := communication.NewHeader(serverMetaInfoHeader, serverMetaInfoBase64)
	headers := []communication.Header{*header}
	headers = append(headers, additionalRequestHeaders...)
	return &MetadataProvider{serverMetaInfo, headers}, nil
}

func validateAdditionalRequestHeader(additionalRequestHeader *communication.Header) error {
	for _, a := range prohibitedHeaders {
		if a == additionalRequestHeader.Name() {
			return errors.New("header is prohibited")
		}
	}
	return nil
}

func validateAdditionalRequestHeaders(additionalRequestHeaders []communication.Header) error {
	for _, additionalRequestHeader := range additionalRequestHeaders {
		err := validateAdditionalRequestHeader(&additionalRequestHeader)
		if err != nil {
			return err
		}
	}
	return nil
}
