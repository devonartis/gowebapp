# Web Development With Golang 

Create Basic Web Page 

## Adding New Pages 

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

## Excercise 

### 3.4.1 Ex1 - Add an FAQ page
This one is pretty straight forward. Try to create an FAQ page to your application
under the path /faq.
You can fill the page with whatever HTML content you prefer, but you should
make it different from the other pages you are certain your code is working as
you intended.

### 3.4.2 Ex2 - Custom 404 page
I mentioned earlier that gorilla/mux has a 404 page by default for paths we
don’t define with our router. You can actually customize this page by setting
the NotFoundHandler attribute on the gorilla/mux.Router.
If you are new to Go this exercise is likely going to prove to be challenging
because you will need to create an implementation of the http.Handler interface
and then assign that to the NotFoundHandler, and this is a little different from
what we have done so far.
To help with this, I have provided an example of how to convert the home
function that we wrote earlier in this chapter into the http.Handler type.


```go
var h http.Handler = http.HandlerFunc(home)
r := mux.NewRouter()
// This will assign the home page to the
// NotFoundHandler
r.NotFoundHandler = h
```

You will want to do something similar, but using your own unique 404 page
function.

### 3.4.3 Ex3 - [HARD] Try out another router

This exercises is labeled hard because it is a little open ended.
Check out another router, like github.com/julienschmidt/httprouter, and try to
replicate the program we have written so far using this router instead of gorilla/
mux.
This is a great way to both (a) ensure you understood what we were doing, and
(b) practice reading docs and using other libraries you are unfamiliar with.****

# VIEW

The work and labs done on MVC could be pulled from the view branch, intially I created a MVC
branch to be used for code but the chapter talked more about what is a MVC.

I have worked with Model View Controller years ago so I am quite familiar with MVC and MVC frameworks. I decided
to just create the view branch. 

## Delete the MVC branch and name this title to View

In the view branch views was added instead of hard coding the web response code.

# Reusable Bootstrap 

In the branch "bootstrap" a reusable bootstrap will be used to enhance the views 

# Branch bs_viewtype 

In this branch we create a the view type later on we will take the named templtes and apply it to create a layout for all pages, but first we are going to address the problem the problem we have with adding the footer template to every file.

This is fine for one or two files but imagine having 20 pages and adding the footer template in every page the view type is going to fix that problem.

**Listing 6.19: Adding Layout to NewView()
**

**Listing 6.20: Using the Layout in handlers.
**
Update home and contact function in main.go to use the new layout



**Listing 6.21: Adding Layout to NewView().** 

6.4 Navigation Bar - Completed

6.5 Cleaning up code ... pending 







