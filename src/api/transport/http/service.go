package http

import (
	"net/http"

	kittransport "github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"

	"github.com/AliasYermukanov/proxy-server/src/api/middleware"
)

// NewHTTPService HTTP service constructor
func NewHTTPService(svcEndpoints *middleware.Endpoints, options []kithttp.ServerOption, logger log.Logger) http.Handler {
	errorEncoder := kithttp.ServerErrorEncoder(
		EncodeError,
	)

	errorLogger := kithttp.ServerErrorHandler(
		kittransport.NewLogErrorHandler(logger),
	)

	options = append(options, errorEncoder, errorLogger)

	return initializeRoutes(svcEndpoints, options)
}
