package main

import (
	"fmt"
	_ "html/template"
	"net/http"

	"github.com/devonartis/gowebapp/views"
	"github.com/gorilla/mux"
)

var h http.Handler

//var homeTemplate, contactTemplate, faqTemplate *template.Template

var homeView, contactView, faqView *views.View

// A helper function that will panic if an error occurs.
func must(err error) {
	if err !=nil {
		panic(err)
	}
}

func faq(w http.ResponseWriter, http *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(faqView.Render(w, nil))
}

func home(w http.ResponseWriter, http *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, nil))
	
}

func contact(w http.ResponseWriter, http *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(contactView.Render(w,nil))
}

func custom404(w http.ResponseWriter, http *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>404 Error</h1>"+"<p>The page your are looking for does not exist."+
		"</p><p><b>Please email support</b></p>")
}
func main() {

	h = http.HandlerFunc(custom404)
	
	homeView = views.NewView("bootstrap","views/home.html")
	contactView = views.NewView("bootstrap","views/contact.html")
	faqView = views.NewView("bootstrap","views/faq.html")

	r := mux.NewRouter()
	r.NotFoundHandler = h

	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	http.ListenAndServe(":3001", r)
}
