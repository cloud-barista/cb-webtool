package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"

	controller "github.com/cloud-barista/cb-webtool/src/controller"
	echosession "github.com/go-session/echo-session"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func init() {

}

type user struct {
	Name  string `json:"name" form:"name" query:"name"`
	Email string `json:"email" form:"email" query:"email"`
}

type connectionInfo struct {
	RegionName     string `json:"regionname"`
	ConfigName     string
	ProviderName   string `json:"ProviderName"`
	CredentialName string `json:"CredentialName"`
	DriverName     string `json:"DriverName"`
}

type TemplateRender struct {
	templates *template.Template
}

func (t *TemplateRender) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func requestApi(method string, restUrl string, body io.Reader) {

}

func main() {
	e := echo.New()
	e.Use(echosession.New())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/assets", "./src/static/assets")

	// paresGlob 를 사용하여 모든 경로에 있는 파일을 가져올 경우 사용하면 되겠다.
	// 사용한다음에 해당 파일을 불러오면 되네.
	// 서브디렉토리에 있는 걸 확인하기가 힘드네....
	renderer := &TemplateRender{
		templates: template.Must(template.ParseGlob(`./src/views/*.html`)),
	}

	e.Renderer = renderer

	e.GET("/", func(c echo.Context) error {
		fmt.Println("어디에 걸리나 보자")
		store := echosession.FromContext(c)
		getUser, ok := store.Get("username")

		if !ok {
			fmt.Println("nothing ")
			//return c.Render(http.StatusNotAcceptable, "login.html", nil)
			return c.Redirect(http.StatusPermanentRedirect, "/login")
		}
		result := map[string]string{}
		getObj, ok := store.Get(getUser.(string))

		if !ok {
			//return c.Render(http.StatusPermanentRedirect, "login.html", nil)
			return c.Redirect(http.StatusPermanentRedirect, "/login")
		}
		for k, v := range getObj.(map[string]string) {
			result[k] = v
		}

		defer func() {
			if e := recover(); e != nil {
				fmt.Printf("error : %s\r\n ", e)
			}
		}()

		return c.Render(http.StatusAccepted, "dashboard.html", result)

	})

	e.GET("/hello", func(c echo.Context) error {
		return c.Render(http.StatusOK, "hello.html", map[string]interface{}{
			"Name": myStruct{Name: "Dennis", Age: 36, Height: 170},
		})
	})

	e.GET("/dashboard", func(c echo.Context) error {
		fmt.Println("=========== DashBoard start ==============")
		if loginInfo := controller.CallLoginInfo(c); loginInfo.Username != "" {
			return c.Render(http.StatusOK, "dashboard.html", map[string]interface{}{
				"LoginInfo": loginInfo,
			})

		}

		return c.Redirect(http.StatusPermanentRedirect, "/login")

	})

	e.GET("/login", controller.LoginForm)
	e.POST("/login/proc", controller.LoginController)
	e.POST("/regUser", controller.RegUserConrtoller)

	e.GET("/MCIS/register", func(c echo.Context) error {
		if loginInfo := controller.CallLoginInfo(c); loginInfo.NameSpace != "" {
			return c.Render(http.StatusOK, "MCISRegister.html", map[string]interface{}{
				"LoginInfo": loginInfo,
			})

		}

		return c.Redirect(http.StatusPermanentRedirect, "/login")

	})

	e.GET("/MCIS/list", func(c echo.Context) error {
		if loginInfo := controller.CallLoginInfo(c); loginInfo.NameSpace != "" {
			return c.Render(http.StatusOK, "MCISlist.html", map[string]interface{}{
				"LoginInfo": loginInfo,
			})

		}
		return c.Redirect(http.StatusPermanentRedirect, "/login")
	})

	// Namespace 관련 rest server
	e.GET("/ns", func(c echo.Context) error {
		res := []map[string]string{
			{
				"id":          "879f1c57-857e-4430-b904-0cda2c16c580",
				"name":        "Seokho Son Name Space",
				"description": "description-2019-10-01",
			},
			{
				"id":          "bce5fb65-f617-4d45-97b7-f6c299559010",
				"name":        "my name spaced",
				"description": "description-2019-08-14",
			},
		}
		return c.JSON(http.StatusOK, map[string][]map[string]string{
			"ns": res,
		})
	})

	e.GET("/ns/:nsid", func(c echo.Context) error {
		nsID := c.Param("nsid")
		fmt.Println("nameSpaceID : ", nsID)
		res := map[string]string{
			"id":          "879f1c57-857e-4430-b904-0cda2c16c580",
			"name":        "Seokho Son Name Space",
			"description": "description-2019-10-01",
		}

		return c.JSON(http.StatusOK, res)
	})

	e.POST("/ns", func(c echo.Context) error {
		res := map[string]string{
			"message": "success",
		}
		return c.JSON(http.StatusOK, res)
	})
	e.DELETE("/ns/:nsid", func(c echo.Context) error {
		res := map[string]string{
			"message": "success",
		}
		return c.JSON(http.StatusOK, res)
	})

	e.GET("/SET/NS/:nsid", controller.SetNameSpace)

	// 웹툴에서 처리할 NameSpace
	e.GET("/NS/list", controller.NsListForm)
	e.GET("/NS/reg", controller.NsRegForm)
	e.POST("/NS/reg/proc", controller.NsRegController)
	e.GET("/GET/ns", controller.GetNameSpace)

	e.Logger.Fatal(e.Start(":1234"))

}

type myStruct struct {
	Name   string
	Age    int
	Height int
}
