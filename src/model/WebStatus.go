package model

type WebStatus struct {
	StatusCode int    `json:"code"`
	Message    string `json:"message"`
}
