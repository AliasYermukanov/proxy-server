package rpc

import (
	"context"

	"github.com/AliasYermukanov/proxy-server/src/api/transport"
	"github.com/AliasYermukanov/proxy-server/src/domain"
)

type ProxyServiceRPC interface {
	SendRequest(ctx context.Context, req *transport.ProxySendRequest) (*domain.Response, error)
}
