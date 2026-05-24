package renderer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatVal(t *testing.T) {
	// These test cases are inspired by CockroachDB's value formatting logic
	// to ensure consistent representation of NULLs, strings, and types.

	testCases := []struct {
		name                 string
		input                any
		showPrintableUnicode bool
		showNewLinesAndTabs  bool
		expected             string
	}{
		{
			name:     "NULL value",
			input:    nil,
			expected: "NULL",
		},
		{
			name:     "Simple ASCII string",
			input:    "hello world",
			expected: "hello world",
		},
		{
			name:                 "Unicode string (Graphic)",
			input:                "哈哈",
			showPrintableUnicode: true,
			expected:             "哈哈",
		},
		{
			name:                 "Unicode string (Non-Graphic fallback)",
			input:                "哈哈\x00",
			showPrintableUnicode: true,
			expected:             "哈哈\x00", // fmt.Sprint fallback
		},
		{
			name:                 "Newlines and Tabs (Allowed)",
			input:                "line1\nline2\ttab",
			showPrintableUnicode: true,
			showNewLinesAndTabs:  true,
			expected:             "line1\nline2\ttab",
		},
		{
			name:                 "Newlines and Tabs (Disallowed/Fallback)",
			input:                "line1\nline2",
			showPrintableUnicode: true,
			showNewLinesAndTabs:  false,
			expected:             "line1\nline2", // Fallback to fmt.Sprint if IndexFunc matches
		},
		{
			name:     "Integer value",
			input:    123,
			expected: "123",
		},
		{
			name:     "Float value",
			input:    0.3333333333333333,
			expected: "0.3333333333333333",
		},
		{
			name:     "Boolean value",
			input:    true,
			expected: "true",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := FormatVal(tc.input, tc.showPrintableUnicode, tc.showNewLinesAndTabs)
			assert.Equal(t, tc.expected, got)
		})
	}
}
