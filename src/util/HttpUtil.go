package util

import (
	"encoding/base64"
	"fmt"
	// "io"
	// "io/ioutil"
	// "log"
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
func CommonHttp(url string, json []byte, httpMethod string) (*http.Response, error) {
	authInfo := AuthenticationHandler()

	fmt.Println("CommonHttp ", url)
	fmt.Println("authInfo ", authInfo)
	client := &http.Client{}
	req, err := http.NewRequest(httpMethod, url, bytes.NewBuffer(json))
	if err != nil {
		panic(err)
	}

	// url = "http://54.248.3.145:1323/tumblebug/ns/ns-01/resources/vNet"

	// set the request header Content-Type for json
	//req.Header.Set("Content-Type", "application/json; charset=utf-8")
	// req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", authInfo)
	// resp, err := client.Do(req)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(resp.StatusCode)
	// //defer resp.Body.Close()
	// // fmt.Println("resp.Body ", resp.Body)
	// // return resp.Body, err
	return client.Do(req)
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
	req, err := http.NewRequest(httpMethod, url, nil)
	if err != nil {
		panic(err)
	}

	// set the request header Content-Type for json
	// req.Header.Set("Content-Type", "application/json; charset=utf-8")	// 사용에 주의할 것.
	req.Header.Add("Authorization", authInfo)
	// resp, err := client.Do(req)
	return client.Do(req)
}
