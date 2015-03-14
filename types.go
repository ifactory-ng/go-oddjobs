package main

//User would hold the user data for retrieving and sending items to the database
type User struct {
	Name      string
	id        string
	About     string
	Email     string
	Location  string
	Address   string
	Bookmarks []BookMark
	Phone     string
	Gender    string
}

//Skill struct holds skill data to be used for adding and retrieving user skills
//from the database
type Skill struct {
	SkillName   string
	TagName     string
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

//Bookmark holds bookmark data
type BookMark struct {
	Name  string
	Phone string
	Email string
}
