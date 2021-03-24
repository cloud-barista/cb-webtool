package service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	// "math"
	"net/http"
	// "strconv"
	// "sync"
	//"io/ioutil"
	//"github.com/davecgh/go-spew/spew"
	model "github.com/cloud-barista/cb-webtool/src/model"
	util "github.com/cloud-barista/cb-webtool/src/util"
)

// 해당 namespace의 vpc 목록 조회
//func GetVnetList(nameSpaceID string) (io.ReadCloser, error) {
func GetVnetList(nameSpaceID string) ([]model.VNetInfo, error) {
	fmt.Println("GetVnetList ************ : ")
	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/resources/vNet"

	pbytes, _ := json.Marshal(nameSpaceID)
	// body, err := util.CommonHttpGet(url)
	body, err := util.CommonHttp(url, pbytes, http.MethodGet)

	if err != nil {
		fmt.Println(err)
	}
	log.Println(body)
	// return body, err

	vNetInfoList := map[string][]model.VNetInfo{}
	defer body.Close()
	json.NewDecoder(body).Decode(&vNetInfoList)
	//spew.Dump(body)
	fmt.Println(vNetInfoList["vNet"])

	return vNetInfoList["vNet"], nil

}

// vpc 상세 조회-> ResourceHandler로 이동
func GetVpcData(nameSpaceID string, vNetID string) (model.VNetInfo, int) {
	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/resources/vNet/" + vNetID

	fmt.Println("nameSpaceID : ", nameSpaceID)

	// pbytes, _ := json.Marshal(nameSpaceID)
	// body, err := util.CommonHttpGet(url)
	resp, err := util.CommonHttpWithoutParam(url, http.MethodGet)
	var respStatus int
	if err != nil {
		fmt.Println(err)
		respStatus = 500
	}

	respBody := resp.Body
	respStatus = resp.StatusCode

	vNetInfo := model.VNetInfo{}
	// json.NewDecoder(body).Decode(&vNetInfo)
	json.NewDecoder(respBody).Decode(&vNetInfo)
	fmt.Println(vNetInfo)

	// return vNetInfo, err
	return vNetInfo, respStatus
}

// vpc 등록
func RegVpc(nameSpaceID string, vnetInfo *model.VNetInfo) (io.ReadCloser, error) {
	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/resources/vNet"

	fmt.Println("nameSpaceID : ", nameSpaceID)

	pbytes, _ := json.Marshal(vnetInfo)
	body, err := util.CommonHttp(url, pbytes, http.MethodPost)

	if err != nil {
		fmt.Println(err)
	}
	return body, err
}

// vpc 삭제
func DelVpc(nameSpaceID string, vNetID string) (io.ReadCloser, error) {
	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/resources/vNet" + vNetID

	fmt.Println("nameSpaceID : ", nameSpaceID)

	pbytes, _ := json.Marshal(vNetID)
	body, err := util.CommonHttp(url, pbytes, http.MethodDelete)

	if err != nil {
		fmt.Println(err)
	}
	return body, err
}
