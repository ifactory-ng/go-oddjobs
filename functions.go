package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

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

//UpdateUser updates a users details
func UpdateUser(data *User) error {
	session, err := mgo.Dial(MONGOSERVER)

	if err != nil {
		return err
	}

	defer session.Close()

	collection := session.DB(MONGODB).C("users")

	err = collection.UpdateId(bson.ObjectIdHex(data.ID), data)

	if err != nil {
		return err
	}

	return nil
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

func GetSkills(id string) ([]Skill, error) {
	session, err := mgo.Dial(MONGOSERVER)

	result := []Skill{}

	if err != nil {
		return result, err
	}

	defer session.Close()

	skillCollection := session.DB(MONGODB).C("skills")

	err = skillCollection.Find(bson.M{"userID": id}).Select(bson.M{"comments": 0}).All(&result)
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

	err = skillCollection.Find(bson.M{"_id": id}).Select(bson.M{"comments": 0}).One(&result)
	if err != nil {
		return result, err
	}
	return result, nil

}

func Authenticate(user *User) error {
	session, err := mgo.Dial(MONGOSERVER)
	if err != nil {
		return err
	}
	defer session.Close()
	result := User{}
	userCollection := session.DB(MONGODB).C("users")

	err = userCollection.Find(bson.M{"ID": user.ID}).One(&result)
	if result.ID != "" {
		return err
	}
	NewUser(user)
	return err
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
