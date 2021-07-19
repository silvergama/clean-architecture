// Package user : usecase/use layer from cleam architecture
package user

import "github.com/silvergama/clean-architecture/entity"

// Reader user reader
type Reader interface {
	Get(id entity.ID) (*entity.User, error)
	Search(query string) ([]*entity.User, error)
	List() ([]*entity.User, error)
}

// Writer user writer
type Writer interface {
	Create(u *entity.User) (entity.ID, error)
	Update(u *entity.User) error
	Delete(id entity.ID) error
}

// Repository interface
type Repository interface {
	Reader
	Writer
}

// UseCase interface
type UseCase interface {
	GetUser(id entity.ID) (*entity.User, error)
	SearchUsers(query string) ([]*entity.User, error)
	ListUsers() ([]*entity.User, error)
	CreateUser(email, password, firstName, lastName string) (entity.ID, error)
	UpdateUser(u *entity.User) error
	DeleteUser(id entity.ID) error
}
