package inmem

import (
	"sync"

	"github.com/AliasYermukanov/proxy-server/src/domain"
	"github.com/AliasYermukanov/proxy-server/src/util"
)

type Repository struct {
	mx        sync.RWMutex
	requests  map[string]*domain.Request
	responses map[string]*domain.Response
}

func NewRepository() *Repository {
	return &Repository{
		requests:  make(map[string]*domain.Request),
		responses: make(map[string]*domain.Response),
	}
}

func (r *Repository) StoreRequest(request *domain.Request) {
	r.mx.Lock()
	defer r.mx.Unlock()

	r.requests[request.ID] = request
}

func (r *Repository) StoreResponse(response *domain.Response) {
	r.mx.Lock()
	defer r.mx.Unlock()

	r.responses[response.ID] = response
}

func (r *Repository) GetAllRequests() []*domain.Request {
	r.mx.RLock()
	defer r.mx.RUnlock()

	var result []*domain.Request
	for _, value := range r.requests {
		result = append(result, value)
	}

	return result
}

func (r *Repository) GetAllResponses() []*domain.Response {
	r.mx.RLock()
	defer r.mx.RUnlock()

	var result []*domain.Response

	for _, value := range r.responses {
		result = append(result, value)
	}

	return result
}

func (r *Repository) GetRequestByID(ID string) (*domain.Request, error) {
	r.mx.RLock()
	defer r.mx.RUnlock()

	if val, ok := r.requests[ID]; ok {
		return val, nil
	}

	return nil, util.NoFound.SetDevMessage("no request data")
}

func (r *Repository) GetResponseByID(ID string) (*domain.Response, error) {
	r.mx.RLock()
	defer r.mx.RUnlock()

	if val, ok := r.responses[ID]; ok {
		return val, nil
	}

	return nil, util.NoFound.SetDevMessage("no response data")
}
