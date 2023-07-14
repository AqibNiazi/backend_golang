package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("Starting the server on port:3000")
	http.ListenAndServe(":3000", nil)

}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello and welcome to course</h1>")
}
