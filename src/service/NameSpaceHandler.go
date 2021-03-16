package service

import (
	"encoding/json"
	"fmt"
	"io"
	// "net/http"
	"errors"
	"log"
	"os"
	// "bytes"
	// "reflect"
	//"github.com/davecgh/go-spew/spew"
	model "github.com/cloud-barista/cb-webtool/src/model"
	util "github.com/cloud-barista/cb-webtool/src/util"
)

// var NameSpaceUrl = "http://15.165.16.67:1323"
var NameSpaceUrl = os.Getenv("TUMBLE_URL")

// type NSInfo struct {
// 	ID          string `json:"id"`
// 	Name        string `json:"name"`
// 	Description string `json:"description"`
// }

// 저장된 namespace가 없을 때 최초 1개 생성하고 해당 namespace 정보를 return  : 검증 필요(TODO : 이미 namespace가 있어서 확인 못함)
func CreateDefaultNamespace() (*model.NSInfo, error) {
	// nsInfo := new(model.NSInfo)
	nsInfo := model.NSInfo{}

	// 사용자의 namespace 목록조회
	nsList, nsErr := GetNameSpaceList()
	if nsErr != nil {
		log.Println(" nsErr  ", nsErr)
		return &nsInfo, nsErr
	}

	if len(nsList) > 0 {
		//return &nsInfo, errors.New(101, "Namespace already exists. size="+len(nsList))
		return &nsInfo, errors.New("aaa")
	}

	// create default namespace
	nsInfo.Name = "NS-01" // default namespace name
	//nsInfo.ID = "NS-01"
	nsInfo.Description = "default name space name"
	respBody, nsCreateErr := RegNameSpace(&nsInfo)
	log.Println(" respBody  ", respBody) // respBody에 namespace Id가 있으면 할당해서 사용할 것
	if nsCreateErr != nil {
		log.Println(" nsCreateErr  ", nsCreateErr)
		return &nsInfo, nsCreateErr
	}

	return &nsInfo, nil
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
// 	nsInfo := map[string][]model.NSInfo{}
// 	fmt.Println("nsInfo type : ", reflect.TypeOf(nsInfo))
// 	json.NewDecoder(resp.Body).Decode(&nsInfo)
// 	fmt.Println("nsInfo : ", nsInfo["ns"][0].ID)
// }

// 사용자의 namespace 목록 조회
func GetNameSpaceList() ([]model.NSInfo, error) {
	fmt.Println("GetNameSpaceList start")
	url := NameSpaceUrl + "/ns"

	body, err := util.CommonHttpGet(url)
	//body := HttpGetHandler(url)

	fmt.Println(body)
	if err != nil {
		// 	// Tumblebug 접속 확인하라고
		fmt.Println(err)
		return nil, err
	}

	nsInfoList := map[string][]model.NSInfo{}
	defer body.Close()
	json.NewDecoder(body).Decode(&nsInfoList)
	//spew.Dump(body)
	fmt.Println(nsInfoList["ns"])

	return nsInfoList["ns"], nil
}

// 성공시 NsInfoList 반환
func RegNameSpace(nsInfo *model.NSInfo) (io.ReadCloser, error) {
	// buff := bytes.NewBuffer(pbytes)
	url := NameSpaceUrl + "/ns"

	fmt.Println("nsInfo : ", nsInfo)

	//body, err := util.CommonHttpPost(url, nsInfo)
	pbytes, _ := json.Marshal(nsInfo)
	body, err := util.CommonHttpPost(url, pbytes)
	if err != nil {
		fmt.Println(err)
	}
	return body, err
}
