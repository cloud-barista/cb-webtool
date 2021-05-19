package controller

import (
	// "encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/cloud-barista/cb-webtool/src/model/ladybug"
	// "github.com/cloud-barista/cb-webtool/src/model/dragonfly"
	// "github.com/cloud-barista/cb-webtool/src/model/tumblebug"

	service "github.com/cloud-barista/cb-webtool/src/service"
	util "github.com/cloud-barista/cb-webtool/src/util"

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

	// 모든 MCKS 조회
	clusterList, _ := service.GetClusterList(defaultNameSpaceID)

	totalClusterCount := len(clusterList)
	if totalClusterCount == 0 {
		return c.Redirect(http.StatusTemporaryRedirect, "/operation/manages/mcksmng/regform")
	}

	totalMcksStatusCountMap := service.GetMcksStatusCountMap(clusterList)
	
	nodeKindCountMapByMcks := make(map[string]map[string]int)   // MCKS UID 별 KindCountMap
	////////////// return value 에 set
	mcksSimpleInfoList := []ladybug.ClusterSimpleInfo{}	// 표에 뿌려줄 정보
	for _, mcksInfo := range clusterList {
		mcksSimpleInfo := ladybug.ClusterSimpleInfo{}
		mcksSimpleInfo.UID = mcksInfo.UID
		mcksSimpleInfo.Status = mcksInfo.Status
		mcksSimpleInfo.McisStatus = util.GetMcksStatus(mcksInfo.Status)
		mcksSimpleInfo.Name = mcksInfo.Name

		resultSimpleNodeList, resultSimpleNodeKindCountMap := service.GetSimpleNodeCountMap(mcksInfo)

		mcksSimpleInfo.Nodes = resultSimpleNodeList
		mcksSimpleInfo.TotalNodeCount = len(resultSimpleNodeList) // 해당 mcks의 모든 node 갯수
		nodeKindCountMapByMcks[mcksInfo.UID] = resultSimpleNodeKindCountMap // MCIS 내 vm 상태별 cnt
		// mcksSimpleInfo.NodeSimpleList = resultSimpleNodeList
		
		mcksSimpleInfoList = append(mcksSimpleInfoList, mcksSimpleInfo)
	}

	// status, filepath, return params
	return echotemplate.Render(c, http.StatusOK,
		"operation/manages/mcksmng/McksMng", // 파일명
		map[string]interface{}{
			"LoginInfo":          loginInfo,
			"DefaultNameSpaceID": defaultNameSpaceID,
			"NameSpaceList":      nsList,

			// "ClusterList": clusterList,
			"ClusterList": mcksSimpleInfoList,
			"TotalMcksStatusCountMap": totalMcksStatusCountMap,
			"NodeKindCountMapByMcks": nodeKindCountMapByMcks,
			"TotalClusterCount": totalClusterCount,
		})
}

// Cluster 등록 처리
func McksRegProc(c echo.Context) error {
	log.Println("McksRegProc : ")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	log.Println("get info")
	//&[]Person{}
	clusterReq := &ladybug.ClusterRegReq{}
	if err := c.Bind(clusterReq); err != nil {
		// if err := c.Bind(mCISInfoList); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
			"status":  "fail",
		})
	}
	log.Println(clusterReq)

	defaultNameSpaceID := loginInfo.DefaultNameSpaceID
	// TODO : defaultNameSpaceID 가 없으면 설정화면으로 보낼 것
	clusterInfo, respStatus := service.RegCluster(defaultNameSpaceID, clusterReq)
	log.Println("RegMcis service returned")
	if respStatus.StatusCode != 200 && respStatus.StatusCode != 201 {
		return c.JSON(respStatus.StatusCode, map[string]interface{}{
			"error":  respStatus.Message,
			"status": respStatus.StatusCode,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":     "success",
		"status":      respStatus.StatusCode,
		"ClusterInfo": clusterInfo,
	})
}

// Node 등록 form
func McksNodeRegForm(c echo.Context) error {
	fmt.Println("McksNodeRegForm ************ : ")

	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}
	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	clusteruid := c.Param("clusteruid")

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

	ndeList, _ := service.GetNodeList(defaultNameSpaceID, clusteruid)

	return echotemplate.Render(c, http.StatusOK,
		"operation/manages/mcksmng/NodeCreate", // 파일명
		map[string]interface{}{
			"LoginInfo":          loginInfo,
			"DefaultNameSpaceID": defaultNameSpaceID,
			"NameSpaceList":      nsList,
			"CloudOSList":        cloudOsList,
			"RegionList":         regionList,

			"CloudConnectionConfigInfoList": cloudConnectionConfigInfoList,
			"NodeList":                      ndeList,
			"ClusterUid":                    clusteruid,
		})
}

// Node 등록 처리
func NodeRegProc(c echo.Context) error {
	log.Println("NodeRegProc : ")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	clusteruid := c.Param("clusteruid")

	nodeRegReq := &ladybug.NodeRegReq{}
	if err := c.Bind(nodeRegReq); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
			"status":  "fail",
		})
	}
	log.Println(nodeRegReq)

	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	nodeInfo, respStatus := service.RegNode(defaultNameSpaceID, clusteruid, nodeRegReq)
	log.Println("RegMcis service returned")
	if respStatus.StatusCode != 200 && respStatus.StatusCode != 201 {
		return c.JSON(respStatus.StatusCode, map[string]interface{}{
			"error":  respStatus.Message,
			"status": respStatus.StatusCode,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":    "success",
		"status":     respStatus.StatusCode,
		"Clusteruid": clusteruid,
		"NodeInfo":   nodeInfo,
	})
}
