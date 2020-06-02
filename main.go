package main

import (
	"employee_system/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {

	l := log.New(os.Stdout, "employee_system", log.LstdFlags)

	// create the handlers
	ph := handlers.NewEmployee(l)

	// create a new serve mux and register the handlers
	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ph.GetEmployee)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id}", ph.UpdateEmployees)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ph.AddEmployee)
	deleteRouter := sm.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/{id}", ph.DeleteEmployee)

	// create a new server
	l.Println("Starting server on port 5000")

	err := http.ListenAndServe(":5000", sm)
	if err != nil {
		l.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}

}
