package service

import (
	"encoding/json"
	"fmt"
	"log"
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
func GetMcisList(nameSpaceID string)  ([]model.MCISInfo, int ) {
// func GetMCISList(nsid string) []MCISInfo {
	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/mcis"
	resp, err := util.CommonHttp(url, nil, http.MethodGet)
	// resp, err := util.CommonHttpWithoutParam(url, http.MethodGet)

	var respStatus int
	if err != nil {
		fmt.Println(err)
		//respStatus = 500
	}

	respBody := resp.Body
	respStatus = resp.StatusCode
	
	mcisList := map[string][]model.MCISInfo{}
	json.NewDecoder(respBody).Decode(&mcisList)
	fmt.Println(mcisList["mcis"])
	log.Println(respBody)
	util.DisplayResponse(resp)// 수신내용 확인

	return mcisList["mcis"], respStatus
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

	// resp, err := util.CommonHttp(url, nil, http.MethodGet)
	resp, err := util.CommonHttpWithoutParam(url, http.MethodGet)
	
	// defer body.Close()

	if err != nil {
		fmt.Println(err)
	}
	util.DisplayResponse(resp)// 수신내용 확인

	respBody := resp.Body
	respStatus := resp.StatusCode

	mcisInfo := model.MCISInfo{}
	json.NewDecoder(respBody).Decode(&mcisInfo)
	fmt.Println(mcisInfo)
	
	
	// resultBody, err := ioutil.ReadAll(respBody)
	// if err == nil {
	// 	str := string(resultBody)
	// 	println(str)
	// }
	// pbytes, _ := json.Marshal(respBody)
	// fmt.Println(string(pbytes))

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
