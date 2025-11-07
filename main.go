package main

import (
	"gofiber-rest-api/dto"
	"gofiber-rest-api/internal/api"
	"gofiber-rest-api/internal/config"
	"gofiber-rest-api/internal/connection"
	"gofiber-rest-api/internal/repository"
	"gofiber-rest-api/internal/service"

	jwtMid "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	conf := config.Get()
	dbConnection := connection.GetDatabase(conf.Database)
	jwtMid_ := jwtMid.New(jwtMid.Config{
		SigningKey: jwtMid.SigningKey{Key: []byte(conf.Jwt.Key)},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(dto.CreateResponseError("Unauthorized"))
		},
	})

	customerRepo := repository.NewCustomer(dbConnection)
	authRepo := repository.NewUser(dbConnection)

	customerService := service.NewCustomer(customerRepo)
	authService := service.NewAuth(conf, authRepo)

	api.NewCustomer(app, jwtMid_, customerService)
	api.NewAuth(app, authService)

	_ = app.Listen(conf.Server.Host + ":" + conf.Server.Port)
}
