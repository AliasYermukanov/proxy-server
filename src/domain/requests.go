package domain

import "encoding/json"

type Request struct {
	ID      string                 `json:"id"`
	Method  *string                `json:"method"`
	URL     string                 `json:"url"`
	Headers map[string]string      `json:"headers"`
	Body    map[string]interface{} `json:"body"`
}

type Response struct {
	ID      string          `json:"id"`
	Status  int             `json:"status"`
	Headers Header          `json:"headers"`
	Body    json.RawMessage `json:"body"`
	Length  int64           `json:"length"`
}

type Header map[string][]string
