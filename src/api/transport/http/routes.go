package http

import (
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"github.com/AliasYermukanov/proxy-server/src/api/middleware"
)

// Initialize HTTP routes
func initializeRoutes(endpoints *middleware.Endpoints, options []kithttp.ServerOption) *mux.Router {

	proxySend := kithttp.NewServer(
		endpoints.ProxySend,
		proxySendRequestDecoder,
		encodeResponse,
		options...,
	)

	getRequestByID := kithttp.NewServer(
		endpoints.GetRequestByID,
		GetRequestByIDRequestDecoder,
		encodeResponse,
		options...,
	)

	router := mux.NewRouter()

	router.Path("/proxy-server/v1/proxy/send").
		Methods("POST").
		Handler(proxySend)

	router.Path("/proxy-server/v1/proxy/{id}").
		Methods("GET").
		Handler(getRequestByID)

	return router
}
