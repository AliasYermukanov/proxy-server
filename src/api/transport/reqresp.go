package transport

import "github.com/AliasYermukanov/proxy-server/src/domain"

type (
	ProxySendRequest struct {
		domain.Request
	}

	ProxySendResponse struct {
		Response domain.Response
	}
)

type (
	GetDataByIDRequest struct {
		ID string `json:"id"`
	}
	GetDataByIDResponse struct {
		Request  domain.Request
		Response domain.Response
	}
)

type (
	GetAllDataRequest  struct{}
	GetAllDataResponse struct {
		Requests  []*domain.Request  `json:"requests"`
		Responses []*domain.Response `json:"responses"`
	}
)
