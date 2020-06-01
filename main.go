package main

import (

	"employee_system/handlers"

	"log"
	"net/http"
	"os"

)


func main() {



	l := log.New(os.Stdout, "employee_system", log.LstdFlags)

	// create the handlers
	ph := handlers.NewEmployee(l)

	// create a new serve mux and register the handlers
	sm := http.NewServeMux()
	sm.Handle("/", ph)

	// create a new server
	l.Println("Starting server on port 5000")

	err :=http.ListenAndServe(":5000",sm)
	if err != nil {
		l.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}





}