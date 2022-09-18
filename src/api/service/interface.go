package service

import (
	"context"

	"github.com/AliasYermukanov/proxy-server/src/api/transport"
)

// ProxyService Main proxy service interface
type ProxyService interface {
	ProxySend(ctx context.Context, req *transport.ProxySendRequest) (*transport.ProxySendResponse, error)
	GetRequestByID(ctx context.Context, req *transport.GetRequestByIDRequest) (*transport.GetRequestByIDResponse, error)
}
