package enterprise_search

import (
	"context"
	"net/http"

	"github.com/Gipphe/enterprise-search-go/internal/base"
)

func (c *Client) GetHealth(ctx context.Context, params *GetHealthRequest) (*GetHealthResponse, error) {
	var result GetHealthResponse
	err := c.DoRequest(ctx, base.RequestParams{
		Method: http.MethodGet,
		Path:   "/api/ent/v1/internal/health",
		Result: &result,
	})
	return &result, err
}

func (c *Client) GetReadOnly(ctx context.Context, params *GetReadOnlyRequest) (*GetReadOnlyResponse, error) {
	var result GetReadOnlyResponse
	err := c.DoRequest(ctx, base.RequestParams{
		Method: http.MethodGet,
		Path:   "/api/ent/v1/internal/read_only_mode",
		Result: &result,
	})
	return &result, err
}

func (c *Client) PutReadOnly(ctx context.Context, params *PutReadOnlyRequest) (*PutReadOnlyResponse, error) {
	var result PutReadOnlyResponse
	err := c.DoRequest(ctx, base.RequestParams{
		Method: http.MethodPut,
		Path:   "/api/ent/v1/internal/read_only_mode",
		Body:   params.Body,
		Result: &result,
	})
	return &result, err
}

func (c *Client) GetStats(ctx context.Context, params *GetStatsRequest) (*GetStatsResponse, error) {
	var result GetStatsResponse
	err := c.DoRequest(ctx, base.RequestParams{
		Method: http.MethodGet,
		Path:   "/api/ent/v1/internal/stats",
		Params: paramsToValues(params),
		Result: &result,
	})
	return &result, err
}

func (c *Client) GetStorage(ctx context.Context, params *GetStorageRequest) (*GetStorageResponse, error) {
	var result GetStorageResponse
	err := c.DoRequest(ctx, base.RequestParams{
		Method: http.MethodGet,
		Path:   "/api/ent/v1/internal/storage",
		Result: &result,
	})
	return &result, err
}

func (c *Client) GetStaleStorage(ctx context.Context, params *GetStaleStorageRequest) (*GetStaleStorageResponse, error) {
	var result GetStaleStorageResponse
	err := c.DoRequest(ctx, base.RequestParams{
		Method: http.MethodGet,
		Path:   "/api/ent/v1/internal/storage/stale",
		Result: &result,
	})
	return &result, err
}

func (c *Client) DeleteStaleStorage(ctx context.Context, params *DeleteStaleStorageRequest) (*DeleteStaleStorageResponse, error) {
	var result DeleteStaleStorageResponse
	err := c.DoRequest(ctx, base.RequestParams{
		Method: http.MethodDelete,
		Path:   "/api/ent/v1/internal/storage/stale",
		Params: paramsToValues(params),
		Result: &result,
	})
	return &result, err
}

func (c *Client) GetVersion(ctx context.Context, params *GetVersionRequest) (*GetVersionResponse, error) {
	var result GetVersionResponse
	err := c.DoRequest(ctx, base.RequestParams{
		Method: http.MethodGet,
		Path:   "/api/ent/v1/internal/version",
		Result: &result,
	})
	return &result, err
}
