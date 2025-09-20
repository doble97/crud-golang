# crud-golang
Project for udacity course. A small REST API built in Go to manage clients (Client) with basic CRUD (Create, Read, Update, Delete) operations. The application uses an in-memory repository and allows you to manage both individual and multiple clients at the same time.

---

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
curl -X GET http://localhost:3333/customers


### Get customer by id
curl -X GET http://localhost:3333/customers/sfa-ffw-111


### Create  a new customer
curl -X POST http://localhost:3333/customers -H "Content-Type: application/json" \
-d '{"name":"John Doe","role":"Basic","email":"john@mail.com","phone":12345,"contacted":true}'


### Create several customers
curl -X POST http://localhost:3333/customers-several -H "Content-Type: application/json" \
-d '[{"name":"Jane","role":"Premium","email":"jane@mail.com","phone":67890,"contacted":false},
{"name":"Carlos","role":"Basic","email":"carlos@mail.com","phone":55566,"contacted":true}]'


### Update a customer
curl -X PUT http://localhost:3333/customers/c1 -H "Content-Type: application/json" \
-d '{"name":"John Updated","role":"Premium","email":"john@mail.com","phone":12345,"contacted":true}'


### Delete a customer
curl -X DELETE http://localhost:3333/customers/c1

**Note:** This API use a slice for store customers. "deleteCustomer" has modified because test was failing.
