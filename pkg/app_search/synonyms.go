package app_search

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Gipphe/enterprise-search-go/internal/base"
)

func (c *Client) ListSynonymSets(ctx context.Context, engineName string, params *ListSynonymSetsParams) (*ListSynonymSetsResponse, error) {
	var result ListSynonymSetsResponse
	err := c.DoRequest(ctx, base.RequestParams{
		Method: http.MethodGet,
		Path:   fmt.Sprintf("/api/as/v1/engines/%s/synonyms", engineName),
		Result: &result,
		Params: paramsToValues(params),
	})
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) CreateSynonymSet(ctx context.Context, engineName string, synonymSet *CreateSynonymSetRequest) (*CreateSynonymSetResponse, error) {
	var result CreateSynonymSetResponse
	err := c.DoRequest(ctx, base.RequestParams{
		Method: http.MethodPost,
		Path:   fmt.Sprintf("/api/as/v1/engines/%s/synonyms", engineName),
		Body:   synonymSet,
		Result: &result,
	})
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetSynonymSet(ctx context.Context, engineName, synonymSetID string) (*GetSynonymSetResponse, error) {
	var result GetSynonymSetResponse
	err := c.DoRequest(ctx, base.RequestParams{
		Method: http.MethodGet,
		Path:   fmt.Sprintf("/api/as/v1/engines/%s/synonyms/%s", engineName, synonymSetID),
		Result: &result,
	})
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) UpdateSynonymSet(ctx context.Context, engineName, synonymSetID string, synonymSet *PutSynonymSetRequest) (*PutSynonymSetResponse, error) {
	var result PutSynonymSetResponse
	err := c.DoRequest(ctx, base.RequestParams{
		Method: http.MethodPut,
		Path:   fmt.Sprintf("/api/as/v1/engines/%s/synonyms/%s", engineName, synonymSetID),
		Body:   synonymSet,
		Result: &result,
	})
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) DeleteSynonymSet(ctx context.Context, engineName, synonymSetID string) (*DeleteSynonymSetResponse, error) {
	var result DeleteSynonymSetResponse
	err := c.DoRequest(ctx, base.RequestParams{
		Method: http.MethodDelete,
		Path:   fmt.Sprintf("/api/as/v1/engines/%s/synonyms/%s", engineName, synonymSetID),
		Result: &result,
	})
	if err != nil {
		return nil, err
	}
	return &result, nil
}
