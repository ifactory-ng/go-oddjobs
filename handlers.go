package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

//HomeHandler serves the home/search page to the user
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	type homestruct struct {
		User LoginDataStruct
	}

	data := homestruct{
		User: LoginData(r),
	}

	renderTemplate(w, "index.html", data)
}

//SearchHandler serves the search results page based on a search query from the
//index page or any search box
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "search-results.html", "")
}

//ProfileHandler might be remove later, its just to test redirection and profile
//data collection after login
func ProfileHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		renderTemplate(w, "profile.html", "")
	} else if r.Method == "POST" {
		fmt.Println("POST request logged")
	}
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
	tmp, _ := GetProfile(id)
	data, _ := json.Marshal(tmp)

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

//UserSkillsHandler to handle all skill related request such as add new skill and getting skill
func UserSkillsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		tmp := strings.Split(r.URL.Path, "/")
		id := tmp[2]
		tmp2, _ := GetSkills(id)
		data, _ := json.Marshal(tmp2)

		w.Header().Set("Content-Type", "application/json")
		w.Write(data)

	case "POST":
		skill := &Skill{

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

//BookmarkHandler handles bookmarking
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

//SingleSkillHandler handlers single skills for the api
func SingleSkillHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	tmp, _ := GetSkill(id)
	data, _ := json.Marshal(tmp)

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)

}
