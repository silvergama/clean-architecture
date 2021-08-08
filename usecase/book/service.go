package book

import (
	"strings"
	"time"

	"github.com/silvergama/clean-architecture/entity"
)

// Service book usecase
type Service struct {
	repo Repository
}

// NewService create new service
func NewService(r Repository) *Service {
	return &Service{repo: r}
}

// CreateBook create a book
func (s *Service) CreateBook(title, author string, pages, quantity int) (entity.ID, error) {
	b, err := entity.NewBook(title, author, pages, quantity)
	if err != nil {
		return b.ID, err
	}

	return s.repo.Create(b)
}

// GetBook Get a book
func (s *Service) GetBook(id entity.ID) (*entity.Book, error) {
	b, err := s.repo.Get(id)
	if b == nil {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return b, nil
}

// SearchBooks searche books
func (s *Service) SearchBooks(query string) ([]*entity.Book, error) {
	books, err := s.repo.Search(strings.ToLower(query))
	if err != nil {
		return nil, err
	}
	if len(books) == 0 {
		return nil, entity.ErrNotFound
	}
	return books, nil
}

// ListBooks list all books
func (s *Service) ListBooks() ([]*entity.Book, error) {
	books, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	if len(books) == 0 {
		return nil, entity.ErrNotFound
	}

	return books, nil
}

// DeleteBook delete a book
func (s *Service) DeleteBook(id entity.ID) error {
	_, err := s.repo.Get(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

// UpdateBook update a book
func (s *Service) UpdateBook(b *entity.Book) error {
	err := b.Validate()
	if err != nil {
		return err
	}

	b.UpdatedAt = time.Now()
	return s.repo.Update(b)
}
