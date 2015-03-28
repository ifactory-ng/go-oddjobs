package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//HomeHandler serves the home/search page to the user
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	type homestruct struct {
		User  LoginDataStruct
		FBURL string
	}

	data := homestruct{
		User:  LoginData(r),
		FBURL: FBURL,
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

//ProfileEditHandler for now just logs the json value sent by the web client for
//debugging purposes
func ProfileEditHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	fmt.Println(r.Method)
	if r.Method == "GET" {
		fmt.Println("Get request")
		user := User{
			Name:  "Anthony Alaribe Test",
			About: "bla bla bla bla",
			Email: "me@me.com",
		}
		x, err := json.Marshal(user)
		if err != nil {
			fmt.Println(err)
		}
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(x)

		if err != nil {
			fmt.Println(err)
		}

	} else if r.Method == "POST" {
		hah, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(w, "%s", err)
		}
		fmt.Println(string(hah))
		user := User{}

		err = json.Unmarshal(hah, &user)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(user)
	}
}

//SkillsHandler would return list of skills via json, and suport editing and
//addition of new skills
func SkillsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	fmt.Println(r.Method)
	hah, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
	}
	fmt.Println(string(hah))
	if r.Method == "GET" {
		fmt.Println("get request")

		skills := []Skill{}

		skill1 := Skill{
			SkillName:   "Electrician",
			Tags:        []string{"tech", "farm"},
			Location:    "Calabar",
			Address:     "QC 28 unical staff quaters",
			Description: "dasfklsdgf sdflksd fdsf sd",
		}
		skill2 := Skill{
			SkillName:   "Programmer",
			Tags:        []string{"code"},
			Location:    "Aba",
			Address:     "QC 20 aba town",
			Description: "sdfdsf sdf sd f dsf ds gf sd  sdfds",
		}

		skills = append(skills, skill1)
		skills = append(skills, skill2)

		x, err := json.Marshal(skills)
		fmt.Print(string(x))
		if err != nil {
			fmt.Println(err)
		}
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(x)

		if err != nil {
			fmt.Println(err)
		}

	}

}
