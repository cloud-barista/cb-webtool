package model

// ConnectionConfigData -> CloudConnectionConfigInfo 로 변경
type CloudConnectionConfigInfo struct {
	//user(username, password, email)
	ConfigName     string `json:"configname"`
	ProviderName   string `json:"providername"`
	DriverName     string `json:"drivername"`
	CredentialName string `json:"CredentialName"`
	RegionName     string `json:"RegionName"`
}
type CloudConnectionConfigInfos []CloudConnectionConfigInfo
