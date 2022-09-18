package domain

import "encoding/json"

type Container struct {
	RequestMap map[string]Request
}

type Request struct {
	ID      string          `json:"id"`
	Status  int             `json:"status"`
	Headers Header          `json:"headers"`
	Body    json.RawMessage `json:"body"`
	Length  int64           `json:"length"`
}

func (c *Container) Init() {
	c.RequestMap = make(map[string]Request)
}

type Header map[string][]string
