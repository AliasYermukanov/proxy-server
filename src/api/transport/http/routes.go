package http

import (
	"github.com/go-chi/chi/v5"
	kithttp "github.com/go-kit/kit/transport/http"

	"github.com/AliasYermukanov/proxy-server/src/api/middleware"
)

// Initialize HTTP routes
func initializeRoutes(endpoints *middleware.Endpoints, options []kithttp.ServerOption) *chi.Mux {

	proxySend := kithttp.NewServer(
		endpoints.ProxySend,
		proxySendRequestDecoder,
		encodeResponse,
		options...,
	)

	getDataByID := kithttp.NewServer(
		endpoints.GetDataByID,
		getDataByIDDecoder,
		encodeResponse,
		options...,
	)

	getAllData := kithttp.NewServer(
		endpoints.GetAllData,
		getAllDataDecoder,
		encodeResponse,
		options...,
	)

	router := chi.NewRouter()

	router.Method("POST","/proxy-server/v1/proxy/send", proxySend)
	router.Method("GET","/proxy-server/v1/proxy/get-by-id/{id}", getDataByID)
	router.Method("GET","/proxy-server/v1/proxy/all", getAllData)

	return router
}
