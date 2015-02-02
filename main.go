package main

import (
	"fmt"
	"os"

	"gopkg.in/mgo.v2"
)

var (
	//MONGOSERVER stores the address of the mongo db server address
	MONGOSERVER string

	//MONGODB stores the name of the database instance
	MONGODB string
)

//User would hold the user data for retrieving and sending items to the database
type User struct {
	Name string
	ID   string

	//Please add the remaining items. I just realised i dont have the schema
}

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

}

//NewUser is for adding a new user to the database. Please note that what you pass to the function is a pointer to the actual data, note the data its self. ie newUser(&NameofVariable)
func NewUser(data *User) error {

	//MONGOSERVER is a variable containing the mongo db instance address
	session, err := mgo.Dial(MONGOSERVER)
	checkPanic(err)
	defer session.Close()

	//MONGODB is the database name while MONGOC is the collection name
	collection := session.DB(MONGODB).C("users")

	err = collection.Insert(data)

	if err != nil {
		return err
	}
	return nil

}

//checkFmt checks the value of an error and prints it to standard output. I'm
//adding it to reduce the number of error checking ifs in my code
func checkFmt(err error) {
	if err != nil {
		fmt.Println(err.Error)
	}
}

//checkPanic checks the value of an error and panics if it isnt empty, thereby
//closing the program.
func checkPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

}