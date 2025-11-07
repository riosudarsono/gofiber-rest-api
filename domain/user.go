package domain

import "context"

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (User, error)
}
