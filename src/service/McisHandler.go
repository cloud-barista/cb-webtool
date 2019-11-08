package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var MCISUrl = "http://127.0.0.1:1234"

type MCISInfo struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
	VMNum  string `json:"vm_num"`
}

func GetMCISList(nsid string) []MCISInfo {
	url := MCISUrl + "/ns/" + nsid + "/mcis"
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
	url := MCISUrl + "/ns/" + nsid + "mcis" + mcisId
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
