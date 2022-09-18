package middleware

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"

	"github.com/AliasYermukanov/proxy-server/src/api/service"
	"github.com/AliasYermukanov/proxy-server/src/api/transport"
)

// Logging middleware entity
type loggingMiddleware struct {
	next   service.ProxyService
	logger log.Logger
}

// NewLoggingMiddleware logging middleware constructor
func NewLoggingMiddleware(logger log.Logger) Middleware {
	return func(next service.ProxyService) service.ProxyService {
		return &loggingMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

// Logging middleware private method
func (lm *loggingMiddleware) logging(begin time.Time, method string, err error) {
	_ = lm.logger.Log("method", method, "took", time.Since(begin), "err", err)
}

func (lm *loggingMiddleware) ProxySend(ctx context.Context, req *transport.ProxySendRequest) (_ *transport.ProxySendResponse, err error) {
	defer lm.logging(time.Now(), "ProxySend", err)
	return lm.next.ProxySend(ctx, req)
}

func (lm *loggingMiddleware) GetRequestByID(ctx context.Context, req *transport.GetRequestByIDRequest) (_ *transport.GetRequestByIDResponse, err error) {
	defer lm.logging(time.Now(), "GetRequestByID", err)
	return lm.next.GetRequestByID(ctx, req)
}
