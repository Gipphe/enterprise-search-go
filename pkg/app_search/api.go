package app_search

import (
	"net/url"

	"github.com/google/go-querystring/query"
)

// paramsToValues converts a struct to url.Values
func paramsToValues(params interface{}) url.Values {
	if params == nil {
		return nil
	}

	values, err := query.Values(params)
	if err != nil {
		return nil
	}
	return values
}
