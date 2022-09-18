package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

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

func GetRequestByIDRequestDecoder(_ context.Context, r *http.Request) (_ interface{}, err error) {
	var request transport.GetRequestByIDRequest
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, util.CommonError.SetDevMessage("mux vars error")
	}

	request.ID = id

	return request, nil
}
