package biz

import "context"

// User is a user model.
type User struct {
	ID        int64
	UserName  string
	FirstName string
	LastName  string
	Gender    bool
	Email     string
	Password  string
	Disabled  bool
}

// UserRepo is a User repo.
type UserRepo interface {
	Save(context.Context, *User) (*User, error)
	FindByUsername(context.Context, string) (*User, error)
	FindByEmail(context.Context, string) (*User, error)
}
