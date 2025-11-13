package repository

import (
	"context"
	"database/sql"
	"gofiber-rest-api/domain"

	"github.com/doug-martin/goqu/v9"
)

type bookStockRepository struct {
	db *goqu.Database
}

func NewBookStock(con *sql.DB) domain.BookStockRepository {
	return &bookStockRepository{
		db: goqu.New("default", con),
	}
}

func (b bookStockRepository) FindByBookID(ctx context.Context, bookID string) (result *domain.BookStock, err error) {
	dataset := b.db.From("book_stocks").Where(goqu.C("book_id").Eq(bookID))
	err = dataset.ScanStructsContext(ctx, &result)
	return
}

func (b bookStockRepository) FindByBookAndCode(ctx context.Context, bookID, code string) (result *domain.BookStock, err error) {
	dataset := b.db.From("book_stocks").Where(
		goqu.C("book_id").Eq(bookID),
		goqu.C("code").Eq(code))
	_, err = dataset.ScanStructContext(ctx, &result)
	return
}

func (b bookStockRepository) Save(ctx context.Context, stocks []domain.BookStock) error {
	executor := b.db.Insert("book_stocks").Rows(stocks).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (b bookStockRepository) Update(ctx context.Context, stock *domain.BookStock) error {
	executor := b.db.Update("book_stocks").
		Where(goqu.C("code").Eq(stock.Code)).
		Set(stock).
		Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (b bookStockRepository) DeleteByBookID(ctx context.Context, bookID string) error {
	executor := b.db.Delete("book_stocks").
		Where(goqu.C("book_id").Eq(bookID)).
		Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (b bookStockRepository) DeleteByCodes(ctx context.Context, codes []string) error {
	executor := b.db.Delete("book_stocks").
		Where(goqu.C("code").In(codes)).
		Executor()
	_, err := executor.ExecContext(ctx)
	return err
}
