package spider

type CredentialInfo struct {
	//user(username, password, email)
	CredentialName   string         `json:"CredentialName"`
	ProviderName     string         `json:"ProviderName"`
	KeyValueInfoList SpKeyValueList `json:"KeyValueInfoList"`
}
type CredentialInfos []CredentialInfo
