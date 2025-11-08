package repository

import (
	"context"
	"database/sql"
	"gofiber-rest-api/domain"
	"time"

	"github.com/doug-martin/goqu/v9"
)

type bookRepository struct {
	db *goqu.Database
}

func NewBook(con *sql.DB) domain.BookRepository {
	return &bookRepository{
		db: goqu.New("default", con),
	}
}

func (b bookRepository) FindAll(ctx context.Context) (result []domain.Book, err error) {
	dataset := b.db.From("books").
		Where(goqu.C("deleted_at").IsNull()).
		Order(goqu.C("created_at").Asc())
	err = dataset.ScanStructsContext(ctx, &result)
	return
}

func (b bookRepository) FindByID(ctx context.Context, id string) (result domain.Book, err error) {
	dataset := b.db.From("books").
		Where(goqu.C("deleted_at").IsNull(), goqu.C("id").Eq(id))
	_, err = dataset.ScanStructContext(ctx, &result)
	return
}

func (b bookRepository) Save(ctx context.Context, book *domain.Book) error {
	executor := b.db.Insert("books").Rows(book).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (b bookRepository) Update(ctx context.Context, book *domain.Book) error {
	executor := b.db.Update("books").
		Where(goqu.C("id").Eq(book.ID)).
		Set(book).
		Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (b bookRepository) Delete(ctx context.Context, id string) error {
	executor := b.db.Update("books").
		Where(goqu.C("id").Eq(id)).
		Set(goqu.Record{"deleted_at": sql.NullTime{Valid: true, Time: time.Now()}}).
		Executor()
	_, err := executor.ExecContext(ctx)
	return err
}
