package main

import (
	"gofiber-rest-api/internal/api"
	"gofiber-rest-api/internal/config"
	"gofiber-rest-api/internal/connection"
	"gofiber-rest-api/internal/repository"
	"gofiber-rest-api/internal/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	conf := config.Get()
	dbConnection := connection.GetDatabase(conf.Database)
	customerRepo := repository.NewCustomer(dbConnection)
	customerService := service.NewCustomer(customerRepo)

	api.NewCustomer(app, customerService)
	_ = app.Listen(conf.Server.Host + ":" + conf.Server.Port)
}
