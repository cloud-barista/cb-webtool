package controller

import (
	"encoding/json"
	"fmt"
	"github.com/cloud-barista/cb-webtool/src/model"
	service "github.com/cloud-barista/cb-webtool/src/service"
	util "github.com/cloud-barista/cb-webtool/src/util"

	"github.com/labstack/echo"
	// "io/ioutil"
	"log"
	"net/http"
	//"github.com/davecgh/go-spew/spew"
	echotemplate "github.com/foolin/echo-template"
	echosession "github.com/go-session/echo-session"
)

func ResourceBoard(c echo.Context) error {
	fmt.Println("=========== ResourceBoard start ==============")
	comURL := service.GetCommonURL()
	if loginInfo := service.CallLoginInfo(c); loginInfo.Username != "" {
		nameSpace := service.GetNameSpaceToString(c)
		fmt.Println("Namespace : ", nameSpace)
		return c.Render(http.StatusOK, "ResourceBoard.html", map[string]interface{}{
			"LoginInfo": loginInfo,
			"NameSpace": nameSpace,
			"comURL":    comURL,
		})
	}

	return c.Redirect(http.StatusTemporaryRedirect, "/login")
}

// func VpcListForm(c echo.Context) error {
func VpcMngForm(c echo.Context) error {
	fmt.Println("ConnectionConfigList ************ : ")

	loginInfo := service.CallLoginInfo(c)
	if loginInfo.Username == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	store := echosession.FromContext(c)

	cloudOsList , _ := service.GetCloudOSList()
	store.Set("cloudos", cloudOsList)
	log.Println(" cloudOsList  ", cloudOsList)

	// 최신 namespacelist 가져오기
	nsList, _ := service.GetNameSpaceList()
	store.Set("namespace", nsList)
	log.Println(" nsList  ", nsList)

	vNetInfoList, respStatus := service.GetVnetList(defaultNameSpaceID)
	// if vNetErr != nil {
	if respStatus != util.HTTP_CALL_SUCCESS && respStatus != util.HTTP_POST_SUCCESS {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid tumblebug connection",
			"status":  respStatus,
		})
	}
	log.Println("VNetList", vNetInfoList)

	return echotemplate.Render(c, http.StatusOK,
		"setting/resources/NetworkMng", // 파일명
		map[string]interface{}{
			"LoginInfo":     loginInfo,
			"CloudOSList":   cloudOsList,
			"NameSpaceList": nsList,
			"VNetList":      vNetInfoList,
		})
}

func GetVpcList(c echo.Context) error {
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.Username == "" {
		// Login 정보가 없으므로 login화면으로
		// return c.JSON(http.StatusBadRequest, map[string]interface{}{
		// 	"message": "invalid tumblebug connection",
		// 	"status":  "403",
		// })
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	// store := echosession.FromContext(c)
	defaultNameSpaceID := loginInfo.DefaultNameSpaceID
	// TODO : defaultNameSpaceID 가 없으면 설정화면으로 보낼 것
	vNetInfoList, respStatus := service.GetVnetList(defaultNameSpaceID)
	// if vNetErr != nil {
	if respStatus != util.HTTP_CALL_SUCCESS && respStatus != util.HTTP_POST_SUCCESS {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid tumblebug connection",
			"status":  respStatus,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":            "success",
		"status":             respStatus,
		"DefaultNameSpaceID": defaultNameSpaceID,
		"VNetList":           vNetInfoList,
	})
}

// Vpc 상세정보
func GetVpcData(c echo.Context) error {
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.Username == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	paramVNetID := c.Param("vNetID")
	vNetInfo, respStatus := service.GetVpcData(defaultNameSpaceID, paramVNetID)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success",
		"status":   respStatus,
		"VNetInfo": vNetInfo,
	})
}

// Vpc 등록 :
func VpcRegProc(c echo.Context) error {
	log.Println("VpcRegProc : ")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.Username == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	vNetRegInfo := new(model.VNetRegInfo)
	if err := c.Bind(vNetRegInfo); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
			"status":  "fail",
		})
	}
	// log.Println(vNetRegInfo)
	resultVNetInfo, respStatus := service.RegVpc(defaultNameSpaceID, vNetRegInfo)
	// respBody, respStatus := service.RegVpc(defaultNameSpaceID, vNetRegInfo)
	// fmt.Println("=============respStatus =============", respStatus)
	// fmt.Println("=============respBody ===============", respBody)

	// if reErr != nil {
	if respStatus != util.HTTP_CALL_SUCCESS && respStatus != util.HTTP_POST_SUCCESS {
		// resultBody, err := ioutil.ReadAll(respBody)
		// if err == nil {
		// 	str := string(resultBody)
		// 	println(str)
		// }
		// pbytes, _ := json.Marshal(respBody)
		// fmt.Println(string(pbytes))
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid tumblebug connection",
			"status":  respStatus,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success",
		"status":   respStatus,
		"VNetInfo": resultVNetInfo,
	})
}

// 삭제
func VpcDelProc(c echo.Context) error {
	log.Println("ConnectionConfigDelProc : ")

	loginInfo := service.CallLoginInfo(c)
	if loginInfo.Username == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	// store := echosession.FromContext(c)
	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	paramVNetID := c.Param("vNetID")

	respBody, respStatus := service.DelVpc(defaultNameSpaceID, paramVNetID)
	fmt.Println("=============respBody =============", respBody)

	// if reErr != nil {
	if respStatus != util.HTTP_CALL_SUCCESS && respStatus != util.HTTP_POST_SUCCESS {
		// resultBody, err := ioutil.ReadAll(respBody)
		// if err == nil {
		// 	str := string(resultBody)
		// 	println(str)
		// }
		pbytes, _ := json.Marshal(respBody)
		fmt.Println(string(pbytes))

		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid tumblebug connection",
			"status":  respStatus,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"status":  respStatus,
	})
}

// SecurityGroup 관리 화면
func SecirityGroupMngForm(c echo.Context) error {
	fmt.Println("SecirityGroupMngForm ************ : ")

	loginInfo := service.CallLoginInfo(c)
	if loginInfo.Username == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	store := echosession.FromContext(c)

	cloudOsList, _ := service.GetCloudOSList()
	store.Set("cloudos", cloudOsList)
	log.Println(" cloudOsList  ", cloudOsList)

	// 최신 namespacelist 가져오기
	nsList, _ := service.GetNameSpaceList()
	store.Set("namespace", nsList)
	log.Println(" nsList  ", nsList)

	securityGroupInfoList, respStatus := service.GetSecurityGroupList(defaultNameSpaceID)
	if respStatus != util.HTTP_CALL_SUCCESS && respStatus != util.HTTP_POST_SUCCESS {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid tumblebug connection",
			"status":  respStatus,
		})
	}
	log.Println("securityGroupInfoList", securityGroupInfoList)

	return echotemplate.Render(c, http.StatusOK,
		"setting/resources/SecurityGroupMng", // 파일명
		map[string]interface{}{
			"LoginInfo":         loginInfo,
			"CloudOSList":       cloudOsList,
			"NameSpaceList":     nsList,
			"SecurityGroupList": securityGroupInfoList,
		})
}

// SecurityGroup 목록
func GetSecirityGroupList(c echo.Context) error {
	log.Println("GetSecirityGroupList : ")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.Username == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	// store := echosession.FromContext(c)
	defaultNameSpaceID := loginInfo.DefaultNameSpaceID
	// TODO : defaultNameSpaceID 가 없으면 설정화면으로 보낼 것
	securityGroupInfoList, respStatus := service.GetSecurityGroupList(defaultNameSpaceID)
	// if vNetErr != nil {
	if respStatus != util.HTTP_CALL_SUCCESS && respStatus != util.HTTP_POST_SUCCESS {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid tumblebug connection",
			"status":  respStatus,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":            "success",
		"status":             respStatus,
		"DefaultNameSpaceID": defaultNameSpaceID,
		"SecurityGroupList":  securityGroupInfoList,
	})
}

// 상세정보
func GetSecirityGroupData(c echo.Context) error {
	log.Println("GetSecirityGroupData : ")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.Username == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	paramSecurityGroupID := c.Param("securityGroupID")
	securityGroupInfo, respStatus := service.GetSecurityGroupData(defaultNameSpaceID, paramSecurityGroupID)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":       "success",
		"status":        respStatus,
		"SecurityGroupInfo": securityGroupInfo,
	})
}

// 등록 :
func SecirityGroupRegProc(c echo.Context) error {
	log.Println("SecirityGroupRegProc : ")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.Username == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	securityGroupRegInfo := new(model.SecurityGroupRegInfo)
	if err := c.Bind(securityGroupRegInfo); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
			"status":  "fail",
		})
	}

	resultSecurityGroupInfo, respStatus := service.RegSecurityGroup(defaultNameSpaceID, securityGroupRegInfo)

	if respStatus != util.HTTP_CALL_SUCCESS && respStatus != util.HTTP_POST_SUCCESS {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid tumblebug connection",
			"status":  respStatus,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success",
		"status":   respStatus,
		"SecurityGroupInfo": resultSecurityGroupInfo,
	})
}

// 삭제
func SecirityGroupDelProc(c echo.Context) error {
	log.Println("SecirityGroupDelProc : ")

	loginInfo := service.CallLoginInfo(c)
	if loginInfo.Username == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	// store := echosession.FromContext(c)
	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	paramSecurityGroupID := c.Param("securityGroupID")

	respBody, respStatus := service.DelSecurityGroup(defaultNameSpaceID, paramSecurityGroupID)
	fmt.Println("=============respBody =============", respBody)

	// if reErr != nil {
	if respStatus != util.HTTP_CALL_SUCCESS && respStatus != util.HTTP_POST_SUCCESS {
		// resultBody, err := ioutil.ReadAll(respBody)
		// if err == nil {
		// 	str := string(resultBody)
		// 	println(str)
		// }
		pbytes, _ := json.Marshal(respBody)
		fmt.Println(string(pbytes))

		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid tumblebug connection",
			"status":  respStatus,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"status":  respStatus,
	})
}


func SshKeyMngForm(c echo.Context) error {
	fmt.Println("SshKeyMngForm ************ : ")

	loginInfo := service.CallLoginInfo(c)
	if loginInfo.Username == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	store := echosession.FromContext(c)

	cloudOsList, _ := service.GetCloudOSList()
	store.Set("cloudos", cloudOsList)
	log.Println(" cloudOsList  ", cloudOsList)

	// 최신 namespacelist 가져오기
	nsList, _ := service.GetNameSpaceList()
	store.Set("namespace", nsList)
	log.Println(" nsList  ", nsList)

	sshKeyInfoList, respStatus := service.GetSshKeyInfoList(defaultNameSpaceID)
	if respStatus != util.HTTP_CALL_SUCCESS && respStatus != util.HTTP_POST_SUCCESS {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid tumblebug connection",
			"status":  respStatus,
		})
	}
	log.Println("sshKeyInfoList", sshKeyInfoList)

	return echotemplate.Render(c, http.StatusOK,
		"setting/resources/SshKeyMng", // 파일명
		map[string]interface{}{
			"LoginInfo":         loginInfo,
			"CloudOSList":       cloudOsList,
			"NameSpaceList":     nsList,
			"SshKeyList": sshKeyInfoList,
		})
}

func GetSshKeyList(c echo.Context) error {
	log.Println("GetSshKeyList : ")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.Username == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	// store := echosession.FromContext(c)
	defaultNameSpaceID := loginInfo.DefaultNameSpaceID
	// TODO : defaultNameSpaceID 가 없으면 설정화면으로 보낼 것
	sshKeyInfoList, respStatus := service.GetSshKeyInfoList(defaultNameSpaceID)
	// if vNetErr != nil {
	if respStatus != util.HTTP_CALL_SUCCESS && respStatus != util.HTTP_POST_SUCCESS {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid tumblebug connection",
			"status":  respStatus,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":            "success",
		"status":             respStatus,
		"DefaultNameSpaceID": defaultNameSpaceID,
		"SshKeyList":  sshKeyInfoList,
	})
}

// SSHKey 상세정보
func GetSshKeyData(c echo.Context) error {
	log.Println("GetSshKeyData : ")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.Username == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	paramSshKey := c.Param("sshKeyID")
	sshKeyInfo, respStatus := service.GetSshKeyData(defaultNameSpaceID, paramSshKey)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":       "success",
		"status":        respStatus,
		"SshKeyInfo": sshKeyInfo,
	})
}

// SSHKey 등록 :
func SshKeyRegProc(c echo.Context) error {
	log.Println("SshKeyRegProc : ")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.Username == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	sshKeyRegInfo := new(model.SshKeyRegInfo)
	if err := c.Bind(sshKeyRegInfo); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
			"status":  "fail",
		})
	}

	resultSshKeyInfo, respStatus := service.RegSshKey(defaultNameSpaceID, sshKeyRegInfo)

	if respStatus != util.HTTP_CALL_SUCCESS && respStatus != util.HTTP_POST_SUCCESS {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid tumblebug connection",
			"status":  respStatus,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success",
		"status":   respStatus,
		"SshKeyInfo": resultSshKeyInfo,
	})
}

// 삭제
func SshKeyDelProc(c echo.Context) error {
	log.Println("SshKeyDelProc : ")

	loginInfo := service.CallLoginInfo(c)
	if loginInfo.Username == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	// store := echosession.FromContext(c)
	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	paramSshKeyID := c.Param("sshKeyID")

	respBody, respStatus := service.DelSshKey(defaultNameSpaceID, paramSshKeyID)
	fmt.Println("=============respBody =============", respBody)

	// if reErr != nil {
	if respStatus != util.HTTP_CALL_SUCCESS && respStatus != util.HTTP_POST_SUCCESS {
		// resultBody, err := ioutil.ReadAll(respBody)
		// if err == nil {
		// 	str := string(resultBody)
		// 	println(str)
		// }
		pbytes, _ := json.Marshal(respBody)
		fmt.Println(string(pbytes))

		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid tumblebug connection",
			"status":  respStatus,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"status":  respStatus,
	})
}

// VirtualMachine Image 등록 form
func VirtualMachineImageMngForm(c echo.Context) error {
	fmt.Println("VirtualMachineImageMngForm ************ : ")

	loginInfo := service.CallLoginInfo(c)
	if loginInfo.Username == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	store := echosession.FromContext(c)

	cloudOsList, _ := service.GetCloudOSList()
	store.Set("cloudos", cloudOsList)
	log.Println(" cloudOsList  ", cloudOsList)

	// 최신 namespacelist 가져오기
	nsList, _ := service.GetNameSpaceList()
	store.Set("namespace", nsList)
	log.Println(" nsList  ", nsList)

	virtualMachineImageInfoList, respStatus := service.GetVirtualMachineImageInfoList(defaultNameSpaceID)
	if respStatus != util.HTTP_CALL_SUCCESS && respStatus != util.HTTP_POST_SUCCESS {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid tumblebug connection",
			"status":  respStatus,
		})
	}
	log.Println("VirtualMachineImageInfoList", virtualMachineImageInfoList)

	return echotemplate.Render(c, http.StatusOK,
		"setting/resources/VirtualMachineImageMng", // 파일명
		map[string]interface{}{
			"LoginInfo":         loginInfo,
			"CloudOSList":       cloudOsList,
			"NameSpaceList":     nsList,
			"VirtualMachineImageList": virtualMachineImageInfoList,
		})
}

func GetVirtualMachineImageList(c echo.Context) error {
	log.Println("GetVirtualMachineImageList : ")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.Username == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	// store := echosession.FromContext(c)
	defaultNameSpaceID := loginInfo.DefaultNameSpaceID
	// TODO : defaultNameSpaceID 가 없으면 설정화면으로 보낼 것
	virtualMachineImageInfoList, respStatus := service.GetVirtualMachineImageInfoList(defaultNameSpaceID)
	// if vNetErr != nil {
	if respStatus != util.HTTP_CALL_SUCCESS && respStatus != util.HTTP_POST_SUCCESS {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid tumblebug connection",
			"status":  respStatus,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":            "success",
		"status":             respStatus,
		"DefaultNameSpaceID": defaultNameSpaceID,
		"VirtualMachineImageList":  virtualMachineImageInfoList,
	})
}

// VirtualMachineImage 상세정보
func GetVirtualMachineImageData(c echo.Context) error {
	log.Println("GetVirtualMachineImageData : ")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.Username == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	paramVirtualMachineImage := c.Param("imageID")
	virtualMachineImageInfo, respStatus := service.GetVirtualMachineImageData(defaultNameSpaceID, paramVirtualMachineImage)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":       "success",
		"status":        respStatus,
		"VirtualMachineImageInfo": virtualMachineImageInfo,
	})
}

// VirtualMachineImage 등록 :
func VirtualMachineImageRegProc(c echo.Context) error {
	log.Println("VirtualMachineImageRegProc : ")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.Username == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	virtualMachineImageRegInfo := new(model.VirtualMachineImageRegInfo)
	if err := c.Bind(virtualMachineImageRegInfo); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
			"status":  "fail",
		})
	}

	resultVirtualMachineImageInfo, respStatus := service.RegVirtualMachineImage(defaultNameSpaceID, virtualMachineImageRegInfo)
	// todo : return message 조치 필요. 중복 등 에러났을 때 message 표시가 제대로 되지 않음
	if respStatus != util.HTTP_CALL_SUCCESS && respStatus != util.HTTP_POST_SUCCESS {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid tumblebug connection",
			"status":  respStatus,
		})
	}
	// respBody := resp.Body
	// respStatusCode := resp.StatusCode
	// respStatus := resp.Status
	// log.Println("respStatusCode = ", respStatusCode)
	// log.Println("respStatus = ", respStatus)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success",
		"status":   respStatus,
		"VirtualMachineImageInfo": resultVirtualMachineImageInfo,
	})
}

// 삭제
func VirtualMachineImageDelProc(c echo.Context) error {
	log.Println("VirtualMachineImageDelProc : ")

	loginInfo := service.CallLoginInfo(c)
	if loginInfo.Username == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	// store := echosession.FromContext(c)
	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	paramVirtualMachineImageID := c.Param("imageID")

	respBody, respStatus := service.DelVirtualMachineImage(defaultNameSpaceID, paramVirtualMachineImageID)
	fmt.Println("=============respBody =============", respBody)

	// if reErr != nil {
	if respStatus != util.HTTP_CALL_SUCCESS && respStatus != util.HTTP_POST_SUCCESS {
		// resultBody, err := ioutil.ReadAll(respBody)
		// if err == nil {
		// 	str := string(resultBody)
		// 	println(str)
		// }
		pbytes, _ := json.Marshal(respBody)
		fmt.Println(string(pbytes))

		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid tumblebug connection",
			"status":  respStatus,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"status":  respStatus,
	})
}

// lookupImage 목록
func LookupVirtualMachineImageList(c echo.Context) error {
	log.Println("GetVirtualMachineImageList : ")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.Username == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	// store := echosession.FromContext(c)
	defaultNameSpaceID := loginInfo.DefaultNameSpaceID
	// TODO : defaultNameSpaceID 가 없으면 설정화면으로 보낼 것
	virtualMachineImageInfoList, respStatus := service.LookupVirtualMachineImageList()
	// if vNetErr != nil {
	if respStatus != util.HTTP_CALL_SUCCESS && respStatus != util.HTTP_POST_SUCCESS {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid tumblebug connection",
			"status":  respStatus,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":            "success",
		"status":             respStatus,
		"DefaultNameSpaceID": defaultNameSpaceID,
		"VirtualMachineImageList":  virtualMachineImageInfoList,
	})
}

// lookupImage 상세정보
func LookupVirtualMachineImageData(c echo.Context) error {
	log.Println("LookupVirtualMachineImageData : ")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.Username == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	paramVirtualMachineImage := c.Param("imageID")
	virtualMachineImageInfo, respStatus := service.LookupVirtualMachineImageData(paramVirtualMachineImage)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":       "success",
		"status":        respStatus,
		"VirtualMachineImageInfo": virtualMachineImageInfo,
	})
}

// lookupImage 상세정보
func SearchVirtualMachineImageList(c echo.Context) error {
	log.Println("SearchVirtualMachineImageList : ")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.Username == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	paramKeywords := c.Param("keywords")
	virtualMachineImageInfoList, respStatus := service.SearchVirtualMachineImageList(defaultNameSpaceID, paramKeywords)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":       "success",
		"status":        respStatus,
		"VirtualMachineImageList": virtualMachineImageInfoList,
	})
}

// InstanceSpecMng 등록 form
func InstanceSpecMngForm(c echo.Context) error {
	fmt.Println("InstanceSpecMngForm ************ : ")

	loginInfo := service.CallLoginInfo(c)
	if loginInfo.Username == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	store := echosession.FromContext(c)

	cloudOsList, _ := service.GetCloudOSList()
	store.Set("cloudos", cloudOsList)
	log.Println(" cloudOsList  ", cloudOsList)

	// 최신 namespacelist 가져오기
	nsList, _ := service.GetNameSpaceList()
	store.Set("namespace", nsList)
	log.Println(" nsList  ", nsList)

	instanceSpecInfoList, respStatus := service.GetInstanceSpecInfoList(defaultNameSpaceID)
	if respStatus != util.HTTP_CALL_SUCCESS && respStatus != util.HTTP_POST_SUCCESS {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid tumblebug connection",
			"status":  respStatus,
		})
	}
	log.Println("instanceSpecInfoList", instanceSpecInfoList)

	return echotemplate.Render(c, http.StatusOK,
		"setting/resources/InstanceSpecMng", // 파일명
		map[string]interface{}{
			"LoginInfo":         loginInfo,
			"CloudOSList":       cloudOsList,
			"NameSpaceList":     nsList,
			"InstanceSpecList": instanceSpecInfoList,
		})
}

func GetInstanceSpecList(c echo.Context) error {
	log.Println("GetInstanceSpecList : ")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.Username == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	// store := echosession.FromContext(c)
	defaultNameSpaceID := loginInfo.DefaultNameSpaceID
	// TODO : defaultNameSpaceID 가 없으면 설정화면으로 보낼 것
	instanceSpecInfoList, respStatus := service.GetInstanceSpecInfoList(defaultNameSpaceID)
	// if vNetErr != nil {
	if respStatus != util.HTTP_CALL_SUCCESS && respStatus != util.HTTP_POST_SUCCESS {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid tumblebug connection",
			"status":  respStatus,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":            "success",
		"status":             respStatus,
		"DefaultNameSpaceID": defaultNameSpaceID,
		"InstanceSpecList":  instanceSpecInfoList,
	})
}

// InstanceSpec 상세정보
func GetInstanceSpecData(c echo.Context) error {
	log.Println("GetInstanceSpecData : ")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.Username == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	paramInstanceSpec := c.Param("specID")
	instanceSpecInfo, respStatus := service.GetInstanceSpecData(defaultNameSpaceID, paramInstanceSpec)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":       "success",
		"status":        respStatus,
		"InstanceSpecInfo": instanceSpecInfo,
	})
}

// InstanceSpec 등록 :
func InstanceSpecRegProc(c echo.Context) error {
	log.Println("InstanceSpecRegProc : ")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.Username == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	instanceSpecRegInfo := new(model.InstanceSpecRegInfo)
	if err := c.Bind(instanceSpecRegInfo); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
			"status":  "fail",
		})
	}

	resultInstanceSpecInfo, respStatus := service.RegInstanceSpec(defaultNameSpaceID, instanceSpecRegInfo)
	// todo : return message 조치 필요. 중복 등 에러났을 때 message 표시가 제대로 되지 않음
	if respStatus != util.HTTP_CALL_SUCCESS && respStatus != util.HTTP_POST_SUCCESS {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid tumblebug connection",
			"status":  respStatus,
		})
	}
	// respBody := resp.Body
	// respStatusCode := resp.StatusCode
	// respStatus := resp.Status
	// log.Println("respStatusCode = ", respStatusCode)
	// log.Println("respStatus = ", respStatus)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success",
		"status":   respStatus,
		"InstanceSpecInfo": resultInstanceSpecInfo,
	})
}

// 삭제
func InstanceSpecDelProc(c echo.Context) error {
	log.Println("InstanceSpecDelProc : ")

	loginInfo := service.CallLoginInfo(c)
	if loginInfo.Username == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	// store := echosession.FromContext(c)
	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	paramInstanceSpecID := c.Param("specID")

	respBody, respStatus := service.DelInstanceSpec(defaultNameSpaceID, paramInstanceSpecID)
	fmt.Println("=============respBody =============", respBody)

	// if reErr != nil {
	if respStatus != util.HTTP_CALL_SUCCESS && respStatus != util.HTTP_POST_SUCCESS {
		// resultBody, err := ioutil.ReadAll(respBody)
		// if err == nil {
		// 	str := string(resultBody)
		// 	println(str)
		// }
		pbytes, _ := json.Marshal(respBody)
		fmt.Println(string(pbytes))

		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid tumblebug connection",
			"status":  respStatus,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"status":  respStatus,
	})
}

// lookupImage 목록
func LookupInstanceSpecList(c echo.Context) error {
	log.Println("GetInstanceSpecList : ")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.Username == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	// store := echosession.FromContext(c)
	defaultNameSpaceID := loginInfo.DefaultNameSpaceID
	// TODO : defaultNameSpaceID 가 없으면 설정화면으로 보낼 것
	instanceSpecInfoList, respStatus := service.LookupInstanceSpecList()
	// if vNetErr != nil {
	if respStatus != util.HTTP_CALL_SUCCESS && respStatus != util.HTTP_POST_SUCCESS {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid tumblebug connection",
			"status":  respStatus,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":            "success",
		"status":             respStatus,
		"DefaultNameSpaceID": defaultNameSpaceID,
		"InstanceSpecList":  instanceSpecInfoList,
	})
}

// lookupImage 상세정보
func LookupInstanceSpecData(c echo.Context) error {
	log.Println("LookupInstanceSpecData : ")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.Username == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	paramInstanceSpec := c.Param("specID")
	instanceSpecInfo, respStatus := service.LookupInstanceSpecData(paramInstanceSpec)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":       "success",
		"status":        respStatus,
		"InstanceSpecInfo": instanceSpecInfo,
	})
}

// resourcesGroup.PUT("/instancespec/put/:specID", controller.InstanceSpecPutProc)	// RegProc _ SshKey 같이 앞으로 넘길까
// resourcesGroup.POST("/instancespec/filterspecs", controller.FilterInstanceSpecList)
// resourcesGroup.POST("/instancespec/filterspecsbyrange", controller.FilterInstanceSpecListByRange)
