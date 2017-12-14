package main

import "fmt"
import "net/http"

func handler(write http.ResponseWriter, reader *http.Request) {
	write.Header().Set("Access-Control-Allow-Origin", "http://localhost:9001");
	fmt.Println("Send received");
}

func main() {
	
	// Connect to mongodb
	/*session, err := mgo.Dial("localhost")

	if err != nil {
		panic(err)
	}
	defer session.Close();

	// Select database
	collection := session.DB("resamvi").C("entries")

	if err != nil {
		log.Fatal(err);
	}*/

	http.HandleFunc("/", handler);
	http.ListenAndServe(":8080", nil);
}