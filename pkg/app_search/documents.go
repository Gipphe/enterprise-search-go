package app_search

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Gipphe/enterprise-search-go/internal/base"
)

func (c *Client) GetDocuments(ctx context.Context, engineName string, params *GetDocumentsParams) (*GetDocumentsResponse, error) {
	var result GetDocumentsResponse
	err := c.DoRequest(ctx, base.RequestParams{
		Method: http.MethodGet,
		Path:   fmt.Sprintf("/api/as/v1/engines/%s/documents", engineName),
		Result: &result,
		Params: paramsToValues(params),
	})
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) IndexDocuments(ctx context.Context, engineName string, documents []map[string]interface{}) (*IndexDocumentsResponse, error) {
	var result IndexDocumentsResponse
	err := c.DoRequest(ctx, base.RequestParams{
		Method: http.MethodPost,
		Path:   fmt.Sprintf("/api/as/v1/engines/%s/documents", engineName),
		Body:   documents,
		Result: &result,
	})
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) DeleteDocuments(ctx context.Context, engineName string, documentIDs []string) (*DeleteDocumentsResponse, error) {
	var result DeleteDocumentsResponse
	err := c.DoRequest(ctx, base.RequestParams{
		Method: http.MethodDelete,
		Path:   fmt.Sprintf("/api/as/v1/engines/%s/documents", engineName),
		Body:   documentIDs,
		Result: &result,
	})
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) UpdateDocuments(ctx context.Context, engineName string, documents []map[string]interface{}) (*UpdateDocumentsResponse, error) {
	var result UpdateDocumentsResponse
	err := c.DoRequest(ctx, base.RequestParams{
		Method: http.MethodPatch,
		Path:   fmt.Sprintf("/api/as/v1/engines/%s/documents", engineName),
		Body:   documents,
		Result: &result,
	})
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) ListDocuments(ctx context.Context, engineName string, params *ListDocumentsParams) (*ListDocumentsResponse, error) {
	var result ListDocumentsResponse
	err := c.DoRequest(ctx, base.RequestParams{
		Method: http.MethodGet,
		Path:   fmt.Sprintf("/api/as/v1/engines/%s/documents/list", engineName),
		Result: &result,
		Params: paramsToValues(params),
	})
	if err != nil {
		return nil, err
	}
	return &result, nil
}
