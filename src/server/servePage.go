package main

import (
	"bytes"
	"fatih/color"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
	"text/template"
)

/* 
 * Same as blogEntry (serveBlog.go)
 *
 * singleEntry is an entry written on the blog
 * served on the start page by endless scrolling (main.go)
 * or as a single page (servePage.go)
 */
type singleEntry struct {
	Number int32
	Title  string
	Date   string
	Text   string
}

/*
 * The servePage handle is responsible for direct queries to a specific
 * entry, say by typing "resamvi.de/apfelkuchen" to get specific information.
 *
 * When redirected the name of entry is converted to its ID
 * which is used here to serve the specific entry
 */
func servePage(write http.ResponseWriter, reader *http.Request) {

	// Print client information that connected
	formatRequest(reader)
	
	// Get the list of entries made
	collection := session.DB("entries").C("blog")

	// Get all the entries from the database
	entry := singleEntry{}

	/*
	 * Convert the URL to int
	 *
	 * When redirected the name of entry is converted to its ID
	 * which is used here to serve the specific entry
	 */ 
	entryNumber, err := strconv.Atoi(string(reader.URL.Path[len(reader.URL.Path)-1]))
	if err != nil {
		color.Cyan("Could not convert URL to int: " + err.Error())
	}

	// Look for the entry in database
	err = collection.Find(bson.M{"number": entryNumber}).One(&entry)
	if err != nil {
		color.Cyan("Could not find entry with given number: " + err.Error())
	}

	// The entry.html file is the given template where entry data inserted
	t := template.Must(template.ParseFiles("../../build/entry.html"))

	// Data inserted into the template is saved here
	buffer := new(bytes.Buffer)
	err = t.Execute(buffer, entry)
	if err != nil {
		fmt.Println(err)
	}

	// Serve back html page
	write.Write(buffer.Bytes())
}
