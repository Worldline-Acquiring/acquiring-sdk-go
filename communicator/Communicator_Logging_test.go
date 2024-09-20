package communicator

import (
	"errors"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/Worldline-Acquiring/acquiring-sdk-go/communicator/communication"
	commErrors "github.com/Worldline-Acquiring/acquiring-sdk-go/communicator/errors"
	"github.com/Worldline-Acquiring/acquiring-sdk-go/json"
	"github.com/Worldline-Acquiring/acquiring-sdk-go/logging"
)

var getJSON = `{
	"result": "OK"
}`

var postWithCreatedResponseJSON = `{
	"payment": {
		"id": "000000123410000595980000100001",
		"status": "PENDING_APPROVAL"
	}
}`

var postWithBadRequestResponseJSON = `{
	"errorId": "0953f236-9e54-4f23-9556-d66bc757dda8",
	"errors": [{
		"code": "21000020",
		"requestId": "24146",
		"message": "VALUE **************** OF FIELD CREDITCARDNUMBER DID NOT PASS THE LUHNCHECK",
		"httpStatusCode": 400
	}]
}`

var unknownServerErrorJSON = `{
	"errorId": "fbff1179-7ba4-4894-9021-d8a0011d23a7",
	"errors": [{
		"code": "9999",
		"message": "UNKNOWN_SERVER_ERROR",
		"httpStatusCode": 500
	}]
}`

var notFoundErrorHTML = `Not Found`

func TestGetWithoutQueryParams(t *testing.T) {
	logPrefix := "TestGetWithoutQueryParams"

	responseHeaders := map[string]string{
		"Content-Type": "application/json",
		"Dummy":        "dummy",
	}
	requestHeaders := map[string][]string{}

	listener, sl, comm, err := createTestEnvironment(
		"/v1/get",
		createRecordRequest(http.StatusOK, getJSON, responseHeaders, requestHeaders))
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}
	defer func(listener net.Listener, sl *stoppableListener, comm *Communicator) {
		_ = listener.Close()
		_ = sl.Close()
		_ = comm.Close()
	}(listener, sl, comm)

	logger := &testLogger{}
	comm.EnableLogging(logger)

	response := make(map[string]interface{})
	err = comm.Get("v1/get", nil, nil, nil, &response)
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}

	if result, ok := response["result"]; !ok || result != "OK" {
		t.Fatalf("%v: responseResult %v", logPrefix, result)
	}

	if len(logger.entries) != 2 {
		t.Fatalf("%v: loggerEntries %v", logPrefix, len(logger.entries))
	}

	requestEntry := logger.entries[0]
	if !strings.Contains(requestEntry.message, "method:       'GET'") {
		t.Fatalf("%v: requestEntryRequestMethod %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "uri:          '/v1/get'") {
		t.Fatalf("%v: requestEntryRequestURL %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "Authorization=\"********\"") {
		t.Fatalf("%v: requestEntryAuthorizationHeader missing %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "Date=\"") {
		t.Fatalf("%v: requestEntryDateHeader missing %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "X-WL-ServerMetaInfo=\"") {
		t.Fatalf("%v: requestEntryDateHeader missing %v", logPrefix, requestEntry.message)
	}
	if strings.Contains(requestEntry.message, "body:") {
		t.Fatalf("%v: requeestEntryRequestBody %v", logPrefix, requestEntry.message)
	}

	responseEntry := logger.entries[1]
	if !strings.Contains(responseEntry.message, "status-code:  '200'") {
		t.Fatalf("%v: responseEntryResponseStatusCode %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "content-type: 'application/json'") {
		t.Fatalf("%v: responseEntryResponseContentType %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "Date=\"") {
		t.Fatalf("%v: responseEntryDateHeader missing %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "Dummy=\"dummy\"") {
		t.Fatalf("%v: responseEntryDummyHeader missing %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "body:         '") || strings.Contains(responseEntry.message, "body:         ''") {
		t.Fatalf("%v: responseEntryResponseBody %v", logPrefix, responseEntry.message)
	}
	if responseEntry.err != nil {
		t.Fatalf("%v: responseEntryErr %v", logPrefix, responseEntry.err)
	}
}

func TestGetWithQueryParams(t *testing.T) {
	logPrefix := "TestGetWithQueryParams"

	responseHeaders := map[string]string{
		"Content-Type": "application/json",
		"Dummy":        "dummy",
	}
	requestHeaders := map[string][]string{}

	listener, sl, comm, err := createTestEnvironment(
		"/v1/get",
		createRecordRequest(http.StatusOK, getJSON, responseHeaders, requestHeaders))
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}
	defer func(listener net.Listener, sl *stoppableListener, comm *Communicator) {
		_ = listener.Close()
		_ = sl.Close()
		_ = comm.Close()
	}(listener, sl, comm)

	logger := &testLogger{}
	comm.EnableLogging(logger)

	query := &testParamRequest{}
	query.addRequestParam("source", "EUR")
	query.addRequestParam("target", "USD")
	query.addRequestParam("amount", "1000")

	response := make(map[string]interface{})
	err = comm.Get("v1/get", nil, query, nil, &response)
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}

	if result, ok := response["result"]; !ok || result != "OK" {
		t.Fatalf("%v: responseResult %v", logPrefix, result)
	}

	if len(logger.entries) != 2 {
		t.Fatalf("%v: loggerEntries %v", logPrefix, len(logger.entries))
	}

	requestEntry := logger.entries[0]
	if !strings.Contains(requestEntry.message, "method:       'GET'") {
		t.Fatalf("%v: requestEntryRequestMethod %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "uri:          '/v1/get'") {
		t.Fatalf("%v: requestEntryRequestURL %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "Authorization=\"********\"") {
		t.Fatalf("%v: requestEntryAuthorizationHeader missing %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "Date=\"") {
		t.Fatalf("%v: requestEntryDateHeader missing %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "X-WL-ServerMetaInfo=\"") {
		t.Fatalf("%v: requestEntryDateHeader missing %v", logPrefix, requestEntry.message)
	}
	if strings.Contains(requestEntry.message, "body:") {
		t.Fatalf("%v: requestEntryRequestBody %v", logPrefix, requestEntry.message)
	}
	if requestEntry.err != nil {
		t.Fatalf("%v: requestEntryErr %v", logPrefix, requestEntry.err)
	}

	responseEntry := logger.entries[1]
	if !strings.Contains(responseEntry.message, "status-code:  '200'") {
		t.Fatalf("%v: responseEntryResponseStatusCode %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "content-type: 'application/json'") {
		t.Fatalf("%v: responseEntryResponseContentType %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "Date=\"") {
		t.Fatalf("%v: responseEntryDateHeader missing %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "Dummy=\"dummy\"") {
		t.Fatalf("%v: responseEntryDummyHeader missing %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "body:         '") || strings.Contains(responseEntry.message, "body:         ''") {
		t.Fatalf("%v: responseEntryResponseBody %v", logPrefix, responseEntry.message)
	}
	if responseEntry.err != nil {
		t.Fatalf("%v: responseEntryErr %v", logPrefix, responseEntry.err)
	}
}

func TestDeleteWithVoidResponse(t *testing.T) {
	logPrefix := "TestDeleteWithVoidResponse"

	responseHeaders := map[string]string{
		"Content-Type": "application/json",
		"Dummy":        "dummy",
	}
	requestHeaders := map[string][]string{}

	listener, sl, comm, err := createTestEnvironment(
		"/v1/void",
		createRecordRequest(http.StatusNoContent, "", responseHeaders, requestHeaders))
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}
	defer func(listener net.Listener, sl *stoppableListener, comm *Communicator) {
		_ = listener.Close()
		_ = sl.Close()
		_ = comm.Close()
	}(listener, sl, comm)

	logger := &testLogger{}
	comm.EnableLogging(logger)

	response := map[string]interface{}{}
	err = comm.Delete("v1/void", nil, nil, nil, &response)
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}

	if len(logger.entries) != 2 {
		t.Fatalf("%v: loggerEntries %v", logPrefix, len(logger.entries))
	}

	requestEntry := logger.entries[0]
	if !strings.Contains(requestEntry.message, "method:       'DELETE'") {
		t.Fatalf("%v: requestEntryRequestMethod %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "uri:          '/v1/void'") {
		t.Fatalf("%v: requestEntryRequestURL %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "Authorization=\"********\"") {
		t.Fatalf("%v: requestEntryAuthorizationHeader missing %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "Date=\"") {
		t.Fatalf("%v: requestEntryDateHeader missing %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "X-WL-ServerMetaInfo=\"") {
		t.Fatalf("%v: requestEntryDateHeader missing %v", logPrefix, requestEntry.message)
	}
	if strings.Contains(requestEntry.message, "body:") {
		t.Fatalf("%v: requestEntryRequestBody %v", logPrefix, requestEntry.message)
	}
	if requestEntry.err != nil {
		t.Fatalf("%v: requestEntryErr %v", logPrefix, requestEntry.err)
	}

	responseEntry := logger.entries[1]
	if !strings.Contains(responseEntry.message, "status-code:  '204'") {
		t.Fatalf("%v: responseEntryResponseStatusCode %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "content-type: 'application/json'") {
		t.Fatalf("%v: responseEntryResponseContentType %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "Date=\"") {
		t.Fatalf("%v: responseEntryDateHeader missing %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "Dummy=\"dummy\"") {
		t.Fatalf("%v: responseEntryDummyHeader missing %v", logPrefix, responseEntry.message)
	}
	if responseEntry.err != nil {
		t.Fatalf("%v: responseEntryErr %v", logPrefix, responseEntry.err)
	}
	if !strings.Contains(responseEntry.message, "body:         ''") {
		t.Fatalf("%v: responseEntryResponseBody %v", logPrefix, responseEntry.message)
	}
}

func TestPostWithCreatedResponse(t *testing.T) {
	logPrefix := "TestPostWithCreatedResponse"

	responseHeaders := map[string]string{
		"Content-Type": "application/json",
		"Dummy":        "dummy",
		"Location":     "http://localhost/v1/created/000000123410000595980000100001",
	}
	requestHeaders := map[string][]string{}

	listener, sl, comm, err := createTestEnvironment(
		"/v1/created",
		createRecordRequest(http.StatusCreated, postWithCreatedResponseJSON, responseHeaders, requestHeaders))
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}
	defer func(listener net.Listener, sl *stoppableListener, comm *Communicator) {
		_ = listener.Close()
		_ = sl.Close()
		_ = comm.Close()
	}(listener, sl, comm)

	logger := &testLogger{}
	comm.EnableLogging(logger)

	cardData := map[string]interface{}{
		"cardNumber":       "1234567890123456",
		"cardSecurityCode": "123",
		"expiryDate":       "122024",
	}

	request := map[string]interface{}{
		"cardData": cardData,
	}

	response := map[string]interface{}{}
	err = comm.Post("v1/created", nil, nil, request, nil, &response)
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}

	if len(logger.entries) != 2 {
		t.Fatalf("%v: loggerEntries %v", logPrefix, len(logger.entries))
	}

	if _, ok := response["payment"]; !ok {
		t.Fatalf("%v: responsePayment nil", logPrefix)
	}
	if payment, isMap := response["payment"].(map[string]interface{}); isMap {
		if _, ok := payment["id"]; !ok {
			t.Fatalf("%v: responsePaymentID nil", logPrefix)
		}
	} else {
		t.Fatalf("%v: responsePayment not a map", logPrefix)
	}

	requestEntry := logger.entries[0]
	if !strings.Contains(requestEntry.message, "method:       'POST'") {
		t.Fatalf("%v: requestEntryRequestMethod %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "uri:          '/v1/created'") {
		t.Fatalf("%v: requestEntryRequestURL %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "Authorization=\"********\"") {
		t.Fatalf("%v: requestEntryAuthorizationHeader missing %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "Date=\"") {
		t.Fatalf("%v: requestEntryDateHeader missing %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "X-WL-ServerMetaInfo=\"") {
		t.Fatalf("%v: requestEntryDateHeader missing %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "body:         '") || strings.Contains(requestEntry.message, "body:         ''") {
		t.Fatalf("%v: requestEntryRequestBody %v", logPrefix, requestEntry.message)
	}
	if strings.Contains(requestEntry.message, "\"1234567890123456\"") || strings.Contains(requestEntry.message, "\"123\"") || strings.Contains(requestEntry.message, "\"122024\"") {
		t.Fatalf("%v: requestEntryRequestBodyObfuscation %v", logPrefix, requestEntry.message)
	}
	if requestEntry.err != nil {
		t.Fatalf("%v: requestEntryErr %v", logPrefix, requestEntry.err)
	}

	responseEntry := logger.entries[1]
	if !strings.Contains(responseEntry.message, "status-code:  '201'") {
		t.Fatalf("%v: responseEntryResponseStatusCode %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "content-type: 'application/json'") {
		t.Fatalf("%v: responseEntryResponseContentType %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "Date=\"") {
		t.Fatalf("%v: responseEntryDateHeader missing %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "Dummy=\"dummy\"") {
		t.Fatalf("%v: responseEntryDummyHeader missing %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "body:         '") || strings.Contains(responseEntry.message, "body:         ''") {
		t.Fatalf("%v: responseEntryResponseBody %v", logPrefix, responseEntry.message)
	}
	if responseEntry.err != nil {
		t.Fatalf("%v: responseEntryErr %v", logPrefix, responseEntry.err)
	}
}

func TestPostWithBadRequestResponse(t *testing.T) {
	logPrefix := "TestPostWithBadRequestResponse"

	responseHeaders := map[string]string{
		"Content-Type": "application/json",
		"Dummy":        "dummy",
	}
	requestHeaders := map[string][]string{}

	listener, sl, comm, err := createTestEnvironment(
		"/v1/bad-request",
		createRecordRequest(http.StatusBadRequest, postWithBadRequestResponseJSON, responseHeaders, requestHeaders))
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}
	defer func(listener net.Listener, sl *stoppableListener, comm *Communicator) {
		_ = listener.Close()
		_ = sl.Close()
		_ = comm.Close()
	}(listener, sl, comm)

	logger := &testLogger{}
	comm.EnableLogging(logger)

	cardData := map[string]interface{}{
		"cardNumber":       "1234567890123456",
		"cardSecurityCode": "123",
		"expiryDate":       "122024",
	}

	request := map[string]interface{}{
		"cardData": cardData,
	}

	response := map[string]interface{}{}
	err = comm.Post("v1/bad-request", nil, nil, request, nil, &response)
	switch ce := err.(type) {
	case *commErrors.ResponseError:
		{
			if ce.StatusCode() != http.StatusBadRequest {
				t.Fatalf("%v: statusCode %v", logPrefix, ce.StatusCode())
			}
			if ce.Body() != postWithBadRequestResponseJSON {
				t.Fatalf("%v: responseBody %v", logPrefix, ce.Body())
			}

			break
		}
	default:
		{
			t.Fatalf("%v: %v", logPrefix, err)
		}
	}

	if len(logger.entries) != 2 {
		t.Fatalf("%v: loggerEntries %v", logPrefix, len(logger.entries))
	}

	requestEntry := logger.entries[0]
	if !strings.Contains(requestEntry.message, "method:       'POST'") {
		t.Fatalf("%v: requestEntryRequestMethod %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "uri:          '/v1/bad-request'") {
		t.Fatalf("%v: requestEntryRequestURL %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "Authorization=\"********\"") {
		t.Fatalf("%v: requestEntryAuthorizationHeader missing %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "Date=\"") {
		t.Fatalf("%v: requestEntryDateHeader missing %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "X-WL-ServerMetaInfo=\"") {
		t.Fatalf("%v: requestEntryDateHeader missing %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "body:         '") || strings.Contains(requestEntry.message, "body:         ''") {
		t.Fatalf("%v: requestEntryRequestBody %v", logPrefix, requestEntry.message)
	}
	if strings.Contains(requestEntry.message, "\"1234567890123456\"") || strings.Contains(requestEntry.message, "\"123\"") || strings.Contains(requestEntry.message, "\"122024\"") {
		t.Fatalf("%v: requestEntryRequestBodyObfuscation %v", logPrefix, requestEntry.message)
	}
	if requestEntry.err != nil {
		t.Fatalf("%v: requestEntryErr %v", logPrefix, requestEntry.err)
	}

	responseEntry := logger.entries[1]
	if !strings.Contains(responseEntry.message, "status-code:  '400'") {
		t.Fatalf("%v: responseEntryResponseStatusCode %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "content-type: 'application/json'") {
		t.Fatalf("%v: responseEntryResponseContentType %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "Date=\"") {
		t.Fatalf("%v: responseEntryDateHeader missing %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "Dummy=\"dummy\"") {
		t.Fatalf("%v: responseEntryDummyHeader missing %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "body:         '") || strings.Contains(responseEntry.message, "body:         ''") {
		t.Fatalf("%v: responseEntryResponseBody %v", logPrefix, responseEntry.message)
	}
	if responseEntry.err != nil {
		t.Fatalf("%v: responseEntryErr %v", logPrefix, responseEntry.err)
	}
}

func TestLoggingUnknownServerError(t *testing.T) {
	logPrefix := "TestLoggingUnknownServerError"

	responseHeaders := map[string]string{
		"Content-Type": "application/json",
		"Dummy":        "dummy",
	}
	requestHeaders := map[string][]string{}

	listener, sl, comm, err := createTestEnvironment(
		"/v1/get",
		createRecordRequest(http.StatusInternalServerError, unknownServerErrorJSON, responseHeaders, requestHeaders))
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}
	defer func(listener net.Listener, sl *stoppableListener, comm *Communicator) {
		_ = listener.Close()
		_ = sl.Close()
		_ = comm.Close()
	}(listener, sl, comm)

	logger := &testLogger{}
	comm.EnableLogging(logger)

	response := make(map[string]interface{})
	err = comm.Get("v1/get", nil, nil, nil, &response)
	switch ce := err.(type) {
	case *commErrors.ResponseError:
		{
			if ce.StatusCode() != http.StatusInternalServerError {
				t.Fatalf("%v: statusCode %v", logPrefix, ce.StatusCode())
			}
			if ce.Body() != unknownServerErrorJSON {
				t.Fatalf("%v: responseBody %v", logPrefix, ce.Body())
			}

			break
		}
	default:
		{
			t.Fatalf("%v: %v", logPrefix, err)
		}
	}

	if len(logger.entries) != 2 {
		t.Fatalf("%v: loggerEntries %v", logPrefix, len(logger.entries))
	}

	requestEntry := logger.entries[0]
	if !strings.Contains(requestEntry.message, "method:       'GET'") {
		t.Fatalf("%v: requestEntryRequestMethod %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "uri:          '/v1/get'") {
		t.Fatalf("%v: requestEntryRequestURL %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "Authorization=\"********\"") {
		t.Fatalf("%v: requestEntryAuthorizationHeader missing %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "Date=\"") {
		t.Fatalf("%v: requestEntryDateHeader missing %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "X-WL-ServerMetaInfo=\"") {
		t.Fatalf("%v: requestEntryDateHeader missing %v", logPrefix, requestEntry.message)
	}
	if requestEntry.err != nil {
		t.Fatalf("%v: requestEntryErr %v", logPrefix, requestEntry.err)
	}

	responseEntry := logger.entries[1]
	if !strings.Contains(responseEntry.message, "status-code:  '500'") {
		t.Fatalf("%v: responseEntryResponseStatusCode %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "content-type: 'application/json'") {
		t.Fatalf("%v: responseEntryResponseContentType %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "Date=\"") {
		t.Fatalf("%v: responseEntryDateHeader missing %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "Dummy=\"dummy\"") {
		t.Fatalf("%v: responseEntryDummyHeader missing %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "body:         '") || strings.Contains(responseEntry.message, "body:         ''") {
		t.Fatalf("%v: responseEntryResponseBody %v", logPrefix, responseEntry.message)
	}
	if responseEntry.err != nil {
		t.Fatalf("%v: responseEntryErr %v", logPrefix, responseEntry.err)
	}
}

func TestNonJson(t *testing.T) {
	logPrefix := "TestNonJson"

	responseHeaders := map[string]string{
		"Dummy": "dummy",
	}
	requestHeaders := map[string][]string{}

	listener, sl, comm, err := createTestEnvironment(
		"/v1/get",
		createRecordRequest(http.StatusNotFound, notFoundErrorHTML, responseHeaders, requestHeaders))
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}
	defer func(listener net.Listener, sl *stoppableListener, comm *Communicator) {
		_ = listener.Close()
		_ = sl.Close()
		_ = comm.Close()
	}(listener, sl, comm)

	logger := &testLogger{}
	comm.EnableLogging(logger)

	response := make(map[string]interface{})
	err = comm.Get("v1/get", nil, nil, nil, &response)
	switch err.(type) {
	case *commErrors.NotFoundError:
		{
			break
		}
	default:
		{
			t.Fatalf("%v: %v", logPrefix, err)
		}
	}

	if len(logger.entries) != 2 {
		t.Fatalf("%v: loggerEntries %v", logPrefix, len(logger.entries))
	}

	requestEntry := logger.entries[0]
	if !strings.Contains(requestEntry.message, "method:       'GET'") {
		t.Fatalf("%v: requestEntryRequestMethod %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "uri:          '/v1/get'") {
		t.Fatalf("%v: requestEntryRequestURL %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "Authorization=\"********\"") {
		t.Fatalf("%v: requestEntryAuthorizationHeader missing %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "Date=\"") {
		t.Fatalf("%v: requestEntryDateHeader missing %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "X-WL-ServerMetaInfo=\"") {
		t.Fatalf("%v: requestEntryDateHeader missing %v", logPrefix, requestEntry.message)
	}
	if requestEntry.err != nil {
		t.Fatalf("%v: requestEntryErr %v", logPrefix, requestEntry.err)
	}

	responseEntry := logger.entries[1]
	if !strings.Contains(responseEntry.message, "status-code:  '404'") {
		t.Fatalf("%v: responseEntryResponseStatusCode %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "content-type: 'text/plain") {
		t.Fatalf("%v: responseEntryResponseContentType %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "Date=\"") {
		t.Fatalf("%v: responseEntryDateHeader missing %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "Dummy=\"dummy\"") {
		t.Fatalf("%v: responseEntryDummyHeader missing %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "body:         '") || strings.Contains(responseEntry.message, "body:         ''") {
		t.Fatalf("%v: responseEntryResponseBody %v", logPrefix, responseEntry.message)
	}
	if responseEntry.err != nil {
		t.Fatalf("%v: responseEntryErr %v", logPrefix, responseEntry.err)
	}
}

func TestReadTimeout(t *testing.T) {
	logPrefix := "TestReadTimeout"

	responseHeaders := map[string]string{
		"Content-Type": "application/json",
		"Dummy":        "dummy",
	}
	requestHeaders := map[string][]string{}

	listener, sl, comm, err := createTimedTestEnvironment(
		"/v1/get",
		createDelayedRecordRequest(http.StatusOK, getJSON, responseHeaders, requestHeaders, 1*time.Second),
		1*time.Millisecond,
		10*time.Millisecond)
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}
	defer func(listener net.Listener, sl *stoppableListener, comm *Communicator) {
		_ = listener.Close()
		_ = sl.Close()
		_ = comm.Close()
	}(listener, sl, comm)

	logger := &testLogger{}
	comm.EnableLogging(logger)

	response := make(map[string]interface{})
	err = comm.Get("v1/get", nil, nil, nil, &response)
	switch ce := err.(type) {
	case *commErrors.CommunicationError:
		{
			internalError := ce.InternalError()

			if uErr, ok := internalError.(*url.Error); ok && uErr.Timeout() {
				break
			}

			t.Fatalf("%v: %v", logPrefix, internalError)
		}
	default:
		{
			t.Fatalf("%v: %v", logPrefix, err)
		}
	}

	if len(logger.entries) != 2 {
		t.Fatalf("%v: loggerEntries %v", logPrefix, len(logger.entries))
	}

	requestEntry := logger.entries[0]
	if !strings.Contains(requestEntry.message, "method:       'GET'") {
		t.Fatalf("%v: requestEntryRequestMethod %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "uri:          '/v1/get'") {
		t.Fatalf("%v: requestEntryRequestURL %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "Authorization=\"********\"") {
		t.Fatalf("%v: requestEntryAuthorizationHeader missing %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "Date=\"") {
		t.Fatalf("%v: requestEntryDateHeader missing %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "X-WL-ServerMetaInfo=\"") {
		t.Fatalf("%v: requestEntryDateHeader missing %v", logPrefix, requestEntry.message)
	}
	if requestEntry.err != nil {
		t.Fatalf("%v: requestEntryErr %v", logPrefix, requestEntry.err)
	}

	errorEntry := logger.entries[1]
	if errorEntry.err == nil {
		t.Fatalf("%v: responseEntryErr %v", logPrefix, errorEntry.err)
	}
}

func TestLogRequestOnly(t *testing.T) {
	logPrefix := "TestLogRequestOnly"

	responseHeaders := map[string]string{
		"Content-Type": "application/json",
		"Dummy":        "dummy",
	}
	requestHeaders := map[string][]string{}

	listener, sl, comm, mux, err := createEmptyTestEnvironment()
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}
	defer func(listener net.Listener, sl *stoppableListener, comm *Communicator) {
		_ = listener.Close()
		_ = sl.Close()
		_ = comm.Close()
	}(listener, sl, comm)

	mux.HandleFunc("/v1/get",
		createNonLoggedRecordRequest(http.StatusOK, getJSON, responseHeaders, requestHeaders, comm))

	logger := &testLogger{}
	comm.EnableLogging(logger)

	response := make(map[string]interface{})
	err = comm.Get("v1/get", nil, nil, nil, &response)
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}

	if len(logger.entries) != 1 {
		t.Fatalf("%v: loggerEntries %v", logPrefix, len(logger.entries))
	}

	requestEntry := logger.entries[0]
	if !strings.Contains(requestEntry.message, "method:       'GET'") {
		t.Fatalf("%v: requestEntryRequestMethod %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "uri:          '/v1/get'") {
		t.Fatalf("%v: requestEntryRequestURL %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "Authorization=\"********\"") {
		t.Fatalf("%v: requestEntryAuthorizationHeader missing %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "Date=\"") {
		t.Fatalf("%v: requestEntryDateHeader missing %v", logPrefix, requestEntry.message)
	}
	if !strings.Contains(requestEntry.message, "X-WL-ServerMetaInfo=\"") {
		t.Fatalf("%v: requestEntryDateHeader missing %v", logPrefix, requestEntry.message)
	}
	if requestEntry.err != nil {
		t.Fatalf("%v: requestEntryErr %v", logPrefix, requestEntry.err)
	}
}

func TestLogResponseOnly(t *testing.T) {
	logPrefix := "TestLogResponseOnly"

	responseHeaders := map[string]string{
		"Content-Type": "application/json",
		"Dummy":        "dummy",
	}
	requestHeaders := map[string][]string{}

	listener, sl, comm, mux, err := createEmptyTestEnvironment()
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}
	defer func(listener net.Listener, sl *stoppableListener, comm *Communicator) {
		_ = listener.Close()
		_ = sl.Close()
		_ = comm.Close()
	}(listener, sl, comm)

	logger := &testLogger{}

	mux.HandleFunc("/v1/get",
		createLoggedRecordRequest(http.StatusOK, getJSON, responseHeaders, requestHeaders, comm, logger))

	response := make(map[string]interface{})
	err = comm.Get("v1/get", nil, nil, nil, &response)
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}

	if len(logger.entries) != 1 {
		t.Fatalf("%v: loggerEntries %v", logPrefix, len(logger.entries))
	}

	responseEntry := logger.entries[0]
	if !strings.Contains(responseEntry.message, "status-code:  '200'") {
		t.Fatalf("%v: responseEntryResponseStatusCode %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "content-type: 'application/json'") {
		t.Fatalf("%v: responseEntryResponseContentType %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "body:         '") || strings.Contains(responseEntry.message, "body:         ''") {
		t.Fatalf("%v: responseEntryResponseBody %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "Date=\"") {
		t.Fatalf("%v: responseEntryDateHeader missing %v", logPrefix, responseEntry.message)
	}
	if !strings.Contains(responseEntry.message, "Dummy=\"dummy\"") {
		t.Fatalf("%v: responseEntryDummyHeader missing %v", logPrefix, responseEntry.message)
	}
	if responseEntry.err != nil {
		t.Fatalf("%v: responseEntryErr %v", logPrefix, responseEntry.err)
	}
}

func TestLogErrorOnly(t *testing.T) {
	logPrefix := "TestLogErrorOnly"

	responseHeaders := map[string]string{
		"Content-Type": "application/json",
		"Dummy":        "dummy",
	}
	requestHeaders := map[string][]string{}

	listener, sl, comm, mux, err := createEmptyTimedTestEnvironment(
		100*time.Millisecond,
		100*time.Millisecond)
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}
	defer func(listener net.Listener, sl *stoppableListener, comm *Communicator) {
		_ = listener.Close()
		_ = sl.Close()
		_ = comm.Close()
	}(listener, sl, comm)

	logger := &testLogger{}

	mux.HandleFunc("/v1/get",
		createLoggedDelayedRecordRequest(http.StatusOK, getJSON, responseHeaders, requestHeaders, 1*time.Second, comm, logger))

	response := make(map[string]interface{})
	err = comm.Get("v1/get", nil, nil, nil, &response)
	switch ce := err.(type) {
	case *commErrors.CommunicationError:
		{
			internalError := ce.InternalError()

			if uErr, ok := internalError.(*url.Error); ok && uErr.Timeout() {
				break
			}

			t.Fatalf("%v: %v", logPrefix, internalError)
		}
	default:
		{
			t.Fatalf("%v: %v", logPrefix, err)
		}
	}

	if len(logger.entries) != 1 {
		t.Fatalf("%v: loggerEntries %v", logPrefix, len(logger.entries))
	}

	errorEntry := logger.entries[0]
	if errorEntry.err == nil {
		t.Fatalf("%v: responseEntryErr %v", logPrefix, errorEntry.err)
	}
}

type stoppableListener struct {
	*net.TCPListener
	stop     chan int
	finished sync.WaitGroup
}

var errStopped = errors.New("listener stopped")

func (sl *stoppableListener) Accept() (net.Conn, error) {
	sl.finished.Add(1)
	defer sl.finished.Done()

	for {
		_ = sl.SetDeadline(time.Now().Add(time.Second))

		newConn, err := sl.TCPListener.Accept()

		select {
		case <-sl.stop:
			return nil, errStopped
		default:
		}

		if err != nil {
			netErr, ok := err.(net.Error)

			if ok && netErr.Timeout() && netErr.Temporary() {
				continue
			}
		}

		return newConn, err
	}
}

func (sl *stoppableListener) Stop() {
	close(sl.stop)
	sl.finished.Wait()
}

func newStoppableListener(l net.Listener) (*stoppableListener, error) {
	tcpL, ok := l.(*net.TCPListener)

	if !ok {
		return nil, errors.New("cannot wrap listener")
	}

	return &stoppableListener{tcpL, make(chan int), sync.WaitGroup{}}, nil
}

func mockServer(server *http.Server, listener net.Listener) (*stoppableListener, error) {
	ls, err := newStoppableListener(listener)
	if err != nil {
		return nil, err
	}

	go func() {
		_ = server.Serve(ls)
	}()

	return ls, nil
}

func createRecordRequest(statusCode int, body string, responseHeaders map[string]string, requestHeaders map[string][]string) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			requestHeaders[k] = v
		}

		for k, v := range responseHeaders {
			rw.Header()[k] = []string{v}
		}

		rw.WriteHeader(statusCode)

		_, _ = rw.Write([]byte(body))
	}
}

func createNonLoggedRecordRequest(statusCode int, body string, responseHeaders map[string]string, requestHeaders map[string][]string, comm *Communicator) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		comm.DisableLogging()

		for k, v := range r.Header {
			requestHeaders[k] = v
		}

		for k, v := range responseHeaders {
			rw.Header()[k] = []string{v}
		}

		rw.WriteHeader(statusCode)

		_, _ = rw.Write([]byte(body))
	}
}

func createLoggedRecordRequest(statusCode int, body string, responseHeaders map[string]string, requestHeaders map[string][]string, comm *Communicator, logger logging.CommunicatorLogger) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		comm.EnableLogging(logger)

		for k, v := range r.Header {
			requestHeaders[k] = v
		}

		for k, v := range responseHeaders {
			rw.Header()[k] = []string{v}
		}

		rw.WriteHeader(statusCode)

		_, _ = rw.Write([]byte(body))
	}
}

func createDelayedRecordRequest(statusCode int, body string, responseHeaders map[string]string, requestHeaders map[string][]string, delay time.Duration) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)

		for k, v := range r.Header {
			requestHeaders[k] = v
		}

		for k, v := range responseHeaders {
			rw.Header()[k] = []string{v}
		}

		rw.WriteHeader(statusCode)

		_, _ = rw.Write([]byte(body))
	}
}

func createLoggedDelayedRecordRequest(statusCode int, body string, responseHeaders map[string]string, requestHeaders map[string][]string, delay time.Duration, comm *Communicator, logger logging.CommunicatorLogger) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		comm.EnableLogging(logger)
		time.Sleep(delay)

		for k, v := range r.Header {
			requestHeaders[k] = v
		}

		for k, v := range responseHeaders {
			rw.Header()[k] = []string{v}
		}

		rw.WriteHeader(statusCode)

		_, _ = rw.Write([]byte(body))
	}
}

func createCommunicator(socketTimeout, connectTimeout time.Duration, port int) (*Communicator, error) {
	connection, err := NewDefaultConnection(socketTimeout, connectTimeout, 30*time.Second, 50*time.Second, 500, nil)
	if err != nil {
		return nil, err
	}

	authenticator := &testAuthenticator{}

	metadataProvider, err := NewMetadataProvider("Worldline")
	if err != nil {
		return nil, err
	}

	endPoint := &url.URL{
		Scheme: "http",
		Host:   "localhost:" + strconv.Itoa(port),
	}

	marshaller := json.DefaultMarshaller()

	return NewCommunicator(endPoint, connection, authenticator, metadataProvider, marshaller)
}

func createTestEnvironment(path string, handleFunc http.HandlerFunc) (net.Listener, *stoppableListener, *Communicator, error) {
	mux := http.NewServeMux()
	mux.Handle(path, handleFunc)

	httpServer := &http.Server{
		Handler: mux,
	}

	randomPort := (1 << 12) + rand.Intn(1<<15)
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(randomPort))
	if err != nil {
		return nil, nil, nil, err
	}

	sl, err := mockServer(httpServer, listener)
	if err != nil {
		return nil, nil, nil, err
	}

	comm, err := createCommunicator(50*time.Second, 50*time.Second, randomPort)
	if err != nil {
		return nil, nil, nil, err
	}

	return listener, sl, comm, nil
}

func createTimedTestEnvironment(path string, handleFunc http.HandlerFunc, socketTimeout, connectTimeout time.Duration) (net.Listener, *stoppableListener, *Communicator, error) {
	mux := http.NewServeMux()
	mux.Handle(path, handleFunc)

	httpServer := &http.Server{
		Handler: mux,
	}

	randomPort := (1 << 12) + rand.Intn(1<<15)
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(randomPort))
	if err != nil {
		return nil, nil, nil, err
	}

	sl, err := mockServer(httpServer, listener)
	if err != nil {
		return nil, nil, nil, err
	}

	comm, err := createCommunicator(socketTimeout, connectTimeout, randomPort)
	if err != nil {
		return nil, nil, nil, err
	}

	return listener, sl, comm, nil
}

func createEmptyTestEnvironment() (net.Listener, *stoppableListener, *Communicator, *http.ServeMux, error) {
	mux := http.NewServeMux()

	httpServer := &http.Server{
		Handler: mux,
	}

	randomPort := (1 << 12) + rand.Intn(1<<15)
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(randomPort))
	if err != nil {
		return nil, nil, nil, nil, err
	}

	sl, err := mockServer(httpServer, listener)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	comm, err := createCommunicator(50*time.Second, 50*time.Second, randomPort)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return listener, sl, comm, mux, nil
}

func createEmptyTimedTestEnvironment(socketTimeout, connectTimeout time.Duration) (net.Listener, *stoppableListener, *Communicator, *http.ServeMux, error) {
	mux := http.NewServeMux()

	httpServer := &http.Server{
		Handler: mux,
	}

	randomPort := (1 << 12) + rand.Intn(1<<15)
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(randomPort))
	if err != nil {
		return nil, nil, nil, nil, err
	}

	sl, err := mockServer(httpServer, listener)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	comm, err := createCommunicator(socketTimeout, connectTimeout, randomPort)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return listener, sl, comm, mux, nil
}

type testAuthenticator struct {
}

func (t *testAuthenticator) GetAuthorization(httpMethod string, resourceURI url.URL, requestHeaders []communication.Header) (string, error) {
	return "Bearer test", nil
}

type testParamRequest struct {
	params communication.RequestParams
}

func (t *testParamRequest) addRequestParam(name, value string) {
	param, _ := communication.NewRequestParam(name, value)
	t.params = append(t.params, *param)
}

func (t *testParamRequest) ToRequestParameters() communication.RequestParams {
	return t.params
}

type testLogger struct {
	entries []testLoggerEntry
}

func (t *testLogger) Log(message string) {
	t.entries = append(t.entries, testLoggerEntry{message, nil})
}

func (t *testLogger) LogError(message string, err error) {
	t.entries = append(t.entries, testLoggerEntry{message, err})
}

type testLoggerEntry struct {
	message string
	err     error
}
