package helpers

type SignupForm struct {
	Fname          string
	Lname          string
	Email          string
	Password       string
	RPassword      string
	RememberDevice bool
}

type LoginForm struct {
	Email          string
	Password       string
	RememberDevice bool
}
