package constructor

import (
	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"

	"github.com/AliasYermukanov/proxy-server/src/api/middleware"
	"github.com/AliasYermukanov/proxy-server/src/api/service"
	"github.com/AliasYermukanov/proxy-server/src/domain"
	"github.com/AliasYermukanov/proxy-server/src/rpc"
)

// NewProxyService Proxy service constructor
func NewProxyService(
	container domain.Container,
	externalRPC rpc.ExternalRPC,
	logger log.Logger) service.ProxyService {
	svc := service.NewService(container, externalRPC, logger)
	svc = middleware.NewLoggingMiddleware(logger)(svc)
	svc = middleware.NewInstrumentingMiddleware(
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "proxy_service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, []string{"method"}),
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "proxy_service",
			Name:      "error_count",
			Help:      "Number of error requests received.",
		}, []string{"method"}),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "proxy_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, []string{"method"}),
	)(svc)

	return svc
}
