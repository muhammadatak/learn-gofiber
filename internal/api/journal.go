package api

import (
	"context"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/muhammadatak/learn-gofiber/domain"
	"github.com/muhammadatak/learn-gofiber/dto"
	"github.com/muhammadatak/learn-gofiber/internal/util"
)

type journalApi struct {
	journalService domain.JournalService
}

func NewJournal(app *fiber.App, authMid fiber.Handler, journalService domain.JournalService) {
	ja := journalApi{
		journalService: journalService,
	}

	app.Get("/journals", authMid, ja.Index)
	app.Post("/journals", authMid, ja.Create)
	app.Post("/journals/return", authMid, ja.Return)

}

func (ja journalApi) Index(ctx fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	customerStr := ctx.Query("customer")
	statusStr := ctx.Query("status")
	search := domain.JournalSearch{}
	if customerStr != "" {
		search.CustomerId = customerStr
	}

	if statusStr != "" {
		search.Status = statusStr
	}

	res, err := ja.journalService.Index(c, search)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.JSON(dto.CreateResponseSuccess(res))
}

func (ja journalApi) Create(ctx fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.CreateJournalRequest

	if err := ctx.Bind().Body(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}

	fails := util.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseErrorData("Validation failed", fails))
	}

	err := ja.journalService.Create(c, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))

	}
	return ctx.Status(http.StatusCreated).JSON("")

}

func (ja journalApi) Return(ctx fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.ReturnJournalRequest

	if err := ctx.Bind().Body(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}

	fails := util.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseErrorData("Validation failed", fails))
	}

	err := ja.journalService.Return(c, req)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))

	}

	return ctx.Status(http.StatusCreated).JSON("")

}
