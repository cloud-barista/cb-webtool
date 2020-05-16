package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	//"io/ioutil"
	//"github.com/davecgh/go-spew/spew"
)

//var CloudConnectionUrl = "http://15.165.16.67:1024"
var CloudConnectionUrl = os.Getenv("SPIDER_URL")

type CloudConnectionInfo struct {
	ID             string `json:"id"`
	ConfigName     string `json:"ConfigName"`
	ProviderName   string `json:"ProviderName"`
	DriverName     string `json:"DriverName"`
	CredentialName string `json:"CredentialName"`
	RegionName     string `json:"RegionName"`
	Description    string `json:"description"`
}
type KeyValueInfo struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
}
type RegionInfo struct {
	RegionName       string `json:"RegionName"`
	ProviderName     string `json:"ProviderName"`
	KeyValueInfoList []KeyValueInfo
}
type RESP struct {
	Region []struct {
		RegionName       string         `json:"RegionName"`
		ProviderName     string         `json:"ProviderName"`
		KeyValueInfoList []KeyValueInfo `json:"KeyValueInfoList"`
	} `json:"region"`
}

func GetConnectionconfig(drivername string) CloudConnectionInfo {
	url := NameSpaceUrl + "/driver/" + drivername

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("request URL : ", url)
	}

	defer resp.Body.Close()
	nsInfo := CloudConnectionInfo{}

	json.NewDecoder(resp.Body).Decode(&nsInfo)
	fmt.Println("nsInfo : ", nsInfo.ID)
	return nsInfo

}

func GetConnectionList() []CloudConnectionInfo {
	url := CloudConnectionUrl + "/connectionconfig"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("request URL : ", url)
	}

	defer resp.Body.Close()
	nsInfo := map[string][]CloudConnectionInfo{}
	json.NewDecoder(resp.Body).Decode(&nsInfo)
	fmt.Println("nsInfo : ", nsInfo["ns"][0].ID)
	return nsInfo["ns"]

}

func GetDriverReg() []CloudConnectionInfo {
	url := NameSpaceUrl + "/driver"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("request URL : ", url)
	}

	defer resp.Body.Close()
	nsInfo := map[string][]CloudConnectionInfo{}
	json.NewDecoder(resp.Body).Decode(&nsInfo)
	// fmt.Println("nsInfo : ", nsInfo["ns"][0].ID)
	return nsInfo["ns"]

}

func GetCredentialList() []CloudConnectionInfo {
	url := CloudConnectionUrl + "/credential"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("request URL : ", url)
	}

	defer resp.Body.Close()
	nsInfo := map[string][]CloudConnectionInfo{}
	json.NewDecoder(resp.Body).Decode(&nsInfo)
	// fmt.Println("nsInfo : ", nsInfo["ns"][0].ID)
	return nsInfo["ns"]

}
func GetRegionList() []RegionInfo {
	url := CloudConnectionUrl + "/region"
	fmt.Println("=========== Get Start Region List : ", url)
	resp, err := http.Get(url)
	//spew.Dump(resp.Body)
	if err != nil {
		fmt.Println("request URL : ", url)
	}

	defer resp.Body.Close()

	//bytes, _ := ioutil.ReadAll(resp.Body)
	//str := string(bytes)
	//fmt.Println(str.region)
	nsInfo := RESP{}
	//spew.Dump(nsInfo)
	json.NewDecoder(resp.Body).Decode(&nsInfo)
	var info []RegionInfo
	// fmt.Println("nsInfo : ", nsInfo["ns"][0].ID)
	for _, item := range nsInfo.Region {
		// var kv
		// for i, v := range item.KeyValueInfoList{
		// 	k := KeyValueInfo{
		// 		Key: v.Key,
		// 		Value: v.Value
		// 	}
		// }
		reg := RegionInfo{
			RegionName:   item.RegionName,
			ProviderName: item.ProviderName,
		}
		info = append(info, reg)
	}
	fmt.Println("info region list : ", info)
	return info

}

func GetCredentialReg() []CloudConnectionInfo {
	url := CloudConnectionUrl + "/credential"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("request URL : ", url)
	}

	defer resp.Body.Close()
	nsInfo := map[string][]CloudConnectionInfo{}
	json.NewDecoder(resp.Body).Decode(&nsInfo)
	fmt.Println("nsInfo : ", nsInfo["ns"][0].ID)
	return nsInfo["ns"]

}

// func RegNS() error {

// }

// func RequestGet(url string) {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		fmt.Println("request URL : ", url)
// 	}

// 	defer resp.Body.Close()
// 	nsInfo := map[string][]NSInfo{}
// 	fmt.Println("nsInfo type : ", reflect.TypeOf(nsInfo))
// 	json.NewDecoder(resp.Body).Decode(&nsInfo)
// 	fmt.Println("nsInfo : ", nsInfo["ns"][0].ID)

// 	// data, err := ioutil.ReadAll(resp.Body)
// 	// if err != nil {
// 	// 	fmt.Println("Get Data Error")
// 	// }
// 	// fmt.Println("GetData : ", string(data))

// }
