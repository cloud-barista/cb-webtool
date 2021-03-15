package controller

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	//"github.com/foolin/echo-template"
	echotemplate "github.com/foolin/echo-template"
	echosession "github.com/go-session/echo-session"
	
	"github.com/dgrijalva/jwt-go"

	"github.com/twinj/uuid"
	
	"github.com/labstack/echo"

	//db "mzc/src/databases/store"
	// "github.com/cloud-barista/cb-webtool/src/service"
	"github.com/cloud-barista/cb-webtool/src/model"
)

type TokenDetails struct {
	AccessToken    string
	RefreshToken   string
	AccessUuid     string
	RefreshUuid    string
	AtExpires      int64
	RtExpires      int64
}
// type ReqInfo struct {
// 	Email    string `email`
// 	Password string `password`
// }


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
		"defaultnamespage": "",
		"accesstoken" : "",
		"refreshtoken" : "",
	}
	store.Set(user, obj)
	store.Save() // 사용자정보를 따로 저장하지 않으므로 설정파일에 유저를 set.
	fmt.Println("user : ", user)
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
	return echotemplate.Render(c, http.StatusOK, "auth/Login", nil)
	//return c.Render(http.StatusOK, "Login.html", map[string]interface{}{})
}

func LoginProc(c echo.Context) error {
	fmt.Println("============== Login proc ===============")
	store := echosession.FromContext(c)

	reqInfo := new(model.ReqInfo)
	if err := c.Bind(reqInfo); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "fail",
			"status":  "fail",
		})
	}

	paramUser := strings.TrimSpace(reqInfo.UserName)
	// paramEmail := strings.TrimSpace(reqInfo.Email)
	paramPass := strings.TrimSpace(reqInfo.Password)
	fmt.Println("paramUser & getPass : ", paramUser, paramPass)
	
	// echoSession에서 가져오기
	result, ok := store.Get(paramUser)
	
	if !ok {
		log.Println(" login proc err  ", ok)
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{ //401
			"message": " 정보가 없으니 다시 등록바랍니다.",
			"status":  "fail",
		})		
	}
	storedUser := result.(map[string]string)
	fmt.Println("Stored USER:", storedUser)
	if paramUser != storedUser["username"] && paramUser != storedUser["password"] {
		log.Println(" invalid id or pass  ")
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{ //401
			"message": "invalid user or password",
			"status":  "fail",
		})
	}
	
	newToken, createTokenErr := createToken(paramUser)
	if createTokenErr != nil {
		log.Println(" login proc err  ", createTokenErr)
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{ //401
			"message": " 로그인 처리 요류",
			"status":  "fail",
		})		
	}

	// "accesstoken" : "",
	// 	"refreshtoken" : "",
	// td.RefreshToken
	storedUser["accesstoken"] = newToken.AccessToken
	storedUser["refreshtoken"] = newToken.RefreshToken
	// store.Set(paramUser, storedUser)
	// store.Save()
	
	//////// 현재구조에서는 nsList 부분을 포함해야 함. 
	//////// TODO : post로 넘기고 login success일 때 namespace 등 설정 popup이 아닌 page로 넘기는게 좋을 듯.
	//contentType := c.Request().Header.Get("Content-Type")
	//AccessToken := c.Request().Header.Get("AccessToken")
	//RefreshToken := c.Request().Header.Get("RefreshToken")
	// return c.JSON(http.StatusOK, map[string]interface{}{
	// 	"message": "success",
	// 	"status":  "200",
	// 	"AccessToken": newToken.AccessToken,
	// 	// "RefreshToken": newToken.RefreshToken,
	// })
	/////////----------------------------         

	// // result := map[string]string{}
	// result := get.(map[string]string)
	// fmt.Println("result mapping : ", result)
	// for k, v := range get.(map[string]string) {
	// 	fmt.Println(k, v)
	// 	result[k] = v
	// }

	// namespace 목록조회 --> 로그인 이후로 이동할 것.
	
		//nsList, nsErr := service.GetNSList()
	// nsList, nsErr := service.GetNameSpaceList()
	// if nsErr != nil {
	// 	log.Println(" nsErr  ", nsErr)
	// 	return c.JSON(http.StatusBadRequest, map[string]interface{}{
	// 		"message": "invalid tumblebug connection",
	// 		"status":  "403",
	// 	})
	// }

	// if len(nsList) == 0 {
	// 	// create default namespace
	// 	nsInfo := new(model.NSInfo)
	// 	nsInfo.ID = "NS-01"
	// 	nsInfo.Name = "NS-01"	// default namespace name
	// 	nsInfo.Description = "default name space name"
	// 	respBody, nsCreateErr := service.RegNameSpace(nsInfo)
	// 	log.Println(" respBody  ", respBody)
	// 	if nsCreateErr != nil {
	// 		log.Println(" nsCreateErr  ", nsCreateErr)
	// 		return c.JSON(http.StatusBadRequest, map[string]interface{}{
	// 			"message": "invalid tumblebug connection",
	// 			"status":  "403",
	// 		})
	// 	}

	// 	storedUser["defaultnamespage"] = nsInfo.ID		
		
	// 	// 저장 성공하면 namespace 목록 조회
	// 	nsList2, nsErr2 := service.GetNameSpaceList()
	// 	if nsErr2 != nil {
	// 		log.Println(" nsErr2  ", nsErr2)
	// 		return c.JSON(http.StatusBadRequest, map[string]interface{}{
	// 			"message": "invalid tumblebug connection",
	// 			"status":  "403",
	// 		})
	// 	}
	// 	log.Println("nsList2  ", nsList2)
	// 	nsList = nsList2
	// }
	// log.Println("nsList  ", nsList)
	// store.Set("namespaceList", nsList)// 이게 유효한가?? 쓸모없을 듯
	store.Set(paramUser, storedUser)
	store.Save()

	loginInfo := model.LoginInfo{
		Username: paramUser,
		//Username:  result["username"],
		// DefaultNameSpace: result["namespace"],
	}
	
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"status":  "200",
		"LoginInfo": loginInfo,
		// "nsList":  nsList,
	})

	// return c.JSON(http.StatusBadRequest, map[string]interface{}{
	// 	"message": "invalid user",
	// 	"status":  "403",
	// })


}



// ----------- 로그인이 성공하면 Namespace가 없으면 생성 ----------/
// ----------- Name Space가 1개 있으면 Dashboard로 이동 ----------/
// ----------- Name Space가 1개 이상 있으면 Dashboard로 이동 및 Namespace선택 Modal 띄우기 ----------/
// ----------- MCIS가 등록되어있지 않으면 등록화면으로 ----------/
func LoginProcess(c echo.Context) error {
	store := echosession.FromContext(c)

	paramUser := c.FormValue("username")	
	paramPass := c.FormValue("password")
	
	// reqInfo := new(model.ReqInfo)
	// if err := c.Bind(reqInfo); err != nil {
	// 	return c.Redirect(http.StatusTemporaryRedirect, "/login")		
	// }

	// paramUser := strings.TrimSpace(reqInfo.UserName)
	// // paramEmail := strings.TrimSpace(reqInfo.Email)
	// paramPass := strings.TrimSpace(reqInfo.Password)
	fmt.Println("paramUser & paramPass : ", paramUser, paramPass)
	

	// echoSession에서 가져오기
	result, ok := store.Get(paramUser)

	if !ok {
		log.Println(" login proc err  ", ok)
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}
	storedUser := result.(map[string]string)
	fmt.Println("Stored USER:", storedUser)
	if paramUser != storedUser["username"] || paramUser != storedUser["password"] {
		log.Println(" invalid id or pass  ")
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}
	
	newToken, createTokenErr := createToken(paramUser)
	if createTokenErr != nil {
		log.Println(" createTokenErr  ", createTokenErr)
		return c.Redirect(http.StatusTemporaryRedirect, "/login")		
	}

	storedUser["accesstoken"] = newToken.AccessToken
	storedUser["refreshtoken"] = newToken.RefreshToken
	store.Set(paramUser, storedUser)
	store.Save()
	// return c.Render(http.StatusBadRequest, "/setting/connections/CloudConnection", map[string]interface{}{
	// return c.Redirect(http.StatusTemporaryRedirect, "/setting/connections/CloudConnection")

	loginInfo := model.LoginInfo{
		Username: storedUser["username"],
		AccessToken: storedUser["accessToken"],
	}
	log.Println(" loginInfo  ", loginInfo)
	// c.Response().Header().Set("logininfo", {"admin"})
  	// c.Response().WriteHeader(http.StatusOK)
	return c.Redirect(http.StatusMovedPermanently, "../setting/connections/CloudConnection")
	// return echotemplate.Render(c, http.StatusOK, 
	// 	"setting/connections/CloudConnection", 
	// 	map[string]interface{}{
	// 		"LoginInfo": loginInfo,
	// })
	// return echotemplate.Render(c, http.StatusOK, 
	// 	"setting/connections/CloudConnection", nil,
	// )
	// return echotemplate.Render(c, http.StatusOK, 
	// 	"Map.html", 
	// 	map[string]interface{}{
	// 			"LoginInfo": loginInfo,
	// 		},
	// )
	// return echotemplate.Render(c, http.StatusOK, 
	// 	"setting/connections/CloudConnection.html", 
	// 	map[string]interface{}{
	// 			"LoginInfo": loginInfo,
	// 		},
	// )
	// return c.Render(http.StatusOK, "/setting/connections/CloudConnection.html", loginInfo)

	// 	storedUser["defaultnamespage"] = nsInfo.ID		
		
	// 	// 저장 성공하면 namespace 목록 조회
	// 	nsList2, nsErr2 := service.GetNameSpaceList()
	// 	if nsErr2 != nil {
	// 		log.Println(" nsErr2  ", nsErr2)		
	// 		return c.Redirect(http.StatusTemporaryRedirect, "/setting/connections/CloudConnection")
	// 	}
	// 	log.Println("nsList2  ", nsList2)
	// 	nsList = nsList2
	// }
	// log.Println("nsList  ", nsList)
	// store.Set("namespaceList", nsList)// 이게 유효한가?? 쓸모없을 듯
	// store.Save()

	// // mcis가 있으면 dashboard로

	// // mcis가 없으면 mcis 등록화면으로
	
	// // return c.Render(http.StatusBadRequest, "/setting/connections/CloudConnection", map[string]interface{}{
	// return c.Redirect(http.StatusTemporaryRedirect, "/setting/connections/CloudConnection")
}


func createToken(username string) (*TokenDetails, error) {

	// var err error
  
	// atClaims := jwt.MapClaims{}
  	// atClaims["authorized"] = true  
	// atClaims["username"] = username  
	// atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()  
	// at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)  
	// token, err := at.SignedString([]byte(os.Getenv("LoginAccessSecret")))
  
	// if err != nil {
	//    return "", err  
	// }  
	// return token, nil  


	// 액세스 토큰(access token)이 만료된 경우 리프레시 토큰(refresh token)을 사용하여 
	// 새 액세스 토큰을 생성하여 액세스 토큰이 만료가 되더라도 사용자가 다시 로그인을 하지 않게
	// refresh token은 30분에 token만료.
	// action이 있을 때마다 at, rt(refresh token)은 갱신
	// 5분 넘어 action이 발생했을 때(at가 expired) rt가 유효하면 로그인 된 것으로 
	// 30분동안 action이 없으면 refresh token이 expire되므로 이후에는 로그인 필요
	// 페이지 호출할 때마다 유효성 검증 후 expired 시간 재할당.
	var err error
	
	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 5).Unix()  
	td.AccessUuid = uuid.NewV4().String()  
	// td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()  
	td.RtExpires = time.Now().Add(time.Minute * 30).Unix()  
	td.RefreshUuid = uuid.NewV4().String()
	
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	//atClaims["user_id"] = userid
	atClaims["username"] = username
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
  
	td.AccessToken, err = at.SignedString([]byte( os.Getenv("LoginAccessSecret") ))
  
	if err != nil {  
		log.Println("create accessToken  ", err)
		return nil, err
  	}
  
	//Creating Refresh Token
 	rtClaims := jwt.MapClaims{}
 	rtClaims["refresh_uuid"] = td.RefreshUuid
 	// rtClaims["user_id"] = userid
	rtClaims["username"] = username
 	rtClaims["exp"] = td.RtExpires
 	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
  
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("LoginRefreshSecret")))
  	
	if err != nil {
		log.Println("create RefreshToken  ", err)
		return nil, err
 	}
  
	return td, nil
  }

// Login 없이 접근가능 
// Login이 필요없는 화면에서 호출하는게 의미 있나? 없이 써도 되는 듯.
func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

// Token이 있어야 접근가능
// login 이 필요한 page에서 호출하여 값이 true일 때만 접근가능
func restricted(c echo.Context) error {
	user := c.Get("UserName").(*jwt.Token)
	// user := c.Get("email").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["username"].(string)
	return c.String(http.StatusOK, "Welcome "+username+"!")
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

	reqInfo := new(model.ReqInfo)

	getUser := strings.TrimSpace(reqInfo.UserName)

	store.Set(getUser, nil)
	store.Save()
	log.Println(" auth expired ")

	return c.Render(http.StatusOK, "login.html", nil)

}
