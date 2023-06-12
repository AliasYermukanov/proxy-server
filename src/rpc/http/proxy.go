package http

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/AliasYermukanov/proxy-server/src/api/transport"
	"github.com/AliasYermukanov/proxy-server/src/domain"
)

type ProxyServiceRPC struct {
	rpc *RPC
}

func (r *ProxyServiceRPC) SendRequest(ctx context.Context, req *transport.ProxySendRequest) (*domain.Response, error) {
	var response domain.Response

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
	response.Headers = domain.Header(resp.Header.Clone())
	response.Body = body

	return &response, nil
}
