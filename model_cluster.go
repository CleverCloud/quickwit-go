package quickwit

type Cluster struct {
	ClusterID  string `json:"cluster_id"`
	SelfNodeID string `json:"self_node_id"`
	ReadyNodes []struct {
		NodeID              string `json:"node_id"`
		GenerationID        int64  `json:"generation_id"`
		GossipAdvertiseAddr string `json:"gossip_advertise_addr"`
	} `json:"ready_nodes"`
	LiveNodes             []any `json:"live_nodes"`
	DeadNodes             []any `json:"dead_nodes"`
	ChitchatStateSnapshot struct {
		NodeStateSnapshots []struct {
			ChitchatID struct {
				NodeID              string `json:"node_id"`
				GenerationID        int64  `json:"generation_id"`
				GossipAdvertiseAddr string `json:"gossip_advertise_addr"`
			} `json:"chitchat_id"`
			NodeState struct {
				ChitchatID struct {
					NodeID              string `json:"node_id"`
					GenerationID        int64  `json:"generation_id"`
					GossipAdvertiseAddr string `json:"gossip_advertise_addr"`
				} `json:"chitchat_id"`
				Heartbeat int `json:"heartbeat"`
				KeyValues struct {
					EnabledServices struct {
						Value   string `json:"value"`
						Version int    `json:"version"`
						Status  string `json:"status"`
					} `json:"enabled_services"`
					GrpcAdvertiseAddr struct {
						Value   string `json:"value"`
						Version int    `json:"version"`
						Status  string `json:"status"`
					} `json:"grpc_advertise_addr"`
					IndexerTask01JKTRV3MFFPE8C0AVFJAW0JVK struct {
						Value   string `json:"value"`
						Version int    `json:"version"`
						Status  string `json:"status"`
					} `json:"indexer.task:01JKTRV3MFFPE8C0AVFJAW0JVK"`
					IndexerTask01JKTRV3MFNMWYHD016JT86KEF struct {
						Value   string `json:"value"`
						Version int    `json:"version"`
						Status  string `json:"status"`
					} `json:"indexer.task:01JKTRV3MFNMWYHD016JT86KEF"`
					IndexerTask01JKTRV3MFSYM43GAZE2C40ZC5 struct {
						Value   string `json:"value"`
						Version int    `json:"version"`
						Status  string `json:"status"`
					} `json:"indexer.task:01JKTRV3MFSYM43GAZE2C40ZC5"`
					IndexingCPUCapacity struct {
						Value   string `json:"value"`
						Version int    `json:"version"`
						Status  string `json:"status"`
					} `json:"indexing_cpu_capacity"`
					Readiness struct {
						Value   string `json:"value"`
						Version int    `json:"version"`
						Status  string `json:"status"`
					} `json:"readiness"`
				} `json:"key_values"`
				MaxVersion    int `json:"max_version"`
				LastGcVersion int `json:"last_gc_version"`
			} `json:"node_state"`
		} `json:"node_state_snapshots"`
		SeedAddrs []string `json:"seed_addrs"`
	} `json:"chitchat_state_snapshot"`
}
