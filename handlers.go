package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

//HomeHandler serves the home/search page to the user
func HomeHandler(w http.ResponseWriter, r *http.Request) {

}

//LoginHandler serves the profile data to the user
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Form)
	user := &User{
		Email:    r.FormValue("email"),
		id:       r.FormValue("ID"),
		Name:     r.FormValue("name"),
		Gender:   r.FormValue("gender"),
		Location: r.FormValue("location"),
	}

	fmt.Println(user)
	Authenticate(user)
}

//UserProfileHandler serves the profile
func UserProfileHandler(w http.ResponseWriter, r *http.Request) {
	//id := r.FormValue("id")
	tmp := strings.Split(r.URL.Path, "/")
	id := tmp[2]
	tmp2, _ := GetProfile(id)
	data, _ := json.Marshal(tmp2)

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

//UserSkillsHandler to handle all skill related request such as add new skill and getting skill
func UserSkillshandler(w http.ResponseWriter, r *http.Request) {
	tmp := strings.Split(r.URL.Path, "/")
	id := tmp[2]
	switch r.Method {
	case "GET":
		//tmp := strings.Split(r.URL.Path, "/")
		//id := tmp[2]
		tmp2, _ := GetSkills(id)
		data, _ := json.Marshal(tmp2)

		w.Header().Set("Content-Type", "application/json")
		w.Write(data)

	case "POST":
		skill := &Skill{

			UserID:      id,
			Location:    r.FormValue("location"),
			Description: r.FormValue("desc"),
			Address:     r.FormValue("address"),
			SkillName:   r.FormValue("skill_name"),
			TagName:     r.FormValue("tag"),
		}
		fmt.Println(r.Form)
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
		bookmarkData, _ := json.Marshal(data)
		w.Header().Set("Content-Type", "application/json")
		w.Write(bookmarkData)
	case "POST":
		bookmark := &BookMark{
			Name:  r.FormValue("name"),
			Phone: r.FormValue("phone"),
			Email: r.FormValue("email"),
		}
		AddBookmark(bookmark, urlID)
	}
}

func SingleSkillHandler(w http.ResponseWriter, r *http.Request) {
	tmp := strings.Split(r.URL.Path, "/")
	urlID := tmp[2]
	tmp2, _ := GetSkill(urlID)
	data, _ := json.Marshal(tmp2)

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)

}
