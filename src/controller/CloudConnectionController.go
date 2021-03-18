package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cloud-barista/cb-webtool/src/model"
	service "github.com/cloud-barista/cb-webtool/src/service"

	"github.com/cloud-barista/cb-webtool/src/util"
	"github.com/labstack/echo"
	//"github.com/davecgh/go-spew/spew"
	echotemplate "github.com/foolin/echo-template"
	echosession "github.com/go-session/echo-session"
)

// 등록 form
// GET LIST
// GET Data
// POST 등록
// DELETE

// func CloudOSListForm
// CloudOS(Provider) 목록
func GetCloudOSList(c echo.Context) error {
	cloudOsList := service.GetCloudOSListData()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"status":  "200",
		"cloudos": cloudOsList,
	})
}

// func GetCloudOS
// func CloudOSRegProc
// func CloudOSDelProc

// func ConnectionConfigListForm  : TODO : method 명을 ConnectionConfigListForm으로 변경할 것
// func ConnectionConfigList(c echo.Context) error {
func ConnectionListForm(c echo.Context) error {
	fmt.Println("ConnectionConfigList ************ : ")

	loginInfo := service.CallLoginInfo(c)
	if loginInfo.Username == "" {
		// Login 정보가 없으므로 login화면으로
		// return c.JSON(http.StatusBadRequest, map[string]interface{}{
		// 	"message": "invalid tumblebug connection",
		// 	"status":  "403",
		// })
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	store := echosession.FromContext(c)
	// result, ok := store.Get(paramUser)
	// storedUser := result.(map[string]string)

	cloudOsList := service.GetCloudOSListData()
	store.Set("cloudos", cloudOsList)
	log.Println(" cloudOsList  ", cloudOsList)

	// connectionconfigList 가져오기
	connectionConfigDataList := service.GetConnectionConfigListData()
	store.Set("connectionconfig", connectionConfigDataList)
	log.Println(" connectionconfig  ", connectionConfigDataList)

	// regionList 가져오기
	regionList := service.GetRegionListData()
	store.Set("region", regionList)
	log.Println(" regionList  ", regionList)

	// credentialList 가져오기
	credentialList := service.GetCredentialListData()
	store.Set("credential", credentialList)
	log.Println(" credentialList  ", credentialList)

	// driverList 가져오기
	driverList := service.GetDriverListData()
	store.Set("driver", driverList)
	log.Println(" driverList  ", driverList)

	// 최신 namespacelist 가져오기
	nsList, _ := service.GetNameSpaceList()
	store.Set("namespace", nsList)
	log.Println(" nsList  ", nsList)

	return echotemplate.Render(c, http.StatusOK,
		"setting/connections/CloudConnection",
		map[string]interface{}{
			"LoginInfo":            loginInfo,
			"CloudOSList":          cloudOsList,
			"NameSpaceList":        nsList,
			"ConnectionConfigList": connectionConfigDataList,
			"RegionList":           regionList,
			"CredentialList":       credentialList,
			"DriverList":           driverList,
		})
}

// 현재 설정된 connection 목록
func GetConnectionConfigList(c echo.Context) error {

	connectionConfigDataList := service.GetConnectionConfigListData()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":          "success",
		"status":           "200",
		"connectionconfig": connectionConfigDataList,
	})
}

// func ConnectionConfigRegProc
// func ConnectionConfigDelProc

// func RegionListForm // Region 등록 form : maing 화면에서 popup형태로 뜸

// 현재 설정된 region 목록
func GetRegionList(c echo.Context) error {

	regionList := service.GetRegionListData()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"status":  "200",
		"region":  regionList,
	})
}

// region 상세정보
func GetRegion(c echo.Context) error {
	paramRegion := c.Param("region")
	resionInfo := service.GetRegionData(paramRegion)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"status":  "200",
		"Region":  resionInfo,
	})
}

// region 등록
func RegionRegProc(c echo.Context) error {
	log.Println("RegionRegProc : ")

	regionInfo := new(model.RegionInfo)
	if err := c.Bind(regionInfo); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
			"status":  "fail",
		})
	}
	log.Println(regionInfo)
	respBody, reErr := service.RegRegion(regionInfo)
	fmt.Println("=============respBody =============", respBody)
	if reErr != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid tumblebug connection",
			"status":  "403",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"status":  "200",
	})
}

// Region 삭제
func RegionDelProc(c echo.Context) error {
	log.Println("RegionRegProc : ")

	paramRegion := c.Param("region")
	log.Println(paramRegion)

	respBody, reErr := service.DelRegion(paramRegion)
	fmt.Println("=============respBody =============", respBody)
	if reErr != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid tumblebug connection",
			"status":  "403",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"status":  "200",
	})
}

// func CredentialListForm // Credential 등록 form : maing 화면에서 popup형태로 뜸

// 현재 설정된 region 목록
func GetCredentialList(c echo.Context) error {

	credentialList := service.GetCredentialListData()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"status":  "200",
		"region":  credentialList,
	})
}

// Credential 상세정보
func GetCredential(c echo.Context) error {
	paramCredential := c.Param("credential")
	credentialInfo := service.GetCredentialData(paramCredential)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":    "success",
		"status":     "200",
		"Credential": credentialInfo,
	})
}

// Credential 등록
func CredentialRegProc(c echo.Context) error {
	log.Println("CredentialRegProc : ")

	credentialInfo := new(model.CredentialInfo)
	if err := c.Bind(credentialInfo); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
			"status":  "fail",
		})
	}
	log.Println(credentialInfo)
	respBody, reErr := service.RegCredential(credentialInfo)
	fmt.Println("=============respBody =============", respBody)
	if reErr != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid tumblebug connection",
			"status":  "403",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"status":  "200",
	})
}

// Credential 삭제
func CredentialDelProc(c echo.Context) error {
	log.Println("CredentialDelProc : ")

	paramCredential := c.Param("credential")
	log.Println(paramCredential)

	respBody, reErr := service.DelCredential(paramCredential)
	fmt.Println("=============respBody =============", respBody)
	if reErr != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid tumblebug connection",
			"status":  "403",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"status":  "200",
	})
}

// func DriverListForm // Driver 등록 form : maing 화면에서 popup형태로 뜸

// 현재 설정된 driver 목록
func GetDriverList(c echo.Context) error {

	driverList := service.GetDriverListData()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"status":  "200",
		"region":  driverList,
	})
}

// Driver 조회
func GetDriver(c echo.Context) error {
	paramDriver := c.Param("driver")
	driverInfo := service.GetDriverData(paramDriver)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"status":  "200",
		"Driver":  driverInfo,
	})
}

// Driver 등록
// func DriverRegProc
func DriverRegProc(c echo.Context) error {
	log.Println("DriverRegProc : ")

	driverInfo := new(model.DriverInfo)
	if err := c.Bind(driverInfo); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
			"status":  "fail",
		})
	}
	log.Println(driverInfo)
	respBody, reErr := service.RegDriver(driverInfo)
	fmt.Println("=============respBody =============", respBody)
	if reErr != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid tumblebug connection",
			"status":  "403",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"status":  "200",
	})
}

// Driver 삭제
func DriverDelProc(c echo.Context) error {
	log.Println("DriverDelProc : ")

	paramDriver := c.Param("driver")
	log.Println(paramDriver)

	respBody, reErr := service.DelDriver(paramDriver)
	fmt.Println("=============respBody =============", respBody)
	if reErr != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid tumblebug connection",
			"status":  "403",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"status":  "200",
	})
}

// Cloud 연결정보 표시(driver)
func ConnectionList(c echo.Context) error {
	fmt.Println("ConnectionList ************ : ")

	loginInfo := service.CallLoginInfo(c)
	if loginInfo.Username == "" {
		// Login 정보가 없으므로 login화면으로
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	nsList, nsErr := service.GetNameSpaceList()
	if nsErr != nil {
		log.Println(" nsErr  ", nsErr)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid tumblebug connection",
			"status":  "403",
		})
	}

	// namespace 가 없으면 1개를 기본으로 생성한다.
	if len(nsList) == 0 {
		// create default namespace
		nsInfo := new(model.NSInfo)
		nsInfo.Name = "NS-01" // default namespace name
		nsInfo.Description = "default name space name"
		respBody, nsCreateErr := service.RegNameSpace(nsInfo)
		log.Println(" respBody  ", respBody)
		if nsCreateErr != nil {
			log.Println(" nsCreateErr  ", nsCreateErr)
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "invalid tumblebug connection",
				"status":  "403",
			})
		}

		// 처음생성했으므로 connection부터
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	} else if len(nsList) == 1 {
		defaultNameSpace := nsList[0]
		// for _, item := range nsList {
		// 	fmt.Println("ID : ", item.ID)
		// }
		loginInfo.DefaultNameSpaceID = defaultNameSpace.ID
		loginInfo.DefaultNameSpaceName = defaultNameSpace.Name
	}

	return echotemplate.Render(c, http.StatusOK,
		"setting/connections/CloudConnection",
		map[string]interface{}{
			"LoginInfo":     loginInfo,
			"NameSpaceList": nsList,
		})
	// return echotemplate.Render(c, http.StatusOK, "CloudConnection", nil)// -> file not found 남. 경로 제대로 적을 것.
}

// Driver Contorller
func DriverRegController(c echo.Context) error {

	username := c.FormValue("username")
	description := c.FormValue("description")

	fmt.Println("DriverRegController : ", username, description)
	return nil
}

// deprecated
// func DriverRegForm(c echo.Context) error {
// 	return c.Redirect(http.StatusTemporaryRedirect, "/login")
// }

// deprecated
// func DriverListForm(c echo.Context) error {
// 	return c.Redirect(http.StatusTemporaryRedirect, "/login")
// }

// deprecated
//Credential Controller
// func CredentialRegForm(c echo.Context) error {
// 	return c.Redirect(http.StatusTemporaryRedirect, "/login")
// }

// deprecated
// func CredertialListForm(c echo.Context) error {
// 	return c.Redirect(http.StatusTemporaryRedirect, "/login")
// }

// deprecated
//Region Controller
// func RegionRegForm(c echo.Context) error {
// 	return c.Redirect(http.StatusTemporaryRedirect, "/login")
// }

// deprecated
// func RegionListForm(c echo.Context) error {
// 	return c.Redirect(http.StatusTemporaryRedirect, "/login")
// }

// deprecated
//Connection Controller
// func ConnectionRegForm(c echo.Context) error {
// 	return c.Redirect(http.StatusTemporaryRedirect, "/login")
// }

// deprecated
// func ConnectionListForm(c echo.Context) error {
// 	return c.Redirect(http.StatusTemporaryRedirect, "/login")
// }

//Image Controller
func ImageRegForm(c echo.Context) error {
	comURL := service.GetCommonURL()
	apiInfo := util.AuthenticationHandler()
	if loginInfo := service.CallLoginInfo(c); loginInfo.Username != "" {
		return c.Render(http.StatusOK, "ImageRegister.html", map[string]interface{}{
			"LoginInfo": loginInfo,
			"comURL":    comURL,
			"apiInfo":   apiInfo,
		})
	}
	// return c.Redirect(http.StatusPermanentRedirect, "/login")
	return c.Redirect(http.StatusTemporaryRedirect, "/login")
}

// func ImageListForm(c echo.Context) error {
// 	comURL := service.GetCommonURL()
// 	loginInfo := service.CallLoginInfo(c)
// 	apiInfo := service.AuthenticationHandler()
// 	if loginInfo.Username != "" {
// 		nsList := service.GetRegionList()
// 		fmt.Println("REGION List : ", nsList)

// 		//spew.Dump(nsList)
// 		return c.Render(http.StatusOK, "Resources_Image.html", map[string]interface{}{
// 			"LoginInfo": loginInfo,
// 			"comURL":    comURL,
// 			"NSList":    nsList,
// 			"apiInfo":   apiInfo,
// 		})
// 	}

// 	fmt.Println("LoginInfo : ", loginInfo)
// 	return c.Redirect(http.StatusTemporaryRedirect, "/login")

// }

//VPC Controller
// func VpcRegForm(c echo.Context) error {
// 	comURL := service.GetCommonURL()
// 	apiInfo := service.AuthenticationHandler()
// 	if loginInfo := service.CallLoginInfo(c); loginInfo.Username != "" {
// 		return c.Render(http.StatusOK, "VpcRegister.html", map[string]interface{}{
// 			"LoginInfo": loginInfo,
// 			"comURL":    comURL,
// 			"apiInfo":   apiInfo,
// 		})
// 	}
// 	// return c.Redirect(http.StatusPermanentRedirect, "/login")
// 	return c.Redirect(http.StatusTemporaryRedirect, "/login")
// }

// func VpcListForm(c echo.Context) error {
// 	comURL := service.GetCommonURL()
// 	loginInfo := service.CallLoginInfo(c)
// 	apiInfo := service.AuthenticationHandler()
// 	if loginInfo.Username != "" {
// 		nsList := service.GetRegionList()
// 		fmt.Println("REGION List : ", nsList)

// 		//spew.Dump(nsList)
// 		return c.Render(http.StatusOK, "Resources_Network.html", map[string]interface{}{
// 			"LoginInfo": loginInfo,
// 			"comURL":    comURL,
// 			"NSList":    nsList,
// 			"apiInfo":   apiInfo,
// 		})
// 	}

// 	fmt.Println("LoginInfo : ", loginInfo)
// 	return c.Redirect(http.StatusTemporaryRedirect, "/login")

// }

// Controller
// func SecurityGroupRegForm(c echo.Context) error {
// 	comURL := service.GetCommonURL()
// 	apiInfo := service.AuthenticationHandler()
// 	if loginInfo := service.CallLoginInfo(c); loginInfo.Username != "" {
// 		return c.Render(http.StatusOK, "SecurityGroupRegister.html", map[string]interface{}{
// 			"LoginInfo": loginInfo,
// 			"comURL":    comURL,
// 			"apiInfo":   apiInfo,
// 		})
// 	}
// 	// return c.Redirect(http.StatusPermanentRedirect, "/login")
// 	return c.Redirect(http.StatusTemporaryRedirect, "/login")
// }

// func SecurityGroupListForm(c echo.Context) error {
// 	comURL := service.GetCommonURL()
// 	loginInfo := service.CallLoginInfo(c)
// 	apiInfo := service.AuthenticationHandler()
// 	if loginInfo.Username != "" {
// 		nsList := service.GetRegionList()
// 		fmt.Println("REGION List : ", nsList)

// 		//spew.Dump(nsList)
// 		return c.Render(http.StatusOK, "Resources_Security.html", map[string]interface{}{
// 			"LoginInfo": loginInfo,
// 			"comURL":    comURL,
// 			"NSList":    nsList,
// 			"apiInfo":   apiInfo,
// 		})
// 	}

// 	fmt.Println("LoginInfo : ", loginInfo)
// 	return c.Redirect(http.StatusTemporaryRedirect, "/login")

// }

// Controller
func SSHRegForm(c echo.Context) error {
	comURL := service.GetCommonURL()
	apiInfo := util.AuthenticationHandler()
	if loginInfo := service.CallLoginInfo(c); loginInfo.Username != "" {
		return c.Render(http.StatusOK, "SSHRegister.html", map[string]interface{}{
			"LoginInfo": loginInfo,
			"comURL":    comURL,
			"apiInfo":   apiInfo,
		})
	}
	// return c.Redirect(http.StatusPermanentRedirect, "/login")
	return c.Redirect(http.StatusTemporaryRedirect, "/login")
}

// func SSHListForm(c echo.Context) error {
// 	comURL := service.GetCommonURL()
// 	loginInfo := service.CallLoginInfo(c)
// 	apiInfo := service.AuthenticationHandler()
// 	if loginInfo.Username != "" {
// 		nsList := service.GetRegionList()
// 		fmt.Println("REGION List : ", nsList)

// 		//spew.Dump(nsList)
// 		return c.Render(http.StatusOK, "Resources_Ssh.html", map[string]interface{}{
// 			"LoginInfo": loginInfo,
// 			"comURL":    comURL,
// 			"NSList":    nsList,
// 			"apiInfo":   apiInfo,
// 		})
// 	}

// 	fmt.Println("LoginInfo : ", loginInfo)
// 	return c.Redirect(http.StatusTemporaryRedirect, "/login")

// }

// Controller
// func SpecRegForm(c echo.Context) error {
// 	comURL := service.GetCommonURL()
// 	apiInfo := service.AuthenticationHandler()
// 	if loginInfo := service.CallLoginInfo(c); loginInfo.Username != "" {
// 		return c.Render(http.StatusOK, "SpecRegister.html", map[string]interface{}{
// 			"LoginInfo": loginInfo,
// 			"comURL":    comURL,
// 			"apiInfo":   apiInfo,
// 		})
// 	}
// 	// return c.Redirect(http.StatusPermanentRedirect, "/login")
// 	return c.Redirect(http.StatusTemporaryRedirect, "/login")
// }

// func SpecListForm(c echo.Context) error {
// 	comURL := service.GetCommonURL()
// 	loginInfo := service.CallLoginInfo(c)
// 	apiInfo := service.AuthenticationHandler()
// 	if loginInfo.Username != "" {
// 		nsList := service.GetRegionList()
// 		fmt.Println("REGION List : ", nsList)

// 		//spew.Dump(nsList)
// 		return c.Render(http.StatusOK, "Resources_Spec.html", map[string]interface{}{
// 			"LoginInfo": loginInfo,
// 			"comURL":    comURL,
// 			"NSList":    nsList,
// 			"apiInfo":   apiInfo,
// 		})
// 	}

// 	fmt.Println("LoginInfo : ", loginInfo)
// 	return c.Redirect(http.StatusTemporaryRedirect, "/login")

// }
