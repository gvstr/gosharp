package json

import (
	"encoding/json"
	"io"
)

// Unmarshals a slice of bytes into type T. Use unmarshal if you have the data in memory already
func UnmarshalJson(data []byte, v interface{}) error {
	return json.Unmarshal(data, &v)
}

// Decodes a slice of bytes into type T. Use Decode if you are streaming data from an http response for example.
func DecodeJson(data io.Reader, v interface{}) error {
	return json.NewDecoder(data).Decode(&v)
}
