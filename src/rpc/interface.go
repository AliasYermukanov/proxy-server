package rpc

// ExternalRPC External RPC Interface
type ExternalRPC interface {
	ProxyService() ProxyServiceRPC
}
