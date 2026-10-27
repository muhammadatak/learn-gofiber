package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/muhammadatak/learn-gofiber/internal/config"
	"github.com/muhammadatak/learn-gofiber/internal/connection"
	"github.com/muhammadatak/learn-gofiber/internal/repository"
)

func main() {
	cnf := config.Get()
	dbConnection := connection.GetDatabase(cnf.Database)

	app := fiber.New()

	customerRepository := repository.NewCustomer(dbConnection)

	app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
}
