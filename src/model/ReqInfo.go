package model

// 사용자 요청(login에서 )
type ReqInfo struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}
