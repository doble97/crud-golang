# crud-golang
Project for udacity course. A small REST API built in Go to manage clients (Client) with basic CRUD (Create, Read, Update, Delete) operations. The application uses an in-memory repository and allows you to manage both individual and multiple clients at the same time.
# Download the dependencies defined in go.mod
go mod tidy

---

## Features

- **GET `/customers`** – Retrieve all customers.  
- **GET `/customers/{id}`** – Retrieve a customer by ID.  
- **POST `/customers`** – Create a new customer.  
- **POST `/customers-several`** – Create multiple customers at once.  
- **PUT `/customers/{id}`** – Update an existing customer.  
- **DELETE `/customers/{id}`** – Delete a customer by ID. 

---

## Example cURL Requests

### Get all customers
