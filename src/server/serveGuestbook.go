package main

import (
	"fmt"
	"bytes"
	"fatih/color"
	"html/template"
	"net/http"
	"strings"
	"time"
)

/*
 * A guestbook entry in the database contains more information than actually
 * displayed on the page (date, name, message only)
 */
type guestEntry struct {
	ID      int
	IP      string
	Date    string
	Name    string
	Email   string
	Message string
}

/*
 * The serveGuestbook handle is responsible for reading the
 * data from the MongoDB database and insert it into the html page.
 *
 * The page is given with a template (src/gaestebuch.html)
 */
func serveGuestbook(write http.ResponseWriter, reader *http.Request) {

	// Log details about the client to console
	formatRequest(reader)

	// All the guestbook entries will be stored in here
	var entries []guestEntry

	// Get the collection of blog entries 
	collection := session.DB("entries").C("guestbook")

	// Get all the entries from the database
	err := collection.Find(nil).All(&entries)
	if err != nil {
		color.Blue("Could not get all guestbook entries: " + err.Error())
	}

	// gaestebuch.html is a go template (TODO: Change path when deploying)
	t := template.Must(template.ParseFiles("../../build/gaestebuch.html"))

	// With given template and the associated data filled into it, save it to this buffer
	buffer := new(bytes.Buffer)

	// Put data into template and save it in buffer
	err = t.Execute(buffer, entries)
	if err != nil {
		color.Blue("Error while filling template with values: " + err.Error())
	}

	// Send back the filled in HTML file
	write.Write(buffer.Bytes())
}

/*
 * sendEntry stores the data given in the form input to the database 
 * to be displayed in the guestbook
 *
 * sendEntry is used by gaestebuch.html when the submit button is pressed
 * via a POST request to /send
 */
func sendEntry(write http.ResponseWriter, reader *http.Request) {

	// Print connection
	formatRequest(reader)

	// Get the collection of blog entries 
	collection := session.DB("entries").C("guestbook")

	// Ready to read form input
	reader.ParseForm()

	// Log given information
	var log []string

	// Current date
	y, m, d := time.Now().UTC().Date()
	date := fmt.Sprintf("%d %s, %d", d, m.String(), y)

	// Calculate ID
	id, err := collection.Count()
	if err != nil {
		color.Blue("Could not count guestbook collection: " + err.Error())
	}

	// List all form data / calculated data
	log = append(log, "ID: " + string(id+1))
	log = append(log, "IP: " + reader.RemoteAddr)
	log = append(log, "Date: " + date)
	log = append(log, "Name: " + reader.Form.Get("name"))
	log = append(log, "email: " + reader.Form.Get("email"))
	log = append(log, "Message: " + reader.Form.Get("content"))
	color.Blue(strings.Join(log, "\n") + "\n")
	
	// Insert form data into database TODO: save for injection?
	data := guestEntry{ID: id+1, IP: reader.RemoteAddr, Date: date, Name: reader.Form.Get("name"), Email: reader.Form.Get("email"), Message: reader.Form.Get("content")}
	
	err = collection.Insert(data)
	if err != nil {
		color.Blue("Writing entry to database failed: " + err.Error())
	}

	// TODO: Write better confirmation message
	write.Write([]byte("Message Sent"))
}