package service

import (
	"encoding/json"
	"fmt"
	"strings"

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

// 해당 namespace의 모든 pmks 목록 조회
func GetPmksNamespaceClusterList(clusterReqInfo spider.ClusterReqInfo) ([]spider.SpClusterInfo, model.WebStatus) {
	// func GetPmksNamespaceClusterList(clusterReqInfo spider.ClusterReqInfo) ([]spider.SpAllClusterInfoList, model.WebStatus) {
	var originalUrl = "/nscluster"

	url := util.SPIDER + originalUrl
	pbytes, _ := json.Marshal(clusterReqInfo)
	resp, err := util.CommonHttp(url, pbytes, http.MethodGet)

	//returnClusterList := spider.SpTotalClusterInfoList{}
	returnClusterList := []spider.SpClusterInfo{}
	//returnClusterList := []spider.SpAllClusterInfoList{}
	if err != nil {
		fmt.Println(err)
		return returnClusterList, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}
	respBody := resp.Body
	respStatus := resp.StatusCode

	fmt.Println(respStatus)
	fmt.Println(respBody)

	if respStatus != 200 && respStatus != 201 { // 호출은 정상이나, 가져온 결과값이 200, 201아닌 경우 message에 담겨있는 것을 WebStatus에 set
		errorInfo := model.ErrorInfo{}
		json.NewDecoder(respBody).Decode(&errorInfo)
		fmt.Println("respStatus != 200 reason ", errorInfo)

		//spew.Dump(respBody)
		return returnClusterList, model.WebStatus{StatusCode: respStatus, Message: errorInfo.Message}
	} else {
		totalClusterList := spider.SpTotalClusterInfoList{}
		json.NewDecoder(respBody).Decode(&totalClusterList)
		//json.NewDecoder(respBody).Decode(&returnClusterList)

		// AllClusterList배열의 Cluster배열 형태를  ClusterInfo의 배열로 변환

		for _, clusterList := range totalClusterList.AllClusterList {
			log.Println(clusterList)

			for _, cluster := range clusterList.ClusterList {
				cluster.ProviderName = clusterList.Provider
				cluster.ConnectionName = clusterList.Connection
				returnClusterList = append(returnClusterList, cluster)
			}

			//returnClusterList = append(returnClusterList, clusterList.ClusterList...)
			//returnClusterList = append(returnClusterList, clusterList)
		}
	}

	return returnClusterList, model.WebStatus{StatusCode: respStatus}
}

// Cluster 목록 조회
func GetPmksClusterList(clusterReqInfo spider.ClusterReqInfo) ([]spider.SpClusterInfo, model.WebStatus) {
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

	clusterList := map[string][]spider.SpClusterInfo{}
	json.NewDecoder(respBody).Decode(&clusterList)
	fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^")
	fmt.Println(clusterList)
	//log.Println(respBody)

	return clusterList["cluster"], model.WebStatus{StatusCode: respStatus}
}

// 특정 Cluster 조회
func GetPmksClusterData(cluster string, clusterReqInfo spider.ClusterReqInfo) (*spider.SpClusterInfo, model.WebStatus) {
	var originalUrl = "/cluster/{cluster}"

	var paramMapper = make(map[string]string)
	paramMapper["{cluster}"] = cluster
	urlParam := util.MappingUrlParameter(originalUrl, paramMapper)

	url := util.SPIDER + urlParam

	pbytes, _ := json.Marshal(clusterReqInfo)
	resp, err := util.CommonHttp(url, pbytes, http.MethodGet)

	// defer body.Close()
	clusterInfo := spider.RespClusterInfo{}
	if err != nil {
		fmt.Println(err)
		return nil, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}
	// util.DisplayResponse(resp) // 수신내용 확인

	respBody := resp.Body
	respStatus := resp.StatusCode

	json.NewDecoder(respBody).Decode(&clusterInfo)
	fmt.Println(clusterInfo)

	clusterInfo.ClusterInfo.ConnectionName = clusterReqInfo.ConnectionName
	return &clusterInfo.ClusterInfo, model.WebStatus{StatusCode: respStatus}
}

// Cluster 생성
func RegPmksCluster(nameSpaceID string, clusterReqInfo *spider.ClusterReqInfo) (*spider.SpClusterInfo, model.WebStatus) {

	var originalUrl = "/cluster"

	url := util.SPIDER + originalUrl

	pbytes, _ := json.Marshal(clusterReqInfo)
	fmt.Println(string(pbytes))
	resp, err := util.CommonHttp(url, pbytes, http.MethodPost)

	returnClusterInfo := spider.SpClusterInfo{}
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
		//Message: {"code":"400","message":"no managedkubernetes ros component exists. version: 1, labels: ap-southeast-1:common:26888:5513479151634744:GC0","requestId":"67CAB7BC-C0E1-3D93-A967-C89AE46138B2","status":400}
		//Message: {"code":"QuotaExceeded.Cluster",
		//"message":"Exceeded the quota for creating a cluster
		//			(quota code: q_ManagedKubernetes_Default_ack.standard_Cluster)
		//			, usage 2/2.","requestId":"EEF6F445-4549-39C7-B8AB-DB83B843A296","status":400}

		errMsg := err.Error()

		//StoreWebsocketMessage(util.TASK_TYPE_PMKS, taskKey, util.PMKS_LIFECYCLE_CREATE, util.TASK_STATUS_FAIL, c)
		StoreWebsocketMessageDetail(util.TASK_TYPE_PMKS, taskKey, util.PMKS_LIFECYCLE_CREATE, util.TASK_STATUS_FAIL, errMsg, c)
	} else {

		respBody := resp.Body
		respStatus := resp.StatusCode
		//spew.Dump(resp)
		if respStatus != 200 && respStatus != 201 { // 호출은 정상이나, 가져온 결과값이 200, 201아닌 경우 message에 담겨있는 것을 WebStatus에 set
			failResultInfo := spider.SpError{}
			json.NewDecoder(respBody).Decode(&failResultInfo)
			log.Println("RegPmksByAsync failed ", failResultInfo)
			log.Println("RegPmksByAsync failed-------- ", failResultInfo.Message)

			//jsonBytes, _ := json.Marshal(failResultInfo) // JSON ENCODING
			//jsonString := string(jsonBytes)
			//fmt.Println("personA String: ", jsonString)

			//var result map[string]interface{}
			//if err := json.Unmarshal([]byte(resp), &result); err != nil {
			//	panic(err)
			//}

			//fmt.Println(result["kind"])       // Event 출력
			//fmt.Println(result["apiVersion"]) // events.k8s.io/v1 출력

			beginIndex := strings.Index(failResultInfo.Message, "Message:")
			var endMessage string
			if beginIndex == -1 {
				fmt.Println("cannot find Message:  ------------")
				endMessage = failResultInfo.Message
			} else {
				findMessageBegin := failResultInfo.Message[beginIndex:]
				fmt.Println("findMessageBegin :  ", findMessageBegin)

				beginIndex2 := strings.Index(findMessageBegin, "{")
				findMessage := findMessageBegin[beginIndex2:]
				fmt.Println("findMessage :  ", findMessage)

				endIndex := strings.Index(findMessage, "}")
				endMessage = findMessage[:endIndex]
				fmt.Println("endMessage :  ", endMessage)
			}
			//findMessageBegin := strings.Index(failResultInfo.Message, "}"))

			//byt := []byte(endMessage)
			//var dat map[string]interface{}
			//if err := json.Unmarshal(byt, &dat); err != nil {
			//panic(err)
			//}
			//fmt.Println("********************")
			//fmt.Println(dat)

			//var errorDetailObj = spider.SpErrorDetail{}
			//err := json.Unmarshal([]byte(endMessage), &errorDetailObj)
			//if err != nil {
			//	fmt.Println(err)
			//}
			//fmt.Println("personObj Object: ", errorDetailObj)
			//detailInfo := spider.SpErrorDetail{}
			//json.NewDecoder(failResultInfo).Decode(&detailInfo)
			//log.Println("detail failed ", detailInfo)

			//StoreWebsocketMessage(util.TASK_TYPE_PMKS, taskKey, util.PMKS_LIFECYCLE_CREATE, util.TASK_STATUS_FAIL, c) // session에 작업내용 저장
			StoreWebsocketMessageDetail(util.TASK_TYPE_PMKS, taskKey, util.PMKS_LIFECYCLE_CREATE, util.TASK_STATUS_FAIL, endMessage, c)
		} else {
			returnClusterInfo := spider.ClusterInfo{}
			json.NewDecoder(respBody).Decode(&returnClusterInfo)
			fmt.Println(returnClusterInfo)
			StoreWebsocketMessage(util.TASK_TYPE_PMKS, taskKey, util.PMKS_LIFECYCLE_CREATE, util.TASK_STATUS_COMPLETE, c) // session에 작업내용 저장
		}
	}

}

// PmksClusterUpdateProc : 현재는 버전만 upgrade. 추후 항목 update가 생기면 function 분리할 것
func UpdatePmksCluster(clusterReqInfo *spider.ClusterReqInfo) (spider.SpClusterInfo, model.WebStatus) {
	var originalUrl = "/cluster/upgrade"

	url := util.SPIDER + originalUrl

	pbytes, _ := json.Marshal(clusterReqInfo)
	resp, err := util.CommonHttp(url, pbytes, http.MethodPut)
	util.DisplayResponse(resp) // 수신내용 확인
	// return body, err
	respBody := resp.Body
	respStatus := resp.StatusCode

	resultClusterInfo := spider.SpClusterInfo{}
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
func DelPmksCluster(cluster string, clusterReqInfo spider.ClusterReqInfo) (*spider.SpClusterInfo, model.WebStatus) {
	var originalUrl = "/cluster/{cluster}"

	var paramMapper = make(map[string]string)
	paramMapper["{cluster}"] = cluster
	urlParam := util.MappingUrlParameter(originalUrl, paramMapper)

	url := util.SPIDER + urlParam

	pbytes, _ := json.Marshal(clusterReqInfo)

	resp, err := util.CommonHttp(url, pbytes, http.MethodDelete)
	util.DisplayResponse(resp) // 수신내용 확인

	resultClusterInfo := spider.SpClusterInfo{}
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
		errMsg := err.Error()
		//StoreWebsocketMessage(util.TASK_TYPE_PMKS, taskKey, util.PMKS_LIFECYCLE_DELETE, util.TASK_STATUS_FAIL, c) // session에 작업내용 저장
		StoreWebsocketMessageDetail(util.TASK_TYPE_PMKS, taskKey, util.PMKS_LIFECYCLE_CREATE, util.TASK_STATUS_FAIL, errMsg, c)

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
func RegPmksNodeGroup(clusterID string, nodeGroupReqInfo *spider.NodeGroupReqInfo) (*spider.NodeGroupInfo, model.WebStatus) {

	var originalUrl = "/cluster/{cluster}/nodegroup"

	var paramMapper = make(map[string]string)
	paramMapper["{cluster}"] = clusterID
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
func DelPmksNodeGroup(clusterID string, nodeGroupID string, nodeGroupReqInfo *spider.NodeGroupReqInfo) (bool, model.WebStatus) {
	var originalUrl = "/cluster/{cluster}/nodegroup/{nodegroup}"

	var paramMapper = make(map[string]string)
	paramMapper["{cluster}"] = clusterID
	paramMapper["{nodegroup}"] = nodeGroupID
	urlParam := util.MappingUrlParameter(originalUrl, paramMapper)
	url := util.SPIDER + urlParam

	if clusterID == "" {
		return false, model.WebStatus{StatusCode: 500, Message: "cluster is required"}
	}
	if nodeGroupID == "" {
		return false, model.WebStatus{StatusCode: 500, Message: "nodeGroup is required"}
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

// NodeGroup 수정 : onAutoScaling
func UpdatePmksNodeGroupAutoScaling(clusterID string, nodeGroupID string, nodeGroupReqInfo *spider.NodeGroupReqInfo) (spider.SpClusterInfo, model.WebStatus) {
	var originalUrl = "/cluster/{cluster}/nodegroup/{nodegroup}/onautoscaling"

	var paramMapper = make(map[string]string)
	paramMapper["{cluster}"] = clusterID
	paramMapper["{nodegroup}"] = nodeGroupID
	urlParam := util.MappingUrlParameter(originalUrl, paramMapper)
	url := util.SPIDER + urlParam

	pbytes, _ := json.Marshal(nodeGroupReqInfo)
	resp, err := util.CommonHttp(url, pbytes, http.MethodPut)
	util.DisplayResponse(resp) // 수신내용 확인
	// return body, err
	respBody := resp.Body
	respStatus := resp.StatusCode

	resultClusterInfo := spider.SpClusterInfo{}
	if err != nil {
		fmt.Println(err)
		failResultInfo := tbcommon.TbSimpleMsg{}
		json.NewDecoder(respBody).Decode(&failResultInfo)
		return resultClusterInfo, model.WebStatus{StatusCode: 500, Message: failResultInfo.Message}
	}

	json.NewDecoder(respBody).Decode(&resultClusterInfo)

	return resultClusterInfo, model.WebStatus{StatusCode: respStatus}
}

// NodeGroup 수정 : node Size
func UpdatePmksNodeGroupAutoscaleSize(clusterID string, nodeGroupID string, nodeGroupReqInfo *spider.NodeGroupReqInfo) (spider.SpClusterInfo, model.WebStatus) {
	var originalUrl = "/cluster/{cluster}/nodegroup/{nodegroup}/autoscalesize"

	var paramMapper = make(map[string]string)
	paramMapper["{cluster}"] = clusterID
	paramMapper["{nodegroup}"] = nodeGroupID
	urlParam := util.MappingUrlParameter(originalUrl, paramMapper)
	url := util.SPIDER + urlParam

	pbytes, _ := json.Marshal(nodeGroupReqInfo)
	resp, err := util.CommonHttp(url, pbytes, http.MethodPut)
	util.DisplayResponse(resp) // 수신내용 확인
	// return body, err
	respBody := resp.Body
	respStatus := resp.StatusCode

	resultClusterInfo := spider.SpClusterInfo{}
	if err != nil {
		fmt.Println(err)
		failResultInfo := tbcommon.TbSimpleMsg{}
		json.NewDecoder(respBody).Decode(&failResultInfo)
		return resultClusterInfo, model.WebStatus{StatusCode: 500, Message: failResultInfo.Message}
	}

	json.NewDecoder(respBody).Decode(&resultClusterInfo)

	return resultClusterInfo, model.WebStatus{StatusCode: respStatus}
}
