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
	query := bson.M{"ID": data.ID}
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

	err = collection.Find(bson.M{"ID": id}).One(&result)
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

//GetComment retrieves the reviews for a particular skill document
func GetComment(id string) (Skill, error) {
	session, err := mgo.Dial(MONGOSERVER)

	result := Skill{}

	if err != nil {
		return result, err
	}

	defer session.Close()

	skillCollection := session.DB(MONGODB).C("skills")

	err = skillCollection.Find(bson.M{"ID": id}).Select(bson.M{"Comments": 1}).One(&result)
	if err != nil {
		return result, err
	}
	return result, nil

}

//Authenticate check if user exists if not create a new user document NewUser function is called within this function
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

	return NewUser(user)
}

func Bookmark(bookmark *Bookmark, id string) error {
	session, err := mgo.Dial(MONGOSERVER)
	if err != nil {
		return err
	}
	defer session.close()
	userCollection := session.DB(MONGODB).C("users")
	query := bson.M{"ID": id}
	change := bson.M{"$push": bson.M{"Bookmarks": bookmark}}
	err = userCollection.Update(query, change)
	if err != nil {
		return err
	}
	return nil
}

//GetBookmarks returns the users bookmarks
func GetBookmarks(id string) (User, error) {
	session, err := mgo.Dial(MONGOSERVER)
	result = User{}

	if err != nil {
		return result, err
	}
	defer session.close()
	userCollection := session.DB(MONGODB).C("users")
	err = userCollection.Find(bson.M{"ID": id}).Select(bson.M{"Bookmarks": 1}).All(&result)
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

func Popular() (Skill, error) {
	session, err := mgo.Dial(MONGOSERVER)

	result := Skill{}

	if err != nil {
		return result, err
	}
	skillCollection := session.DB(MONGODB).C("skills")
	err = skillCollection.Find().Limit(15).Sort("rating").All(&result)
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
