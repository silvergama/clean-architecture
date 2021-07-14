package entity

import "errors"

// ErrNotFound not found
var ErrNotFound = errors.New("Not found")

// ErrInvalidEntity invalid entity
var ErrInvalidEntity = errors.New("Invalid entity")

// ErrCannotBeDeleted cannot be deleted
var ErrCannotBeDeleted = errors.New("Cannot be deleted")

// ErrEnoughBooks enough books
var ErrNotEnoughBooks = errors.New("Not enough books")

// ErrBookAlreadyBorrowed book already borrowed
var ErrBookAlreadyBorrowed = errors.New("Books already borrowed")

// ErrBookNotBorrowed book not borrowed
var ErrBookNotBorrowed = errors.New("Book not borrowed")
