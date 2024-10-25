package workplace_search

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Gipphe/enterprise-search-go/internal/base"
)

func (c *Client) DeleteDocumentsByQuery(ctx context.Context, params *DeleteDocumentsByQueryRequest) (*DeleteDocumentsByQueryResponse, error) {
	var result DeleteDocumentsByQueryResponse
	err := c.DoRequest(ctx, base.RequestParams{
		Method: http.MethodDelete,
		Path:   fmt.Sprintf("/api/ws/v1/sources/%s/documents", params.ContentSourceID),
		Body:   params.Body,
		Result: &result,
	})
	return &result, err
}

func (c *Client) ListDocuments(ctx context.Context, params *ListDocumentsRequest) (*ListDocumentsResponse, error) {
	var result ListDocumentsResponse
	err := c.DoRequest(ctx, base.RequestParams{
		Method: http.MethodPost,
		Path:   fmt.Sprintf("/api/ws/v1/sources/%s/documents", params.ContentSourceID),
		Body:   params.Body,
		Result: &result,
	})
	return &result, err
}

func (c *Client) IndexDocuments(ctx context.Context, params *IndexDocumentsRequest) (*IndexDocumentsResponse, error) {
	var result IndexDocumentsResponse
	err := c.DoRequest(ctx, base.RequestParams{
		Method: http.MethodPost,
		Path:   fmt.Sprintf("/api/ws/v1/sources/%s/documents/bulk_create", params.ContentSourceID),
		Body:   params.Documents,
		Result: &result,
	})
	return &result, err
}

func (c *Client) DeleteDocuments(ctx context.Context, params *DeleteDocumentsRequest) (*DeleteDocumentsResponse, error) {
	var result DeleteDocumentsResponse
	err := c.DoRequest(ctx, base.RequestParams{
		Method: http.MethodPost,
		Path:   fmt.Sprintf("/api/ws/v1/sources/%s/documents/bulk_destroy", params.ContentSourceID),
		Body:   params.DocumentIDs,
		Result: &result,
	})
	return &result, err
}

func (c *Client) GetDocument(ctx context.Context, params *GetDocumentRequest) (*GetDocumentResponse, error) {
	var result GetDocumentResponse
	err := c.DoRequest(ctx, base.RequestParams{
		Method: http.MethodGet,
		Path:   fmt.Sprintf("/api/ws/v1/sources/%s/documents/%s", params.ContentSourceID, params.DocumentID),
		Result: &result,
	})
	return &result, err
}
