package http

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"

	"github.com/AliasYermukanov/proxy-server/src/api/transport"
	"github.com/AliasYermukanov/proxy-server/src/util"
)

func proxySendRequestDecoder(_ context.Context, r *http.Request) (interface{}, error) {
	var request transport.ProxySendRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, err
	}

	return request, nil
}

func getDataByIDDecoder(_ context.Context, r *http.Request) (_ interface{}, err error) {
	var request transport.GetDataByIDRequest
	id := chi.URLParam(r, "id")
	if id == "" {
		return nil, util.CommonError.SetDevMessage("mux vars error")
	}

	request.ID = id

	return request, nil
}

func getAllDataDecoder(_ context.Context, r *http.Request) (_ interface{}, err error) {
	return transport.GetAllDataRequest{}, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errored); ok && e.error() != nil {
		EncodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type errored interface {
	error() error
}

// EncodeError encode errors from business-logic
func EncodeError(_ context.Context, err error, w http.ResponseWriter) {
	switch err {
	case util.NoFound:
		w.WriteHeader(http.StatusNotFound)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_ = json.NewEncoder(w).Encode(err)
}