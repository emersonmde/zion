package zion

type Room struct {
	AccountData struct {
		Events []Event
	} `json:"account_data"`
	Ephemeral struct {
		Events []Event
	}
	State struct {
		Events []Event
	}
	Timeline struct {
		Events    []Event
		Limited   bool
		PrevBatch string `json:"prev_batch"`
	}
	UnreadNotifications interface{}
}
