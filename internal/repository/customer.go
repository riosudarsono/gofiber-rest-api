package repository

import (
	"context"
	"database/sql"
	"gofiber-rest-api/domain"
	"time"

	"github.com/doug-martin/goqu/v9"
)

type customerRepository struct {
	db *goqu.Database
}

func NewCustomer(con *sql.DB) domain.CustomerRepository {
	return &customerRepository{
		db: goqu.New("default", con),
	}
}

func (cr customerRepository) FindAll(ctx context.Context) (result []domain.Customer, err error) {
	dataset := cr.db.From("customers").
		Where(goqu.C("deleted_at").IsNull()).
		Order(goqu.C("id").Asc())
	err = dataset.ScanStructsContext(ctx, &result)
	return
}

func (cr customerRepository) FindByID(ctx context.Context, id int64) (result domain.CustomerUpdate, err error) {
	dataset := cr.db.From("customers").
		Where(goqu.C("deleted_at").IsNull(), goqu.C("id").Eq(id))
	_, err = dataset.ScanStructContext(ctx, &result)
	return
}

func (cr customerRepository) Save(ctx context.Context, customer *domain.Customer) error {
	executor := cr.db.Insert("customers").Rows(customer).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (cr customerRepository) Update(ctx context.Context, customer *domain.CustomerUpdate) error {
	executor := cr.db.Update("customers").
		Where(goqu.C("id").Eq(customer.ID)).
		Set(customer).
		Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (cr customerRepository) Delete(ctx context.Context, id int64) error {
	executor := cr.db.Update("customers").
		Where(goqu.C("id").Eq(id)).
		Set(goqu.Record{"deleted_at": sql.NullTime{Valid: true, Time: time.Now()}}).
		Executor()
	_, err := executor.ExecContext(ctx)
	return err
}
