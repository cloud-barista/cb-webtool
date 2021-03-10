package models

type User struct {
	//user(username, password, email)
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Users []User
