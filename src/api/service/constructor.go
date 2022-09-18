package service

import (
	"github.com/go-kit/kit/log"

	"github.com/AliasYermukanov/proxy-server/src/domain"
	"github.com/AliasYermukanov/proxy-server/src/rpc"
)

// Service entity
type service struct {
	externalRPC rpc.ExternalRPC
	logger      log.Logger
	container   domain.Container
}

// NewService service constructor
func NewService(
	container domain.Container,
	externalRPC rpc.ExternalRPC,
	logger log.Logger,
) ProxyService {
	return &service{
		container:   container,
		externalRPC: externalRPC,
		logger:      logger,
	}
}
