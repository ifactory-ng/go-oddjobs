package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//NewUser is for adding a new user to the database. Please note that what you pass to the function is a pointer to the actual data, note the data its self. ie newUser(&NameofVariable)
func NewUser(data *User, socialProvider string) (bson.ObjectId, error) {

	//MONGOSERVER is a variable containing the mongo db instance address
	session, err := mgo.Dial(MONGOSERVER)
	checkPanic(err)
	defer session.Close()
	//user := User{}
	//userId := data._id
	lookUpSession := session.DB(MONGODB).C("lookup")

	//MONGODB is the database name while MONGOC is the collection name
	collection := session.DB(MONGODB).C("users")

	//i := bson.NewObjectId()

	err = collection.Insert(data)
	if err != nil {
		return "error", err
	}
	//collection.FindId(i).One(&user)
	//fmt.Println(i)
	//fmt.Println(user)
	//err = collection.Find(bson.M{"ID": data.ID}).One(&user)
	//checkFmt(err)

	lookup := &LookUp{
		Provider:       socialProvider,
		IdFromProvider: data.ID,
		UserId:         data._id,
	}
	fmt.Println(data)
	err = lookUpSession.Insert(lookup)
if err != nil {
		return "error", err
	}
	return data._id, nil
}

//Authenticate check if user exists if not create a new user document NewUser function is called within this function. note the user struct being passed
//to this function should alredi contain a self generated objectid
func Authenticate(user *User, provider string) (bson.ObjectId, error) {
	fmt.Println("test")
	session, err := mgo.Dial(MONGOSERVER)
	if err != nil {
		return " ", err
	}
	defer session.Close()
	result := LookUp{}
	lookupCollection := session.DB(MONGODB).C("lookup")

	err = lookupCollection.Find(bson.M{"IdFromProvider": user.ID, "provider": provider}).One(&result)
	//checkFmt(err)

	if result.UserId != "" {
		return result.UserId, nil
	}

	return NewUser(user, provider)
}

//UpdateUser updates a users details
func UpdateUser(data *User, id string) error {
	session, err := mgo.Dial(MONGOSERVER)

	if err != nil {
		return err
	}

	defer session.Close()

	collection := session.DB(MONGODB).C("users")
	query := bson.ObjectIdHex(id)
	change := bson.M{"$set": data}

	err = collection.Update(query, change)

	if err != nil {
		return err
	}

	return nil
}

//GetProfile returns the authenticated users profile. this does not include users skills
func GetProfile(id string) (User, error) {
	session, err := mgo.Dial(MONGOSERVER)

	result := User{}

	if err != nil {
		return result, err
	}
	defer session.Close()
	collection := session.DB(MONGODB).C("users")

	err = collection.FindId(bson.ObjectIdHex(id)).One(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

//AddSkill adds a skill to the collection
func AddSkill(data *Skill) error {
	session, err := mgo.Dial(MONGOSERVER)

	if err != nil {
		return err
	}

	defer session.Close()

	skillCollection := session.DB(MONGODB).C("skills")

	err = skillCollection.Insert(data)
	if err != nil {
		return err
	}
	return nil
}

//GetSkills gets all the skills added by user
func GetSkills(ID string) ([]Skill, error) {
	session, err := mgo.Dial(MONGOSERVER)

	result := []Skill{}
	if err != nil {
		return result, err
	}

	defer session.Close()

	skillCollection := session.DB(MONGODB).C("skills")

	err = skillCollection.Find(bson.M{"userid": ID}).Select(bson.M{"comments": 0}).All(&result)
	if err != nil {
		return result, err
	}
	return result, nil

}

//GetSkill return a single skill document
func GetSkill(id string) (Skill, error) {
	session, err := mgo.Dial(MONGOSERVER)

	result := Skill{}

	if err != nil {
		return result, err
	}

	defer session.Close()

	skillCollection := session.DB(MONGODB).C("skills")

	err = skillCollection.FindId(bson.ObjectIdHex(id)).Select(bson.M{"comments": 0}).One(&result)
	if err != nil {
		return result, err
	}

	return result, nil

}

//GetComment retrieves the reviews for a particular skill document
func GetComment(id string) ([]Skill, error) {
	session, err := mgo.Dial(MONGOSERVER)

	result := []Skill{}

	if err != nil {
		return result, err
	}

	defer session.Close()

	skillCollection := session.DB(MONGODB).C("skills")

	err = skillCollection.FindId(bson.ObjectIdHex(id)).Select(bson.M{"Comments": 1}).One(&result)
	if err != nil {
		return result, err
	}
	return result, nil

}

//AddBookmark is a utility function for adding bookmarks
func AddBookmark(bookmark *BookMark, id string) error {
	session, err := mgo.Dial(MONGOSERVER)
	if err != nil {
		return err
	}
	defer session.Close()
	userCollection := session.DB(MONGODB).C("users")
	query := bson.ObjectIdHex(id)
	change := bson.M{"$push": bson.M{"Bookmarks": bookmark}}
	err = userCollection.Update(query, change)
	if err != nil {
		return err
	}
	return nil
}

//GetBookmarks returns the users bookmarks
func GetBookmarks(id string) ([]User, error) {
	session, err := mgo.Dial(MONGOSERVER)
	result := []User{}

	if err != nil {
		return result, err
	}
	defer session.Close()
	userCollection := session.DB(MONGODB).C("users")
	err = userCollection.FindId(bson.ObjectIdHex(id)).Select(bson.M{"Bookmarks": 1}).All(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

//AddComment adds a comment to a skill
func AddComment(comment *Comment, id string) error {
	session, err := mgo.Dial(MONGOSERVER)

	if err != nil {
		return err
	}

	defer session.Close()

	skillCollection := session.DB(MONGODB).C("skills")

	err = skillCollection.UpdateId(
		bson.ObjectIdHex(id),
		bson.M{
			"$push": bson.M{"Comments": comment},
		})

	return nil

}

//Search takes a location and a search query and returns a slice of structs that
//match the query
func Search(location string, query string, count int, page int, perPage int) ([]Skill, Page, error) {
	var Results []Skill
	var Page Page
	session, err := mgo.Dial(MONGOSERVER)

	if err != nil {
		return Results, Page, err
	}

	defer session.Close()

	skillCollection := session.DB(MONGODB).C("skills")

	index := mgo.Index{
		Key: []string{"$text:skillname", "$text:description", "$text:location"},
	}

	err = skillCollection.EnsureIndex(index)
	if err != nil {
		return Results, Page, err
	}

	q := skillCollection.Find(
		bson.M{
			"location": location,
			"$text": bson.M{
				"$search": query,
			},
		},
	)

	//SearchPagination gives us a struct that tells us if the data has a
	//next page or previous page, as well as the page number
	Page = SearchPagination(count, page, perPage)

	err = q.Limit(perPage).Skip(Page.Skip).All(&Results)
	if err != nil {
		return Results, Page, err
	}
	return Results, Page, nil

}

//Popular does something i dont know
func Popular() ([]Skill, error) {
	session, err := mgo.Dial(MONGOSERVER)

	result := []Skill{}

	if err != nil {
		return result, err
	}
	defer session.Close()
	skillCollection := session.DB(MONGODB).C("skills")
	err = skillCollection.Find(bson.M{}).Limit(30).Sort("rating").All(&result)
	if err != nil {
		return result, err
	}
	return result, nil
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
