package main

import (
	"html/template"
	"os"
	
)

func main() {
	t, err := template.ParseFiles("hello.html")
	if err != nil {
		panic(err)
	}

	data := struct {
		Name string 
	}{"John Smith"}
	
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

}