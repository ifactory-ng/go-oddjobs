package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//HomeHandler serves the home/search page to the user
func HomeHandler(w http.ResponseWriter, r *http.Request) {

}

//LoginHandler serves the home/search page to the user
func LoginHandler(w http.ResponseWriter, r *http.Request) {

	user := new(User)
	user.Email = r.FormValue("email")
	user.ID = r.FormValue("ID")
	user.Name = r.FormValue("name")
	user.Gender = r.FormValue("gender")
	user.Location = r.FormValue("location")
	stats := Authenticate(user)
	fmt.Println(w, stats)
}

func UserProfileHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	data, _ := json.Marshal(GetProfile(id))

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
func AddSkillHandler(w http.ResponseWriter, r *http.Request) {
	skill := new(Skill)
	skill.UserID = r.FormValue("id")
	skill.Location = r.FormValue("location")
	skill.Description = r.FormValue("desc")
	skill.Address = r.FormValue("address")
	skill.SkillName = r.FormValue("skill_name")
	skill.TagName = r.FormValue("tag_name")
	resp := AddSkill(skill)
	fmt.Println(w, resp)
}
func UserSkillshandler(w http.ResponseWriter, r *http.Request) {

	GetSkills(id)
}
