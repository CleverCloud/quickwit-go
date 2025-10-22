package quickwit

type Source struct {
	ID           string `json:"source_id" yaml:"source_id"`
	Version      string `json:"version" yaml:"version"`
	NumPipelines int    `json:"num_pipelines" yaml:"num_pipelines"`
	Enabled      bool   `json:"enabled" yaml:"enabled"`
	SourceType   string `json:"source_type" yaml:"source_type"`
	InputFormat  string `json:"input_format" yaml:"input_format"`
}

type SourceConfig struct {
	Version       string         `json:"version"`
	ID            string         `json:"source_id"`
	Type          string         `json:"source_type"`
	PipelineCount int            `json:"num_pipelines,omitempty"`
	Params        map[string]any `json:"params"`
}

func NewPulsarSourceConfig(sourceID, endpoint, token, topic string) SourceConfig {
	src := SourceConfig{
		Version:       "0.9",
		ID:            sourceID,
		Type:          "pulsar",
		PipelineCount: 1,
		Params: map[string]any{
			"address":       endpoint,
			"topics":        []string{topic},
			"consumer_name": "quickwit",
		},
	}

	if token != "" {
		src.Params["authentication"] = map[string]any{
			"token": token,
		}
	}

	return src
}
