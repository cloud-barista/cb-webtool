package controller

import (
	// "encoding/json"
	"fmt"
	"log"
	"net/http"

	// model "github.com/cloud-barista/cb-webtool/src/model"
	"github.com/cloud-barista/cb-webtool/src/model"

	// spider "github.com/cloud-barista/cb-webtool/src/model/spider"
	// "github.com/cloud-barista/cb-webtool/src/model/tumblebug"
	// tbcommon "github.com/cloud-barista/cb-webtool/src/model/tumblebug/common"
	// tbmcir "github.com/cloud-barista/cb-webtool/src/model/tumblebug/mcir"
	tbmcis "github.com/cloud-barista/cb-webtool/src/model/tumblebug/mcis"

	service "github.com/cloud-barista/cb-webtool/src/service"
	util "github.com/cloud-barista/cb-webtool/src/util"

	echotemplate "github.com/foolin/echo-template"
	echosession "github.com/go-session/echo-session"
	"github.com/labstack/echo"
)

// type SecurityGroup struct {
// 	Id []string `form:"sg"`
// }

func NlbRegForm(c echo.Context) error {
	fmt.Println("NlbRegForm ************ : ")

	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}
	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	// namespacelist 가져오기
	// nsList, _ := service.GetNameSpaceList()
	nsList, _ := service.GetStoredNameSpaceList(c)
	log.Println(" nsList  ", nsList)

	// connectionconfigList 가져오기
	cloudOsList, _ := service.GetCloudOSList()
	log.Println(" cloudOsList  ", cloudOsList)

	// regionList 가져오기
	regionList, _ := service.GetRegionList()
	log.Println(" regionList  ", regionList)

	cloudConnectionConfigInfoList, _ := service.GetCloudConnectionConfigList() // 등록된 모든 connection 정보
	log.Println("---------------------- GetCloudConnectionConfigList ", defaultNameSpaceID)

	//// namespace에 등록 된 resource 정보들 //////
	virtualMachineImageInfoList, _ := service.GetVirtualMachineImageInfoList(defaultNameSpaceID)
	vmSpecInfoList, _ := service.GetVmSpecInfoList(defaultNameSpaceID)
	vNetInfoList, _ := service.GetVnetList(defaultNameSpaceID)
	securityGroupInfoList, _ := service.GetSecurityGroupList(defaultNameSpaceID)
	sshKeyInfoList, _ := service.GetSshKeyInfoList(defaultNameSpaceID)

	return echotemplate.Render(c, http.StatusOK,
		"operation/manages/mcismng/McisCreate", // 파일명
		map[string]interface{}{
			"LoginInfo":          loginInfo,
			"DefaultNameSpaceID": defaultNameSpaceID,
			"NameSpaceList":      nsList,
			"CloudOSList":        cloudOsList,
			"RegionList":         regionList,

			"CloudConnectionConfigInfoList": cloudConnectionConfigInfoList,
			"VMImageList":                   virtualMachineImageInfoList,
			"VMSpecList":                    vmSpecInfoList,
			"VNetList":                      vNetInfoList,
			"SecurityGroupList":             securityGroupInfoList,
			"SshKeyList":                    sshKeyInfoList,
		})
}

// MCIS 관리 화면 McisListForm 에서 이름 변경 McisMngForm으로
// func McisListForm(c echo.Context) error {
func NlbMngForm(c echo.Context) error {
	fmt.Println("NlbMngForm ************ : ")

	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	defaultNameSpaceID := loginInfo.DefaultNameSpaceID
	selectedMcisID := c.QueryParam("mcisid") // Dashboard 등에서 선택한 mcis가 있는경우 mng 화면에 해당 mcis만 보이기 위해(실제로는 filterling을 위해서만 사용)

	mcisErr := model.WebStatus{}

	store := echosession.FromContext(c)

	cloudOsList, _ := service.GetCloudOSList()
	store.Set("cloudos", cloudOsList)
	log.Println(" cloudOsList  ", cloudOsList)

	// 최신 namespacelist 가져오기
	nsList, _ := service.GetNameSpaceList()
	store.Set("namespace", nsList)
	store.Save()
	log.Println(" nsList  ", nsList)

	return echotemplate.Render(c, http.StatusOK,
		"operation/services/nlbmng/NlbMng", // 파일명
		map[string]interface{}{
			"Message":            mcisErr.Message,
			"Status":             200, //mcisErr.StatusCode, // 주요한 객체 return message 를 사용
			"LoginInfo":          loginInfo,
			"CloudOSList":        cloudOsList,
			"NameSpaceList":      nsList,
			"DefaultNameSpaceID": defaultNameSpaceID,
			"SelectedMcisID":     selectedMcisID, // 선택한 MCIS ID
		})
}

// Namespace의 모든 NLB 목록 조회
func AllNlbListOfNamespace(c echo.Context) error {
	log.Println("GetAllNlbList : ")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	mcisIdList, respStatus := service.GetMcisListByID(defaultNameSpaceID, "", "")
	if respStatus.StatusCode != 200 && respStatus.StatusCode != 201 {
		return c.JSON(respStatus.StatusCode, map[string]interface{}{
			"error":  respStatus.Message,
			"status": respStatus.StatusCode,
		})
	}

	totalNlbList := []tbmcis.TbNLBInfo{}
	for _, mcisID := range mcisIdList {
		nlbList, respStatus := service.GetNlbListByOption(defaultNameSpaceID, mcisID, "")
		if respStatus.StatusCode != 200 && respStatus.StatusCode != 201 {
			continue
		}
		// mcisID set
		for _, nlb := range nlbList {
			nlb.McisID = mcisID
			log.Println("nlb.McisID : ", nlb.McisID)
			totalNlbList = append(totalNlbList, nlb)
		}
		//totalNlbList = append(totalNlbList, nlbList...) // mcisID가 set 안됨.

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":            "success",
		"status":             respStatus.StatusCode,
		"DefaultNameSpaceID": defaultNameSpaceID,
		"NlbList":            totalNlbList,
	})

}

// NLB 목록 조회
func NlbList(c echo.Context) error {
	log.Println("GetNlbList : ")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	defaultNameSpaceID := loginInfo.DefaultNameSpaceID
	mcisID := c.Param("mcisID")
	optionParam := c.QueryParam("option")

	if optionParam == "id" {
		nlbList, respStatus := service.GetNlbIdListByMcisID(defaultNameSpaceID, mcisID)
		if respStatus.StatusCode != 200 && respStatus.StatusCode != 201 {
			return c.JSON(respStatus.StatusCode, map[string]interface{}{
				"error":  respStatus.Message,
				"status": respStatus.StatusCode,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message":            "success",
			"status":             respStatus.StatusCode,
			"DefaultNameSpaceID": defaultNameSpaceID,
			"NlbList":            nlbList,
		})
	} else {
		nlbList, respStatus := service.GetNlbListByOption(defaultNameSpaceID, mcisID, optionParam)
		if respStatus.StatusCode != 200 && respStatus.StatusCode != 201 {
			return c.JSON(respStatus.StatusCode, map[string]interface{}{
				"error":  respStatus.Message,
				"status": respStatus.StatusCode,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message":            "success",
			"status":             respStatus.StatusCode,
			"DefaultNameSpaceID": defaultNameSpaceID,
			"NlbList":            nlbList,
		})
	}
}

// Nlb 등록
func NlbRegProc(c echo.Context) error {
	log.Println("NlbRegProc : ")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	nlbReq := &tbmcis.TbNLBReq{}
	if err := c.Bind(nlbReq); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
			"status":  "5001",
		})
	}
	defaultNameSpaceID := loginInfo.DefaultNameSpaceID
	mcisID := c.Param("mcisID")

	// // socket의 key 생성 : ns + 구분 + id
	taskKey := defaultNameSpaceID + "||" + "nlb" + "||" + nlbReq.Name // TODO : 공통 function으로 뺄 것.

	service.StoreWebsocketMessage(util.TASK_TYPE_MCIS, taskKey, util.MCIS_LIFECYCLE_CREATE, util.TASK_STATUS_REQUEST, c) // session에 작업내용 저장

	// // go routin, channel
	go service.RegNlbByAsync(defaultNameSpaceID, mcisID, nlbReq, c) // 오래걸리므로 요청여부만 return, 결과는 notice로 확인
	// 원래는 호출 결과를 return하나 go routine으로 바꾸면서 요청성공으로 return
	log.Println("before return")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"status":  200,
	})

}

// Nlb 삭제
func NlbDelProc(c echo.Context) error {
	log.Println("NlbDelProc : ")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	nlbID := c.Param("nlbID")
	mcisID := c.Param("mcisID")
	optionParam := c.QueryParam("option")
	log.Println("nlbID= " + nlbID)
	_, respStatus := service.DelNlb(defaultNameSpaceID, mcisID, nlbID, optionParam)
	log.Println("RegMcis service returned")
	if respStatus.StatusCode != 200 && respStatus.StatusCode != 201 {
		return c.JSON(respStatus.StatusCode, map[string]interface{}{
			"error":  respStatus.Message,
			"status": respStatus.StatusCode,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"status":  respStatus.StatusCode,
	})
}

func NlbAllDelProc(c echo.Context) error {
	log.Println("NlbAllDelProc : ")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	// store := echosession.FromContext(c)
	defaultNameSpaceID := loginInfo.DefaultNameSpaceID
	// TODO : defaultNameSpaceID 가 없으면 설정화면으로 보낼 것

	mcisID := c.Param("mcisID")
	_, respStatus := service.DelAllNlb(defaultNameSpaceID, mcisID)
	log.Println("NlbAllDelProc service returned")
	if respStatus.StatusCode != 200 && respStatus.StatusCode != 201 {
		return c.JSON(respStatus.StatusCode, map[string]interface{}{
			"error":  respStatus.Message,
			"status": respStatus.StatusCode,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"status":  respStatus.StatusCode,
	})
}

// GetNlbInfoData
// 특정 MCIS의 상세정보를 가져온다.
func NlbGet(c echo.Context) error {
	log.Println("GetNlbInfoData")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login") // 조회기능에서 바로 login화면으로 돌리지말고 return message로 하는게 낫지 않을까?
	}
	defaultNameSpaceID := loginInfo.DefaultNameSpaceID
	mcisID := c.Param("mcisID")
	nlbID := c.Param("nlbID")
	log.Println("nlbID= " + nlbID)

	resultNlbInfo, respStatus := service.GetNlbData(defaultNameSpaceID, mcisID, nlbID)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"status":  respStatus.StatusCode,
		"NlbInfo": resultNlbInfo,
	})

}

//NlbHealthGet

// Nlb에 VM 추가 등록
func NlbVmRegProc(c echo.Context) error {
	log.Println("NlbVmRegProc : ")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	// mCISInfo := &tumblebug.McisInfo{}
	vmInfo := &tbmcis.TbVmReq{}
	if err := c.Bind(vmInfo); err != nil {
		// if err := c.Bind(mCISInfoList); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
			"status":  "fail",
		})
	}
	log.Println(vmInfo)

	defaultNameSpaceID := loginInfo.DefaultNameSpaceID
	mcisID := c.Param("mcisID")

	taskKey := defaultNameSpaceID + "||" + "vm" + "||" + mcisID + "||" + vmInfo.Name
	service.StoreWebsocketMessage(util.TASK_TYPE_VM, taskKey, util.VM_LIFECYCLE_CREATE, util.TASK_STATUS_REQUEST, c) // session에 작업내용 저장

	// go 루틴 호출 : return 값은 session에 저장
	go service.AsyncRegVm(defaultNameSpaceID, mcisID, vmInfo, c)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Call success",
		"status":  200,
	})

}

// Nlb에 VM 삭제
func NlbVmDelProc(c echo.Context) error {
	log.Println("VmRegProc : ")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	// mCISInfo := &tumblebug.McisInfo{}
	vmInfo := &tbmcis.TbVmReq{}
	if err := c.Bind(vmInfo); err != nil {
		// if err := c.Bind(mCISInfoList); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
			"status":  "fail",
		})
	}
	log.Println(vmInfo)

	defaultNameSpaceID := loginInfo.DefaultNameSpaceID
	mcisID := c.Param("mcisID")

	taskKey := defaultNameSpaceID + "||" + "vm" + "||" + mcisID + "||" + vmInfo.Name
	service.StoreWebsocketMessage(util.TASK_TYPE_VM, taskKey, util.VM_LIFECYCLE_CREATE, util.TASK_STATUS_REQUEST, c) // session에 작업내용 저장

	// go 루틴 호출 : return 값은 session에 저장
	go service.AsyncRegVm(defaultNameSpaceID, mcisID, vmInfo, c)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Call success",
		"status":  200,
	})

}

func NlbHealthGet(c echo.Context) error {
	log.Println("NlbHealthGet")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login") // 조회기능에서 바로 login화면으로 돌리지말고 return message로 하는게 낫지 않을까?
	}
	defaultNameSpaceID := loginInfo.DefaultNameSpaceID
	mcisID := c.Param("mcisID")
	nlbID := c.Param("nlbID")
	log.Println("nlbID= " + nlbID)

	resultNlbInfo, respStatus := service.GetNlbHealth(defaultNameSpaceID, mcisID, nlbID)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"status":  respStatus.StatusCode,
		"NlbInfo": resultNlbInfo,
	})

}

// nlb targetgroup에 vm추가
func AddVmToNLBTargetGroup(c echo.Context) error {
	log.Println("AddVmToNLBTargetGroup : ")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	nlbTargetGroupReq := new(tbmcis.TbNLBAddRemoveVMReq)
	if err := c.Bind(nlbTargetGroupReq); err != nil {

		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
			"status":  "fail",
		})
	}

	mcisID := c.Param("mcisID")
	nlbID := c.Param("nlbID")
	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	nlbInfo, respStatus := service.AddVmToNLBTargetGroup(defaultNameSpaceID, mcisID, nlbID, nlbTargetGroupReq)
	log.Println("AddVmToNLBTargetGroup result")
	if respStatus.StatusCode != 200 && respStatus.StatusCode != 201 {

		return c.JSON(respStatus.StatusCode, map[string]interface{}{
			"error":  respStatus.Message,
			"status": respStatus.StatusCode,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": respStatus.Message,
		"status":  respStatus.StatusCode,
		"nlbInfo": nlbInfo,
	})
}

// nlb targetgroup에서 vm제거
func RemoveVmToNLBTargetGroup(c echo.Context) error {
	log.Println("RemoveVmToNLBTargetGroup : ")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.UserID == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	nlbTargetGroupReq := new(tbmcis.TbNLBAddRemoveVMReq)
	if err := c.Bind(nlbTargetGroupReq); err != nil {

		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
			"status":  "fail",
		})
	}

	mcisID := c.Param("mcisID")
	nlbID := c.Param("nlbID")
	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	respMessage, respStatus := service.RemoveVmToNLBTargetGroup(defaultNameSpaceID, mcisID, nlbID, nlbTargetGroupReq)
	log.Println("RemoveVmToNLBTargetGroup result")
	if respStatus.StatusCode != 200 && respStatus.StatusCode != 201 {

		return c.JSON(respStatus.StatusCode, map[string]interface{}{
			"error":  respStatus.Message,
			"status": respStatus.StatusCode,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": respMessage.Message,
		"status":  respStatus.StatusCode,
	})
}
