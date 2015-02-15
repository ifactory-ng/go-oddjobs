package main

import "testing"

//TestNewUser tests whether the user item actually gets sent to the database
func TestNewUser(t *testing.T) {
	user := User{
		ID:   "12345678890",
		Name: "Anthony Alaribe",
	}
	err := NewUser(&user)
	if err != nil {
		t.Error("Unable to add user to database in test")
	}
}
<<<<<<< HEAD
    func TestNewSkill(t *testing.T) {
=======

func TestUpdateUser(t *testing.T) {
	user := User{
		ID:   "12345678890",
		Name: "Anthony Alaribe",
	}
	err := UpdateUser(&user)
	if err != nil {
		t.Error("Unable to update user data in test")
	}
}

func TestNewSkill(t *testing.T) {
>>>>>>> 8dfd2da576f21f4a77c50639b7ef4fe7f1e77089
	skill := Skill{
		UserID:      "12345678890",
		SkillName:   "Programmer",
		Description: "ajadjkadjkad",
		Price:       "10000",
		Location:    "fall",
		Address:     "Cali",
		TagName:     "kslfksl",
	}
	err := AddSkill(&skill)
	if err != nil {
		t.Error("Unable to add skill to database")
	}

}
func TestGetSkill(t *testing.T) {

	skill := GetSkills("1234567890")
	if skill != nil {
		t.Error("Unable to retrieve skill from database")
	}

}
