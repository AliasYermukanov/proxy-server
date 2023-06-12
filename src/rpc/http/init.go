package http

import (
	"net/http"
	"sync"

	"github.com/go-kit/log"

	"github.com/AliasYermukanov/proxy-server/src/rpc"
)

type RPC struct {
	mu sync.Mutex

	client *http.Client
	logger log.Logger

	ProxyServiceRPC rpc.ProxyServiceRPC
}

func NewRPC(client *http.Client, logger log.Logger) *RPC {
	return &RPC{
		client: client,
		logger: log.With(logger, "rpc", "http"),
	}
}

// InitHTTPClient Initialize HTTP client
func InitHTTPClient() *http.Client {
	httpClient := &http.Client{
		Transport: &http.Transport{},
	}

	return httpClient
}

func (r *RPC) ProxyService() rpc.ProxyServiceRPC {
	if r.ProxyServiceRPC != nil {
		return r.ProxyServiceRPC
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	r.ProxyServiceRPC = &ProxyServiceRPC{
		rpc: r,
	}

	return r.ProxyServiceRPC
}
