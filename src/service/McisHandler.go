package service

import (
	"encoding/json"
	"fmt"
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
func GetMcisList(nameSpaceID string) ([]model.MCISInfo, int) {
	// func GetMCISList(nsid string) []MCISInfo {
	url := util.TUMBLEBUG + "/ns/" + nameSpaceID + "/mcis"
	resp, err := util.CommonHttp(url, nil, http.MethodGet)
	// resp, err := util.CommonHttpWithoutParam(url, http.MethodGet)

	var respStatus int
	if err != nil {
		fmt.Println(err)
		//respStatus = 500
	}

	respBody := resp.Body
	respStatus = resp.StatusCode

	mcisList := map[string][]model.MCISInfo{}
	json.NewDecoder(respBody).Decode(&mcisList)
	fmt.Println(mcisList["mcis"])
	log.Println(respBody)
	util.DisplayResponse(resp) // 수신내용 확인

	return mcisList["mcis"], respStatus
}

func GetMCIS(nameSpaceID string, mcisID string) (model.MCISInfo, int) {
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

	if err != nil {
		fmt.Println(err)
	}
	util.DisplayResponse(resp) // 수신내용 확인

	respBody := resp.Body
	respStatus := resp.StatusCode

	mcisInfo := model.MCISInfo{}
	json.NewDecoder(respBody).Decode(&mcisInfo)
	fmt.Println(mcisInfo)

	// resultBody, err := ioutil.ReadAll(respBody)
	// if err == nil {
	// 	str := string(resultBody)
	// 	println(str)
	// }
	// pbytes, _ := json.Marshal(respBody)
	// fmt.Println(string(pbytes))

	return mcisInfo, respStatus
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
	// mcisStatusTotalMap[mcisInfo.ID] = mcisStatusMap

	return mcisStatusMap
}

// MCIS 목록에서 vm 상태별 count
func GetVMStatusCountMap(mcisInfo model.MCISInfo) map[string]int {
	vmStatusRunning := 0
	// vmStatusResuming := 0
	vmStatusInclude := 0
	vmStatusSuspended := 0
	vmStatusTerminated := 0
	vmStatusUndefined := 0
	vmStatusPartial := 0
	vmStatusEtc := 0

	// log.Println(" mcisInfo  ", index, mcisInfo)
	vmList := mcisInfo.VMs
	for vmIndex, vmInfo := range vmList {
		// log.Println(" vmInfo  ", vmIndex, vmInfo)
		vmStatus := util.GetVmStatus(vmInfo.Status)
		if vmStatus == util.VM_STATUS_RUNNING {
			vmStatusRunning++
			// }else if vmStatus == util.VM_STATUS_RESUMING {
			// 	vmStatusResuming++
		} else if vmStatus == util.VM_STATUS_INCLUDE {
			vmStatusInclude++
			// } else if vmStatus == util.VM_STATUS_SUSPENDED {
			// 	vmStatusSuspended++
		} else if vmStatus == util.VM_STATUS_TERMINATED {
			vmStatusTerminated++
			// }else if vmStatus == util.VM_STATUS_UNDEFINED {
			// 	vmStatusUndefined++
			// }else if vmStatus == util.VM_STATUS_PARTIAL {
			// 	vmStatusPartial++
		} else {
			vmStatusEtc++
			log.Println("vmStatus  ", vmIndex, vmStatus)
		}
	}
	// vmStatusMap := make(map[string]int)
	// UI에서 사칙연산이 되지 않아 controller에서 계산한 뒤 넘겨 줌.
	// vmStatusMap[util.VM_STATUS_RUNNING] = vmStatusRunning
	// vmStatusMap[util.VM_STATUS_RESUMING] = vmStatusResuming
	// vmStatusMap[util.VM_STATUS_INCLUDE] = vmStatusInclude
	// vmStatusMap[util.VM_STATUS_SUSPENDED] = vmStatusSuspended
	// vmStatusMap[util.VM_STATUS_TERMINATED] = vmStatusTerminated
	// vmStatusMap[util.VM_STATUS_UNDEFINED] = vmStatusUndefined
	// vmStatusMap[util.VM_STATUS_PARTIAL] = vmStatusPartial
	// vmStatusMap[util.VM_STATUS_ETC] = vmStatusEtc
	log.Println("mcisInfo.ID  ", mcisInfo.ID)
	// mcisIdArr[mcisIndex] = mcisInfo.ID	// 바로 넣으면 Runtime Error구만..
	// vmStatusArr[mcisIndex] = vmStatusMap

	// UI에서는 3가지로 통합하여 봄
	// vmStatusMap["RUNNING"] = vmStatusRunning
	// vmStatusMap["STOPPED"] = vmStatusInclude + vmStatusSuspended + vmStatusUndefined + vmStatusPartial + vmStatusEtc
	// vmStatusMap["TERMINATED"] = vmStatusTerminated
	// vmStatusTotalMap[mcisInfo.ID] = vmStatusMap
	// vmIdArr = append(vmIdArr, vmInfo.ID)
	// vmStatusArr = append(vmStatusArr, vmStatusMap)

	log.Println("mcisIndex  ", mcisIndex)

	vmStatusMap := make(map[string]int)
	vmStatusMap["RUNNING"] = vmStatusRunning
	vmStatusMap["STOPPED"] = vmStatusInclude + vmStatusSuspended + vmStatusUndefined + vmStatusPartial + vmStatusEtc
	vmStatusMap["TERMINATED"] = vmStatusTerminated
	vmStatusMap["TOTAL"] = vmStatusMap["RUNNING"] + vmStatusMap["STOPPED"] + vmStatusMap["TERMINATED"]

	return vmStatusMap

}

// MCIS별 connection count
func GetVMConnectionCountMap(mcisInfo model.MCISInfo) map[string]int {
	connectionCountTotal := 0
	connectionCountByMcis := 0
	vmCountTotal := 0
	vmRunningCountByMcis := 0
	vmStoppedCountByMcis := 0
	vmTerminatedCountByMcis := 0

	// log.Println(" mcisInfo  ", index, mcisInfo)
	vmList := mcisInfo.VMs
	for vmIndex, vmInfo := range vmList {
		// log.Println(" vmInfo  ", vmIndex, vmInfo)
		vmConnection := util.GetVmConnection(vmInfo.ConnectionName)
		if vmStatus == util.VM_STATUS_RUNNING {
			vmStatusRunning++
			// }else if vmStatus == util.VM_STATUS_RESUMING {
			// 	vmStatusResuming++
		} else if vmStatus == util.VM_STATUS_INCLUDE {
			vmStatusInclude++
			// } else if vmStatus == util.VM_STATUS_SUSPENDED {
			// 	vmStatusSuspended++
		} else if vmStatus == util.VM_STATUS_TERMINATED {
			vmStatusTerminated++
			// }else if vmStatus == util.VM_STATUS_UNDEFINED {
			// 	vmStatusUndefined++
			// }else if vmStatus == util.VM_STATUS_PARTIAL {
			// 	vmStatusPartial++
		} else {
			vmStatusEtc++
			log.Println("vmStatus  ", vmIndex, vmStatus)
		}
	}
	// vmStatusMap := make(map[string]int)
	// UI에서 사칙연산이 되지 않아 controller에서 계산한 뒤 넘겨 줌.
	// vmStatusMap[util.VM_STATUS_RUNNING] = vmStatusRunning
	// vmStatusMap[util.VM_STATUS_RESUMING] = vmStatusResuming
	// vmStatusMap[util.VM_STATUS_INCLUDE] = vmStatusInclude
	// vmStatusMap[util.VM_STATUS_SUSPENDED] = vmStatusSuspended
	// vmStatusMap[util.VM_STATUS_TERMINATED] = vmStatusTerminated
	// vmStatusMap[util.VM_STATUS_UNDEFINED] = vmStatusUndefined
	// vmStatusMap[util.VM_STATUS_PARTIAL] = vmStatusPartial
	// vmStatusMap[util.VM_STATUS_ETC] = vmStatusEtc
	log.Println("mcisInfo.ID  ", mcisInfo.ID)
	// mcisIdArr[mcisIndex] = mcisInfo.ID	// 바로 넣으면 Runtime Error구만..
	// vmStatusArr[mcisIndex] = vmStatusMap

	// UI에서는 3가지로 통합하여 봄
	// vmStatusMap["RUNNING"] = vmStatusRunning
	// vmStatusMap["STOPPED"] = vmStatusInclude + vmStatusSuspended + vmStatusUndefined + vmStatusPartial + vmStatusEtc
	// vmStatusMap["TERMINATED"] = vmStatusTerminated
	// vmStatusTotalMap[mcisInfo.ID] = vmStatusMap
	// vmIdArr = append(vmIdArr, vmInfo.ID)
	// vmStatusArr = append(vmStatusArr, vmStatusMap)

	log.Println("mcisIndex  ", mcisIndex)

	vmStatusMap := make(map[string]int)
	vmStatusMap["RUNNING"] = vmStatusRunning
	vmStatusMap["STOPPED"] = vmStatusInclude + vmStatusSuspended + vmStatusUndefined + vmStatusPartial + vmStatusEtc
	vmStatusMap["TERMINATED"] = vmStatusTerminated

	return vmStatusMap

}
