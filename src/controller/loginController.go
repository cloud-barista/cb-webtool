package controller

import (
	"fmt"
	"net/http"
	"strings"

	echosession "github.com/go-session/echo-session"
	"github.com/labstack/echo"
)

type ReqInfo struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func RegUserConrtoller(c echo.Context) error {

	reqInfo := new(ReqInfo)
	if err := c.Bind(reqInfo); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "fail",
		})
	}
	user := reqInfo.UserName
	pass := reqInfo.Password
	fmt.Println("c.Request : ", user, pass)
	store := echosession.FromContext(c)
	get, ok := store.Get(user)
	fmt.Println(get)
	obj := map[string]string{
		"username": user,
		// "namespace": MakeNameSpace(user),
		"password": pass,
	}
	if !ok {

		store.Set(user, obj)
		err := store.Save()
		if err != nil {
			return c.JSON(http.StatusServiceUnavailable, map[string]string{
				"message": "Fail",
			})
		}
		return c.JSON(http.StatusOK, map[string]string{
			"message": "SUCCESS",
		})
	} else {
		return c.JSON(301, map[string]string{
			"message": "already register",
		})
	}

}

func LoginController(c echo.Context) error {
	store := echosession.FromContext(c)
	reqInfo := new(ReqInfo)
	if err := c.Bind(reqInfo); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "fail",
		})
	}
	getUser := strings.TrimSpace(reqInfo.UserName)
	getPass := strings.TrimSpace(reqInfo.Password)
	fmt.Println("getUser & getPass : ", getUser, getPass)

	get, ok := store.Get(getUser)
	fmt.Println("GEt USER:", get)
	if !ok {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": " 정보가 없으니 다시 등록 해라",
			"status":  "fail",
		})
	}
	//result := map[string]string{}
	result := get.(map[string]string)
	fmt.Println("result mapping : ", result)
	// for k, v := range get.(map[string]string) {
	// 	fmt.Println(k, v)
	// 	result[k] = v

	// }

	fmt.Println("result : ", result["password"])
	if result["password"] == getPass && result["username"] == getUser {
		store.Set("username", result["username"])
		store.Save()
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Login Success",
			//	"nameSpace": result["namespace"],
			"status": "success",
		})
	} else {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "wrong password of ID",
			"status":  "fail",
		})
	}

	// var result map[string]string
	// for k, item := range getObj {
	// 	fmt.Println("GetItem : ", item)
	// 	result[k] = item
	// }
	fmt.Println("getObj :", get)
	// if sesEmail := session.Get(getUser); sesEmail != nil {
	// 	if sesEmail == getUser {

	// 	}
	// }

	return c.JSON(http.StatusOK, map[string]string{
		"message": "success",
		"status":  "200",
	})
}
