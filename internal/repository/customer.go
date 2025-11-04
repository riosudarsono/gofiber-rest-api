package repository

import (
	"context"
	"gofiber-rest-api/domain"
)

type customerRepository struct {
}

func NewCutomer() domain.CustomerRepository {
	return &customerRepository{}
}

func (cr customerRepository) FindAll(ctx context.Context) ([]domain.Customer, error) {
	//TODO implement me
	panic("implement me")
}

func (cr customerRepository) FindByID(ctx context.Context, id string) (domain.Customer, error) {
	//TODO implement me
	panic("implement me")
}

func (cr customerRepository) Save(ctx context.Context, customer *domain.Customer) error {
	//TODO implement me
	panic("implement me")
}

func (cr customerRepository) Update(ctx context.Context, customer *domain.Customer) error {
	//TODO implement me
	panic("implement me")
}

func (cr customerRepository) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}
