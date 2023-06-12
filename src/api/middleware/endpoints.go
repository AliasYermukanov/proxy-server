package middleware

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/AliasYermukanov/proxy-server/src/api/service"
	"github.com/AliasYermukanov/proxy-server/src/api/transport"
)

// Endpoints middleware entity
type Endpoints struct {
	ProxySend   endpoint.Endpoint
	GetDataByID endpoint.Endpoint
	GetAllData  endpoint.Endpoint
}

// MakeEndpoints endpoints middleware constructor
func MakeEndpoints(s service.ProxyService) *Endpoints {
	return &Endpoints{
		ProxySend:   makeProxySendEndpoint(s),
		GetDataByID: makeGetDataByIDEndpoint(s),
		GetAllData:  makeGetAllDataEndpoint(s),
	}
}

func makeProxySendEndpoint(s service.ProxyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(transport.ProxySendRequest)
		return s.ProxySend(ctx, &req)
	}
}

func makeGetDataByIDEndpoint(s service.ProxyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(transport.GetDataByIDRequest)
		return s.GetDataByID(ctx, &req)
	}
}

func makeGetAllDataEndpoint(s service.ProxyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(transport.GetAllDataRequest)
		return s.GetAllData(ctx, &req)
	}
}
