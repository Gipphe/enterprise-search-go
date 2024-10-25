package app_search

import (
	"net/http"

	"github.com/Gipphe/enterprise-search-go/internal/base"
)

type Client struct {
	base.BaseClient
}

func NewClient(baseURL string, auth base.Auth, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	return &Client{
		BaseClient: base.BaseClient{
			BaseURL:    baseURL,
			Auth:       auth,
			HTTPClient: httpClient,
		},
	}
}
