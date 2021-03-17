package model

type Driver struct {
	//user(username, password, email)
	DriverName        string `json:"DriverName"`
	ProviderName      string `json:"ProviderName"`
	DriverLibFileName string `json:"DriverLibFileName"`
}
type Drivers []Driver
