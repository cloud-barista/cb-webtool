package controller

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	//"github.com/foolin/echo-template"
	echotemplate "github.com/foolin/echo-template"
	echosession "github.com/go-session/echo-session"
	"github.com/labstack/echo"

	//db "mzc/src/databases/store"
	"github.com/cloud-barista/cb-webtool/src/service"
)

// type ReqInfo struct {
// 	Email    string `email`
// 	Password string `password`
// }
type ReqInfo struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

// func Index(c echo.Context) error {

// 	// fmt.Println("=========== DashBoard start ==============")
// 	// if loginInfo := CallLoginInfo(c); loginInfo.Username != "" {

// 	// 	return c.Redirect(http.StatusTemporaryRedirect, "/dashboard")

// 	// }
// 	fmt.Println("=========== Index Controller nothing ==============")
// 	return c.Redirect(http.StatusTemporaryRedirect, "/login")
// }

func Index(c echo.Context) error {
	fmt.Println("============== index ===============")
	user := os.Getenv("LoginUser")
	email := os.Getenv("LoginEmail")
	pass := os.Getenv("LoginPassword")

	store := echosession.FromContext(c)
	obj := map[string]string{
		"username": user,
		"email":    email,
		"password": pass,
	}
	store.Set(user, obj)
	store.Save() // 사용자정보를 따로 저장하지 않으므로 설정파일에 유저를 set.

	return c.Redirect(http.StatusTemporaryRedirect, "/login")
}

func About(c echo.Context) error {
	return c.Render(http.StatusOK, "About.html", map[string]interface{}{})
}

func Test(c echo.Context) error {
	return c.Render(http.StatusOK, "Test.html", map[string]interface{}{})
}

func LoginForm(c echo.Context) error {
	fmt.Println("============== Login Form ===============")
	return echotemplate.Render(c, http.StatusOK, "Login", nil)
	//return c.Render(http.StatusOK, "Login.html", map[string]interface{}{})
}

func LoginProc(c echo.Context) error {
	fmt.Println("============== Login proc ===============")
	store := echosession.FromContext(c)

	reqInfo := new(ReqInfo)
	comURL := service.GetCommonURL()
	apiInfo := service.AuthenticationHandler()
	if err := c.Bind(reqInfo); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
			"comURL":  comURL,
			"apiInfo": apiInfo,
		})
	}

	getUser := strings.TrimSpace(reqInfo.UserName)
	getPass := strings.TrimSpace(reqInfo.Password)
	fmt.Println("getUser & getPass : ", getUser, getPass)

	get, ok := store.Get(getUser)
	fmt.Println("Stored USER:", get)
	if !ok {
		// //return c.JSON(http.StatusNotFound, map[string]interface{}{	//404
		// //return c.JSON(http.StatusOK, map[string]interface{}{			..200
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{ //401
			"message": " 정보가 없으니 다시 등록 해라",
			"status":  "fail",
			// 	"comURL":  comURL,
			// 	"apiInfo": apiInfo,
		})
	}

	// // result := map[string]string{}
	// result := get.(map[string]string)
	// fmt.Println("result mapping : ", result)
	// for k, v := range get.(map[string]string) {
	// 	fmt.Println(k, v)
	// 	result[k] = v

	// }

	// namespace 목록조회 --> 로그인 이후로 이동할 것.
	//nsList, nsErr := service.GetNSList()
	nsList, nsErr := service.GetNameSpaceList()
	if nsErr != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid tumblebug connection",
			"status":  "403",
		})
	} else {
		store.Set("namespaceList", nsList)
		store.Save()
	}

	log.Println(" auth  ")
	// userpass, erruser := db.GetUser(getUser)
	// if erruser != nil {
	// 	log.Println(erruser)
	// 	return c.JSON(http.StatusBadRequest, map[string]interface{}{
	// 		"message": "invalid user",
	// 		"status":  "403",
	// 	})
	// }

	// if getPass == userpass {
	// 	store.Set("username", getUser)
	// 	store.Save()
	// 	log.Println(" userName---  ", getUser)
	// 	return c.JSON(http.StatusOK, map[string]interface{}{
	// 		"message":  "success",
	// 		"status":   "200",
	// 		"userInfo": getUser,
	// 	})
	// }

	// if (getUser != "" && getPass != "") && db.ValidUser(getUser, getPass) {
	if getUser != "" && getPass != "" {
		store.Set("username", getUser) //이렇게 하면 login이 1사람만 가능.
		store.Save()
		fmt.Println(" userName---  ", getUser)
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success",
			"status":  "200",
			"nsList":  nsList,
		})
	}

	return c.JSON(http.StatusBadRequest, map[string]interface{}{
		"message": "invalid user",
		"status":  "403",
	})

}

func RegUser(c echo.Context) error {
	//comURL := GetCommonURL()

	user := os.Getenv("LoginEmail")
	pass := os.Getenv("LoginPassword")

	store := echosession.FromContext(c)
	obj := map[string]string{
		"username": user,
		"password": pass,
	}
	store.Set(user, obj)
	err := store.Save()
	if err != nil {
		return c.JSON(http.StatusServiceUnavailable, map[string]string{
			"message": "Fail",
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "SUCCESS",
		"user":    user,
	})

}

func LogoutForm(c echo.Context) error {
	fmt.Println("============== Logout Form ===============")
	//comURL := GetCommonURL()
	return c.Render(http.StatusOK, "logout.html", nil)
}

func LogoutProc(c echo.Context) error {
	fmt.Println("============== Logout proc ===============")
	store := echosession.FromContext(c)

	reqInfo := new(ReqInfo)

	getUser := strings.TrimSpace(reqInfo.UserName)

	store.Set(getUser, nil)
	store.Save()
	log.Println(" auth expired ")

	return c.Render(http.StatusOK, "login.html", nil)

}
