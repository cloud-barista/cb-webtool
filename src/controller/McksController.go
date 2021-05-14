package controller

import (
	// "encoding/json"
	"fmt"
	"log"
	"net/http"

	// model "github.com/cloud-barista/cb-webtool/src/model"
	// "github.com/cloud-barista/cb-webtool/src/model/dragonfly"
	// "github.com/cloud-barista/cb-webtool/src/model/tumblebug"

	service "github.com/cloud-barista/cb-webtool/src/service"
	// util "github.com/cloud-barista/cb-webtool/src/util"

	echotemplate "github.com/foolin/echo-template"
	// echosession "github.com/go-session/echo-session"
	"github.com/labstack/echo"
	// echosession "github.com/go-session/echo-session"
)

func McksRegForm(c echo.Context) error {
	fmt.Println("McksRegForm ************ : ")

	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}
	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	// 최신 namespacelist 가져오기
	nsList, _ := service.GetNameSpaceList()
	log.Println(" nsList  ", nsList)

	// connectionconfigList 가져오기
	cloudOsList, _ := service.GetCloudOSList()
	log.Println(" cloudOsList  ", cloudOsList)

	// regionList 가져오기
	regionList, _ := service.GetRegionList()
	log.Println(" regionList  ", regionList)

	cloudConnectionConfigInfoList, _ := service.GetCloudConnectionConfigList() // 등록된 모든 connection 정보
	log.Println("---------------------- GetCloudConnectionConfigList ", defaultNameSpaceID)

	clusterList, _ := service.GetClusterList(defaultNameSpaceID)

	return echotemplate.Render(c, http.StatusOK,
		"operation/manages/mcksmng/McksCreate", // 파일명
		map[string]interface{}{
			"LoginInfo":          loginInfo,
			"DefaultNameSpaceID": defaultNameSpaceID,
			"NameSpaceList":      nsList,
			"CloudOSList":        cloudOsList,
			"RegionList":         regionList,

			"CloudConnectionConfigInfoList": cloudConnectionConfigInfoList,
			"ClusterList":                   clusterList,
		})
}

func McksMngForm(c echo.Context) error {
	fmt.Println("McksMngForm ************ : ")

	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	// 최신 namespacelist 가져오기
	nsList, _ := service.GetNameSpaceList()
	log.Println(" nsList  ", nsList)

	clusterList, _ := service.GetClusterList(defaultNameSpaceID)

	totalClusterCount := len(clusterList)
	if totalClusterCount == 0 {
		return c.Redirect(http.StatusTemporaryRedirect, "/operation/manages/mcksmng/regform")
	}
	// provider 별 연결정보 count(MCIS 무관)
	// cloudConnectionConfigInfoList, _ := service.GetCloudConnectionConfigList()
	// connectionConfigCountMap, providerCount := service.GetCloudConnectionCountMap(cloudConnectionConfigInfoList)
	// totalConnectionCount := len(cloudConnectionConfigInfoList)

	// 모든 MCKS 조회
	// mcksList, _ := service.GetMcksList(defaultNameSpaceID)
	// log.Println(" mcisList  ", mcisList)

	// store := echosession.FromContext(c)
	// store.Set("MCKS_"+loginInfo.UserID, mcksList)

	// status, filepath, return params
	return echotemplate.Render(c, http.StatusOK,
		"operation/manages/mcksmng/McksMng", // 파일명
		map[string]interface{}{
			"LoginInfo":          loginInfo,
			"DefaultNameSpaceID": defaultNameSpaceID,
			"NameSpaceList":      nsList,

			"ClusterList": clusterList,
		})
}
