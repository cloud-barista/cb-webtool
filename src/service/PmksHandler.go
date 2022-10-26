package service

import (
	"encoding/json"
	"fmt"

	tbcommon "github.com/cloud-barista/cb-webtool/src/model/tumblebug/common"
	"github.com/labstack/echo"

	// "io"
	"log"
	"net/http"

	// "os"
	"github.com/cloud-barista/cb-webtool/src/model"
	spider "github.com/cloud-barista/cb-webtool/src/model/spider"

	util "github.com/cloud-barista/cb-webtool/src/util"
)

// Cluster 목록 조회
func GetPmksClusterList(clusterReqInfo spider.ClusterReqInfo) ([]spider.ClusterInfo, model.WebStatus) {
	var originalUrl = "/cluster"

	url := util.SPIDER + originalUrl
	pbytes, _ := json.Marshal(clusterReqInfo)
	resp, err := util.CommonHttp(url, pbytes, http.MethodGet)

	if err != nil {
		fmt.Println(err)
		return nil, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode

	clusterList := []spider.ClusterInfo{}
	json.NewDecoder(respBody).Decode(&clusterList)
	fmt.Println(clusterList)
	log.Println(respBody)

	return clusterList, model.WebStatus{StatusCode: respStatus}
}

// 특정 Cluster 조회
func GetPmksClusterData(cluster string, clusterReqInfo spider.ClusterReqInfo) (*spider.ClusterInfo, model.WebStatus) {
	var originalUrl = "/cluster/{cluster}"

	var paramMapper = make(map[string]string)
	paramMapper["{cluster}"] = cluster
	urlParam := util.MappingUrlParameter(originalUrl, paramMapper)

	url := util.SPIDER + urlParam

	pbytes, _ := json.Marshal(clusterReqInfo)
	resp, err := util.CommonHttp(url, pbytes, http.MethodGet)

	// defer body.Close()
	clusterInfo := spider.ClusterInfo{}
	if err != nil {
		fmt.Println(err)
		return &clusterInfo, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}
	// util.DisplayResponse(resp) // 수신내용 확인

	respBody := resp.Body
	respStatus := resp.StatusCode

	json.NewDecoder(respBody).Decode(&clusterInfo)
	fmt.Println(clusterInfo)

	return &clusterInfo, model.WebStatus{StatusCode: respStatus}
}

// Cluster 생성
func RegPmksCluster(nameSpaceID string, clusterReqInfo *spider.ClusterReqInfo) (*spider.ClusterInfo, model.WebStatus) {

	var originalUrl = "/cluster"

	url := util.SPIDER + originalUrl

	pbytes, _ := json.Marshal(clusterReqInfo)
	fmt.Println(string(pbytes))
	resp, err := util.CommonHttp(url, pbytes, http.MethodPost)

	returnClusterInfo := spider.ClusterInfo{}
	returnStatus := model.WebStatus{}

	if err != nil {
		fmt.Println(err)
		return &returnClusterInfo, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode

	if respStatus != 200 && respStatus != 201 { // 호출은 정상이나, 가져온 결과값이 200, 201아닌 경우 message에 담겨있는 것을 WebStatus에 set
		errorInfo := model.ErrorInfo{}
		json.NewDecoder(respBody).Decode(&errorInfo)
		fmt.Println("respStatus != 200 reason ", errorInfo)
		returnStatus.Message = errorInfo.Message
	} else {
		json.NewDecoder(respBody).Decode(&returnClusterInfo)
		fmt.Println(returnClusterInfo)
	}
	returnStatus.StatusCode = respStatus

	return &returnClusterInfo, returnStatus
}

// PMKS Cluster 생성
func RegPmksClusterByAsync(clusterReqInfo *spider.ClusterReqInfo, c echo.Context) {

	var originalUrl = "/cluster"

	url := util.SPIDER + originalUrl

	pbytes, _ := json.Marshal(clusterReqInfo)
	fmt.Println(string(pbytes))
	resp, err := util.CommonHttp(url, pbytes, http.MethodPost)

	nameSpaceID := clusterReqInfo.NameSpace
	clusterName := clusterReqInfo.ReqInfo.Name
	taskKey := nameSpaceID + "||" + "pmks" + "||" + clusterName // TODO : 공통 function으로 뺄 것.

	if err != nil {
		fmt.Println(err)
		StoreWebsocketMessage(util.TASK_TYPE_PMKS, taskKey, util.PMKS_LIFECYCLE_CREATE, util.TASK_STATUS_FAIL, c) // session에 작업내용 저장
	} else {

		respBody := resp.Body
		respStatus := resp.StatusCode

		if respStatus != 200 && respStatus != 201 { // 호출은 정상이나, 가져온 결과값이 200, 201아닌 경우 message에 담겨있는 것을 WebStatus에 set
			failResultInfo := tbcommon.TbSimpleMsg{}
			json.NewDecoder(respBody).Decode(&failResultInfo)
			log.Println("RegPmksByAsync ", failResultInfo)
			StoreWebsocketMessage(util.TASK_TYPE_PMKS, taskKey, util.PMKS_LIFECYCLE_CREATE, util.TASK_STATUS_FAIL, c) // session에 작업내용 저장

		} else {
			returnClusterInfo := spider.ClusterInfo{}
			json.NewDecoder(respBody).Decode(&returnClusterInfo)
			fmt.Println(returnClusterInfo)
			StoreWebsocketMessage(util.TASK_TYPE_PMKS, taskKey, util.PMKS_LIFECYCLE_CREATE, util.TASK_STATUS_COMPLETE, c) // session에 작업내용 저장
		}
	}

}

// PmksClusterUpdateProc : 현재는 버전만 upgrade. 추후 항목 update가 생기면 function 분리할 것
func UpdatePmksCluster(clusterReqInfo *spider.ClusterReqInfo) (spider.ClusterInfo, model.WebStatus) {
	var originalUrl = "/cluster/upgrade"

	url := util.SPIDER + originalUrl

	pbytes, _ := json.Marshal(clusterReqInfo)
	resp, err := util.CommonHttp(url, pbytes, http.MethodPut)
	util.DisplayResponse(resp) // 수신내용 확인
	// return body, err
	respBody := resp.Body
	respStatus := resp.StatusCode

	resultClusterInfo := spider.ClusterInfo{}
	if err != nil {
		fmt.Println(err)
		failResultInfo := tbcommon.TbSimpleMsg{}
		json.NewDecoder(respBody).Decode(&failResultInfo)
		return resultClusterInfo, model.WebStatus{StatusCode: 500, Message: failResultInfo.Message}
	}

	json.NewDecoder(respBody).Decode(&resultClusterInfo)

	return resultClusterInfo, model.WebStatus{StatusCode: respStatus}
}

// PMKS Cluster 삭제
func DelPmksCluster(cluster string, clusterReqInfo spider.ClusterReqInfo) (*spider.ClusterInfo, model.WebStatus) {
	var originalUrl = "/cluster/{cluster}"

	var paramMapper = make(map[string]string)
	paramMapper["{cluster}"] = cluster
	urlParam := util.MappingUrlParameter(originalUrl, paramMapper)

	url := util.SPIDER + urlParam

	pbytes, _ := json.Marshal(clusterReqInfo)

	resp, err := util.CommonHttp(url, pbytes, http.MethodDelete)
	util.DisplayResponse(resp) // 수신내용 확인

	resultClusterInfo := spider.ClusterInfo{}
	if err != nil {
		fmt.Println("delCluster ", err)
		return &resultClusterInfo, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode

	json.NewDecoder(respBody).Decode(&resultClusterInfo)
	fmt.Println(resultClusterInfo)

	if respStatus != 200 && respStatus != 201 {
		failResultInfo := tbcommon.TbSimpleMsg{}
		json.NewDecoder(respBody).Decode(&failResultInfo)
		fmt.Println(failResultInfo)
		return &resultClusterInfo, model.WebStatus{StatusCode: respStatus, Message: failResultInfo.Message}
	}
	return &resultClusterInfo, model.WebStatus{StatusCode: respStatus}
}

// Cluster 삭제 비동기 처리
func DelPmksClusterByAsync(cluster string, clusterReqInfo *spider.ClusterReqInfo, c echo.Context) {
	var originalUrl = "/cluster/{cluster}"

	var paramMapper = make(map[string]string)
	paramMapper["{cluster}"] = cluster
	urlParam := util.MappingUrlParameter(originalUrl, paramMapper)

	url := util.SPIDER + urlParam

	pbytes, _ := json.Marshal(clusterReqInfo)

	resp, err := util.CommonHttp(url, pbytes, http.MethodDelete)

	nameSpaceID := clusterReqInfo.NameSpace
	taskKey := nameSpaceID + "||" + "pmks" + "||" + cluster

	if err != nil {
		fmt.Println(err)
		StoreWebsocketMessage(util.TASK_TYPE_PMKS, taskKey, util.PMKS_LIFECYCLE_DELETE, util.TASK_STATUS_FAIL, c) // session에 작업내용 저장
	}

	respBody := resp.Body
	respStatus := resp.StatusCode

	if respStatus != 200 && respStatus != 201 { // 호출은 정상이나, 가져온 결과값이 200, 201아닌 경우 message에 담겨있는 것을 WebStatus에 set
		failResultInfo := tbcommon.TbSimpleMsg{}
		json.NewDecoder(respBody).Decode(&failResultInfo)
		log.Println("DelPmksByAsync ", failResultInfo)
		StoreWebsocketMessage(util.TASK_TYPE_PMKS, taskKey, util.PMKS_LIFECYCLE_DELETE, util.TASK_STATUS_FAIL, c) // session에 작업내용 저장

	} else {
		log.Println("DelPmksByAsync respBody : ", respBody)
		StoreWebsocketMessage(util.TASK_TYPE_PMKS, taskKey, util.PMKS_LIFECYCLE_DELETE, util.TASK_STATUS_COMPLETE, c) // session에 작업내용 저장
	}
}

// NodeGroup 생성
func RegPmksNodeGroup(clusterName string, nodeGroupReqInfo *spider.NodeGroupReqInfo) (*spider.NodeGroupInfo, model.WebStatus) {

	var originalUrl = "/cluster/{cluster}/nodegroup"

	var paramMapper = make(map[string]string)
	paramMapper["{cluster}"] = clusterName
	urlParam := util.MappingUrlParameter(originalUrl, paramMapper)
	url := util.SPIDER + urlParam

	pbytes, _ := json.Marshal(nodeGroupReqInfo)
	fmt.Println(string(pbytes))
	resp, err := util.CommonHttp(url, pbytes, http.MethodPost)

	returnNodeGroupInfo := spider.NodeGroupInfo{}
	returnStatus := model.WebStatus{}

	respBody := resp.Body
	respStatus := resp.StatusCode

	if err != nil {
		fmt.Println(err)
		return &returnNodeGroupInfo, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	if respStatus != 200 && respStatus != 201 { // 호출은 정상이나, 가져온 결과값이 200, 201아닌 경우 message에 담겨있는 것을 WebStatus에 set
		errorInfo := model.ErrorInfo{}
		json.NewDecoder(respBody).Decode(&errorInfo)
		fmt.Println("respStatus != 200 reason ", errorInfo)
		returnStatus.Message = errorInfo.Message
	} else {
		json.NewDecoder(respBody).Decode(&returnNodeGroupInfo)
		fmt.Println(returnNodeGroupInfo)
	}
	returnStatus.StatusCode = respStatus

	return &returnNodeGroupInfo, returnStatus
}

// NodeGroup 삭제
func DelPmksNodeGroup(clusterName string, nodeGroupName string, nodeGroupReqInfo *spider.NodeGroupReqInfo) (bool, model.WebStatus) {
	var originalUrl = "/cluster/{cluster}/nodegroup/{nodegroup}"

	var paramMapper = make(map[string]string)
	paramMapper["{cluster}"] = clusterName
	paramMapper["{nodegroup}"] = nodeGroupName
	urlParam := util.MappingUrlParameter(originalUrl, paramMapper)
	url := util.SPIDER + urlParam

	if clusterName == "" {
		return false, model.WebStatus{StatusCode: 500, Message: "cluster is required"}
	}
	if nodeGroupName == "" {
		return false, model.WebStatus{StatusCode: 500, Message: "nodeGroupName is required"}
	}

	pbytes, _ := json.Marshal(nodeGroupReqInfo)
	resp, err := util.CommonHttp(url, pbytes, http.MethodDelete)
	if err != nil {
		fmt.Println(err)
		return false, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	util.DisplayResponse(resp) // 수신내용 확인

	//respBody := resp.Body
	respStatus := resp.StatusCode

	return true, model.WebStatus{StatusCode: respStatus}
}
