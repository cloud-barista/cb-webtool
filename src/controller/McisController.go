package controller

import (
	"fmt"
	"net/http"
	"log"
	service "github.com/cloud-barista/cb-webtool/src/service"
	model "github.com/cloud-barista/cb-webtool/src/model"
	
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
	
	// 최신 namespacelist 가져오기
	nsList, _ := service.GetNameSpaceList()
	log.Println(" nsList  ", nsList)

	// status, filepath, return params
	return echotemplate.Render(c, http.StatusOK,
		"operation/manage/McisMng", // 파일명
		map[string]interface{}{
			"LoginInfo":                 loginInfo,
			"NameSpaceList":             nsList,			
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
