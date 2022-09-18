package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/AliasYermukanov/proxy-server/src/util"
)

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
