package handlers

import (
	"employee_system/data"
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
func (e *Employees) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet{
		e.getEmployee(rw, r)
		return
	}
	if r.Method == http.MethodPost {
		e.addEmployee(rw, r)
		return
	}

	// catch all
	// if no method is satisfied return an error

	rw.WriteHeader(http.StatusMethodNotAllowed)

}

func (p *Employees) getEmployee(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle Get Product")
	r.Header.Add("content-type","application/json")

	data.GetEmployee( rw)
}

func (p *Employees) addEmployee(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")
	r.Header.Add("content-type","application/json")

	prod := &data.Employee{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

   data.AddEmployee(prod,rw)
}

// getEmployees returns the products from the data store
