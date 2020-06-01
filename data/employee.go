package data

import (
	"context"
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"

	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"

)

// Employee defines the structure for an API product
type Employee struct {
	ID  primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name          string  `json:"name,omitempty" bson:"name,omitempty" `
	//Age           int   `json:"age,omitempty" bson:"age,omitempty" `
	//Salary        float32 `json:"salary,omitempty" bson:"salary,omitempty" `
	 Address       string  `json:"address,omitempty" bson:"address,omitempty" `
	 Phone         string  `json:"phone,omitempty" bson:"phone,omitempty" `

}

// Employees is a collection of Product
type Employees []*Employee

// ToJSON serializes the contents of the collection to JSON
// NewEncoder provides better performance than json.Unmarshal as it does not
// have to buffer the output into an in memory slice of bytes
// this reduces allocations and the overheads of the service
//
// https://golang.org/pkg/encoding/json/#NewEncoder
func (em *Employees) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(em)
}
func (em *Employee) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(em)
}



// GetEmployees returns a list of products
func GetEmployee( w http.ResponseWriter ) {
	client,err:= mongo.NewClient(options.Client().ApplyURI("mongodb+srv://ahsanfayaz52:Ahsan_khan425@employeesystem-mwxmj.mongodb.net/test?retryWrites=true&w=majority"))
    if (err!=nil){
    	http.Error(w, "Connection probelm", http.StatusBadRequest)
	}
	ctx,_ := context.WithTimeout(context.Background(),10*time.Second)
	err = client.Connect(ctx)
	if (err!=nil){
		log.Fatal(err)
	}

	collection := client.Database("employeeSystemdb" ).Collection("employees")

    var people []Employee
	cursor,err:= collection.Find(ctx,bson.M{})
	if (err!=nil){
		http.Error(w, "database error", http.StatusBadRequest)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx){
		var Employee Employee
		cursor.Decode(&Employee)


		people  = append(people,Employee)

	}
	if err:= cursor.Err(); err !=nil {
		http.Error(w, "data cannot be shown", http.StatusBadRequest)
	}
	json.NewEncoder(w).Encode(people)








}

func AddEmployee(em *Employee,w http.ResponseWriter ) {
	client,err:= mongo.NewClient(options.Client().ApplyURI("mongodb+srv://ahsanfayaz52:Ahsan_khan425@employeesystem-mwxmj.mongodb.net/test?retryWrites=true&w=majority"))
	if (err!=nil){
		log.Fatal(err)
	}



   ctx,_ := context.WithTimeout(context.Background(),10*time.Second)
	err = client.Connect(ctx)
	if (err!=nil){
		log.Fatal(err)
	}

	collection  := client.Database("employeeSystemdb" ).Collection("employees")
   result,_:= collection.InsertOne(ctx,em)
   json.NewEncoder(w).Encode(result)

}





