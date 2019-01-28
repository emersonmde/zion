package zion

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

type State struct {
	AccountData struct {
		Events []Event
	} `json:"account_data"`
	NextBatch string `json:"next_batch"`
	Rooms     struct {
		Invite map[string]Room
		Join   map[string]Room
	}
}

func DecodeState(r io.Reader) (*State, error) {
	d := json.NewDecoder(r)
	var res State
	if err := d.Decode(&res); err != nil {
		return nil, errors.New(fmt.Sprintf("error decoding state: %s", err))
	}
	return &res, nil
}
