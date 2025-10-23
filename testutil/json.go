package testutil

import (
	"encoding/json"
	"testing"
)

// AssertValidJSON checks the given string is valid JSON.
func AssertValidJSON(tb testing.TB, s string) {
	tb.Helper()

	AssertValidJSONBytes(tb, []byte(s))
}

// AssertValidJSONBytes checks the given buffer is valid JSON.
func AssertValidJSONBytes(tb testing.TB, b []byte) {
	tb.Helper()

	if err := ValidJSONBytes(b); err != nil {
		tb.Errorf("invalid JSON: %v\njson:\n%s", err, b)
	}
}

// ValidJSONBytes checks b is valid JSON.
func ValidJSONBytes(b []byte) error {
	var v any
	return json.Unmarshal(b, &v)
}

// ValidJSON checks s is valid JSON.
func ValidJSON(s string) error {
	return ValidJSONBytes([]byte(s))
}
