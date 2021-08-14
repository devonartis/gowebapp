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

func faq(w http.ResponseWriter, http *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := faqView.Template.ExecuteTemplate(w, faqView.Layout, nil)
	if err != nil {
		panic(err)
	}

}
func home(w http.ResponseWriter, http *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := homeView.Template.ExecuteTemplate(w, homeView.Layout, nil)
	if err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, http *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := contactView.Template.ExecuteTemplate(w, contactView.Layout, nil)
	if err != nil {
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
	/*
		var err error


		homeTemplate, err = template.ParseFiles(
			"views/home.html",
			"views/layout/footer.html")
		if err != nil {
			panic(err)
		}


		contactTemplate, err = template.ParseFiles(
			"views/contact.html",
			"views/layout/footer.html")

		if err != nil {
			panic(err)
		}

		faqTemplate, err = template.ParseFiles(
			"views/faq.html",
			"views/layout/footer.html")
		if err != nil {
			panic(err)
		}
	*/

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
