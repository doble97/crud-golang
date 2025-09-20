package repoimpl

import (
	"crud_project/domain/models"
	"crud_project/domain/repository"
	"crud_project/infrastructure/utils"
	"errors"
)

var data = []models.Customer{
	{ID: "sfa-ffw-111", Name: "Pruebas", Role: "Basic", Email: "pruebas@gmail.com", Phone: 666888999, Contacted: false},
	{ID: "sfa-ffw-222", Name: "Pruebas2", Role: "Premium", Email: "pruebas2@gmail.com", Phone: 666111999, Contacted: true},
	{ID: "sfa-ffw-333", Name: "Pruebas3", Role: "Basic", Email: "pruebas3@gmail.com", Phone: 666444555, Contacted: true},
}

type customerRepoMemory struct {
	customers []models.Customer
}

// CreateSeveral implements repository.CustomerRepository.
func (c *customerRepoMemory) CreateSeveral(customers []models.Customer) error {
	for key := range customers {
		customers[key].ID = utils.GenerateID()
	}
	c.customers = append(c.customers, customers...)
	return nil
}

// Create implements repository.CustomerRepository.
func (c *customerRepoMemory) Create(customer models.Customer) (*models.Customer, error) {
	customer.ID = utils.GenerateID()
	c.customers = append(c.customers, customer)
	return &customer, nil
}

// Delete implements repository.CustomerRepository.
func (c *customerRepoMemory) Delete(id string) error {
	sizeCustomers := len(c.customers)
	existCustomer := false
	for key, value := range c.customers {
		if value.ID == id {
			if key+1 == sizeCustomers {
				c.customers = c.customers[:key]
			} else {
				c.customers = append(c.customers[:key], c.customers[key:]...)
			}
			existCustomer = true
			break
		}
	}
	if existCustomer {
		return nil
	}
	return errors.New("user is not exist")
}

// GetAll implements repository.CustomerRepository.
func (c *customerRepoMemory) GetAll() []models.Customer {
	return c.customers
}

// GetById implements repository.CustomerRepository.
func (c *customerRepoMemory) GetById(id string) (*models.Customer, error) {
	for _, value := range c.customers {
		if value.ID == id {
			return &value, nil
		}
	}
	return nil, errors.New("user not found")
}

// Update implements repository.CustomerRepository.
func (c *customerRepoMemory) Update(id string, customer models.Customer) error {
	existCustomer := false
	customer.ID = id
	for key, value := range c.customers {
		if value.ID == id {
			c.customers[key] = customer
			existCustomer = true
			break
		}
	}
	if existCustomer {
		return nil
	}
	return errors.New("customer not found")
}

func NewCustomerRepoMemory() repository.CustomerRepository {
	return &customerRepoMemory{
		customers: data,
	}
}
