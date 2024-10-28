package app_search

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParamsToValues(t *testing.T) {
	tests := []struct {
		name     string
		params   interface{}
		expected url.Values
	}{
		{
			name: "ListEnginesParams with page parameters",
			params: ListEnginesParams{
				Page: &PageParams{
					Current: intPtr(2),
					Size:    intPtr(20),
				},
			},
			expected: url.Values{
				"page[current]": []string{"2"},
				"page[size]":    []string{"20"},
			},
		},
		{
			name: "GetDocumentsParams with multiple IDs",
			params: GetDocumentsParams{
				IDs: []string{"doc1", "doc2", "doc3"},
			},
			expected: url.Values{
				"ids": []string{"doc1", "doc2", "doc3"},
			},
		},
		{
			name: "ListSynonymSetsParams with only size",
			params: ListSynonymSetsParams{
				Page: &PageParams{
					Size: intPtr(50),
				},
			},
			expected: url.Values{
				"page[size]": []string{"50"},
			},
		},
		{
			name:     "nil params",
			params:   nil,
			expected: nil,
		},
		{
			name: "empty struct",
			params: ListEnginesParams{
				Page: nil,
			},
			expected: url.Values{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := paramsToValues(tt.params)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// Helper function to create int pointers
func intPtr(i int) *int {
	return &i
}
