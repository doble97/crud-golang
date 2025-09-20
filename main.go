package main

import (
	"crud_project/domain/models"
	usecases "crud_project/domain/use_cases"
	repoimpl "crud_project/infrastructure/repo_impl"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

var custUseCase = usecases.NewCustomerUseCase(repoimpl.NewCustomerRepoMemory())

func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(custUseCase.GetAllCustomers())
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(errors.New("incorrect error"))
		return
	}
	result, err := custUseCase.GetCustomer(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "ko",
			"message": err.Error(),
		})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
	}
}

func addCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var tempCustomer models.Customer
	body, errBody := io.ReadAll(r.Body)
	if errBody != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "ko",
			"message": "Error in body",
		})
	}
	defer r.Body.Close()
	if err := json.Unmarshal(body, &tempCustomer); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "ko",
			"message": "Incorrect data",
		})
		return
	}
	result, err := custUseCase.CreateCustomer(tempCustomer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "ko",
			"message": err.Error(),
		})
	} else {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(result)
	}

}
func updateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var tempCustomer models.Customer
	body, errBody := io.ReadAll(r.Body)
	if errBody != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "ko",
			"message": "Incorrect data",
		})
		return
	}
	if err := json.Unmarshal(body, &tempCustomer); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "ko",
			"message": "Incorrect data",
		})
		return
	}
	vars := mux.Vars(r)
	if _, ok := vars["id"]; !ok {
		w.WriteHeader(http.StatusBadRequest)
	}
	err := custUseCase.UpdateCustomer(vars["id"], tempCustomer)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "ko",
			"message": err.Error(),
		})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "ok",
		"message": "Updated user",
	})
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	var idCustomer string
	if _, ok := vars["id"]; !ok {
		// Contingence for  pass tests
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) < 3 || parts[2] == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"status":  "ko",
				"message": "Incorrect request",
			})
			return
		}
		idCustomer = parts[2]

	}
	result := custUseCase.DeleteCustomer(idCustomer)
	if result != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "ko",
			"message": result.Error(),
		})
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "ok",
		"message": "Deleted user",
	})

}

func addCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")
	data, errData := io.ReadAll(r.Body)
	if errData != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "ko",
			"message": "Error in server",
		})
		return
	}
	defer r.Body.Close()
	var tempCustomers []models.Customer
	errData = json.Unmarshal(data, &tempCustomers)
	if errData != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "ko",
			"message": "Error in body",
		})
	}
	custUseCase.CreateCustomers(tempCustomers)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "ok",
		"message": "Created users",
	})
}
func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/customers", getCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	router.HandleFunc("/customers", addCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PUT")
	router.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")
	router.HandleFunc("/customers-several", addCustomers).Methods("POST")
	router.PathPrefix("/").Handler(fileServer)
	fmt.Println("LISTENING ON PORT 3333")
	http.ListenAndServe(":3333", router)
}
