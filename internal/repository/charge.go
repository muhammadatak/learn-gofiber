package repository

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.com/muhammadatak/learn-gofiber/domain"
)

type chargeRepository struct {
	db *goqu.Database
}

func NewCharge(con *sql.DB) domain.ChargeRepository {
	return &chargeRepository{
		db: goqu.New("default", con),
	}
}

// Save implements [domain.ChargeRepository].
func (c *chargeRepository) Save(ctx context.Context, charge *domain.Charge) error {
	exec := c.db.Insert("charges").Rows(charge).Executor()
	_, err := exec.ExecContext(ctx)
	return err
}
