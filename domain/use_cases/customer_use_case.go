package usecases

import (
	"crud_project/domain/models"
	"crud_project/domain/repository"
)

type CustomerUseCase interface {
	GetCustomer(id string) (*models.Customer, error)
	GetAllCustomers() []models.Customer
	CreateCustomer(customer models.Customer) (*models.Customer, error)
	UpdateCustomer(id string, customer models.Customer) error
	DeleteCustomer(id string) error
	CreateCustomers(customers []models.Customer) error
}

type customerUseCase struct {
	repo repository.CustomerRepository
}

// CreateCustomers implements CustomerUseCase.
func (c *customerUseCase) CreateCustomers(customers []models.Customer) error {
	return c.repo.CreateSeveral(customers)
}

// CreateCustomer implements CustomerUseCase.
func (c *customerUseCase) CreateCustomer(customer models.Customer) (*models.Customer, error) {
	return c.repo.Create(customer)
}

// DeleteCustomer implements CustomerUseCase.
func (c *customerUseCase) DeleteCustomer(id string) error {
	return c.repo.Delete(id)
}

// GetAllCustomers implements CustomerUseCase.
func (c *customerUseCase) GetAllCustomers() []models.Customer {
	return c.repo.GetAll()
}

// GetCustomer implements CustomerUseCase.
func (c *customerUseCase) GetCustomer(id string) (*models.Customer, error) {
	return c.repo.GetById(id)
}

// UpdateCustomer implements CustomerUseCase.
func (c *customerUseCase) UpdateCustomer(id string, customer models.Customer) error {
	return c.repo.Update(id, customer)
}

func NewCustomerUseCase(repo repository.CustomerRepository) CustomerUseCase {
	return &customerUseCase{repo: repo}
}
