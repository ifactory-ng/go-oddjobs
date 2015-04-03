package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

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

	oauthstring = "hellother"
)

//AccessToken is where the facebook authentication data would be stored
type AccessToken struct {
	Token  string
	Expiry int64
}

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

	FBURL = fbConfig.AuthCodeURL(oauthstring)
}

func readHTTPBody(response *http.Response) string {

	fmt.Println("Reading body")

	bodyBuffer := make([]byte, 5000)
	var str string

	count, err := response.Body.Read(bodyBuffer)

	for ; count > 0; count, err = response.Body.Read(bodyBuffer) {

		if err != nil {

		}

		str += string(bodyBuffer[:count])
	}

	return str

}

// GetAccessToken Converts a code to an Auth_Token
func GetAccessToken(clientID string, code string, secret string, callbackURI string) AccessToken {
	fmt.Println("GetAccessToken")
	//https://graph.facebook.com/oauth/access_token?client_id=YOUR_APP_ID&redirect_uri=YOUR_REDIRECT_URI&client_secret=YOUR_APP_SECRET&code=CODE_GENERATED_BY_FACEBOOK
	response, err := http.Get("https://graph.facebook.com/oauth/access_token?client_id=" +
		clientID + "&redirect_uri=" + callbackURI +
		"&client_secret=" + secret + "&code=" + code)

	if err == nil {

		auth := readHTTPBody(response)
		//a, err := ioutil.ReadAll(response.Body)
		fmt.Println(auth)
		var token AccessToken

		tokenArr := strings.Split(auth, "&")
		fmt.Println(tokenArr)

		token.Token = strings.Split(tokenArr[0], "=")[1]
		expireInt, err := strconv.Atoi(strings.Split(tokenArr[1], "=")[1])

		if err == nil {
			token.Expiry = int64(expireInt)
		}

		return token
	}

	var token AccessToken

	return token
}

//FacebookOAUTH is the handler that would be redirected to
func FacebookOAUTH(w http.ResponseWriter, r *http.Request) {

	state := r.FormValue("state")
	if state != oauthstring {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthstring, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	// grab the code fragment

	code := r.FormValue("code")

	//RedirectURL := RootURL + "/fblogin"
	fmt.Println(code)
	accessToken := GetAccessToken(FBClientID, code, FBClientSecret, fbConfig.RedirectURL)
	//accessToken, err := fbConfig.Exchange(oauth2.NoContext, code)
	/*if err != nil {
		fmt.Printf("oauthConf.Exchange() failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	fmt.Println("Expect access token next")*/
	fmt.Println(accessToken)
	//client := fbConfig.Client(oauth2.NoContext, accessToken)

	resp, err := http.Get("https://graph.facebook.com/me?access_token=" + accessToken.AccessToken)

	//resp, err := client.Get("https://graph.facebook.com/me?access_token="+accessToken.Acc)

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
