package http

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/google/uuid"

	"github.com/AliasYermukanov/proxy-server/src/api/transport"
	"github.com/AliasYermukanov/proxy-server/src/domain"
)

type ProxyServiceRPC struct {
	rpc *RPC
}

func (r *ProxyServiceRPC) SendRequest(ctx context.Context, req *transport.ProxySendRequest) (*domain.Request, error) {
	var response domain.Request

	jsonData, err := json.Marshal(req.Body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequestWithContext(ctx, *req.Method, req.URL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	for key, value := range req.Headers {
		request.Header.Set(key, value)
	}

	resp, err := r.rpc.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	body, _ := io.ReadAll(resp.Body)

	response.Status = resp.StatusCode
	response.Length = resp.ContentLength

	uid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	response.ID = uid.String()
	response.Headers = domain.Header(resp.Header.Clone())
	response.Body = body

	return &response, nil
}
