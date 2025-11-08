package domain

import "context"

type Book struct {
	ID        string `db:"id"`
	Name      string `db:"name"`
	Desc      string `db:"desc"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
	DeletedAt string `db:"deleted_at"`
}

type BookRepository interface {
	FindAll(ctx context.Context) ([]Book, error)
	FindByID(ctx context.Context, id string) (Book, error)
	Save(ctx context.Context, book *Book) error
	Update(ctx context.Context, book *Book) error
	Delete(ctx context.Context, id string) error
}
