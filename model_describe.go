package quickwit

// Describe index res body
type Describe struct {
	IndexID                       string `json:"index_id" yaml:"index_id"`
	IndexURI                      string `json:"index_uri" yaml:"index_uri"`
	NbPublishedSplits             int    `json:"num_published_splits" yaml:"num_published_splits"`
	SizePublishedSplits           int    `json:"size_published_splits" yaml:"size_published_splits"`
	NbPublishedDocs               int    `json:"num_published_docs" yaml:"num_published_docs"`
	SizePublishedDocsUncompressed int    `json:"size_published_docs_uncompressed" yaml:"size_published_docs_uncompressed"`
	TimestampFieldName            string `json:"timestamp_field_name" yaml:"timestamp_field_name"`
	MinTimestamp                  *int64 `json:"min_timestamp" yaml:"min_timestamp"`
	MaxTimestamp                  *int64 `json:"max_timestamp" yaml:"max_timestamp"`
}
