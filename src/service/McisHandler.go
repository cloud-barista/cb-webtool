package service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	// "os"
	// model "github.com/cloud-barista/cb-webtool/src/model"
	"github.com/cloud-barista/cb-webtool/src/model"
	// spider "github.com/cloud-barista/cb-webtool/src/model/spider"
	"github.com/cloud-barista/cb-webtool/src/model/tumblebug"

	util "github.com/cloud-barista/cb-webtool/src/util"
)

//var MCISUrl = "http://15.165.16.67:1323"
//var SPiderUrl = "http://15.165.16.67:1024"

// var SpiderUrl = os.Getenv("SPIDER_URL")// util.SPIDER
// var MCISUrl = os.Getenv("TUMBLE_URL")// util.TUMBLEBUG

// MCIS 목록 조회
func GetMcisList(nameSpaceID string) ([]tumblebug.McisInfo, model.WebStatus) {

	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/mcis"
	resp, err := util.CommonHttp(url, nil, http.MethodGet)
	// resp, err := util.CommonHttpWithoutParam(url, http.MethodGet)

	if err != nil {
		fmt.Println(err)
		return nil, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode

	mcisList := map[string][]tumblebug.McisInfo{}
	json.NewDecoder(respBody).Decode(&mcisList)
	fmt.Println(mcisList["mcis"])
	log.Println(respBody)
	util.DisplayResponse(resp) // 수신내용 확인

	return mcisList["mcis"], model.WebStatus{StatusCode: respStatus}
}

// 특정 MCIS 조회
func GetMcisData(nameSpaceID string, mcisID string) (*tumblebug.McisInfo, model.WebStatus) {

	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/mcis/" + mcisID
	// /ns/{nsId}/mcis/{mcisId}

	// 	// resp, err := http.Get(url)
	// 	// if err != nil {
	// 	// 	fmt.Println("request URL : ", url)
	// 	// }

	// 	// defer resp.Body.Close()
	// 	body := HttpGetHandler(url)
	// 	defer body.Close()
	// 	info := map[string][]McisInfo{}
	// 	json.NewDecoder(body).Decode(&info)
	// 	fmt.Println("info : ", info["mcis"][0].ID)
	// 	return info["ns"]

	// resp, err := util.CommonHttp(url, nil, http.MethodGet)
	resp, err := util.CommonHttpWithoutParam(url, http.MethodGet)

	// defer body.Close()
	mcisInfo := tumblebug.McisInfo{}
	if err != nil {
		fmt.Println(err)
		return &mcisInfo, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}
	// util.DisplayResponse(resp) // 수신내용 확인

	respBody := resp.Body
	respStatus := resp.StatusCode

	json.NewDecoder(respBody).Decode(&mcisInfo)
	fmt.Println(mcisInfo)

	// resultBody, err := ioutil.ReadAll(respBody)
	// if err == nil {
	// 	str := string(resultBody)
	// 	println(str)
	// }
	// pbytes, _ := json.Marshal(respBody)
	// fmt.Println(string(pbytes))

	return &mcisInfo, model.WebStatus{StatusCode: respStatus}
}

// MCIS 등록. VM도 함께 등록
func RegMcis(nameSpaceID string, mcisInfo *tumblebug.McisInfo) (*tumblebug.McisInfo, model.WebStatus) {

	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/mcis"

	pbytes, _ := json.Marshal(mcisInfo)
	resp, err := util.CommonHttp(url, pbytes, http.MethodPost)

	returnMcisInfo := tumblebug.McisInfo{}
	returnStatus := model.WebStatus{}

	respBody := resp.Body
	respStatus := resp.StatusCode

	if err != nil {
		fmt.Println(err)
		return &returnMcisInfo, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	if respStatus != 200 && respStatus != 201 { // 호출은 정상이나, 가져온 결과값이 200, 201아닌 경우 message에 담겨있는 것을 WebStatus에 set
		errorInfo := model.ErrorInfo{}
		json.NewDecoder(respBody).Decode(&errorInfo)
		fmt.Println("respStatus != 200 reason ", errorInfo)
		returnStatus.Message = errorInfo.Message
	} else {
		json.NewDecoder(respBody).Decode(&returnMcisInfo)
		fmt.Println(returnMcisInfo)
	}
	returnStatus.StatusCode = respStatus

	// return respBody, respStatusCode
	return &returnMcisInfo, returnStatus

	// resultMcisInfo := tumblebug.McisInfo{}
	// if err != nil {
	// 	fmt.Println(err)
	// 	return &resultMcisInfo, model.WebStatus{StatusCode: 500, Message: err.Error()}
	// }

	// if err != nil {
	// 	fmt.Println(err)
	// 	return nil, model.WebStatus{StatusCode: 500, Message: err.Error()}
	// }

	// respBody := resp.Body
	// respStatus := resp.StatusCode

	// return respBody, model.WebStatus{StatusCode: respStatus}

	// respBody := resp.Body
	// respStatus := resp.StatusCode
	// // respStatus := resp.Status
	// // log.Println("respStatusCode = ", respStatusCode)
	// // log.Println("respStatus = ", respStatus)

	// // 응답에 생성한 객체값이 옴

	// json.NewDecoder(respBody).Decode(resultMcisInfo)
	// fmt.Println(resultMcisInfo)

	// // return body, err
	// // respBody := resp.Body
	// // respStatus := resp.StatusCode
	// // return respBody, respStatus
	// return &resultMcisInfo, model.WebStatus{StatusCode: respStatus}

}

// func GetVMStatus(vm_name string, connectionConfig string) string {
// 	url := SpiderUrl + "/vmstatus/" + vm_name + "?connection_name=" + connectionConfig
// 	// resp, err := http.Get(url)
// 	// if err != nil {
// 	// 	fmt.Println("request URL : ", url)
// 	// }

// 	// defer resp.Body.Close()
// 	body := HttpGetHandler(url)
// 	defer body.Close()
// 	info := map[string]McisInfo{}
// 	json.NewDecoder(body).Decode(&info)
// 	fmt.Println("VM Status : ", info["status"].Status)
// 	return info["status"].Status

// }

// MCIS 목록에서 mcis 상태별 count map반환
func GetMcisStatusCountMap(mcisInfo tumblebug.McisInfo) map[string]int {
	mcisStatusRunning := 0
	mcisStatusStopped := 0
	mcisStatusTerminated := 0

	// log.Println(" mcisInfo  ", index, mcisInfo)
	mcisStatus := util.GetMcisStatus(mcisInfo.Status)
	if mcisStatus == util.MCIS_STATUS_RUNNING {
		mcisStatusRunning++
	} else if mcisStatus == util.MCIS_STATUS_TERMINATED {
		mcisStatusTerminated++
	} else {
		mcisStatusStopped++
	}
	mcisStatusMap := make(map[string]int)
	mcisStatusMap["RUNNING"] = mcisStatusRunning
	mcisStatusMap["STOPPED"] = mcisStatusStopped
	mcisStatusMap["TERMINATED"] = mcisStatusTerminated
	mcisStatusMap["TOTAL"] = mcisStatusRunning + mcisStatusStopped + mcisStatusTerminated
	// mcisStatusTotalMap[mcisInfo.ID] = mcisStatusMap

	return mcisStatusMap
}

// MCIS의 vm별 statun와 vm 상태별 count
// key는 vmID + vmName, value는 vmStatus
func GetSimpleVmWithStatusCountMap(mcisInfo tumblebug.McisInfo) ([]tumblebug.VmSimpleInfo, map[string]int) {
	// log.Println(" mcisInfo  ", index, mcisInfo)
	// vmStatusMap := make(map[string]int)
	// vmStatusMap := map[string]string{} // vmName : vmStatus

	vmStatusCountMap := map[string]int{}
	totalVmStatusCount := 0
	vmList := mcisInfo.Vms
	var vmSimpleList []tumblebug.VmSimpleInfo
	for vmIndex, vmInfo := range vmList {
		// log.Println(" vmInfo  ", vmIndex, vmInfo)
		vmStatus := util.GetVmStatus(vmInfo.Status) // lowercase로 변환

		locationInfo := vmInfo.Location
		vmLatitude := locationInfo.Latitude
		vmLongitude := locationInfo.Longitude

		log.Println(locationInfo)
		//
		vmSimpleObj := tumblebug.VmSimpleInfo{
			VmIndex:   vmIndex + 1,
			VmID:      vmInfo.ID,
			VmName:    vmInfo.Name,
			VmStatus:  vmStatus,
			Latitude:  vmLatitude,
			Longitude: vmLongitude,
		}
		vmSimpleList = append(vmSimpleList, vmSimpleObj)

		log.Println("vmStatus " + vmStatus + ", Status " + vmInfo.Status)
		vmStatusCount := 0
		val, exists := vmStatusCountMap[vmStatus]
		if exists {
			vmStatusCount = val + 1
			totalVmStatusCount += 1
		} else {
			vmStatusCount = 1
			totalVmStatusCount += 1
		}
		vmStatusCountMap[vmStatus] = vmStatusCount
	}

	vmStatusCountMap["TOTAL"] = totalVmStatusCount
	log.Println(vmStatusCountMap)
	// vmStatusCountMap := make(map[string]int)
	// UI에서 사칙연산이 되지 않아 controller에서 계산한 뒤 넘겨 줌. 아나면 function을 정의해서 넘겨야 함
	// vmStatusCountMap[util.VM_STATUS_RUNNING] = vmStatusRunning
	// vmStatusCountMap[util.VM_STATUS_RESUMING] = vmStatusResuming
	// vmStatusCountMap[util.VM_STATUS_INCLUDE] = vmStatusInclude
	// vmStatusCountMap[util.VM_STATUS_SUSPENDED] = vmStatusSuspended
	// vmStatusCountMap[util.VM_STATUS_TERMINATED] = vmStatusTerminated
	// vmStatusCountMap[util.VM_STATUS_UNDEFINED] = vmStatusUndefined
	// vmStatusCountMap[util.VM_STATUS_PARTIAL] = vmStatusPartial
	// vmStatusCountMap[util.VM_STATUS_ETC] = vmStatusEtc
	// log.Println("mcisInfo.ID  ", mcisInfo.ID)
	// mcisIdArr[mcisIndex] = mcisInfo.ID	// 바로 넣으면 Runtime Error구만..
	// vmStatusArr[mcisIndex] = vmStatusCountMap

	// UI에서는 3가지로 통합하여 봄
	// vmStatusCountMap["RUNNING"] = vmStatusRunning
	// vmStatusCountMap["STOPPED"] = vmStatusInclude + vmStatusSuspended + vmStatusUndefined + vmStatusPartial + vmStatusEtc
	// vmStatusCountMap["TERMINATED"] = vmStatusTerminated
	// vmStatusTotalMap[mcisInfo.ID] = vmStatusCountMap
	// vmIdArr = append(vmIdArr, vmInfo.ID)
	// vmStatusArr = append(vmStatusArr, vmStatusCountMap)

	// log.Println("mcisIndex  ", mcisIndex)

	// vmStatusCountMap["RUNNING"] = vmStatusRunning
	// vmStatusCountMap["STOPPED"] = vmStatusInclude + vmStatusSuspended + vmStatusUndefined + vmStatusPartial + vmStatusEtc
	// vmStatusCountMap["TERMINATED"] = vmStatusTerminated
	// vmStatusCountMap["TOTAL"] = vmStatusCountMap["RUNNING"] + vmStatusCountMap["STOPPED"] + vmStatusCountMap["TERMINATED"]

	return vmSimpleList, vmStatusCountMap

}

// MCIS별 connection count
func GetVmConnectionCountMap(mcisInfo tumblebug.McisInfo) map[string]int {
	// connectionCountTotal := 0
	// connectionCountByMcis := 0
	// vmCountTotal := 0
	// vmRunningCountByMcis := 0
	// vmStoppedCountByMcis := 0
	// vmTerminatedCountByMcis := 0
	// vmStatusUndefined := 0
	// vmStatusPartial := 0
	// vmStatusEtc := 0
	// vmStatusTerminated := 0

	// log.Println(" mcisInfo  ", index, mcisInfo)
	// vmList := mcisInfo.VMs
	// for vmIndex, vmInfo := range vmList {
	// 	// log.Println(" vmInfo  ", vmIndex, vmInfo)
	// 	vmConnection := util.GetVmConnectionName(vmInfo.ConnectionName)

	// }
	vmStatusMap := make(map[string]int)
	// UI에서 사칙연산이 되지 않아 controller에서 계산한 뒤 넘겨 줌.
	// vmStatusMap[util.VM_STATUS_RUNNING] = vmStatusRunning
	// vmStatusMap[util.VM_STATUS_RESUMING] = vmStatusResuming
	// vmStatusMap[util.VM_STATUS_INCLUDE] = vmStatusInclude
	// vmStatusMap[util.VM_STATUS_SUSPENDED] = vmStatusSuspended
	// vmStatusMap[util.VM_STATUS_TERMINATED] = vmStatusTerminated
	// vmStatusMap[util.VM_STATUS_UNDEFINED] = vmStatusUndefined
	// vmStatusMap[util.VM_STATUS_PARTIAL] = vmStatusPartial
	// vmStatusMap[util.VM_STATUS_ETC] = vmStatusEtc
	// log.Println("mcisInfo.ID  ", mcisInfo.ID)
	// mcisIdArr[mcisIndex] = mcisInfo.ID	// 바로 넣으면 Runtime Error구만..
	// vmStatusArr[mcisIndex] = vmStatusMap

	// UI에서는 3가지로 통합하여 봄
	// vmStatusMap["RUNNING"] = vmStatusRunning
	// vmStatusMap["STOPPED"] = vmStatusInclude + vmStatusSuspended + vmStatusUndefined + vmStatusPartial + vmStatusEtc
	// vmStatusMap["TERMINATED"] = vmStatusTerminated
	// vmStatusTotalMap[mcisInfo.ID] = vmStatusMap
	// vmIdArr = append(vmIdArr, vmInfo.ID)
	// vmStatusArr = append(vmStatusArr, vmStatusMap)

	// log.Println("mcisIndex  ", mcisIndex)

	// vmStatusMap := make(map[string]int)
	// vmStatusMap["RUNNING"] = vmStatusRunning
	// vmStatusMap["STOPPED"] = vmStatusInclude + vmStatusSuspended + vmStatusUndefined + vmStatusPartial + vmStatusEtc
	// vmStatusMap["TERMINATED"] = vmStatusTerminated

	return vmStatusMap

}

// 해당 MCIS의 VM 연결 수
func GetVmConnectionCountByMcis(mcisInfo tumblebug.McisInfo) map[string]int {
	// log.Println(" mcisInfo  ", index, mcisInfo)
	vmList := mcisInfo.Vms
	// mcisConnectionCountMap := make(map[string]int)
	mcisConnectionCountMap := map[string]int{}

	totalConnectionCount := 0
	log.Println("GetVMConnectionCountByMcis map length ", len(mcisConnectionCountMap))
	for _, vmInfo := range vmList {
		// log.Println(" vmInfo  ", vmIndex, vmInfo)
		locationInfo := vmInfo.Location
		// cloudType := locationInfo.CloudType // CloudConnection
		providerCount := 0
		val, exists := mcisConnectionCountMap[util.GetProviderName(locationInfo.CloudType)]
		if exists {
			providerCount = val + 1
			// totalConnectionCount += 1 // 이미 있는 경우에는 count추가필요없음
		} else {
			providerCount = 1
			totalConnectionCount += 1
		}
		log.Println("GetProviderName ", locationInfo.CloudType)
		mcisConnectionCountMap[util.GetProviderName(locationInfo.CloudType)] = providerCount
	}
	log.Println("GetVMConnectionCountByMcis map length ", len(mcisConnectionCountMap))
	log.Println("GetVMConnectionCountByMcis map ", mcisConnectionCountMap)
	return mcisConnectionCountMap
}

// MCIS의 특정 VM 조회
func GetVMofMcisData(nameSpaceID string, mcisID string, vmID string) (*tumblebug.VmInfo, model.WebStatus) {
	///ns/{nsId}/mcis/{mcisId}/vm/{vmId}
	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/mcis/" + mcisID + "/vm/" + vmID

	// resp, err := util.CommonHttp(url, nil, http.MethodGet)
	resp, err := util.CommonHttpWithoutParam(url, http.MethodGet)

	// defer body.Close()
	vmInfo := tumblebug.VmInfo{}
	if err != nil {
		fmt.Println(err)
		return &vmInfo, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}
	// util.DisplayResponse(resp) // 수신내용 확인

	respBody := resp.Body
	respStatus := resp.StatusCode

	json.NewDecoder(respBody).Decode(&vmInfo)
	fmt.Println(vmInfo)

	// resultBody, err := ioutil.ReadAll(respBody)
	// if err == nil {
	// 	str := string(resultBody)
	// 	println(str)
	// }
	// pbytes, _ := json.Marshal(respBody)
	// fmt.Println(string(pbytes))

	return &vmInfo, model.WebStatus{StatusCode: respStatus}
}

// MCIS의 Status변경
func McisLifeCycle(mcisLifeCycle *tumblebug.McisLifeCycle) (*tumblebug.McisLifeCycle, model.WebStatus) {
	url := util.TUMBLEBUG + "/ns/" + mcisLifeCycle.NameSpaceID + "/mcis/" + mcisLifeCycle.McisID + "?action=" + mcisLifeCycle.LifeCycleType
	//// var url = CommonURL+"/ns/"+nameSpace+"/mcis/"+mcis_id+"?action="+type
	pbytes, _ := json.Marshal(mcisLifeCycle)
	resp, err := util.CommonHttp(url, pbytes, http.MethodGet) // POST로 받기는 했으나 실제로는 Get으로 날아감.
	resultMcisLifeCycle := tumblebug.McisLifeCycle{}
	if err != nil {
		fmt.Println(err)
		return &resultMcisLifeCycle, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode
	// respStatus := resp.Status
	// log.Println("respStatusCode = ", respStatusCode)
	// log.Println("respStatus = ", respStatus)

	// 응답에 생성한 객체값이 옴

	json.NewDecoder(respBody).Decode(resultMcisLifeCycle)
	fmt.Println(resultMcisLifeCycle)

	// return body, err
	// respBody := resp.Body
	// respStatus := resp.StatusCode
	// return respBody, respStatus
	return &resultMcisLifeCycle, model.WebStatus{StatusCode: respStatus}

}

// MCIS의 VM Status변경
func McisVmLifeCycle(vmLifeCycle *tumblebug.VmLifeCycle) (*tumblebug.VmLifeCycle, model.WebStatus) {

	url := util.TUMBLEBUG + "/ns/" + vmLifeCycle.NameSpaceID + "/mcis/" + vmLifeCycle.McisID + "/vm/" + vmLifeCycle.VmID + "?action=" + vmLifeCycle.LifeCycleType
	///url = CommonURL+"/ns/"+nameSpace+"/mcis/"+mcis_id+"/vm/"+vm_id+"?action="+type
	pbytes, _ := json.Marshal(vmLifeCycle)
	resp, err := util.CommonHttp(url, pbytes, http.MethodGet) // POST로 받기는 했으나 실제로는 Get으로 날아감.
	resultVmLifeCycle := tumblebug.VmLifeCycle{}
	if err != nil {
		fmt.Println(err)
		return &resultVmLifeCycle, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode
	// respStatus := resp.Status
	// log.Println("respStatusCode = ", respStatusCode)
	// log.Println("respStatus = ", respStatus)

	// 응답에 생성한 객체값이 옴

	json.NewDecoder(respBody).Decode(resultVmLifeCycle)
	fmt.Println(resultVmLifeCycle)

	// return body, err
	// respBody := resp.Body
	// respStatus := resp.StatusCode
	// return respBody, respStatus
	return &resultVmLifeCycle, model.WebStatus{StatusCode: respStatus}

}

// 벤치마크?? MCIS 조회. 근데 왜 결과는 resultarray지?
// TODO : 여러개 return되면 method이름을 xxxData -> xxxList 로 바꿀 것
func GetBenchmarkMcisData(nameSpaceID string, mcisID string, host string) (*[]tumblebug.McisBenchmarkInfo, model.WebStatus) {
	// func GetMCIS(nsid string, mcisId string) []McisInfo {
	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/benchmark/mcis/" + mcisID
	// /ns/{nsId}/benchmark/mcis/{mcisId}
	pbytes, _ := json.Marshal(host)
	resp, err := util.CommonHttp(url, pbytes, http.MethodGet)

	// defer body.Close()
	resultBenchmarkInfos := []tumblebug.McisBenchmarkInfo{}
	if err != nil {
		fmt.Println(err)
		return &resultBenchmarkInfos, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}
	// util.DisplayResponse(resp) // 수신내용 확인

	respBody := resp.Body
	respStatus := resp.StatusCode

	json.NewDecoder(respBody).Decode(&resultBenchmarkInfos)
	fmt.Println(resultBenchmarkInfos)

	return &resultBenchmarkInfos, model.WebStatus{StatusCode: respStatus}
}

func GetBenchmarkAllMcisList(nameSpaceID string, mcisID string, host string) (*[]tumblebug.McisBenchmarkInfo, model.WebStatus) {

	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/benchmark/mcis/" + mcisID
	// /ns/{nsId}/benchmark/mcis/{mcisId}
	pbytes, _ := json.Marshal(host)
	resp, err := util.CommonHttp(url, pbytes, http.MethodGet)

	// defer body.Close()
	resultBenchmarkInfos := []tumblebug.McisBenchmarkInfo{}
	if err != nil {
		fmt.Println(err)
		return &resultBenchmarkInfos, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}
	// util.DisplayResponse(resp) // 수신내용 확인

	respBody := resp.Body
	respStatus := resp.StatusCode

	json.NewDecoder(respBody).Decode(&resultBenchmarkInfos)
	fmt.Println(resultBenchmarkInfos)

	return &resultBenchmarkInfos, model.WebStatus{StatusCode: respStatus}
}

// MCIS에 명령 내리기
func CommandMcis(nameSpaceID string, mcisCommandInfo *tumblebug.McisCommandInfo) (*tumblebug.McisCommandResult, model.WebStatus) {

	mcisID := mcisCommandInfo.McisID
	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/cmd/mcis/" + mcisID
	// /ns/{nsId}/cmd/mcis/{mcisId}
	pbytes, _ := json.Marshal(mcisCommandInfo)
	resp, err := util.CommonHttp(url, pbytes, http.MethodPost)

	returnMcisCommandResult := tumblebug.McisCommandResult{}
	returnStatus := model.WebStatus{}

	respBody := resp.Body
	respStatus := resp.StatusCode

	if err != nil {
		fmt.Println(err)
		return &returnMcisCommandResult, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	if respStatus != 200 && respStatus != 201 { // 호출은 정상이나, 가져온 결과값이 200, 201아닌 경우 message에 담겨있는 것을 WebStatus에 set
		errorInfo := model.ErrorInfo{}
		json.NewDecoder(respBody).Decode(&errorInfo)
		fmt.Println("respStatus != 200 reason ", errorInfo)
		returnStatus.Message = errorInfo.Message
	} else {
		json.NewDecoder(respBody).Decode(&returnMcisCommandResult)
		fmt.Println(returnMcisCommandResult)
	}
	returnStatus.StatusCode = respStatus

	return &returnMcisCommandResult, returnStatus

	// resultMcisCommandResult := tumblebug.McisCommandResult{}
	// if err != nil {
	// 	fmt.Println(err)
	// 	return &resultMcisCommandResult, model.WebStatus{StatusCode: 500, Message: err.Error()}
	// }

	// respBody := resp.Body
	// respStatus := resp.StatusCode

	// json.NewDecoder(respBody).Decode(resultMcisCommandResult)
	// fmt.Println(resultMcisCommandResult)
	// return &resultMcisCommandResult, model.WebStatus{StatusCode: respStatus}

	// if err != nil {
	// 	fmt.Println(err)
	// 	return nil, model.WebStatus{StatusCode: 500, Message: err.Error()}
	// }

	// respBody := resp.Body
	// respStatus := resp.StatusCode

	// return respBody, model.WebStatus{StatusCode: respStatus}
}

// 특정 VM에 명령내리기
func CommandVMOfMcis(nameSpaceID string, mcisCommandInfo *tumblebug.McisCommandInfo) (*tumblebug.McisCommandResult, model.WebStatus) {

	mcisID := mcisCommandInfo.McisID
	vmID := mcisCommandInfo.VmID
	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/mcis/" + mcisID + "/vm/" + vmID
	// /ns/{nsId}/cmd/mcis/{mcisId}/vm/{vmId}
	pbytes, _ := json.Marshal(mcisCommandInfo)
	resp, err := util.CommonHttp(url, pbytes, http.MethodPost)

	returnMcisCommandResult := tumblebug.McisCommandResult{}
	returnStatus := model.WebStatus{}

	respBody := resp.Body
	respStatus := resp.StatusCode

	if err != nil {
		fmt.Println(err)
		return &returnMcisCommandResult, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	if respStatus != 200 && respStatus != 201 { // 호출은 정상이나, 가져온 결과값이 200, 201아닌 경우 message에 담겨있는 것을 WebStatus에 set
		errorInfo := model.ErrorInfo{}
		json.NewDecoder(respBody).Decode(&errorInfo)
		fmt.Println("respStatus != 200 reason ", errorInfo)
		returnStatus.Message = errorInfo.Message
	} else {
		json.NewDecoder(respBody).Decode(&returnMcisCommandResult)
		fmt.Println(returnMcisCommandResult)
	}
	returnStatus.StatusCode = respStatus

	// return respBody, respStatusCode
	return &returnMcisCommandResult, returnStatus

	// resultMcisCommandResult := tumblebug.McisCommandResult{}
	// if err != nil {
	// 	fmt.Println(err)
	// 	return &resultMcisCommandResult, model.WebStatus{StatusCode: 500, Message: err.Error()}
	// }

	// respBody := resp.Body
	// respStatus := resp.StatusCode

	// // TODO : result는 string으로 "result" 만 있는데...
	// json.NewDecoder(respBody).Decode(resultMcisCommandResult)
	// fmt.Println(resultMcisCommandResult)
	// return &resultMcisCommandResult, model.WebStatus{StatusCode: respStatus}
	// if err != nil {
	// 	fmt.Println(err)
	// 	return nil, model.WebStatus{StatusCode: 500, Message: err.Error()}
	// }

	// respBody := resp.Body
	// respStatus := resp.StatusCode

	// return respBody, model.WebStatus{StatusCode: respStatus}
}

//Install the benchmark agent to specified MCIS
func InstallBenchmarkAgentToMcis(nameSpaceID string, mcisCommandInfo *tumblebug.McisCommandInfo) (*tumblebug.McisCommandResult, model.WebStatus) {

	mcisID := mcisCommandInfo.McisID
	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/mcis/" + mcisID
	// /ns/{nsId}/install/mcis/{mcisId}
	pbytes, _ := json.Marshal(mcisCommandInfo)
	resp, err := util.CommonHttp(url, pbytes, http.MethodPost)

	returnMcisCommandResult := tumblebug.McisCommandResult{}
	returnStatus := model.WebStatus{}

	respBody := resp.Body
	respStatus := resp.StatusCode

	if err != nil {
		fmt.Println(err)
		return &returnMcisCommandResult, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	if respStatus != 200 && respStatus != 201 { // 호출은 정상이나, 가져온 결과값이 200, 201아닌 경우 message에 담겨있는 것을 WebStatus에 set
		errorInfo := model.ErrorInfo{}
		json.NewDecoder(respBody).Decode(&errorInfo)
		fmt.Println("respStatus != 200 reason ", errorInfo)
		returnStatus.Message = errorInfo.Message
	} else {
		json.NewDecoder(respBody).Decode(&returnMcisCommandResult)
		fmt.Println(returnMcisCommandResult)
	}
	returnStatus.StatusCode = respStatus

	// return respBody, respStatusCode
	return &returnMcisCommandResult, returnStatus

	// resultMcisCommandResult := tumblebug.McisCommandResult{}
	// if err != nil {
	// 	fmt.Println(err)
	// 	return &resultMcisCommandResult, model.WebStatus{StatusCode: 500, Message: err.Error()}
	// }

	// respBody := resp.Body
	// respStatus := resp.StatusCode

	// // TODO : result는 resultArray인데....
	// json.NewDecoder(respBody).Decode(resultMcisCommandResult)
	// fmt.Println(resultMcisCommandResult)
	// return &resultMcisCommandResult, model.WebStatus{StatusCode: respStatus}
	// if err != nil {
	// 	fmt.Println(err)
	// 	return nil, model.WebStatus{StatusCode: 500, Message: err.Error()}
	// }

	// respBody := resp.Body
	// respStatus := resp.StatusCode

	// return respBody, model.WebStatus{StatusCode: respStatus}
}

func DelAllMcis(nameSpaceID string) (io.ReadCloser, model.WebStatus) {
	// buff := bytes.NewBuffer(pbytes)
	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/mcis/"
	// /ns/{nsId}/mcis/{mcisId}
	fmt.Println("url : ", url)

	//body, err := util.CommonHttpPost(url, nsInfo)

	// 경로안에 parameter가 있어 추가 param없이 호출 함.
	resp, err := util.CommonHttp(url, nil, http.MethodDelete)
	// body, err := util.CommonHttpDelete(url, pbytes)
	if err != nil {
		fmt.Println(err)
		return nil, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}
	// return body, err
	respBody := resp.Body
	respStatus := resp.StatusCode

	return respBody, model.WebStatus{StatusCode: respStatus}
}

// MCIS 삭제. TODO : 해당 namespace의 MCIS만 삭제 가능... 창 두개에서 1개는 MCIS삭제, 1개는 namespace 변경이 있을 수 있으므로 UI에서 namespace도 넘겨서 비교할 것.
func DelMcis(nameSpaceID string, mcisID string) (io.ReadCloser, model.WebStatus) {
	// buff := bytes.NewBuffer(pbytes)
	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/mcis/" + mcisID
	// /ns/{nsId}/mcis/{mcisId}
	fmt.Println("url : ", url)

	if mcisID == "" {
		return nil, model.WebStatus{StatusCode: 500, Message: "MCIS ID is required"}
	}

	//body, err := util.CommonHttpPost(url, nsInfo)

	// 경로안에 parameter가 있어 추가 param없이 호출 함.
	resp, err := util.CommonHttp(url, nil, http.MethodDelete)
	// body, err := util.CommonHttpDelete(url, pbytes)
	if err != nil {
		fmt.Println(err)
		return nil, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}
	// return body, err
	respBody := resp.Body
	respStatus := resp.StatusCode
	return respBody, model.WebStatus{StatusCode: respStatus}
}

// MCIS에 VM 생성. path에 mcisID가 있음. VMInfo에는 mcisID가 없음.
func RegVM(nameSpaceID string, mcisID string, vmInfo *tumblebug.VmInfo) (*tumblebug.VmInfo, model.WebStatus) {
	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/mcis/" + mcisID + "/vm"
	// /ns/{nsId}/mcis/{mcisId}/vm
	pbytes, _ := json.Marshal(vmInfo)
	resp, err := util.CommonHttp(url, pbytes, http.MethodPost)

	returnVmInfo := tumblebug.VmInfo{}
	returnStatus := model.WebStatus{}

	respBody := resp.Body
	respStatus := resp.StatusCode

	if err != nil {
		fmt.Println(err)
		return &returnVmInfo, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	if respStatus != 200 && respStatus != 201 { // 호출은 정상이나, 가져온 결과값이 200, 201아닌 경우 message에 담겨있는 것을 WebStatus에 set
		errorInfo := model.ErrorInfo{}
		json.NewDecoder(respBody).Decode(&errorInfo)
		fmt.Println("respStatus != 200 reason ", errorInfo)
		returnStatus.Message = errorInfo.Message
	} else {
		json.NewDecoder(respBody).Decode(&returnVmInfo)
		fmt.Println(returnVmInfo)
	}
	returnStatus.StatusCode = respStatus

	// return respBody, respStatusCode
	return &returnVmInfo, returnStatus

	// resultVmResult := tumblebug.VmInfo{}
	// if err != nil {
	// 	fmt.Println(err)
	// 	return &resultVmResult, model.WebStatus{StatusCode: 500, Message: err.Error()}
	// }

	// respBody := resp.Body
	// respStatus := resp.StatusCode

	// // TODO : result는 resultArray인데....
	// json.NewDecoder(respBody).Decode(resultVmResult)
	// fmt.Println(resultVmResult)
	// return &resultVmResult, model.WebStatus{StatusCode: respStatus}
	// if err != nil {
	// 	fmt.Println(err)
	// 	return nil, model.WebStatus{StatusCode: 500, Message: err.Error()}
	// }

	// respBody := resp.Body
	// respStatus := resp.StatusCode

	// return respBody, model.WebStatus{StatusCode: respStatus}
}

func DelVM(nameSpaceID string, mcisID string, vmID string) (io.ReadCloser, model.WebStatus) {

	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/mcis/" + mcisID + "/vm"
	// /ns/{nsId}/mcis/{mcisId}/vm/{vmId}
	fmt.Println("url : ", url)

	if vmID == "" {
		return nil, model.WebStatus{StatusCode: 500, Message: "vmID ID is required"}
	}

	// 경로안에 parameter가 있어 추가 param없이 호출 함.
	resp, err := util.CommonHttp(url, nil, http.MethodDelete)
	// body, err := util.CommonHttpDelete(url, pbytes)
	if err != nil {
		fmt.Println(err)
		return nil, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}
	// return body, err
	respBody := resp.Body
	respStatus := resp.StatusCode
	return respBody, model.WebStatus{StatusCode: respStatus}
}

// 특정 VM 조회
func GetVmData(nameSpaceID string, mcisID string, vmID string) (*tumblebug.VmInfo, model.WebStatus) {

	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/mcis/" + mcisID + "/vm/" + vmID
	// /ns/{nsId}/mcis/{mcisId}/vm/{vmId}

	resp, err := util.CommonHttpWithoutParam(url, http.MethodGet)

	// defer body.Close()
	vmInfo := tumblebug.VmInfo{}
	if err != nil {
		fmt.Println(err)
		return &vmInfo, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}
	// util.DisplayResponse(resp) // 수신내용 확인

	respBody := resp.Body
	respStatus := resp.StatusCode

	json.NewDecoder(respBody).Decode(&vmInfo)
	fmt.Println(vmInfo)

	return &vmInfo, model.WebStatus{StatusCode: respStatus}
}

// Get MCIS recommendation
func GetMcisRecommand(nameSpaceID string, mcisID string, mcisRecommandReq *tumblebug.McisRecommendReq) (*tumblebug.McisRecommendInfo, model.WebStatus) {
	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/mcis/recommend"
	// /ns/{nsId}/mcis/recommend
	pbytes, _ := json.Marshal(mcisRecommandReq)
	resp, err := util.CommonHttp(url, pbytes, http.MethodPost)

	returnMcisRecommendInfo := tumblebug.McisRecommendInfo{}
	returnStatus := model.WebStatus{}

	respBody := resp.Body
	respStatus := resp.StatusCode

	if err != nil {
		fmt.Println(err)
		return &returnMcisRecommendInfo, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	if respStatus != 200 && respStatus != 201 { // 호출은 정상이나, 가져온 결과값이 200, 201아닌 경우 message에 담겨있는 것을 WebStatus에 set
		errorInfo := model.ErrorInfo{}
		json.NewDecoder(respBody).Decode(&errorInfo)
		fmt.Println("respStatus != 200 reason ", errorInfo)
		returnStatus.Message = errorInfo.Message
	} else {
		json.NewDecoder(respBody).Decode(&returnMcisRecommendInfo)
		fmt.Println(returnMcisRecommendInfo)
	}
	returnStatus.StatusCode = respStatus

	return &returnMcisRecommendInfo, returnStatus

	// mcisRecommandesult := tumblebug.McisRecommendInfo{}
	// if err != nil {
	// 	fmt.Println(err)
	// 	return &mcisRecommandesult, model.WebStatus{StatusCode: 500, Message: err.Error()}
	// }

	// respBody := resp.Body
	// respStatus := resp.StatusCode

	// // TODO : result는 resultArray인데....
	// json.NewDecoder(respBody).Decode(mcisRecommandesult)
	// fmt.Println(mcisRecommandesult)
	// return &mcisRecommandesult, model.WebStatus{StatusCode: respStatus}
	// if err != nil {
	// 	fmt.Println(err)
	// 	return nil, model.WebStatus{StatusCode: 500, Message: err.Error()}
	// }

	// respBody := resp.Body
	// respStatus := resp.StatusCode

	// return respBody, model.WebStatus{StatusCode: respStatus}
}
