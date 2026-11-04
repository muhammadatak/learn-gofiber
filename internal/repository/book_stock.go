package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/muhammadatak/learn-gofiber/domain"
)

type bookStockRepository struct {
	db *goqu.Database
}

func NewBookStock(con *sql.DB) domain.BookStockRepository {
	return &bookStockRepository{
		db: goqu.New("default", con),
	}
}

func (b *bookStockRepository) FindByBookAndCode(ctx context.Context, id string, code string) (result domain.BookStock, err error) {
	dataset := b.db.From("book_stocks").Where(goqu.C("id").Eq(id), goqu.C("c0de").Eq(code))
	_, err = dataset.ScanStructContext(ctx, &result)
	return
}

func (b *bookStockRepository) FindByBookId(ctx context.Context, id string) (result domain.BookStock, err error) {
	dataset := b.db.From("book_stocks").Where(goqu.C("id").Eq(id), goqu.C("deleted_at").IsNull())
	_, err = dataset.ScanStructContext(ctx, &result)
	return
}

// Save implements [domain.BookStockRepository].
func (b *bookStockRepository) Save(ctx context.Context, data []domain.BookStock) error {
	executor := b.db.Insert("book_stocks").Rows(data).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

// Update implements [domain.BookStockRepository].
func (b *bookStockRepository) Update(ctx context.Context, stock *domain.BookStock) error {
	executor := b.db.Update("book_stocks").Where(goqu.C("code").Eq(stock.Code)).Set(stock).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (b *bookStockRepository) Delete(ctx context.Context, id string) error {
	executor := b.db.Update("book_stocks").Where(goqu.C("id").Eq(id)).Set(goqu.Record{"deleted_at": sql.NullTime{Valid: true, Time: time.Now()}}).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (b *bookStockRepository) DeleteByCodes(ctx context.Context, codes []string) error {
	executor := b.db.Update("book_stocks").Where(goqu.C("id").Eq(codes)).Set(goqu.Record{"deleted_at": sql.NullTime{Valid: true, Time: time.Now()}}).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}
