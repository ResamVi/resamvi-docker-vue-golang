package main

import "fmt"
import "net/http"


func handler(write http.ResponseWriter, reader *http.Request) {
	write.Header().Set("Access-Control-Allow-Origin", "http://localhost:9001");
	fmt.Println("Send received");
}

func main() {
	http.HandleFunc("/", handler);
	http.ListenAndServe(":8080", nil);
}