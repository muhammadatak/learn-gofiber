package service

import (
	"context"
	"database/sql"
	"errors"
	"path"
	"time"

	"github.com/google/uuid"
	"github.com/muhammadatak/learn-gofiber/domain"
	"github.com/muhammadatak/learn-gofiber/dto"
	"github.com/muhammadatak/learn-gofiber/internal/config"
)

type bookService struct {
	cnf                 *config.Config
	bookRepository      domain.BookRepository
	bookStockRepository domain.BookStockRepository
	mediaRepository     domain.MediaRepository
}

func NewBook(cnf *config.Config, bookRepository domain.BookRepository, bookStockRepository domain.BookStockRepository, mediaRepository domain.MediaRepository) domain.BookService {
	return &bookService{
		bookRepository:      bookRepository,
		bookStockRepository: bookStockRepository,
		mediaRepository:     mediaRepository,
		cnf:                 cnf,
	}
}

// Index implements [domain.BookService].
func (b *bookService) Index(ctx context.Context) ([]dto.BookData, error) {
	books, err := b.bookRepository.FindAll(ctx)

	if err != nil {
		return nil, err
	}

	coverId := make([]string, 0)
	for _, v := range books {
		if v.CoverId.Valid {
			coverId = append(coverId, v.CoverId.String)
		}
	}

	covers := make(map[string]string)
	if len(coverId) > 0 {
		coverDb, _ := b.mediaRepository.FindByIds(ctx, coverId)
		for _, v := range coverDb {
			covers[v.Id] = path.Join(b.cnf.Server.Asset, v.Path)
		}
	}

	var bookData []dto.BookData

	for _, v := range books {
		var coverUrl string
		if v2, e := covers[v.CoverId.String]; e {
			coverUrl = v2
		}
		bookData = append(bookData, dto.BookData{
			Id:          v.Id,
			Isbn:        v.Isbn,
			Title:       v.Title,
			Description: v.Description,
			CoverUrl:    coverUrl,
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

	var coverUrl string
	if book.CoverId.Valid {
		cover, _ := b.mediaRepository.FindById(ctx, book.CoverId.String)
		if cover.Path != "" {
			coverUrl = path.Join(b.cnf.Server.Asset, cover.Path)
		}
	}

	return dto.BookShowData{
		Stocks: stocksData,
		BookData: dto.BookData{Id: book.Id,
			Isbn:        book.Isbn,
			Title:       book.Title,
			Description: book.Description,
			CoverUrl:    coverUrl},
	}, nil
}

func (b *bookService) Create(ctx context.Context, book dto.CreateBookRequest) error {
	coverId := sql.NullString{Valid: false, String: book.CoverId}
	if book.CoverId != "" {
		coverId.Valid = true
	}

	data := domain.Book{
		Id:          uuid.NewString(),
		Isbn:        book.Isbn,
		Description: book.Description,
		Title:       book.Title,
		CoverId:     coverId,
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
	coverId := sql.NullString{Valid: false, String: book.CoverId}
	if book.CoverId != "" {
		coverId.Valid = true
	}
	exists.Isbn = book.Isbn
	exists.Title = book.Title
	exists.CoverId = coverId
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
