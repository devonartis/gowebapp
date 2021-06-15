package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

var homeTemplate *template.Template

func faq(w http.ResponseWriter, http *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>FAQ Placeholder!</h1>")
}
func home(w http.ResponseWriter, http *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, http *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "To get in touch please send an email "+
		"to <a href=\"mailto:support@devonartis.com\">"+"support@devonartis.com</a>.")
}

func main() {
	var err error

	homeTemplate, err = template.ParseFiles("views/home.html")

	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	http.ListenAndServe(":3000", r)
}
