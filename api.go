package main

// The main resource of this API
type Api struct {
	Version int
	//User    *User
}

func (a *Api) GET() *Api {
	return a
}
