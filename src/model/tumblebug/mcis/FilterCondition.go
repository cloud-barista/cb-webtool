package mcis

type FilterCondition struct {
	condition Operation `json:"condition"`
	metric    string    `json:"metric"`
}
