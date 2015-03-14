package main

import (
	"errors"
	"net/http"
)

func sessionCheck(r *http.Request) (User, error) {
	var user User
	session, err := store.Get(r, "user")
	if err != nil {
		return user, err
	}

	if session.Values["email"] == nil {
		return user, errors.New("The User session is empty or hasn't been created")
	}

	user = User{
		Name:  session.Values["name"].(string),
		Email: session.Values["email"].(string),
		Image: session.Values["image"].(string),
	}

	return user, nil
}

//LoginData returns the LoginDataStruct with either carries an auth url or a
//user sttruct if authenticated
func LoginData(r *http.Request) LoginDataStruct {
	user, err := sessionCheck(r)

	if err != nil {
		data := LoginDataStruct{
			URL: FBURL,
		}
		return data
	}

	data := LoginDataStruct{
		User: user,
	}

	return data
}

//SearchPagination returns a page strict which carries details about the
//pagination of any given search result or page
func SearchPagination(count int, page int, perPage int) Page {
	var pg Page
	var total int

	if count%perPage != 0 {
		total = count/perPage + 1
	} else {
		total = count / perPage
	}

	if total < page {
		page = total
	}

	if page == 1 {
		pg.Prev = false
		pg.Next = true
	}

	if page != 1 {
		pg.Prev = true
	}

	if total > page {
		pg.Next = true
	}

	if total == page {
		pg.Next = false
	}

	var pgs = make([]string, total)

	//The number of number of documents to skip
	skip := perPage * (page - 1)

	pg.Total = total
	pg.Skip = skip
	pg.Count = count
	pg.NextVal = page + 1
	pg.PrevVal = page - 1
	pg.Pages = pgs

	return pg

}
