package handlers

import (
	"employee_system/data"

	"github.com/gorilla/mux"

	"log"

	"net/http"
)

// employees is a http.Handler
type Employees struct {
	l *log.Logger
}

// NewEmployee creates a products handler with the given logger
func NewEmployee(l *log.Logger) *Employees {

	return &Employees{l}
}

// ServeHTTP is the main entry point for the handler and staisfies the http.Handler
// interface

func (em *Employees) GetEmployee(rw http.ResponseWriter, r *http.Request) {
	em.l.Println("Handle Get Employee")
	r.Header.Add("content-type", "application/json")

	data.GetEmployee(rw)
}

func (em *Employees) AddEmployee(rw http.ResponseWriter, r *http.Request) {
	em.l.Println("Handle POST Employee")
	r.Header.Add("content-type", "application/json")

	DATA := &data.Employee{}

	err := DATA.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	data.AddEmployee(DATA, rw)
}

func (em *Employees) UpdateEmployees(rw http.ResponseWriter, r *http.Request) {
	em.l.Println("Handle PUT Employee")
	r.Header.Add("content-type", "application/json")
	vars := mux.Vars(r)

	DATA := &data.Employee{}

	err := DATA.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	data.UpdateEmployees(DATA, vars["id"], rw)

}
func (em *Employees) DeleteEmployee(rw http.ResponseWriter, r *http.Request) {
	em.l.Println("Handle Delete Employee")
	r.Header.Add("content-type", "application/json")
	vars := mux.Vars(r)

	data.DeleteEmployee(vars["id"], rw)

}

// getEmployees returns the products from the data store
