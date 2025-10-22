package quickwit

type SearchResponse struct {
	Hits              any `json:"hits"`
	NumHits           int `json:"num_hits"`
	ElapsedTimeMicros int `json:"elapsed_time_micros"`
}
