package service

import (
	"context"
	"database/sql"
	"errors"
	"gofiber-rest-api/domain"
	"gofiber-rest-api/dto"
	"time"

	"github.com/google/uuid"
)

type bookService struct {
	bookRepository      domain.BookRepository
	bookStockRepository domain.BookStockRepository
}

func NewBook(bookRepository domain.BookRepository, bookStockRepository domain.BookStockRepository) domain.BookService {
	return &bookService{
		bookRepository:      bookRepository,
		bookStockRepository: bookStockRepository,
	}
}

func (b bookService) Index(ctx context.Context) ([]dto.BookData, error) {
	result, err := b.bookRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	var books []dto.BookData
	for _, book := range result {
		books = append(books, dto.BookData{
			ID:          book.ID,
			Isbn:        book.Isbn,
			Title:       book.Title,
			Description: book.Description,
		})
	}
	return books, nil
}

func (b bookService) Show(ctx context.Context, id string) (dto.BookData, error) {
	data, err := b.bookRepository.FindByID(ctx, id)
	if err != nil {
		return dto.BookData{}, err
	}
	if data.ID == "" {
		return dto.BookData{}, errors.New("book not found")
	}
	return dto.BookData{
		ID:          data.ID,
		Isbn:        data.Isbn,
		Title:       data.Title,
		Description: data.Description,
	}, nil
}

func (b bookService) Create(ctx context.Context, req dto.CreateBookRequest) error {
	book := domain.Book{
		ID:          uuid.NewString(),
		Title:       req.Title,
		Description: req.Description,
		Isbn:        req.Isbn,
		CreatedAt: sql.NullTime{
			Valid: true,
			Time:  time.Now(),
		},
	}
	return b.bookRepository.Save(ctx, &book)
}

func (b bookService) Update(ctx context.Context, req dto.UpdateBookRequest) error {
	persisted, err := b.bookRepository.FindByID(ctx, req.ID)
	if err != nil {
		return err
	}
	if persisted.ID == "" {
		return errors.New("book not found")
	}
	persisted.Isbn = req.Isbn
	persisted.Title = req.Title
	persisted.Description = req.Description
	persisted.UpdatedAt = sql.NullTime{
		Valid: true,
		Time:  time.Now(),
	}
	return b.bookRepository.Save(ctx, &persisted)
}

func (b bookService) Delete(ctx context.Context, id string) error {
	persisted, err := b.bookRepository.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if persisted.ID == "" {
		return errors.New("book not found")
	}
	err = b.bookRepository.Delete(ctx, persisted.ID)
	if err != nil {
		return err
	}
	return b.bookStockRepository.DeleteByBookID(ctx, persisted.ID)
}
