package main

import (
	"gopkg.in/mgo.v2/bson"
)

//User would hold the user data for retrieving and sending items to the database
type User struct {
	_id       bson.ObjectId
	Name      string
	ID        string
	About     string
	Email     string
	Location  string
	Address   string
	Bookmarks []BookMark
	Phone     string
	Gender    string
	Image     string
}

//Skill struct holds skill data to be used for adding and retrieving user skills
//from the database
type Skill struct {
	SkillName   string
	UserName    string
	TagName     []string
	UserID      string
	Location    string
	Address     string
	Price       string
	Description string
	Comments    []Comment
	Rating      int
}

//Comment holds comment data
type Comment struct {
	Name    string
	Email   string
	Comment string
	Rating  int
}

//LoginDataStruct carries information about a user if logged in, or an
//authentication url if not logged in
type LoginDataStruct struct {
	URL  string
	User User
}
type LookUp struct {
	Provider       string
	IdFromProvider string
	UserId         bson.ObjectId
}

//Page carries pagination info to aid in knowing whether any given page has a
//next or previous page, and to know its page number
type Page struct {
	Prev    bool
	PrevVal int

	Next    bool
	NextVal int

	pages int
	Pages []string
	Total int
	Count int
	Skip  int
}

//BookMark holds bookmark data
type BookMark struct {
	Name      string
	SkillName string
	id        string
}
