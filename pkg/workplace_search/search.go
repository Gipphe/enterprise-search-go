package workplace_search

import (
	"context"
	"net/http"

	"github.com/Gipphe/enterprise-search-go/internal/base"
)

func (c *Client) Search(ctx context.Context, params *SearchRequest) (*SearchResponse, error) {
	var result SearchResponse
	err := c.DoRequest(ctx, base.RequestParams{
		Method: http.MethodPost,
		Path:   "/api/ws/v1/search",
		Body:   params.Body,
		Result: &result,
	})
	return &result, err
}
