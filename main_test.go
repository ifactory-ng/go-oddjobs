package main

import "testing"

//TestNewUser tests whether the user item actually gets sent to the database
func TestNewUser(t *testing.T) {
	user := User{
		ID:   "12345678890",
		Name: "Anthony Alaribe",
	}
	err := NewUser(user)
	if err != nil {
		t.Error("Unable to add user to database")
	}
    
    func TestNewSkill(t *testing.T) {
	skill := Skill{
		User_id:   "12345678890",
		Skill_name: "Programmer",
        Description: "ajadjkadjkad",
        Price: "10000",
        Location: "fall",
        Address: "Cali",
        Tag_name: "kslfksl",
	}
	err := AddSkill(skill)
	if err != nil {
		t.Error("Unable to add skill to database")
	}
        
}    
    func TestGetSkill(t *testing.T) {
            err := GetSkills("1234567890")
        if err != nil {
		t.Error("Unable to add skill to database")
	}
        
        
}
