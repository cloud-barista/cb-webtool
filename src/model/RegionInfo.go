package model

type RegionInfo struct {
	RegionName   string             `json:"RegionName"`
	ProviderName string             `json:"ProviderName"`
	KeyValueInfo []KeyValueInfoList `json:"KeyValueInfoList"`
}

type RegionInfos []RegionInfo
