package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cloud-barista/cb-webtool/src/model/tumblebug"
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

///mcis/:mcisID/vm/:vmID/agent/mngform
// vm에 monitoring Agent 등록 하는 폼.
// TODO : 이거 지금 쓰는데가 없는데???
func VmMonitoringAgentRegForm(c echo.Context) error {
	fmt.Println("VmMonitoringAgentRegForm ************ : ")

	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}
	mcisID := c.Param("mcisID")
	vmID := c.Param("vmID")
	//publicIp := c.Param("public_ip")

	namespace := service.GetNameSpaceToString(c)
	return c.Render(http.StatusOK, "InstallAgent.html", map[string]interface{}{
		// return c.Render(http.StatusOK, "InstallAgent.html", map[string]interface{}{
		"LoginInfo": loginInfo,
		"NameSpace": namespace,
		"mcisID":    mcisID,
		"vmID":      vmID,
		// "publicIp":  publicIp,
	})

}

// VM에 모니터링 Agent 설치
// /ns/{nsId}/monitoring/install/mcis/{mcisId}
func VmMonitoringAgentRegProc(c echo.Context) error {
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	vmMonitoringAgentReg := &tumblebug.VmMonitoringAgentReg{}
	if err := c.Bind(vmMonitoringAgentReg); err != nil {
		// if err := c.Bind(mCISInfoList); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
			"status":  "fail",
		})
	}
	log.Println(vmMonitoringAgentReg)

	// store := echosession.FromContext(c)
	defaultNameSpaceID := loginInfo.DefaultNameSpaceID
	// TODO : defaultNameSpaceID 가 없으면 설정화면으로 보낼 것

	mcisID := c.Param("mcisID")
	vmMonitoringAgentInfo, respStatus := service.RegMonitoringAgentInVm(defaultNameSpaceID, mcisID, vmMonitoringAgentReg)
	if respStatus.StatusCode != 200 && respStatus.StatusCode != 201 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": respStatus.Message,
			"status":  respStatus.StatusCode,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":               "success",
		"status":                respStatus.StatusCode,
		"VmMonitoringAgentInfo": vmMonitoringAgentInfo,
	})
}

// InstallAgent.html
// func AgentRegForm(c echo.Context) error {
// }

// GetMcisInfoData
// 특정 MCIS의 상세정보를 가져온다.
func GetVmMonitoringInfoData(c echo.Context) error {
	log.Println("GetVmMonitoringInfoData")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login") // 조회기능에서 바로 login화면으로 돌리지말고 return message로 하는게 낫지 않을까?
	}

	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	mcisID := c.Param("mcisID")
	metric := c.Param("metric")
	log.Println("mcisID= " + mcisID)
	//monitoringGroup.GET("/mcis/:mcisID/metric/:metric", controller.GetVmMonitoringInfoData)

	resultMcisInfo, _ := service.GetVmMonitoringInfoData(defaultNameSpaceID, mcisID, metric)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success",
		"status":   200,
		"McisInfo": resultMcisInfo,
	})
}
