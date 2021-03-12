package model

type NSInfo struct {
	ID          string `json:"id"`
	Name        string `json:"name" validate:"required`
	Description string `json:"description"`
}

type NSInfos []NSInfo
