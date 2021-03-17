package model

type Credential struct {
	//user(username, password, email)
	CredentialName string             `json:"CredentialName"`
	ProviderName   string             `json:"ProviderName"`
	KeyValueInfoList   []KeyValueInfoList `json:"KeyValueInfoList"`
}
type Credentials []Credential
