// Package user : usecase/use layer from cleam architecture
package user

import "github.com/silvergama/clean-architecture/entity"

// Reader interface
type Reader interface {
	Get(id entity.ID) (*entity.User, error)
	Search(query string) ([]*entity.User, error)
	List() ([]*entity.User, error)
}

type Repository interface {
	Reader
}

type UseCase interface {
	GetUser(id entity.ID) (*entity.User, error)
	SearchUsers(query string) ([]*entity.User, error)
	ListUsers() ([]*entity.User, error)
}
