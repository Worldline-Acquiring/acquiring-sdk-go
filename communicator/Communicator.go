package communicator

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/Worldline-Acquiring/acquiring-sdk-go/authentication"
	"github.com/Worldline-Acquiring/acquiring-sdk-go/communicator/communication"
	commErrors "github.com/Worldline-Acquiring/acquiring-sdk-go/communicator/errors"
	"github.com/Worldline-Acquiring/acquiring-sdk-go/json"
	"github.com/Worldline-Acquiring/acquiring-sdk-go/logging"
	"github.com/Worldline-Acquiring/acquiring-sdk-go/logging/obfuscation"
)

// A Communicator is used to communicate with the Worldline Acquiring platform web services.
// It contains all the logic to transform a request object to an HTTP request and an HTTP response to a response object.
// It is also thread safe.
type Communicator struct {
	apiEndpoint      *url.URL
	connection       Connection
	metadataProvider *MetadataProvider
	authenticator    authentication.Authenticator
	marshaller       json.Marshaller
}

// APIEndpoint returns the apiEndpoint of the Communicator
func (c *Communicator) APIEndpoint() *url.URL {
	return c.apiEndpoint
}

// Connection returns the connection of the Communicator
func (c *Communicator) Connection() Connection {
	return c.connection
}

// MetadataProvider returns the metadataProvider of the Communicator
func (c *Communicator) MetadataProvider() *MetadataProvider {
	return c.metadataProvider
}

// Authenticator returns the authenticator of the Communicator
func (c *Communicator) Authenticator() authentication.Authenticator {
	return c.authenticator
}

// Marshaller returns the marshaller of the Communicator
func (c *Communicator) Marshaller() json.Marshaller {
	return c.marshaller
}

// Close closes the connection of the Communicator
func (c *Communicator) Close() error {
	return c.connection.Close()
}

// SetBodyObfuscator sets the body obfuscator to use.
func (c *Communicator) SetBodyObfuscator(bodyObfuscator obfuscation.BodyObfuscator) {
	c.connection.SetBodyObfuscator(bodyObfuscator)
}

// SetHeaderObfuscator sets the header obfuscator to use.
func (c *Communicator) SetHeaderObfuscator(headerObfuscator obfuscation.HeaderObfuscator) {
	c.connection.SetHeaderObfuscator(headerObfuscator)
}

// EnableLogging turns on logging using the given communicator logger.
func (c *Communicator) EnableLogging(logger logging.CommunicatorLogger) {
	c.connection.EnableLogging(logger)
}

// DisableLogging turns off logging.
func (c *Communicator) DisableLogging() {
	c.connection.DisableLogging()
}

// Get corresponds to the HTTP Get method
//
// relativePath is the Path to call, relative to the base URI
// requestHeaders is an optimal list of request headers
// requestParameters is an optional set of request parameters. If not used set to nil
// context is an optional Call context which can be used
// expectedObject is a reference to the expected response object
func (c *Communicator) Get(relativePath string, requestHeaders []communication.Header, requestParameters communication.ParamRequest, context *CallContext, expectedObject interface{}) error {
	var requestParameterList []communication.RequestParam
	if requestParameters != nil {
		requestParameterList = requestParameters.ToRequestParameters()
	}
	uri, err := c.toAbsoluteURI(relativePath, requestParameterList)
	if err != nil {
		return err
	}

	requestHeaders, err = c.addGenericHeaders(http.MethodGet, uri, requestHeaders, context)
	if err != nil {
		return err
	}

	_, err = c.connection.Get(uri, requestHeaders, func(statusCode int, headers []communication.Header, reader io.Reader) (interface{}, error) {
		return nil, c.processResponse(statusCode, reader, headers, relativePath, context, expectedObject)
	})
	return err
}

// GetWithHandler corresponds to the HTTP Get method
//
// relativePath is the Path to call, relative to the base URI
// requestHeaders is an optimal list of request headers
// requestParameters is an optional set of request parameters. If not used set to nil
// context is an optional Call context which can be used
// expectedObject is a reference to the expected response object
// bodyHandler is a BodyHandler that handles the body stream
func (c *Communicator) GetWithHandler(relativePath string, requestHeaders []communication.Header, requestParameters communication.ParamRequest, context *CallContext, bodyHandler BodyHandler) error {
	var requestParameterList []communication.RequestParam
	if requestParameters != nil {
		requestParameterList = requestParameters.ToRequestParameters()
	}
	uri, err := c.toAbsoluteURI(relativePath, requestParameterList)
	if err != nil {
		return err
	}

	requestHeaders, err = c.addGenericHeaders(http.MethodGet, uri, requestHeaders, context)
	if err != nil {
		return err
	}

	_, err = c.connection.Get(uri, requestHeaders, func(statusCode int, headers []communication.Header, reader io.Reader) (interface{}, error) {
		return nil, c.processResponseWithHandler(statusCode, reader, headers, relativePath, context, bodyHandler)
	})
	return err
}

// Delete corresponds to the HTTP Delete method
//
// relativePath is the Path to call, relative to the base URI
// requestHeaders is an optimal list of request headers
// requestParameters is an optional set of request parameters. If not used set to nil
// context is an optional Call context which can be used
// expectedObject is a reference to the expected response object
func (c *Communicator) Delete(relativePath string, requestHeaders []communication.Header, requestParameters communication.ParamRequest, context *CallContext, expectedObject interface{}) error {
	var requestParameterList []communication.RequestParam
	if requestParameters != nil {
		requestParameterList = requestParameters.ToRequestParameters()
	}
	uri, err := c.toAbsoluteURI(relativePath, requestParameterList)
	if err != nil {
		return err
	}

	requestHeaders, err = c.addGenericHeaders(http.MethodDelete, uri, requestHeaders, context)
	if err != nil {
		return err
	}

	_, err = c.connection.Delete(uri, requestHeaders, func(statusCode int, headers []communication.Header, reader io.Reader) (interface{}, error) {
		return nil, c.processResponse(statusCode, reader, headers, relativePath, context, expectedObject)
	})
	return err
}

// DeleteWithHandler corresponds to the HTTP Delete method
//
// relativePath is the Path to call, relative to the base URI
// requestHeaders is an optimal list of request headers
// requestParameters is an optional set of request parameters. If not used set to nil
// context is an optional Call context which can be used
// expectedObject is a reference to the expected response object
// bodyHandler is a BodyHandler that handles the body stream
func (c *Communicator) DeleteWithHandler(relativePath string, requestHeaders []communication.Header, requestParameters communication.ParamRequest, context *CallContext, bodyHandler BodyHandler) error {
	var requestParameterList []communication.RequestParam
	if requestParameters != nil {
		requestParameterList = requestParameters.ToRequestParameters()
	}
	uri, err := c.toAbsoluteURI(relativePath, requestParameterList)
	if err != nil {
		return err
	}

	requestHeaders, err = c.addGenericHeaders(http.MethodDelete, uri, requestHeaders, context)
	if err != nil {
		return err
	}

	_, err = c.connection.Delete(uri, requestHeaders, func(statusCode int, headers []communication.Header, reader io.Reader) (interface{}, error) {
		return nil, c.processResponseWithHandler(statusCode, reader, headers, relativePath, context, bodyHandler)
	})
	return err
}

// Post corresponds to the HTTP Post method
//
// relativePath is the Path to call, relative to the base URI
// requestHeaders is an optimal list of request headers
// requestParameters is an optional set of request parameters. If not used set it to nil
// requestBody is the body of the request. If not used set to nil
// context is an optional Call context which can be used. If not used set it to nil
// expectedObject is a reference to the expected response object
func (c *Communicator) Post(relativePath string, requestHeaders []communication.Header, requestParameters communication.ParamRequest, requestBody interface{}, context *CallContext, expectedResponse interface{}) error {
	if multipartObject, ok := requestBody.(communication.MultipartFormDataObject); ok {
		return c.postMultipart(relativePath, requestHeaders, requestParameters, &multipartObject, context, expectedResponse)
	}
	if multipartObject, ok := requestBody.(*communication.MultipartFormDataObject); ok {
		return c.postMultipart(relativePath, requestHeaders, requestParameters, multipartObject, context, expectedResponse)
	}
	if multipartRequest, ok := requestBody.(communication.MultipartFormDataRequest); ok {
		return c.postMultipart(relativePath, requestHeaders, requestParameters, multipartRequest.ToMultipartFormDataObject(), context, expectedResponse)
	}

	var requestParameterList []communication.RequestParam
	if requestParameters != nil {
		requestParameterList = requestParameters.ToRequestParameters()
	}
	uri, err := c.toAbsoluteURI(relativePath, requestParameterList)
	if err != nil {
		return err
	}

	var requestJSON string
	if requestBody != nil {
		contentTypeHeader, _ := communication.NewHeader("Content-Type", "application/json")
		requestHeaders = append(requestHeaders, *contentTypeHeader)

		requestJSON, err = c.marshaller.Marshal(requestBody)
		if err != nil {
			return err
		}
	}

	requestHeaders, err = c.addGenericHeaders(http.MethodPost, uri, requestHeaders, context)
	if err != nil {
		return err
	}

	_, err = c.connection.Post(uri, requestHeaders, requestJSON, func(statusCode int, headers []communication.Header, reader io.Reader) (interface{}, error) {
		return nil, c.processResponse(statusCode, reader, headers, relativePath, context, expectedResponse)
	})
	return err
}

func (c *Communicator) postMultipart(relativePath string, requestHeaders []communication.Header, requestParameters communication.ParamRequest, requestBody *communication.MultipartFormDataObject, context *CallContext, expectedResponse interface{}) error {
	var requestParameterList []communication.RequestParam
	if requestParameters != nil {
		requestParameterList = requestParameters.ToRequestParameters()
	}
	uri, err := c.toAbsoluteURI(relativePath, requestParameterList)
	if err != nil {
		return err
	}

	contentTypeHeader, _ := communication.NewHeader("Content-Type", requestBody.GetContentType())
	requestHeaders = append(requestHeaders, *contentTypeHeader)

	requestHeaders, err = c.addGenericHeaders(http.MethodPost, uri, requestHeaders, context)
	if err != nil {
		return err
	}

	_, err = c.connection.PostMultipart(uri, requestHeaders, requestBody, func(statusCode int, headers []communication.Header, reader io.Reader) (interface{}, error) {
		return nil, c.processResponse(statusCode, reader, headers, relativePath, context, expectedResponse)
	})
	return err
}

// PostWithHandler corresponds to the HTTP Post method
//
// relativePath is the Path to call, relative to the base URI
// requestHeaders is an optimal list of request headers
// requestParameters is an optional set of request parameters. If not used set it to nil
// requestBody is the body of the request. If not used set to nil
// context is an optional Call context which can be used. If not used set it to nil
// expectedObject is a reference to the expected response object
// bodyHandler is a BodyHandler that handles the body stream
func (c *Communicator) PostWithHandler(relativePath string, requestHeaders []communication.Header, requestParameters communication.ParamRequest, requestBody interface{}, context *CallContext, bodyHandler BodyHandler) error {
	if multipartObject, ok := requestBody.(communication.MultipartFormDataObject); ok {
		return c.postMultipartWithHandler(relativePath, requestHeaders, requestParameters, &multipartObject, context, bodyHandler)
	}
	if multipartObject, ok := requestBody.(*communication.MultipartFormDataObject); ok {
		return c.postMultipartWithHandler(relativePath, requestHeaders, requestParameters, multipartObject, context, bodyHandler)
	}
	if multipartRequest, ok := requestBody.(communication.MultipartFormDataRequest); ok {
		return c.postMultipartWithHandler(relativePath, requestHeaders, requestParameters, multipartRequest.ToMultipartFormDataObject(), context, bodyHandler)
	}

	var requestParameterList []communication.RequestParam
	if requestParameters != nil {
		requestParameterList = requestParameters.ToRequestParameters()
	}
	uri, err := c.toAbsoluteURI(relativePath, requestParameterList)
	if err != nil {
		return err
	}

	var requestJSON string
	if requestBody != nil {
		contentTypeHeader, _ := communication.NewHeader("Content-Type", "application/json")
		requestHeaders = append(requestHeaders, *contentTypeHeader)

		requestJSON, err = c.marshaller.Marshal(requestBody)
		if err != nil {
			return err
		}
	}

	requestHeaders, err = c.addGenericHeaders(http.MethodPost, uri, requestHeaders, context)
	if err != nil {
		return err
	}

	_, err = c.connection.Post(uri, requestHeaders, requestJSON, func(statusCode int, headers []communication.Header, reader io.Reader) (interface{}, error) {
		return nil, c.processResponseWithHandler(statusCode, reader, headers, relativePath, context, bodyHandler)
	})

	return err
}

func (c *Communicator) postMultipartWithHandler(relativePath string, requestHeaders []communication.Header, requestParameters communication.ParamRequest, requestBody *communication.MultipartFormDataObject, context *CallContext, bodyHandler BodyHandler) error {
	var requestParameterList []communication.RequestParam
	if requestParameters != nil {
		requestParameterList = requestParameters.ToRequestParameters()
	}
	uri, err := c.toAbsoluteURI(relativePath, requestParameterList)
	if err != nil {
		return err
	}

	contentTypeHeader, _ := communication.NewHeader("Content-Type", requestBody.GetContentType())
	requestHeaders = append(requestHeaders, *contentTypeHeader)

	requestHeaders, err = c.addGenericHeaders(http.MethodPost, uri, requestHeaders, context)
	if err != nil {
		return err
	}

	_, err = c.connection.PostMultipart(uri, requestHeaders, requestBody, func(statusCode int, headers []communication.Header, reader io.Reader) (interface{}, error) {
		return nil, c.processResponseWithHandler(statusCode, reader, headers, relativePath, context, bodyHandler)
	})
	return err
}

// Put corresponds to the HTTP Put method
//
// relativePath is the Path to call, relative to the base URI
// requestHeaders is an optimal list of request headers
// requestParameters is an optional set of request parameters. If not used set it to nil
// requestBody is the body of the request. If not used set to nil
// context is an optional Call context which can be used. If not used set it to nil
// expectedObject is a reference to the expected response object
func (c *Communicator) Put(relativePath string, requestHeaders []communication.Header, requestParameters communication.ParamRequest, requestBody interface{}, context *CallContext, expectedObject interface{}) error {
	if multipartObject, ok := requestBody.(communication.MultipartFormDataObject); ok {
		return c.putMultipart(relativePath, requestHeaders, requestParameters, &multipartObject, context, expectedObject)
	}
	if multipartObject, ok := requestBody.(*communication.MultipartFormDataObject); ok {
		return c.putMultipart(relativePath, requestHeaders, requestParameters, multipartObject, context, expectedObject)
	}
	if multipartRequest, ok := requestBody.(communication.MultipartFormDataRequest); ok {
		return c.putMultipart(relativePath, requestHeaders, requestParameters, multipartRequest.ToMultipartFormDataObject(), context, expectedObject)
	}

	var requestParameterList []communication.RequestParam
	if requestParameters != nil {
		requestParameterList = requestParameters.ToRequestParameters()
	}
	uri, err := c.toAbsoluteURI(relativePath, requestParameterList)
	if err != nil {
		return err
	}

	var requestJSON string
	if requestBody != nil {
		contentTypeHeader, _ := communication.NewHeader("Content-Type", "application/json")
		requestHeaders = append(requestHeaders, *contentTypeHeader)

		requestJSON, err = c.marshaller.Marshal(requestBody)
		if err != nil {
			return err
		}
	}

	requestHeaders, err = c.addGenericHeaders(http.MethodPut, uri, requestHeaders, context)
	if err != nil {
		return err
	}

	_, err = c.connection.Put(uri, requestHeaders, requestJSON, func(statusCode int, headers []communication.Header, reader io.Reader) (interface{}, error) {
		return nil, c.processResponse(statusCode, reader, headers, relativePath, context, expectedObject)
	})
	return err
}

func (c *Communicator) putMultipart(relativePath string, requestHeaders []communication.Header, requestParameters communication.ParamRequest, requestBody *communication.MultipartFormDataObject, context *CallContext, expectedResponse interface{}) error {
	var requestParameterList []communication.RequestParam
	if requestParameters != nil {
		requestParameterList = requestParameters.ToRequestParameters()
	}
	uri, err := c.toAbsoluteURI(relativePath, requestParameterList)
	if err != nil {
		return err
	}

	contentTypeHeader, _ := communication.NewHeader("Content-Type", requestBody.GetContentType())
	requestHeaders = append(requestHeaders, *contentTypeHeader)

	requestHeaders, err = c.addGenericHeaders(http.MethodPost, uri, requestHeaders, context)
	if err != nil {
		return err
	}

	_, err = c.connection.PutMultipart(uri, requestHeaders, requestBody, func(statusCode int, headers []communication.Header, reader io.Reader) (interface{}, error) {
		return nil, c.processResponse(statusCode, reader, headers, relativePath, context, expectedResponse)
	})
	return err
}

// PutWithHandler corresponds to the HTTP Put method
//
// relativePath is the Path to call, relative to the base URI
// requestHeaders is an optimal list of request headers
// requestParameters is an optional set of request parameters. If not used set it to nil
// requestBody is the body of the request. If not used set to nil
// context is an optional Call context which can be used. If not used set it to nil
// expectedObject is a reference to the expected response object
// bodyHandler is a BodyHandler that handles the body stream
func (c *Communicator) PutWithHandler(relativePath string, requestHeaders []communication.Header, requestParameters communication.ParamRequest, requestBody interface{}, context *CallContext, bodyHandler BodyHandler) error {
	if multipartObject, ok := requestBody.(communication.MultipartFormDataObject); ok {
		return c.putMultipartWithHandler(relativePath, requestHeaders, requestParameters, &multipartObject, context, bodyHandler)
	}
	if multipartObject, ok := requestBody.(*communication.MultipartFormDataObject); ok {
		return c.putMultipartWithHandler(relativePath, requestHeaders, requestParameters, multipartObject, context, bodyHandler)
	}
	if multipartRequest, ok := requestBody.(communication.MultipartFormDataRequest); ok {
		return c.putMultipartWithHandler(relativePath, requestHeaders, requestParameters, multipartRequest.ToMultipartFormDataObject(), context, bodyHandler)
	}

	var requestParameterList []communication.RequestParam
	if requestParameters != nil {
		requestParameterList = requestParameters.ToRequestParameters()
	}
	uri, err := c.toAbsoluteURI(relativePath, requestParameterList)
	if err != nil {
		return err
	}

	var requestJSON string
	if requestBody != nil {
		contentTypeHeader, _ := communication.NewHeader("Content-Type", "application/json")
		requestHeaders = append(requestHeaders, *contentTypeHeader)

		requestJSON, err = c.marshaller.Marshal(requestBody)
		if err != nil {
			return err
		}
	}

	requestHeaders, err = c.addGenericHeaders(http.MethodPut, uri, requestHeaders, context)
	if err != nil {
		return err
	}

	_, err = c.connection.Put(uri, requestHeaders, requestJSON, func(statusCode int, headers []communication.Header, reader io.Reader) (interface{}, error) {
		return nil, c.processResponseWithHandler(statusCode, reader, headers, relativePath, context, bodyHandler)
	})
	return err
}

func (c *Communicator) putMultipartWithHandler(relativePath string, requestHeaders []communication.Header, requestParameters communication.ParamRequest, requestBody *communication.MultipartFormDataObject, context *CallContext, bodyHandler BodyHandler) error {
	var requestParameterList []communication.RequestParam
	if requestParameters != nil {
		requestParameterList = requestParameters.ToRequestParameters()
	}
	uri, err := c.toAbsoluteURI(relativePath, requestParameterList)
	if err != nil {
		return err
	}

	contentTypeHeader, _ := communication.NewHeader("Content-Type", requestBody.GetContentType())
	requestHeaders = append(requestHeaders, *contentTypeHeader)

	requestHeaders, err = c.addGenericHeaders(http.MethodPost, uri, requestHeaders, context)
	if err != nil {
		return err
	}

	_, err = c.connection.PutMultipart(uri, requestHeaders, requestBody, func(statusCode int, headers []communication.Header, reader io.Reader) (interface{}, error) {
		return nil, c.processResponseWithHandler(statusCode, reader, headers, relativePath, context, bodyHandler)
	})
	return err
}

// CloseExpiredConnections is a utility method that delegates the call to this communicator's connection.
// Also see Connection.CloseExpiredConnections
func (c *Communicator) CloseExpiredConnections() {
	c.connection.CloseExpiredConnections()
}

// CloseIdleConnections is a utility method that delegates the call to this communicator's connection.
// The duration argument is a specification of how long the connection has to be Idle.
// Also see Connection.CloseIdleConnections
func (c *Communicator) CloseIdleConnections(duration time.Duration) {
	c.connection.CloseIdleConnections(duration)
}

func (c *Communicator) toAbsoluteURI(relativePath string, requestParameters []communication.RequestParam) (url.URL, error) {
	if c.apiEndpoint.Path != "" {
		return url.URL{}, errors.New("apiEndpoint should not contain a path")
	}
	if c.apiEndpoint.User != nil || c.apiEndpoint.Query().Encode() != "" || c.apiEndpoint.Fragment != "" {
		return url.URL{}, errors.New("apiEndpoint should not contain user info, query or fragment")
	}

	var absolutePath string
	if strings.Index(relativePath, "/") == 0 {
		absolutePath = relativePath
	} else {
		absolutePath = "/" + relativePath
	}

	var rawQuery = ""
	for index, element := range requestParameters {
		if index != 0 {
			rawQuery += "&"
		}
		id := url.QueryEscape(element.Name())
		value := url.QueryEscape(element.Value())
		rawQuery += id + "=" + value
	}

	var absoluteURI = url.URL{Scheme: c.apiEndpoint.Scheme,
		Host:     c.apiEndpoint.Host,
		Path:     absolutePath,
		RawQuery: rawQuery,
	}
	return absoluteURI, nil
}

// Adds the necessary headers to the given list of headers. This includes the authorization header,
// which uses other headers, so when you need to override this method,
// make sure to call AddGenericHeaders at the end of your overridden method.
func (c *Communicator) addGenericHeaders(httpMethod string, url url.URL, requestHeaders []communication.Header, context *CallContext) ([]communication.Header, error) {
	requestHeaders = append(requestHeaders, c.metadataProvider.MetadataHeaders()...)

	header, _ := communication.NewHeader("Date", getHeaderDateString())
	requestHeaders = append(requestHeaders, *header)

	// no context specific headers at this time

	//add authorization
	authorization, err := c.authenticator.GetAuthorization(httpMethod, url, requestHeaders)
	if err != nil {
		return nil, err
	}

	header, _ = communication.NewHeader("Authorization", authorization)
	requestHeaders = append(requestHeaders, *header)

	traceID, err := communication.PseudoUUID()
	if err != nil {
		return nil, err
	}
	header, _ = communication.NewHeader("Trace-Id", traceID)
	requestHeaders = append(requestHeaders, *header)

	return requestHeaders, nil
}

// Gets the date in the preferred format for the HTTP date header (RFC1123).
func getHeaderDateString() string {
	return time.Now().UTC().Format(http.TimeFormat)
}

// Checks the Response for errors and creates an error if necessary.
func (c *Communicator) createErrorIfNecessary(statusCode int, reader io.Reader, headers []communication.Header, requestPath string) error {
	if statusCode < http.StatusOK || statusCode >= http.StatusMultipleChoices {
		bodyBuff, err := ioutil.ReadAll(reader)
		if err != nil {
			return err
		}

		body := string(bodyBuff)
		if body != "" && !isJSON(headers) {
			cause, _ := commErrors.NewResponseError(statusCode, body, headers)
			if statusCode == http.StatusNotFound {
				err, _ := commErrors.NewNotFoundErrorVerbose("The requested resource was not found; invalid path: "+requestPath, cause)
				return err
			}
			err, _ := commErrors.NewCommunicationError(cause)
			return err
		}
		responseErr, _ := commErrors.NewResponseError(statusCode, body, headers)
		return responseErr
	}
	return nil
}

func (c *Communicator) processResponse(statusCode int, reader io.Reader, headers []communication.Header, requestPath string, context *CallContext, expectedObject interface{}) error {
	err := c.updateContext(headers, context)
	if err != nil {
		return err
	}

	err = c.createErrorIfNecessary(statusCode, reader, headers, requestPath)
	if err != nil {
		return err
	}

	return c.marshaller.UnmarshalFromReader(reader, expectedObject)
}

func (c *Communicator) processResponseWithHandler(statusCode int, reader io.Reader, headers []communication.Header, requestPath string, context *CallContext, bodyHandler BodyHandler) error {
	err := c.updateContext(headers, context)
	if err != nil {
		return err
	}

	err = c.createErrorIfNecessary(statusCode, reader, headers, requestPath)
	if err != nil {
		return err
	}

	return bodyHandler(headers, reader)
}

func (c *Communicator) updateContext(headers []communication.Header, context *CallContext) error {
	// no context specific headers at this time
	return nil
}

func isJSON(headers []communication.Header) bool {
	header := communication.Headers(headers).GetHeader("Content-Type")
	if header == nil {
		return true
	}

	contentType := strings.ToLower(header.Value())
	return contentType == "application/json" || strings.HasPrefix(contentType, "application/json") ||
		contentType == "application/problem+json" || strings.HasPrefix(contentType, "application/problem+json")
}

// NewCommunicator creates a communicator with the given apiEndpoint, connection, authenticator, metadata provider and marshaller
func NewCommunicator(apiEndpoint *url.URL, connection Connection, authenticator authentication.Authenticator, metadataProvider *MetadataProvider, marshaller json.Marshaller) (*Communicator, error) {
	if apiEndpoint == nil {
		return nil, errors.New("apiEndpoint is required")
	}
	if len(apiEndpoint.Path) > 0 && apiEndpoint.Path != "/" {
		return nil, errors.New("apiEndpoint should not contain a path")
	}
	if apiEndpoint.User != nil || apiEndpoint.RawQuery != "" || apiEndpoint.Fragment != "" {
		return nil, errors.New("apiEndpoint should not contain user info, query or fragment")
	}
	if connection == nil {
		return nil, errors.New("connection is required")
	}
	if authenticator == nil {
		return nil, errors.New("authenticator is required")
	}
	if metadataProvider == nil {
		return nil, errors.New("metadataProvider is required")
	}
	if marshaller == nil {
		return nil, errors.New("marshaller is required")
	}

	return &Communicator{apiEndpoint, connection, metadataProvider, authenticator, marshaller}, nil
}
