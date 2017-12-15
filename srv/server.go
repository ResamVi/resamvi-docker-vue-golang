package main

import "fmt"
import "net/http"
import "log"
import "gopkg.in/mgo.v2"
import "time"

func handler(write http.ResponseWriter, reader *http.Request) {
	write.Header().Set("Access-Control-Allow-Origin", "nginx:9001");
	fmt.Println("Send received");
}

func main() {
	
	fmt.Println("Trying localhost:27017")

	for i := 0; i < 5; i++ {
		fmt.Printf("%d Sekunden\n", i)
		time.Sleep(1 * time.Second);
	}

	// Connect to mongodb
	session, err := mgo.Dial("mongodb:27017")

	if err != nil {
		panic(err)
	}
	defer session.Close();

	names, err := session.DatabaseNames();

	fmt.Println();

	for _, name := range names {
		fmt.Println(name);
	}

	// Select database
	//collection := session.DB("resamvi").C("entries")

	if err != nil {
		log.Fatal(err);
	}

	http.HandleFunc("/", handler);
	http.ListenAndServe(":8080", nil);
}