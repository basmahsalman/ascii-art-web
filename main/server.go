package main

import (
	"fmt"
	"text/template"

	//"html/template"
	ar "asciiartweb"
	"log"
	"net/http"
)

// html/teplate
var tmpl *template.Template
var anss *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("../static/index.html"))
	anss = template.Must(template.ParseFiles("../static/form.html"))

}

type textBanner struct {
	converted string
}

func main() {

	http.HandleFunc("/", Handler)
	http.HandleFunc("/ascii-art", formHandler)

	fmt.Printf("Starting the server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

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
	// ascii is the converted text we want to print
	ascii := textBanner{
		converted: r.FormValue("converted"),
	}
	tmpl.Execute(w, struct {
		Success bool
		ascii   textBanner
	}{true, ascii})
}

// handler for the server
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Method is not supported", http.StatusMethodNotAllowed)
		return
	}
	if r.FormValue("text") == "" {
		http.Error(w, "400 Bad Request.", http.StatusBadRequest)
		return
	}
	banner := r.Form.Get("banner")
	// if (banner != "standard") || (banner != "shadow") || (banner != "thinkertoy") {
	// 	http.Error(w, "500 Internal server error.", http.StatusInternalServerError)
	// 	return
	// }
	switch banner {
	case "standard":
	case "shadow":
	case "thinkertoy":
		break
	default:
		http.Error(w, "500 Internal server error.", http.StatusInternalServerError)
		return
	}

	text := r.FormValue("text")

	// Call Validating Inputs
	ar.ValidatingInput(text)

	// Calling Storing Function
	ans := ar.PrintArray(text, banner)

	anss.Execute(w, ans)
}
