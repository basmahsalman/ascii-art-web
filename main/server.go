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
	tmpl = template.Must(template.ParseFiles("../static/form.html"))
}

type textBanner struct {
	//converted [][]string
	converted string
}

func main() {

	// fileServer := http.FileServer(http.Dir("./static"))
	//1
	http.HandleFunc("/", formHandler)
	// http.Handle("/", fileServer)
	// http.HandleFunc("/form", formHandler)
	http.HandleFunc("/ascii-art-web", Handler)

	fmt.Printf("Starting the server at port 8080\n")
	//1
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// handler for the gathering the data
// func Handler(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/ascii-art-web" {
// 		// if template / banner are not found
// 		http.Error(w, "404 Not Found.", http.StatusNotFound)
// 		return
// 	}

// 	if r.Method != "GET" {
// 		http.Error(w, "Method is not supported", http.StatusNotFound)
// 		return
// 	}

// 	// Bad Request 400 - incorrect request
// 	// blank

// 	// Internal server error 500 - unhadled errors
// 	// arabic
// }

// again handler for gathering data
func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}
	// ascii is the converted text we want to print
	ascii := textBanner{
		// hmmm, mismatch type string and [][]string
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
