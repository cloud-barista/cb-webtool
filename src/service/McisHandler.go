package service

import (
	"encoding/json"
	"fmt"
	// "io"
	"log"
	"net/http"

	// "os"
	model "github.com/cloud-barista/cb-webtool/src/model"
	util "github.com/cloud-barista/cb-webtool/src/util"
)

//var MCISUrl = "http://15.165.16.67:1323"
//var SPiderUrl = "http://15.165.16.67:1024"

// var SpiderUrl = os.Getenv("SPIDER_URL")// util.SPIDER
// var MCISUrl = os.Getenv("TUMBLE_URL")// util.TUMBLEBUG

// MCIS 목록 조회
func GetMcisList(nameSpaceID string) ([]model.MCISInfo, model.WebStatus) {
	// func GetMCISList(nsid string) []MCISInfo {
	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/mcis"
	resp, err := util.CommonHttp(url, nil, http.MethodGet)
	// resp, err := util.CommonHttpWithoutParam(url, http.MethodGet)

	if err != nil {
		fmt.Println(err)
		return nil, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode

	mcisList := map[string][]model.MCISInfo{}
	json.NewDecoder(respBody).Decode(&mcisList)
	fmt.Println(mcisList["mcis"])
	log.Println(respBody)
	util.DisplayResponse(resp) // 수신내용 확인

	return mcisList["mcis"], model.WebStatus{StatusCode: respStatus}
}

// 특정 MCIS 조회
func GetMcisData(nameSpaceID string, mcisID string) (*model.MCISInfo, model.WebStatus) {
	// func GetMCIS(nsid string, mcisId string) []MCISInfo {
	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/mcis/" + mcisID
	// 	// resp, err := http.Get(url)
	// 	// if err != nil {
	// 	// 	fmt.Println("request URL : ", url)
	// 	// }

	// 	// defer resp.Body.Close()
	// 	body := HttpGetHandler(url)
	// 	defer body.Close()
	// 	info := map[string][]MCISInfo{}
	// 	json.NewDecoder(body).Decode(&info)
	// 	fmt.Println("info : ", info["mcis"][0].ID)
	// 	return info["ns"]

	// resp, err := util.CommonHttp(url, nil, http.MethodGet)
	resp, err := util.CommonHttpWithoutParam(url, http.MethodGet)

	// defer body.Close()
	mcisInfo := model.MCISInfo{}
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

func RegMcis(nameSpaceID string, mCISInfo *model.MCISInfo) (*model.MCISInfo, model.WebStatus) {

	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/mcis"

	pbytes, _ := json.Marshal(mCISInfo)
	resp, err := util.CommonHttp(url, pbytes, http.MethodPost)

	resultMcisInfo := model.MCISInfo{}
	if err != nil {
		fmt.Println(err)
		return &resultMcisInfo, model.WebStatus{StatusCode: 500, Message: err.Error()}
	}

	respBody := resp.Body
	respStatus := resp.StatusCode
	// respStatus := resp.Status
	// log.Println("respStatusCode = ", respStatusCode)
	// log.Println("respStatus = ", respStatus)

	// 응답에 생성한 객체값이 옴
	
	json.NewDecoder(respBody).Decode(resultMcisInfo)
	fmt.Println(resultMcisInfo)

	// return body, err
	// respBody := resp.Body
	// respStatus := resp.StatusCode
	// return respBody, respStatus
	return &resultMcisInfo, model.WebStatus{StatusCode: respStatus}

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
// 	info := map[string]MCISInfo{}
// 	json.NewDecoder(body).Decode(&info)
// 	fmt.Println("VM Status : ", info["status"].Status)
// 	return info["status"].Status

// }

// MCIS 목록에서 mcis 상태별 count map반환
func GetMcisStatusCountMap(mcisInfo model.MCISInfo) map[string]int {
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
func GetSimpleVmWithStatusCountMap(mcisInfo model.MCISInfo) ([]model.VMSimpleInfo, map[string]int) {
	// log.Println(" mcisInfo  ", index, mcisInfo)
	// vmStatusMap := make(map[string]int)
	// vmStatusMap := map[string]string{} // vmName : vmStatus

	vmStatusCountMap := map[string]int{}
	totalVmStatusCount := 0
	vmList := mcisInfo.VMs
	var vmSimpleList []model.VMSimpleInfo
	for vmIndex, vmInfo := range vmList {
		// log.Println(" vmInfo  ", vmIndex, vmInfo)
		vmStatus := util.GetVmStatus(vmInfo.Status) // lowercase로 변환

		locationInfo := vmInfo.Location
		vmLatitude := locationInfo.Latitude
		vmLongitude := locationInfo.Longitude

		log.Println(locationInfo)
		//
		vmSimpleObj := model.VMSimpleInfo{
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
func GetVMConnectionCountMap(mcisInfo model.MCISInfo) map[string]int {
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
func GetVMConnectionCountByMcis(mcisInfo model.MCISInfo) map[string]int {
	// log.Println(" mcisInfo  ", index, mcisInfo)
	vmList := mcisInfo.VMs
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
func GetVMofMcisData(nameSpaceID string, mcisID string, vmID string) (*model.VMInfo, model.WebStatus) {
	///ns/{nsId}/mcis/{mcisId}/vm/{vmId}
	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/mcis/" + mcisID + "/vm/" + vmID

	// resp, err := util.CommonHttp(url, nil, http.MethodGet)
	resp, err := util.CommonHttpWithoutParam(url, http.MethodGet)

	// defer body.Close()
	vmInfo := model.VMInfo{}
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
func McisLifeCycle(mcisLifeCycle *model.McisLifeCycle) (*model.McisLifeCycle, model.WebStatus) {
	url := util.TUMBLEBUG + "/ns/" + mcisLifeCycle.NameSpaceID + "/mcis/" + mcisLifeCycle.McisID + "?action=" + mcisLifeCycle.LifeCycleType
	//// var url = CommonURL+"/ns/"+nameSpace+"/mcis/"+mcis_id+"?action="+type
	pbytes, _ := json.Marshal(mcisLifeCycle)
	resp, err := util.CommonHttp(url, pbytes, http.MethodGet) // POST로 받기는 했으나 실제로는 Get으로 날아감.
	resultMcisLifeCycle := model.McisLifeCycle{}
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
func McisVmLifeCycle(vmLifeCycle *model.VMLifeCycle) (*model.VMLifeCycle, model.WebStatus) {

	url := util.TUMBLEBUG + "/ns/" + vmLifeCycle.NameSpaceID + "/mcis/" + vmLifeCycle.McisID + "/vm/" + vmLifeCycle.VmID + "?action=" + vmLifeCycle.LifeCycleType
	///url = CommonURL+"/ns/"+nameSpace+"/mcis/"+mcis_id+"/vm/"+vm_id+"?action="+type
	pbytes, _ := json.Marshal(vmLifeCycle)
	resp, err := util.CommonHttp(url, pbytes, http.MethodGet) // POST로 받기는 했으나 실제로는 Get으로 날아감.
	resultVmLifeCycle := model.VMLifeCycle{}
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

// VM monitoring
func GetVmMonitoring(vmMonitoring *model.VMMonitoring) (*model.VMMonitoringInfo, model.WebStatus) {
	////var url = DragonFlyURL+"/ns/"+NAMESPACE+
	//"/mcis/"+mcis_id+"/vm/"+vm_id+"/metric/"+metric+"/info?periodType="+periodType+"&statisticsCriteria="+statisticsCriteria+"&duration="+duration;
	urlParam := "periodType=" + vmMonitoring.PeriodType + "&statisticsCriteria=" + vmMonitoring.StatisticsCriteria + "&duration=" + vmMonitoring.Duration
	url := util.DRAGONFLY + "/ns/" + vmMonitoring.NameSpaceID + "/mcis/" + vmMonitoring.McisID + "/vm/" + vmMonitoring.VmID + "/metric/" + vmMonitoring.Metric + "/info?" + urlParam

	// resp, err := util.CommonHttp(url, nil, http.MethodGet)
	resp, err := util.CommonHttpWithoutParam(url, http.MethodGet)

	// defer body.Close()
	vmMonitoringInfo := model.VMMonitoringInfo{}
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
