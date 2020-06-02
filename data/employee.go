package data

import (
	"context"
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"net/http"
	"time"
)

func fm() {}

// Employee defines the structure for an API product
type Employee struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name,omitempty" bson:"name,omitempty" `
	//Age           int   `json:"age,omitempty" bson:"age,omitempty" `
	//Salary        float32 `json:"salary,omitempty" bson:"salary,omitempty" `
	Address string `json:"address,omitempty" bson:"address,omitempty" `
	Phone   string `json:"phone,omitempty" bson:"phone,omitempty" `
}

// Employees is a collection of employee
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

// GetEmployees returns a list of employees
func GetEmployee(w http.ResponseWriter) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://ahsanfayaz:ahsan123@employeesystem-mwxmj.mongodb.net/test?retryWrites=true&w=majority",
	))
	if err != nil {
		fmt.Println("error pccored")
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("employeeSystemdb").Collection("employees")

	var people []Employee
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, "database error", http.StatusBadRequest)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var Employee Employee
		err = cursor.Decode(&Employee)
		if err != nil {
			log.Fatal(err)
		}
		people = append(people, Employee)

	}
	if err := cursor.Err(); err != nil {
		http.Error(w, "data cannot be shown", http.StatusBadRequest)
	}
	err = json.NewEncoder(w).Encode(people)
	if err != nil {
		log.Fatal(err)
	}

}

func AddEmployee(em *Employee, w http.ResponseWriter) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://ahsanfayaz:ahsan123@employeesystem-mwxmj.mongodb.net/test?retryWrites=true&w=majority",
	))
	if err != nil {
		fmt.Println("error pccored")
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("employeeSystemdb").Collection("employees")
	result, _ := collection.InsertOne(ctx, em)
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		log.Fatal(err)
	}

}

func UpdateEmployees(em *Employee, ID string, w http.ResponseWriter) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://ahsanfayaz:ahsan123@employeesystem-mwxmj.mongodb.net/test?retryWrites=true&w=majority",
	))
	if err != nil {
		fmt.Println("error pccored")
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("employeeSystemdb").Collection("employees")
	id, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": id}}
	update := bson.M{"$set": em}

	result, _ := collection.UpdateOne(
		ctx,
		filter,
		update,
	)
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteEmployee(ID string, w http.ResponseWriter) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://ahsanfayaz:ahsan123@employeesystem-mwxmj.mongodb.net/test?retryWrites=true&w=majority",
	))
	if err != nil {
		fmt.Println("error pccored")
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("employeeSystemdb").Collection("employees")
	id, _ := primitive.ObjectIDFromHex(ID)
	result, _ := collection.DeleteOne(
		ctx,
		bson.M{"_id": id},
	)
	fmt.Printf("DeleteOne removed %v document(s)\n", result.DeletedCount)
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		log.Fatal(err)
	}
}
