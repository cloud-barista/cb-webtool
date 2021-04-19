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

// func CommonHttpGet(url string) (io.ReadCloser, error) {
// 	authInfo := AuthenticationHandler()

// 	fmt.Println("CommonHttpGet ", url)
// 	req, _ := http.NewRequest("GET", url, nil)
// 	req.Header.Add("Authorization", authInfo)

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	// defer resp.Body.Close()

// 	return resp.Body, err
// }

// func CommonHttpPost(url string, json []byte) (io.ReadCloser, error) {
// 	authInfo := AuthenticationHandler()

// 	fmt.Println("CommonHttpPost ", url)
// 	client := &http.Client{}
// 	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))
// 	if err != nil {
// 		panic(err)
// 	}

// 	// set the request header Content-Type for json
// 	// req.Header.Set("Content-Type", "application/json; charset=utf-8")
// 	req.Header.Add("Authorization", authInfo)
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(resp.StatusCode)
// 	defer resp.Body.Close()

// 	return resp.Body, err
// }

// 호출 전 json.Marshal로 byte형태로 바꾸어 호출. json []byte로 받으면 공통으로 사용가능하므로
// func CommonHttpDelete(url string, json []byte) (io.ReadCloser, error) {
// 	authInfo := AuthenticationHandler()

// 	fmt.Println("CommonHttpDelete ", url)
// 	client := &http.Client{}
// 	req, err := http.NewRequest(http.MethodDelete, url, bytes.NewBuffer(json))
// 	if err != nil {
// 		panic(err)
// 	}

// 	// set the request header Content-Type for json
// 	// req.Header.Set("Content-Type", "application/json; charset=utf-8")
// 	req.Header.Add("Authorization", authInfo)
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(resp.StatusCode)
// 	defer resp.Body.Close()

// 	return resp.Body, err
// }

// http 호출
// func CommonHttp(url string, json []byte, httpMethod string) (io.ReadCloser, int) {
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
	// if err != nil {
	// 	log.Println("*********************Error*****************************")
	// 	// v := reflect.ValueOf(err)
	// 	// log.Println("**********************1************************")
	// 	// log.Println(v)
	// 	// for i := 0; i < v.NumField(); i++ {
	// 	// 	log.Println("**********************222************************")
	// 	// 	log.Println(v.Type().Field(i).Name)
	// 	// 	log.Println("\t", v.Field(i))
	// 	// }

	// log.Println(" after client.do ")
	// log.Println(err)
	// log.Println(" after client.do22 ")
	// log.Println(resp)

	// 	// log.Println("http %d - %s ", resp.StatusCode, err)
	// 	log.Println("********************Error End******************************")
	// 	// log.Println(err)
	// 	// // log.Println(errorRespBody)
	// 	// log.Println("resp.StatusCode ", resp.StatusCode)
	// 	// log.Println("resp.Body ", resp.Body)
	// 	//DisplayResponse(resp) // return message 확인 용

	// 	// panic(err)
	// }
	// DisplayResponse(resp) // return message 확인 용

	// fmt.Println(resp.StatusCode)
	// //defer resp.Body.Close()
	// // fmt.Println("resp.Body ", resp.Body)
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

	fmt.Println("CommonHttp ", url)
	client := &http.Client{}
	fmt.Println("111")
	req, err := http.NewRequest(httpMethod, url, nil)
	fmt.Println("222")
	if err != nil {
		fmt.Println("456")
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("333")

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
