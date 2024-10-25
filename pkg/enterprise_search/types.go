package enterprise_search

type GetHealthRequest struct{}

type GetHealthResponse struct {
	Name        string `json:"name"`
	ClusterUUID string `json:"cluster_uuid"`
	Version     struct {
		Number    string `json:"number"`
		BuildHash string `json:"build_hash"`
		BuildDate string `json:"build_date"`
	} `json:"version"`
	JVM struct {
		PID         int   `json:"pid"`
		Uptime      int64 `json:"uptime"`
		MemoryUsage struct {
			HeapInit                       int64 `json:"heap_init"`
			HeapUsed                       int64 `json:"heap_used"`
			HeapCommitted                  int64 `json:"heap_committed"`
			HeapMax                        int64 `json:"heap_max"`
			NonHeapInit                    int64 `json:"non_heap_init"`
			NonHeapCommitted               int64 `json:"non_heap_committed"`
			ObjectPendingFinalizationCount int   `json:"object_pending_finalization_count"`
		} `json:"memory_usage"`
		MemoryPools []string `json:"memory_pools"`
		Threads     struct {
			ThreadCount             int   `json:"thread_count"`
			PeakThreadCount         int   `json:"peak_thread_count"`
			TotalStartedThreadCount int64 `json:"total_started_thread_count"`
			DaemonThreadCount       int   `json:"daemon_thread_count"`
		} `json:"threads"`
		VMVersion string `json:"vm_version"`
		VMVendor  string `json:"vm_vendor"`
		VMName    string `json:"vm_name"`
		GC        struct {
			CollectionCount   int64 `json:"collection_count"`
			CollectionTime    int64 `json:"collection_time"`
			GarbageCollectors map[string]struct {
				CollectionCount *int64 `json:"collection_count,omitempty"`
				CollectionTime  *int64 `json:"collection_time,omitempty"`
			} `json:"garbage_collectors"`
		} `json:"gc"`
	} `json:"jvm"`
	Filebeat struct {
		PID                     *int `json:"pid,omitempty"`
		Alive                   bool `json:"alive"`
		RestartCount            *int `json:"restart_count,omitempty"`
		SecondsSinceLastRestart *int `json:"seconds_since_last_restart,omitempty"`
	} `json:"filebeat"`
	Metricbeat *struct {
		PID                     *int `json:"pid,omitempty"`
		Alive                   bool `json:"alive"`
		RestartCount            *int `json:"restart_count,omitempty"`
		SecondsSinceLastRestart *int `json:"seconds_since_last_restart,omitempty"`
	} `json:"metricbeat,omitempty"`
	System struct {
		JavaVersion  string `json:"java_version"`
		JrubyVersion string `json:"jruby_version"`
		OSName       string `json:"os_name"`
		OSVersion    string `json:"os_version"`
	} `json:"system"`
	EsqueuesMe *map[string]interface{} `json:"esqueues_me,omitempty"`
	Crawler    *struct {
		Running bool `json:"running"`
		Workers struct {
			PoolSize  int `json:"pool_size"`
			Active    int `json:"active"`
			Available int `json:"available"`
		} `json:"workers"`
	} `json:"crawler,omitempty"`
}

type GetReadOnlyRequest struct{}

type GetReadOnlyResponse struct {
	Enabled bool `json:"enabled"`
}

type PutReadOnlyRequest struct {
	Body struct {
		Enabled bool `json:"enabled"`
	} `json:"body"`
}

type PutReadOnlyResponse struct {
	Enabled bool `json:"enabled"`
}

type GetStatsRequest struct {
	Include *[]string `json:"include,omitempty"`
}

type GetStatsResponse struct {
	ClusterUUID string `json:"cluster_uuid"`
	HTTP        struct {
		Connections struct {
			Current int `json:"current"`
			Max     int `json:"max"`
			Total   int `json:"total"`
		} `json:"connections"`
		RequestDurationMS struct {
			Max    float64 `json:"max"`
			Mean   float64 `json:"mean"`
			StdDev float64 `json:"std_dev"`
		} `json:"request_duration_ms"`
		NetworkBytes struct {
			ReceivedTotal int64   `json:"received_total"`
			ReceivedRate  float64 `json:"received_rate"`
			SentTotal     int64   `json:"sent_total"`
			SentRate      float64 `json:"sent_rate"`
		} `json:"network_bytes"`
		Responses struct {
			XX1 int `json:"1xx"`
			XX2 int `json:"2xx"`
			XX3 int `json:"3xx"`
			XX4 int `json:"4xx"`
			XX5 int `json:"5xx"`
		} `json:"responses"`
	} `json:"http"`
	App struct {
		PID     int                    `json:"pid"`
		Start   string                 `json:"start"`
		End     string                 `json:"end"`
		Metrics map[string]interface{} `json:"metrics"`
	} `json:"app"`
	Queues     map[string]interface{} `json:"queues"`
	Connectors struct {
		Alive bool `json:"alive"`
		Pool  struct {
			ExtractWorkerPool struct {
				Running        bool `json:"running"`
				QueueDepth     int  `json:"queue_depth"`
				Size           int  `json:"size"`
				Busy           int  `json:"busy"`
				Idle           int  `json:"idle"`
				TotalScheduled int  `json:"total_scheduled"`
				TotalCompleted int  `json:"total_completed"`
			} `json:"extract_worker_pool"`
			SubextractWorkerPool struct {
				Running        bool `json:"running"`
				QueueDepth     int  `json:"queue_depth"`
				Size           int  `json:"size"`
				Busy           int  `json:"busy"`
				Idle           int  `json:"idle"`
				TotalScheduled int  `json:"total_scheduled"`
				TotalCompleted int  `json:"total_completed"`
			} `json:"subextract_worker_pool"`
			PublishWorkerPool struct {
				Running        bool `json:"running"`
				QueueDepth     int  `json:"queue_depth"`
				Size           int  `json:"size"`
				Busy           int  `json:"busy"`
				Idle           int  `json:"idle"`
				TotalScheduled int  `json:"total_scheduled"`
				TotalCompleted int  `json:"total_completed"`
			} `json:"publish_worker_pool"`
		} `json:"pool"`
		JobStore struct {
			Waiting  int `json:"waiting"`
			Working  int `json:"working"`
			JobTypes struct {
				Full        int `json:"full"`
				Incremental int `json:"incremental"`
				Delete      int `json:"delete"`
				Permissions int `json:"permissions"`
			} `json:"job_types"`
		} `json:"job_store"`
	} `json:"connectors"`
	Crawler *struct {
		Global struct {
			CrawlRequests struct {
				Pending    int `json:"pending"`
				Active     int `json:"active"`
				Successful int `json:"successful"`
				Failed     int `json:"failed"`
			} `json:"crawl_requests"`
		} `json:"global"`
		Node struct {
			ActiveThreads int `json:"active_threads"`
			PagesVisited  int `json:"pages_visited"`
			URLsAllowed   int `json:"urls_allowed"`
			QueueSize     struct {
				Primary int `json:"primary"`
				Purge   int `json:"purge"`
			} `json:"queue_size"`
			URLsDenied  map[string]int `json:"urls_denied"`
			StatusCodes map[string]int `json:"status_codes"`
			Workers     struct {
				PoolSize  int `json:"pool_size"`
				Active    int `json:"active"`
				Available int `json:"available"`
			} `json:"workers"`
		} `json:"node"`
	} `json:"crawler,omitempty"`
	ProductUsage *struct {
		AppSearch struct {
			TotalEngines int `json:"total_engines"`
		} `json:"app_search"`
		WorkplaceSearch struct {
			TotalOrgSources        int `json:"total_org_sources"`
			TotalPrivateSources    int `json:"total_private_sources"`
			TotalQueriesLast30Days int `json:"total_queries_last_30_days"`
		} `json:"workplace_search"`
		EnterpriseSearch struct {
			TotalSearchIndices int `json:"total_search_indices"`
		} `json:"enterprise_search"`
	} `json:"product_usage,omitempty"`
}

type GetStorageRequest struct{}

type GetStorageResponse struct {
	Indices []struct {
		Name        string `json:"name"`
		SizeInBytes int64  `json:"size_in_bytes"`
	} `json:"indices"`
	Summary struct {
		IndexCount  int   `json:"index_count"`
		SizeInBytes int64 `json:"size_in_bytes"`
	} `json:"summary"`
}

type GetStaleStorageRequest struct{}

type GetStaleStorageResponse struct {
	Indices []struct {
		Name        string `json:"name"`
		SizeInBytes int64  `json:"size_in_bytes"`
	} `json:"indices"`
	Summary struct {
		IndexCount  int   `json:"index_count"`
		SizeInBytes int64 `json:"size_in_bytes"`
	} `json:"summary"`
}

type DeleteStaleStorageRequest struct {
	Force *bool `json:"force,omitempty"`
}

type DeleteStaleStorageResponse struct {
	Indices    []string `json:"indices"`
	IndexCount int      `json:"index_count"`
}

type GetVersionRequest struct{}

type GetVersionResponse struct {
	Number    string `json:"number"`
	BuildHash string `json:"build_hash"`
	BuildDate string `json:"build_date"`
}
