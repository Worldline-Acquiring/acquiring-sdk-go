package acquiringsdk

import (
	"bytes"
	encodingjson "encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/textproto"
	"net/url"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/Worldline-Acquiring/acquiring-sdk-go/communicator"
	"github.com/Worldline-Acquiring/acquiring-sdk-go/communicator/communication"
	commErrors "github.com/Worldline-Acquiring/acquiring-sdk-go/communicator/errors"
	"github.com/Worldline-Acquiring/acquiring-sdk-go/domain"
	"github.com/Worldline-Acquiring/acquiring-sdk-go/json"
)

// These tests verify the multipart/form-data requests that the SDK puts on the wire, using a local mock server.

const testAuthorization = "Bearer test"

func TestMultipartFormDataUpload(t *testing.T) {
	methods := []struct {
		method          string
		path            string
		send            func(comm *communicator.Communicator, path string, body interface{}, response interface{}) error
		sendWithHandler func(comm *communicator.Communicator, path string, body interface{}, handler communicator.BodyHandler) error
	}{
		{
			method: http.MethodPost,
			path:   "/post",
			send: func(comm *communicator.Communicator, path string, body interface{}, response interface{}) error {
				return comm.Post(path, nil, nil, body, nil, response)
			},
			sendWithHandler: func(comm *communicator.Communicator, path string, body interface{}, handler communicator.BodyHandler) error {
				return comm.PostWithHandler(path, nil, nil, body, nil, handler)
			},
		},
		{
			method: http.MethodPut,
			path:   "/put",
			send: func(comm *communicator.Communicator, path string, body interface{}, response interface{}) error {
				return comm.Put(path, nil, nil, body, nil, response)
			},
			sendWithHandler: func(comm *communicator.Communicator, path string, body interface{}, handler communicator.BodyHandler) error {
				return comm.PutWithHandler(path, nil, nil, body, nil, handler)
			},
		},
	}
	bodies := []struct {
		name   string
		toBody func(multipart *communication.MultipartFormDataObject) interface{}
	}{
		{"MultipartFormDataObject", func(multipart *communication.MultipartFormDataObject) interface{} { return *multipart }},
		{"MultipartFormDataObjectPointer", func(multipart *communication.MultipartFormDataObject) interface{} { return multipart }},
		{"MultipartFormDataRequest", func(multipart *communication.MultipartFormDataObject) interface{} {
			return &MultipartFormDataObjectWrapper{multipart}
		}},
	}

	for _, m := range methods {
		m := m
		for _, b := range bodies {
			b := b

			t.Run(m.method+"/"+b.name+"/WithResponse", func(t *testing.T) {
				server := newMockServer(t)
				comm := createCommunicator(t, server)
				multipart := newFileAndValueMultipart(t)

				response := echoResponse{}
				if err := m.send(comm, m.path, b.toBody(multipart), &response); err != nil {
					t.Fatal(err)
				}

				assertFileAndValueEchoed(t, response)
				assertFileAndValueRequest(t, server.takeRequest(t), m.method, m.path, multipart)
			})

			t.Run(m.method+"/"+b.name+"/WithBodyHandler", func(t *testing.T) {
				server := newMockServer(t)
				comm := createCommunicator(t, server)
				multipart := newFileAndValueMultipart(t)

				handled := false
				response := echoResponse{}
				err := m.sendWithHandler(comm, m.path, b.toBody(multipart), func(headers []communication.Header, reader io.Reader) error {
					handled = true
					return json.DefaultMarshaller().UnmarshalFromReader(reader, &response)
				})
				if err != nil {
					t.Fatal(err)
				}

				if !handled {
					t.Fatal("body handler was not called")
				}
				assertFileAndValueEchoed(t, response)
				assertFileAndValueRequest(t, server.takeRequest(t), m.method, m.path, multipart)
			})
		}
	}
}

func TestMultipartFormDataUploadTwoFiles(t *testing.T) {
	multipart := newMultipart(t)
	addFile(t, multipart, "firstFile", "first.txt", "text/plain", []byte("firstContent"))
	addFile(t, multipart, "secondFile", "second.txt", "text/plain", []byte("secondContent"))

	parts := postAndTakeParts(t, multipart)

	assertPartCount(t, parts, 2)
	assertFilePart(t, parts, "firstFile", "first.txt", "text/plain", []byte("firstContent"))
	assertFilePart(t, parts, "secondFile", "second.txt", "text/plain", []byte("secondContent"))
}

func TestMultipartFormDataUploadThreeFilesWithDifferentTypes(t *testing.T) {
	multipart := newMultipart(t)
	addFile(t, multipart, "textFile", "file.txt", "text/plain", []byte("text"))
	addFile(t, multipart, "jsonFile", "file.json", "application/json", []byte(`{"key": "json"}`))
	addFile(t, multipart, "xmlFile", "file.xml", "application/xml", []byte("<key>xml</key>"))

	parts := postAndTakeParts(t, multipart)

	assertPartCount(t, parts, 3)
	assertFilePart(t, parts, "textFile", "file.txt", "text/plain", []byte("text"))
	assertFilePart(t, parts, "jsonFile", "file.json", "application/json", []byte(`{"key": "json"}`))
	assertFilePart(t, parts, "xmlFile", "file.xml", "application/xml", []byte("<key>xml</key>"))
}

func TestMultipartFormDataUploadTwoValues(t *testing.T) {
	multipart := newMultipart(t)
	addValue(t, multipart, "firstKey", "firstValue")
	addValue(t, multipart, "secondKey", "secondValue")

	parts := postAndTakeParts(t, multipart)

	assertPartCount(t, parts, 2)
	assertValuePart(t, parts, "firstKey", "firstValue")
	assertValuePart(t, parts, "secondKey", "secondValue")
}

func TestMultipartFormDataUploadThreeValues(t *testing.T) {
	multipart := newMultipart(t)
	addValue(t, multipart, "name", "John")
	addValue(t, multipart, "age", "30")
	addValue(t, multipart, "city", "NYC")

	parts := postAndTakeParts(t, multipart)

	assertPartCount(t, parts, 3)
	assertValuePart(t, parts, "name", "John")
	assertValuePart(t, parts, "age", "30")
	assertValuePart(t, parts, "city", "NYC")
}

func TestMultipartFormDataUploadNonASCIIValue(t *testing.T) {
	multipart := newMultipart(t)
	addValue(t, multipart, "greeting", "Hello, Šđčćž €")

	parts := postAndTakeParts(t, multipart)

	assertPartCount(t, parts, 1)
	assertValuePart(t, parts, "greeting", "Hello, Šđčćž €")
}

func TestMultipartFormDataUploadSingleFileWithoutValues(t *testing.T) {
	multipart := newMultipart(t)
	addFile(t, multipart, "document", "doc.pdf", "application/pdf", []byte("doc content"))

	parts := postAndTakeParts(t, multipart)

	assertPartCount(t, parts, 1)
	assertFilePart(t, parts, "document", "doc.pdf", "application/pdf", []byte("doc content"))
}

func TestMultipartFormDataUploadSingleValueWithoutFiles(t *testing.T) {
	multipart := newMultipart(t)
	addValue(t, multipart, "message", "Hello")

	parts := postAndTakeParts(t, multipart)

	assertPartCount(t, parts, 1)
	assertValuePart(t, parts, "message", "Hello")
}

func TestMultipartFormDataUploadContentTypes(t *testing.T) {
	binaryContent := make([]byte, 64*1024)
	for i := range binaryContent {
		binaryContent[i] = byte(i)
	}

	tests := []struct {
		name          string
		parameterName string
		fileName      string
		contentType   string
		content       []byte
	}{
		{"pdf", "pdf", "document.pdf", "application/pdf", []byte("pdf content")},
		{"image", "image", "photo.jpg", "image/jpeg", []byte("image content")},
		{"json", "data", "data.json", "application/json", []byte(`{"key": "value"}`)},
		{"binary", "binary", "data.bin", "application/octet-stream", binaryContent},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			multipart := newMultipart(t)
			addFile(t, multipart, tt.parameterName, tt.fileName, tt.contentType, tt.content)

			parts := postAndTakeParts(t, multipart)

			assertPartCount(t, parts, 1)
			assertFilePart(t, parts, tt.parameterName, tt.fileName, tt.contentType, tt.content)
		})
	}
}

func TestMultipartFormDataUploadErrorResponse(t *testing.T) {
	errorBody := `{"status": 400, "title": "Bad Request"}`
	server := newMockServer(t)
	server.enqueue(mockResponse{statusCode: http.StatusBadRequest, contentType: "application/json", body: errorBody})
	comm := createCommunicator(t, server)

	multipart := newMultipart(t)
	addFile(t, multipart, "file", "file.txt", "text/plain", []byte("file-content"))

	err := comm.Post("/post", nil, nil, multipart, nil, &echoResponse{})

	var responseError *commErrors.ResponseError
	if !errors.As(err, &responseError) {
		t.Fatalf("expected a ResponseError, got %T: %v", err, err)
	}
	if responseError.StatusCode() != http.StatusBadRequest {
		t.Errorf("expected status code %d, got %d", http.StatusBadRequest, responseError.StatusCode())
	}
	if responseError.Body() != errorBody {
		t.Errorf("expected body %q, got %q", errorBody, responseError.Body())
	}

	parts := assertMultipartRequest(t, server.takeRequest(t), http.MethodPost, "/post", multipart)
	assertPartCount(t, parts, 1)
	assertFilePart(t, parts, "file", "file.txt", "text/plain", []byte("file-content"))
}

func createCommunicator(t *testing.T, server *mockServer) *communicator.Communicator {
	t.Helper()

	apiEndpoint, err := url.Parse(server.server.URL)
	if err != nil {
		t.Fatal(err)
	}

	connection, err := communicator.NewDefaultConnection(10*time.Second, 5*time.Second, 30*time.Second, 5*time.Second, 10, nil)
	if err != nil {
		t.Fatal(err)
	}

	metadataProvider, err := communicator.NewMetadataProvider("Worldline")
	if err != nil {
		t.Fatal(err)
	}

	comm, err := communicator.NewCommunicator(apiEndpoint, connection, &testAuthenticator{}, metadataProvider, json.DefaultMarshaller())
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		_ = comm.Close()
	})
	return comm
}

func postAndTakeParts(t *testing.T, multipart *communication.MultipartFormDataObject) map[string]recordedPart {
	t.Helper()

	server := newMockServer(t)
	comm := createCommunicator(t, server)

	response := echoResponse{}
	if err := comm.Post("/post", nil, nil, multipart, nil, &response); err != nil {
		t.Fatal(err)
	}

	return assertMultipartRequest(t, server.takeRequest(t), http.MethodPost, "/post", multipart)
}

func newMultipart(t *testing.T) *communication.MultipartFormDataObject {
	t.Helper()

	multipart, err := communication.NewMultipartFormDataObject()
	if err != nil {
		t.Fatal(err)
	}
	return multipart
}

func addFile(t *testing.T, multipart *communication.MultipartFormDataObject, parameterName, fileName, contentType string, content []byte) {
	t.Helper()

	file, err := domain.NewUploadableFile(fileName, bytes.NewReader(content), contentType)
	if err != nil {
		t.Fatal(err)
	}
	if err := multipart.AddFile(parameterName, *file); err != nil {
		t.Fatal(err)
	}
}

func addValue(t *testing.T, multipart *communication.MultipartFormDataObject, parameterName, value string) {
	t.Helper()

	if err := multipart.AddValue(parameterName, value); err != nil {
		t.Fatal(err)
	}
}

func newFileAndValueMultipart(t *testing.T) *communication.MultipartFormDataObject {
	t.Helper()

	multipart := newMultipart(t)
	addFile(t, multipart, "file", "file.txt", "text/plain", []byte("file-content"))
	addValue(t, multipart, "value", "Hello World")
	return multipart
}

func assertFileAndValueEchoed(t *testing.T, response echoResponse) {
	t.Helper()

	if len(response.Files) != 1 || response.Files["file"] != "file-content" {
		t.Errorf("expected echoed files {file: file-content}, got %v", response.Files)
	}
	if len(response.Form) != 1 || response.Form["value"] != "Hello World" {
		t.Errorf("expected echoed values {value: Hello World}, got %v", response.Form)
	}
}

func assertFileAndValueRequest(t *testing.T, request *recordedRequest, expectedMethod, expectedPath string, multipart *communication.MultipartFormDataObject) {
	t.Helper()

	parts := assertMultipartRequest(t, request, expectedMethod, expectedPath, multipart)
	assertPartCount(t, parts, 2)
	assertFilePart(t, parts, "file", "file.txt", "text/plain", []byte("file-content"))
	assertValuePart(t, parts, "value", "Hello World")
}

func assertMultipartRequest(t *testing.T, request *recordedRequest, expectedMethod, expectedPath string, multipart *communication.MultipartFormDataObject) map[string]recordedPart {
	t.Helper()

	if request.method != expectedMethod {
		t.Errorf("expected method %s, got %s", expectedMethod, request.method)
	}
	if request.path != expectedPath {
		t.Errorf("expected path %s, got %s", expectedPath, request.path)
	}
	if authorization := request.header.Get("Authorization"); authorization != testAuthorization {
		t.Errorf("expected Authorization %q, got %q", testAuthorization, authorization)
	}
	if contentType := request.header.Get("Content-Type"); contentType != multipart.GetContentType() {
		t.Errorf("expected Content-Type %q, got %q", multipart.GetContentType(), contentType)
	}
	if request.parseErr != nil {
		t.Fatalf("invalid multipart body: %v", request.parseErr)
	}
	return request.parts
}

func assertPartCount(t *testing.T, parts map[string]recordedPart, expected int) {
	t.Helper()

	if len(parts) != expected {
		t.Errorf("expected %d parts, got %d", expected, len(parts))
	}
}

func assertFilePart(t *testing.T, parts map[string]recordedPart, name, fileName, contentType string, content []byte) {
	t.Helper()

	part, ok := parts[name]
	if !ok {
		t.Errorf("missing part: %s", name)
		return
	}
	if part.fileName != fileName {
		t.Errorf("part %s: expected filename %q, got %q", name, fileName, part.fileName)
	}
	if partContentType := part.header.Get("Content-Type"); partContentType != contentType {
		t.Errorf("part %s: expected Content-Type %q, got %q", name, contentType, partContentType)
	}
	if !bytes.Equal(part.content, content) {
		t.Errorf("part %s: content differs; expected %d bytes, got %d bytes", name, len(content), len(part.content))
	}
}

func assertValuePart(t *testing.T, parts map[string]recordedPart, name, value string) {
	t.Helper()

	part, ok := parts[name]
	if !ok {
		t.Errorf("missing part: %s", name)
		return
	}
	if part.fileName != "" {
		t.Errorf("value part %s should not have a filename, got %q", name, part.fileName)
	}
	if string(part.content) != value {
		t.Errorf("part %s: expected value %q, got %q", name, value, string(part.content))
	}
}

// mock server

// mockServer is a local HTTP server that records every request it receives.
type mockServer struct {
	server    *httptest.Server
	requests  chan *recordedRequest
	mutex     sync.Mutex
	responses []mockResponse
}

type mockResponse struct {
	statusCode  int
	contentType string
	body        string
}

type recordedRequest struct {
	method   string
	path     string
	header   http.Header
	parts    map[string]recordedPart
	parseErr error
}

type recordedPart struct {
	fileName string
	header   textproto.MIMEHeader
	content  []byte
}

type echoResponse struct {
	Method string            `json:"method"`
	Path   string            `json:"path"`
	Form   map[string]string `json:"form"`
	Files  map[string]string `json:"files"`
}

func newMockServer(t *testing.T) *mockServer {
	t.Helper()

	s := &mockServer{requests: make(chan *recordedRequest, 10)}
	s.server = httptest.NewServer(http.HandlerFunc(s.handle))
	t.Cleanup(s.server.Close)
	return s
}

func (s *mockServer) enqueue(response mockResponse) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.responses = append(s.responses, response)
}

func (s *mockServer) nextResponse() (mockResponse, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if len(s.responses) == 0 {
		return mockResponse{}, false
	}
	response := s.responses[0]
	s.responses = s.responses[1:]
	return response, true
}

func (s *mockServer) takeRequest(t *testing.T) *recordedRequest {
	t.Helper()

	select {
	case request := <-s.requests:
		return request
	case <-time.After(5 * time.Second):
		t.Fatal("no request was received by the mock server")
		return nil
	}
}

func (s *mockServer) handle(w http.ResponseWriter, r *http.Request) {
	request := &recordedRequest{
		method: r.Method,
		path:   r.URL.RequestURI(),
		header: r.Header.Clone(),
	}
	if strings.HasPrefix(r.Header.Get("Content-Type"), "multipart/form-data") {
		request.parts, request.parseErr = parseMultipart(r)
	} else {
		_, _ = io.Copy(io.Discard, r.Body)
	}
	// record the request before responding, so it can be taken as soon as the client received the response
	s.requests <- request

	response, ok := s.nextResponse()
	if !ok {
		response = createEchoResponse(request)
	}
	w.Header().Set("Content-Type", response.contentType)
	w.WriteHeader(response.statusCode)
	_, _ = io.WriteString(w, response.body)
}

func parseMultipart(r *http.Request) (map[string]recordedPart, error) {
	reader, err := r.MultipartReader()
	if err != nil {
		return nil, err
	}

	parts := map[string]recordedPart{}
	for {
		part, err := reader.NextRawPart()
		if err == io.EOF {
			return parts, nil
		}
		if err != nil {
			return nil, err
		}

		content, err := io.ReadAll(part)
		if err != nil {
			return nil, err
		}
		name := part.FormName()
		if _, exists := parts[name]; exists {
			return nil, fmt.Errorf("duplicate part: %s", name)
		}
		parts[name] = recordedPart{fileName: part.FileName(), header: part.Header, content: content}
	}
}

func createEchoResponse(request *recordedRequest) mockResponse {
	echo := echoResponse{Method: request.method, Path: request.path}
	if request.parts != nil {
		echo.Form = map[string]string{}
		echo.Files = map[string]string{}
		for name, part := range request.parts {
			if part.fileName == "" {
				echo.Form[name] = string(part.content)
			} else {
				echo.Files[name] = string(part.content)
			}
		}
	}

	body, err := encodingjson.Marshal(echo)
	if err != nil {
		return mockResponse{statusCode: http.StatusInternalServerError, contentType: "text/plain", body: err.Error()}
	}
	return mockResponse{statusCode: http.StatusOK, contentType: "application/json", body: string(body)}
}

type MultipartFormDataObjectWrapper struct {
	Multipart *communication.MultipartFormDataObject
}

func (w *MultipartFormDataObjectWrapper) ToMultipartFormDataObject() *communication.MultipartFormDataObject {
	return w.Multipart
}

type testAuthenticator struct {
}

func (t *testAuthenticator) GetAuthorization(httpMethod string, resourceURI url.URL, requestHeaders []communication.Header) (string, error) {
	return testAuthorization, nil
}
