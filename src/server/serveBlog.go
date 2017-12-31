package main

import (
	"encoding/json"
	"fatih/color"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"strconv"
	"net/http"
)

/* 
 * BlogEntry is a single entry written on the blog
 * served on the start page by endless scrolling (main.go)
 * or as a single page (servePage.go)
 */
type blogEntry struct {
	Number int32
	Title  string
	Date   string
	Text   string
}

/*
 * The serveBlog handle is responsible for the endless loading
 * of entries.
 *
 * It handles HTTP requests sent by vue on a scroll-down event.
 *
 * The client has reached the end of the post and a new
 * entry has to be dynamically loaded next.
 */
func serveBlog(write http.ResponseWriter, reader *http.Request) {

	// Log details about the client to console
	formatRequest(reader)
	
	/*
	 * Handle CORS preflight requests.
	 * A CORS preflight request is a CORS request that checks to see if the CORS protocol is understood.
	 * TODO: STill necessary?
	 */
	if reader.Method == "OPTIONS" {

		write.Header().Set("Access-Control-Allow-Origin", "*")
		write.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		write.Header().Set("Content-Type", "application/json;charset=utf-8")

		// OK
		write.WriteHeader(http.StatusOK)

		return
	}

	// Read request body
	body, err := ioutil.ReadAll(reader.Body)
	if err != nil {
		color.Green("Could not read HTTP request body: " + err.Error())
	}

	// Convert body to int
	convert, err := strconv.Atoi(string(body))
	if err != nil {
		color.Green("Not a valid body/could not convert to number: " + err.Error())
	}

	// Get the collection of blog entries 
	entries := session.DB("entries").C("blog")
	
	// Get latest entry number
	total, err := entries.Count()
	if err != nil {
		color.Green("Could not count amount of entries: " + err.Error())
	}

	/*
	 * Return most recent entry first.
	 * Numbering of entries (i.e. convert) are numbered from 1 (the oldest) upwards to total
	 */
	entryNumber := total + 1 - convert

	// Select entry, as long as there are entries left
	result := blogEntry{}
	if entryNumber > 0 {

		err = entries.Find(bson.M{"number": entryNumber}).One(&result)
		if err != nil {
			color.Green("The entry with the given number could not be found: " + err.Error())
		}
	
	} else {
		color.Green("Reached end of blog") // TODO: Test this
	}

	// Prepare response
	write.Header().Set("Access-Control-Allow-Origin", "*") // TODO: Do not allow all headers!
	write.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	write.Header().Set("Content-Type", "application/json;charset=utf-8")
	
	// json-ify the result
	out, err := json.Marshal(result)
	if err != nil {
		color.Green("Failed at converting the entry to JSON: " + err.Error())
	}

	// Send back
	write.Write(out)
}