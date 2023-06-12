package inmem

import (
	"context"
	"sync"

	"github.com/go-kit/log"

	"github.com/AliasYermukanov/proxy-server/src/repository"
)

type Store struct {
	mu sync.Mutex

	RequestsRepo repository.RequestsRepository

	logger log.Logger
}

func NewStore(_ context.Context, logger log.Logger) (*Store, error) {
	store := &Store{
		logger: log.With(logger, "rep", "inmem"),
	}

	return store, nil
}

func (s *Store) Requests() repository.RequestsRepository {
	if s.RequestsRepo != nil {
		return s.RequestsRepo
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.RequestsRepo = NewRepository()

	return s.RequestsRepo
}
