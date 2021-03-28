package model

type SshKeyInfo struct {
	ConnectionName           string        `json:"connectionName"`
	CspSshKeyName     string         `json:"cspSshKeyName"`

	Description     string         `json:"description"`
	Fingerprint     string         `json:"fingerprint"`
	ID     string         `json:"id"`
	Name     string         `json:"name"`
	PrivateKey     string         `json:"privateKey"`
	PublicKey     string         `json:"publicKey"`
	Username     string         `json:"username"`

	KeyValueInfos []KeyValueInfo `json:"keyValueList"`
}

type SshKeyInfos []SshKeyInfo