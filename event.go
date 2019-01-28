package zion

type Event struct {
	Content struct {
		Creator string
	}
	EventId        string  `json:"event_id"`
	OriginServerTs float64 `json:"origin_server_ts"`
	Sender         string
	StateKey       string `json:"state_key"`
	Type           string
	Unsigned       struct {
		Age float64
	}
}
