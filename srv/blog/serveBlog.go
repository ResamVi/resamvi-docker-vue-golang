package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	
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

func main() {

	fmt.Println("Setup serveBlog on port 8080")

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

func handler(write http.ResponseWriter, reader *http.Request) {
	
	fmt.Println(formatRequest(reader))

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

	// Return most recent entry first. 
	// Entries (i.e. convert) are numbered from 1 (the oldest) upwards to total
	entryNumber := total + 1 - convert

	// Select entry, as long as there are entries left
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

	collection = session.DB("entries").C("blog")

	fmt.Println("Connection established.")
	return session
}

// formatRequest generates ascii representation of a request
func formatRequest(r *http.Request) string {
	// Create return string
	var request []string
   
	// Add the request string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)
   
	// Add the host
	request = append(request, fmt.Sprintf("Host: %v", r.Host))
   
	// Loop through headers
	for name, headers := range r.Header {
	  name = strings.ToLower(name)
	  for _, h := range headers {
		request = append(request, fmt.Sprintf("%v: %v", name, h))
	  }
	}
   
	 // Return the request as a string
	 return strings.Join(request, "\n")
}