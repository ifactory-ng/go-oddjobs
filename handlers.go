package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

//UserProfileHandler serves the profile
func UserProfileHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	data, _ := json.Marshal(GetProfile(id))

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

//UserSkillsHandler to handle all skill related request such as add new skill and getting skill
func UserSkillshandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		tmp := r.URL.Query("id")
		id := tmp[2]
		data, _ := json.Marshal(GetSkills(id))

		w.Header().Set("Content-Type", "application/json")
		w.Write(data)

	case "POST":
		skill := Skill{

			UserID:      r.FormValue("id"),
			Location:    r.FormValue("location"),
			Description: r.FormValue("desc"),
			Address:     r.FormValue("address"),
			SkillName:   r.FormValue("skill_name"),
			TagName:     r.FormValue("tag_name"),
		}
		resp := AddSkill(skill)
		fmt.Println(w, resp)

	}

}

func BookmarkHandler(w http.ResponseWriter, r *http.Request) {
	tmp := strings.Split(r.URL.Path, "/")
	urlID := tmp[2]
	switch r.Method {
	case "GET":
		data, _ := GetBookmarks(urlID)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	case "POST":
		bookmark := Bookmark{
			Name:  r.FormValue("name"),
			Phone: r.FormValue("phone"),
			Email: r.FormValue("email"),
		}
		Bookmark(bookmark, urlID)
	}
}

func SingleSkillHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	data, _ := json.Marshal(GetSkill(id))

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)

}
