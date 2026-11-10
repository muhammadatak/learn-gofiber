package service

import (
	"context"
	"database/sql"
	"path"
	"time"

	"github.com/google/uuid"
	"github.com/muhammadatak/learn-gofiber/domain"
	"github.com/muhammadatak/learn-gofiber/dto"
	"github.com/muhammadatak/learn-gofiber/internal/config"
)

type mediaService struct {
	mediaRepository domain.MediaRepository
	conf            *config.Config
}

func NewMedia(conf *config.Config, mediaRepository domain.MediaRepository) domain.MediaService {
	return &mediaService{
		conf:            conf,
		mediaRepository: mediaRepository,
	}
}

// Create implements [domain.MediaService].
func (m *mediaService) Create(ctx context.Context, req dto.CreateMediaRequest) (dto.MediaData, error) {
	data := domain.Media{
		Id:        uuid.NewString(),
		Path:      req.Path,
		CreatedAt: sql.NullTime{Valid: true, Time: time.Now()},
	}

	err := m.mediaRepository.Save(ctx, &data)

	if err != nil {
		return dto.MediaData{}, err
	}
	url := path.Join(m.conf.Server.Asset, data.Path)
	media := dto.MediaData{
		Id:   data.Id,
		Path: data.Path,
		Url:  url,
	}

	return media, nil
}
