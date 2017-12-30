	package main

import (
	"fmt"
	"bytes"
	"net/http"
	"text/template"
	"log"
	"strings"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fatih/color"
	"strconv"
)

type singleEntry struct {
	Number  int32	// TODO: Necessary?
	URL		string
	Title	string
	Date  	string
	Text  	string
}

func main() {

	fmt.Println("Setup serveBlog on port 8082")

	// Setup handlers
	http.HandleFunc("/", handler)

	// Listen
	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handler(write http.ResponseWriter, reader *http.Request) {

	//fmt.Println(formatRequest(reader))
	color.Red(reader.URL.Path[1:])

	// Print client information that connected
	fmt.Println(formatRequest(reader) + "\n")

	// Connect to database
	fmt.Println("Connecting to database")

	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		fmt.Println("Could not connect to database", err)
	}

	collection := session.DB("entries").C("blog")

	fmt.Println("Connection established.")

	// Get all the entries from the database
	entry := singleEntry{}
	
	entryNumber, _ := strconv.Atoi(reader.URL.Path[1:])
	
	err = collection.Find(bson.M{"number": entryNumber}).One(&entry)
	if err != nil {
		fmt.Println("COULD NOT FIND:", err)
	}

	t := template.Must(template.ParseFiles("../../build/entry.html"))
	
	buffer := new(bytes.Buffer)
	
	err = t.Execute(buffer, entry)
	
	if err != nil {
		fmt.Println(err)
	}
	
	write.Write(buffer.Bytes())

}

func formatRequest(r *http.Request) string {
	// Create return string
	var request []string
   
	// Add the request string
	url := fmt.Sprintf("\n%v %v %v", r.Method, r.URL, r.Proto)
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
	
	// If this is a POST, add post data
	if r.Method == "POST" {
	   r.ParseForm()
	   request = append(request, "\n")
	   request = append(request, r.Form.Encode())
	} 
   
	 // Return the request as a string
	 return strings.Join(request, "\n")
}