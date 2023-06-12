package service

import (
	"github.com/go-kit/log"

	"github.com/AliasYermukanov/proxy-server/src/repository"
	"github.com/AliasYermukanov/proxy-server/src/rpc"
)

// Service entity
type service struct {
	externalRPC rpc.ExternalRPC
	logger      log.Logger
	store       repository.InmemStore
}

// NewService service constructor
func NewService(
	store repository.InmemStore,
	externalRPC rpc.ExternalRPC,
	logger log.Logger,
) ProxyService {
	return &service{
		store:       store,
		externalRPC: externalRPC,
		logger:      logger,
	}
}
