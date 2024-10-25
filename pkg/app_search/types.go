package app_search

type CreateEngineRequest = GetEngineResponse
type CreateEngineResponse = GetEngineResponse
type GetEngineResponse struct {
	Name                        string                       `json:"name"`
	Language                    *string                      `json:"language,omitempty"`
	Type                        *string                      `json:"type,omitempty"`
	SourceEngines               []string                     `json:"source_engines,omitempty"`
	DocumentCount               *int                         `json:"document_count,omitempty"`
	IndexCreateSettingsOverride *IndexCreateSettingsOverride `json:"index_create_settings_override,omitempty"`
}

type IndexCreateSettingsOverride struct {
	NumberOfShards *int `json:"number_of_shards,omitempty"`
}

type ListEnginesParams struct {
	Page *PageParams `json:"page,omitempty"`
}

type PageParams struct {
	Current *int `json:"current,omitempty"`
	Size    *int `json:"size,omitempty"`
}

type ListEnginesResponse struct {
	Meta    Meta     `json:"meta"`
	Results []Engine `json:"results"`
}

type Meta struct {
	Page PageInfo `json:"page"`
}

type PageInfo struct {
	Current      int `json:"current"`
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
	Size         int `json:"size"`
}

type Engine struct {
	Name     string  `json:"name"`
	Type     *string `json:"type,omitempty"`
	Language *string `json:"language,omitempty"`
}

type DeleteEngineResponse struct {
	Deleted bool `json:"deleted"`
}

type DeleteMetaEngineSourceResponse = GetEngineResponse
type AddMetaEngineSourceResponse = GetEngineResponse

type GetDocumentsParams struct {
	IDs []string `json:"ids,omitempty"`
}

type GetDocumentsResponse []map[string]interface{}

type IndexDocumentsResponse []struct {
	ID string `json:"id"`
}

type DeleteDocumentsResponse []struct {
	ID      string `json:"id"`
	Deleted bool   `json:"deleted"`
}

type UpdateDocumentsResponse = IndexDocumentsResponse

type ListDocumentsParams struct {
	Page *PageParams `json:"page,omitempty"`
}

type ListDocumentsResponse struct {
	Meta    Meta                     `json:"meta"`
	Results []map[string]interface{} `json:"results"`
}

type SearchRequest struct {
	Query        string                 `json:"query"`
	Page         *PageParams            `json:"page,omitempty"`
	SearchFields map[string]interface{} `json:"search_fields,omitempty"`
	ResultFields map[string]interface{} `json:"result_fields,omitempty"`
	Filters      map[string]interface{} `json:"filters,omitempty"`
	Facets       map[string]interface{} `json:"facets,omitempty"`
	Sort         interface{}            `json:"sort,omitempty"`
	Boosts       map[string]interface{} `json:"boosts,omitempty"`
	Analytics    map[string]interface{} `json:"analytics,omitempty"`
	GroupBy      map[string]interface{} `json:"group,omitempty"`
}

type SearchResponse struct {
	Meta    SearchMeta               `json:"meta"`
	Results []map[string]interface{} `json:"results"`
}

type SearchMeta struct {
	Page      PageInfo `json:"page"`
	Alerts    []string `json:"alerts"`
	Warnings  []string `json:"warnings"`
	Precision *float64 `json:"precision,omitempty"`
	Engine    Engine   `json:"engine"`
	RequestID *string  `json:"request_id,omitempty"`
}

type MultiSearchRequest struct {
	Queries []SearchRequest `json:"queries"`
}

type MultiSearchResponse []SearchResponse

type SearchExplainRequest = SearchRequest

type SearchExplainResponse struct {
	Meta        SearchMeta             `json:"meta"`
	QueryString string                 `json:"query_string"`
	QueryBody   map[string]interface{} `json:"query_body"`
}

type QuerySuggestionRequest struct {
	Query string                 `json:"query"`
	Types map[string]interface{} `json:"types,omitempty"`
	Size  *int                   `json:"size,omitempty"`
}

type QuerySuggestionResponse struct {
	Results map[string]interface{} `json:"results"`
	Meta    map[string]interface{} `json:"meta"`
}

type GetSchemaResponse map[string]string

type PutSchemaRequest map[string]string
type PutSchemaResponse map[string]string

type ListSynonymSetsParams struct {
	Page *PageParams `json:"page,omitempty"`
}

type ListSynonymSetsResponse struct {
	Meta    Meta         `json:"meta"`
	Results []SynonymSet `json:"results"`
}

type SynonymSet struct {
	ID       string   `json:"id,omitempty"`
	Synonyms []string `json:"synonyms"`
}

type CreateSynonymSetRequest = SynonymSet
type CreateSynonymSetResponse = SynonymSet
type GetSynonymSetResponse = SynonymSet
type PutSynonymSetRequest = SynonymSet
type PutSynonymSetResponse = SynonymSet

type DeleteSynonymSetResponse struct {
	Deleted bool `json:"deleted"`
}
