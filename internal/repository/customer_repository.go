package repository

import (
	"customer/internal/entity"
	"errors"
	"strings"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	FindOneBy(c map[string]interface{}) (entity.Customer, error)
	Store(c entity.Customer) (entity.Customer, error)
}

type customerRepository struct {
	connection *gorm.DB
}

func NewCustomerRepository(dbConn *gorm.DB) CustomerRepository {
	return &customerRepository{
		connection: dbConn,
	}
}

func (r *customerRepository) FindOneBy(c map[string]interface{}) (entity.Customer, error) {
	var customer entity.Customer
	if result := r.connection.Where(c).Order("string_to_array(nomor, '.')::int[] DESC").Limit(1).Find(&customer); result.Error != nil {
		return entity.Customer{}, result.Error
	}

	if customer.ID.String() == "00000000-0000-0000-0000-000000000000" {
		return entity.Customer{}, errors.New("Customer not found")
	}

	customer = entity.Customer{
		ID:       customer.ID,
		ParentId: customer.ParentId,
		Name:     customer.Name,
		Nomor:    strings.TrimSpace(customer.Nomor),
	}

	return customer, nil
}

func (r customerRepository) Store(c entity.Customer) (entity.Customer, error) {
	customer := entity.Customer{
		ParentId: c.ParentId,
		Name:     c.Name,
		Nomor:    c.Nomor,
	}
	if result := r.connection.Create(&customer); result.Error != nil {
		return entity.Customer{}, result.Error
	}

	return customer, nil
}
