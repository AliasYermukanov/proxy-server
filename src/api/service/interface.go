package service

import (
	"context"

	"github.com/AliasYermukanov/proxy-server/src/api/transport"
)

// ProxyService Main proxy service interface
type ProxyService interface {
	ProxySend(ctx context.Context, req *transport.ProxySendRequest) (*transport.ProxySendResponse, error)
	GetDataByID(ctx context.Context, req *transport.GetDataByIDRequest) (*transport.GetDataByIDResponse, error)
	GetAllData(ctx context.Context, req *transport.GetAllDataRequest) (*transport.GetAllDataResponse, error)
}
