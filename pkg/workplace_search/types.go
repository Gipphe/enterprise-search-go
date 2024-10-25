package workplace_search

import "time"

type CreateContentSourceRequest struct {
	Body CreateContentSourceBody `json:"body"`
}

type CreateContentSourceBody struct {
	Name                     string               `json:"name"`
	Schema                   map[string]string    `json:"schema,omitempty"`
	Display                  *DisplayDetails      `json:"display,omitempty"`
	IsSearchable             *bool                `json:"is_searchable,omitempty"`
	Indexing                 *IndexingDetails     `json:"indexing,omitempty"`
	Facets                   *FacetCustomizations `json:"facets,omitempty"`
	AutomaticQueryRefinement *AutoQueryRefinement `json:"automatic_query_refinement,omitempty"`
}

type DisplayDetails struct {
	TitleField       string        `json:"title_field"`
	URLField         string        `json:"url_field"`
	Color            *string       `json:"color,omitempty"`
	DescriptionField *string       `json:"description_field,omitempty"`
	SubtitleField    *string       `json:"subtitle_field,omitempty"`
	TypeField        *string       `json:"type_field,omitempty"`
	MediaTypeField   *string       `json:"media_type_field,omitempty"`
	CreatedByField   *string       `json:"created_by_field,omitempty"`
	UpdatedByField   *string       `json:"updated_by_field,omitempty"`
	DetailFields     []DetailField `json:"detail_fields,omitempty"`
}

type DetailField struct {
	Label     string `json:"label"`
	FieldName string `json:"field_name"`
}

type IndexingDetails struct {
	Enabled       *bool             `json:"enabled,omitempty"`
	Features      *IndexingFeatures `json:"features,omitempty"`
	DefaultAction string            `json:"default_action,omitempty"`
	Rules         []IndexingRule    `json:"rules,omitempty"`
	Schedule      *IndexingSchedule `json:"schedule,omitempty"`
}

type IndexingFeatures struct {
	Thumbnails        ThumbnailFeature         `json:"thumbnails"`
	ContentExtraction ContentExtractionFeature `json:"content_extraction"`
}

type ThumbnailFeature struct {
	Enabled bool `json:"enabled"`
}

type ContentExtractionFeature struct {
	Enabled bool `json:"enabled"`
}

type IndexingRule struct {
	Exclude    *string        `json:"exclude,omitempty"`
	Include    *string        `json:"include,omitempty"`
	FilterType *string        `json:"filter_type,omitempty"`
	Fields     []FieldMapping `json:"fields,omitempty"`
}

type FieldMapping struct {
	Remote string `json:"remote"`
	Target string `json:"target"`
}

type IndexingSchedule struct {
	Full           *string         `json:"full,omitempty"`
	Incremental    *string         `json:"incremental,omitempty"`
	Delete         *string         `json:"delete,omitempty"`
	Permissions    *string         `json:"permissions,omitempty"`
	BlockedWindows []BlockedWindow `json:"blocked_windows,omitempty"`
}

type BlockedWindow struct {
	JobType string  `json:"job_type"`
	Day     string  `json:"day"`
	Start   *string `json:"start,omitempty"`
	End     *string `json:"end,omitempty"`
}

type FacetCustomizations struct {
	Overrides []FacetOverride `json:"overrides"`
}

type FacetOverride struct {
	Field       string  `json:"field"`
	Enabled     bool    `json:"enabled"`
	DisplayName *string `json:"display_name,omitempty"`
	Transform   *string `json:"transform,omitempty"`
}

type AutoQueryRefinement struct {
	Overrides []AutoQueryRefinementOverride `json:"overrides"`
}

type AutoQueryRefinementOverride struct {
	Field                 string   `json:"field"`
	Enabled               bool     `json:"enabled"`
	QueryExpansionPhrases []string `json:"query_expansion_phrases"`
	IsPerson              *bool    `json:"is_person,omitempty"`
}

type CreateContentSourceResponse struct {
	ID            string                   `json:"id"`
	ServiceType   string                   `json:"service_type"`
	CreatedAt     time.Time                `json:"created_at"`
	LastUpdatedAt time.Time                `json:"last_updated_at"`
	IsRemote      bool                     `json:"is_remote"`
	Details       []map[string]interface{} `json:"details"`
	Groups        []Group                  `json:"groups"`
	Name          string                   `json:"name"`
	Schema        map[string]string        `json:"schema,omitempty"`
	Display       *DisplayDetails          `json:"display,omitempty"`
	Context       string                   `json:"context"`
	IsSearchable  bool                     `json:"is_searchable"`
	Indexing      *IndexingDetails         `json:"indexing,omitempty"`
}

type Group struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Add other request and response types here...

type SearchRequest struct {
	Body SearchRequestBody `json:"body"`
}

type SearchRequestBody struct {
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

type PageParams struct {
	Current *int `json:"current,omitempty"`
	Size    *int `json:"size,omitempty"`
}

type SearchResponse struct {
	Meta    SearchMeta               `json:"meta"`
	Results []map[string]interface{} `json:"results"`
	Facets  interface{}              `json:"facets,omitempty"`
}

type SearchMeta struct {
	Page             PageInfo                     `json:"page"`
	RequestID        string                       `json:"request_id"`
	Warnings         []string                     `json:"warnings,omitempty"`
	QueryRefinements *QueryRefinements            `json:"query_refinements,omitempty"`
	ContentSources   map[string]ContentSourceInfo `json:"content_sources,omitempty"`
	Timeout          *int                         `json:"timeout,omitempty"`
}

type PageInfo struct {
	Current      int `json:"current"`
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
	Size         int `json:"size"`
}

type QueryRefinements struct {
	SubmittedQuery     string       `json:"submitted_query"`
	DecoratedQueryHTML string       `json:"decorated_query_html"`
	Refinements        []Refinement `json:"refinements"`
}

type Refinement struct {
	Term              *string                `json:"term,omitempty"`
	Position          []int                  `json:"position,omitempty"`
	TriggerType       *string                `json:"trigger_type,omitempty"`
	TriggerFilterType *string                `json:"trigger_filter_type,omitempty"`
	Filter            map[string]interface{} `json:"filter,omitempty"`
}

type ContentSourceInfo struct {
	Name        *string `json:"name,omitempty"`
	ServiceType *string `json:"service_type,omitempty"`
}

// Add these new types to the existing file

type ListContentSourcesRequest struct {
	Page *PageParams `json:"page,omitempty"`
}

type ListContentSourcesResponse struct {
	Meta    Meta                          `json:"meta"`
	Results []CreateContentSourceResponse `json:"results"`
}

type Meta struct {
	Page PageInfo `json:"page"`
}

type GetContentSourceRequest struct {
	ContentSourceID string `json:"content_source_id"`
}

type GetContentSourceResponse CreateContentSourceResponse

type PutContentSourceRequest struct {
	ContentSourceID string                  `json:"content_source_id"`
	Body            CreateContentSourceBody `json:"body"`
}

type PutContentSourceResponse CreateContentSourceResponse

type DeleteContentSourceRequest struct {
	ContentSourceID string `json:"content_source_id"`
}

type DeleteContentSourceResponse struct {
	Deleted bool `json:"deleted"`
}

type DeleteDocumentsByQueryRequest struct {
	ContentSourceID string                      `json:"content_source_id"`
	Body            *DeleteDocumentsByQueryBody `json:"body,omitempty"`
}

type DeleteDocumentsByQueryBody struct {
	Filters *DeleteDocumentsByQueryFilters `json:"filters,omitempty"`
}

type DeleteDocumentsByQueryFilters struct {
	LastUpdated *LastUpdatedFilter `json:"last_updated,omitempty"`
}

type LastUpdatedFilter struct {
	From string `json:"from,omitempty"`
	To   string `json:"to,omitempty"`
}

type DeleteDocumentsByQueryResponse struct {
	Total    int                      `json:"total"`
	Deleted  int                      `json:"deleted"`
	Failures []map[string]interface{} `json:"failures"`
}

type ListDocumentsRequest struct {
	ContentSourceID string             `json:"content_source_id"`
	Body            *ListDocumentsBody `json:"body,omitempty"`
}

type ListDocumentsBody struct {
	Page    *PageParams            `json:"page,omitempty"`
	Filters map[string]interface{} `json:"filters,omitempty"`
	Sort    interface{}            `json:"sort,omitempty"`
	Cursor  *string                `json:"cursor,omitempty"`
}

type ListDocumentsResponse struct {
	Meta    ListDocumentsMeta        `json:"meta"`
	Results []map[string]interface{} `json:"results"`
}

type ListDocumentsMeta struct {
	Page     PageInfo   `json:"page"`
	Cursor   CursorInfo `json:"cursor"`
	Warnings []string   `json:"warnings"`
}

type CursorInfo struct {
	Current *string `json:"current"`
	Next    *string `json:"next"`
}

type IndexDocumentsRequest struct {
	ContentSourceID string                   `json:"content_source_id"`
	Documents       []map[string]interface{} `json:"documents"`
}

type IndexDocumentsResponse struct {
	Results []IndexDocumentResult `json:"results"`
}

type IndexDocumentResult struct {
	ID     string   `json:"id"`
	Errors []string `json:"errors"`
}

type DeleteDocumentsRequest struct {
	ContentSourceID string   `json:"content_source_id"`
	DocumentIDs     []string `json:"document_ids"`
}

type DeleteDocumentsResponse struct {
	Results []DeleteDocumentResult `json:"results"`
}

type DeleteDocumentResult struct {
	ID      string `json:"id"`
	Success bool   `json:"success"`
}

type GetDocumentRequest struct {
	ContentSourceID string `json:"content_source_id"`
	DocumentID      string `json:"document_id"`
}

type GetDocumentResponse map[string]interface{}
