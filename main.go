package main

import (
	"fmt"
	"log"
	"net/http"
)

// net/http to create go server and routes
// * is pointer

func helloHandler(w http.ResponseWriter, r *http.Request) {

	// if suppose by mistake a diifferent url redirects here instead of /hello, which will never happen, bt if it does
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not Supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "HELLO!!!") // Fprintf formats according to a format specifier and writes to w.
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	// submit form and we want it to be parsed to POST request
	// we will take these form values from req and store it in variables and display on response

	// http.Error for sending structured error responses with specific status codes and
	// fmt.Fprintf for generating custom response messages

	if err := r.ParseForm(); err != nil { // r.ParseForm() is called to parse the form data from the HTTP request.
		fmt.Fprintf(w, "ParseForm() err : %v", err)
		return
	}

	fmt.Fprintf(w, "POST request successful")

	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func main() {
	// http.FileServer creates a handler that serves HTTP requests with the contents of the given directory.
	// This fileServer variable now holds an http.Handler that can handle HTTP requests for files in the ./static directory.
	// Now it will automatically look at index.html, as all languages have this default behaviour to go thru index.html as servers

	fileServer := http.FileServer(http.Dir("./static"))
	// creating routes
	http.Handle("/", fileServer)            // root path i.e. index.html, fileServer is handler func for root path
	http.HandleFunc("/hello", helloHandler) // helloHandler is the handler func for /hello path
	http.HandleFunc("/form", formHandler)

	fmt.Printf("Starting server at Port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
