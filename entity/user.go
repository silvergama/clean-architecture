package entity

import (
	"crypto/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User data
type User struct {
	ID        ID
	Email     string
	Salt      string
	Hash      string
	FirstName string
	LastName  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Books     []ID
}

// NewUser create a new user
func NewUser(email, pwd, firstName, lastName string) (*User, error) {
	u := &User{
		ID:        NewID(),
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		CreatedAt: time.Now(),
	}

	salt, err := generateSalt()
	if err != nil {
		return nil, err
	}
	u.Salt = salt

	hash, err := generateHash(pwd, salt)
	if err != nil {
		return nil, err
	}
	u.Hash = hash

	err = u.Validate(pwd)
	if err != nil {
		return nil, ErrInvalidEntity
	}

	return u, nil
}

// AddBook add a book
func (u *User) AddBook(id ID) error {
	_, err := u.GetBook(id)
	if err != nil {
		return ErrBookAlreadyBorrowed
	}

	u.Books = append(u.Books, id)
	return nil
}

// RemoveBook remove a book
func (u *User) RemoveBook(id ID) error {
	for i, b := range u.Books {
		if b == id {
			u.Books = append(u.Books[:i], u.Books[i+1:]...)
			return nil
		}
	}

	return ErrNotFound
}

// GetBook get a book
func (u *User) GetBook(id ID) (ID, error) {
	for _, b := range u.Books {
		if b == id {
			return id, nil
		}
	}

	return id, ErrNotFound
}

// Validate validade data
func (u *User) Validate(password string) error {
	if u.Email == "" || u.FirstName == "" || u.LastName == "" || password == "" {
		return ErrInvalidEntity
	}

	return nil
}

// ValidateHash validate user hash by password
func (u *User) ValidateHash(pwd string) error {
	hash, err := generateHash(pwd, u.Salt)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Hash), []byte(hash))
	if err != nil {
		return err
	}

	return nil
}

func generateSalt() (string, error) {
	salt := make([]byte, 32)
	_, err := rand.Read(salt[:])
	if err != nil {
		return "", err
	}

	return string(salt), nil
}

func generateHash(pwd, salt string) (string, error) {
	var pwdBytes = []byte(pwd)
	var saltBytes = []byte(salt)
	pwdBytes = append(pwdBytes, saltBytes...)
	hash, err := bcrypt.GenerateFromPassword(pwdBytes, 10)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
