package repository

import "crud_project/domain/models"

type CustomerRepository interface {
	GetById(id string) (*models.Customer, error)
	GetAll() []models.Customer
	Create(customer models.Customer) (*models.Customer, error)
	Update(id string, customer models.Customer) error
	Delete(id string) error
	CreateSeveral(customers []models.Customer) error
}
