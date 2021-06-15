package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var h http.Handler

var homeTemplate, contactTemplate, faqTemplate *template.Template

func faq(w http.ResponseWriter, http *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := faqTemplate.Execute(w, nil); err != nil {
		panic(err)
	}

}
func home(w http.ResponseWriter, http *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, http *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := contactTemplate.Execute(w, nil); err != nil {
		panic(err)
	}

}

func custom404(w http.ResponseWriter, http *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>404 Error</h1>"+"<p>The page your are looking for does not exist."+
		"</p><p><b>Please email support</b></p>")
}
func main() {

	h = http.HandlerFunc(custom404)
	var err error

	homeTemplate, err = template.ParseFiles("views/home.html")
	if err != nil {
		panic(err)
	}

	contactTemplate, err = template.ParseFiles("views/contact.html")
	if err != nil {
		panic(err)
	}

	faqTemplate, err = template.ParseFiles("views/faq.html")
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.NotFoundHandler = h

	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	http.ListenAndServe(":3000", r)
}
