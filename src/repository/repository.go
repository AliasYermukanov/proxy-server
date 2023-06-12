package repository

import (
	"github.com/AliasYermukanov/proxy-server/src/domain"
)

// RequestsRepository In memory store interface
type RequestsRepository interface {
	StoreRequest(request *domain.Request)
	StoreResponse(response *domain.Response)
	GetAllRequests() []*domain.Request
	GetAllResponses() []*domain.Response
	GetRequestByID(ID string) (*domain.Request, error)
	GetResponseByID(ID string) (*domain.Response, error)
}
