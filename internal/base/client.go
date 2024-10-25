package base

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type RequestParams struct {
	Method string
	Path   string
	Body   interface{}
	Result interface{}
	Params url.Values
}

type BaseClient struct {
	BaseURL    string
	Auth       Auth
	HTTPClient *http.Client
}

type Auth interface {
	ToHeader() (map[string]string, error)
}

// BasicAuth implements the Auth interface for basic authentication
type BasicAuth struct {
	Username string
	Password string
}

func (b BasicAuth) ToHeader() (map[string]string, error) {
	return map[string]string{
		"Authorization": "Basic " + basicAuth(b.Username, b.Password),
	}, nil
}

// ApiKeyAuth implements the Auth interface for API key authentication
type ApiKeyAuth struct {
	ApiKey string
}

func (a ApiKeyAuth) ToHeader() (map[string]string, error) {
	return map[string]string{
		"Authorization": "Bearer " + a.ApiKey,
	}, nil
}

func (c *BaseClient) DoRequest(ctx context.Context, params RequestParams) error {
	var bodyReader io.Reader
	if params.Body != nil {
		jsonBody, err := json.Marshal(params.Body)
		if err != nil {
			return err
		}
		bodyReader = bytes.NewReader(jsonBody)
	}

	u, err := url.Parse(c.BaseURL + params.Path)
	if err != nil {
		return err
	}

	if params.Params != nil {
		u.RawQuery = params.Params.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, params.Method, u.String(), bodyReader)
	if err != nil {
		return err
	}

	if params.Body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	authHeaders, err := c.Auth.ToHeader()
	if err != nil {
		return err
	}
	for k, v := range authHeaders {
		req.Header.Set(k, v)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return handleErrorResponse(resp)
	}

	if params.Result != nil {
		decoder := json.NewDecoder(resp.Body)
		if err := decoder.Decode(params.Result); err != nil {
			return err
		}
	}

	return nil
}

func handleErrorResponse(resp *http.Response) error {
	var errResp ErrorResponse
	if err := json.NewDecoder(resp.Body).Decode(&errResp); err != nil {
		return fmt.Errorf("error decoding error response: %v", err)
	}
	return &errResp
}

type ErrorResponse struct {
	Errors  []string          `json:"errors"`
	Message string            `json:"message"`
	Details map[string]string `json:"details,omitempty"`
}

func (e *ErrorResponse) Error() string {
	return e.Message
}

// Helper function to create base64 encoded string for basic auth
func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
