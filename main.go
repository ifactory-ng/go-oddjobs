package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	//MONGOSERVER stores the address of the mongo db server address
	MONGOSERVER string

	//MONGODB stores the name of the database instance
	MONGODB string

	//PORT stores the port number
	PORT string
)

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

}

func main() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/authenticate", LoginHandler)
	log.Fatal(http.ListenAndServe(PORT, nil))

}
