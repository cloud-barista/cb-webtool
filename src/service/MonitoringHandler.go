package service

import (
	"encoding/json"
	"fmt"
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
