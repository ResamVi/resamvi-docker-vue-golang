package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
	"io/ioutil"
)

type entry struct {
	Number  int32	// TODO: Necessary?
	Title	string
	Date  	string
	Text  	string
}

var collection *mgo.Collection

func handler(write http.ResponseWriter, reader *http.Request) {
	
	/* 
	 * Handle CORS preflight requests.
	 * A CORS preflight request is a CORS request that checks to see if the CORS protocol is understood.
	 */
	if reader.Method == "OPTIONS" {
		
		write.Header().Set("Access-Control-Allow-Origin", "*")
		write.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		write.Header().Set("Content-Type", "application/json;charset=utf-8")
		
		// OK
		write.WriteHeader(http.StatusOK)

		return
	}
	
	// Logging (read request body)
	body, err := ioutil.ReadAll(reader.Body)
	if err != nil {
		fmt.Println(err)
	}

	convert, err := strconv.Atoi(string(body))
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Printf("%s, COUNT: %d\n", reader.Method, convert)

	// Get latest entry number
	total, err := collection.Count()
	if err != nil {
		fmt.Println(err)
	}

	entryNumber := total + 1 - convert

	// Select entry
	result := entry{}
	if entryNumber > 0 {
		
		err = collection.Find(bson.M{"number": entryNumber}).One(&result)
		if err != nil {
			log.Fatal(err)
		}
	}
	
	// Prepare response
	write.Header().Set("Access-Control-Allow-Origin", "*")
	write.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	write.Header().Set("Content-Type", "application/json;charset=utf-8")
	out, err := json.Marshal(result)

	write.Write(out)
}

func connectDB() *mgo.Session {

	fmt.Println("Connecting to database")
	
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		fmt.Println("Could not connect to database", err)
		return nil
	}

	collection = session.DB("resamvi").C("entries")

	fmt.Println("Connection established.")
	return session
}

func main() {

	fmt.Println("Setup server on port 8080")

	session := connectDB()
	defer session.Close()
	
	// Setup handlers
	http.HandleFunc("/", handler)

	// Listen
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}


}
