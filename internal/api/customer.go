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

type customerApi struct {
	customerService domain.CustomerService
}

func NewCustomer(app *fiber.App, jwtMid fiber.Handler, customerService domain.CustomerService) {
	ca := customerApi{customerService}
	app.Get("/customers", jwtMid, ca.Index)
	app.Post("/customers", jwtMid, ca.Create)
	app.Put("/customers/:id", jwtMid, ca.Update)
	app.Delete("/customers/:id", jwtMid, ca.Delete)
	app.Get("/customers/:id", jwtMid, ca.Show)
}

func (ca customerApi) Index(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	response, err := ca.customerService.Index(c)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.JSON(dto.CreateResponseSuccess(response))
}

func (ca customerApi) Create(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.CreateCustomerRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseError(err.Error()))
	}

	fails := util.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseErrorData("Validate Error", fails))
	}

	err := ca.customerService.Create(c, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.Status(http.StatusCreated).JSON(dto.CreateResponseSuccess(""))
}

func (ca customerApi) Update(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.UpdateCustomerRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseError(err.Error()))
	}

	fails := util.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseErrorData("Validate Error", fails))
	}

	idInt, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseError("ID must be an Number"))
	}

	req.ID = int64(idInt)
	err = ca.customerService.Update(c, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.Status(http.StatusOK).JSON(dto.CreateResponseSuccess(""))
}

func (ca customerApi) Delete(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	idInt, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseError("ID must be an Number"))
	}

	errDelete := ca.customerService.Delete(c, int64(idInt))
	if errDelete != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(errDelete.Error()))
	}

	return ctx.SendStatus(http.StatusNoContent)
}

func (ca customerApi) Show(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	idInt, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseError("ID must be an Number"))
	}

	data, errShow := ca.customerService.Show(c, int64(idInt))
	if errShow != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(errShow.Error()))
	}

	return ctx.Status(http.StatusOK).JSON(dto.CreateResponseSuccess(data))
}
