package main

import (
	"github.com/gofiber/fiber/v3"
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

	customerRepository := repository.NewCustomer(dbConnection)

	customerService := service.NewCustomer(customerRepository)

	api.NewCustomer(app, customerService)

	_ = app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
}
