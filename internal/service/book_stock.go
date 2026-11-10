package service

import (
	"context"
	"errors"

	"github.com/muhammadatak/learn-gofiber/domain"
	"github.com/muhammadatak/learn-gofiber/dto"
)

type bookStockService struct {
	bookRepository      domain.BookRepository
	bookStockRepository domain.BookStockRepository
}

func NewBookStock(bookRepository domain.BookRepository,
	bookStockRepository domain.BookStockRepository) domain.BookStockService {
	return &bookStockService{
		bookRepository:      bookRepository,
		bookStockRepository: bookStockRepository,
	}
}

// Create implements [domain.BookStockService].
func (b *bookStockService) Create(ctx context.Context, req dto.CreateBookStockRequest) error {
	exists, err := b.bookRepository.FindById(ctx, req.BookId)

	if err != nil {
		return err
	}

	if exists.Id == "" {
		return errors.New("book tidak ada")
	}

	stocks := make([]domain.BookStock, 0)

	for _, v := range req.Codes {
		stocks = append(stocks, domain.BookStock{
			Code:   v,
			BookId: req.BookId,
			Status: domain.BookStockAvailable,
		})
	}
	return b.bookStockRepository.Save(ctx, stocks)

}

// Delete implements [domain.BookStockService].
func (b *bookStockService) Delete(ctx context.Context, req dto.DeleteBookStockRequest) error {
	return b.bookStockRepository.DeleteByCodes(ctx, req.Codes)
}
