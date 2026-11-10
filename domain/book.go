package domain

import (
	"context"
	"database/sql"

	"github.com/muhammadatak/learn-gofiber/dto"
)

type Book struct {
	Id          string       `db:"id"`
	Title       string       `db:"title"`
	Description string       `db:"description"`
	Isbn        string       `db:"isbn"`
	CreatedAt   sql.NullTime `db:"created_at"`
	UpdatedAt   sql.NullTime `db:"updated_at"`
	DeletedAt   sql.NullTime `db:"deleted_at"`
}

type BookRepository interface {
	FindAll(ctx context.Context) ([]Book, error)
	FindById(ctx context.Context, id string) (Book, error)
	FindByIds(ctx context.Context, ids []string) ([]Book, error)
	Save(ctx context.Context, book *Book) error
	Update(ctx context.Context, book *Book) error
	Delete(ctx context.Context, id string) error
}

type BookService interface {
	Index(ctx context.Context) ([]dto.BookData, error)
	Get(ctx context.Context, id string) (dto.BookShowData, error)
	Create(ctx context.Context, book dto.CreateBookRequest) error
	Update(ctx context.Context, book dto.UpdateBookRequest) error
	Delete(ctx context.Context, id string) error
}
