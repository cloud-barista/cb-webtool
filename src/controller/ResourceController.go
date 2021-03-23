package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cloud-barista/cb-webtool/src/model"
	service "github.com/cloud-barista/cb-webtool/src/service"
	"github.com/labstack/echo"

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
	// log.Println(" loginInfo  ", loginInfo)
	// if defaultNameSpaceID == "" {
	// 	log.Println(" loginInfo2  ", loginInfo)
	// 	return c.JSON(http.StatusBadRequest, map[string]interface{}{
	// 		"message": "select namespace",
	// 		"status":  "403",
	// 	})
	// }

	store := echosession.FromContext(c)

	cloudOsList := service.GetCloudOSListData()
	store.Set("cloudos", cloudOsList)
	log.Println(" cloudOsList  ", cloudOsList)

	// 최신 namespacelist 가져오기
	nsList, _ := service.GetNameSpaceList()
	store.Set("namespace", nsList)
	log.Println(" nsList  ", nsList)

	vNetInfoList, vNetErr := service.GetVnetList(defaultNameSpaceID)
	if vNetErr != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid tumblebug connection",
			"status":  "403",
		})
	}
	log.Println("VNetList", vNetInfoList)
	//spew.Dump(nsList)
	// return c.Render(http.StatusOK, "NetworkMng.html", map[string]interface{}{
	// 	"LoginInfo": loginInfo,
	// 	"NSList":    regionList,
	// })

	return echotemplate.Render(c, http.StatusOK,
		"setting/resources/NetworkMng", // 파일명
		map[string]interface{}{
			"LoginInfo":     loginInfo,
			"CloudOSList":   cloudOsList,
			"NameSpaceList": nsList,
			"VNetList":      vNetInfoList,
		})
}

//////////////////////////////
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
	vNetInfoList, vNetErr := service.GetVnetList(defaultNameSpaceID)
	if vNetErr != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid tumblebug connection",
			"status":  "403",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":            "success",
		"status":             "200",
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

	// store := echosession.FromContext(c)
	defaultNameSpaceID := loginInfo.DefaultNameSpaceID

	paramVNetID := c.Param("vNetID")

	vNetInfo, vNetErr := service.GetVpcData(defaultNameSpaceID, paramVNetID)
	if vNetErr != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid tumblebug connection",
			"status":  "403",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "success",
		"status":   "200",
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

	vNetInfo := new(model.VNetInfo)
	if err := c.Bind(vNetInfo); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
			"status":  "fail",
		})
	}
	log.Println(vNetInfo)
	respBody, reErr := service.RegVpc(defaultNameSpaceID, vNetInfo)
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

	respBody, reErr := service.DelVpc(defaultNameSpaceID, paramVNetID)
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
