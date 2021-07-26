// Package loan : usecase/loan layer from cleam architecture
package loan

import "github.com/silvergama/clean-architecture/entity"

// UseCase use case interface
type UseCase interface {
	Borrow(u *entity.User, b *entity.Book) error
	Return(b *entity.Book) error
}
