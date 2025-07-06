package client

import "encoding/json"

type ResponseWrapper struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}
