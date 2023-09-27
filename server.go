package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fileServer := http.FileServer(http.Dir("./static"))
	http.HandleFunc("/form", formHandler)
	http.Handle("/", fileServer)
	// http.HandleFunc("/form", formHandler)
	http.HandleFunc("/ascii-art-web", Handler)

	fmt.Printf("Starting the server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art-web" {
		// if template / banner are not found
		http.Error(w, "404 Not Found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}

	// Bad Request - incorrect request
	// Internal server error - unhadled errors
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	banner := r.FormValue("banner")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Banner = %s\n", banner)
}
