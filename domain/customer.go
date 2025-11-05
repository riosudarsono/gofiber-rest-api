package domain

import (
	"context"
	"database/sql"
	"gofiber-rest-api/dto"
)

type Customer struct {
	ID        sql.NullInt64 `db:"id" goqu:"skipinsert"`
	Code      string        `db:"code"`
	Name      string        `db:"name"`
	CreatedAt sql.NullTime  `db:"created_at" goqu:"skipinsert"`
	UpdatedAt sql.NullTime  `db:"updated_at"`
	DeletedAt sql.NullTime  `db:"deleted_at"`
}

type CustomerRepository interface {
	FindAll(ctx context.Context) ([]Customer, error)
	FindByID(ctx context.Context, id string) (Customer, error)
	Save(ctx context.Context, customer *Customer) error
	Update(ctx context.Context, customer *Customer) error
	Delete(ctx context.Context, id string) error
}
type CustomerService interface {
	Index(ctx context.Context) ([]dto.CustomerData, error)
	Create(ctx context.Context, req dto.CreateCustomerRequest) error
}
