package controller

import (
	"fmt"
	"net/http"
	"log"
	service "github.com/cloud-barista/cb-webtool/src/service"
	model "github.com/cloud-barista/cb-webtool/src/model"
	util "github.com/cloud-barista/cb-webtool/src/util"

	"github.com/labstack/echo"
	echotemplate "github.com/foolin/echo-template"
	// echosession "github.com/go-session/echo-session"
)



// type SecurityGroup struct {
// 	Id []string `form:"sg"`
// }

// MCIS 관리 화면 McisListForm 에서 이름 변경 McisMngForm으로
// func McisListForm(c echo.Context) error {
func McisMngForm(c echo.Context) error {
	// comURL := service.GetCommonURL()
	// apiInfo := util.AuthenticationHandler()
	// if loginInfo := service.CallLoginInfo(c); loginInfo.Username != "" {
	// 	namespace := service.GetNameSpaceToString(c)
	// 	if namespace != "" {
	// 		return c.Render(http.StatusOK, "Manage_Mcis.html", map[string]interface{}{
	// 			"LoginInfo": loginInfo,
	// 			"NameSpace": namespace,
	// 			"comURL":    comURL,
	// 			"apiInfo":   apiInfo,
	// 		})
	// 	} else {
	// 		return c.Redirect(http.StatusTemporaryRedirect, "/NS/reg")
	// 	}
	// }

	// //return c.Render(http.StatusOK, "MCISlist.html", nil)
	// return c.Redirect(http.StatusTemporaryRedirect, "/login")

	fmt.Println("McisMngForm ************ : ")

	loginInfo := service.CallLoginInfo(c)
	if loginInfo.Username == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}
	
	// store := echosession.FromContext(c)
	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	// 최신 namespacelist 가져오기
	nsList, _ := service.GetNameSpaceList()
	log.Println(" nsList  ", nsList)

	mcisList, _ := service.GetMcisList(defaultNameSpaceID)
	log.Println(" mcisList  ", mcisList)
	totoalMcisLength := len(mcisList)
	
	log.Println(" totoalMcisLength  ", totoalMcisLength)
	
	var mcisIdArr []string
	vmStatusArr := []map[string]int{}
	for mcisIndex, mcisInfo := range mcisList {
		// log.Println(" mcisInfo  ", index, mcisInfo)
		vmList := mcisInfo.VMs
		vmStatusRunning := 0
		vmStatusResuming := 0
		vmStatusInclude := 0
		vmStatusSuspended := 0
		vmStatusTerminated := 0
		vmStatusUndefined := 0
		vmStatusEtc := 0		
		for vmIndex, vmInfo := range vmList {
			// log.Println(" vmInfo  ", vmIndex, vmInfo)
			vmStatus := vmInfo.Status

			if vmStatus == util.VM_STATUS_RUNNING {
				vmStatusRunning++				
			}else if vmStatus == util.VM_STATUS_RESUMING {
				vmStatusResuming++
			}else if vmStatus == util.VM_STATUS_INCLUDE {
				vmStatusInclude++
			}else if vmStatus == util.VM_STATUS_SUSPENDED {
				vmStatusSuspended++
			}else if vmStatus == util.VM_STATUS_TERMINATED {
				vmStatusTerminated++
			}else if vmStatus == util.VM_STATUS_UNDEFINED {
				vmStatusUndefined++
			}else {
				vmStatusEtc++
				log.Println("vmStatus  ", vmIndex, vmStatus)
			}								
		}
		vmStatusMap := make(map[string]int)
		vmStatusMap[util.VM_STATUS_RUNNING] = vmStatusRunning
		vmStatusMap[util.VM_STATUS_RESUMING] = vmStatusResuming	
		vmStatusMap[util.VM_STATUS_INCLUDE] = vmStatusInclude
		vmStatusMap[util.VM_STATUS_SUSPENDED] = vmStatusSuspended
		vmStatusMap[util.VM_STATUS_TERMINATED] = vmStatusTerminated
		vmStatusMap[util.VM_STATUS_UNDEFINED] = vmStatusUndefined
		vmStatusMap[util.VM_STATUS_ETC] = vmStatusEtc
		log.Println("mcisInfo.ID  ", mcisInfo.ID)
		// mcisIdArr[mcisIndex] = mcisInfo.ID	// 바로 넣으면 Runtime Error구만..
		// vmStatusArr[mcisIndex] = vmStatusMap
		mcisIdArr = append(mcisIdArr, mcisInfo.ID)
		vmStatusArr = append(vmStatusArr, vmStatusMap)
		
		log.Println("mcisIndex  ", mcisIndex)
	}

	


	// status, filepath, return params
	return echotemplate.Render(c, http.StatusOK,
		"operation/manage/McisMng", // 파일명
		map[string]interface{}{
			"LoginInfo":                 loginInfo,
			"DefaultNameSpaceID": defaultNameSpaceID,
			"NameSpaceList":             nsList,
			"McisList":  mcisList,
			"McisIDList":  mcisIdArr,			
			"VMStatusList":  vmStatusArr,
		})

}

func GetMcisList(c echo.Context) error {
	log.Println("GetMcisList : ")
	loginInfo := service.CallLoginInfo(c)
	if loginInfo.Username == "" {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	// store := echosession.FromContext(c)
	defaultNameSpaceID := loginInfo.DefaultNameSpaceID
	// TODO : defaultNameSpaceID 가 없으면 설정화면으로 보낼 것
	mcisList, respStatus := service.GetMcisList(defaultNameSpaceID)
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
		"McisList":  mcisList,
	})
}
// func McisListFormWithParam(c echo.Context) error {
// 	mcis_id := c.Param("mcis_id")
// 	mcis_name := c.Param("mcis_name")
// 	comURL := service.GetCommonURL()
// 	apiInfo := service.AuthenticationHandler()
// 	if mcis_id == "" && mcis_name == "" {
// 		mcis_id = ""
// 		mcis_name = ""
// 	}
// 	if loginInfo := service.CallLoginInfo(c); loginInfo.Username != "" {
// 		namespace := service.GetNameSpaceToString(c)
// 		return c.Render(http.StatusOK, "Manage_Mcis.html", map[string]interface{}{
// 			"LoginInfo": loginInfo,
// 			"NameSpace": namespace,
// 			"McisID":    mcis_id,
// 			"McisName":  mcis_name,
// 			"comURL":    comURL,
// 			"apiInfo":   apiInfo,
// 		})

// 	}

// 	//return c.Render(http.StatusOK, "MCISlist.html", nil)
// 	return c.Redirect(http.StatusTemporaryRedirect, "/login")
// }

// func VMAddForm(c echo.Context) error {
// 	mcis_id := c.Param("mcis_id")
// 	mcis_name := c.Param("mcis_name")
// 	comURL := service.GetCommonURL()
// 	apiInfo := service.AuthenticationHandler()
// 	if mcis_id == "" && mcis_name == "" {
// 		mcis_id = ""
// 		mcis_name = ""
// 	}
// 	if loginInfo := service.CallLoginInfo(c); loginInfo.Username != "" {
// 		namespace := service.GetNameSpaceToString(c)
// 		return c.Render(http.StatusOK, "Manage_Create_VM.html", map[string]interface{}{
// 			"LoginInfo": loginInfo,
// 			"NameSpace": namespace,
// 			"McisID":    mcis_id,
// 			"McisName":  mcis_name,
// 			"comURL":    comURL,
// 			"apiInfo":   apiInfo,
// 		})

// 	}

// 	//return c.Render(http.StatusOK, "MCISlist.html", nil)
// 	return c.Redirect(http.StatusTemporaryRedirect, "/login")
// }

// func McisRegForm(c echo.Context) error {
// 	comURL := service.GetCommonURL()
// 	apiInfo := service.AuthenticationHandler()
// 	if loginInfo := service.CallLoginInfo(c); loginInfo.Username != "" {
// 		namespace := service.GetNameSpaceToString(c)
// 		return c.Render(http.StatusOK, "Manage_Create_Mcis.html", map[string]interface{}{
// 			"LoginInfo": loginInfo,
// 			"NameSpace": namespace,
// 			"comURL":    comURL,
// 			"apiInfo":   apiInfo,
// 		})

// 	}

// 	return c.Redirect(http.StatusTemporaryRedirect, "/login")
// }

func McisRegController(c echo.Context) error {
	m := new(model.MCISRequest)

	vmspec := c.FormValue("vmspec")
	namespace := c.FormValue("namespace")
	mcis_name := c.FormValue("mcis_name")
	provider := c.FormValue("provider")
	sg := c.FormValue("sg")

	fmt.Println("namespace : ", namespace)
	fmt.Println("mcis_name : ", mcis_name)
	fmt.Println("vmSpec : ", vmspec)
	fmt.Println("provider : ", provider)
	fmt.Println("sg : ", sg)

	if err := c.Bind(m); err != nil {
		fmt.Println("bind Error")
		return err
	}
	fmt.Println("Bind Form : ", m)
	fmt.Println("nameSPace:", m.NameSpace)
	fmt.Println("vmName 0 : ", m.VMName[0])
	fmt.Println("vmName 1 : ", m.VMName[1])
	fmt.Println("vmSpec 0 : ", m.VMSpec[0])
	fmt.Println("vmspec 1 : ", m.VMSpec[1])

	//spew.Dump(m)
	//return c.Redirect(http.StatusTemporaryRedirect, "/MCIS/list")
	return nil
}
