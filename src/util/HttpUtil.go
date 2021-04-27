package util

import (
	"encoding/base64"
	"fmt"
	// "reflect"
	// "io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	// "time"
	"bytes"
	"encoding/json"
	"math"
	// "io/ioutil"
	// echosession "github.com/go-session/echo-session"
	// "github.com/labstack/echo"
	// "github.com/cloud-barista/cb-webtool/src/model"
)

type KeepZero float64

func (f KeepZero) MarshalJSON() ([]byte, error) {
	if float64(f) == float64(int(f)) {
		return []byte(strconv.FormatFloat(float64(f), 'f', 1, 32)), nil
	}
	return []byte(strconv.FormatFloat(float64(f), 'f', -1, 32)), nil
}

type myFloat64 float64

func (mf myFloat64) MarshalJSON() ([]byte, error) {
	const ε = 1e-12
	v := float64(mf)
	w, f := math.Modf(v)
	if f < ε {
		return []byte(fmt.Sprintf(`%v.0`, math.Trunc(w))), nil
	}
	return json.Marshal(v)
}

// ajax 호출할 때 header key 생성
func AuthenticationHandler() string {

	// conf 파일에 정의
	api_username := os.Getenv("API_USERNAME")
	api_password := os.Getenv("API_PASSWORD")
	// api_username := "default"
	// api_password := "default"

	//The header "KEY: VAL" is "Authorization: Basic {base64 encoded $USERNAME:$PASSWORD}".
	apiUserInfo := api_username + ":" + api_password
	encA := base64.StdEncoding.EncodeToString([]byte(apiUserInfo))
	//req.Header.Add("Authorization", "Basic"+encA)
	return "Basic " + encA

}

// originalUrl 은 API의 전체 경로
// parammapper 의 Key는 replace할 모든 text
// ex1) path인 경우 {abc}
// ex2) path인 경우 :abc
func MappingUrlParameter(originalUrl string, paramMapper map[string]string) string {
	returnUrl := originalUrl
	if paramMapper != nil {
		for key, replaceValue := range paramMapper {
			returnUrl := strings.Replace(originalUrl, key, replaceValue, -1)
			fmt.Println("Key:", key, "=>", "Element:", replaceValue+":"+returnUrl)
		}
	}
	return returnUrl
}

// http 호출
func CommonHttp(url string, json []byte, httpMethod string) (*http.Response, error) {

	authInfo := AuthenticationHandler()

	log.Println("CommonHttp "+httpMethod+", ", url)
	log.Println("authInfo ", authInfo)
	client := &http.Client{}
	req, err1 := http.NewRequest(httpMethod, url, bytes.NewBuffer(json))
	if err1 != nil {
		panic(err1)
	}

	// url = "http://54.248.3.145:1323/tumblebug/ns/ns-01/resources/vNet"

	// set the request header Content-Type for json
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	// req.Header.Set("Content-Type", "application/json")

	req.Header.Add("Authorization", authInfo)
	resp, err := client.Do(req) // err 자체는 nil 이고 resp 내에 statusCode가 500임...

	return resp, err
}

// func CommonHttpWithoutParam1(url string, httpMethod string) (io.ReadCloser, error) {
// 	authInfo := AuthenticationHandler()

// 	fmt.Println("CommonHttp ", url)
// 	client := &http.Client{}
// 	req, err := http.NewRequest(httpMethod, url, nil)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// set the request header Content-Type for json
// 	req.Header.Set("Content-Type", "application/json; charset=utf-8")
// 	req.Header.Add("Authorization", authInfo)
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(resp.StatusCode)
// 	defer resp.Body.Close()

// 	return resp.Body, err
// }

// parameter 없이 호출하는 경우 사용.받은대로 return하면 호출하는 method에서 가공하여 사용
// func CommonHttpWithoutParam(url string, httpMethod string) (io.ReadCloser, error) {
// 	authInfo := AuthenticationHandler()

// 	fmt.Println("CommonHttp ", url)
// 	client := &http.Client{}
// 	req, err := http.NewRequest(httpMethod, url, nil)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// set the request header Content-Type for json
// 	// req.Header.Set("Content-Type", "application/json; charset=utf-8")	// 사용에 주의할 것.
// 	req.Header.Add("Authorization", authInfo)
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		panic(err)
// 	}
// 	// respBody := resp.Body
// 	// robots, _ := ioutil.ReadAll(resp.Body)
// 	// defer resp.Body.Close()
// 	// log.Println(fmt.Print(string(robots)))
// 	// fmt.Println(resp.StatusCode)

// 	return resp.Body, err
// }

// parameter 없이 호출하는 경우 사용.받은대로 return하면 호출하는 method에서 가공하여 사용
func CommonHttpWithoutParam(url string, httpMethod string) (*http.Response, error) {
	authInfo := AuthenticationHandler()

	log.Println("CommonHttpWithoutParam "+httpMethod+", ", url)
	log.Println("authInfo ", authInfo)
	client := &http.Client{}
	req, err := http.NewRequest(httpMethod, url, nil)
	if err != nil {
		fmt.Println("CommonHttpWithoutParam error")
		fmt.Println(err)
		panic(err)
	}

	// set the request header Content-Type for json
	// req.Header.Set("Content-Type", "application/json; charset=utf-8")	// 사용에 주의할 것.
	req.Header.Add("Authorization", authInfo)
	// resp, err := client.Do(req)
	return client.Do(req)
}

// return message 확인용
func DisplayResponse(resp *http.Response) {
	fmt.Println("*****DisplayResponse begin****")
	if resp == nil {
		log.Println(" response is nil ")
	} else {
		// resultMessage, err1 := ioutil.ReadAll(resp.Message)
		// if err1 != nil {
		// 	str := string(resultMessage)
		// 	println("nil ", str)
		// 	println("err1 ", err1)
		// }
		// fmt.Println(string(resultMessage))
		// log.Println(" 11111111111111111111111111111 ")

		fmt.Println(resp.StatusCode)
		log.Println(" 22222222222222222222222222 ")

		fmt.Println(string(resp.Status))
		log.Println(" 3333333333333333333 ")
		// data, err := ioutil.ReadAll(resp.Body)
		// if err != nil {
		//     panic(err)
		// }
		// fmt.Printf("%s\n", string(data))

		// resultStatus := resp.StatusCode
		// fmt.Println("resultStatus ", resultStatus)
		// // fmt.Println("body ",  resp.Body)
		resultBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			str := string(resultBody)
			println("nil ", str)
			println("err ", err)
		}
		fmt.Println(string(resultBody))

		var target interface{}
		body, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body, &target)
		fmt.Println(fmt.Println(target))
		// // json.NewDecoder(respBody).Decode(&stringMap)
		// pbytes, _ := json.Marshal(resultBody)
		// fmt.Println(string(pbytes))

		fmt.Println("*****DisplayResponse end****")
	}
}

// Response 객체의 내용
// type Response struct {
//     Status     string // e.g. "200 OK"
//     StatusCode int    // e.g. 200
//     Proto      string // e.g. "HTTP/1.0"
//     ProtoMajor int    // e.g. 1
//     ProtoMinor int    // e.g. 0

//     // response headers
//     Header http.Header
//     // response body
//     Body io.ReadCloser
//     // request that was sent to obtain the response
//     Request *http.Request
// }
