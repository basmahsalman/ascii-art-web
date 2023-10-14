package main

import (
	"fmt"
	"text/template"	
	ar "asciiartweb"
	"log"
	"net/http" //proiveds all the functionality for creating the web server
)

/* text/template  *1 */
var tmpl *template.Template
var anss *template.Template
var text string

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
	//ListAndServe method allows us to start the web server and specify the port to listen for incoming requests
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// HandleFunc function to add route handlers to the web server
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
	// to read the data from the form
	ascii := textBanner{
		converted: r.FormValue("converted"),
	}
	tmpl.Execute(w, struct {
		// to confirm
		Success bool
		ascii   textBanner
	}{true, ascii})
}

// formHandler function, holds all the logic related to the /ascii-art request.
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "405 Method is not supported", http.StatusMethodNotAllowed)
		return
	}
	text = r.FormValue("text")
	text = ar.ValidatingInput(text)
	if text == "" {
		http.Error(w, "400 Bad Request.", http.StatusBadRequest)
		return
	}
	banner := r.Form.Get("banner")
	switch banner {
	case "standard":
	case "shadow":
	case "thinkertoy":
		break
	default:
		http.Error(w, "404 banner not found", http.StatusNotFound)
		return
	}
	
	ans := ar.PrintArray(text, banner)

	anss.Execute(w, ans)
}