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
	
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"

	//db "mzc/src/databases/store"
	"github.com/cloud-barista/cb-webtool/src/service"
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
	return echotemplate.Render(c, http.StatusOK, "Login", nil)
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

	getUser := strings.TrimSpace(reqInfo.UserName)
	getEmail := strings.TrimSpace(reqInfo.Email)
	getPass := strings.TrimSpace(reqInfo.Password)
	fmt.Println("getUser & getPass : ", getUser, getPass)

	

	// echoSession에서 가져오기
	get, ok := store.Get(getUser)
	fmt.Println("Stored USER:", get)


	// jwt token 사용
	if !ok {
		log.Println(" login proc err  ", ok)
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{ //401
			"message": " 정보가 없으니 다시 등록바랍니다.",
			"status":  "fail",
		})		
	}

	newToken, createTokenErr := createToken(getUser string)
	if createTokenErr != nil {
		log.Println(" login proc err  ", createTokenErr)
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{ //401
			"message": " 로그인 처리 요류",
			"status":  "fail",
		})		
	}

	get.accesstoken = newToken.AccessToken
	get.refreshtoken = newToken.RefreshToken
	store.Set(getUser, get)
	store.Save()
	
	//contentType := c.Request().Header.Get("Content-Type")
	//AccessToken := c.Request().Header.Get("AccessToken")
	//RefreshToken := c.Request().Header.Get("RefreshToken")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"status":  "200",
		"AccessToken": newToken.AccessToken,
		// "RefreshToken": newToken.RefreshToken,
	})
	//////// 현재구조에서는 nsList 부분을 포함해야 함. 
	//////// TODO : post로 넘기고 login success일 때 namespace 등 설정 popup이 아닌 page로 넘기는게 좋을 듯.
	///////         

	// echo store 사용
	// if !ok {
	// 	// //return c.JSON(http.StatusNotFound, map[string]interface{}{	//404
	// 	// //return c.JSON(http.StatusOK, map[string]interface{}{			..200
	// 	return c.JSON(http.StatusUnauthorized, map[string]interface{}{ //401
	// 		"message": " 정보가 없으니 다시 등록 해라",
	// 		"status":  "fail"	
	// 	})
	// }

	// return c.JSON(http.StatusOK, map[string]string{
	// 	return c.JSON(http.StatusOK, map[string]interface{}{
	// 		"message": "success",
	// 		"status":  "200",
	// 		"token": t,
	// 	})
		
	// })

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

	log.Println(" auth  ", nsList)


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

func createToken(username string) (string, error) {

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
	atClaims["user_id"] = userid
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
  
	td.AccessToken, err = at.SignedString([]byte( os.Getenv("LoginAccessSecret") ))
  
	if err != nil {  
	   return nil, err
  	}
  
	//Creating Refresh Token
 	rtClaims := jwt.MapClaims{}
 	rtClaims["refresh_uuid"] = td.RefreshUuid
 	rtClaims["user_id"] = userid
 	rtClaims["exp"] = td.RtExpires
 	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
  
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("LoginRefreshSecret")))
  	
	if err != nil {
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

	reqInfo := new(ReqInfo)

	getUser := strings.TrimSpace(reqInfo.UserName)

	store.Set(getUser, nil)
	store.Save()
	log.Println(" auth expired ")

	return c.Render(http.StatusOK, "login.html", nil)

}
