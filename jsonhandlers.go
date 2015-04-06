package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

//LoginHandler serves the profile data to the user
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	fmt.Println(r.Form)
	id := bson.NewObjectId()
	//nw := strings.(id)
	user := &User{
		Email: r.FormValue("email"),
		ID:    r.FormValue("ID"),
		Name:  r.FormValue("name"),
		_id:   id,
	}

	fmt.Println(user)
	i, _ := Authenticate(user, r.FormValue("provider"))
	i2, _ := json.Marshal(i)
	w.Header().Set("Content-Type", "application/json")
	w.Write(i2)
}

//UserProfileHandler serves the profile
func UserProfileHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	//id := r.FormValue("id")
	tmp := strings.Split(r.URL.Path, "/")
	id := tmp[3]
	switch r.Method {
	case "GET":
		tmp2, _ := GetProfile(id)
		data, _ := json.Marshal(tmp2)

		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	case "POST":
		user := &User{
			Location: r.FormValue("location"),
			About:    r.FormValue("about"),
			Address:  r.FormValue("address"),
			Phone:    r.FormValue("phone"),
		}
		UpdateUser(user, id)
	}
}

//UserSkillsHandler to handle all skill related request such as add new skill and getting skill
func UserSkillshandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

	tmp := strings.Split(r.URL.Path, "/")
	id := tmp[3]
	switch r.Method {
	case "GET":
		//tmp := strings.Split(r.URL.Path, "/")
		//id := tmp[2]
		tmp2, _ := GetSkills(id)
		data, _ := json.Marshal(tmp2)
		fmt.Println(data)

		w.Header().Set("Content-Type", "application/json")
		w.Write(data)

	case "POST":
		userName, _ := GetProfile(id)
		tags := strings.Split(r.FormValue("tag"), ",")
		skill := &Skill{

			UserID:      id,
			UserName:    userName.Name,
			Location:    r.FormValue("location"),
			Description: r.FormValue("desc"),
			Address:     r.FormValue("address"),
			SkillName:   r.FormValue("skill_name"),
			Tags:        tags,
			Rating:      0,
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
			id:        r.FormValue("id"),
			Name:      r.FormValue("phone"),
			SkillName: r.FormValue("email"),
		}
		AddBookmark(bookmark, urlID)
	}
}

func SingleSkillHandler(w http.ResponseWriter, r *http.Request) {
	tmp := strings.Split(r.URL.Path, "/")
	urlID := tmp[3]
	fmt.Println(urlID)
	tmp2, _ := GetSkill(urlID)
	data, _ := json.Marshal(tmp2)
	fmt.Println(tmp2)

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)

}

func FeedsHandler(w http.ResponseWriter, r *http.Request) {
	v, _ := Popular()
	w.Header().Set("Content_Type", "application/json")
	data, _ := json.Marshal(v)
	fmt.Println(data)
	w.Write(data)
}

func ApiSearchHandler(w http.ResponseWriter, r *http.Request) {
	tmp := r.URL.Query().Get("location")
	fmt.Println(tmp)
	fmt.Println(r.URL)
	fmt.Println(r.URL.Query())
	tmp2 := r.URL.Query().Get("query")
	v, _, _ := Search(tmp, tmp2, 50, 50, 50)

	w.Header().Set("Content_Type", "application/json")
	data, _ := json.Marshal(v)
	fmt.Println(data)
	w.Write(data)
}
