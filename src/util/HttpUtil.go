package util

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	// "strconv"
	// "time"
	"bytes"
    // "encoding/json"
    // "io/ioutil"

	// echosession "github.com/go-session/echo-session"
	// "github.com/labstack/echo"
	// "github.com/cloud-barista/cb-webtool/src/model"
)


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


func CommonHttpGet(url string) (io.ReadCloser, error) {
	authInfo := AuthenticationHandler()

	fmt.Println("CommonHttpGet ", url)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", authInfo)

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	return resp.Body, err
}

func CommonHttpPost(url string, json []byte ) (io.ReadCloser, error) {
	authInfo := AuthenticationHandler()
	fmt.Println("CommonHttpPost ", url)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))
    if err != nil {
        panic(err)
    }

    // set the request header Content-Type for json
    req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("Authorization", authInfo)
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }

    fmt.Println(resp.StatusCode)
	defer resp.Body.Close()

	return resp.Body, err
}
