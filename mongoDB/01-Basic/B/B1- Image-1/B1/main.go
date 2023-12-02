// main.go

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	// "time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Person represents a simple data structure for our MongoDB documents
type Person struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

var client *mongo.Client

func main() {
	// Connect to MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	var err error
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// Initialize router
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{name}", GetPerson).Methods("GET")
	router.HandleFunc("/people", CreatePerson).Methods("POST")

	// Start server
	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// GetPeople returns a list of people from the database
func GetPeople(w http.ResponseWriter, r *http.Request) {
	var people []Person
	collection := client.Database("test").Collection("people")
	cur, err := collection.Find(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var person Person
		err := cur.Decode(&person)
		if err != nil {
			log.Fatal(err)
		}
		people = append(people, person)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(people)
}

// GetPerson returns a single person by name from the database
func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]

	var person Person
	collection := client.Database("test").Collection("people")
	err := collection.FindOne(context.Background(), Person{Name: name}).Decode(&person)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(person)
}

// CreatePerson adds a new person to the database
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person Person
	json.NewDecoder(r.Body).Decode(&person)

	collection := client.Database("test").Collection("people")
	_, err := collection.InsertOne(context.Background(), person)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(person)
}
