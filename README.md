# Basic Web 

Create Basic Web Page 

# Adding New Pages 

A web application would be boring with only one page, in this section I will work on adding new pages. 

In the first example I used if else statement to determine routing, this is definitely not the right way to do routing.


```go
w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome to My Simple Golang Site!</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "To get in touch please send an email "+
			"to <a href=\"mailto:support@devonartis.com\">"+"devon@devonartis.com</a>.")
	}
```
> [!TIP]
If you are interested in learning about dynamic reloading, check out the Go library
fresh - github.com/pilu/fresh

> 

## Routers 
In this section we focus on using routers as ifelse statements is not the route you want to go for routing.

There are multiple routers but in this section we will be using github.com/gorilla/mux.Router which is part of the Gorilla Toolkit


You will need to run the command 

`go get -u github.com/gorilla/mux`  

