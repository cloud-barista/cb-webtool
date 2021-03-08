package controller

// 기본화면 + Login 관련 화면 통합한 Controller

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	model "github.com/dogfootman/cb-webtool/src/model"

	echosession "github.com/go-session/echo-session"
	"github.com/labstack/echo"
)

// type ReqInfo struct {
// 	UserName string `json:"username"`
// 	Password string `json:"password"`
// }

// html 화면이동은 삭제처리
// func LoginForm(c echo.Context) error {
// 	fmt.Println("============== Login Form ===============")
// 	comURL := GetCommonURL()
// 	apiInfo := AuthenticationHandler()
// 	return c.Render(http.StatusOK, "login.html", map[string]interface{}{
// 		"comURL":  comURL,
// 		"apiInfo": apiInfo,
// 	})
// }

// func LogoutForm(c echo.Context) error {
// 	fmt.Println("============== Logout Form ===============")
// 	//comURL := GetCommonURL()
// 	return c.Render(http.StatusOK, "logout.html", nil)
// }

func Index(c echo.Context) error {
	store := echosession.FromContext(c)
	userId := os.Getenv("LoginUserId")
	store.Set(userId, nil)
	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"status":  "success",
	})
}


// Logout 처리 : 세션에서 정보 삭제
func LogoutProc(c echo.Context) error {
	store := echosession.FromContext(c)
	userId := os.Getenv("LoginUserId")
	store.Set(userId, nil)
	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"message": "You have been logged out",
		"status":  "success",
	})
}

// 사용자 등록 : 해당 사용자는 세션에 추가
func RegUser(c echo.Context) error {
// func RegUserConrtoller(c echo.Context) error {
	//comURL := GetCommonURL()

	userId := os.Getenv("LoginUserId")
	pass := os.Getenv("LoginPassword")

	store := echosession.FromContext(c)
	obj := map[string]string{
		"userid": userId,
		"password": pass,
	}
	store.Set(userId, obj)
	err := store.Save()
	if err != nil {
		return c.JSON(http.StatusServiceUnavailable, map[string]string{
			"message": "Fail",
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "SUCCESS",
		"user":    userId,
	})

}

// func name이 Controller가 들어가서 LoginProc로 변경
//func LoginController(c echo.Context) error {
func LoginProc(c echo.Context) error {
	store := echosession.FromContext(c)
	reqInfo := new(model.ReqInfo)
	// comURL := GetCommonURL()
	// apiInfo := AuthenticationHandler()
	if err := c.Bind(reqInfo); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
		})
	}
	getUserId := strings.TrimSpace(reqInfo.UserId)
	getPass := strings.TrimSpace(reqInfo.Password)
	fmt.Println("getUserId & getPass : ", getUserId, getPass)

	get, ok := store.Get(getUserId)
	fmt.Println("GEt USER:", get)
	if !ok {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
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
	if result["userid"] == getUserId && result["password"] == getPass {
		store.Set("userid", result["userid"])
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

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"status":  "200",
	})
}

// 여기서 둘다 다 되게 처리 해야 한다.
// 로그인체크와, ns check
// func LoginProc(c echo.Context) error {

// 	inputName := c.FormValue("username")
// 	inputPass := c.FormValue("password")
// 	//username에저장되어 있는 크리덴셜 정보를 가져 오자.
// 	credentialInfo := GetCredentialInfo(c, inputName)
// 	if credentialInfo.Username == inputName && credentialInfo.Password == inputPass {

// 	}
// }

