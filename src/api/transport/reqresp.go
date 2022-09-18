package transport

import "github.com/AliasYermukanov/proxy-server/src/domain"

type (
	ProxySendRequest struct {
		Method  *string                `json:"method"`
		URL     string                 `json:"url"`
		Headers map[string]string      `json:"headers"`
		Body    map[string]interface{} `json:"body"`
	}

	ProxySendResponse struct {
		Response domain.Request
	}
)

type (
	GetRequestByIDRequest struct {
		ID string `json:"id"`
	}
	GetRequestByIDResponse struct {
		Response domain.Request
	}
)
