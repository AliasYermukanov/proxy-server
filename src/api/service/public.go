package service

import (
	"context"

	"github.com/AliasYermukanov/proxy-server/src/api/transport"
	"github.com/AliasYermukanov/proxy-server/src/util"
)

func (s *service) ProxySend(ctx context.Context, req *transport.ProxySendRequest) (*transport.ProxySendResponse, error) {
	if req.Method != nil {
		err := s.checkMethod(*req.Method)
		if err != nil {
			return nil, err
		}
	}
	resp, err := s.externalRPC.ProxyService().SendRequest(ctx, req)
	if err != nil {
		return nil, err
	}
	s.container.RequestMap[resp.ID] = *resp

	return &transport.ProxySendResponse{Response: *resp}, nil
}

func (s *service) GetRequestByID(ctx context.Context, req *transport.GetRequestByIDRequest) (*transport.GetRequestByIDResponse, error) {
	if val, ok := s.container.RequestMap[req.ID]; ok {
		return &transport.GetRequestByIDResponse{Response: val}, nil
	}
	util.CommonError.Code = 404
	return nil, util.CommonError.SetDevMessage("no data")
}
