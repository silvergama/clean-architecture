// Package book : usecase/book layer from cleam architecture
package book

import "github.com/silvergama/clean-architecture/entity"

// Reader book interface
type Reader interface {
	Get(id entity.ID) (*entity.Book, error)
	Search(query string) ([]*entity.Book, error)
	List() ([]*entity.Book, error)
}

// Writer book interface
type Writer interface {
	Create(b *entity.Book) (entity.ID, error)
	Update(b *entity.Book) error
	Delete(entity.ID) error
}

// Repository book interface
type Repository interface {
	Reader
	Writer
}

// UseCase book interface
type UseCase interface {
	Get(id entity.ID) (*entity.Book, error)
	Search(query string) ([]*entity.Book, error)
	List() ([]*entity.Book, error)
	Create(b *entity.Book) (entity.ID, error)
	Update(b *entity.Book) error
	Delete(entity.ID) error
}
