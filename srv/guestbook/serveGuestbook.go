package main

import (
	"text/template"
	"os"
	"fmt"
)

func main() {

	var i interface{}

	t := template.Must(template.ParseFiles("gaestebuch.html"))

	err := t.Execute(os.Stdout, i)

	if err != nil {
		fmt.Println(err)
	}

}