package entity

import "errors"

// ErrNotFound not found
var ErrNotFound = errors.New("not found")

// ErrInvalidEntity invalid entity
var ErrInvalidEntity = errors.New("invalid entity")

// ErrCannotBeDeleted cannot be deleted
var ErrCannotBeDeleted = errors.New("cannot be deleted")

// ErrEnoughBooks enough books
var ErrNotEnoughBooks = errors.New("not enough books")

// ErrBookAlreadyBorrowed book already borrowed
var ErrBookAlreadyBorrowed = errors.New("books already borrowed")

// ErrBookNotBorrowed book not borrowed
var ErrBookNotBorrowed = errors.New("book not borrowed")
