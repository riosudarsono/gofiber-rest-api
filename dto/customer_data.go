package dto

type CustomerData struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type CreateCustomerRequest struct {
	Name string `json:"name" validate:"required"`
	Code string `json:"code" validate:"required"`
}

type UpdateCustomerRequest struct {
	ID   int64  `json:"-"`
	Name string `json:"name" validate:"required"`
	Code string `json:"code" validate:"required"`
}
