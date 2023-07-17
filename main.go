package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", pathhandler)
	// http.HandleFunc("/contact", contactHandler)

	fmt.Println("Starting the server on port:3000")
	http.ListenAndServe(":3000", nil)
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	fmt.Fprintf(w, "<h1>This is heading 2</h1>")
}
func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	fmt.Fprintf(w, "<h2>This is contact page</h2>")
}
func pathhandler(w http.ResponseWriter, r *http.Request) {

	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		http.Error(w, "page not found", http.StatusNotFound)
		// w.WriteHeader(http.StatusNotFound)
		// fmt.Fprintf(w, "page not found")
	}
}
