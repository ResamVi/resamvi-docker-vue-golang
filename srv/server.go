package main

import (
	"fmt"
	"log"
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
)

var id int = 0

type entry struct {
	Number  int32
	Title	string
	Date  	string
	Text  	string
}

func handler(write http.ResponseWriter, reader *http.Request) {
	
	// Log
	fmt.Printf("%s %s\nHost: %s %s\n\n", reader.Method, reader.Proto, reader.URL.Host, reader.URL.Path)
	
	// Connect to mongodb
	session, err := mgo.Dial("localhost:27017")

	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Connect database
	collection := session.DB("resamvi").C("entries")

	// Select entry
	result := entry{}
	err = collection.Find(bson.M{"number": id}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	// Prepare response
	write.Header().Set("Access-Control-Allow-Origin", "*")
	write.Header().Set("Content-Type", "application/json;charset=utf-8")
	out, err := json.Marshal(result)

	write.Write(out)

	id = 0

	// OK
	write.WriteHeader(http.StatusOK)
}

func main() {

	fmt.Println("Setup server on port 8080")

	// Setup handlers
	http.HandleFunc("/", handler)

	// Listen
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
