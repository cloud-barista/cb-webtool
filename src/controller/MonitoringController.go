package controller

import (
	"net/http"

	service "github.com/cloud-barista/cb-webtool/src/service"

	"github.com/labstack/echo"
)

func MornitoringListForm(c echo.Context) error {
	comURL := service.GetCommonURL()
	apiInfo := service.AuthenticationHandler()
	if loginInfo := service.CallLoginInfo(c); loginInfo.Username != "" {
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

func AgentRegForm(c echo.Context) error {
	comURL := service.GetCommonURL()
	apiInfo := service.AuthenticationHandler()
	mcis_id := c.Param("mcis_id")
	vm_id := c.Param("vm_id")
	public_ip := c.Param("public_ip")

	if loginInfo := service.CallLoginInfo(c); loginInfo.Username != "" {
		namespace := service.GetNameSpaceToString(c)
		return c.Render(http.StatusOK, "InstallAgent.html", map[string]interface{}{
			"LoginInfo": loginInfo,
			"NameSpace": namespace,
			"comURL":    comURL,
			"mcis_id":   mcis_id,
			"vm_id":     vm_id,
			"public_ip": public_ip,
			"apiInfo":   apiInfo,
		})

	}

	//return c.Render(http.StatusOK, "MCISlist.html", nil)
	return c.Redirect(http.StatusTemporaryRedirect, "/login")
}
