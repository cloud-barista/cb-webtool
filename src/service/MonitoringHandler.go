package service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	// "os"
	model "github.com/cloud-barista/cb-webtool/src/model"
	util "github.com/cloud-barista/cb-webtool/src/util"
)

// VM 에 모니터링 Agent 설치
///ns/{nsId}/monitoring/install/mcis/{mcisId}
func RegMonitoringAgentInVm(nameSpaceID string, vmMonitoringAgentReg *model.VmMonitoringAgentReg) (*model.VmMonitoringAgentInfo, model.WebStatus) {
	fmt.Println("RegMonitoringAgentInVm ************ : ")

	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/monitoring/install/mcis/"
	// /ns/{nsId}/monitoring/install/mcis/{mcisId}

	pbytes, _ := json.Marshal(vmMonitoringAgentReg)
	fmt.Println(string(pbytes))
	resp, err := util.CommonHttp(url, pbytes, http.MethodPost)

	vmMonitoringAgentInfo := model.VmMonitoringAgentInfo{}
	if err != nil {
		fmt.Println(err)
		return &vmMonitoringAgentInfo, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode

	returnStatus := model.WebStatus{}
	if respStatus != 200 && respStatus != 201 { // 호출은 정상이나, 가져온 결과값이 200, 201아닌 경우 message에 담겨있는 것을 WebStatus에 set
		errorInfo := model.ErrorInfo{}
		json.NewDecoder(respBody).Decode(&errorInfo)
		fmt.Println("respStatus != 200 reason ", errorInfo)
		returnStatus.Message = errorInfo.Message
	} else {
		json.NewDecoder(respBody).Decode(&vmMonitoringAgentInfo)
		fmt.Println(vmMonitoringAgentInfo)
	}
	returnStatus.StatusCode = respStatus

	return &vmMonitoringAgentInfo, returnStatus
}

// Get Monitoring Data
func GetVmMonitoringInfoData(nameSpaceID string, mcisID string, metric string) (*model.VmMonitoringAgentInfo, model.WebStatus) {

	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/monitoring/mcis/" + mcisID + "/metric/" + metric
	// /ns/{nsId}/monitoring/mcis/{mcisId}/metric/{metric}

	resp, err := util.CommonHttpWithoutParam(url, http.MethodGet)

	// defer body.Close()
	vmMonitoringAgentInfo := model.VmMonitoringAgentInfo{}
	if err != nil {
		fmt.Println(err)
		return &vmMonitoringAgentInfo, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}
	// util.DisplayResponse(resp) // 수신내용 확인

	respBody := resp.Body
	respStatus := resp.StatusCode

	json.NewDecoder(respBody).Decode(&vmMonitoringAgentInfo)
	fmt.Println(vmMonitoringAgentInfo)

	return &vmMonitoringAgentInfo, model.WebStatus{StatusCode: respStatus}
}

// VM monitoring
func GetVmMonitoring(vmMonitoring *model.VmMonitoring) (*model.VmMonitoringInfo, model.WebStatus) {
	////var url = DragonFlyURL+"/ns/"+NAMESPACE+
	//"/mcis/"+mcis_id+"/vm/"+vm_id+"/metric/"+metric+"/info?periodType="+periodType+"&statisticsCriteria="+statisticsCriteria+"&duration="+duration;
	urlParam := "periodType=" + vmMonitoring.PeriodType + "&statisticsCriteria=" + vmMonitoring.StatisticsCriteria + "&duration=" + vmMonitoring.Duration
	url := util.DRAGONFLY + "/ns/" + vmMonitoring.NameSpaceID + "/mcis/" + vmMonitoring.McisID + "/vm/" + vmMonitoring.VmID + "/metric/" + vmMonitoring.Metric + "/info?" + urlParam

	// resp, err := util.CommonHttp(url, nil, http.MethodGet)
	resp, err := util.CommonHttpWithoutParam(url, http.MethodGet)

	// defer body.Close()
	vmMonitoringInfo := model.VmMonitoringInfo{}
	if err != nil {
		fmt.Println(err)
		return &vmMonitoringInfo, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}
	// util.DisplayResponse(resp) // 수신내용 확인

	respBody := resp.Body
	respStatus := resp.StatusCode

	json.NewDecoder(respBody).Decode(&vmMonitoringInfo)
	fmt.Println(vmMonitoringInfo)

	return &vmMonitoringInfo, model.WebStatus{StatusCode: respStatus}
}

// 멀티 클라우드 인프라 VM 온디맨드 모니터링 정보 조회
func GetMcisMonitoringMetricInfo(vmMonitoring *model.VmMonitoring) (*model.VmMonitoringInfo, model.WebStatus) {

	url := util.DRAGONFLY + "/ns/" + vmMonitoring.NameSpaceID + "/mcis/" + vmMonitoring.McisID + "/vm/" + vmMonitoring.VmID + "/agent_ip/" + vmMonitoring.AgentIP + "/mcis_metric/" + vmMonitoring.MetricName + "/mcis-monitoring-info"
	//{{ip}}:{{port}}/dragonfly/ns/:ns_id/mcis/:mcis_id/vm/:vm_id/agent_ip/:agent_ip/mcis_metric/:metric_name/mcis-monitoring-info
	resp, err := util.CommonHttpWithoutParam(url, http.MethodGet)

	// defer body.Close()
	vmMonitoringInfo := model.VmMonitoringInfo{}
	if err != nil {
		fmt.Println(err)
		return &vmMonitoringInfo, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}
	// util.DisplayResponse(resp) // 수신내용 확인

	respBody := resp.Body
	respStatus := resp.StatusCode

	json.NewDecoder(respBody).Decode(&vmMonitoringInfo)
	fmt.Println(vmMonitoringInfo)

	return &vmMonitoringInfo, model.WebStatus{StatusCode: respStatus}
}

// 모니터링 정책 조회
func GetMonitoringConfig() (*model.MonitoringConfig, model.WebStatus) {

	url := util.DRAGONFLY + "/config"
	// {{ip}}:{{port}}/dragonfly/config
	resp, err := util.CommonHttpWithoutParam(url, http.MethodGet)

	// defer body.Close()
	monitoringConfig := model.MonitoringConfig{}
	if err != nil {
		fmt.Println(err)
		return &monitoringConfig, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}
	// util.DisplayResponse(resp) // 수신내용 확인

	respBody := resp.Body
	respStatus := resp.StatusCode

	json.NewDecoder(respBody).Decode(&monitoringConfig)
	fmt.Println(monitoringConfig)

	return &monitoringConfig, model.WebStatus{StatusCode: respStatus}
}

// 모니터링 정책 설정
func PutMonigoringConfig(monitoringConfig *model.MonitoringConfig) (*model.MonitoringConfig, model.WebStatus) {
	url := util.DRAGONFLY + "/config"

	fmt.Println("UpdateMonigoringConfig : ")

	pbytes, _ := json.Marshal(monitoringConfig)
	fmt.Println(string(pbytes))
	resp, err := util.CommonHttp(url, pbytes, http.MethodPut)
	resultMonitoringConfig := model.MonitoringConfig{}
	if err != nil {
		log.Println("-----")
		fmt.Println(err)
		log.Println("-----1111")
		fmt.Println(err.Error())
		log.Println("-----222")
		return &resultMonitoringConfig, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode
	log.Println("respStatusCode = ", resp.StatusCode)
	log.Println("respStatus = ", resp.Status)
	if respStatus != 200 && respStatus != 201 {
		errorInfo := model.ErrorInfo{}
		json.NewDecoder(respBody).Decode(&errorInfo)
		fmt.Println(errorInfo)
		return nil, model.WebStatus{StatusCode: 500, Message: errorInfo.Message}
	}

	// 응답에 생성한 객체값이 옴
	json.NewDecoder(respBody).Decode(&resultMonitoringConfig)
	fmt.Println(resultMonitoringConfig)
	// return respBody, respStatusCode
	return &resultMonitoringConfig, model.WebStatus{StatusCode: respStatus}
}

// 모니터링 정책 초기화
func ResetMonigoringConfig(monitoringConfig *model.MonitoringConfig) (*model.MonitoringConfig, model.WebStatus) {
	url := util.DRAGONFLY + "/config/reset"

	fmt.Println("ResetMonigoringConfig : ", url)

	resp, err := util.CommonHttp(url, nil, http.MethodPut)
	resultMonitoringConfig := model.MonitoringConfig{}
	if err != nil {
		log.Println("-----")
		fmt.Println(err)
		log.Println("-----1111")
		fmt.Println(err.Error())
		log.Println("-----222")
		return &resultMonitoringConfig, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode
	log.Println("respStatusCode = ", resp.StatusCode)
	log.Println("respStatus = ", resp.Status)
	if respStatus != 200 && respStatus != 201 {
		errorInfo := model.ErrorInfo{}
		json.NewDecoder(respBody).Decode(&errorInfo)
		fmt.Println(errorInfo)
		return nil, model.WebStatus{StatusCode: 500, Message: errorInfo.Message}
	}

	// 응답에 생성한 객체값이 옴
	json.NewDecoder(respBody).Decode(&resultMonitoringConfig)
	fmt.Println(resultMonitoringConfig)
	// return respBody, respStatusCode
	return &resultMonitoringConfig, model.WebStatus{StatusCode: respStatus}
}

// Install agent to vm
// 모니터링 에이전트 설치 : 위에 RegMonitoringAgentInVm 와 뭐가 다른거지?
func InstallAgentToVm(nameSpaceID string, vmMonitoringInstallReg *model.VmMonitoringInstallReg) (*model.VmMonitoringInstallReg, model.WebStatus) {

	url := util.DRAGONFLY + "/agent/install/"
	// {{ip}}:{{port}}/dragonfly/agent/install

	pbytes, _ := json.Marshal(vmMonitoringInstallReg)
	resp, err := util.CommonHttp(url, pbytes, http.MethodPost)

	returnVmMonitoringInstallReg := model.VmMonitoringInstallReg{}
	returnStatus := model.WebStatus{}

	respBody := resp.Body
	respStatus := resp.StatusCode

	if err != nil {
		fmt.Println(err)
		return &returnVmMonitoringInstallReg, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	if respStatus != 200 && respStatus != 201 { // 호출은 정상이나, 가져온 결과값이 200, 201아닌 경우 message에 담겨있는 것을 WebStatus에 set
		errorInfo := model.ErrorInfo{}
		json.NewDecoder(respBody).Decode(&errorInfo)
		fmt.Println("respStatus != 200 reason ", errorInfo)
		returnStatus.Message = errorInfo.Message
	} else {
		json.NewDecoder(respBody).Decode(&returnVmMonitoringInstallReg)
		fmt.Println(returnVmMonitoringInstallReg)
	}
	returnStatus.StatusCode = respStatus

	return &returnVmMonitoringInstallReg, returnStatus
}

// 모니터링 에이전트 제거
// Uninstall agent to vm
func UnInstallAgentToVm(nameSpaceID string, vmMonitoringInstallReg *model.VmMonitoringInstallReg) (*model.VmMonitoringInstallReg, model.WebStatus) {

	url := util.DRAGONFLY + "/agent/uninstall/"
	// {{ip}}:{{port}}/dragonfly/agent/uninstall

	pbytes, _ := json.Marshal(vmMonitoringInstallReg)
	resp, err := util.CommonHttp(url, pbytes, http.MethodPost)

	returnVmMonitoringInstallReg := model.VmMonitoringInstallReg{}
	returnStatus := model.WebStatus{}

	respBody := resp.Body
	respStatus := resp.StatusCode

	if err != nil {
		fmt.Println(err)
		return &returnVmMonitoringInstallReg, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	if respStatus != 200 && respStatus != 201 { // 호출은 정상이나, 가져온 결과값이 200, 201아닌 경우 message에 담겨있는 것을 WebStatus에 set
		errorInfo := model.ErrorInfo{}
		json.NewDecoder(respBody).Decode(&errorInfo)
		fmt.Println("respStatus != 200 reason ", errorInfo)
		returnStatus.Message = errorInfo.Message
	} else {
		json.NewDecoder(respBody).Decode(&returnVmMonitoringInstallReg)
		fmt.Println(returnVmMonitoringInstallReg)
	}
	returnStatus.StatusCode = respStatus

	return &returnVmMonitoringInstallReg, returnStatus
}

// 알람 목록 조회
// List monitoring alert
func GetMonitoringAlertList() ([]model.VmMonitoringAlertInfo, model.WebStatus) {

	url := util.DRAGONFLY + "/alert/tasks"
	resp, err := util.CommonHttp(url, nil, http.MethodGet)
	// {{ip}}:{{port}}/dragonfly/alert/tasks

	if err != nil {
		fmt.Println(err)
		return nil, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode

	vmMonitoringAlertInfoList := map[string][]model.VmMonitoringAlertInfo{}
	json.NewDecoder(respBody).Decode(&vmMonitoringAlertInfoList)

	return vmMonitoringAlertInfoList["mcis"], model.WebStatus{StatusCode: respStatus}
}

// 알람  조회
// monitoring alert
func GetMonitoringAlertData(taskName string) (model.VmMonitoringAlertInfo, model.WebStatus) {

	url := util.DRAGONFLY + "/alert/task/" + taskName
	resp, err := util.CommonHttp(url, nil, http.MethodGet)
	// {{ip}}:{{port}}/dragonfly/alert/task/:task_name

	vmMonitoringAlertInfo := model.VmMonitoringAlertInfo{}
	if err != nil {
		fmt.Println(err)
		return vmMonitoringAlertInfo, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode

	json.NewDecoder(respBody).Decode(&vmMonitoringAlertInfo)

	return vmMonitoringAlertInfo, model.WebStatus{StatusCode: respStatus}
}

// 알람 생성
// Create Monitoring Alert
func RegMonitoringAlert(vmMonitoringAlertInfo *model.VmMonitoringAlertInfo) (*model.VmMonitoringAlertInfo, model.WebStatus) {
	fmt.Println("RegMonitoringAlert ************ : ")

	url := util.DRAGONFLY + "/alert/task"
	// {{ip}}:{{port}}/dragonfly/alert/task

	pbytes, _ := json.Marshal(vmMonitoringAlertInfo)
	fmt.Println(string(pbytes))
	resp, err := util.CommonHttp(url, pbytes, http.MethodPost)

	resultVmMonitoringAlertInfo := model.VmMonitoringAlertInfo{}
	if err != nil {
		fmt.Println(err)
		return &resultVmMonitoringAlertInfo, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode

	returnStatus := model.WebStatus{}
	if respStatus != 200 && respStatus != 201 { // 호출은 정상이나, 가져온 결과값이 200, 201아닌 경우 message에 담겨있는 것을 WebStatus에 set
		errorInfo := model.ErrorInfo{}
		json.NewDecoder(respBody).Decode(&errorInfo)
		fmt.Println("respStatus != 200 reason ", errorInfo)
		returnStatus.Message = errorInfo.Message
	} else {
		json.NewDecoder(respBody).Decode(&resultVmMonitoringAlertInfo)
		fmt.Println(resultVmMonitoringAlertInfo)
	}
	returnStatus.StatusCode = respStatus

	return &resultVmMonitoringAlertInfo, returnStatus
}

// 알람 수정
// Update Monitoring Alert
func PutMonitoringAlert(taskName string, vmMonitoringAlertInfo *model.VmMonitoringAlertInfo) (*model.VmMonitoringAlertInfo, model.WebStatus) {
	fmt.Println("PutMonitoringAlert ************ : ")

	url := util.DRAGONFLY + "/alert/task/" + taskName
	// {{ip}}:{{port}}/dragonfly/alert/task

	pbytes, _ := json.Marshal(vmMonitoringAlertInfo)
	fmt.Println(string(pbytes))
	resp, err := util.CommonHttp(url, pbytes, http.MethodPut)

	resultVmMonitoringAlertInfo := model.VmMonitoringAlertInfo{}
	if err != nil {
		fmt.Println(err)
		return &resultVmMonitoringAlertInfo, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode

	returnStatus := model.WebStatus{}
	if respStatus != 200 && respStatus != 201 { // 호출은 정상이나, 가져온 결과값이 200, 201아닌 경우 message에 담겨있는 것을 WebStatus에 set
		errorInfo := model.ErrorInfo{}
		json.NewDecoder(respBody).Decode(&errorInfo)
		fmt.Println("respStatus != 200 reason ", errorInfo)
		returnStatus.Message = errorInfo.Message
	} else {
		json.NewDecoder(respBody).Decode(&resultVmMonitoringAlertInfo)
		fmt.Println(resultVmMonitoringAlertInfo)
	}
	returnStatus.StatusCode = respStatus

	return &resultVmMonitoringAlertInfo, returnStatus
}

// 알람 제거
// Delete Monitoring Alert
func DelMonitoringAlert(taskName string) (io.ReadCloser, model.WebStatus) {
	// buff := bytes.NewBuffer(pbytes)
	url := util.DRAGONFLY + "/alert/task/" + taskName
	// /ns/{nsId}/mcis/{mcisId}
	fmt.Println("url : ", url)

	if taskName == "" {
		return nil, model.WebStatus{StatusCode: 500, Message: "TaskName is required"}
	}

	// 경로안에 parameter가 있어 추가 param없이 호출 함.
	resp, err := util.CommonHttp(url, nil, http.MethodDelete)
	if err != nil {
		fmt.Println(err)
		return nil, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}
	// return body, err
	respBody := resp.Body
	respStatus := resp.StatusCode
	return respBody, model.WebStatus{StatusCode: respStatus}
}

// 알람 이벤트 핸들러 조회
// Get monitoring alert event-handler
// type : 이벤트 핸들러 유형 ( "slack" | "smtp" )
// name : slackHandler(EventHandlerName)
func GetMonitoringAlertEventHandlerData(eventHandlerType string, eventName string) (model.VmMonitoringAlertInfo, model.WebStatus) {

	url := util.DRAGONFLY + "/alert/eventhandler/type/" + eventHandlerType + "/event/" + eventName
	//{{ip}}:{{port}}/dragonfly/alert/eventhandler/type/:type/event/:name
	resp, err := util.CommonHttp(url, nil, http.MethodGet)

	vmMonitoringAlertInfo := model.VmMonitoringAlertInfo{}
	if err != nil {
		fmt.Println(err)
		return vmMonitoringAlertInfo, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode

	json.NewDecoder(respBody).Decode(&vmMonitoringAlertInfo)

	return vmMonitoringAlertInfo, model.WebStatus{StatusCode: respStatus}
}

// 알람 이벤트 핸들러 생성
// Create monitoring alert event-handler
func RegMonitoringAlertEventHandler(vmMonitoringAlertEventHandlerInfo *model.VmMonitoringAlertEventHandlerInfo) (*model.VmMonitoringAlertEventHandlerInfo, model.WebStatus) {
	fmt.Println("RegMonitoringAlertEventHandler ************ : ")

	url := util.DRAGONFLY + "/alert/eventhandler"
	// {{ip}}:{{port}}/dragonfly/alert/eventhandler

	pbytes, _ := json.Marshal(vmMonitoringAlertEventHandlerInfo)
	fmt.Println(string(pbytes))
	resp, err := util.CommonHttp(url, pbytes, http.MethodPost)

	resultVmMonitoringAlertEventHandlerInfo := model.VmMonitoringAlertEventHandlerInfo{}
	if err != nil {
		fmt.Println(err)
		return &resultVmMonitoringAlertEventHandlerInfo, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode

	returnStatus := model.WebStatus{}
	if respStatus != 200 && respStatus != 201 { // 호출은 정상이나, 가져온 결과값이 200, 201아닌 경우 message에 담겨있는 것을 WebStatus에 set
		errorInfo := model.ErrorInfo{}
		json.NewDecoder(respBody).Decode(&errorInfo)
		fmt.Println("respStatus != 200 reason ", errorInfo)
		returnStatus.Message = errorInfo.Message
	} else {
		json.NewDecoder(respBody).Decode(&resultVmMonitoringAlertEventHandlerInfo)
		fmt.Println(resultVmMonitoringAlertEventHandlerInfo)
	}
	returnStatus.StatusCode = respStatus

	return &resultVmMonitoringAlertEventHandlerInfo, returnStatus
}

// 알람 이벤트 핸들러 수정( handlerType=slack)
func PutMonitoringAlertEventHandlerSlack(eventHandlerType string, eventName string, vmMonitoringAlertEventHandlerInfo *model.VmMonitoringAlertEventHandlerInfo) (*model.VmMonitoringAlertEventHandlerInfo, model.WebStatus) {
	fmt.Println("PutMonitoringAlertEventHandler ************ : ")

	url := util.DRAGONFLY + "/alert/eventhandler/type/" + eventHandlerType + "/event/" + eventName
	// {{ip}}:{{port}}/dragonfly/alert/eventhandler/type/:type/event/:name

	pbytes, _ := json.Marshal(vmMonitoringAlertEventHandlerInfo)
	fmt.Println(string(pbytes))
	resp, err := util.CommonHttp(url, pbytes, http.MethodPut)

	resultVmMonitoringAlertEventHandlerInfo := model.VmMonitoringAlertEventHandlerInfo{}
	if err != nil {
		fmt.Println(err)
		return &resultVmMonitoringAlertEventHandlerInfo, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode

	returnStatus := model.WebStatus{}
	if respStatus != 200 && respStatus != 201 { // 호출은 정상이나, 가져온 결과값이 200, 201아닌 경우 message에 담겨있는 것을 WebStatus에 set
		errorInfo := model.ErrorInfo{}
		json.NewDecoder(respBody).Decode(&errorInfo)
		fmt.Println("respStatus != 200 reason ", errorInfo)
		returnStatus.Message = errorInfo.Message
	} else {
		json.NewDecoder(respBody).Decode(&resultVmMonitoringAlertEventHandlerInfo)
		fmt.Println(resultVmMonitoringAlertEventHandlerInfo)
	}
	returnStatus.StatusCode = respStatus

	return &resultVmMonitoringAlertEventHandlerInfo, returnStatus
}

// 알람 이벤트 핸들러 수정( handlerType=smtp)
func PutMonitoringAlertEventHandlerSmtp(eventHandlerType string, eventName string, vmMonitoringAlertEventHandlerInfo *model.VmMonitoringAlertEventHandlerInfo) (*model.VmMonitoringAlertEventHandlerInfo, model.WebStatus) {
	fmt.Println("PutMonitoringAlertEventHandlerSmtp ************ : ")

	url := util.DRAGONFLY + "/alert/eventhandler/type/" + eventHandlerType + "/event/" + eventName
	// {{ip}}:{{port}}/dragonfly/alert/eventhandler/type/:type/event/:name

	pbytes, _ := json.Marshal(vmMonitoringAlertEventHandlerInfo)
	fmt.Println(string(pbytes))
	resp, err := util.CommonHttp(url, pbytes, http.MethodPut)

	resultVmMonitoringAlertEventHandlerInfo := model.VmMonitoringAlertEventHandlerInfo{}
	if err != nil {
		fmt.Println(err)
		return &resultVmMonitoringAlertEventHandlerInfo, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode

	returnStatus := model.WebStatus{}
	if respStatus != 200 && respStatus != 201 { // 호출은 정상이나, 가져온 결과값이 200, 201아닌 경우 message에 담겨있는 것을 WebStatus에 set
		errorInfo := model.ErrorInfo{}
		json.NewDecoder(respBody).Decode(&errorInfo)
		fmt.Println("respStatus != 200 reason ", errorInfo)
		returnStatus.Message = errorInfo.Message
	} else {
		json.NewDecoder(respBody).Decode(&resultVmMonitoringAlertEventHandlerInfo)
		fmt.Println(resultVmMonitoringAlertEventHandlerInfo)
	}
	returnStatus.StatusCode = respStatus

	return &resultVmMonitoringAlertEventHandlerInfo, returnStatus
}

// 알람 제거
// Delete monitoring alert event-handler
func DelMonitoringAlertEventHandler(eventHandlerType string, eventHandlerName string) (io.ReadCloser, model.WebStatus) {

	url := util.DRAGONFLY + "/alert/eventhandler/type/" + eventHandlerType + "/event/" + eventHandlerName
	// {{ip}}:{{port}}/dragonfly/alert/eventhandler/type/:type/event/:name
	fmt.Println("url : ", url)

	if eventHandlerType == "" {
		return nil, model.WebStatus{StatusCode: 500, Message: "eventHandlerType is required"}
	}
	if eventHandlerName == "" {
		return nil, model.WebStatus{StatusCode: 500, Message: "eventHandlerName is required"}
	}

	// 경로안에 parameter가 있어 추가 param없이 호출 함.
	resp, err := util.CommonHttp(url, nil, http.MethodDelete)
	if err != nil {
		fmt.Println(err)
		return nil, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}
	// return body, err
	respBody := resp.Body
	respStatus := resp.StatusCode
	return respBody, model.WebStatus{StatusCode: respStatus}
}

// 알람 로그 정보 목록 조회
// List monitoring alert event
func GetMonitoringAlertLogList() ([]model.VmMonitoringAlertInfo, model.WebStatus) {

	url := util.DRAGONFLY + "/alert/task"
	resp, err := util.CommonHttp(url, nil, http.MethodGet)
	// {{ip}}:{{port}}/dragonfly/alert/task/:task_name/events?level=warning

	if err != nil {
		fmt.Println(err)
		return nil, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode

	vmMonitoringAlertInfoList := map[string][]model.VmMonitoringAlertInfo{}
	json.NewDecoder(respBody).Decode(&vmMonitoringAlertInfoList)

	return vmMonitoringAlertInfoList["mcis"], model.WebStatus{StatusCode: respStatus}
}
