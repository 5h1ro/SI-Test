package service

import (
	"customer/internal/entity"
	"customer/internal/repository"
	"fmt"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type CustomerService interface {
	Store(any) (any, error)
}

type customerService struct {
	customerRepository repository.CustomerRepository
}

func NewCustomerService(customerRepository repository.CustomerRepository) CustomerService {
	return &customerService{
		customerRepository: customerRepository,
	}
}

func incrementLastNumber(input string) (string, error) {
	parts := strings.Split(input, ".")
	lastPart := parts[len(parts)-1]
	number, err := strconv.Atoi(lastPart)
	if err != nil {
		return "", err
	}
	number++
	parts[len(parts)-1] = strconv.Itoa(number)
	return strings.Join(parts, "."), nil
}

func (s customerService) Store(parentID any) (any, error) {
	lastData, err := s.customerRepository.FindOneBy(map[string]interface{}{
		"parent_id": fmt.Sprintf("%v", parentID),
	})
	var name, nomor string
	var parentId uuid.NullUUID
	if err != nil {
		parent, err := s.customerRepository.FindOneBy(map[string]interface{}{
			"id": fmt.Sprintf("%v", parentID),
		})
		if err != nil {
			lastParentData, err := s.customerRepository.FindOneBy(map[string]interface{}{
				"parent_id": nil,
			})
			if err != nil {
				nomor = "1"
			} else {
				nomor, err = incrementLastNumber(lastParentData.Nomor)
				if err != nil {
					return nil, err
				}
				if err != nil {
					return nil, err
				}
			}
			parentId = uuid.NullUUID{
				Valid: false,
			}
			name = fmt.Sprintf("Data %v", nomor)
		} else {
			name = fmt.Sprintf("%v.1", parent.Name)
			nomor = fmt.Sprintf("%v.1", parent.Nomor)
			parentId = uuid.NullUUID{
				UUID:  parent.ID,
				Valid: true,
			}
		}
	} else {
		nextNumber, err := incrementLastNumber(lastData.Nomor)
		if err != nil {
			return nil, err
		}
		name = fmt.Sprintf("Data %v", nextNumber)
		nomor = fmt.Sprintf("%v", nextNumber)
		parentId = uuid.NullUUID{
			UUID:  lastData.ParentId.UUID,
			Valid: true,
		}
	}
	customer, err := s.customerRepository.Store(entity.Customer{
		ParentId: parentId,
		Name:     name,
		Nomor:    nomor,
	})
	if err != nil {
		return nil, err
	}
	return customer, nil
}
