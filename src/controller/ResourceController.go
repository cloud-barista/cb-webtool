package controller

import (
	"fmt"
	"net/http"

	service "github.com/dogfootman/cb-webtool/src/service"
	"github.com/labstack/echo"
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
