package model

// Namespace 목록
type NSInfo struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type NSInfoList []NSInfo
