package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//var TumblebugUrl = "http://15.165.16.67:1323"
//var SpiderUrl = "http://15.165.16.67:1024"

//var SpiderUrl = os.Getenv("SPIDER_URL")
//var TumblebugUrl = os.Getenv("TUMBLE_URL")

type MCISInfo struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
	VMNum  string `json:"vm_num"`
}

func GetMCISList(nsid string) []MCISInfo {
	url := TumblebugUrl + "/ns/" + nsid + "/mcis"
	fmt.Println("GETMCISLIST URL : ", url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("request URL : ", url)
	}

	defer resp.Body.Close()
	info := map[string][]MCISInfo{}
	json.NewDecoder(resp.Body).Decode(&info)
	fmt.Println("nsInfo : ", info["mcis"][0].ID)
	return info["mcis"]

}

func GetMCIS(nsid string, mcisId string) []MCISInfo {
	url := TumblebugUrl + "/ns/" + nsid + "/mcis/" + mcisId
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("request URL : ", url)
	}

	defer resp.Body.Close()
	info := map[string][]MCISInfo{}
	json.NewDecoder(resp.Body).Decode(&info)
	fmt.Println("info : ", info["mcis"][0].ID)
	return info["ns"]

}

func GetVMStatus(vm_name string, connectionConfig string) string {
	url := SpiderUrl + "/vmstatus/" + vm_name + "?connection_name=" + connectionConfig
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("request URL : ", url)
	}

	defer resp.Body.Close()
	info := map[string]MCISInfo{}
	json.NewDecoder(resp.Body).Decode(&info)
	fmt.Println("VM Status : ", info["status"].Status)
	return info["status"].Status

}
