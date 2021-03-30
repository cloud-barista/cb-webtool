package service

import (
	"encoding/json"
	"fmt"
	"io"
	// "math"
	"net/http"
	// "strconv"
	// "sync"
	//"io/ioutil"
	//"github.com/davecgh/go-spew/spew"
	model "github.com/cloud-barista/cb-webtool/src/model"
	util "github.com/cloud-barista/cb-webtool/src/util"
)

//var CloudConnectionUrl = "http://15.165.16.67:1024"
// var CloudConnectionUrl = os.Getenv("SPIDER_URL")	// Const.go로 이동
// var TumbleUrl = os.Getenv("TUMBLE_URL") // Const.go로 이동

// type KeyValueInfo struct {
// 	Key   string `json:"Key"`
// 	Value string `json:"Value"`
// }
// type RegionInfo struct {
// 	RegionName       string `json:"RegionName"`
// 	ProviderName     string `json:"ProviderName"`
// 	KeyValueInfoList []KeyValueInfo
// }

// 뭐에쓰는 거지?
type RESP struct {
	Region []struct {
		RegionName       string                   `json:"RegionName"`
		ProviderName     string                   `json:"ProviderName"`
		KeyValueInfoList []model.KeyValueInfoList `json:"KeyValueInfoList"`
	} `json:"region"`
}

// 뭐에쓰는 거지?
type ImageRESP struct {
	Image []struct {
		id               string                   `json:"id"`
		name             string                   `json:"name"`
		connectionName   string                   `json:"connectionName"`
		cspImageId       string                   `json:"cspImageId"`
		cspImageName     string                   `json:"cspImageName"`
		description      string                   `json:"description"`
		guestOS          string                   `json:"guestOS"`
		status           string                   `json:"status"`
		KeyValueInfoList []model.KeyValueInfoList `json:"KeyValueList"`
	} `json:"image"`
}
type Image struct {
	id               string                   `json:"id"`
	name             string                   `json:"name"`
	connectionName   string                   `json:"connectionName"`
	cspImageId       string                   `json:"cspImageId"`
	cspImageName     string                   `json:"cspImageName"`
	description      string                   `json:"description"`
	guestOS          string                   `json:"guestOS"`
	status           string                   `json:"status"`
	KeyValueInfoList []model.KeyValueInfoList `json:"KeyValueList"`
}
type IPStackInfo struct {
	IP          string  `json:"ip"`
	Lat         float64 `json:"latitude"`
	Long        float64 `json:"longitude"`
	CountryCode string  `json:"country_code"`
	VMName      string
	VMID        string
	Status      string
}

// 목록 : ListData
// 1개 : Data
// 등록 : Reg
// 삭제 : Del

// Cloud Provider 목록
func GetCloudOSList() ([]string, int) {

	// CloudConnectionUrl == SPIDER
	url := util.SPIDER + "/" + "cloudos"
	// fmt.Println("=========== GetCloudOSList : ", url)

	resp, err := util.CommonHttp(url, nil, http.MethodGet)
	// defer resp.Close()

	var respStatus int
	if err != nil {
		fmt.Println(err)
		//respStatus = 500
	}

	respBody := resp.Body
	respStatus = resp.StatusCode

	cloudOs := map[string][]string{}
	json.NewDecoder(respBody).Decode(&cloudOs)
	fmt.Println(cloudOs["cloudos"])
	// TODO : mock을 추가할 것
	return cloudOs["cloudos"], respStatus
}

// 현재 설정된 connection 목록 GetConnectionConfigListData -> GetCloudConnectionConfigList로 변경
func GetCloudConnectionConfigList() ([]model.CloudConnectionConfigInfo, int ) {

	// CloudConnectionUrl == SPIDER
	url := util.SPIDER + "/" + "connectionconfig"
	// fmt.Println("=========== GetCloudConnectionConfigList : ", url)

	resp, err := util.CommonHttp(url, nil, http.MethodGet)
	// defer body.Close()

	if err != nil {
		fmt.Println(err)
	}
	
	respBody := resp.Body
	respStatus := resp.StatusCode

	cloudConnectionConfigInfo := map[string][]model.CloudConnectionConfigInfo{}
	json.NewDecoder(respBody).Decode(&cloudConnectionConfigInfo)
	fmt.Println(cloudConnectionConfigInfo["connectionconfig"])

	return cloudConnectionConfigInfo["connectionconfig"], respStatus
}

// Connection 상세
func GetCloudConnectionConfigData(configName string) (model.CloudConnectionConfigInfo, int ) {
	url := util.SPIDER + "/connectionconfig/" + configName
	fmt.Println("=========== GetCloudConnectionConfigData : ", configName)

	resp, err := util.CommonHttp(url, nil, http.MethodGet)
	// defer body.Close()

	if err != nil {
		fmt.Println(err)
	}

	respBody := resp.Body
	respStatus := resp.StatusCode

	cloudConnectionConfigInfo := model.CloudConnectionConfigInfo{}
	json.NewDecoder(respBody).Decode(&cloudConnectionConfigInfo)
	fmt.Println(cloudConnectionConfigInfo)
	return cloudConnectionConfigInfo, respStatus
}

// CloudConnectionConfigInfo 등록
func RegCloudConnectionConfig(cloudConnectionConfigInfo *model.CloudConnectionConfigInfo) (io.ReadCloser, int) {
	// buff := bytes.NewBuffer(pbytes)
	url := util.SPIDER + "/connectionconfig"

	fmt.Println("cloudConnectionConfigInfo : ", cloudConnectionConfigInfo)

	// body, err := util.CommonHttpPost(url, regionInfo)
	pbytes, _ := json.Marshal(cloudConnectionConfigInfo)
	// body, err := util.CommonHttpPost(url, pbytes)
	resp, err := util.CommonHttp(url, pbytes, http.MethodPost)
	if err != nil {
		fmt.Println(err)
	}

	respBody := resp.Body
	respStatus := resp.StatusCode

	return respBody, respStatus
}

// CloudConnectionConfigInfo 삭제
func DelCloudConnectionConfig(configName string) (io.ReadCloser, int) {
	// buff := bytes.NewBuffer(pbytes)
	url := util.SPIDER + "/connectionconfig/" + configName

	fmt.Println("DelCloudConnectionConfig : ", configName)

	// body, err := util.CommonHttpPost(url, regionInfo)

	pbytes, _ := json.Marshal(configName)
	// body, err := util.CommonHttpDelete(url, pbytes)
	resp, err := util.CommonHttp(url, pbytes, http.MethodDelete)

	if err != nil {
		fmt.Println(err)
	}
	// return body, err
	respBody := resp.Body
	respStatus := resp.StatusCode

	return respBody, respStatus
}

// 현재 설정된 region 목록
func GetRegionList() ([]model.RegionInfo, int ) {

	// SPIDER == SPIDER
	url := util.SPIDER + "/" + "region"
	// fmt.Println("=========== GetRegionListData : ", url)

	resp, err := util.CommonHttp(url, nil, http.MethodGet)
	// defer body.Close()

	if err != nil {
		fmt.Println(err)
	}

	respBody := resp.Body
	respStatus := resp.StatusCode

	regionList := map[string][]model.RegionInfo{}
	json.NewDecoder(respBody).Decode(&regionList)
	fmt.Println(regionList["region"])

	return regionList["region"], respStatus
}

func GetRegionData(regionName string) ( model.RegionInfo, int ) {
	url := util.SPIDER + "/region/" + regionName
	fmt.Println("=========== GetRegionData : ", regionName)

	resp, err := util.CommonHttp(url, nil, http.MethodGet)
	// defer body.Close()

	if err != nil {
		fmt.Println(err)
	}

	respBody := resp.Body
	respStatus := resp.StatusCode

	// regionList := map[string][]string{}
	// // regionList := map[string][]model.RegionInfo{}
	// json.NewDecoder(body).Decode(&regionList)
	// fmt.Println(regionList)	// map[KeyValueInfoList:[] ProviderName:[] RegionName:[]]
	// // fmt.Println(regionList["connectionconfig"])
	regionInfo := model.RegionInfo{}
	json.NewDecoder(respBody).Decode(&regionInfo)
	fmt.Println(regionInfo)
	fmt.Println(regionInfo.KeyValueInfoList)
	return regionInfo, respStatus
}

// Region 등록
func RegRegion(regionInfo *model.RegionInfo) (io.ReadCloser, int) {
	// buff := bytes.NewBuffer(pbytes)
	url := util.SPIDER + "/region"

	fmt.Println("RegRegion : ", regionInfo)

	// body, err := util.CommonHttpPost(url, regionInfo)
	pbytes, _ := json.Marshal(regionInfo)
	// body, err := util.CommonHttpPost(url, pbytes)
	resp, err := util.CommonHttp(url, pbytes, http.MethodPost)

	if err != nil {
		fmt.Println(err)
	}

	respBody := resp.Body
	respStatus := resp.StatusCode

	return respBody, respStatus
}

// Region 삭제
func DelRegion(regionName string) (io.ReadCloser, int) {
	// buff := bytes.NewBuffer(pbytes)
	url := util.SPIDER + "/region/" + regionName

	fmt.Println("DelRegion : ", regionName)

	// body, err := util.CommonHttpPost(url, regionInfo)

	pbytes, _ := json.Marshal(regionName)
	// body, err := util.CommonHttpDelete(url, pbytes)
	resp, err := util.CommonHttp(url, pbytes, http.MethodDelete)

	if err != nil {
		fmt.Println(err)
	}
	// return body, err
	respBody := resp.Body
	respStatus := resp.StatusCode

	return respBody, respStatus
}

// 현재 설정된 credential 목록 : 목록에서는 key의 value는 ...으로 표시
func GetCredentialList() ([]model.CredentialInfo, int ) {

	// SPIDER == SPIDER
	url := util.SPIDER + "/" + "credential"
	// fmt.Println("=========== GetRegionListData : ", url)

	resp, err := util.CommonHttp(url, nil, http.MethodGet)
	// defer body.Close()

	if err != nil {
		fmt.Println(err)
	}

	respBody := resp.Body
	respStatus := resp.StatusCode

	credentialList := map[string][]model.CredentialInfo{}
	json.NewDecoder(respBody).Decode(&credentialList)
	fmt.Println(credentialList["credential"])
	// TODO : key의 value에 ...표시
	for _, credentialInfo := range credentialList["credential"] {
		fmt.Println("credentialInfo : ", credentialInfo)
		keyValueInfoList := credentialInfo.KeyValueInfoList
		fmt.Println("before keyValueInfoList : ", keyValueInfoList)
		for _, keyValueInfo := range keyValueInfoList {
			keyValueInfo.Value = "..."
		}
		fmt.Println("after keyValueInfoList : ", keyValueInfoList)
	}

	return credentialList["credential"], respStatus
}

// Credential 상세조회
func GetCredentialData(credentialName string) (model.CredentialInfo, int ) {
	url := util.SPIDER + "/credential/" + credentialName
	fmt.Println("=========== GetCredentialData : ", credentialName)

	resp, err := util.CommonHttp(url, nil, http.MethodGet)
	// defer body.Close()

	if err != nil {
		fmt.Println(err)
	}
	respBody := resp.Body
	respStatus := resp.StatusCode

	credentialInfo := model.CredentialInfo{}
	json.NewDecoder(respBody).Decode(&credentialInfo)
	fmt.Println(credentialInfo)
	fmt.Println(credentialInfo.KeyValueInfoList)
	return credentialInfo, respStatus
}

// Credential 등록
func RegCredential(credentialInfo *model.CredentialInfo) (io.ReadCloser, int) {
	// buff := bytes.NewBuffer(pbytes)
	url := util.SPIDER + "/credential"

	fmt.Println("RegCredential : ", credentialInfo)

	// body, err := util.CommonHttpPost(url, regionInfo)
	pbytes, _ := json.Marshal(credentialInfo)
	resp, err := util.CommonHttp(url, pbytes, http.MethodPost)
	if err != nil {
		fmt.Println(err)
	}

	respBody := resp.Body
	respStatus := resp.StatusCode
	return respBody, respStatus
}

// Credential 삭제
func DelCredential(credentialName string) (io.ReadCloser, int) {

	// buff := bytes.NewBuffer(pbytes)
	url := util.SPIDER + "/credential/" + credentialName

	fmt.Println("DelCredential : ", credentialName)

	pbytes, _ := json.Marshal(credentialName)
	// body, err := util.CommonHttpDelete(url, pbytes)
	resp, err := util.CommonHttp(url, pbytes, http.MethodDelete)

	if err != nil {
		fmt.Println(err)
	}
	// return body, err
	respBody := resp.Body
	respStatus := resp.StatusCode
	return respBody, respStatus
}

// 현재 설정된 Driver 목록
func GetDriverList() ([]model.DriverInfo, int ) {
	url := util.SPIDER + "/" + "driver"
	fmt.Println("=========== GetDriverListData : ", url)

	resp, err := util.CommonHttp(url, nil, http.MethodGet)
	// defer body.Close()

	if err != nil {
		fmt.Println(err)
	}

	respBody := resp.Body
	respStatus := resp.StatusCode

	driverList := map[string][]model.DriverInfo{}
	json.NewDecoder(respBody).Decode(&driverList)
	fmt.Println(driverList["driver"])

	return driverList["driver"], respStatus
}

// Driver 상세조회
func GetDriverData(driverlName string) (model.DriverInfo, int ) {
	url := util.SPIDER + "/driver/" + driverlName
	fmt.Println("=========== GetDriverData : ", url)

	resp, err := util.CommonHttp(url, nil, http.MethodGet)
	// defer body.Close()

	if err != nil {
		fmt.Println(err)
	}

	respBody := resp.Body
	respStatus := resp.StatusCode

	driverInfo := model.DriverInfo{}
	json.NewDecoder(respBody).Decode(&driverInfo)
	fmt.Println(driverInfo)
	return driverInfo, respStatus
}

// Driver 등록
func RegDriver(driverInfo *model.DriverInfo) (io.ReadCloser, int) {
	// buff := bytes.NewBuffer(pbytes)
	url := util.SPIDER + "/driver"

	fmt.Println("driverInfo : ", driverInfo)

	// body, err := util.CommonHttpPost(url, regionInfo)
	pbytes, _ := json.Marshal(driverInfo)
	resp, err := util.CommonHttp(url, pbytes, http.MethodPost)
	if err != nil {
		fmt.Println(err)
	}

	respBody := resp.Body
	respStatus := resp.StatusCode
	return respBody, respStatus
}

// Driver 삭제
func DelDriver(driverName string) (io.ReadCloser, int) {

	// buff := bytes.NewBuffer(pbytes)
	url := util.SPIDER + "/driver/" + driverName

	fmt.Println("driverName : ", driverName)

	pbytes, _ := json.Marshal(driverName)
	// body, err := util.CommonHttpDelete(url, pbytes)
	resp, err := util.CommonHttp(url, pbytes, http.MethodDelete)

	if err != nil {
		fmt.Println(err)
	}
	// return body, err
	respBody := resp.Body
	respStatus := resp.StatusCode
	return respBody, respStatus
}

// 해당 namespace의 vpc 목록 조회 -> ResourceHandler로 이동
// func GetVnetList(nameSpaceID string) (io.ReadCloser, error) {
// url := TumbleUrl + "ns/" + nameSpaceID + "/resources/vNet"

// fmt.Println("nameSpaceID : ", nameSpaceID)

// pbytes, _ := json.Marshal(nameSpaceID)
// body, err := util.CommonHttp(url, pbytes, http.MethodGet)

// if err != nil {
// 	fmt.Println(err)
// }
// return body, err
// }

// // vpc 상세 조회-> ResourceHandler로 이동
// func GetVpcData(nameSpaceID string, vNetID string) (io.ReadCloser, error) {
// 	url := TumbleUrl + "ns/" + nameSpaceID + "/resources/vNet"

// 	fmt.Println("nameSpaceID : ", nameSpaceID)

// 	pbytes, _ := json.Marshal(nameSpaceID)
// 	body, err := util.CommonHttp(url, pbytes, http.MethodGet)

// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return body, err
// }

// vpc 등록 -> ResourceHandler로 이동
// func RegVpc(nameSpaceID string, vnetInfo *model.VNetInfo) (io.ReadCloser, error) {
// 	url := TumbleUrl + "ns/" + nameSpaceID + "/resources/vNet"

// 	fmt.Println("nameSpaceID : ", nameSpaceID)

// 	pbytes, _ := json.Marshal(vnetInfo)
// 	body, err := util.CommonHttp(url, pbytes, http.MethodPost)

// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return body, err
// }

// vpc 삭제 -> ResourceHandler로 이동
// func DelVpc(nameSpaceID string, vNetID string) (io.ReadCloser, error) {
// 	url := TumbleUrl + "ns/" + nameSpaceID + "/resources/vNet" + vNetID

// 	fmt.Println("nameSpaceID : ", nameSpaceID)

// 	pbytes, _ := json.Marshal(vNetID)
// 	body, err := util.CommonHttp(url, pbytes, http.MethodDelete)

// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return body, err
// }

// func GetConnectionconfig(drivername string) CloudConnectionInfo {
// 	url := NameSpaceUrl + "/driver/" + drivername

// 	// resp, err := http.Get(url)

// 	// if err != nil {
// 	fmt.Println("request URL : ", url)
// 	// }

// 	// defer resp.Body.Close()
// 	body := HttpGetHandler(url)
// 	nsInfo := CloudConnectionInfo{}

// 	json.NewDecoder(body).Decode(&nsInfo)
// 	fmt.Println("nsInfo : ", nsInfo.ID)
// 	return nsInfo

// }
// func GetImageList() []Image {
// 	url := CloudConnectionUrl + "/connectionconfig"
// 	// resp, err := http.Get(url)
// 	// if err != nil {
// 	// 	fmt.Println("request URL : ", url)
// 	// }

// 	// defer resp.Body.Close()
// 	body := HttpGetHandler(url)
// 	defer body.Close()

// 	nsInfo := ImageRESP{}
// 	json.NewDecoder(body).Decode(&nsInfo)
// 	fmt.Println("nsInfo : ", nsInfo.Image[0].id)
// 	var info []Image
// 	for _, item := range nsInfo.Image {
// 		reg := Image{
// 			id:             item.id,
// 			name:           item.name,
// 			connectionName: item.connectionName,
// 			cspImageId:     item.cspImageId,
// 			cspImageName:   item.cspImageName,
// 			description:    item.description,
// 			guestOS:        item.guestOS,
// 			status:         item.status,
// 		}
// 		info = append(info, reg)
// 	}
// 	return info

// }

// func GetConnectionList() []CloudConnectionInfo {
// 	url := CloudConnectionUrl + "/connectionconfig"
// 	// resp, err := http.Get(url)
// 	// if err != nil {
// 	// 	fmt.Println("request URL : ", url)
// 	// }

// 	// defer resp.Body.Close()
// 	body := HttpGetHandler(url)
// 	defer body.Close()

// 	nsInfo := map[string][]CloudConnectionInfo{}
// 	json.NewDecoder(body).Decode(&nsInfo)
// 	//fmt.Println("nsInfo : ", nsInfo["connectionconfig"][0].ID)
// 	return nsInfo["connectionconfig"]

// }

// func GetDriverReg() []CloudConnectionInfo {
// 	url := NameSpaceUrl + "/driver"

// 	body := HttpGetHandler(url)
// 	defer body.Close()
// 	nsInfo := map[string][]CloudConnectionInfo{}
// 	json.NewDecoder(body).Decode(&nsInfo)

// 	return nsInfo["driver"]

// }

// func GetCredentialList() []CloudConnectionInfo {
// 	url := CloudConnectionUrl + "/credential"
// 	// resp, err := http.Get(url)
// 	// if err != nil {
// 	// 	fmt.Println("request URL : ", url)
// 	// }

// 	// defer resp.Body.Close()
// 	body := HttpGetHandler(url)
// 	defer body.Close()
// 	nsInfo := map[string][]CloudConnectionInfo{}
// 	json.NewDecoder(body).Decode(&nsInfo)
// 	// fmt.Println("nsInfo : ", nsInfo["credential"][0].ID)
// 	return nsInfo["credential"]

// }

// func GetRegionList() []RegionInfo {
// 	url := CloudConnectionUrl + "/region"
// 	fmt.Println("=========== Get Start Region List : ", url)

// 	body := HttpGetHandler(url)
// 	defer body.Close()

// 	nsInfo := RESP{}

// 	json.NewDecoder(body).Decode(&nsInfo)
// 	var info []RegionInfo

// 	for _, item := range nsInfo.Region {

// 		reg := RegionInfo{
// 			RegionName:   item.RegionName,
// 			ProviderName: item.ProviderName,
// 		}
// 		info = append(info, reg)
// 	}
// 	fmt.Println("info region list : ", info)
// 	return info

// }

// func GetCredentialReg() []CloudConnectionInfo {
// 	url := CloudConnectionUrl + "/credential"

// 	body := HttpGetHandler(url)
// 	defer body.Close()

// 	nsInfo := map[string][]CloudConnectionInfo{}
// 	json.NewDecoder(body).Decode(&nsInfo)
// 	fmt.Println("nsInfo : ", nsInfo["credential"][0].ID)
// 	return nsInfo["credential"]

// }

// func GetGeoMetryInfo(wg *sync.WaitGroup, ip_address string, status string, vm_id string, vm_name string, returnResult *[]IPStackInfo) {
// 	defer wg.Done() //goroutin sync done

// 	apiUrl := "http://api.ipstack.com/"
// 	access_key := "86c895286435070c0369a53d2d0b03d1"
// 	url := apiUrl + ip_address + "?access_key=" + access_key
// 	resp, err := http.Get(url)
// 	fmt.Println("GetGeoMetryInfo request URL : ", url)
// 	if err != nil {
// 		fmt.Println("GetGeoMetryInfo request URL : ", url)
// 	}
// 	defer resp.Body.Close()

// 	//그냥 스트링으로 반환해서 프론트에서 JSON.parse로 처리 하는 방법도 괜찮네
// 	//spew.Dump(resp.Body)
// 	// bytes, _ := ioutil.ReadAll(resp.Body)
// 	// str := string(bytes)
// 	// fmt.Println(str)
// 	// *returnStr = append(*returnStr, str)

// 	ipStackInfo := IPStackInfo{
// 		VMID:   vm_id,
// 		Status: status,
// 		VMName: vm_name,
// 	}

// 	json.NewDecoder(resp.Body).Decode(&ipStackInfo)
// 	fmt.Println("Get GeoMetry INFO :", ipStackInfo)

// 	*returnResult = append(*returnResult, ipStackInfo)
// }

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
