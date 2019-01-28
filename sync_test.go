package zion

import (
	"strings"
	"testing"
)

func TestState(t *testing.T) {
	s := `{
		"account_data":{
			"events":[]
		},
		"next_batch": "NEXTBATCHVALUE",
		"rooms": {
			"invite": {},
			"join": {
				"!asfLdzLnOdGRkdPZWu:localhost": {
					"account_data": {
						"events": []
					},
					"ephemeral": {
						"events": []
					},
					"state": {
						"events": []
					},
					"timeline": {
						"events": [
							{
								"content": {
									"creator": "@example:localhost"
								},
								"event_id": "$14606534990LhqHt:localhost",
								"origin_server_ts": 1460653499699,
								"sender": "@example:localhost",
								"state_key": "",
								"type": "m.room.create",
								"unsigned": {
									"age": 239192
								}
							},
							{
								"content": {
									"avatar_url": null,
									"displayname": null,
									"membership": "join"
								},
								"event_id": "$14606534991nsZKk:localhost",
								"membership": "join",
								"origin_server_ts": 1460653499727,
								"sender": "@example:localhost",
								"state_key": "@example:localhost",
								"type": "m.room.member",
								"unsigned": {
									"age": 239164
								}
							}
						],
						"limited": false,
						"prev_batch": "s9_3_0_1_1_1"
					},
					"unread_notifications": {}
				}
			}
		}
	}`
	res, err := DecodeState(strings.NewReader(s))
	if err != nil {
		t.Fatalf("unable to decode json: %s", err)
	}
	if res.NextBatch != "NEXTBATCHVALUE" {
		t.Fatalf("res.NextBatch expected NEXTBATCHVALUE got: %s", res.NextBatch)
	}
}