package controller

import (
	"fmt"
	_ "fmt"
	"net/http"
	_ "net/http"

	// "github.com/cloud-barista/cb-webtool/src/service"
	service "github.com/dogfootman/cb-webtool/src/service"
	
	echosession "github.com/go-session/echo-session"
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
	apiInfo := service.AuthenticationKey()
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

func NsListForm(c echo.Context) error {
	fmt.Println("=============start NsListForm =============")
	comURL := service.GetCommonURL()
	apiInfo := service.AuthenticationKey()
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

func SetNameSpace(c echo.Context) error {
	fmt.Println("====== SET NAME SPACE ========")
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

func GetNameSpace(c echo.Context) error {
	fmt.Println("====== GET NAME SPACE ========")
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

