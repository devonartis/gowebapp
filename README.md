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

Add tag



