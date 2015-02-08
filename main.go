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

<<<<<<< HEAD
type Skill struct {
    Skillname string
    Userid int
    Location string
    Address string
    Price string
    Tagname string
    Description string
    
    }
=======
	//PORT stores the port number
	PORT string
)
>>>>>>> 8dfd2da576f21f4a77c50639b7ef4fe7f1e77089

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

<<<<<<< HEAD
func AddSkill(data *Skill) error{
    session, err := mgo.Dial(MONGOSERVER)
    checkPanic(err)
    defer session.Close()
    
    skillCollection := session.DB(MONGODB).C("skills")
    
    err = skillCollection.Insert(data)
    
    if err != nil {
        return err
    }
    return nil
}

func GetSkills(id string) Skill{
    session, err := mgo.Dial(MONGOSERVER)
    checkPanic(err)
    defer session.Close()
    
    skillCollection := session.DB(MONGODB).C("skills")
    
    result := Skill{}
    err = skillCollection.Find(bson.M{"User_id": id}).One(&result)
    checkFmt(err)
return result

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
=======
>>>>>>> 8dfd2da576f21f4a77c50639b7ef4fe7f1e77089
}

func main() {
	http.HandleFunc("/", HomeHandler)

	log.Fatal(http.ListenAndServe(PORT, nil))

}
