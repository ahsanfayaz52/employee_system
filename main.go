package main

import (

	"fmt"
	"log"
	"net/http"


)


func homePage(writer http.ResponseWriter,request *http.Request) {


		fmt.Fprintf(writer,"hello world" )




}



func handleRequests(){
	http.HandleFunc("/",homePage)
	log.Fatal(http.ListenAndServe(":5000",nil))

}
func main(){
	handleRequests();

}

