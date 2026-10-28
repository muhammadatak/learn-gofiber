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

type customerService struct {
	customerRepository domain.CustomerRepository
}

func NewCustomer(customerRepository domain.CustomerRepository) domain.CustomerService {
	return &customerService{
		customerRepository: customerRepository,
	}
}

// Index implements [domain.CustomerService].
func (c *customerService) Index(ctx context.Context) ([]dto.CustomerData, error) {
	customers, err := c.customerRepository.FindAll(ctx)

	if err != nil {
		return nil, err
	}

	var customerData []dto.CustomerData

	for _, v := range customers {
		customerData = append(customerData, dto.CustomerData{
			ID:   v.ID,
			Code: v.Code,
			Name: v.Name,
		})
	}

	return customerData, nil
}

func (c *customerService) Get(ctx context.Context, id string) (dto.CustomerData, error) {
	customer, err := c.customerRepository.FindById(ctx, id)

	if err != nil {
		return dto.CustomerData{}, err
	}

	if customer.ID == "" {
		return dto.CustomerData{}, errors.New("data customer not found")
	}

	return dto.CustomerData{
		ID:   customer.ID,
		Code: customer.Code,
		Name: customer.Name,
	}, nil
}

func (c *customerService) Create(ctx context.Context, data dto.CreateCustomerRequest) error {
	customer := domain.Customer{
		ID:        uuid.NewString(),
		Code:      data.Code,
		Name:      data.Name,
		CreatedAt: sql.NullTime{Valid: true, Time: time.Now()},
	}

	return c.customerRepository.Save(ctx, &customer)
}

func (c *customerService) Update(ctx context.Context, data dto.UpdateCustomerRequest) error {
	exists, err := c.customerRepository.FindById(ctx, data.ID)

	if err != nil {
		return err
	}

	if exists.ID == "" {
		return errors.New("data customer not found")
	}

	exists.Code = data.Code
	exists.Name = data.Name

	err = c.customerRepository.Update(ctx, &exists)

	return err
}

func (c *customerService) Delete(ctx context.Context, id string) error {
	exists, err := c.customerRepository.FindById(ctx, id)

	if err != nil {
		return err
	}

	if exists.ID == "" {
		return errors.New("data customer not found")
	}

	err = c.customerRepository.Delete(ctx, id)

	return err

}
