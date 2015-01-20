package main

// The main resource of this API
type Api struct {
	Version int
	Env     string
	//User    *User
}

func (a *Api) GET() *Api {
	if env.Production {
		a.Env = "Production"
	} else {
		a.Env = "Development"
	}
	return a
}
