package user

import (
	"strings"

	"github.com/silvergama/clean-architecture/entity"
)

// inmen in memory repo
type inmen struct {
	m map[entity.ID]*entity.User
}

// newInmen create a new repository
func newInmen() *inmen {
	var m = map[entity.ID]*entity.User{}
	return &inmen{
		m: m,
	}
}

// create an user
func (i *inmen) Create(u *entity.User) (entity.ID, error) {
	i.m[u.ID] = u
	return u.ID, nil
}

// Get an user by ID
func (i *inmen) Get(id entity.ID) (*entity.User, error) {
	if i.m[id] != nil {
		return i.m[id], nil
	}
	return i.m[id], entity.ErrNotFound
}

// Update an user
func (i *inmen) Update(u *entity.User) error {
	_, err := i.Get(u.ID)
	if err != nil {
		return err
	}
	i.m[u.ID] = u
	return nil
}

// Search users
func (i *inmen) Search(query string) ([]*entity.User, error) {
	var res []*entity.User
	for _, j := range i.m {
		if strings.Contains(strings.ToLower(j.FirstName), query) {
			res = append(res, j)
		}
	}
	if len(res) == 0 {
		return nil, entity.ErrNotFound
	}
	return res, nil
}

// List users
func (i *inmen) List() ([]*entity.User, error) {
	var users []*entity.User
	for _, j := range i.m {
		users = append(users, j)
	}
	return users, nil
}

// Delete an user
func (i *inmen) Delete(id entity.ID) error {
	if i.m[id] == nil {
		return entity.ErrNotFound
	}
	i.m[id] = nil
	return nil
}
