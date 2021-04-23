package controller

import (
	"fmt"
	"log"
	"net/http"

	service "github.com/cloud-barista/cb-webtool/src/service"
	"github.com/cloud-barista/cb-webtool/src/util"
	"github.com/labstack/echo"

	echotemplate "github.com/foolin/echo-template"
	echosession "github.com/go-session/echo-session"
)

// MCIS Monitoring 화면
func McisMonitoringMngForm(c echo.Context) error {
	fmt.Println("McisMonitoringMngForm ************ : ")

	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	store := echosession.FromContext(c)

	// 최신 namespacelist 가져오기
	nsList, _ := service.GetNameSpaceList()
	store.Set("namespace", nsList)
	log.Println(" nsList  ", nsList)

	// 해당 Namespace의 모든 MCIS 조회
	mcisList, _ := service.GetMcisList(defaultNameSpaceID)
	log.Println(" mcisList  ", mcisList)

	return echotemplate.Render(c, http.StatusOK,
		"operation/monitorings/McisMonitoringMng", // 파일명
		map[string]interface{}{
			"LoginInfo":          loginInfo,
			"DefaultNameSpaceID": defaultNameSpaceID,
			"NameSpaceList":      nsList,
			"McisList":           mcisList,
		})

}
func MornitoringListForm(c echo.Context) error {
	comURL := service.GetCommonURL()
	apiInfo := util.AuthenticationHandler()
	if loginInfo := service.CallLoginInfo(c); loginInfo.UserID != "" {
		namespace := service.GetNameSpaceToString(c)
		return c.Render(http.StatusOK, "Monitoring_Mcis.html", map[string]interface{}{
			// return c.Render(http.StatusOK, "Monitoring.html", map[string]interface{}{
			"LoginInfo": loginInfo,
			"NameSpace": namespace,
			"comURL":    comURL,
			"apiInfo":   apiInfo,
		})

	}

	//return c.Render(http.StatusOK, "MCISlist.html", nil)
	return c.Redirect(http.StatusTemporaryRedirect, "/login")
}

// func AgentRegForm(c echo.Context) error {
// 	comURL := service.GetCommonURL()
// 	apiInfo := service.AuthenticationHandler()
// 	mcis_id := c.Param("mcis_id")
// 	vm_id := c.Param("vm_id")
// 	public_ip := c.Param("public_ip")

// 	if loginInfo := service.CallLoginInfo(c); loginInfo.Username != "" {
// 		namespace := service.GetNameSpaceToString(c)
// 		return c.Render(http.StatusOK, "InstallAgent.html", map[string]interface{}{
// 			"LoginInfo": loginInfo,
// 			"NameSpace": namespace,
// 			"comURL":    comURL,
// 			"mcis_id":   mcis_id,
// 			"vm_id":     vm_id,
// 			"public_ip": public_ip,
// 			"apiInfo":   apiInfo,
// 		})

// 	}

// 	//return c.Render(http.StatusOK, "MCISlist.html", nil)
// 	return c.Redirect(http.StatusTemporaryRedirect, "/login")
// }
