package main

import (
	"fmt"
	//"html/template"
	"text/template"
	//"html/template"
	ar "asciiartweb"
	"log"
	"net/http"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("../static/index.html"))

}

type textBanner struct {
	//converted [][]string
	converted string
}

func main() {

	// fileServer := http.FileServer(http.Dir("./static"))
	//1
	http.HandleFunc("/", Handler)
	// http.Handle("/", fileServer)
	// http.HandleFunc("/form", formHandler)
	http.HandleFunc("/ascii-art", formHandler)

	fmt.Printf("Starting the server at port 8080\n")
	//1
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// func Handler(w http.ResponseWriter, r *http.Request) {
// if r.URL.Path != "/ascii-art-web" {
// 	// if template / banner are not found
// 	http.Error(w, "404 Not Found.", http.StatusNotFound)
// 	return
// }

// if r.Method != "GET" {
// 	http.Error(w, "Method is not supported", http.StatusNotFound)
// 	return
// }

// 	// Bad Request 400 - incorrect request
// 	// blank

// 	// Internal server error 500 - unhadled errors
// 	// arabic
// }

// again handler for gathering data
func Handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 Page not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodPost {
		err := tmpl.Execute(w, nil)
		if err != nil {
			log.Println(err)
		}
		return
	}
	// if r.URL.Path != "/ascii-art-web" {
	// 	// if template / banner are not found
	// 	http.Error(w, "404 Not Found.", http.StatusNotFound)
	// 	return
	// }

	// ascii is the converted text we want to print
	ascii := textBanner{
		converted: r.FormValue("converted"),
	}
	tmpl.Execute(w, struct {
		Success bool
		ascii   textBanner
	}{true, ascii})
}

// after this function is written we can run the server

// 2
// handler for the server
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}
	if r.FormValue("text") == "" {
		http.Error(w, "400 Bad Request.", http.StatusNotFound)
		return
	}
	if r.FormValue("banner") != "standard" || r.FormValue("banner") != "shadow" || r.FormValue("banner") != "thinkertoy" {
		http.Error(w, "500 Internal server error.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "POST request successful")

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	fmt.Fprintf(w, "Text = %s\n", text)
	fmt.Fprintf(w, "Banner = %s\n", banner)

	// Call Validating Inputs
	ar.ValidatingInput(text)

	// Calling Storing Function
	ar.PrintArray(text, banner)
}
