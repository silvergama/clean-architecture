package user

import (
	"strings"

	"github.com/silvergama/clean-architecture/entity"
)

// Service interface
type Service struct {
	repo Repository
}

// NewService create new use case
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

// CreateUser create a new user
//func (s *Service) CreateUser(email, password, firstName, lastName string) (entity.ID, error) {
//	u, err := entity.NewUser(email, password, firstName, lastName)
//	if err != nil {
//		return u.ID, err
//	}
//
//	return s.repo.Create(u)
//}

// GetUser get an user
func (s *Service) GetUser(id entity.ID) (*entity.User, error) {
	return s.repo.Get(id)
}

// SearchUsers search users
func (s *Service) SearchUsers(query string) ([]*entity.User, error) {
	return s.repo.Search(strings.ToLower(query))
}

// ListUsers list users
func (s *Service) ListUsers() ([]*entity.User, error) {
	return s.repo.List()
}
