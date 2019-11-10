package controller

import (
	"fmt"
	"net/http"

	"github.com/davecgh/go-spew/spew"

	"github.com/labstack/echo"
)

type MCISRequest struct {
	VMSpec    []string `form:"vmspec"`
	NameSpace string   `form:"namespace"`
	McisName  string   `form:"mcis_name"`
	VMName    []string `form:"vmName"`
	Provider  []string `form:"provider"`
}

func McisListForm(c echo.Context) error {
	if loginInfo := CallLoginInfo(c); loginInfo.Username != "" {
		namespace := GetNameSpaceToString(c)
		return c.Render(http.StatusOK, "MCISList.html", map[string]interface{}{
			"LoginInfo": loginInfo,
			"NameSpace": namespace,
		})

	}

	//return c.Render(http.StatusOK, "MCISlist.html", nil)
	return c.Redirect(http.StatusTemporaryRedirect, "/login")
}

func McisRegForm(c echo.Context) error {
	if loginInfo := CallLoginInfo(c); loginInfo.Username != "" {
		namespace := GetNameSpaceToString(c)
		return c.Render(http.StatusOK, "MCISRegister.html", map[string]interface{}{
			"LoginInfo": loginInfo,
			"NameSpace": namespace,
		})

	}

	return c.Redirect(http.StatusTemporaryRedirect, "/login")
}

func McisRegController(c echo.Context) error {
	m := new(MCISRequest)

	vmspec := c.FormValue("vmspec")
	namespace := c.FormValue("namespace")
	mcis_name := c.FormValue("mcis_name")
	provider := c.FormValue("provider")

	fmt.Println("namespace : ", namespace)
	fmt.Println("mcis_name : ", mcis_name)
	fmt.Println("vmSpec : ", vmspec)
	fmt.Println("provider : ", provider)

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

	spew.Dump(m)
	//return c.Redirect(http.StatusTemporaryRedirect, "/MCIS/list")
	return nil
}
