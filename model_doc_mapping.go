package quickwit

type DocMapping struct {
	FieldMappings      []FieldMapping `json:"field_mappings" yaml:"field_mappings"`
	TagFields          []any          `json:"tag_fields,omitempty" yaml:"tag_fields,omitempty"`
	StoreSource        bool           `json:"store_source,omitempty" yaml:"store_source,omitempty"`
	IndexFieldPresence bool           `json:"index_field_presence,omitempty" yaml:"index_field_presence,omitempty"`
	TimestampField     string         `json:"timestamp_field,omitempty" yaml:"timestamp_field,omitempty"`
	Mode               string         `json:"mode" yaml:"mode"`
	MaxNumPartitions   int            `json:"max_num_partitions,omitempty" yaml:"max_num_partitions,omitempty"`
	Tokenizers         []any          `json:"tokenizers,omitempty" yaml:"tokenizers,omitempty"`
	// UID                string                  `json:"doc_mapping_uid,omitempty" yaml:"doc_mapping_uid,omitempty"`
	// StoreDocumentSize  bool           `json:"store_document_size" yaml:"store_document_size"`
}

type FieldMapping struct {
	Name          string   `json:"name" yaml:"name"`
	Type          string   `json:"type" yaml:"type"`
	Fast          any      `json:"fast,omitempty" yaml:"fast,omitempty"`
	FastPrecision string   `json:"fast_precision,omitempty" yaml:"fast_precision,omitempty"`
	Indexed       bool     `json:"indexed" yaml:"indexed"`
	InputFormat   string   `json:"input_format,omitempty" yaml:"input_format,omitempty"`   // TODO: accept arrays ?
	InputFormats  []string `json:"input_formats,omitempty" yaml:"input_formats,omitempty"` // TODO: accept arrays ?
	OutputFormat  string   `json:"output_format,omitempty" yaml:"output_format,omitempty"`
	Stored        bool     `json:"stored" yaml:"stored"`
	Fieldnorms    bool     `json:"fieldnorms,omitempty" yaml:"fieldnorms,omitempty"`
	Record        string   `json:"record,omitempty" yaml:"record,omitempty"`
	Tokenizer     string   `json:"tokenizer,omitempty" yaml:"tokenizer,omitempty"`
	Coerce        bool     `json:"coerce,omitempty" yaml:"coerce,omitempty"`
	ExpandDots    bool     `json:"expand_dots,omitempty" yaml:"expand_dots,omitempty"`
}
