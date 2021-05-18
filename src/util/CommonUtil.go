package util

// 공통으로 사용할 function 정의
import (
	// "encoding/base64"
	// "fmt"
	// "io"
	// "io/ioutil"
	"log"
	// "net/http"
	"net/url"
	// "os"
	"reflect"
	"strconv"
	"strings"
	// "time"
	// "bytes"
	// "encoding/json"
	// "math"
	// "io/ioutil"
	// echosession "github.com/go-session/echo-session"
	// "github.com/labstack/echo"
	// "github.com/cloud-barista/cb-webtool/src/model"
)

// providerName 소문자로
func GetProviderName(provider string) string {
	return strings.ToLower(provider)
}

// MCIS 상태값의 앞부분만 사용. 소문자로
func GetMcisStatus(mcisStatus string) string {
	statusArr := strings.Split(mcisStatus, "-")
	returnStatus := strings.ToLower(statusArr[0])

	if returnStatus == MCIS_STATUS_RUNNING {
		returnStatus = "running"
	} else if returnStatus == MCIS_STATUS_INCLUDE {
		returnStatus = "stop"
	} else if returnStatus == MCIS_STATUS_SUSPENDED {
		returnStatus = "stop"
	} else if returnStatus == MCIS_STATUS_TERMINATED {
		returnStatus = "terminate"
	} else if returnStatus == MCIS_STATUS_PARTIAL {
		returnStatus = "stop"
	} else if returnStatus == MCIS_STATUS_ETC {
		returnStatus = "stop"
	} else {
		returnStatus = "stop"
	}
	return returnStatus
}

// VM 상태를 UI에서 표현하는 방식으로 변경
func GetVmStatus(vmStatus string) string {
	returnVmStatus := strings.ToLower(vmStatus) // 소문자로 변환

	if returnVmStatus == VM_STATUS_RUNNING {
		returnVmStatus = VM_STATUS_RUNNING
		// }else if vmStatus == util.VM_STATUS_RESUMING {
		// 	vmStatusResuming++
	} else if returnVmStatus == VM_STATUS_INCLUDE {
		returnVmStatus = VM_STATUS_INCLUDE
	} else if returnVmStatus == VM_STATUS_SUSPENDED {
		returnVmStatus = VM_STATUS_SUSPENDED
	} else if returnVmStatus == VM_STATUS_TERMINATED {
		returnVmStatus = VM_STATUS_TERMINATED
		// }else if returnVmStatus == util.VM_STATUS_UNDEFINED {
		// 	vmStatusUndefined++
		// }else if returnVmStatus == util.VM_STATUS_PARTIAL {
		// 	vmStatusPartial++
	} else {
		returnVmStatus = VM_STATUS_ETC
	}
	return returnVmStatus
}

func GetVmConnectionName(vmConnectionName string) string {
	return strings.ToLower(vmConnectionName)
}

/////////// Map Control 참고

// var mcisIdArr []string
// var vmIdArr []string
// // mcisStatusTotalMap := make(map[string]map[string]int)
// // vmStatusTotalMap := make(map[string]map[string]int)

// mcisStatusRunning := 0
// mcisStatusStopped := 0
// mcisStatusTerminated := 0

// vmStatusRunning := 0
// // vmStatusResuming := 0
// vmStatusInclude := 0
// vmStatusSuspended := 0
// vmStatusTerminated := 0
// vmStatusUndefined := 0
// vmStatusPartial := 0
// vmStatusEtc := 0

// for mcisIndex, mcisInfo := range mcisList {
// 	// log.Println(" mcisInfo  ", index, mcisInfo)
// 	mcisIdArr = append(mcisIdArr, mcisInfo.ID)
// 	mcisStatus := util.GetMcisStatus(mcisInfo.Status)
// 	if mcisStatus == util.MCIS_STATUS_RUNNING {
// 		mcisStatusRunning++
// 	} else if mcisStatus == util.MCIS_STATUS_TERMINATED {
// 		mcisStatusTerminated++
// 	} else {
// 		mcisStatusStopped++
// 	}

// 	vmList := mcisInfo.Vms

// 	for vmIndex, vmInfo := range vmList {
// 		// log.Println(" vmInfo  ", vmIndex, vmInfo)
// 		vmStatus := util.GetVmStatus(vmInfo.Status)
// 		vmIdArr = append(vmIdArr, vmInfo.ID)
// 		if vmStatus == util.VM_STATUS_RUNNING {
// 			vmStatusRunning++
// 			// }else if vmStatus == util.VM_STATUS_RESUMING {
// 			// 	vmStatusResuming++
// 		} else if vmStatus == util.VM_STATUS_INCLUDE {
// 			vmStatusInclude++
// 			// } else if vmStatus == util.VM_STATUS_SUSPENDED {
// 			// 	vmStatusSuspended++
// 		} else if vmStatus == util.VM_STATUS_TERMINATED {
// 			vmStatusTerminated++
// 			// }else if vmStatus == util.VM_STATUS_UNDEFINED {
// 			// 	vmStatusUndefined++
// 			// }else if vmStatus == util.VM_STATUS_PARTIAL {
// 			// 	vmStatusPartial++
// 		} else {
// 			vmStatusEtc++
// 			log.Println("vmStatus  ", vmIndex, vmStatus)
// 		}
// 	}
// 	// vmStatusMap := make(map[string]int)
// 	// UI에서 사칙연산이 되지 않아 controller에서 계산한 뒤 넘겨 줌.
// 	// vmStatusMap[util.VM_STATUS_RUNNING] = vmStatusRunning
// 	// vmStatusMap[util.VM_STATUS_RESUMING] = vmStatusResuming
// 	// vmStatusMap[util.VM_STATUS_INCLUDE] = vmStatusInclude
// 	// vmStatusMap[util.VM_STATUS_SUSPENDED] = vmStatusSuspended
// 	// vmStatusMap[util.VM_STATUS_TERMINATED] = vmStatusTerminated
// 	// vmStatusMap[util.VM_STATUS_UNDEFINED] = vmStatusUndefined
// 	// vmStatusMap[util.VM_STATUS_PARTIAL] = vmStatusPartial
// 	// vmStatusMap[util.VM_STATUS_ETC] = vmStatusEtc
// 	log.Println("mcisInfo.ID  ", mcisInfo.ID)
// 	// mcisIdArr[mcisIndex] = mcisInfo.ID	// 바로 넣으면 Runtime Error구만..
// 	// vmStatusArr[mcisIndex] = vmStatusMap

// 	// UI에서는 3가지로 통합하여 봄
// 	// vmStatusMap["RUNNING"] = vmStatusRunning
// 	// vmStatusMap["STOPPED"] = vmStatusInclude + vmStatusSuspended + vmStatusUndefined + vmStatusPartial + vmStatusEtc
// 	// vmStatusMap["TERMINATED"] = vmStatusTerminated
// 	// vmStatusTotalMap[mcisInfo.ID] = vmStatusMap
// 	// vmIdArr = append(vmIdArr, vmInfo.ID)
// 	// vmStatusArr = append(vmStatusArr, vmStatusMap)

// 	log.Println("mcisIndex  ", mcisIndex)
// }
// mcisStatusMap := make(map[string]int)
// mcisStatusMap["RUNNING"] = mcisStatusRunning
// mcisStatusMap["STOPPED"] = mcisStatusStopped
// mcisStatusMap["TERMINATED"] = mcisStatusTerminated
// // mcisStatusTotalMap[mcisInfo.ID] = mcisStatusMap

// vmStatusMap := make(map[string]int)
// vmStatusMap["RUNNING"] = vmStatusRunning
// vmStatusMap["STOPPED"] = vmStatusInclude + vmStatusSuspended + vmStatusUndefined + vmStatusPartial + vmStatusEtc
// vmStatusMap["TERMINATED"] = vmStatusTerminated
// // vmStatusTotalMap[mcisInfo.ID] = vmStatusMap
///////////

func StructToUrlValues(i interface{}) (values url.Values) {
	values = url.Values{}
	iVal := reflect.ValueOf(i).Elem()
	typ := iVal.Type()
	for i := 0; i < iVal.NumField(); i++ {
		f := iVal.Field(i)
		// You ca use tags here...
		// tag := typ.Field(i).Tag.Get("tagname")
		// Convert each type into a string for the url.Values string map
		var v string
		switch f.Interface().(type) {
		case int, int8, int16, int32, int64:
			v = strconv.FormatInt(f.Int(), 10)
		case uint, uint8, uint16, uint32, uint64:
			v = strconv.FormatUint(f.Uint(), 10)
		case float32:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 32)
		case float64:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 64)
		case []byte:
			v = string(f.Bytes())
		case string:
			v = f.String()
		}
		values.Set(typ.Field(i).Name, v)
	}
	return
}

// func Struct2Map(obj interface{}) map[string]interface{} {
// 	t := reflect.TypeOf(obj)
// 	v := reflect.ValueOf(obj)

// 	var data = make(map[string]interface{})
// 	for i := 0; i < t.NumField(); i++ {
// 		data[t.Field(i).Name] = v.Field(i).Interface()
// 	}
// 	return data
// }

func Struct2Map(obj interface{}) map[string]interface{} {
	var data = make(map[string]interface{})

	target := reflect.ValueOf(obj)
	elements := target.Elem()

	log.Printf("Type: %s\n", target.Type()) // 구조체 타입명

	for i := 0; i < elements.NumField(); i++ {
		mValue := elements.Field(i)
		mType := elements.Type().Field(i)
		tag := mType.Tag

		log.Printf("%10s %10s ==> %10v, json: %10s\n",
			mType.Name,         // 이름
			mType.Type,         // 타입
			mValue.Interface(), // 값
			tag.Get("json"))    // json 태그

		data[mType.Name] = mValue.Interface()
	}
	return data
}

func Struct2MapString(obj interface{}) map[string]string {
	var data = make(map[string]string)

	target := reflect.ValueOf(obj)
	target = reflect.Indirect(target)
	elements := target.Elem()

	var stringType = reflect.TypeOf("")
	log.Printf("Type: %s\n", target.Type()) // 구조체 타입명

	for i := 0; i < elements.NumField(); i++ {
		mValue := elements.Field(i)
		mType := elements.Type().Field(i)
		tag := mType.Tag

		log.Printf("%10s %10s ==> %10v, json: %10s\n",
			mType.Name,         // 이름
			mType.Type,         // 타입
			mValue.Interface(), // 값
			tag.Get("json"))    // json 태그

		log.Println("vv ", mValue.Convert(stringType))
		// data[mType.Name] = string(mValue.Interface().(int))

		data[mType.Name] = "1"
	}
	return data
}

// func StructToMap(i interface{}) (values url.Values) {
// 	values = map[string]
// 	iVal := reflect.ValueOf(i).Elem()
// 	typ := iVal.Type()
// 	for i := 0; i < iVal.NumField(); i++ {
// 		f := iVal.Field(i)
// 		// You ca use tags here...
// 		// tag := typ.Field(i).Tag.Get("tagname")
// 		// Convert each type into a string for the url.Values string map
// 		var v string
// 		switch f.Interface().(type) {
// 		case int, int8, int16, int32, int64:
// 			v = strconv.FormatInt(f.Int(), 10)
// 		case uint, uint8, uint16, uint32, uint64:
// 			v = strconv.FormatUint(f.Uint(), 10)
// 		case float32:
// 			v = strconv.FormatFloat(f.Float(), 'f', 4, 32)
// 		case float64:
// 			v = strconv.FormatFloat(f.Float(), 'f', 4, 64)
// 		case []byte:
// 			v = string(f.Bytes())
// 		case string:
// 			v = f.String()
// 		}
// 		values.Set(typ.Field(i).Name, v)
// 	}
// 	return
// }
