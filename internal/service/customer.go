package service

import (
	"context"
	"gofiber-rest-api/domain"
	"gofiber-rest-api/dto"
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
		customerData = append(customerData, dto.CustomerData{
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
