package controller

import (
	"fmt"
	"strconv"
	"time"

	echosession "github.com/go-session/echo-session"
	"github.com/labstack/echo"
)

type LoginInfo struct {
	Username  string
	NameSpace string
}

func CallLoginInfo(c echo.Context) LoginInfo {
	store := echosession.FromContext(c)
	getUser, ok := store.Get("username")
	if !ok {
		return LoginInfo{}
	}
	getObj, ok := store.Get(getUser.(string))
	if !ok {
		return LoginInfo{}
	}
	result := getObj.(map[string]string)
	loginInfo := LoginInfo{
		Username:  result["username"],
		NameSpace: result["namespace"],
	}

	return loginInfo

}

func MakeNameSpace(name string) string {
	now := time.Now()
	nanos := strconv.FormatInt(now.UnixNano(), 10)

	result := name + "-" + nanos
	fmt.Println("makeNameSpace : ", result)
	return result
}

// func RequestTumBleBug(method string, url string, s ) {
// 	proxyReq, err := http.NewRequest(method, url, nil)
// 	if err != nil {
// 		//log.Fatal(err)
// 	}
// 	client := &http.Client{}
// 	proxyRes, err := client.Do(proxyReq)
// 	if err != nil {
// 		//log.Fatal(err)
// 	}

// 	defer proxyRes.Body.Close()
// 	var cInfo []connectionInfo
// 	e := json.NewDecoder(proxyRes.Body).Decode(&cInfo)
// 	if e != nil {
// 		//http.Error(w, e.Error(), http.StatusBadRequest)
// 		//log.Fatal(e)
// 	}
// 	fmt.Println("bind :", cInfo[0])
// 	spew.Dump(cInfo)
// }

// func requestPost
