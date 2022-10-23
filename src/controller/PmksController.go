package controller

import (
	// "encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/cloud-barista/cb-webtool/src/model/ladybug"
	spider "github.com/cloud-barista/cb-webtool/src/model/spider"

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

	// 모든 PMKS 조회
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

	// status, filepath, return params
	return echotemplate.Render(c, http.StatusOK,
		"operation/manages/pmksmng/PmksMng", // 파일명
		map[string]interface{}{
			"Message":            clusterErr.Message,
			"Status":             clusterErr.StatusCode,
			"LoginInfo":          loginInfo,
			"DefaultNameSpaceID": defaultNameSpaceID,
			"NameSpaceList":      nsList,

			"ClusterList": pmksSimpleInfoList,
		})
}

// PMKS 목록 조회
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

// PMKS 단건 조회
func GetPmksInfoData(c echo.Context) error {
	log.Println("GetPmksInfoData")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login") // 조회기능에서 바로 login화면으로 돌리지말고 return message로 하는게 낫지 않을까?
	}
	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	pmksID := c.Param("pmksID")
	log.Println("pmksID= " + pmksID)
	optionParam := c.QueryParam("option")
	log.Println("optionParam= " + optionParam)

	resultPmksInfo, _ := service.GetPmksClusterData(defaultNameSpaceID, pmksID)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success",
		"status":   200,
		"PmksInfo": resultPmksInfo,
	})
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
	service.StoreWebsocketMessage(util.TASK_TYPE_PMKS, taskKey, util.PMKS_LIFECYCLE_CREATE, util.TASK_STATUS_REQUEST, c) // session에 작업내용 저장
	go service.RegPmksClusterByAsync(defaultNameSpaceID, clusterReq, c)
	// 원래는 호출 결과를 return하나 go routine으로 바꾸면서 요청성공으로 return
	log.Println("before return")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"status":  200,
	})
}

// Cluster Update
func PmksClusterUpdateProc(c echo.Context) error {
	log.Println("PmksClusterUpdateProc : ")

	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		// Login 정보가 없으므로 login화면으로
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}
	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	clusterInfo := new(spider.ClusterReqInfo)
	if err := c.Bind(clusterInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	respBody, respStatus := service.UpdatePmksCluster(defaultNameSpaceID, clusterInfo)
	fmt.Println("=============respBody =============", respBody)
	if respStatus.StatusCode != 200 && respStatus.StatusCode != 201 {
		return c.JSON(respStatus.StatusCode, map[string]interface{}{
			"error":  respStatus.Message,
			"status": respStatus.StatusCode,
		})
	}

	// 저장 성공하면 namespace 목록 조회
	nameSpaceList, nsStatus := service.GetNameSpaceList()
	if nsStatus.StatusCode == 500 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":       nsStatus.Message,
			"status":        nsStatus.StatusCode,
			"NameSpaceList": nil,
		})
	}

	storeNameSpaceErr := service.SetStoreNameSpaceList(c, nameSpaceList)
	if storeNameSpaceErr != nil {
		log.Println("Store NameSpace Err")
	}

	// return namespace 목록
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":       "success",
		"status":        respStatus.StatusCode,
		"NameSpaceList": nameSpaceList,
	})
}

// PMKS 삭제처리
func PmksDelProc(c echo.Context) error {
	log.Println("PmksDelProc : ")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	//clusteruID := c.Param("clusteruID")
	clusterName := c.Param("clusterName")
	log.Println("clusterName= " + clusterName)

	taskKey := defaultNameSpaceID + "||" + "pmks" + "||" + clusterName
	service.StoreWebsocketMessage(util.TASK_TYPE_PMKS, taskKey, util.PMKS_LIFECYCLE_DELETE, util.TASK_STATUS_REQUEST, c) // session에 작업내용 저장

	go service.DelPmksClusterByAsync(defaultNameSpaceID, clusterName, c)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"status":  200,
	})
}

// NodeGroup 등록 처리
func PmksNodeGroupRegProc(c echo.Context) error {
	log.Println("PmksNodeGroupRegProc : ")
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

// NodeGroup 삭제 처리
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
	log.Println("DelPMKS service returned")
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
