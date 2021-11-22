package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type authInfo struct {
	Usr string
	Pwd string
}

func (api *API) authJWT(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var auth authInfo
	err = json.Unmarshal(body, &auth)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, user := range users {
		if auth.Usr == user.username && auth.Pwd == user.password {
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"usr": auth.Usr,
				"nbf": time.Now().Unix(),
			})

			// Sign and get the complete encoded token as a string using the secret
			tokenString, err := token.SignedString([]byte("secret-password"))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write([]byte(tokenString))
			return
		}
	}
}

type user struct {
	username    string
	password    string
	permissions int
}

var users = []user{
	{
		username:    "user1",
		password:    "pwd1",
		permissions: 10,
	},
	{
		username:    "user2",
		password:    "pwd2",
		permissions: 3,
	},
}
