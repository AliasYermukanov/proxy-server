package repository

type InmemStore interface {
	Requests() RequestsRepository
}
