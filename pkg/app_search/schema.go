package app_search

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Gipphe/enterprise-search-go/internal/base"
)

func (c *Client) GetSchema(ctx context.Context, engineName string) (*GetSchemaResponse, error) {
	var result GetSchemaResponse
	err := c.DoRequest(ctx, base.RequestParams{
		Method: http.MethodGet,
		Path:   fmt.Sprintf("/api/as/v1/engines/%s/schema", engineName),
		Result: &result,
	})
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) UpdateSchema(ctx context.Context, engineName string, schema *PutSchemaRequest) (*PutSchemaResponse, error) {
	var result PutSchemaResponse
	err := c.DoRequest(ctx, base.RequestParams{
		Method: http.MethodPost,
		Path:   fmt.Sprintf("/api/as/v1/engines/%s/schema", engineName),
		Body:   schema,
		Result: &result,
	})
	if err != nil {
		return nil, err
	}
	return &result, nil
}
