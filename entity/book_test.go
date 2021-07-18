package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBook(t *testing.T) {
	b, err := NewBook("American Gods", "Neil Gaiman", 100, 1)
	assert.Nil(t, err)
	assert.NotNil(t, b.ID)
	assert.Equal(t, b.Title, "American Gods")
	assert.Equal(t, b.Author, "Neil Gaiman")
	assert.Equal(t, b.Pages, 100)
	assert.Equal(t, b.Quantity, 1)

}

func TestBook_Validate(t *testing.T) {
	type fields struct {
		Title    string
		Author   string
		Pages    int
		Quantity int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr error
	}{
		{
			name: "Should not return error",
			fields: fields{
				Title:    "American Gods",
				Author:   "Neil Gaiman",
				Pages:    100,
				Quantity: 1,
			},
			wantErr: nil,
		},
		{
			name: "Should return an error when title is empty",
			fields: fields{
				Title:    "",
				Author:   "Neil Gaiman",
				Pages:    100,
				Quantity: 1,
			},
			wantErr: ErrInvalidEntity,
		},
		{
			name: "Should return an error when author is empty",
			fields: fields{
				Title:    "American Gods",
				Author:   "",
				Pages:    100,
				Quantity: 1,
			},
			wantErr: ErrInvalidEntity,
		},
		{
			name: "Should return an error when the page is less than or equal to zero",
			fields: fields{
				Title:    "American Gods",
				Author:   "Neil Gaiman",
				Pages:    0,
				Quantity: 1,
			},
			wantErr: ErrInvalidEntity,
		},
		{
			name: "Should return an error when the quantity is less than or equal to zero",
			fields: fields{
				Title:    "American Gods",
				Author:   "Neil Gaiman",
				Pages:    100,
				Quantity: 0,
			},
			wantErr: ErrInvalidEntity,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewBook(
				tt.fields.Title,
				tt.fields.Author,
				tt.fields.Pages,
				tt.fields.Quantity,
			)
			assert.Equal(t, err, tt.wantErr)
		})
	}
}
