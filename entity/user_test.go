package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("flash@dccomics.com", "The@Flash", "Barry", "Allen")
	assert.Nil(t, err)
	assert.NotNil(t, user.ID)
	assert.Equal(t, user.Email, "flash@dccomics.com")
	assert.Equal(t, user.FirstName, "Barry")
	assert.Equal(t, user.LastName, "Allen")
	assert.NotNil(t, user.CreatedAt)
	assert.NotNil(t, user.UpdatedAt)
	assert.Nil(t, user.Books)
}

func TestUserAddBook(t *testing.T) {
	user, _ := NewUser("flash@dccomics.com", "The@Flash", "Barry", "Allen")
	bID := NewID()
	bErr := user.AddBook(bID)
	assert.Nil(t, bErr)
	assert.Equal(t, 1, len(user.Books))

	err := user.AddBook(bID)
	assert.Equal(t, ErrBookAlreadyBorrowed, err)
}

func TestUserRemoveBook(t *testing.T) {
	user, _ := NewUser("flash@dccomics.com", "The@Flash", "Barry", "Allen")
	err := user.RemoveBook(NewID())
	assert.Equal(t, ErrNotFound, err)

	bID := NewID()
	_ = user.AddBook(bID)
	bErr := user.RemoveBook(bID)
	assert.Nil(t, bErr)
}

func TestUserGetBook(t *testing.T) {
	user, _ := NewUser("flash@dccomics.com", "The@Flash", "Barry", "Allen")
	bID := NewID()
	_ = user.AddBook(bID)
	id, err := user.GetBook(bID)
	assert.Nil(t, err)
	assert.Equal(t, id, bID)

	_, err = user.GetBook(NewID())
	assert.Equal(t, ErrNotFound, err)
}

func TestUserValidate(t *testing.T) {
	type args struct {
		Email     string
		FirstName string
		LastName  string
		Password  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "Should return a new user",
			args: args{
				Email:     "flash@dccomics.com",
				FirstName: "Barry",
				LastName:  "Allen",
				Password:  "The@Flash",
			},
			wantErr: nil,
		},
		{
			name: "Should return a error when missing email",
			args: args{
				Email:     "",
				FirstName: "Barry",
				LastName:  "Allen",
				Password:  "The@Flash",
			},
			wantErr: ErrInvalidEntity,
		},
		{
			name: "Should return a error when a missing first name",
			args: args{
				Email:     "flash@dccomics.com",
				FirstName: "",
				LastName:  "Allen",
				Password:  "The@Flash",
			},
			wantErr: ErrInvalidEntity,
		},
		{
			name: "Should return a error when a missing last name",
			args: args{
				Email:     "flash@dccomics.com",
				FirstName: "Barry",
				LastName:  "",
				Password:  "The@Flash",
			},
			wantErr: ErrInvalidEntity,
		},
		{
			name: "Should return a error when a missing password",
			args: args{
				Email:     "flash@dccomics.com",
				FirstName: "Barry",
				LastName:  "Allen",
				Password:  "",
			},
			wantErr: ErrInvalidEntity,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewUser(tt.args.Email, tt.args.Password, tt.args.FirstName, tt.args.LastName)
			assert.Equal(t, err, tt.wantErr)
		})
	}
}

func TestUserValidateHash(t *testing.T) {
	salt, _ := generateSalt()
	type fields struct {
		Salt     string
		Password string
		Hash     string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				Salt:     salt,
				Password: "user@Password",
				Hash:     "$2a$10$NEuhbLRV16n.idJwSH7yze/9DC0Sv2UjzDOlWF/o.Iws9xgIYcR6W",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				Salt:     tt.fields.Salt,
				Hash:     tt.fields.Hash,
				Password: tt.fields.Password,
			}
			hash, _ := generateHash(u.Password, u.Salt)
			u.Hash = hash
			err := u.ValidateHash()
			assert.Nil(t, err)
		})
	}
}
