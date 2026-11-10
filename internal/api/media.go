package api

import (
	"context"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/google/uuid"
	"github.com/muhammadatak/learn-gofiber/domain"
	"github.com/muhammadatak/learn-gofiber/dto"
	"github.com/muhammadatak/learn-gofiber/internal/config"
)

type mediaApi struct {
	mediaService domain.MediaService
	cnf          *config.Config
}

func NewMedia(cnf *config.Config, app *fiber.App, mediaService domain.MediaService, authMidd fiber.Handler) {

	ma := mediaApi{
		cnf:          cnf,
		mediaService: mediaService,
	}
	app.Post("/media", authMidd, ma.Create)
	app.Use("/media", static.New(cnf.Storage.BasePath))
}

func (ma mediaApi) Create(ctx fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	file, err := ctx.FormFile("media")
	if err != nil {
		return ctx.SendStatus(http.StatusBadRequest)
	}
	filename := uuid.NewString() + file.Filename
	path := filepath.Join(ma.cnf.Storage.BasePath, filename)
	err = ctx.SaveFile(file, path)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}

	res, err := ma.mediaService.Create(c, dto.CreateMediaRequest{
		Path: filename,
	})

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.JSON(dto.CreateResponseSuccess(res))
}
