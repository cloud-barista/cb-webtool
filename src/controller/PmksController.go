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

func PmksRegForm(c echo.Context) error {
	fmt.Println("PmksRegForm ************ : ")

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

	clusterList, _ := service.GetPmksClusterList(defaultNameSpaceID)

	return echotemplate.Render(c, http.StatusOK,
		"operation/manages/pmksmng/PmksCreate", // 파일명
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

func PmksMngForm(c echo.Context) error {
	fmt.Println("PmksMngForm ************ : ")

	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	pmksSimpleInfoList := []ladybug.ClusterSimpleInfo{} // 표에 뿌려줄 정보
	totalPmksStatusCountMap := make(map[string]int)
	totalClusterCount := 0

	// 최신 namespacelist 가져오기
	nsList, _ := service.GetNameSpaceList()
	log.Println(" nsList  ", nsList)

	// provider 별 연결정보 count
	cloudConnectionConfigInfoList, _ := service.GetCloudConnectionConfigList()
	connectionConfigCountMap, providerCount := service.GetCloudConnectionCountMap(cloudConnectionConfigInfoList)
	totalConnectionCount := len(cloudConnectionConfigInfoList)

	// 모든 MCKS 조회
	clusterList, clusterErr := service.GetPmksClusterList(defaultNameSpaceID)
	if clusterErr.StatusCode != 200 && clusterErr.StatusCode != 201 {
		echotemplate.Render(c, http.StatusOK,
			"operation/manages/pmksmng/PmksMng", // 파일명
			map[string]interface{}{
				"Message":            clusterErr.Message,
				"Status":             clusterErr.StatusCode,
				"LoginInfo":          loginInfo,
				"DefaultNameSpaceID": defaultNameSpaceID,
				"NameSpaceList":      nsList,

				// cp count 영역
				"TotalProviderCount":         providerCount,
				"TotalConnectionConfigCount": totalConnectionCount,     // 총 connection 갯수
				"ConnectionConfigCountMap":   connectionConfigCountMap, // provider별 connection 수

				// "ClusterList": clusterList,
				"ClusterList":             pmksSimpleInfoList,
				"TotalPmksStatusCountMap": totalPmksStatusCountMap,
				"TotalClusterCount":       totalClusterCount,
			})
	}

	totalClusterCount = len(clusterList)
	if totalClusterCount == 0 {
		return c.Redirect(http.StatusTemporaryRedirect, "/operation/manages/pmksmng/regform")
	}

	// totalPmksStatusCountMap = service.GetPmksStatusCountMap(clusterList)
	// ////////////// return value 에 set

	// for _, pmksInfo := range clusterList {
	// 	pmksSimpleInfo := ladybug.ClusterSimpleInfo{}
	// 	pmksSimpleInfo.UID = pmksInfo.UID
	// 	pmksSimpleInfo.Status = pmksInfo.Status
	// 	pmksSimpleInfo.PmksStatus = util.GetPmksStatus(pmksInfo.Status)
	// 	pmksSimpleInfo.Name = pmksInfo.Name
	// 	pmksSimpleInfo.ClusterConfig = pmksInfo.ClusterConfig
	// 	pmksSimpleInfo.CpLeader = pmksInfo.CpLeader
	// 	pmksSimpleInfo.Kind = pmksInfo.Kind
	// 	pmksSimpleInfo.Mcis = pmksInfo.Mcis
	// 	pmksSimpleInfo.NameSpace = pmksInfo.NameSpace
	// 	pmksSimpleInfo.NetworkCni = pmksInfo.NetworkCni

	// 	resultSimpleNodeList, resultSimpleNodeRoleCountMap := service.GetSimpleNodeCountMap(pmksInfo)

	// 	pmksSimpleInfo.Nodes = resultSimpleNodeList
	// 	pmksSimpleInfo.TotalNodeCount = len(resultSimpleNodeList) // 해당 pmks의 모든 node 갯수
	// 	//nodeKindCountMapByPmks[pmksInfo.UID] = resultSimpleNodeKindCountMap // MCIS 내 vm 상태별 cnt
	// 	pmksSimpleInfo.NodeCountMap = resultSimpleNodeRoleCountMap // MCKS UID 별 KindCountMap
	// 	// pmksSimpleInfo.NodeSimpleList = resultSimpleNodeList

	// 	// log.Println("**************")
	// 	// mapValues, _ := util.StructToMapByJson(pmksSimpleInfo)
	// 	// log.Println(mapValues)
	// 	// log.Println("**************")

	// 	pmksSimpleInfoList = append(pmksSimpleInfoList, pmksSimpleInfo)
	// }

	// status, filepath, return params
	return echotemplate.Render(c, http.StatusOK,
		"operation/manages/pmksmng/PmksMng", // 파일명
		map[string]interface{}{
			"Message":            clusterErr.Message,
			"Status":             clusterErr.StatusCode,
			"LoginInfo":          loginInfo,
			"DefaultNameSpaceID": defaultNameSpaceID,
			"NameSpaceList":      nsList,

			// cp count 영역
			// "TotalProviderCount":         providerCount,
			// "TotalConnectionConfigCount": totalConnectionCount,     // 총 connection 갯수
			// "ConnectionConfigCountMap":   connectionConfigCountMap, // provider별 connection 수

			// "ClusterList": clusterList,
			"ClusterList": pmksSimpleInfoList,
			// "TotalPmksStatusCountMap": totalPmksStatusCountMap,
			// "TotalClusterCount":       totalClusterCount,
		})
}

// MCKS 목록 조회
func GetPmksList(c echo.Context) error {
	log.Println("GetPmksList : ")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	optionParam := c.QueryParam("option")
	if optionParam == "id" {
		pmksList, respStatus := service.GetPmksClusterListByID(defaultNameSpaceID)
		if respStatus.StatusCode != 200 && respStatus.StatusCode != 201 {
			return c.JSON(respStatus.StatusCode, map[string]interface{}{
				"error":  respStatus.Message,
				"status": respStatus.StatusCode,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message":            "success",
			"status":             respStatus.StatusCode,
			"DefaultNameSpaceID": defaultNameSpaceID,
			"PmksList":           pmksList,
		})
	} else {
		pmksList, respStatus := service.GetPmksClusterList(defaultNameSpaceID)
		if respStatus.StatusCode != 200 && respStatus.StatusCode != 201 {
			return c.JSON(respStatus.StatusCode, map[string]interface{}{
				"error":  respStatus.Message,
				"status": respStatus.StatusCode,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message":            "success",
			"status":             respStatus.StatusCode,
			"DefaultNameSpaceID": defaultNameSpaceID,
			"PmksList":           pmksList,
		})
	}
}

// Cluster 등록 처리
func PmksRegProc(c echo.Context) error {
	log.Println("PmksRegProc : ")
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

	// Async로 변경
	taskKey := defaultNameSpaceID + "||" + "pmks" + "||" + clusterReq.Name                                               // TODO : 공통 function으로 뺄 것.
	service.StoreWebsocketMessage(util.TASK_TYPE_MCKS, taskKey, util.MCKS_LIFECYCLE_CREATE, util.TASK_STATUS_REQUEST, c) // session에 작업내용 저장
	go service.RegPmksClusterByAsync(defaultNameSpaceID, clusterReq, c)
	// 원래는 호출 결과를 return하나 go routine으로 바꾸면서 요청성공으로 return
	log.Println("before return")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"status":  200,
	})
}

// MCKS 삭제처리
func PmksDelProc(c echo.Context) error {
	log.Println("PmksDelProc : ")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	defaultNameSpaceID := loginInfo.DefaultNameSpaceID
	// TODO : defaultNameSpaceID 가 없으면 설정화면으로 보낼 것

	//clusteruID := c.Param("clusteruID")
	clusterName := c.Param("clusterName")
	log.Println("clusterName= " + clusterName)

	taskKey := defaultNameSpaceID + "||" + "pmks" + "||" + clusterName
	service.StoreWebsocketMessage(util.TASK_TYPE_MCKS, taskKey, util.MCKS_LIFECYCLE_DELETE, util.TASK_STATUS_REQUEST, c) // session에 작업내용 저장

	// resultStatusInfo, respStatus := service.DelCluster(defaultNameSpaceID, clusterName)
	//log.Println("DelMCKS service returned")
	//if respStatus.StatusCode != 200 && respStatus.StatusCode != 201 {
	//	return c.JSON(respStatus.StatusCode, map[string]interface{}{
	//		"error":  respStatus.Message,
	//		"status": respStatus.StatusCode,
	//	})
	//}
	//
	//return c.JSON(http.StatusOK, map[string]interface{}{
	//	"message":    "success",
	//	"status":     respStatus.StatusCode,
	//	"StatusInfo": resultStatusInfo,
	//})

	go service.DelPmksClusterByAsync(defaultNameSpaceID, clusterName, c)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"status":  200,
	})
}

// Node 등록 form
func PmksNodeGroupRegForm(c echo.Context) error {
	fmt.Println("PmksNodeGroupRegForm ************ : ")

	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}
	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	clusterUID := c.Param("clusterUID")
	clusterName := c.Param("clusterName")

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

	nodeList, _ := service.GetPmksNodeList(defaultNameSpaceID, clusterName)
	nodeListLength := len(nodeList.Items)
	log.Println("---------------------- nodeListLength ", nodeListLength)

	return echotemplate.Render(c, http.StatusOK,
		"operation/manages/pmksmng/NodeCreate", // 파일명
		map[string]interface{}{
			"LoginInfo":          loginInfo,
			"DefaultNameSpaceID": defaultNameSpaceID,
			"NameSpaceList":      nsList,
			"CloudOSList":        cloudOsList,
			"RegionList":         regionList,

			"CloudConnectionConfigInfoList": cloudConnectionConfigInfoList,
			"NodeList":                      nodeList,
			"PmksID":                        clusterUID,
			"PmksName":                      clusterName,
		})
}

// NodeGroup 등록 처리
func PmksNodeGroupRegProc(c echo.Context) error {
	log.Println("NodeRegProc : ")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	clusteruID := c.Param("clusteruID")
	clusterName := c.Param("clusterName")

	nodeRegReq := &ladybug.NodeRegReq{}
	// nodeRegReq := &ladybug.NodeReq{}
	if err := c.Bind(nodeRegReq); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
			"status":  "fail",
		})
	}
	log.Println(nodeRegReq)

	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	nodeInfo, respStatus := service.RegPmksNodeGroup(defaultNameSpaceID, clusterName, nodeRegReq)
	log.Println("RegNodeGroup service returned")
	if respStatus.StatusCode != 200 && respStatus.StatusCode != 201 {
		return c.JSON(respStatus.StatusCode, map[string]interface{}{
			"error":  respStatus.Message,
			"status": respStatus.StatusCode,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":    "success",
		"status":     respStatus.StatusCode,
		"ClusteruID": clusteruID,
		"NodeInfo":   nodeInfo,
	})
}

// Node 삭제 처리
func PmksNodeGroupDelProc(c echo.Context) error {
	log.Println("NodeGroupRegProc : ")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	clusterName := c.Param("clusterName")
	nodeName := c.Param("nodeName")

	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	resultStatusInfo, respStatus := service.DelPmksNodeGroup(defaultNameSpaceID, clusterName, nodeName)
	log.Println("DelMCKS service returned")
	if respStatus.StatusCode != 200 && respStatus.StatusCode != 201 {
		return c.JSON(respStatus.StatusCode, map[string]interface{}{
			"error":  respStatus.Message,
			"status": respStatus.StatusCode,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":    "success",
		"status":     respStatus.StatusCode,
		"StatusInfo": resultStatusInfo,
	})
}
