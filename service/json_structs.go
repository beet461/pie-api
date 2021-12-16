package service

type Test struct {
	TestRespond string
}

type UserData struct {
	Email     string
	Password  string
	Firstname string
	Lastname  string
	Id        string
}

type Customise struct {
	Colorscheme string
	Id          string
}

type Account struct {
	Signin UserData
	Cust   Customise
}

type Response struct {
	Code    int
	Account Account
}
