package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

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
		return c.Render(http.StatusOK, "MCISRegister.html", map[string]interface{}{
			"LoginInfo": loginInfo,
		})

	}

	return c.Redirect(http.StatusTemporaryRedirect, "/login")
}
