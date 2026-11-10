package api

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/muhammadatak/learn-gofiber/domain"
	"github.com/muhammadatak/learn-gofiber/dto"
	"github.com/muhammadatak/learn-gofiber/internal/util"
)

type bookStockApi struct {
	bookStockService domain.BookStockService
}

func NewBookStock(app *fiber.App, authMid fiber.Handler, bookStockService domain.BookStockService) {
	bsa := bookStockApi{
		bookStockService: bookStockService,
	}

	app.Post("/book-stocks", authMid, bsa.Create)
	app.Delete("/book-stocks", authMid, bsa.Delete)
}

func (bsa bookStockApi) Create(ctx fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.CreateBookStockRequest
	if err := ctx.Bind().Body(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}
	fails := util.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadRequest).
			JSON(dto.CreateResponseErrorData("validasi gagal", fails))
	}
	err := bsa.bookStockService.Create(c, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.Status(http.StatusCreated).JSON(dto.CreateResponseSuccess("ok"))

}

func (bsa bookStockApi) Delete(ctx fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.DeleteBookStockRequest

	codeStr := ctx.Query("code")
	if codeStr == "" {
		return ctx.Status(http.StatusBadRequest).JSON(dto.CreateResponseError("parameter delete wajib diisi"))
	}
	codes := strings.Split(codeStr, ";")

	req.Codes = codes
	err := bsa.bookStockService.Delete(c, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.SendStatus(http.StatusNoContent)

}
