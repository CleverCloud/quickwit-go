package quickwit

type SplitsRes struct {
	Offset int     `json:"offset" yaml:"offset"`
	Size   int     `json:"size" yaml:"size"`
	Splits []Split `json:"splits" yaml:"splits"`
}

type Split struct {
	SplitState                  string        `json:"split_state"`
	UpdateTimestamp             int64         `json:"update_timestamp"`
	PublishTimestamp            int64         `json:"publish_timestamp"`
	Version                     string        `json:"version"`
	SplitID                     string        `json:"split_id"`
	IndexUID                    string        `json:"index_uid"`
	PartitionID                 int           `json:"partition_id"`
	SourceID                    string        `json:"source_id"`
	NodeID                      string        `json:"node_id"`
	NumDocs                     int           `json:"num_docs"`
	UncompressedDocsSizeInBytes int           `json:"uncompressed_docs_size_in_bytes"`
	TimeRange                   TimeRange     `json:"time_range"`
	CreateTimestamp             int64         `json:"create_timestamp"`
	Maturity                    Maturity      `json:"maturity"`
	Tags                        []string      `json:"tags"`
	FooterOffsets               FooterOffsets `json:"footer_offsets"`
	DeleteOpstamp               int           `json:"delete_opstamp"`
	NumMergeOps                 int           `json:"num_merge_ops"`
	DocMappingUID               string        `json:"doc_mapping_uid"`
}

type TimeRange struct {
	Start int64 `json:"start"`
	End   int64 `json:"end"`
}

type Maturity struct {
	Type                   string `json:"type"`
	MaturationPeriodMillis int64  `json:"maturation_period_millis"`
}

type FooterOffsets struct {
	Start int `json:"start"`
	End   int `json:"end"`
}
