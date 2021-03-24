package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	// "os"
	// "bytes"
	// "reflect"
	//"github.com/davecgh/go-spew/spew"
	model "github.com/cloud-barista/cb-webtool/src/model"
	util "github.com/cloud-barista/cb-webtool/src/util"
)

// var NameSpaceUrl = "http://15.165.16.67:1323"
// var NameSpaceUrl = os.Getenv("TUMBLE_URL")

// type NSInfo struct {
// 	ID          string `json:"id"`
// 	Name        string `json:"name"`
// 	Description string `json:"description"`
// }

// 저장된 namespace가 없을 때 최초 1개 생성하고 해당 namespace 정보를 return  : 검증 필요(TODO : 이미 namespace가 있어서 확인 못함)
func CreateDefaultNamespace() (*model.NameSpaceInfo, error) {
	// nsInfo := new(model.NSInfo)
	nameSpaceInfo := model.NameSpaceInfo{}

	// 사용자의 namespace 목록조회
	nsList, nsErr := GetNameSpaceList()
	if nsErr != nil {
		log.Println(" nsErr  ", nsErr)
		return &nameSpaceInfo, nsErr
	}

	if len(nsList) > 0 {
		//return &nameSpaceInfo, errors.New(101, "Namespace already exists. size="+len(nsList))
		return &nameSpaceInfo, errors.New("aaa")
	}

	// create default namespace
	nameSpaceInfo.Name = "NS-01" // default namespace name
	//nameSpaceInfo.ID = "NS-01"
	nameSpaceInfo.Description = "default name space name"
	respBody, nsCreateErr := RegNameSpace(&nameSpaceInfo)
	log.Println(" respBody  ", respBody) // respBody에 namespace Id가 있으면 할당해서 사용할 것
	if nsCreateErr != nil {
		log.Println(" nsCreateErr  ", nsCreateErr)
		return &nameSpaceInfo, nsCreateErr
	}

	return &nameSpaceInfo, nil
}

// func GetNS(nsID string) model.NSInfo {
// 	url := NameSpaceUrl + "ns" + nsID

// 	body := HttpGetHandler(url)
// 	defer body.Close()
// 	nsInfo := model.NSInfo{}
// 	json.NewDecoder(body).Decode(&nsInfo)
// 	fmt.Println("nsInfo : ", nsInfo.ID)
// 	return nsInfo

// }

// func GetNSList() []model.NSInfo {
// 	url := NameSpaceUrl + "/ns"
// 	fmt.Println("============= NameSpace URL =============", url)
// 	// authInfo := controller.AuthenticationHandler()
// 	// req, err := http.NewRequest("GET", url, nil)
// 	// if err != nil {

// 	// }
// 	// req.Header.Add("Authorization", authInfo)
// 	// client := &http.Client{}
// 	// resp, err := client.Do(req)
// 	// fmt.Println("=============result GetNSList =============", resp)
// 	// //spew.Dump(resp)
// 	// if err != nil {
// 	// 	fmt.Println("========= GetNSList Error : ", err)
// 	// 	fmt.Println("request URL : ", url)
// 	// 	return nil
// 	// }

// 	// defer resp.Body.Close()
// 	body := HttpGetHandler(url)
// 	nsInfo := map[string][]model.NSInfo{}
// 	defer body.Close()
// 	json.NewDecoder(body).Decode(&nsInfo)
// 	//spew.Dump(body)
// 	return nsInfo["ns"]

// }

// func GetNSCnt() int {
// 	url := NameSpaceUrl + "/ns"
// 	fmt.Println("============= NameSpace URL =============", url)

// 	// defer resp.Body.Close()
// 	body := HttpGetHandler(url)
// 	nsInfo := map[string][]model.NSInfo{}
// 	defer body.Close()
// 	json.NewDecoder(body).Decode(&nsInfo)
// 	//spew.Dump(body)
// 	if nsInfo["ns"] == nil {
// 		return 0
// 	} else {
// 		return len(nsInfo["ns"])

// 	}

// }

// 안쓰는 function인듯.
// func RequestGet(url string) {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		fmt.Println("request URL : ", url)
// 	}

// 	defer resp.Body.Close()
// 	nameSpaceInfo := map[string][]model.NSInfo{}
// 	fmt.Println("nameSpaceInfo type : ", reflect.TypeOf(nameSpaceInfo))
// 	json.NewDecoder(resp.Body).Decode(&nameSpaceInfo)
// 	fmt.Println("nameSpaceInfo : ", nameSpaceInfo["ns"][0].ID)
// }

// 사용자의 namespace 목록 조회
func GetNameSpaceList() ([]model.NameSpaceInfo, error) {
	fmt.Println("GetNameSpaceList start")
	url := util.TUMBLEBUG + "/ns"

	body, err := util.CommonHttpGet(url)
	//body := HttpGetHandler(url)

	fmt.Println(body)
	if err != nil {
		// 	// Tumblebug 접속 확인하라고
		fmt.Println(err)
		return nil, err
	}

	nameSpaceInfoList := map[string][]model.NameSpaceInfo{}
	defer body.Close()
	json.NewDecoder(body).Decode(&nameSpaceInfoList)
	//spew.Dump(body)
	fmt.Println(nameSpaceInfoList["ns"])

	return nameSpaceInfoList["ns"], nil
}

// NameSpace 등록
func RegNameSpace(nameSpaceInfo *model.NameSpaceInfo) (io.ReadCloser, error) {
	// buff := bytes.NewBuffer(pbytes)
	url := util.TUMBLEBUG + "/ns"

	fmt.Println("nameSpaceInfo : ", nameSpaceInfo)

	//body, err := util.CommonHttpPost(url, nameSpaceInfo)
	pbytes, _ := json.Marshal(nameSpaceInfo)
	body, err := util.CommonHttpPost(url, pbytes)
	if err != nil {
		fmt.Println(err)
	}
	return body, err
}

// NameSpace 수정
func UpdateNameSpace(nameSpaceInfo *model.NameSpaceInfo) (io.ReadCloser, error) {
	// buff := bytes.NewBuffer(pbytes)
	url := util.TUMBLEBUG + "/ns"

	fmt.Println("nameSpaceInfo : ", nameSpaceInfo)

	//body, err := util.CommonHttpPost(url, nameSpaceInfo)
	pbytes, _ := json.Marshal(nameSpaceInfo)
	body, err := util.CommonHttp(url, pbytes, http.MethodPut)

	if err != nil {
		fmt.Println(err)
	}
	return body, err
}

// NameSpace 삭제
func DelNameSpace(nameSpaceID string) (io.ReadCloser, error) {
	// buff := bytes.NewBuffer(pbytes)
	url := util.TUMBLEBUG + "/ns/" + nameSpaceID

	fmt.Println("nameSpaceID : ", nameSpaceID)

	//body, err := util.CommonHttpPost(url, nsInfo)

	// 경로안에 parameter가 있어 추가 param없이 호출 함.
	body, err := util.CommonHttpDelete(url, nil)
	// body, err := util.CommonHttpDelete(url, pbytes)
	if err != nil {
		fmt.Println(err)
	}
	return body, err
}
