package json

import (
	"encoding/json"
	"io"
)

// defaultMarshaller represents the default implementation of the JSON Marshaller
type defaultMarshaller struct {
}

func (m defaultMarshaller) Marshal(v interface{}) (string, error) {
	dataBytes, err := json.Marshal(v)
	data := string(dataBytes)

	return data, err
}

func (m defaultMarshaller) Unmarshal(data string, v interface{}) error {
	if len(data) < 1 {
		return nil
	}

	return json.Unmarshal([]byte(data), v)
}

func (m defaultMarshaller) UnmarshalFromReader(reader io.Reader, v interface{}) error {
	decoder := json.NewDecoder(reader)

	err := decoder.Decode(v)
	if err != io.EOF {
		return err
	}

	return nil
}

var marshaller = &defaultMarshaller{}

// DefaultMarshaller returns the default implementation of the JSON Marshaller
func DefaultMarshaller() Marshaller {
	return marshaller
}
