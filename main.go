package main

import (
	"net/http"

	jwtware "github.com/gofiber/contrib/v3/jwt"
	"github.com/gofiber/fiber/v3"
	"github.com/muhammadatak/learn-gofiber/dto"
	"github.com/muhammadatak/learn-gofiber/internal/api"
	"github.com/muhammadatak/learn-gofiber/internal/config"
	"github.com/muhammadatak/learn-gofiber/internal/connection"
	"github.com/muhammadatak/learn-gofiber/internal/repository"
	"github.com/muhammadatak/learn-gofiber/internal/service"
)

func main() {
	cnf := config.Get()
	dbConnection := connection.GetDatabase(cnf.Database)

	app := fiber.New()

	jwtMidd := jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(cnf.Jwt.Key)},
		ErrorHandler: func(ctx fiber.Ctx, err error) error {
			return ctx.Status(http.StatusUnauthorized).
				JSON(dto.CreateResponseError("endpoint perlu token, silakan login dulu"))
		},
	})

	customerRepository := repository.NewCustomer(dbConnection)

	customerService := service.NewCustomer(customerRepository)

	userRepository := repository.NewUser(dbConnection)
	userService := service.NewUser(userRepository)
	authService := service.NewAuth(cnf, userRepository)

	bookRepository := repository.NewBook(dbConnection)
	bookStockRepository := repository.NewBookStock(dbConnection)
	bookService := service.NewBook(bookRepository, bookStockRepository)

	api.NewAuth(app, authService)
	api.NewCustomer(app, jwtMidd, customerService)

	api.NewUser(app, userService)

	api.NewBook(app, bookService)

	_ = app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
}
