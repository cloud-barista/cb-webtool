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
func GetVnetList(nameSpaceID string) ([]model.VNetInfo, int) {
	fmt.Println("GetVnetList ************ : ")
	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/resources/vNet"

	pbytes, _ := json.Marshal(nameSpaceID)
	// body, err := util.CommonHttpGet(url)
	resp, err := util.CommonHttp(url, pbytes, http.MethodGet)

	if err != nil {
		fmt.Println(err)
	}
	// defer body.Close()
	respBody := resp.Body
	respStatus := resp.StatusCode

	// return respBody, respStatus
	log.Println(respBody)
	vNetInfoList := map[string][]model.VNetInfo{}
	json.NewDecoder(respBody).Decode(&vNetInfoList)
	//spew.Dump(body)
	fmt.Println(vNetInfoList["vNet"])

	return vNetInfoList["vNet"], respStatus

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
		//respStatus = 500
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
// func RegVpc(nameSpaceID string, vnetRegInfo *model.VNetRegInfo) (io.ReadCloser, int) {
func RegVpc(nameSpaceID string, vnetRegInfo *model.VNetRegInfo) (model.VNetInfo, int) {
	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/resources/vNet"

	// fmt.Println("vnetInfo : ", vnetInfo)

	pbytes, _ := json.Marshal(vnetRegInfo)
	fmt.Println(string(pbytes))
	resp, err := util.CommonHttp(url, pbytes, http.MethodPost)

	if err != nil {
		fmt.Println(err)
	}

	respBody := resp.Body
	respStatusCode := resp.StatusCode
	respStatus := resp.Status
	log.Println("respStatusCode = ", respStatusCode)
	log.Println("respStatus = ", respStatus)

	// 응답에 생성한 객체값이 옴
	vNetInfo := model.VNetInfo{}
	json.NewDecoder(respBody).Decode(&vNetInfo)
	fmt.Println(vNetInfo)
	// return respBody, respStatusCode
	return vNetInfo, respStatusCode
}

// vpc 삭제
func DelVpc(nameSpaceID string, vNetID string) (io.ReadCloser, int) {
	// if ValidateString(vNetID) != nil {
	if len(vNetID) == 0 {
		log.Println("vNetID 가 없으면 해당 namespace의 모든 vpc가 삭제되므로 처리할 수 없습니다.")
		return nil, 4040
	}
	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/resources/vNet/" + vNetID

	fmt.Println("vNetID : ", vNetID)

	pbytes, _ := json.Marshal(vNetID)
	resp, err := util.CommonHttp(url, pbytes, http.MethodDelete)

	if err != nil {
		fmt.Println(err)
	}
	// return body, err
	respBody := resp.Body
	respStatusCode := resp.StatusCode
	respStatus := resp.Status
	log.Println("respStatusCode = ", respStatusCode)
	log.Println("respStatus = ", respStatus)

	// resultBody, err := ioutil.ReadAll(respBody)
	// if err == nil {
	// 	str := string(resultBody)
	// 	println(str)
	// }
	// pbytes, _ := json.Marshal(respBody)
	// fmt.Println(string(pbytes))

	return respBody, respStatusCode

}

// 해당 namespace의 SecurityGroup 목록 조회
func GetSecurityGroupList(nameSpaceID string) ([]model.SecurityGroupInfo, int) {
	fmt.Println("GetSecurityGroupList ************ : ")
	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/resources/securityGroup"

	pbytes, _ := json.Marshal(nameSpaceID)
	// body, err := util.CommonHttpGet(url)
	resp, err := util.CommonHttp(url, pbytes, http.MethodGet)

	if err != nil {
		fmt.Println(err)
	}
	// defer body.Close()
	respBody := resp.Body
	respStatus := resp.StatusCode

	// return respBody, respStatus
	log.Println(respBody)
	securityGroupList := map[string][]model.SecurityGroupInfo{}

	json.NewDecoder(respBody).Decode(&securityGroupList)
	//spew.Dump(body)
	fmt.Println(securityGroupList["securityGroup"])

	return securityGroupList["securityGroup"], respStatus

}

// SecurityGroup 상세 조회
func GetSecurityGroupData(nameSpaceID string, securityGroupID string) (model.SecurityGroupInfo, int) {
	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/resources/securityGroup/" + securityGroupID

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

	securityGroupInfo := model.SecurityGroupInfo{}
	json.NewDecoder(respBody).Decode(&securityGroupInfo)
	fmt.Println(securityGroupInfo)

	return securityGroupInfo, respStatus
}

// SecurityGroup 등록
func RegSecurityGroup(nameSpaceID string, securityGroupRegInfo *model.SecurityGroupRegInfo) (model.SecurityGroupInfo, int) {
	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/resources/securityGroup"

	// fmt.Println("vnetInfo : ", vnetInfo)

	pbytes, _ := json.Marshal(securityGroupRegInfo)
	fmt.Println(string(pbytes))
	resp, err := util.CommonHttp(url, pbytes, http.MethodPost)

	if err != nil {
		fmt.Println(err)
	}

	respBody := resp.Body
	respStatusCode := resp.StatusCode
	respStatus := resp.Status
	log.Println("respStatusCode = ", respStatusCode)
	log.Println("respStatus = ", respStatus)

	// 응답에 생성한 객체값이 옴
	securityGroupInfo := model.SecurityGroupInfo{}
	json.NewDecoder(respBody).Decode(&securityGroupInfo)
	fmt.Println(securityGroupInfo)
	// return respBody, respStatusCode
	return securityGroupInfo, respStatusCode
}

// SecurityGroup 삭제
func DelSecurityGroup(nameSpaceID string, securityGroupID string) (io.ReadCloser, int) {
	// if ValidateString(vNetID) != nil {
	if len(securityGroupID) == 0 {
		log.Println("securityGroupID 가 없으면 해당 namespace의 모든 securityGroup이 삭제되므로 처리할 수 없습니다.")
		return nil, 4040
	}
	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/resources/securityGroup/" + securityGroupID

	fmt.Println("securityGroupID : ", securityGroupID)

	pbytes, _ := json.Marshal(securityGroupID)
	resp, err := util.CommonHttp(url, pbytes, http.MethodDelete)

	if err != nil {
		fmt.Println(err)
	}
	// return body, err
	respBody := resp.Body
	respStatusCode := resp.StatusCode
	respStatus := resp.Status
	log.Println("respStatusCode = ", respStatusCode)
	log.Println("respStatus = ", respStatus)

	return respBody, respStatusCode

}
