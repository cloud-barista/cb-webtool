package service

import (
	"encoding/json"
	"fmt"
	"io"
	// "math"
	"net/http"
	"os"
	// "strconv"
	// "sync"
	//"io/ioutil"
	//"github.com/davecgh/go-spew/spew"
	model "github.com/cloud-barista/cb-webtool/src/model"
	util "github.com/cloud-barista/cb-webtool/src/util"
)

//var CloudConnectionUrl = "http://15.165.16.67:1024"
var CloudConnectionUrl = os.Getenv("SPIDER_URL")
var TumbleUrl = os.Getenv("TUMBLE_URL")

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
func GetCloudOSListData() []string {

	// CloudConnectionUrl == SPIDER
	url := CloudConnectionUrl + "/" + "cloudos"
	// fmt.Println("=========== GetConnectionConfigData : ", url)

	body, err := util.CommonHttpGet(url)
	defer body.Close()

	if err != nil {
		fmt.Println(err)
	}

	cloudOs := map[string][]string{}
	json.NewDecoder(body).Decode(&cloudOs)
	fmt.Println(cloudOs["cloudos"])

	return cloudOs["cloudos"]
}

// 현재 설정된 connection 목록
func GetConnectionConfigListData() []model.ConnectionConfigData {

	// CloudConnectionUrl == SPIDER
	url := CloudConnectionUrl + "/" + "connectionconfig"
	// fmt.Println("=========== GetConnectionConfigData : ", url)

	body, err := util.CommonHttpGet(url)
	defer body.Close()

	if err != nil {
		fmt.Println(err)
	}

	connectionConfigDataList := map[string][]model.ConnectionConfigData{}
	json.NewDecoder(body).Decode(&connectionConfigDataList)
	fmt.Println(connectionConfigDataList["connectionconfig"])

	return connectionConfigDataList["connectionconfig"]
}

// 현재 설정된 region 목록
func GetRegionListData() []model.RegionInfo {

	// CloudConnectionUrl == SPIDER
	url := CloudConnectionUrl + "/" + "region"
	// fmt.Println("=========== GetRegionListData : ", url)

	body, err := util.CommonHttpGet(url)
	defer body.Close()

	if err != nil {
		fmt.Println(err)
	}

	regionList := map[string][]model.RegionInfo{}
	json.NewDecoder(body).Decode(&regionList)
	fmt.Println(regionList["region"])

	return regionList["region"]
}

func GetRegionData(regionName string) model.RegionInfo {
	url := CloudConnectionUrl + "/region/" + regionName
	fmt.Println("=========== GetRegionData : ", regionName)

	body, err := util.CommonHttpGet(url)
	defer body.Close()

	if err != nil {
		fmt.Println(err)
	}

	// region := map[string][]string{}
	// // retion := map[string][]model.ConnectionConfigData{}
	// json.NewDecoder(body).Decode(&region)
	// fmt.Println(region)	// map[KeyValueInfoList:[] ProviderName:[] RegionName:[]]
	// // fmt.Println(connectionConfigDataList["connectionconfig"])
	regionInfo := model.RegionInfo{}
	json.NewDecoder(body).Decode(&regionInfo)
	fmt.Println(regionInfo)
	fmt.Println(regionInfo.KeyValueInfoList)
	return regionInfo
}

// Region 등록
func RegRegion(regionInfo *model.RegionInfo) (io.ReadCloser, error) {
	// buff := bytes.NewBuffer(pbytes)
	url := CloudConnectionUrl + "/region"

	fmt.Println("RegRegion : ", regionInfo)

	// body, err := util.CommonHttpPost(url, regionInfo)
	pbytes, _ := json.Marshal(regionInfo)
	body, err := util.CommonHttpPost(url, pbytes)
	if err != nil {
		fmt.Println(err)
	}
	return body, err
}

// Region 삭제
func DelRegion(regionName string) (io.ReadCloser, error) {
	// buff := bytes.NewBuffer(pbytes)
	url := CloudConnectionUrl + "/region/" + regionName

	fmt.Println("DelRegion : ", regionName)

	// body, err := util.CommonHttpPost(url, regionInfo)

	pbytes, _ := json.Marshal(regionName)
	// body, err := util.CommonHttpDelete(url, pbytes)
	body, err := util.CommonHttp(url, pbytes, http.MethodDelete)

	if err != nil {
		fmt.Println(err)
	}
	return body, err
}

// 현재 설정된 credential 목록
func GetCredentialListData() []model.CredentialInfo {

	// CloudConnectionUrl == SPIDER
	url := CloudConnectionUrl + "/" + "credential"
	// fmt.Println("=========== GetRegionListData : ", url)

	body, err := util.CommonHttpGet(url)
	defer body.Close()

	if err != nil {
		fmt.Println(err)
	}

	credentialList := map[string][]model.CredentialInfo{}
	json.NewDecoder(body).Decode(&credentialList)
	fmt.Println(credentialList["credential"])

	return credentialList["credential"]
}

// Credential 상세조회
func GetCredentialData(credentialName string) model.CredentialInfo {
	url := CloudConnectionUrl + "/region/" + credentialName
	fmt.Println("=========== GetCredentialData : ", credentialName)

	body, err := util.CommonHttpGet(url)
	defer body.Close()

	if err != nil {
		fmt.Println(err)
	}

	credentialInfo := model.CredentialInfo{}
	json.NewDecoder(body).Decode(&credentialInfo)
	fmt.Println(credentialInfo)
	fmt.Println(credentialInfo.KeyValueInfoList)
	return credentialInfo
}

// Credential 등록
func RegCredential(credentialInfo *model.CredentialInfo) (io.ReadCloser, error) {
	// buff := bytes.NewBuffer(pbytes)
	url := CloudConnectionUrl + "/credential"

	fmt.Println("RegCredential : ", credentialInfo)

	// body, err := util.CommonHttpPost(url, regionInfo)
	pbytes, _ := json.Marshal(credentialInfo)
	body, err := util.CommonHttpPost(url, pbytes)
	if err != nil {
		fmt.Println(err)
	}
	return body, err
}

// Credential 삭제
func DelCredential(credentialName string) (io.ReadCloser, error) {

	// buff := bytes.NewBuffer(pbytes)
	url := CloudConnectionUrl + "/credential/" + credentialName

	fmt.Println("DelCredential : ", credentialName)

	pbytes, _ := json.Marshal(credentialName)
	// body, err := util.CommonHttpDelete(url, pbytes)
	body, err := util.CommonHttp(url, pbytes, http.MethodDelete)

	if err != nil {
		fmt.Println(err)
	}
	return body, err
}

// 현재 설정된 Driver 목록
func GetDriverListData() []model.DriverInfo {
	url := CloudConnectionUrl + "/" + "driver"
	fmt.Println("=========== GetDriverListData : ", url)

	body, err := util.CommonHttpGet(url)
	defer body.Close()

	if err != nil {
		fmt.Println(err)
	}

	driverList := map[string][]model.DriverInfo{}
	json.NewDecoder(body).Decode(&driverList)
	fmt.Println(driverList["driver"])

	return driverList["driver"]
}

// Driver 상세조회
func GetDriverData(driverlName string) model.DriverInfo {
	url := CloudConnectionUrl + "/driver/" + driverlName
	fmt.Println("=========== GetDriverData : ", url)

	body, err := util.CommonHttpGet(url)
	defer body.Close()

	if err != nil {
		fmt.Println(err)
	}

	driverInfo := model.DriverInfo{}
	json.NewDecoder(body).Decode(&driverInfo)
	fmt.Println(driverInfo)
	return driverInfo
}

// Driver 등록
func RegDriver(driverInfo *model.DriverInfo) (io.ReadCloser, error) {
	// buff := bytes.NewBuffer(pbytes)
	url := CloudConnectionUrl + "/driver"

	fmt.Println("driverInfo : ", driverInfo)

	// body, err := util.CommonHttpPost(url, regionInfo)
	pbytes, _ := json.Marshal(driverInfo)
	body, err := util.CommonHttpPost(url, pbytes)
	if err != nil {
		fmt.Println(err)
	}
	return body, err
}

// Driver 삭제
func DelDriver(driverName string) (io.ReadCloser, error) {

	// buff := bytes.NewBuffer(pbytes)
	url := CloudConnectionUrl + "/driver/" + driverName

	fmt.Println("driverName : ", driverName)

	pbytes, _ := json.Marshal(driverName)
	// body, err := util.CommonHttpDelete(url, pbytes)
	body, err := util.CommonHttp(url, pbytes, http.MethodDelete)

	if err != nil {
		fmt.Println(err)
	}
	return body, err
}

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
