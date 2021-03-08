package service

// controller에서 service로 이동
import (
	"encoding/base64"
	"fmt"
	// "net/http"
	"os"
	"strconv"
	"time"

	model "github.com/dogfootman/cb-webtool/src/model"

	// model "github.com/dogfootman/cb-webtool/src/model"
	// service "github.com/dogfootman/cb-webtool/src/service"

	echosession "github.com/go-session/echo-session"
	"github.com/labstack/echo"
)

// type LoginInfo struct {
// 	Username  string
// 	NameSpace string
// }

// type CredentialInfo struct {
// 	Username string
// 	Password string
// }

// model package로 이동
// type CommonURL struct {
// 	SpiderURL    string
// 	TumbleBugURL string
// 	DragonFlyURL string
// 	LadyBugURL   string
// }

func GetCommonURL() model.CommonURL {
	common_url := model.CommonURL{
		SpiderURL:    os.Getenv("SPIDER_URL"),
		TumbleBugURL: os.Getenv("TUMBLE_URL"),
		DragonFlyURL: os.Getenv("DRAGONFLY_URL"),
		LadyBugURL:   os.Getenv("LADYBUG_URL"),
	}
	return common_url
}

func GetCredentialInfo(c echo.Context, username string) model.CredentialInfo {
	store := echosession.FromContext(c)
	getObj, ok := store.Get(username)
	if !ok {
		return model.CredentialInfo{}
	}
	result := getObj.(map[string]string)
	credentialInfo := model.CredentialInfo{
		Username: result["username"],
		Password: result["password"],
	}
	return credentialInfo
}

// func SetLoginInfo(c echo.Context) LoginInfo {
// 	store := echosession.FromContext(c)
// 	nsList := service.GetNSList()
// 	store.Set("username")
// }

// NamespaceController로 이동
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

// NamespaceController로 이동
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

func CallLoginInfo(c echo.Context) model.LoginInfo {
	store := echosession.FromContext(c)
	// getUser, ok := store.Get("username")
	getUser, ok := store.Get("userid")
	if !ok {
		fmt.Println("========= CallLoginInfo Nothing =========")
		return model.LoginInfo{}
	}
	fmt.Println("GETUSER : ", getUser.(string))
	getObj, ok := store.Get(getUser.(string))

	if !ok {
		return model.LoginInfo{}
	}

	result := getObj.(map[string]string)
	loginInfo := model.LoginInfo{
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

// pkg.go.dev 에 이미 있으므로 다른 이름 필요
// func AuthenticationHandler() string {
func AuthenticationKey() string {
	api_username := os.Getenv("API_USERNAME")
	api_password := os.Getenv("API_PASSWORD")

	//The header "KEY: VAL" is "Authorization: Basic {base64 encoded $USERNAME:$PASSWORD}".
	apiUserInfo := api_username + ":" + api_password
	encA := base64.StdEncoding.EncodeToString([]byte(apiUserInfo))
	//req.Header.Add("Authorization", "Basic"+encA)
	return "Basic " + encA

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
