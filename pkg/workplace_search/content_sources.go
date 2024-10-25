package workplace_search

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Gipphe/enterprise-search-go/internal/base"
)

func (c *Client) CreateContentSource(ctx context.Context, params *CreateContentSourceRequest) (*CreateContentSourceResponse, error) {
	var result CreateContentSourceResponse
	err := c.DoRequest(ctx, base.RequestParams{
		Method: http.MethodPost,
		Path:   "/api/ws/v1/sources",
		Body:   params.Body,
		Result: &result,
	})
	return &result, err
}

func (c *Client) ListContentSources(ctx context.Context, params *ListContentSourcesRequest) (*ListContentSourcesResponse, error) {
	var result ListContentSourcesResponse
	err := c.DoRequest(ctx, base.RequestParams{
		Method: http.MethodGet,
		Path:   "/api/ws/v1/sources",
		Params: paramsToValues(params),
		Result: &result,
	})
	return &result, err
}

func (c *Client) GetContentSource(ctx context.Context, params *GetContentSourceRequest) (*GetContentSourceResponse, error) {
	var result GetContentSourceResponse
	err := c.DoRequest(ctx, base.RequestParams{
		Method: http.MethodGet,
		Path:   fmt.Sprintf("/api/ws/v1/sources/%s", params.ContentSourceID),
		Result: &result,
	})
	return &result, err
}

func (c *Client) PutContentSource(ctx context.Context, params *PutContentSourceRequest) (*PutContentSourceResponse, error) {
	var result PutContentSourceResponse
	err := c.DoRequest(ctx, base.RequestParams{
		Method: http.MethodPut,
		Path:   fmt.Sprintf("/api/ws/v1/sources/%s", params.ContentSourceID),
		Body:   params.Body,
		Result: &result,
	})
	return &result, err
}

func (c *Client) DeleteContentSource(ctx context.Context, params *DeleteContentSourceRequest) (*DeleteContentSourceResponse, error) {
	var result DeleteContentSourceResponse
	err := c.DoRequest(ctx, base.RequestParams{
		Method: http.MethodDelete,
		Path:   fmt.Sprintf("/api/ws/v1/sources/%s", params.ContentSourceID),
		Result: &result,
	})
	return &result, err
}
