package controller

import (
	// "encoding/json"
	"fmt"
	"log"
	"net/http"

	// model "github.com/cloud-barista/cb-webtool/src/model"
	// "github.com/cloud-barista/cb-webtool/src/model/dragonfly"
	// "github.com/cloud-barista/cb-webtool/src/model/tumblebug"

	"github.com/cloud-barista/cb-webtool/src/model/dragonfly"
	service "github.com/cloud-barista/cb-webtool/src/service"

	// util "github.com/cloud-barista/cb-webtool/src/util"

	echotemplate "github.com/foolin/echo-template"
	// echosession "github.com/go-session/echo-session"
	"github.com/labstack/echo"
	// echosession "github.com/go-session/echo-session"
)

// [MCIS] Auto control policy management (WIP) 참조

// PolicyMonitoring 등록화면
func MonitoringConfigPolicyRegForm(c echo.Context) error {
	fmt.Println("PolicyMonitoringRegForm ************ : ")

	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}
	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	// 최신 namespacelist 가져오기
	nsList, _ := service.GetNameSpaceList()
	log.Println(" nsList  ", nsList)

	//

	return echotemplate.Render(c, http.StatusOK,
		"operation/policies/monitoring/MonitoringPolicyCreate",
		map[string]interface{}{
			"LoginInfo":          loginInfo,
			"DefaultNameSpaceID": defaultNameSpaceID,
			"NameSpaceList":      nsList,
		})
}

// Policy Monitoring 관리 화면
func MonitoringConfigPolicyMngForm(c echo.Context) error {
	fmt.Println("PolicyMonitoringMngForm ************ : ")

	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	// 최신 namespacelist 가져오기
	nsList, _ := service.GetNameSpaceList()
	log.Println(" nsList  ", nsList)

	//MonitoringConfig

	monitoringConfig, _ := service.GetMonitoringConfig()
	// status, filepath, return params
	return echotemplate.Render(c, http.StatusOK,
		"/operation/policies/monitoring/MonitoringConfigPolicyMng", // 파일명
		map[string]interface{}{
			"LoginInfo":          loginInfo,
			"DefaultNameSpaceID": defaultNameSpaceID,
			"NameSpaceList":      nsList,
			"MonitoringConfig":   monitoringConfig,
		})
}

// MonitoringPolicy 목록 조회
func GetMonitoringConfigPolicyList(c echo.Context) error {
	log.Println("GetMonitoringConfigPolicyList : ")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	// store := echosession.FromContext(c)
	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	// monitoringPolicyList, respStatus := service.GetMonitoringPolicyList(defaultNameSpaceID)
	// if respStatus.StatusCode != 200 && respStatus.StatusCode != 201 {
	// 	return c.JSON(respStatus.StatusCode, map[string]interface{}{
	// 		"error":  respStatus.Message,
	// 		"status": respStatus.StatusCode,
	// 	})
	// }

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		// "status":               respStatus.StatusCode,
		"DefaultNameSpaceID": defaultNameSpaceID,
		// "MonitoringPolicyList": monitoringPolicyList,
	})
}

// MonitoringPolicy 등록 처리
func MonitoringConfigPolicyPutProc(c echo.Context) error {
	log.Println("MonitoringConfigPolicyRegProc : ")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	monitoringConfigRegInfo := &dragonfly.MonitoringConfigReg{}
	if err := c.Bind(monitoringConfigRegInfo); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
			"status":  "fail",
		})
	}
	log.Println(monitoringConfigRegInfo)

	resultMonitoringConfigInfo, respStatus := service.PutMonigoringConfig(monitoringConfigRegInfo)
	log.Println("MonitoringPolicyReg service returned")
	if respStatus.StatusCode != 200 && respStatus.StatusCode != 201 {
		return c.JSON(respStatus.StatusCode, map[string]interface{}{
			"error":  respStatus.Message,
			"status": respStatus.StatusCode,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":          "success",
		"status":           respStatus.StatusCode,
		"MonitoringConfig": resultMonitoringConfigInfo,
	})
}

//PolicyThresholdMngForm
// PolicyThreshold 등록화면
func ThresholdPolicyRegForm(c echo.Context) error {
	fmt.Println("PolicyThresholdRegForm ************ : ")

	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}
	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	// 최신 namespacelist 가져오기
	nsList, _ := service.GetNameSpaceList()
	log.Println(" nsList  ", nsList)

	//

	return echotemplate.Render(c, http.StatusOK,
		"operation/policies/threshold/ThresholdPolicyCreate",
		map[string]interface{}{
			"LoginInfo":          loginInfo,
			"DefaultNameSpaceID": defaultNameSpaceID,
			"NameSpaceList":      nsList,
		})
}

// Policy Threshold 관리 화면
func MonitoringAlertPolicyMngForm(c echo.Context) error {
	fmt.Println("PolicyThresholdMngForm ************ : ")

	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	// 최신 namespacelist 가져오기
	nsList, _ := service.GetNameSpaceList()
	log.Println(" nsList  ", nsList)

	// status, filepath, return params
	return echotemplate.Render(c, http.StatusOK,
		"/operation/policies/threshold/MonitoringAlertPolicyMng", // 파일명
		map[string]interface{}{
			"LoginInfo":          loginInfo,
			"DefaultNameSpaceID": defaultNameSpaceID,
			"NameSpaceList":      nsList,
		})
}

// Monitoring Threshold 목록 조회
func GetThresholdPolicyList(c echo.Context) error {
	log.Println("GetMonitoringThresholdList : ")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	// store := echosession.FromContext(c)
	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	// monitoringPolicyList, respStatus := service.GetMonitoringPolicyList(defaultNameSpaceID)
	// if respStatus.StatusCode != 200 && respStatus.StatusCode != 201 {
	// 	return c.JSON(respStatus.StatusCode, map[string]interface{}{
	// 		"error":  respStatus.Message,
	// 		"status": respStatus.StatusCode,
	// 	})
	// }

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		// "status":               respStatus.StatusCode,
		"DefaultNameSpaceID": defaultNameSpaceID,
		// "MonitoringPolicyList": monitoringPolicyList,
	})
}

// Threshold 등록 처리
func ThresholdPolicyRegProc(c echo.Context) error {
	log.Println("MonitoringThresholdRegProc : ")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	// defaultNameSpaceID := loginInfo.DefaultNameSpaceID
	// _, respStatus := service.RegMonitoringPolicy(defaultNameSpaceID, mCISInfo)
	// log.Println("MonitoringPolicyReg service returned")
	// if respStatus.StatusCode != 200 && respStatus.StatusCode != 201 {
	// 	return c.JSON(respStatus.StatusCode, map[string]interface{}{
	// 		"error":  respStatus.Message,
	// 		"status": respStatus.StatusCode,
	// 	})
	// }

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		// "status":  respStatus.StatusCode,
	})
}

//
// PolicyPlacement 등록화면
func PlacementPolicyRegForm(c echo.Context) error {
	fmt.Println("PolicyPlacementRegForm ************ : ")

	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}
	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	// 최신 namespacelist 가져오기
	nsList, _ := service.GetNameSpaceList()
	log.Println(" nsList  ", nsList)

	//

	return echotemplate.Render(c, http.StatusOK,
		"operation/policies/placement/PlacementPolicyCreate",
		map[string]interface{}{
			"LoginInfo":          loginInfo,
			"DefaultNameSpaceID": defaultNameSpaceID,
			"NameSpaceList":      nsList,
		})
}

// Policy Monitoring 관리 화면
func PlacementPolicyMngForm(c echo.Context) error {
	fmt.Println("PolicyPlacementMngForm ************ : ")

	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	// 최신 namespacelist 가져오기
	nsList, _ := service.GetNameSpaceList()
	log.Println(" nsList  ", nsList)

	// status, filepath, return params
	return echotemplate.Render(c, http.StatusOK,
		"/operation/policies/placement/PlacementPolicyMng", // 파일명
		map[string]interface{}{
			"LoginInfo":          loginInfo,
			"DefaultNameSpaceID": defaultNameSpaceID,
			"NameSpaceList":      nsList,
		})
}

// Placement Policy 목록 조회
func GetPlacementPolicyList(c echo.Context) error {
	log.Println("GetMonitoringPolicyList : ")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	// store := echosession.FromContext(c)
	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	// monitoringPolicyList, respStatus := service.GetMonitoringPolicyList(defaultNameSpaceID)
	// if respStatus.StatusCode != 200 && respStatus.StatusCode != 201 {
	// 	return c.JSON(respStatus.StatusCode, map[string]interface{}{
	// 		"error":  respStatus.Message,
	// 		"status": respStatus.StatusCode,
	// 	})
	// }

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		// "status":               respStatus.StatusCode,
		"DefaultNameSpaceID": defaultNameSpaceID,
		// "MonitoringPolicyList": monitoringPolicyList,
	})
}

// Placement 등록 처리
func PlacementPolicyRegProc(c echo.Context) error {
	log.Println("PlacementolicyRegProc : ")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	// defaultNameSpaceID := loginInfo.DefaultNameSpaceID
	// _, respStatus := service.RegMonitoringPolicy(defaultNameSpaceID, mCISInfo)
	// log.Println("MonitoringPolicyReg service returned")
	// if respStatus.StatusCode != 200 && respStatus.StatusCode != 201 {
	// 	return c.JSON(respStatus.StatusCode, map[string]interface{}{
	// 		"error":  respStatus.Message,
	// 		"status": respStatus.StatusCode,
	// 	})
	// }

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		// "status":  respStatus.StatusCode,
	})
}
