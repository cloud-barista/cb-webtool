package model

// 사용자 요청(login에서 )
type ReqInfo struct {
	UserId	 string `json:"userid`
	UserName string `json:"username"`
	Password string `json:"password"`
}
