package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/facebook"
)

func init() {
	goth.UseProviders(
		facebook.New("", "", "http://localhost:8080/auth/facebook/callback"),
	)
}

//OAuthHandler would be in charge of authentication via oauth
func OAuthHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path

	variables := strings.Split(url, "/")
	//provider := variables[1]
	fmt.Println(len(variables))
	if len(variables) > 2 {
		fmt.Println(r.URL.Query().Get("state"))

		user, err := gothic.CompleteUserAuth(w, r)
		if err != nil {
			fmt.Println(err)

		}
		fmt.Println(user)
	} else {
		url, err := gothic.GetAuthURL(w, r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, err)
			return
		}

		http.Redirect(w, r, url, http.StatusTemporaryRedirect)

	}

}
