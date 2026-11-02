package api

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/muhammadatak/learn-gofiber/domain"
	"github.com/muhammadatak/learn-gofiber/dto"
)

type userApi struct {
	userService domain.UserService
}

func NewUser(app *fiber.App, userService domain.UserService) {
	ua := userApi{
		userService: userService,
	}

	app.Get("/users/:email", ua.Get)
}

func (ua userApi) Get(ctx fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	email, err := url.PathUnescape(ctx.Params("email"))
	if err != nil {
		return err
	}
	res, err := ua.userService.Get(c, email)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.JSON(dto.CreateResponseSuccess(res))
}
