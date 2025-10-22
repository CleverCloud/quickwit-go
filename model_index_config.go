package quickwit

type IndexConfig struct {
	Version          string          `json:"version" yaml:"version"`
	ID               string          `json:"index_id" yaml:"index_id"`
	URI              string          `json:"index_uri,omitempty" yaml:"index_uri,omitempty"`
	DocMapping       DocMapping      `json:"doc_mapping" yaml:"doc_mapping"`
	IndexingSettings Settings        `json:"indexing_settings" yaml:"indexing_settings"`
	SearchSettings   SearchSettings  `json:"search_settings" yaml:"search_settings"`
	Retention        *IndexRetention `json:"retention,omitempty" yaml:"retention,omitempty"`
}

type IndexRetention struct {
	// Duration after which splits are dropped, expressed in a human-readable way (1 day, 2 hours, a week, ...).
	Period string `json:"period,omitempty" yaml:"period,omitempty"`
	//Frequency at which the retention policy is evaluated and applied
	// expressed as a cron expression (0 0 * * * *) or human-readable form (hourly, daily, weekly, monthly, yearly).
	Schedule string `json:"schedule,omitempty" yaml:"retention,omitempty"`
}
