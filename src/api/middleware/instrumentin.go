package middleware

import (
	"context"
	"time"

	"github.com/go-kit/kit/metrics"

	"github.com/AliasYermukanov/proxy-server/src/api/service"
	"github.com/AliasYermukanov/proxy-server/src/api/transport"
)

// Instrumenting middleware entity
type instrumentingMiddleware struct {
	next           service.ProxyService
	requestCount   metrics.Counter
	requestError   metrics.Counter
	requestLatency metrics.Histogram
}

// Instrumenting middleware private method
func (im *instrumentingMiddleware) instrumenting(begin time.Time, method string, err error) {
	im.requestCount.With("method", method).Add(1)
	if err != nil {
		im.requestError.With("method", method).Add(1)
	}
	im.requestLatency.With("method", method).Observe(time.Since(begin).Seconds())
}

func NewInstrumentingMiddleware(counter, counterErr metrics.Counter, latency metrics.Histogram) Middleware {
	return func(next service.ProxyService) service.ProxyService {
		return &instrumentingMiddleware{
			next:           next,
			requestCount:   counter,
			requestError:   counterErr,
			requestLatency: latency,
		}
	}
}

func (im *instrumentingMiddleware) ProxySend(ctx context.Context, req *transport.ProxySendRequest) (_ *transport.ProxySendResponse, err error) {
	defer im.instrumenting(time.Now(), "ProxySend", err)
	return im.next.ProxySend(ctx, req)
}

func (im *instrumentingMiddleware) GetRequestByID(ctx context.Context, req *transport.GetRequestByIDRequest) (_ *transport.GetRequestByIDResponse, err error) {
	defer im.instrumenting(time.Now(), "GetRequestByID", err)
	return im.next.GetRequestByID(ctx, req)
}
