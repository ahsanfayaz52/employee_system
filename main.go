package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"


)
type userData struct {
	Name    string
	Address string
}


func homePage(writer http.ResponseWriter,request *http.Request) {
	switch request.Method {
	case "GET":
		http.ServeFile(writer, request, "form.html")
	case "POST":

		name := request.FormValue("name")
		address := request.FormValue("address")
		data := userData{Name: name, Address:address}
		result,err := json.Marshal( data)
		if (err != nil){
			fmt.Fprintf(writer,"Error occured" ,err )
		}
		fmt.Fprintf(writer,string(result ))


	default:
		fmt.Fprintf(writer, "Sorry, only GET and POST methods are supported.")

	}



}



func handleRequests(){
	http.HandleFunc("/",homePage)
	log.Fatal(http.ListenAndServe(":5000",nil))

}
func main(){
	handleRequests();

}

