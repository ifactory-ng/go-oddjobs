package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/antonholmquist/jason"
	"golang.org/x/oauth2"
)

var (
	// FBClientID facebook clientid
	FBClientID string
	// FBClientSecret ish
	FBClientSecret string
	//RootURL ish
	RootURL string

	//FBURL is the link to facebook redirect page
	FBURL string

	fbConfig oauth2.Config
)

func init() {

	FBClientID = os.Getenv("FBClientID")
	FBClientSecret = os.Getenv("FBClientSecret")
	RootURL = os.Getenv("RootURL")

	fbConfig := &oauth2.Config{
		// ClientId: FBAppID(string), ClientSecret : FBSecret(string)
		// Example - ClientId: "1234567890", ClientSecret: "red2drdff6e2321e51aedcc94e19c76ee"

		ClientID:     FBClientID, // change this to yours
		ClientSecret: FBClientSecret,
		RedirectURL:  RootURL + "/fblogin", // change this to your webserver adddress
		Scopes:       []string{"email"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://www.facebook.com/dialog/oauth",
			TokenURL: "https://graph.facebook.com/oauth/access_token",
		},
	}
	FBURL = fbConfig.AuthCodeURL("hellothereasimasdfkjhaskjdf")
}

//FacebookOAUTH is the handler that would be redirected to
func FacebookOAUTH(w http.ResponseWriter, r *http.Request) {
	// grab the code fragment

	code := r.FormValue("code")

	//RedirectURL := RootURL + "/fblogin"
	fmt.Println(code)

	accessToken, err := fbConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Error")
	}
	fmt.Println("Expect access token next")
	//fmt.Println(accessToken.AccessToken)

	//response, err := http.Get("https://graph.facebook.com/me?access_token=" + accessToken.AccessToken)

	client := fbConfig.Client(oauth2.NoContext, token)
	resp, err := client.Get("https://graph.facebook.com/me")

	// handle err. You need to change this into something more robust
	// such as redirect back to home page with error message
	if err != nil {
		fmt.Println(err.Error())
	}

	str, err := ioutil.ReadAll(resp.Body)
	fmt.Println(str)
	user, err := jason.NewObjectFromBytes([]byte(str))
	if err != nil {
		fmt.Println(err)
	}

	id, err := user.GetString("id")
	if err != nil {
		fmt.Println(err)
	}

	email, err := user.GetString("email")
	if err != nil {
		fmt.Println(err)
	}

	name, err := user.GetString("name")
	if err != nil {
		fmt.Println(err)
	}

	img := "https://graph.facebook.com/" + id + "/picture?width=180&height=180"
	fmt.Println("It got this far; the ID is")
	fmt.Println(id)
	session, err := store.Get(r, "user")
	if err != nil {
		fmt.Println(err)
	}

	session.Values["email"] = email
	session.Values["name"] = name
	session.Values["image"] = img
	session.Values["FBID"] = id
	err = session.Save(r, w)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Checking the session values")
	fmt.Println(session.Values["email"])

	http.Redirect(w, r, "/", http.StatusFound)
}
