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

type journalService struct {
	journalRepository   domain.JournalRepository
	bookRepository      domain.BookRepository
	bookStockRepository domain.BookStockRepository
	customerRepository  domain.CustomerRepository
	chargeRepository    domain.ChargeRepository
}

func NewJournal(journalRepository domain.JournalRepository,
	bookRepository domain.BookRepository,
	bookStockRepository domain.BookStockRepository, customerRepository domain.CustomerRepository, chargeRepository domain.ChargeRepository,
) domain.JournalService {

	return &journalService{
		journalRepository:   journalRepository,
		bookRepository:      bookRepository,
		bookStockRepository: bookStockRepository,
		customerRepository:  customerRepository,
		chargeRepository:    chargeRepository,
	}
}

// Create implements [domain.JournalService].
func (j *journalService) Create(ctx context.Context, req dto.CreateJournalRequest) error {
	// validaate customer and book and stock

	customer, err := j.customerRepository.FindById(ctx, req.CustomerId)
	if err != nil {
		return err
	}

	if customer.ID == "" {
		return errors.New("customer not found")
	}

	book, err := j.bookRepository.FindById(ctx, req.BookId)
	if err != nil {
		return err
	}

	if book.Id == "" {
		return errors.New("book not found")
	}

	stock, err := j.bookStockRepository.FindByBookAndCode(ctx, book.Id, req.BookStock)
	if err != nil {
		return err
	}
	if stock.Code == "" {
		return errors.New("code not avail")
	}
	if stock.Status != domain.BookStockAvailable {
		return errors.New("book is not available")
	}

	journal := domain.Journal{
		Id:         uuid.NewString(),
		BookId:     req.BookId,
		StockCode:  req.BookStock,
		CustomerId: req.CustomerId,
		Status:     domain.JournalStatusInProgress,
		BorrowedAt: sql.NullTime{Valid: false, Time: time.Now()},
		ReturnedAt: sql.NullTime{Valid: false, Time: time.Now()},
		DueAt:      sql.NullTime{Valid: true, Time: time.Now().Add(domain.JournalLoanDuration)},
	}

	stock.Status = domain.BookStockBorrowed
	stock.BorrowerId = sql.NullString{String: customer.ID, Valid: true}
	stock.BorrowedAt = sql.NullTime{Valid: true, Time: time.Now()}

	err = j.bookStockRepository.Update(ctx, &stock)
	if err != nil {
		return err
	}
	err = j.journalRepository.Save(ctx, &journal)

	return err
}

// Index implements [domain.JournalService].
func (j *journalService) Index(ctx context.Context, se domain.JournalSearch) ([]dto.JournalData, error) {
	journals, err := j.journalRepository.Find(ctx, se)

	if err != nil {
		return nil, err
	}

	customerIds := make([]string, 0)
	bookIds := make([]string, 0)

	for _, v := range journals {
		customerIds = append(customerIds, v.CustomerId)
		bookIds = append(bookIds, v.BookId)
	}

	customers := make(map[string]domain.Customer)
	books := make(map[string]domain.Book)
	// list customer id di journal
	if len(customerIds) > 0 {
		customerDb, _ := j.customerRepository.FindByIds(ctx, customerIds)
		for _, v := range customerDb {
			customers[v.ID] = v
		}
	}

	if len(bookIds) > 0 {
		bookDb, _ := j.bookRepository.FindByIds(ctx, bookIds)
		for _, v := range bookDb {
			books[v.Id] = v
		}
	}

	var journalData []dto.JournalData

	for _, v := range journals {

		book := books[v.BookId]
		customer := customers[v.CustomerId]

		journalData = append(journalData, dto.JournalData{
			Id:        v.Id,
			BookStock: v.StockCode,
			Book: dto.BookData{
				Id:          book.Id,
				Isbn:        book.Isbn,
				Title:       book.Title,
				Description: book.Description,
			},
			Customer: dto.CustomerData{
				ID:   customer.ID,
				Code: customer.Code,
				Name: customer.Name,
			},
			Status:     v.Status,
			BorrowedAt: v.BorrowedAt.Time,
			ReturnedAt: v.ReturnedAt.Time,
			DueAt:      v.DueAt.Time,
		})
	}
	return journalData, nil
}

// Return implements [domain.JournalService].
func (j *journalService) Return(ctx context.Context, req dto.ReturnJournalRequest) error {
	journal, err := j.journalRepository.FindById(ctx, req.JournalId)

	if err != nil {
		return err
	}

	if journal.Id == "" {
		return errors.New("journal not found")
	}

	stock, err := j.bookStockRepository.FindByBookAndCode(ctx, journal.BookId, journal.StockCode)

	if err != nil {
		return err
	}

	if stock.Code == "" {
		return errors.New("stock not found")
	}
	stock.Status = domain.BookStockAvailable
	stock.BorrowerId = sql.NullString{Valid: false}
	stock.BorrowedAt = sql.NullTime{Valid: false}
	err = j.bookStockRepository.Update(ctx, &stock)

	if err != nil {
		return err
	}

	journal.Status = domain.JournalStatusCompleted
	journal.ReturnedAt = sql.NullTime{Valid: true, Time: time.Now()}
	err = j.journalRepository.Update(ctx, &journal)

	if err != nil {
		return err
	}

	hoursLate := time.Now().Sub(journal.DueAt.Time).Hours()
	if hoursLate >= 24 {
		daysLate := int(hoursLate / 24)
		charge := domain.Charge{
			Id:           uuid.NewString(),
			JournalId:    journal.Id,
			DaysLate:     daysLate,
			DailyLateFee: 5000,
			Total:        5000 * daysLate,
			UserId:       req.UserId,
			CreatedAt:    sql.NullTime{Valid: true, Time: time.Now()},
		}
		err = j.chargeRepository.Save(ctx, &charge)
	}
	return err

}
