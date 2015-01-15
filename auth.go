package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/goquadro/core"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/martini-contrib/render"
)

type Session struct {
	Username string
	Uid      string
	Name     string
	Role     int
	Exp      time.Time
	Email    string
}

var InvalidJwt = errors.New("Invalid Token.")

// http://tools.ietf.org/html/draft-ietf-oauth-json-web-token-20

// Extracts uid from request
func JwtGetUser(req *http.Request) (*core.User, error) {
	u := new(core.User)
	gqtoken := req.Header.Get("gqtoken")

	if gqtoken == "" {
		return u, nil
	}

	token, err := jwt.Parse(gqtoken, func(token *jwt.Token) (interface{}, error) {
		//return myLookupKey(token.Header["kid"])
		return gqConfig.jwtVerifyKey, nil
	})
	if err != nil {
		return u, err
	}
	if !token.Valid {
		return u, InvalidJwt
	}
	uid := fmt.Sprint(token.Claims["uid"])

	err = u.GetById(uid)

	return u, err
}

// Provided a user object with username and password, validates auth and generates JWT
func getToken(user core.User) (string, error) {
	err := user.CheckPassword()
	if err != nil {
		return "", errors.New("401")
	}
	t := jwt.New(jwt.GetSigningMethod("RS256"))

	// set our claims
	t.Claims["uid"] = user.ID
	time.Now().Add(time.Hour * 72).Unix()
	t.Header["kid"] = "1"
	return t.SignedString(gqConfig.jwtSignKey)
}

func ApiUserLogin(r render.Render, w http.ResponseWriter, req *http.Request, user core.User) {
	user.Username = strings.ToLower(user.Username)
	tokenString, err := getToken(user)
	if err != nil {
		if err.Error() == "401" {
			r.JSON(http.StatusUnauthorized, nil)
			return
		}
		r.JSON(http.StatusInternalServerError, nil)
		return
	}
	token := map[string]string{"gqtoken": tokenString}
	r.JSON(http.StatusOK, token)
	return
}

func ApiUserSignup(w http.ResponseWriter, req *http.Request, r render.Render, user core.User) {
	user.Username = strings.ToLower(user.Username)

	err := user.SignupWithCode(req.FormValue("signupcode"))

	if err != nil {
		r.JSON(http.StatusUnauthorized, err)
		return
	}
	r.JSON(http.StatusAccepted, nil)
	return
}
