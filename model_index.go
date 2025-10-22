package quickwit

type Index struct {
	Version         string          `json:"version" yaml:"version"`
	UID             string          `json:"index_uid" yaml:"index_uid"`
	Config          IndexConfig     `json:"index_config" yaml:"index_config"`
	Checkpoint      IndexCheckpoint `json:"checkpoint" yaml:"checkpoint"`
	CreateTimestamp UnixTime        `json:"create_timestamp" yaml:"create_timestamp"`
	Sources         []Source        `json:"sources" yaml:"sources"`
}

type IndexCheckpoint map[string]any
