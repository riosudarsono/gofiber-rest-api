package api

import (
	"context"
	"gofiber-rest-api/domain"
	"gofiber-rest-api/dto"
	"gofiber-rest-api/internal/util"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type bookApi struct {
	bookService domain.BookService
}

func NewBook(app *fiber.App, authMid fiber.Handler, bookService domain.BookService) {
	ba := &bookApi{bookService: bookService}

	app.Get("/books", authMid, ba.Index)
	app.Post("/books", authMid, ba.Create)
	app.Get("/books/:id", authMid, ba.Show)
	app.Put("/books/:id", authMid, ba.Update)
	app.Delete("/books/:id", authMid, ba.Delete)
}

func (ba bookApi) Index(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 5*time.Second)
	defer cancel()

	res, err := ba.bookService.Index(c)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(dto.CreateResponseSuccess(res))
}

func (ba bookApi) Create(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 5*time.Second)
	defer cancel()

	var book dto.CreateBookRequest
	if err := ctx.BodyParser(&book); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseError(err.Error()))
	}

	fails := util.Validate(book)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseErrorData("Validate Error", fails))
	}

	err := ba.bookService.Create(c, book)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.Status(http.StatusCreated).JSON(dto.CreateResponseSuccess(""))
}

func (ba bookApi) Show(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 5*time.Second)
	defer cancel()

	res, err := ba.bookService.Show(c, ctx.Params("id"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.Status(http.StatusOK).JSON(dto.CreateResponseSuccess(res))
}

func (ba bookApi) Update(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 5*time.Second)
	defer cancel()

	var book dto.UpdateBookRequest
	if err := ctx.BodyParser(&book); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseError(err.Error()))
	}
	fails := util.Validate(book)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseErrorData("Validate Error", fails))
	}
	err := ba.bookService.Update(c, book)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.Status(http.StatusOK).JSON(dto.CreateResponseSuccess(""))
}

func (ba bookApi) Delete(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 5*time.Second)
	defer cancel()

	err := ba.bookService.Delete(c, ctx.Params("id"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.Status(http.StatusOK).JSON(dto.CreateResponseSuccess(""))
}
