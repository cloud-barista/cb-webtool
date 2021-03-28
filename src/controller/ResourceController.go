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

	cloudOsList , _ := service.GetCloudOSListData()
	store.Set("cloudos", cloudOsList)
	log.Println(" cloudOsList  ", cloudOsList)

	// 최신 namespacelist 가져오기
	nsList, _ := service.GetNameSpaceList()
	store.Set("namespace", nsList)
	log.Println(" nsList  ", nsList)

	vNetInfoList, vNetStatus := service.GetVnetList(defaultNameSpaceID)
	// if vNetErr != nil {
	if vNetStatus != util.HTTP_CALL_SUCCESS && vNetStatus != util.HTTP_POST_SUCCESS {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid tumblebug connection",
			"status":  vNetStatus,
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
	vNetInfoList, vNetStatus := service.GetVnetList(defaultNameSpaceID)
	// if vNetErr != nil {
	if vNetStatus != util.HTTP_CALL_SUCCESS && vNetStatus != util.HTTP_POST_SUCCESS {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid tumblebug connection",
			"status":  vNetStatus,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":            "success",
		"status":             vNetStatus,
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
	vNetInfo, vNetStatus := service.GetVpcData(defaultNameSpaceID, paramVNetID)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success",
		"status":   vNetStatus,
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

///////////////////
func SecirityGroupMngForm(c echo.Context) error {
	fmt.Println("SecirityGroupMngForm ************ : ")

	loginInfo := service.CallLoginInfo(c)
	if loginInfo.Username == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	store := echosession.FromContext(c)

	cloudOsList, _ := service.GetCloudOSListData()
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

// Vpc 상세정보
func GetSecirityGroupData(c echo.Context) error {
	log.Println("GetSecirityGroupData : ")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.Username == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	paramVNetID := c.Param("vNetID")
	securityGroupInfo, vNetStatus := service.GetSecurityGroupData(defaultNameSpaceID, paramVNetID)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":       "success",
		"status":        vNetStatus,
		"SecurityGroup": securityGroupInfo,
	})
}

// Vpc 등록 :
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

	resultVNetInfo, respStatus := service.RegSecurityGroup(defaultNameSpaceID, securityGroupRegInfo)

	if respStatus != util.HTTP_CALL_SUCCESS && respStatus != util.HTTP_POST_SUCCESS {
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
func SecirityGroupDelProc(c echo.Context) error {
	log.Println("SecirityGroupDelProc : ")

	loginInfo := service.CallLoginInfo(c)
	if loginInfo.Username == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	// store := echosession.FromContext(c)
	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	paramVNetID := c.Param("vNetID")

	respBody, respStatus := service.DelSecurityGroup(defaultNameSpaceID, paramVNetID)
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
