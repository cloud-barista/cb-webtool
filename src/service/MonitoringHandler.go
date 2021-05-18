package service

import (
	"encoding/json"
	"fmt"
	"io"

	// "io/ioutil"
	"log"
	"net/http"
	"strconv"

	// "os"
	model "github.com/cloud-barista/cb-webtool/src/model"
	// "github.com/cloud-barista/cb-webtool/src/model/spider"
	"github.com/cloud-barista/cb-webtool/src/model/dragonfly"
	"github.com/cloud-barista/cb-webtool/src/model/tumblebug"

	util "github.com/cloud-barista/cb-webtool/src/util"
)

// VM 에 모니터링 Agent 설치
///ns/{nsId}/monitoring/install/mcis/{mcisId}
func RegMonitoringAgentInVm(nameSpaceID string, mcisID string, vmMonitoringAgentReg *tumblebug.VmMonitoringAgentReg) (*tumblebug.VmMonitoringAgentInfo, model.WebStatus) {
	fmt.Println("RegMonitoringAgentInVm ************ : ")
	var originalUrl = "/ns/{nsId}/monitoring/install/mcis/{mcisId}"

	var paramMapper = make(map[string]string)
	paramMapper["{nsId}"] = nameSpaceID
	paramMapper["{mcisId}"] = mcisID
	urlParam := util.MappingUrlParameter(originalUrl, paramMapper)

	url := util.TUMBLEBUG + urlParam
	// url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/monitoring/install/mcis/"

	pbytes, _ := json.Marshal(vmMonitoringAgentReg)
	fmt.Println(string(pbytes))
	resp, err := util.CommonHttp(url, pbytes, http.MethodPost)

	vmMonitoringAgentInfo := tumblebug.VmMonitoringAgentInfo{}
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
func GetVmMonitoringInfoData(nameSpaceID string, mcisID string, metric string) (*tumblebug.VmMonitoringAgentInfo, model.WebStatus) {
	var originalUrl = "/ns/{nsId}/monitoring/mcis/{mcisId}/metric/{metric}"

	var paramMapper = make(map[string]string)
	paramMapper["{nsId}"] = nameSpaceID
	paramMapper["{mcisId}"] = mcisID
	paramMapper["{metric}"] = metric
	urlParam := util.MappingUrlParameter(originalUrl, paramMapper)

	url := util.TUMBLEBUG + urlParam
	// url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/monitoring/mcis/" + mcisID + "/metric/" + metric

	resp, err := util.CommonHttpWithoutParam(url, http.MethodGet)

	// defer body.Close()
	vmMonitoringAgentInfo := tumblebug.VmMonitoringAgentInfo{}
	if err != nil {
		fmt.Println(err)
		return &vmMonitoringAgentInfo, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode

	json.NewDecoder(respBody).Decode(&vmMonitoringAgentInfo)
	fmt.Println(vmMonitoringAgentInfo)

	return &vmMonitoringAgentInfo, model.WebStatus{StatusCode: respStatus}
}

// VM monitoring
// Get vm monitoring info
// 멀티 클라우드 인프라 VM 모니터링 정보 조회
func GetVmMonitoring(vmMonitoring *dragonfly.VmMonitoring) (*dragonfly.VmMonitoringInfo, model.WebStatus) {
	nameSpaceID := vmMonitoring.NameSpaceID
	mcisID := vmMonitoring.McisID
	vmID := vmMonitoring.VmID
	metric := vmMonitoring.Metric
	periodType := vmMonitoring.PeriodType
	statisticsCriteria := vmMonitoring.StatisticsCriteria
	duration := vmMonitoring.Duration

	var originalUrl = "/ns/:ns_id/mcis/:mcis_id/vm/:vm_id/metric/:metric_name/info?periodType={periodType}&statisticsCriteria={statisticsCriteria}&duration={duration}"
	//{{ip}}:{{port}}/dragonfly/ns/:ns_id/mcis/:mcis_id/vm/:vm_id/metric/:metric_name/info?periodType=m&statisticsCriteria=last&duration=5m
	var paramMapper = make(map[string]string)
	paramMapper[":ns_id"] = nameSpaceID
	paramMapper[":mcis_id"] = mcisID
	paramMapper[":vm_id"] = vmID
	paramMapper["{metric}"] = metric
	paramMapper["{periodType}"] = periodType
	paramMapper["{statisticsCriteria}"] = statisticsCriteria
	paramMapper["{duration}"] = duration
	urlParam := util.MappingUrlParameter(originalUrl, paramMapper)

	url := util.DRAGONFLY + urlParam

	//"/mcis/"+mcis_id+"/vm/"+vm_id+"/metric/"+metric+"/info?periodType="+periodType+"&statisticsCriteria="+statisticsCriteria+"&duration="+duration;
	// urlParam := "periodType=" + vmMonitoring.PeriodType + "&statisticsCriteria=" + vmMonitoring.StatisticsCriteria + "&duration=" + vmMonitoring.Duration
	// url := util.DRAGONFLY + "/ns/" + vmMonitoring.NameSpaceID + "/mcis/" + vmMonitoring.McisID + "/vm/" + vmMonitoring.VmID + "/metric/" + vmMonitoring.Metric + "/info?" + urlParam

	// resp, err := util.CommonHttp(url, nil, http.MethodGet)
	resp, err := util.CommonHttpWithoutParam(url, http.MethodGet)

	// defer body.Close()
	vmMonitoringInfo := dragonfly.VmMonitoringInfo{}
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
// Get MCIS on-demand monitoring metric info
func GetMcisOnDemandMonitoringMetricInfo(agentIp string, metricName string, vmMonitoring *dragonfly.VmMonitoring) (*dragonfly.McisMonitoringOnDemandInfo, model.WebStatus) {
	nameSpaceID := vmMonitoring.NameSpaceID
	mcisID := vmMonitoring.McisID
	vmID := vmMonitoring.VmID
	// agentIp := vmMonitoring.AgentIp
	// metricName := vmMonitoring.MetricName

	var originalUrl = "/ns/:ns_id/mcis/:mcis_id/vm/:vm_id/agent_ip/:agent_ip/mcis_metric/:metric_name/mcis-monitoring-info"
	//{{ip}}:{{port}}/dragonfly/ns/:ns_id/mcis/:mcis_id/vm/:vm_id/agent_ip/:agent_ip/mcis_metric/:metric_name/mcis-monitoring-info
	var paramMapper = make(map[string]string)
	paramMapper[":ns_id"] = nameSpaceID
	paramMapper[":mcis_id"] = mcisID
	paramMapper[":vm_id"] = vmID
	paramMapper[":agent_ip"] = agentIp       // 에이전트 아이피
	paramMapper[":metric_name"] = metricName // 메트릭 정보 ( "InitDB" | "ResetDB" | "CpuM" | "CpuS" | "MemR" | "MemW" | "FioW" | "FioR" | "DBW" | DBR" | "Rtt" | "Mrtt" )
	urlParam := util.MappingUrlParameter(originalUrl, paramMapper)

	url := util.DRAGONFLY + urlParam

	// url := util.DRAGONFLY + "/ns/" + vmMonitoring.NameSpaceID + "/mcis/" + vmMonitoring.McisID + "/vm/" + vmMonitoring.VmID + "/agent_ip/" + vmMonitoring.AgentIP + "/mcis_metric/" + vmMonitoring.MetricName + "/mcis-monitoring-info"
	// url := util.DRAGONFLY + "/ns/" + vmMonitoring.NameSpaceID + "/mcis/" + vmMonitoring.McisID + "/vm/" + vmMonitoring.VmID // TODO : 객체에 parameter추가해야 함

	resp, err := util.CommonHttpWithoutParam(url, http.MethodGet)

	// defer body.Close()
	mcisMonitoringInfo := dragonfly.McisMonitoringOnDemandInfo{}
	if err != nil {
		fmt.Println(err)
		return &mcisMonitoringInfo, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}
	// util.DisplayResponse(resp) // 수신내용 확인

	respBody := resp.Body
	respStatus := resp.StatusCode

	json.NewDecoder(respBody).Decode(&mcisMonitoringInfo)
	fmt.Println(mcisMonitoringInfo)

	return &mcisMonitoringInfo, model.WebStatus{StatusCode: respStatus}
}

// 멀티 클라우드 인프라 VM 온디맨드 모니터링 정보 조회
// Get vm on-demand monitoring metric info
func GetVmOnDemandMonitoringMetricInfo(agentIp string, metricName string, vmMonitoring *dragonfly.VmMonitoring) (*dragonfly.VmMonitoringOnDemandInfo, model.WebStatus) {
	nameSpaceID := vmMonitoring.NameSpaceID
	mcisID := vmMonitoring.McisID
	vmID := vmMonitoring.VmID
	// agentIp := vmMonitoring.AgentIp
	// metricName := vmMonitoring.MetricName

	var originalUrl = "/ns/:ns_id/mcis/:mcis_id/vm/:vm_id/agent_ip/:agent_ip/metric/:metric_name/ondemand-monitoring-info"
	// {{ip}}:{{port}}/dragonfly/ns/:ns_id/mcis/:mcis_id/vm/:vm_id/agent_ip/:agent_ip/metric/:metric_name/ondemand-monitoring-info
	var paramMapper = make(map[string]string)
	paramMapper[":ns_id"] = nameSpaceID
	paramMapper[":mcis_id"] = mcisID
	paramMapper[":vm_id"] = vmID
	paramMapper[":agent_ip"] = agentIp
	paramMapper[":metric_name"] = metricName
	urlParam := util.MappingUrlParameter(originalUrl, paramMapper)

	url := util.DRAGONFLY + urlParam

	// url := util.DRAGONFLY + "/ns/" + vmMonitoring.NameSpaceID + "/mcis/" + vmMonitoring.McisID + "/vm/" + vmMonitoring.VmID + "/agent_ip/" + vmMonitoring.AgentIP + "/mcis_metric/" + vmMonitoring.MetricName + "/mcis-monitoring-info"
	// url := util.DRAGONFLY + "/ns/" + vmMonitoring.NameSpaceID + "/mcis/" + vmMonitoring.McisID + "/vm/" + vmMonitoring.VmID // TODO : 객체에 parameter추가해야 함

	resp, err := util.CommonHttpWithoutParam(url, http.MethodGet)

	// defer body.Close()
	vmMonitoringInfo := dragonfly.VmMonitoringOnDemandInfo{}
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
// Get monitoring config
func GetMonitoringConfig() (*dragonfly.MonitoringConfig, model.WebStatus) {
	var originalUrl = "/config"
	//{{ip}}:{{port}}/dragonfly/config
	urlParam := util.MappingUrlParameter(originalUrl, nil)

	url := util.DRAGONFLY + urlParam
	// url := util.DRAGONFLY + "/config"
	resp, err := util.CommonHttpWithoutParam(url, http.MethodGet)

	// defer body.Close()
	monitoringConfig := dragonfly.MonitoringConfig{}
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
func PutMonigoringConfig(monitoringConfigReg *dragonfly.MonitoringConfigReg) (*dragonfly.MonitoringConfig, model.WebStatus) {
	var originalUrl = "/config"
	//{{ip}}:{{port}}/dragonfly/config
	urlParam := util.MappingUrlParameter(originalUrl, nil)

	url := util.DRAGONFLY + urlParam
	// url := util.DRAGONFLY + "/config"

	fmt.Println("Update MonigoringConfigReg : ", url)

	// pbytes, _ := json.Marshal(monitoringConfig)
	// fmt.Println(string(pbytes))
	fmt.Println(monitoringConfigReg)
	//urlValues := util.Struct2MapString(monitoringConfigReg)
	urlValues := map[string]string{}
	urlValues["agent_interval"] = strconv.Itoa(monitoringConfigReg.AgentInterval)
	urlValues["collector_interval"] = strconv.Itoa(monitoringConfigReg.CollectorInterval)
	urlValues["max_host_count"] = strconv.Itoa(monitoringConfigReg.MaxHostCount)
	// urlValues := util.Struct2MapString(monitoringConfigReg)  : TODO : struct -> map으로 변환하도록
	//urlValues := map[string]int{}
	fmt.Println(urlValues)
	resp, err := util.CommonHttpFormData(url, urlValues, http.MethodPut)
	// resp, err := util.CommonHttp(url, pbytes, http.MethodPut)
	resultMonitoringConfig := dragonfly.MonitoringConfig{}
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
func ResetMonigoringConfig(monitoringConfig *dragonfly.MonitoringConfig) (*dragonfly.MonitoringConfig, model.WebStatus) {
	var originalUrl = "/config/reset"
	//{{ip}}:{{port}}/dragonfly/config/reset
	urlParam := util.MappingUrlParameter(originalUrl, nil)

	url := util.DRAGONFLY + urlParam
	// url := util.DRAGONFLY + "/config/reset"

	resp, err := util.CommonHttp(url, nil, http.MethodPut)
	resultMonitoringConfig := dragonfly.MonitoringConfig{}
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
func InstallAgentToVm(nameSpaceID string, vmMonitoringInstallReg *dragonfly.VmMonitoringInstallReg) (*dragonfly.VmMonitoringInstallReg, model.WebStatus) {
	var originalUrl = "/agent/install"
	//{{ip}}:{{port}}/dragonfly/agent/install
	urlParam := util.MappingUrlParameter(originalUrl, nil)

	url := util.DRAGONFLY + urlParam
	// url := util.DRAGONFLY + "/agent/install/"

	pbytes, _ := json.Marshal(vmMonitoringInstallReg)
	resp, err := util.CommonHttp(url, pbytes, http.MethodPost)

	returnVmMonitoringInstallReg := dragonfly.VmMonitoringInstallReg{}
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
func UnInstallAgentToVm(nameSpaceID string, vmMonitoringInstallReg *dragonfly.VmMonitoringInstallReg) (*dragonfly.VmMonitoringInstallReg, model.WebStatus) {
	var originalUrl = "/agent/uninstall"
	//{{ip}}:{{port}}/dragonfly/agent/uninstall
	urlParam := util.MappingUrlParameter(originalUrl, nil)

	url := util.DRAGONFLY + urlParam
	// url := util.DRAGONFLY + "/agent/uninstall/"

	pbytes, _ := json.Marshal(vmMonitoringInstallReg)
	resp, err := util.CommonHttp(url, pbytes, http.MethodPost)

	returnVmMonitoringInstallReg := dragonfly.VmMonitoringInstallReg{}
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
func GetMonitoringAlertList() ([]dragonfly.VmMonitoringAlertInfo, model.WebStatus) {
	fmt.Print("#########GetMonitoringAlertList############")
	var originalUrl = "/alert/tasks"
	// {{ip}}:{{port}}/dragonfly/alert/tasks
	urlParam := util.MappingUrlParameter(originalUrl, nil)

	url := util.DRAGONFLY + urlParam
	// url := util.DRAGONFLY + "/alert/tasks"
	resp, _ := util.CommonHttp(url, nil, http.MethodGet)

	// vmMonitoringAlertInfoList := dragonfly.VmMonitoringAlertInfo{}

	// if err != nil {
	// 	fmt.Println(err)
	// 	return vmMonitoringAlertInfoList, model.WebStatus{StatusCode: 500, Message: err.Error()}
	// }

	respBody := resp.Body
	respStatus := resp.StatusCode

	vmMonitoringAlertInfoList := []dragonfly.VmMonitoringAlertInfo{}
	json.NewDecoder(respBody).Decode(&vmMonitoringAlertInfoList)

	// robots, _ := ioutil.ReadAll(resp.Body)
	// log.Println(fmt.Print(string(robots)))

	// json.NewDecoder(respBody).Decode(&vmMonitoringAlertInfoList)
	// fmt.Println(vmMonitoringAlertInfoList)

	return vmMonitoringAlertInfoList, model.WebStatus{StatusCode: respStatus}
}

// 알람  조회
// monitoring alert
func GetMonitoringAlertData(taskName string) (dragonfly.VmMonitoringAlertInfo, model.WebStatus) {
	var originalUrl = "/alert/task/:task_name"
	// {{ip}}:{{port}}/dragonfly/alert/task/:task_name
	var paramMapper = make(map[string]string)
	paramMapper[":task_name"] = taskName
	urlParam := util.MappingUrlParameter(originalUrl, paramMapper)

	url := util.DRAGONFLY + urlParam
	// url := util.DRAGONFLY + "/alert/task/" + taskName
	resp, err := util.CommonHttp(url, nil, http.MethodGet)

	vmMonitoringAlertInfo := dragonfly.VmMonitoringAlertInfo{}
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
func RegMonitoringAlert(vmMonitoringAlertInfo *dragonfly.VmMonitoringAlertInfo) (*dragonfly.VmMonitoringAlertInfo, model.WebStatus) {
	fmt.Println("RegMonitoringAlert ************ : ")
	var originalUrl = "/alert/task"
	// {{ip}}:{{port}}/dragonfly/alert/task
	urlParam := util.MappingUrlParameter(originalUrl, nil)

	url := util.DRAGONFLY + urlParam
	// url := util.DRAGONFLY + "/alert/task"

	pbytes, _ := json.Marshal(vmMonitoringAlertInfo)
	fmt.Println(string(pbytes))
	resp, err := util.CommonHttp(url, pbytes, http.MethodPost)

	resultVmMonitoringAlertInfo := dragonfly.VmMonitoringAlertInfo{}
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
func PutMonitoringAlert(taskName string, vmMonitoringAlertInfo *dragonfly.VmMonitoringAlertInfo) (*dragonfly.VmMonitoringAlertInfo, model.WebStatus) {
	fmt.Println("PutMonitoringAlert ************ : ")
	var originalUrl = "/alert/task"
	// {{ip}}:{{port}}/dragonfly/alert/task
	urlParam := util.MappingUrlParameter(originalUrl, nil)

	url := util.DRAGONFLY + urlParam
	// url := util.DRAGONFLY + "/alert/task/" + taskName

	pbytes, _ := json.Marshal(vmMonitoringAlertInfo)
	fmt.Println(string(pbytes))
	resp, err := util.CommonHttp(url, pbytes, http.MethodPut)

	resultVmMonitoringAlertInfo := dragonfly.VmMonitoringAlertInfo{}
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
	var originalUrl = "/alert/task/:task_name"
	// {{ip}}:{{port}}/dragonfly/alert/task/:task_name
	var paramMapper = make(map[string]string)
	paramMapper[":task_name"] = taskName
	urlParam := util.MappingUrlParameter(originalUrl, paramMapper)

	url := util.DRAGONFLY + urlParam
	// url := util.DRAGONFLY + "/alert/task/" + taskName

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
func GetMonitoringAlertEventHandlerData(eventHandlerType string, eventName string) (dragonfly.VmMonitoringAlertInfo, model.WebStatus) {
	var originalUrl = "/alert/eventhandler/type/:type/event/:name"
	//{{ip}}:{{port}}/dragonfly/alert/eventhandler/type/:type/event/:name
	var paramMapper = make(map[string]string)
	paramMapper[":type"] = eventHandlerType
	paramMapper[":name"] = eventName
	urlParam := util.MappingUrlParameter(originalUrl, paramMapper)

	url := util.DRAGONFLY + urlParam
	// url := util.DRAGONFLY + "/alert/eventhandler/type/" + eventHandlerType + "/event/" + eventName

	resp, err := util.CommonHttp(url, nil, http.MethodGet)

	vmMonitoringAlertInfo := dragonfly.VmMonitoringAlertInfo{}
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
func RegMonitoringAlertEventHandler(vmMonitoringAlertEventHandlerInfo *dragonfly.VmMonitoringAlertEventHandlerInfo) (*dragonfly.VmMonitoringAlertEventHandlerInfo, model.WebStatus) {
	fmt.Println("RegMonitoringAlertEventHandler ************ : ")
	var originalUrl = "/alert/eventhandler"
	// {{ip}}:{{port}}/dragonfly/alert/eventhandler
	urlParam := util.MappingUrlParameter(originalUrl, nil)

	url := util.DRAGONFLY + urlParam
	// url := util.DRAGONFLY + "/alert/eventhandler"

	pbytes, _ := json.Marshal(vmMonitoringAlertEventHandlerInfo)
	fmt.Println(string(pbytes))
	resp, err := util.CommonHttp(url, pbytes, http.MethodPost)

	resultVmMonitoringAlertEventHandlerInfo := dragonfly.VmMonitoringAlertEventHandlerInfo{}
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
func PutMonitoringAlertEventHandlerSlack(eventHandlerType string, eventName string, vmMonitoringAlertEventHandlerSlackInfo *dragonfly.EventHandlerOptionSlack) (*dragonfly.VmMonitoringAlertEventHandlerSlackInfo, model.WebStatus) {
	fmt.Println("PutMonitoringAlertEventHandler ************ : ")
	var originalUrl = "/alert/eventhandler/type/:type/event/:name"
	// {{ip}}:{{port}}/dragonfly/alert/eventhandler/type/:type/event/:name
	var paramMapper = make(map[string]string)
	paramMapper[":type"] = eventHandlerType
	paramMapper[":name"] = eventName
	urlParam := util.MappingUrlParameter(originalUrl, paramMapper)

	url := util.DRAGONFLY + urlParam
	// url := util.DRAGONFLY + "/alert/eventhandler/type/" + eventHandlerType + "/event/" + eventName

	pbytes, _ := json.Marshal(vmMonitoringAlertEventHandlerSlackInfo)
	fmt.Println(string(pbytes))
	resp, err := util.CommonHttp(url, pbytes, http.MethodPut)

	resultVmMonitoringAlertEventHandlerInfo := dragonfly.VmMonitoringAlertEventHandlerSlackInfo{}
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
func PutMonitoringAlertEventHandlerSmtp(eventHandlerType string, eventName string, vmMonitoringAlertEventHandlerInfo *dragonfly.EventHandlerOptionSmtp) (*dragonfly.VmMonitoringAlertEventHandlerSmtpInfo, model.WebStatus) {
	fmt.Println("PutMonitoringAlertEventHandlerSmtp ************ : ")
	var originalUrl = "/alert/eventhandler/type/:type/event/:name"
	// {{ip}}:{{port}}/dragonfly/alert/eventhandler/type/:type/event/:name
	var paramMapper = make(map[string]string)
	paramMapper[":type"] = eventHandlerType
	paramMapper[":name"] = eventName
	urlParam := util.MappingUrlParameter(originalUrl, paramMapper)

	url := util.DRAGONFLY + urlParam
	// url := util.DRAGONFLY + "/alert/eventhandler/type/" + eventHandlerType + "/event/" + eventName
	// {{ip}}:{{port}}/dragonfly/alert/eventhandler/type/:type/event/:name

	pbytes, _ := json.Marshal(vmMonitoringAlertEventHandlerInfo)
	fmt.Println(string(pbytes))
	resp, err := util.CommonHttp(url, pbytes, http.MethodPut)

	resultVmMonitoringAlertEventHandlerInfo := dragonfly.VmMonitoringAlertEventHandlerSmtpInfo{}
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
func DelMonitoringAlertEventHandler(eventHandlerType string, eventName string) (io.ReadCloser, model.WebStatus) {
	var originalUrl = "/alert/eventhandler/type/:type/event/:name"
	// {{ip}}:{{port}}/dragonfly/alert/eventhandler/type/:type/event/:name
	var paramMapper = make(map[string]string)
	paramMapper[":type"] = eventHandlerType
	paramMapper[":name"] = eventName
	urlParam := util.MappingUrlParameter(originalUrl, paramMapper)

	url := util.DRAGONFLY + urlParam
	// url := util.DRAGONFLY + "/alert/eventhandler/type/" + eventHandlerType + "/event/" + eventName

	if eventHandlerType == "" {
		return nil, model.WebStatus{StatusCode: 500, Message: "eventHandlerType is required"}
	}
	if eventName == "" {
		return nil, model.WebStatus{StatusCode: 500, Message: "eventName is required"}
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
func GetMonitoringAlertLogList(taskName string, logLevel string) ([]dragonfly.VmMonitoringAlertInfo, model.WebStatus) {
	if logLevel == "" {
		logLevel = "warning"
	}
	var originalUrl = "/alert/task/:task_name/events?level={logLevel}"
	// {{ip}}:{{port}}/dragonfly/alert/task/:task_name/events?level=warning
	var paramMapper = make(map[string]string)
	paramMapper[":task_name"] = taskName
	paramMapper["{logLevel}"] = logLevel
	urlParam := util.MappingUrlParameter(originalUrl, paramMapper)

	url := util.DRAGONFLY + urlParam
	resp, err := util.CommonHttp(url, nil, http.MethodGet)

	if err != nil {
		fmt.Println(err)
		return nil, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode

	vmMonitoringAlertInfoList := map[string][]dragonfly.VmMonitoringAlertInfo{}
	json.NewDecoder(respBody).Decode(&vmMonitoringAlertInfoList)

	return vmMonitoringAlertInfoList["mcis"], model.WebStatus{StatusCode: respStatus}
}
