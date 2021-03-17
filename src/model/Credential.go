package model

type Credential struct {
	//user(username, password, email)
	CredentialName string             `json:"CredentialName"`
	ProviderName   string             `json:"ProviderName"`
	KeyValueInfo   []KeyValueInfoList `json:"KeyValueInfoList"`
}
type Credentials []Credential
