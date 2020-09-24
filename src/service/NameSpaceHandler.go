package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"reflect"

	"github.com/cloud-barista/cb-webtool/src/controller"
)

// var NameSpaceUrl = "http://15.165.16.67:1323"
var NameSpaceUrl = os.Getenv("TUMBLE_URL")

type NSInfo struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func GetNS(nsID string) NSInfo {
	url := NameSpaceUrl + "ns" + nsID

	body := httpGetHandler(url)
	nsInfo := NSInfo{}
	json.NewDecoder(body).Decode(&nsInfo)
	fmt.Println("nsInfo : ", nsInfo.ID)
	return nsInfo

}

func GetNSList() []NSInfo {
	url := NameSpaceUrl + "/ns"
	fmt.Println("============= NameSpace URL =============", url)
	// authInfo := controller.AuthenticationHandler()
	// req, err := http.NewRequest("GET", url, nil)
	// if err != nil {

	// }
	// req.Header.Add("Authorization", authInfo)
	// client := &http.Client{}
	// resp, err := client.Do(req)
	// fmt.Println("=============result GetNSList =============", resp)
	// //spew.Dump(resp)
	// if err != nil {
	// 	fmt.Println("========= GetNSList Error : ", err)
	// 	fmt.Println("request URL : ", url)
	// 	return nil
	// }

	// defer resp.Body.Close()
	body := httpGetHandler(url)
	nsInfo := map[string][]NSInfo{}

	json.NewDecoder(body).Decode(&nsInfo)
	//fmt.Println("nsInfo : ", nsInfo["ns"][0].ID)
	return nsInfo["ns"]

}

// func RegNS() error {

// }

func RequestGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("request URL : ", url)
	}

	defer resp.Body.Close()
	nsInfo := map[string][]NSInfo{}
	fmt.Println("nsInfo type : ", reflect.TypeOf(nsInfo))
	json.NewDecoder(resp.Body).Decode(&nsInfo)
	fmt.Println("nsInfo : ", nsInfo["ns"][0].ID)

	// data, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println("Get Data Error")
	// }
	// fmt.Println("GetData : ", string(data))

}

func httpGetHandler(url string) io.ReadCloser {
	authInfo := controller.AuthenticationHandler()
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", authInfo)

	resp, _ := client.Do(req)
	fmt.Println("=============result httpGetHandler =============", resp)
	defer resp.Body.Close()

	return resp.Body
}
