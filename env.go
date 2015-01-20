package main

type Env struct {
	Url        string
	Port       int
	Production bool
}

var devEnv = Env{
	Url:        "http://localhost:8080/",
	Port:       8080,
	Production: false,
}

// TODO
var prodEnv = Env{
	Url:        "http://localhost:8080/",
	Port:       8080,
	Production: true,
}
