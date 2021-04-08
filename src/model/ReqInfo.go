package model

type ReqInfo struct {
	UserID   string `json:"userid"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
