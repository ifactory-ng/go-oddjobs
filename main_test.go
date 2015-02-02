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
}
