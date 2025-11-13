package domain

import (
	"context"
	"database/sql"
)

type BookStock struct {
	BookID     string         `db:"book_id"`
	Code       string         `db:"code"`
	Status     string         `db:"status"`
	BorrowerID sql.NullString `db:"borrower_id"`
	BorrowedAt sql.NullTime   `db:"borrowed_at"`
}

type BookStockRepository interface {
	FindByBookID(ctx context.Context, bookID string) (*BookStock, error)
	FindByBookAndCode(ctx context.Context, bookID, code string) (*BookStock, error)
	Save(ctx context.Context, stocks []BookStock) error
	Update(ctx context.Context, stock *BookStock) error
	DeleteByBookID(ctx context.Context, bookID string) error
	DeleteByCodes(ctx context.Context, codes []string) error
}
