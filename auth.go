package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"code.google.com/p/goauth2/oauth"
)

type Auth struct{}

var config = &oauth.Config{
	ClientId:     "400554146779733",
	ClientSecret: "f3860a201ccbdfb0a95bb32ba1cfe131",
	//RedirectURL:   // It is setted dynamicaly on LoginHandler based on server current url
	Scope:    "public_profile user_friends  email",
	AuthURL:  "https://www.facebook.com/dialog/oauth",
	TokenURL: "https://graph.facebook.com/oauth/access_token",
}

var transport = &oauth.Transport{
	Config:    config,
	Transport: http.DefaultTransport,
}

// Will come to this method
// when user ask for [GET] /auth/login
func (l *Auth) GETLogin(w http.ResponseWriter, req *http.Request, env Env) {
	state := req.URL.Query().Get("method")
	if state == "" { // Here we can send something to the GETCallback... but...
		state = "nothing to transfer..."
	}

	// Set based on the server current url
	transport.Config.RedirectURL = env.Url + "api/auth/callback"

	log.Println("*** R:", transport.Config.RedirectURL)

	http.Redirect(w, req, transport.Config.AuthCodeURL(state), 302)
}

// Facebok server will call this method to authenticate user
func (l *Auth) GETCallback(db *DB, w http.ResponseWriter, req *http.Request) (*interface{}, error) {
	// Nothing yet received by GETLogin method
	//state := req.URL.Query().Get("state")

	code := req.URL.Query().Get("code")

	// Token = oauth.Token, err = oauth.OAuthError
	_, err := transport.Exchange(code)
	if err != nil {
		return nil, fmt.Errorf("Desculpe, mas nao conseguimos autentica-lo.\n Err: %s", err)
	}

	// Here we have our http client configured to make calls
	client := transport.Client()

	res, err := client.Get("https://graph.facebook.com/me")
	if err != nil {
		return nil, fmt.Errorf("Desculpe, mas nao foi possivel pegar perfil do facebook.\n Err: %s", err)
	}

	defer res.Body.Close()

	data := new(interface{}) // Its a pointer to an interface
	// The Decoder can handle a stream, unlike Unmarshal
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(data) // Passing a pointer
	if err != nil {
		return nil, fmt.Errorf("Desculpe, mas a resposta do facebook nao foi um json valido.\n Err: %s", err)
	}

	return data, nil

	/*
		profile, err := extractProfile(transport.Token, data)
		if err != nil {
			return nil, fmt.Errorf("DDesculpe, mas nao conseguimos acessar seu perfil.\n Err: %s", err)
		}
		user, err := getOrCreateUser(db, profile)
		if err != nil {
			renderHtmlOrJson(r, w, req, state, err, "Desculpe, mas nao conseguimos processar seu perfil")
			return
		}

		// After answare the request, check if this user has an pic
		err = checkOrGetPic(db, client, user)
		if err != nil {
			renderHtmlOrJson(r, w, req, state, err, "Desculpe, mas ocorreu um erro ao processar sua foto do perfil.")
			return
		}

		token, err := newToken(db, user)
		// log.Printf("user: %#v \n", user)
		// log.Printf("Token: %#v \n", token)
		if err != nil {
			renderHtmlOrJson(r, w, req, state, err, "Desculpe, mas ocorreu um erro ao criar seu token de acesso.")
			return
		}
		credentials := encodeAuth(token)

		if state == "json" {
			r.JSON(http.StatusOK, map[string]interface{}{"credentials": credentials})
			return
		}
		// COOKIES CAN'T HAVE VALUES WITH COMMA OR SPACE !!!
		cookieCredentials := &http.Cookie{Name: "credentials", Value: credentials, MaxAge: 60 * 60 * 24 * 30 * 12, Path: "/"}
		cookieMessage := &http.Cookie{Name: "message", Value: "Bem-Vindo!", MaxAge: 60 * 30, Path: "/"}
		http.SetCookie(w, cookieCredentials)
		http.SetCookie(w, cookieMessage)
		http.Redirect(w, req, "/", 302)
	*/
}
