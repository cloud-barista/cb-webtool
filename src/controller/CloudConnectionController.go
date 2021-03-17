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

// CloudOS(Provider) 목록
func GetCloudOSList(c echo.Context) error {
	cloudOsList := service.GetCloudOSListData()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"status":  "200",
		"cloudos": cloudOsList,
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

// 현재 설정된 region 목록
func GetResionList(c echo.Context) error {

	regionList := service.GetRegionListData()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"status":  "200",
		"region":  regionList,
	})
}

// region 상세정보
func GetResion(c echo.Context) error {
	paramRegion := c.Param("region")
	resionInfo := service.GetRegionData(paramRegion)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"status":  "200",
		"Region":  resionInfo,
	})
}

// 현재 설정된 region 목록
func GetCredentialList(c echo.Context) error {

	credentialList := service.GetCredentialListData()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"status":  "200",
		"region":  credentialList,
	})
}

// 현재 설정된 driver 목록
func GetDriverList(c echo.Context) error {

	driverList := service.GetDriverListData()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"status":  "200",
		"region":  driverList,
	})
}

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


func RegionRegProc(c echo.Context) error {
	// RegionMoalRegionName
	// RegionModalProviderName
	// RegionModalRegionID
	// RegionModalZoneID

	regionInfo := new(model.RegionInfo)
	regionInfo.RegionName = c.Param("RegionName")
	regionInfo.ProviderName = c.Param("ProviderName")
	// nsInfo.RegionName = c.Param("RegionName")
	// nsInfo.RegionName = c.Param("RegionName")
	// if err := c.Bind(nsInfo); err != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	// }
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"status":  "200",
	})

}

// func DriverRegForm(c echo.Context) error {
// 	comURL := service.GetCommonURL()
// 	apiInfo := service.AuthenticationHandler()
// 	if loginInfo := service.CallLoginInfo(c); loginInfo.Username != "" {
// 		return c.Render(http.StatusOK, "DriverRegister.html", map[string]interface{}{
// 			"LoginInfo": loginInfo,
// 			"comURL":    comURL,
// 			"apiInfo":   apiInfo,
// 		})
// 	}
// 	return c.Redirect(http.StatusTemporaryRedirect, "/login")
// }

// func DriverListForm(c echo.Context) error {
// 	fmt.Println("=============start NsListForm =============")
// 	comURL := service.GetCommonURL()
// 	apiInfo := service.AuthenticationHandler()
// 	loginInfo := service.CallLoginInfo(c)

// 	if loginInfo.Username != "" {
// 		//nsList := service.GetDriverList()
// 		return c.Render(http.StatusOK, "DriverList.html", map[string]interface{}{
// 			"LoginInfo": loginInfo,
// 			"comURL":    comURL,
// 			"apiInfo":   apiInfo,
// 			//"NSList": nsList,
// 		})
// 	}

// 	fmt.Println("LoginInfo : ", loginInfo)

// 	return c.Redirect(http.StatusTemporaryRedirect, "/login")
// }

//Credential Controller
func CredentialRegForm(c echo.Context) error {
	comURL := service.GetCommonURL()
	apiInfo := util.AuthenticationHandler()
	if loginInfo := service.CallLoginInfo(c); loginInfo.Username != "" {
		return c.Render(http.StatusOK, "CredentialRegister.html", map[string]interface{}{
			"LoginInfo": loginInfo,
			"comURL":    comURL,
			"apiInfo":   apiInfo,
		})
	}
	return c.Redirect(http.StatusTemporaryRedirect, "/login")
}

// func CredertialListForm(c echo.Context) error {

// 	fmt.Println("=============start CredertialRegForm =============")
// 	loginInfo := service.CallLoginInfo(c)
// 	comURL := service.GetCommonURL()
// 	apiInfo := service.AuthenticationHandler()
// 	if loginInfo.Username != "" {
// 		//nsList := service.GetCredentialList()
// 		return c.Render(http.StatusOK, "CredentialList.html", map[string]interface{}{
// 			"LoginInfo": loginInfo,
// 			"comURL":    comURL,
// 			"apiInfo":   apiInfo,
// 			// "NSList": nsList,
// 		})
// 	}

// 	fmt.Println("LoginInfo : ", loginInfo)
// 	return c.Redirect(http.StatusTemporaryRedirect, "/login")

// }

//Region Controller
// func RegionRegForm(c echo.Context) error {
// 	comURL := service.GetCommonURL()
// 	apiInfo := service.AuthenticationHandler()
// 	if loginInfo := service.CallLoginInfo(c); loginInfo.Username != "" {
// 		return c.Render(http.StatusOK, "RegionRegister.html", map[string]interface{}{
// 			"LoginInfo": loginInfo,
// 			"comURL":    comURL,
// 			"apiInfo":   apiInfo,
// 		})
// 	}
// 	// return c.Redirect(http.StatusPermanentRedirect, "/login")
// 	return c.Redirect(http.StatusTemporaryRedirect, "/login")
// }

// func RegionListForm(c echo.Context) error {
// 	comURL := service.GetCommonURL()
// 	loginInfo := service.CallLoginInfo(c)
// 	apiInfo := service.AuthenticationHandler()
// 	if loginInfo.Username != "" {
// 		nsList := service.GetRegionList()
// 		fmt.Println("REGION List : ", nsList)

// 		//spew.Dump(nsList)
// 		return c.Render(http.StatusOK, "RegionList.html", map[string]interface{}{
// 			"LoginInfo": loginInfo,
// 			"comURL":    comURL,
// 			"NSList":    nsList,
// 			"apiInfo":   apiInfo,
// 		})
// 	}

// 	fmt.Println("LoginInfo : ", loginInfo)
// 	return c.Redirect(http.StatusTemporaryRedirect, "/login")

// }

//Connection Controller
// func ConnectionRegForm(c echo.Context) error {
// 	comURL := service.GetCommonURL()
// 	apiInfo := service.AuthenticationHandler()
// 	if loginInfo := service.CallLoginInfo(c); loginInfo.Username != "" {

// 		return c.Render(http.StatusOK, "ConnectionRegister.html", map[string]interface{}{
// 			"LoginInfo": loginInfo,
// 			"comURL":    comURL,
// 			"apiInfo":   apiInfo,
// 		})
// 	}
// 	return c.Redirect(http.StatusTemporaryRedirect, "/login")
// 	//return c.Render(http.StatusOK, "RegionRegister.html", nil)
// }

// func ConnectionListForm(c echo.Context) error {
// 	comURL := service.GetCommonURL()
// 	apiInfo := service.AuthenticationHandler()
// 	loginInfo := service.CallLoginInfo(c)
// 	if loginInfo.Username != "" {
// 		cList := service.GetConnectionList()
// 		fmt.Println("=============info GetConnectionList =============", cList)
// 		return c.Render(http.StatusOK, "CloudConnection.html", map[string]interface{}{
// 			// return c.Render(http.StatusOK, "ConnectionList.html", map[string]interface{}{
// 			"LoginInfo": loginInfo,
// 			"cList":     cList,
// 			"comURL":    comURL,
// 			"apiInfo":   apiInfo,
// 		})
// 	}

// 	fmt.Println("LoginInfo : ", loginInfo)

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
