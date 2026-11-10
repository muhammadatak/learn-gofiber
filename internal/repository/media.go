package repository

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.com/muhammadatak/learn-gofiber/domain"
)

type MediaRepository struct {
	db *goqu.Database
}

func NewMedia(con *sql.DB) domain.MediaRepository {
	return &MediaRepository{
		db: goqu.New("default", con),
	}
}

// FindById implements [domain.MediaRepository].
func (m *MediaRepository) FindById(ctx context.Context, id string) (result domain.Media, err error) {
	dataset := m.db.From("media").Where(goqu.C("id").Eq(id))
	_, err = dataset.ScanStructContext(ctx, &result)
	return
}

// FindByIds implements [domain.MediaRepository].
func (m *MediaRepository) FindByIds(ctx context.Context, ids []string) (result []domain.Media, err error) {
	dataset := m.db.From("media").Where(goqu.C("id").Eq(ids))
	err = dataset.ScanStructsContext(ctx, &result)
	return
}

// Save implements [domain.MediaRepository].
func (m *MediaRepository) Save(ctx context.Context, media *domain.Media) error {
	executor := m.db.Insert("media").Rows(media).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}
