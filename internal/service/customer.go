package service

import (
	"context"
	"database/sql"
	"errors"
	"gofiber-rest-api/domain"
	"gofiber-rest-api/dto"
	"time"
)

type customerService struct {
	customerRepo domain.CustomerRepository
}

func NewCustomer(customerRepo domain.CustomerRepository) domain.CustomerService {
	return &customerService{
		customerRepo: customerRepo,
	}
}

func (c customerService) Index(ctx context.Context) ([]dto.CustomerData, error) {
	customers, err := c.customerRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	var customerData []dto.CustomerData
	for _, customer := range customers {
		if !customer.ID.Valid {
			continue
		}
		customerData = append(customerData, dto.CustomerData{
			ID:   customer.ID.Int64,
			Name: customer.Name,
			Code: customer.Code,
		})
	}
	return customerData, nil
}

func (c customerService) Create(ctx context.Context, req dto.CreateCustomerRequest) error {
	customer := domain.Customer{
		Name: req.Name,
		Code: req.Code,
	}
	return c.customerRepo.Save(ctx, &customer)
}

func (c customerService) Update(ctx context.Context, req dto.UpdateCustomerRequest) error {
	persisted, err := c.customerRepo.FindByID(ctx, req.ID)
	if err != nil {
		return err
	}
	if persisted.ID == 0 {
		return errors.New("customer not found")
	}
	persisted.Name = req.Name
	persisted.Code = req.Code
	persisted.UpdatedAt = sql.NullTime{
		Valid: true,
		Time:  time.Now(),
	}
	return c.customerRepo.Update(ctx, &persisted)
}
