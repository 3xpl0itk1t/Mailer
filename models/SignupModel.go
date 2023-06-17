package models

type SignupUser struct {
	Auth      string `json:"auth"`
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
