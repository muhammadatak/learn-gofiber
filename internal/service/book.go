package service

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/muhammadatak/learn-gofiber/domain"
	"github.com/muhammadatak/learn-gofiber/dto"
)

type bookService struct {
	bookRepository      domain.BookRepository
	bookStockRepository domain.BookStockRepository
}

func NewBook(bookRepository domain.BookRepository, bookStockRepository domain.BookStockRepository) domain.BookService {
	return &bookService{
		bookRepository:      bookRepository,
		bookStockRepository: bookStockRepository,
	}
}

// Index implements [domain.BookService].
func (b *bookService) Index(ctx context.Context) ([]dto.BookData, error) {
	books, err := b.bookRepository.FindAll(ctx)

	if err != nil {
		return nil, err
	}

	var bookData []dto.BookData

	for _, v := range books {
		bookData = append(bookData, dto.BookData{
			Id:          v.Id,
			Isbn:        v.Isbn,
			Title:       v.Title,
			Description: v.Description,
		})
	}

	return bookData, nil
}

// Get implements [domain.BookService].
func (b *bookService) Get(ctx context.Context, id string) (dto.BookShowData, error) {
	book, err := b.bookRepository.FindById(ctx, id)

	if err != nil {
		return dto.BookShowData{}, err
	}

	if book.Id == "" {
		return dto.BookShowData{}, errors.New("book tidak ada")
	}

	stocks, err := b.bookStockRepository.FindByBookId(ctx, book.Id)
	if err != nil {
		return dto.BookShowData{}, err
	}

	stocksData := make([]dto.BookStockData, 0)

	for _, v := range stocks {
		stocksData = append(stocksData, dto.BookStockData{
			Code:   v.Code,
			Status: v.Status,
		})
	}

	return dto.BookShowData{
		Stocks: stocksData,
		BookData: dto.BookData{Id: book.Id,
			Isbn:        book.Isbn,
			Title:       book.Title,
			Description: book.Description},
	}, nil
}

func (b *bookService) Create(ctx context.Context, book dto.CreateBookRequest) error {
	data := domain.Book{
		Id:          uuid.NewString(),
		Isbn:        book.Isbn,
		Description: book.Description,
		Title:       book.Title,
		CreatedAt:   sql.NullTime{Valid: true, Time: time.Now()},
	}

	return b.bookRepository.Save(ctx, &data)
}

// Update implements [domain.BookService].
func (b *bookService) Update(ctx context.Context, book dto.UpdateBookRequest) error {
	exists, err := b.bookRepository.FindById(ctx, book.Id)

	if err != nil {
		return err
	}

	if exists.Id == "" {
		return errors.New("book tidak ada")
	}

	exists.Isbn = book.Isbn
	exists.Title = book.Title
	exists.Description = book.Description
	exists.UpdatedAt = sql.NullTime{Valid: true, Time: time.Now()}

	return b.bookRepository.Update(ctx, &exists)
}

// Delete implements [domain.BookService].
func (b *bookService) Delete(ctx context.Context, id string) error {
	exists, err := b.bookRepository.FindById(ctx, id)

	if err != nil {
		return err
	}

	if exists.Id == "" {
		return errors.New("data book not found")
	}

	err = b.bookRepository.Delete(ctx, id)

	return err
}
