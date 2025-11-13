package dto

type BookData struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Isbn        string `json:"isbn"`
}

type CreateBookRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Isbn        string `json:"isbn" validate:"required"`
}

type UpdateBookRequest struct {
	ID string `json:"-"`
	CreateBookRequest
}
