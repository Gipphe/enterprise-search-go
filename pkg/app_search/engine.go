package app_search

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Gipphe/enterprise-search-go/internal/base"
)

func (c *Client) ListEngines(ctx context.Context, params *ListEnginesParams) (*ListEnginesResponse, error) {
	var result ListEnginesResponse
	err := c.DoRequest(ctx, base.RequestParams{
		Method: http.MethodGet,
		Path:   "/api/as/v1/engines",
		Result: &result,
		Params: paramsToValues(params),
	})
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) CreateEngine(ctx context.Context, engine *CreateEngineRequest) (*CreateEngineResponse, error) {
	var result CreateEngineResponse
	err := c.DoRequest(ctx, base.RequestParams{
		Method: http.MethodPost,
		Path:   "/api/as/v1/engines",
		Body:   engine,
		Result: &result,
	})
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetEngine(ctx context.Context, engineName string) (*GetEngineResponse, error) {
	var result GetEngineResponse
	err := c.DoRequest(ctx, base.RequestParams{
		Method: http.MethodGet,
		Path:   fmt.Sprintf("/api/as/v1/engines/%s", engineName),
		Result: &result,
	})
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) DeleteEngine(ctx context.Context, engineName string) (*DeleteEngineResponse, error) {
	var result DeleteEngineResponse
	err := c.DoRequest(ctx, base.RequestParams{
		Method: http.MethodDelete,
		Path:   fmt.Sprintf("/api/as/v1/engines/%s", engineName),
		Result: &result,
	})
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) DeleteMetaEngineSource(ctx context.Context, engineName string, sourceEngines []string) (*DeleteMetaEngineSourceResponse, error) {
	var result DeleteMetaEngineSourceResponse
	err := c.DoRequest(ctx, base.RequestParams{
		Method: http.MethodDelete,
		Path:   fmt.Sprintf("/api/as/v1/engines/%s/source_engines", engineName),
		Body:   sourceEngines,
		Result: &result,
	})
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) AddMetaEngineSource(ctx context.Context, engineName string, sourceEngines []string) (*AddMetaEngineSourceResponse, error) {
	var result AddMetaEngineSourceResponse
	err := c.DoRequest(ctx, base.RequestParams{
		Method: http.MethodPost,
		Path:   fmt.Sprintf("/api/as/v1/engines/%s/source_engines", engineName),
		Body:   sourceEngines,
		Result: &result,
	})
	if err != nil {
		return nil, err
	}
	return &result, nil
}
