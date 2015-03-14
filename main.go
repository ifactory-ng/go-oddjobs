package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
)

var (
	//MONGOSERVER stores the address of the mongo db server address
	MONGOSERVER string

	//MONGODB stores the name of the database instance
	MONGODB string

	//PORT stores the port number
	PORT string

	store = sessions.NewCookieStore([]byte("bla-bla-bla-sheep-is-very-secret"))
)

//pre parse the template files, and store them in memory. Fail if
//they're not found
var templates = template.Must(template.ParseFiles("templates/index.html", "templates/search-results.html", "templates/profile.html"))

func init() {
	MONGOSERVER = os.Getenv("MONGOSERVER")
	if MONGOSERVER == "" {
		fmt.Println("No mongo server address set, resulting to default address")
		MONGOSERVER = "localhost"
	}

	MONGODB = os.Getenv("MONGODB")
	if MONGODB == "" {
		fmt.Println("No Mongo database name set, resulting to default")
		MONGODB = "oddjobs"
	}

	PORT = os.Getenv("PORT")
	if PORT == "" {
		fmt.Println("No Global port has been defined, using default")

		PORT = ":8080"

	}

	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600 * 1000, //3600 is 1 hour
		HttpOnly: false,
	}
}

//renderTemplate is simply a helper function that takes in the response writer
//interface, the template file name and the data to be passed in, as an
//interface. It causes an internal server error if any of the templates is not
//found. Better fail now than fail later, or display rubbish.
func renderTemplate(w http.ResponseWriter, tmpl string, q interface{}) {
	err := templates.ExecuteTemplate(w, tmpl, q)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func main() {
	//serve assets
	fs := http.FileServer(http.Dir("templates/assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	//serving public views
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/search", SearchHandler)
	http.HandleFunc("/profile", ProfileHandler)

	http.HandleFunc("/authenticate", LoginHandler)
	http.HandleFunc("/user/bookmark/", BookmarkHandler)
	log.Fatal(http.ListenAndServe(PORT, nil))

	fmt.Println("serving on http://localhost" + PORT)
	log.Fatal(http.ListenAndServe(PORT, context.ClearHandler(http.DefaultServeMux)))
}
