package main

import (
	"fmt"
	"net/http"
)

//HomeHandler serves the home/search page to the user
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	type homestruct struct {
		User LoginDataStruct
	}

	data := homestruct{
		User: LoginData(r),
	}

	renderTemplate(w, "index.html", data)
}

//SearchHandler serves the search results page based on a search query from the
//index page or any search box
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "search-results.html", "")
}

//ProfileHandler might be remove later, its just to test redirection and profile
//data collection after login
func ProfileHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		renderTemplate(w, "profile.html", "")
	} else if r.Method == "POST" {
		fmt.Println("POST request logged")
	}
}
