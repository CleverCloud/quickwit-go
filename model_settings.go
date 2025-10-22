package quickwit

type Settings struct {
	CommitTimeoutSecs        int          `json:"commit_timeout_secs,omitempty" yaml:"commit_timeout_secs,omitempty"`
	DocstoreCompressionLevel int          `json:"docstore_compression_level,omitempty" yaml:"docstore_compression_level,omitempty"`
	DocstoreBlocksize        int          `json:"docstore_blocksize,omitempty" yaml:"docstore_blocksize,omitempty"`
	SplitNumDocsTarget       int          `json:"split_num_docs_target,omitempty" yaml:"split_num_docs_target,omitempty"`
	MergePolicy              *MergePolicy `json:"merge_policy,omitempty" yaml:"merge_policy,omitempty"`
	Resources                Resources    `json:"resources" yaml:"resources"`
}

type SearchSettings struct {
	DefaultSearchFields []string `json:"default_search_fields,omitempty" yaml:"default_search_fields,omitempty"`
}

type MergePolicy struct {
	Type             string `json:"type" yaml:"type"`
	MinLevelNumDocs  int    `json:"min_level_num_docs,omitempty" yaml:"min_level_num_docs,omitempty"`
	MergeFactor      int    `json:"merge_factor,omitempty" yaml:"merge_factor,omitempty"`
	MaxMergeFactor   int    `json:"max_merge_factor,omitempty" yaml:"max_merge_factor,omitempty"`
	MaturationPeriod string `json:"maturation_period,omitempty" yaml:"maturation_period,omitempty"`
}

type Resources struct {
	HeapSize string `json:"heap_size,omitempty" yaml:"heap_size,omitempty"`
}
