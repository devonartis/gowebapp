package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var h http.Handler

func faq(w http.ResponseWriter, http *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>FAQ Placeholder!</h1>")
}
func home(w http.ResponseWriter, http *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to My Simple Golang Site!</h1>")
}

func contact(w http.ResponseWriter, http *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "To get in touch please send an email "+
		"to <a href=\"mailto:support@devonartis.com\">"+"devon@devonartis.com</a>.")
}

func custom404(w http.ResponseWriter, http *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>404 Error</h1>"+"<p>The page your are looking for does not exist."+
		"</p><p><b>Please email support</b></p>")
}

func main() {

	router := httprouter.New()
	router.GET("/", home)
	router.GET("/contact", contact)
	router.GET("/faq", faq)
	log.Fatal(http.ListenAndServe(":3000", router))
}
