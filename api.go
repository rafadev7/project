package main

// The main resource of this API
type Api struct {
	Version int
	Env     Env
	//User       *User
	Categories *Categories
	DB         *DB
	Auth       *Auth
}

func (a *Api) GET() *Api {
	return a
}
