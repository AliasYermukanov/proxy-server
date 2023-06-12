package service

import (
	"context"

	"github.com/google/uuid"

	"github.com/AliasYermukanov/proxy-server/src/api/transport"
)

func (s *service) ProxySend(ctx context.Context, req *transport.ProxySendRequest) (*transport.ProxySendResponse, error) {
	uid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	req.ID = uid.String()

	resp, err := s.externalRPC.ProxyService().SendRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	resp.ID = req.ID

	s.store.Requests().StoreRequest(&req.Request)
	s.store.Requests().StoreResponse(resp)

	return &transport.ProxySendResponse{Response: *resp}, nil
}

func (s *service) GetDataByID(_ context.Context, req *transport.GetDataByIDRequest) (*transport.GetDataByIDResponse, error) {
	request, err := s.store.Requests().GetRequestByID(req.ID)
	if err != nil {
		return nil, err
	}

	response, err := s.store.Requests().GetResponseByID(req.ID)
	if err != nil {
		return nil, err
	}

	return &transport.GetDataByIDResponse{
		Request:  *request,
		Response: *response,
	}, nil
}

func (s *service) GetAllData(_ context.Context, _ *transport.GetAllDataRequest) (*transport.GetAllDataResponse, error) {
	requests := s.store.Requests().GetAllRequests()
	responses := s.store.Requests().GetAllResponses()

	return &transport.GetAllDataResponse{
		Requests:  requests,
		Responses: responses,
	}, nil
}
