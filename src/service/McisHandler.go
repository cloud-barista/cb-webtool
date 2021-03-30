package service

import (
	"encoding/json"
	"fmt"

	"net/http"
	// "os"
	util "github.com/cloud-barista/cb-webtool/src/util"
	model "github.com/cloud-barista/cb-webtool/src/model"
)

//var MCISUrl = "http://15.165.16.67:1323"
//var SPiderUrl = "http://15.165.16.67:1024"

// var SpiderUrl = os.Getenv("SPIDER_URL")// util.SPIDER
// var MCISUrl = os.Getenv("TUMBLE_URL")// util.TUMBLEBUG

// MCIS 목록 조회
func GetMCISList(nameSpaceID string)  ([]model.MCISInfo, int ) {
// func GetMCISList(nsid string) []MCISInfo {
	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/mcis"
	resp, err := util.CommonHttp(url, nil, http.MethodGet)

	var respStatus int
	if err != nil {
		fmt.Println(err)
		//respStatus = 500
	}

	respBody := resp.Body
	respStatus = resp.StatusCode

	mcisInfo := map[string][]model.MCISInfo{}
	json.NewDecoder(respBody).Decode(&mcisInfo)
	fmt.Println(mcisInfo["connectionconfig"])

	return mcisInfo["mcis"], respStatus
}

func GetMCIS(nameSpaceID string, mcisID string) (model.MCISInfo, int ) {
// func GetMCIS(nsid string, mcisId string) []MCISInfo {
	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/mcis/" + mcisID
// 	// resp, err := http.Get(url)
// 	// if err != nil {
// 	// 	fmt.Println("request URL : ", url)
// 	// }

// 	// defer resp.Body.Close()
// 	body := HttpGetHandler(url)
// 	defer body.Close()
// 	info := map[string][]MCISInfo{}
// 	json.NewDecoder(body).Decode(&info)
// 	fmt.Println("info : ", info["mcis"][0].ID)
// 	return info["ns"]

	resp, err := util.CommonHttp(url, nil, http.MethodGet)
	// defer body.Close()

	if err != nil {
		fmt.Println(err)
	}

	respBody := resp.Body
	respStatus := resp.StatusCode

	mcisInfo := model.MCISInfo{}
	json.NewDecoder(respBody).Decode(&mcisInfo)
	fmt.Println(mcisInfo)
	return mcisInfo, respStatus
}

// func GetVMStatus(vm_name string, connectionConfig string) string {
// 	url := SpiderUrl + "/vmstatus/" + vm_name + "?connection_name=" + connectionConfig
// 	// resp, err := http.Get(url)
// 	// if err != nil {
// 	// 	fmt.Println("request URL : ", url)
// 	// }

// 	// defer resp.Body.Close()
// 	body := HttpGetHandler(url)
// 	defer body.Close()
// 	info := map[string]MCISInfo{}
// 	json.NewDecoder(body).Decode(&info)
// 	fmt.Println("VM Status : ", info["status"].Status)
// 	return info["status"].Status

// }
