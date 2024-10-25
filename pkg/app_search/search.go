package app_search

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Gipphe/enterprise-search-go/internal/base"
)

func (c *Client) Search(ctx context.Context, engineName string, query *SearchRequest) (*SearchResponse, error) {
	var result SearchResponse
	err := c.DoRequest(ctx, base.RequestParams{
		Method: http.MethodPost,
		Path:   fmt.Sprintf("/api/as/v1/engines/%s/search", engineName),
		Body:   query,
		Result: &result,
	})
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) MultiSearch(ctx context.Context, engineName string, queries *MultiSearchRequest) (*MultiSearchResponse, error) {
	var result MultiSearchResponse
	err := c.DoRequest(ctx, base.RequestParams{
		Method: http.MethodPost,
		Path:   fmt.Sprintf("/api/as/v1/engines/%s/multi_search", engineName),
		Body:   queries,
		Result: &result,
	})
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) SearchExplain(ctx context.Context, engineName string, query *SearchExplainRequest) (*SearchExplainResponse, error) {
	var result SearchExplainResponse
	err := c.DoRequest(ctx, base.RequestParams{
		Method: http.MethodPost,
		Path:   fmt.Sprintf("/api/as/v1/engines/%s/search_explain", engineName),
		Body:   query,
		Result: &result,
	})
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) QuerySuggestion(ctx context.Context, engineName string, query *QuerySuggestionRequest) (*QuerySuggestionResponse, error) {
	var result QuerySuggestionResponse
	err := c.DoRequest(ctx, base.RequestParams{
		Method: http.MethodPost,
		Path:   fmt.Sprintf("/api/as/v1/engines/%s/query_suggestion", engineName),
		Body:   query,
		Result: &result,
	})
	if err != nil {
		return nil, err
	}
	return &result, nil
}
