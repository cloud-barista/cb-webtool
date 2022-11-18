package controller

import (
	// "encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	spider "github.com/cloud-barista/cb-webtool/src/model/spider"

	service "github.com/cloud-barista/cb-webtool/src/service"
	util "github.com/cloud-barista/cb-webtool/src/util"

	echotemplate "github.com/foolin/echo-template"
	"github.com/labstack/echo"
)

// PMKS Cluster 등록 form
func PmksClusterRegForm(c echo.Context) error {
	fmt.Println("PmksClusterRegForm ************ : ")

	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}
	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	// 최신 namespacelist 가져오기
	nsList, _ := service.GetNameSpaceList()
	log.Println(" nsList  ", nsList)

	cloudOsList, _ := service.GetCloudOSList()
	log.Println(" cloudOsList  ", cloudOsList)

	return echotemplate.Render(c, http.StatusOK,
		"operation/manages/pmksmng/PmksClusterCreate", // 파일명
		map[string]interface{}{
			"LoginInfo":          loginInfo,
			"DefaultNameSpaceID": defaultNameSpaceID,
			"NameSpaceList":      nsList,
			"CloudOSList":        cloudOsList,
		})
}

// PMKS NodeGroup 등록 form
func PmksNodeGroupRegForm(c echo.Context) error {
	fmt.Println("PmksNodeGroupRegForm ************ : ")

	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}
	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	clusterID := c.Param("clusterID")
	log.Println("clusterID= " + clusterID)
	connectionName := c.QueryParam("connectionName")
	log.Println("connectionName= " + connectionName)

	// 최신 namespacelist 가져오기
	nsList, _ := service.GetNameSpaceList()
	log.Println(" nsList  ", nsList)

	cloudOsList, _ := service.GetCloudOSList()
	log.Println(" cloudOsList  ", cloudOsList)

	//clusterReqInfo := spider.ClusterReqInfo{}
	// cluster 조회 // onload시 가져옴 대신 받은 Param은 바로 되돌려 줌.

	//clusterReqInfo := spider.ClusterReqInfo{}
	//clusterReqInfo.NameSpace = defaultNameSpaceID
	//clusterReqInfo.ConnectionName = connectionName
	//resultPmksInfo, _ := service.GetPmksClusterData(clusterID, clusterReqInfo)

	return echotemplate.Render(c, http.StatusOK,
		"operation/manages/pmksmng/PmksNodeGroupCreate", // 파일명
		map[string]interface{}{
			"LoginInfo":          loginInfo,
			"DefaultNameSpaceID": defaultNameSpaceID,
			"NameSpaceList":      nsList,
			"CloudOSList":        cloudOsList,
			"ClusterID":          clusterID,
			"ConnectionName":     connectionName,
			//"PmksInfo":           resultPmksInfo,
		})
}

// PMKS 관리화면
// 보통 등록된 것이 없으면 RegForm으로 보내는데 전체조회해서 redirect하는게 애매해서 그냥 mng화면을 보여줌
func PmksMngForm(c echo.Context) error {
	fmt.Println("PmksMngForm ************ : ")

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

	return echotemplate.Render(c, http.StatusOK,
		"operation/manages/pmksmng/PmksMng", // 파일명
		map[string]interface{}{
			"Message":            "success",
			"LoginInfo":          loginInfo,
			"DefaultNameSpaceID": defaultNameSpaceID,
			"NameSpaceList":      nsList,
			"CloudOSList":        cloudOsList,
			"RegionList":         regionList,

			"CloudConnectionConfigInfoList": cloudConnectionConfigInfoList,
		})
}

// namespace 내 모든 pmks 목록.
func GetPmksListOfNamespace(c echo.Context) error {
	log.Println("GetPmksListOfNamespace : ")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	clusterReqInfo := spider.ClusterReqInfo{}
	clusterReqInfo.NameSpace = defaultNameSpaceID

	// ","를 구분자로 하는 connectionNames 잘라서 다시 string 배열에 넣고 써야 함.
	connectionNames := c.QueryParam("connectionNames")
	if connectionNames != "" {
		connNameArr := strings.Split(connectionNames, ",")
		for _, connName := range connNameArr {
			clusterReqInfo.ConnectionName = connName
		}
	}

	pmksList, respStatus := service.GetPmksNamespaceClusterList(clusterReqInfo)
	log.Println("---------------------- result  ", pmksList)
	if respStatus.StatusCode != 200 && respStatus.StatusCode != 201 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": respStatus.Message,
			"status":  "fail",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":            "success",
		"status":             200,
		"DefaultNameSpaceID": defaultNameSpaceID,
		"PmksList":           pmksList,
	})
}

// PMKS 목록 조회 :

func GetPmksList(c echo.Context) error {
	log.Println("GetPmksList : ")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	clusterReqInfo := spider.ClusterReqInfo{}
	clusterReqInfo.NameSpace = defaultNameSpaceID

	connectionName := c.QueryParam("connectionName")
	if connectionName != "" {
		clusterReqInfo.ConnectionName = connectionName
	}

	pmksList, respStatus := service.GetPmksClusterList(clusterReqInfo)
	if respStatus.StatusCode != 200 && respStatus.StatusCode != 201 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": respStatus.Message,
			"status":  "fail",
		})
	}
	//cloudConnectionConfigInfoList, _ := service.GetCloudConnectionConfigList() // 등록된 모든 connection 정보
	//cloudConnectionConfigInfoList := []spider.CloudConnectionConfigInfo{}
	//conn := spider.CloudConnectionConfigInfo{}
	//conn.ConfigName = "ali-test-conn"
	//cloudConnectionConfigInfoList = append(cloudConnectionConfigInfoList, conn)
	//log.Println("---------------------- GetCloudConnectionConfigList ", defaultNameSpaceID)
	//totalPmksList := []spider.SpClusterInfo{}
	// totalPmksList := map[string][]spider.SpClusterInfo{}
	// 모든 connection의 pmks목록 조회
	//for _, connectionInfo := range cloudConnectionConfigInfoList {
	//	clusterReqInfo.ConnectionName = connectionInfo.ConfigName
	//	pmksList, respStatus := service.GetPmksClusterList(clusterReqInfo)
	//	log.Println("---------------------- result  ", pmksList)
	//	if respStatus.StatusCode != 200 && respStatus.StatusCode != 201 {
	//		continue
	//	}
	//	totalPmksList = append(totalPmksList, pmksList...)
	//totalPmksList[connectionInfo.ConfigName] = pmksList
	//}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":            "success",
		"status":             200,
		"DefaultNameSpaceID": defaultNameSpaceID,
		"PmksList":           pmksList,
	})

}

// PMKS 단건 조회
func GetPmksInfoData(c echo.Context) error {
	log.Println("GetPmksInfoData")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login") // 조회기능에서 바로 login화면으로 돌리지말고 return message로 하는게 낫지 않을까?
	}
	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	clusterID := c.Param("clusterID")
	log.Println("clusterID= " + clusterID)
	optionParam := c.QueryParam("connectionName")
	log.Println("optionParam= " + optionParam)

	clusterReqInfo := spider.ClusterReqInfo{}
	clusterReqInfo.NameSpace = defaultNameSpaceID
	clusterReqInfo.ConnectionName = optionParam
	resultPmksInfo, _ := service.GetPmksClusterData(clusterID, clusterReqInfo)

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

	clusterReqInfo := &spider.ClusterReqInfo{}
	if err := c.Bind(clusterReqInfo); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
			"status":  "fail",
		})
	}
	log.Println(clusterReqInfo)

	defaultNameSpaceID := loginInfo.DefaultNameSpaceID
	clusterReqInfo.NameSpace = defaultNameSpaceID

	taskKey := defaultNameSpaceID + "||" + "pmks" + "||" + clusterReqInfo.ReqInfo.Name                                   // TODO : 공통 function으로 뺄 것.
	service.StoreWebsocketMessage(util.TASK_TYPE_PMKS, taskKey, util.PMKS_LIFECYCLE_CREATE, util.TASK_STATUS_REQUEST, c) // session에 작업내용 저장
	go service.RegPmksClusterByAsync(clusterReqInfo, c)
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
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}
	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	clusterReqInfo := new(spider.ClusterReqInfo)
	if err := c.Bind(clusterReqInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	clusterReqInfo.NameSpace = defaultNameSpaceID

	taskKey := defaultNameSpaceID + "||" + "pmks" + "||" + clusterReqInfo.ReqInfo.Name
	service.StoreWebsocketMessage(util.TASK_TYPE_PMKS, taskKey, util.PMKS_CLUSTER_UPDATE, util.TASK_STATUS_REQUEST, c)

	respBody, respStatus := service.UpdatePmksCluster(clusterReqInfo)
	fmt.Println("=============respBody =============", respBody)
	if respStatus.StatusCode != 200 && respStatus.StatusCode != 201 {
		service.StoreWebsocketMessage(util.TASK_TYPE_PMKS, taskKey, util.PMKS_CLUSTER_UPDATE, util.TASK_STATUS_FAIL, c)

		return c.JSON(respStatus.StatusCode, map[string]interface{}{
			"error":  respStatus.Message,
			"status": respStatus.StatusCode,
		})
	}

	service.StoreWebsocketMessage(util.TASK_TYPE_PMKS, taskKey, util.PMKS_CLUSTER_UPDATE, util.TASK_STATUS_COMPLETE, c)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"status":  respStatus.StatusCode,
		"Result":  respBody,
	})
}

// PMKS 삭제처리
func PmksDelProc(c echo.Context) error {
	log.Println("PmksDelProc : ")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	clusterReqInfo := new(spider.ClusterReqInfo)
	if err := c.Bind(clusterReqInfo); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
			"status":  "fail",
		})
	}

	defaultNameSpaceID := loginInfo.DefaultNameSpaceID
	clusterReqInfo.NameSpace = defaultNameSpaceID

	clusterID := c.Param("clusterID")
	log.Println("clusterID= " + clusterID)
	connectionName := c.QueryParam("connectionName")
	log.Println("connectionName= " + connectionName)
	clusterReqInfo.ConnectionName = connectionName

	taskKey := defaultNameSpaceID + "||" + "pmks" + "||" + clusterID
	service.StoreWebsocketMessage(util.TASK_TYPE_PMKS, taskKey, util.PMKS_LIFECYCLE_DELETE, util.TASK_STATUS_REQUEST, c) // session에 작업내용 저장

	go service.DelPmksClusterByAsync(clusterID, clusterReqInfo, c)

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

	nodeGroupReqInfo := new(spider.NodeGroupReqInfo)
	if err := c.Bind(nodeGroupReqInfo); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
			"status":  "fail",
		})
	}

	clusterID := c.Param("clusterID")

	defaultNameSpaceID := loginInfo.DefaultNameSpaceID
	nodeGroupReqInfo.NameSpace = defaultNameSpaceID

	nodeInfo, respStatus := service.RegPmksNodeGroup(clusterID, nodeGroupReqInfo)
	log.Println("RegNodeGroup service returned")
	if respStatus.StatusCode != 200 && respStatus.StatusCode != 201 {
		return c.JSON(respStatus.StatusCode, map[string]interface{}{
			"error":  respStatus.Message,
			"status": respStatus.StatusCode,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success",
		"status":   respStatus.StatusCode,
		"NodeInfo": nodeInfo,
	})
}

// NodeGroup 삭제 처리
func PmksNodeGroupDelProc(c echo.Context) error {
	log.Println("NodeGroupRegProc : ")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	nodeGroupReqInfo := new(spider.NodeGroupReqInfo)
	if err := c.Bind(nodeGroupReqInfo); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
			"status":  "fail",
		})
	}

	clusterID := c.Param("clusterID")
	nodeGroupID := c.Param("nodeGroupID")

	defaultNameSpaceID := loginInfo.DefaultNameSpaceID
	nodeGroupReqInfo.NameSpace = defaultNameSpaceID

	resultStatusInfo, respStatus := service.DelPmksNodeGroup(clusterID, nodeGroupID, nodeGroupReqInfo)
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

// NodeGroup Update
// reqParameter에서 onautoscaling, nodesize, ...
func PmksNodeGroupUpdateProc(c echo.Context) error {
	log.Println("PmksNodeGroupUpdateProc : ")

	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	nodeGroupReqInfo := new(spider.NodeGroupReqInfo)
	if err := c.Bind(nodeGroupReqInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	clusterID := c.Param("clusterID")
	nodeGroupID := c.Param("nodeGroupID")
	defaultNameSpaceID := loginInfo.DefaultNameSpaceID
	nodeGroupReqInfo.NameSpace = defaultNameSpaceID

	//taskKey := defaultNameSpaceID + "||" + "pmks" + "||" + nodeGroupReqInfo.ReqInfo.Name
	//service.StoreWebsocketMessage(util.TASK_TYPE_PMKS, taskKey, util.PMKS_CLUSTER_UPDATE, util.TASK_STATUS_REQUEST, c) // session에 작업내용 저장

	// onautoscaling update
	if nodeGroupReqInfo.ReqInfo.OnAutoScaling != "" {
		respBody, respStatus := service.UpdatePmksNodeGroupAutoScaling(clusterID, nodeGroupID, nodeGroupReqInfo)
		fmt.Println("=============respBody =============", respBody)
		if respStatus.StatusCode != 200 && respStatus.StatusCode != 201 {
			//service.StoreWebsocketMessage(util.TASK_TYPE_PMKS, taskKey, util.PMKS_CLUSTER_UPDATE, util.TASK_STATUS_FAIL, c) // session에 작업내용 저장

			return c.JSON(respStatus.StatusCode, map[string]interface{}{
				"error":  respStatus.Message,
				"status": respStatus.StatusCode,
			})
		}
	}

	// node size update
	if nodeGroupReqInfo.ReqInfo.DesiredNodeSize != "" || nodeGroupReqInfo.ReqInfo.MaxNodeSize != "" || nodeGroupReqInfo.ReqInfo.MinNodeSize != "" {
		desiredNodeSize := nodeGroupReqInfo.ReqInfo.DesiredNodeSize
		MaxNodeSize := nodeGroupReqInfo.ReqInfo.MaxNodeSize
		MinNodeSize := nodeGroupReqInfo.ReqInfo.MinNodeSize
		desiredNodeSizeInt, err := strconv.ParseInt(desiredNodeSize, 10, 64)
		if err != nil {

		}
		MaxNodeSizeInt, err := strconv.ParseInt(MaxNodeSize, 10, 64)
		if err != nil {

		}
		MinNodeSizeInt, err := strconv.ParseInt(MinNodeSize, 10, 64)
		if err != nil {

		}

		if desiredNodeSizeInt > 0 && MaxNodeSizeInt > 0 && MinNodeSizeInt > 0 {
			respBody, respStatus := service.UpdatePmksNodeGroupAutoscaleSize(clusterID, nodeGroupID, nodeGroupReqInfo)
			fmt.Println("=============respBody =============", respBody)
			if respStatus.StatusCode != 200 && respStatus.StatusCode != 201 {
				//service.StoreWebsocketMessage(util.TASK_TYPE_PMKS, taskKey, util.PMKS_CLUSTER_UPDATE, util.TASK_STATUS_FAIL, c) // session에 작업내용 저장

				return c.JSON(respStatus.StatusCode, map[string]interface{}{
					"error":  respStatus.Message,
					"status": respStatus.StatusCode,
				})
			}
		}
	}

	//service.StoreWebsocketMessage(util.TASK_TYPE_PMKS, taskKey, util.PMKS_CLUSTER_UPDATE, util.TASK_STATUS_COMPLETE, c) // session에 작업내용 저장

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"status":  "200",
		"Result":  "Update Success",
	})
}
