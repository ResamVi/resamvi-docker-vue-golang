package main

import (
	"github.com/fatih/color"
	"net/http"
	"gopkg.in/mgo.v2"
)

/* 
 * Port is part of the network address (URI)
 * on which this server listens to.
 */
const port = "8080"

/*
 * Global Variable.
 * 
 * Once the connection to the database has ben setup
 * access to the collections inside the database are made
 * by each service of the server. 
 */
var session *mgo.Session

/*
 * Start the server.
 *
 * This requires a connection to the MongoDB Database to access
 * blog entries and guestbook entries (connectDB in util.go)
 *
 * Each service logs its messages in a different color on the console
 * 
 * RED:		Initializing server (main.go/util.go)
 * YELLOW:	Incoming client request information (util.go)
 * GREEN:	Endless loading blog entries on index.html (serveBlog.go)
 * BLUE:	Loading guestbook (serveGuestbook.go)
 * CYAN:	Serving single page entries on direct access (servePage.go)
 */
func main() {

	// Start
	color.Red("Setup server on port " + port)

	// Connect DB
	session = connectDB()

	// Close when this function is terminated
	defer session.Close()

	// Setup handlers
	http.HandleFunc("/blog", serveBlog)
	http.HandleFunc("/gaestebuch", serveGuestbook)
	http.HandleFunc("/send", sendEntry)
	http.HandleFunc("/entry/", servePage)

	// Listen
	err := http.ListenAndServe(":" + port, nil)
	if err != nil {
		color.Red("Could not setup handlers: " + err.Error())
	}
}