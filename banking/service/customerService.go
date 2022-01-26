package service

import "github.com/federicoalonso/go-microservices-udemy/banking/domain"

type CustomerService interface {
	GetAll() ([]domain.Customer, error)
}

type DefaultCustomerService struct {
	repository domain.CustomerRepository
}

func (s DefaultCustomerService) GetAll() ([]domain.Customer, error) {
	return s.repository.FindAll()
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository: repository}
}
