package core

import (
	"encoding/json"
)

// JSON definition.
type JSON map[string]interface{}

// DecodeJSON returns a JSON object.
func DecodeJSON(body []byte) (JSON, error) {
	var resp JSON

	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// DecodeJSONArray returns an array of JSON objects.
func DecodeJSONArray(body []byte) ([]JSON, error) {
	var resp []JSON

	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// StringFromJSON extracts a string from a JSON.
func StringFromJSON(j JSON, key string) *string {
	if v, exists := j[key]; exists {
		str := v.(string)
		return &str
	}

	return nil
}
