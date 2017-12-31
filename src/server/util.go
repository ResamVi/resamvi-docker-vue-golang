package main

import(
	"gopkg.in/mgo.v2"
	"fatih/color"
	"fmt"
	"net/http"
	"strings"
)

// DatabaseAddress is the location
// to access the MongoDB database containing
// blog entries and guestbook entries
const databaseAddress = "localhost:27017"

/*
 * Connect to the MongoDB database
 * on port 27017.
 *
 * The database is called "entries" and
 * has two collections: "blog" and "guestbook"
 */
func connectDB() *mgo.Session {

	color.Red("Connecting to database on " + databaseAddress)

	// Connect to database
	session, err := mgo.Dial(databaseAddress)
	if err != nil {
		color.Red("Could not connect to database:", err)
	}

	// Return collection
	color.Red("Connection established.")

	return session
}

// formatRequest generates ascii representation of a request
func formatRequest(r *http.Request) {
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
	color.Yellow(strings.Join(request, "\n") + "\n")
}
