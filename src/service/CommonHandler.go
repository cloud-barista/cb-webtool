package service

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	echosession "github.com/go-session/echo-session"
	"github.com/labstack/echo"
)

var SpiderURL = os.Getenv("SPIDER_URL")
var TumbleBugURL = os.Getenv("TUMBLE_URL")
var DragonFlyURL = os.Getenv("DRAGONFLY_URL")
var LadyBugURL = os.Getenv("LADYBUG_URL")

type LoginInfo struct {
	Username  string
	NameSpace string
}

type CredentialInfo struct {
	Username string
	Password string
}
type CommonURL struct {
	SpiderURL    string
	TumbleBugURL string
	DragonFlyURL string
	LadyBugURL   string
}

func GetCommonURL() CommonURL {
	common_url := CommonURL{
		SpiderURL:    os.Getenv("SPIDER_URL"),
		TumbleBugURL: os.Getenv("TUMBLE_URL"),
		DragonFlyURL: os.Getenv("DRAGONFLY_URL"),
		LadyBugURL:   os.Getenv("LADYBUG_URL"),
	}
	return common_url
}

func GetCredentialInfo(c echo.Context, username string) CredentialInfo {
	store := echosession.FromContext(c)
	getObj, ok := store.Get(username)
	if !ok {
		return CredentialInfo{}
	}
	result := getObj.(map[string]string)
	credentialInfo := CredentialInfo{
		Username: result["username"],
		Password: result["password"],
	}
	return credentialInfo
}

func HttpGetHandler(url string) io.ReadCloser {
	authInfo := AuthenticationHandler()

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", authInfo)

	client := &http.Client{}
	resp, _ := client.Do(req)

	//defer resp.Body.Close()

	return resp.Body
}

// SPIDER 호출. 결과값이 Object 일 때
func GetSpiderObject(targetUri string) io.ReadCloser {
	url := SpiderURL + "/" + targetUri
	authInfo := AuthenticationHandler()

	// resp, err := http.Get(url)

	// if err != nil {
	fmt.Println("request URL : ", url)
	// }

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", authInfo)

	client := &http.Client{}
	resp, _ := client.Do(req)

	//defer resp.Body.Close()

	return resp.Body
}

func GetSpiderList(targetUri string) (io.ReadCloser, error) {
	url := SpiderURL + targetUri // targetUrl 은 "/" 부터 시작이다. ex) targetUrl = "/ns/reg"
	authInfo := AuthenticationHandler()

	// resp, err := http.Get(url)

	// if err != nil {
	fmt.Println("request URL : ", url)
	// }

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", authInfo)

	client := &http.Client{}
	resp, err := client.Do(req)

	//defer resp.Body.Close()

	return resp.Body, err
}

// func SetLoginInfo(c echo.Context) LoginInfo {
// 	store := echosession.FromContext(c)
// 	nsList := service.GetNSList()
// 	store.Set("username")
// }

// func SetNameSpace(c echo.Context) error {
// 	fmt.Println("====== SET NAME SPACE ========")
// 	store := echosession.FromContext(c)
// 	ns := c.Param("nsid")
// 	fmt.Println("SetNameSpaceID : ", ns)
// 	store.Set("namespace", ns)
// 	err := store.Save()
// 	res := map[string]string{
// 		"message": "success",
// 	}
// 	if err != nil {
// 		res["message"] = "fail"
// 		return c.JSON(http.StatusNotAcceptable, res)
// 	}
// 	return c.JSON(http.StatusOK, res)
// }

// move to NameSpaceController.go
// func GetNameSpace(c echo.Context) error {
// 	fmt.Println("====== GET NAME SPACE ========")
// 	store := echosession.FromContext(c)

// 	getInfo, ok := store.Get("namespace")
// 	if !ok {
// 		return c.JSON(http.StatusNotAcceptable, map[string]string{
// 			"message": "Not Exist",
// 		})
// 	}
// 	nsId := getInfo.(string)

// 	res := map[string]string{
// 		"message": "success",
// 		"nsID":    nsId,
// 	}

// 	return c.JSON(http.StatusOK, res)
// }

func GetNameSpaceToString(c echo.Context) string {
	fmt.Println("====== GET NAME SPACE ========")
	store := echosession.FromContext(c)

	getInfo, ok := store.Get("namespace")
	if !ok {
		return ""
	}
	nsId := getInfo.(string)

	return nsId
}

func CallLoginInfo(c echo.Context) LoginInfo {
	store := echosession.FromContext(c)
	getUser, ok := store.Get("username")
	if !ok {
		fmt.Println("========= CallLoginInfo Nothing =========")
		return LoginInfo{}
	}
	fmt.Println("GETUSER : ", getUser.(string))
	getObj, ok := store.Get(getUser.(string))

	if !ok {
		return LoginInfo{}
	}

	result := getObj.(map[string]string)
	loginInfo := LoginInfo{
		Username: "admin",
		//Username:  result["username"],
		NameSpace: result["namespace"],
	}
	getNs, ok := store.Get("namespace")
	if !ok {
		return loginInfo
	}
	loginInfo.NameSpace = getNs.(string)

	return loginInfo

}

func LoginCheck(c echo.Context) bool {
	store := echosession.FromContext(c)

	inputName := c.FormValue("username")
	inputPass := c.FormValue("password")

	getInfo, ok := store.Get(inputName)
	if !ok {
		return false
	}
	result := getInfo.(map[string]string)
	if result["password"] == inputPass && result["username"] == inputName {
		return true
	}

	return false
}

func MakeNameSpace(name string) string {
	now := time.Now()
	nanos := strconv.FormatInt(now.UnixNano(), 10)

	result := name + "-" + nanos
	fmt.Println("makeNameSpace : ", result)
	return result
}

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
