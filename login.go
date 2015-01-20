package main

import (
	"net/http"

	"code.google.com/p/goauth2/oauth"
)

type Login struct{}

var config = &oauth.Config{
	ClientId:     "560046220732101",
	ClientSecret: "18d10b619523227e65ecf5b38fc18f90",
	//RedirectURL:   // It is setted dynamicaly on LoginHandler based on server current url
	Scope:    "basic_info email",
	AuthURL:  "https://www.facebook.com/dialog/oauth",
	TokenURL: "https://graph.facebook.com/oauth/access_token",
}

var transport = &oauth.Transport{
	Config:    config,
	Transport: http.DefaultTransport,
}

// Will come to this method
// when user ask for [GET] /user/login
// "Using methods of the resource"
// Resource should have no child called login...
func (l *Login) GET(w http.ResponseWriter, req *http.Request, env Env) {
	state := req.URL.Query().Get("method")
	if state == "" { // The default is Json method, but it could be html too
		state = "json"
	}

	// Set based on the server current url
	transport.Config.RedirectURL = env.Url + "logincallback/"

	http.Redirect(w, req, transport.Config.AuthCodeURL(state), 302)
}
