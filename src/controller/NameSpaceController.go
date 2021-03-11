package controller

import (
	"fmt"
	_ "fmt"
	"net/http"
	_ "net/http"

	// echotemplate "github.com/foolin/echo-template"
	echosession "github.com/go-session/echo-session"

	"github.com/cloud-barista/cb-webtool/src/service"
	"github.com/labstack/echo"
)

func NsRegController(c echo.Context) error {
	username := c.FormValue("username")
	description := c.FormValue("description")

	fmt.Println("NSRegController : ", username, description)
	return nil
}

func NsRegForm(c echo.Context) error {
	comURL := service.GetCommonURL()
	apiInfo := service.AuthenticationHandler()
	if loginInfo := service.CallLoginInfo(c); loginInfo.Username != "" {
		return c.Render(http.StatusOK, "NSRegister.html", map[string]interface{}{
			"LoginInfo": loginInfo,
			"comURL":    comURL,
			"apiInfo":   apiInfo,
		})
	}
	return c.Redirect(http.StatusTemporaryRedirect, "/login")
	//return c.Render(http.StatusOK, "NSRegister.html", nil)
}

// namespace 등록 처리
func NsRegProc(c echo.Context) error {
	username := c.FormValue("username")
	description := c.FormValue("description")

	fmt.Println("NSRegController : ", username, description)
	// return nil

	// Tubblebug 호출하여 namespace 생성

	// 성공하면 namespace 목록 조회	
	nsList, err := service.GetNameSpaceList()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
			"status":  "403",
			"nsList":  nil,
		})
	}
	// return namespace 목록
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"status":  "200",
		"nsList":  nsList,
	})
}

func NsListForm(c echo.Context) error {
	fmt.Println("=============start NsListForm =============")
	comURL := service.GetCommonURL()
	apiInfo := service.AuthenticationHandler()
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.Username != "" {
		fmt.Println("=============start GetNSList =============")
		nsList := service.GetNSList()
		fmt.Println("=============start GetNSList =============", nsList)
		if nsList != nil {
			return c.Render(http.StatusOK, "NameSpace.html", map[string]interface{}{
				"LoginInfo": loginInfo,
				"NSList":    nsList,
				"comURL":    comURL,
				"apiInfo":   apiInfo,
			})
		} else {
			return c.Redirect(http.StatusTemporaryRedirect, "/NS/reg")
		}

	}

	fmt.Println("LoginInfo : ", loginInfo)
	//return c.Redirect(http.StatusPermanentRedirect, "/login")
	return c.Redirect(http.StatusTemporaryRedirect, "/login")
}

func GetNameSpaceList(c echo.Context) error {
	fmt.Println("====== GET NAMESPACE LIST ========")
	store := echosession.FromContext(c)
	nsInfoList, nsErr := service.GetNameSpaceList()
	if nsErr != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid tumblebug connection",
			"status":  "403",
		})
	} else {
		store.Set("namespaceList", nsInfoList)
		store.Save()
	}

	return c.JSON(http.StatusOK, nsInfoList)
}

// 기본 namespace set. set default Namespace
func SetNameSpace(c echo.Context) error {
	fmt.Println("====== SET SELECTED NAME SPACE ========")
	store := echosession.FromContext(c)
	ns := c.Param("nsid")
	fmt.Println("SetNameSpaceID : ", ns)
	store.Set("namespace", ns)
	err := store.Save()
	res := map[string]string{
		"message": "success",
	}
	if err != nil {
		res["message"] = "fail"
		return c.JSON(http.StatusNotAcceptable, res)
	}
	return c.JSON(http.StatusOK, res)
}

// 기본 namespace get. get default Namespace
func GetNameSpace(c echo.Context) error {
	fmt.Println("====== GET SELECTED NAME SPACE ========")
	store := echosession.FromContext(c)

	getInfo, ok := store.Get("namespace")
	if !ok {
		return c.JSON(http.StatusNotAcceptable, map[string]string{
			"message": "Not Exist",
		})
	}
	nsId := getInfo.(string)

	res := map[string]string{
		"message": "success",
		"nsID":    nsId,
	}

	return c.JSON(http.StatusOK, res)
}
