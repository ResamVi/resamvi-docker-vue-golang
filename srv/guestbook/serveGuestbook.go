package main

import (
	"bytes"
	"html/template"
	"log"
	"fmt"
	"net/http"
	"strings"
	
	"gopkg.in/mgo.v2"
)

type entry struct {
	ID  	int32
	IP		string
	Date  	string
	Name  	string
	Email	string
	Message	string
}

var collection *mgo.Collection

func main() {

	session := connectDB()
	defer session.Close()

	fmt.Println("Setup serveBlog on port 8081")

	// Setup handlers
	http.HandleFunc("/", handler)
	http.HandleFunc("/send", sendEntry)

	// Listen
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handler(write http.ResponseWriter, reader *http.Request) {

	// Print client information that connected
	fmt.Println(formatRequest(reader) + "\n")

	// All the guestbook entries will be stored in here
	var entries []entry

	// Get all the entries from the database
	err := collection.Find(nil).All(&entries)
	if err != nil {
		fmt.Println(err)
	}

	t := template.Must(template.ParseFiles("../../build/gaestebuch.html"))
	
	buffer := new(bytes.Buffer)
	
	err = t.Execute(buffer, entries)
	
	if err != nil {
		fmt.Println(err)
	}
	
	write.Write(buffer.Bytes())
}

func connectDB() *mgo.Session {

	fmt.Println("Connecting to database")
	
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		fmt.Println("Could not connect to database", err)
		return nil
	}

	collection = session.DB("entries").C("guestbook")

	fmt.Println("Connection established.")
	return session
}

func sendEntry(write http.ResponseWriter, reader *http.Request) {
	
	// Print connection
	fmt.Println(formatRequest(reader))
	
	// Connect to database
	fmt.Println("Connecting to database")
	
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		fmt.Println("Could not connect to database", err)
	}
	
	collection = session.DB("entries").C("guestbook")
	
	fmt.Println("Connection established.")
	
	// Ready to read form input
	reader.ParseForm()
	
	err = collection.Insert(entry{ID: 5, IP: "TODO", Date: "TODO", Name: reader.Form.Get("name"),
		Email: reader.Form.Get("email"), Message: reader.Form.Get("content")})
		
	if err != nil {
		fmt.Println(err)
	}
	
	write.Write([]byte("Message Sent"))
}

// formatRequest generates ascii representation of a request
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