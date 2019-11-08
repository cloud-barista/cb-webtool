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

	e.GET("/", controller.IndexController)
	e.GET("/dashboard", controller.DashBoard)

	//login 관련
	e.GET("/login", controller.LoginForm)
	e.POST("/login/proc", controller.LoginController)
	e.POST("/regUser", controller.RegUserConrtoller)

	// MCIS
	e.GET("/MCIS/register", controller.McisRegForm)
	e.GET("/MCIS/list", controller.McisListForm)
	// MCIS지울것
	//예가 리스트 전부
	e.GET("/ns/:nsid/mcis", func(c echo.Context) error {
		res := map[string]interface{}{
			"mcis": []map[string]string{
				{
					"id":     "7e3130a0-a811-47b8-a82c-b155267edef5",
					"name":   "mcis-1-t001",
					"vm_num": "3",
					"status": "launching",
				},
				{
					"id":     "423123123-a811-47b8-a82c-b155267edef5",
					"name":   "mcis-2-t002",
					"vm_num": "4",
					"status": "launching",
				},
				{
					"id":     "087070987-a811-47b8-a82c-b155267edef5",
					"name":   "mcis-3-t003",
					"vm_num": "2",
					"status": "launching",
				},
			},
		}
		return c.JSON(http.StatusOK, res)
	})
	e.GET("/ns/:nsid/mcis/:mcis_id", func(c echo.Context) error {
		res := map[string]interface{}{
			"id":     "7e3130a0-a811-47b8-a82c-b155267edef5",
			"name":   "mcis-2-t003",
			"vm_num": "3",
			"status": "launching",
			"vm": []map[string]string{
				{
					"id":                "04b9a6f1-c210-4941-bae1-545fb76fbb63",
					"csp_vm_id":         "azureshson0",
					"name":              "azure-t09",
					"status":            "Running",
					"public_ip":         "52.231.161.89",
					"private_ip":        "192.168.0.2",
					"domain_name":       "Not assigned yet",
					"config_name":       "aws-connection-config-01",
					"spec_id":           "17c12631-d29c-46c9-8390-322ad065cc39",
					"image_id":          "UUID-for-aws-ubuntu-image",
					"vnet_id":           "17c12631-d29c-46c9-8390-322ad065cc39",
					"vnic_id":           "17c12631-d29c-46c9-8390-322ad065cc39",
					"security_group_id": "17c12631-d29c-46c9-8390-322ad065cc39",
					"ssh_key_id":        "17c12631-d29c-46c9-8390-322ad065cc39",
					"description":       "description",
				},
				{
					"id":                "66074602-bc67-4604-b736-fc75205afeb3",
					"csp_vm_id":         "etri-shson0",
					"name":              "gcp-vmt05",
					"status":            "Running",
					"public_ip":         "34.97.218.87",
					"private_ip":        "192.168.0.6",
					"domain_name":       "Not assigned yet",
					"config_name":       "aws-connection-config-01",
					"spec_id":           "17c12631-d29c-46c9-8390-322ad065cc39",
					"image_id":          "UUID-for-aws-ubuntu-image",
					"vnet_id":           "17c12631-d29c-46c9-8390-322ad065cc39",
					"vnic_id":           "17c12631-d29c-46c9-8390-322ad065cc39",
					"security_group_id": "17c12631-d29c-46c9-8390-322ad065cc39",
					"ssh_key_id":        "17c12631-d29c-46c9-8390-322ad065cc39",
					"description":       "description",
				},
				{
					"id":                "7a76be36-f8aa-4ac9-a715-e6201c73365a",
					"csp_vm_id":         "i-08b5318cb5c61fa9c",
					"name":              "aws-vmtest06",
					"status":            "Running",
					"public_ip":         "52.78.122.12",
					"private_ip":        "192.168.0.9",
					"domain_name":       "Not assigned yet",
					"config_name":       "aws-connection-config-01",
					"spec_id":           "17c12631-d29c-46c9-8390-322ad065cc39",
					"image_id":          "UUID-for-aws-ubuntu-image",
					"vnet_id":           "17c12631-d29c-46c9-8390-322ad065cc39",
					"vnic_id":           "17c12631-d29c-46c9-8390-322ad065cc39",
					"security_group_id": "17c12631-d29c-46c9-8390-322ad065cc39",
					"ssh_key_id":        "17c12631-d29c-46c9-8390-322ad065cc39",
					"description":       "description",
				},
			},
			"description": "Test description",
		}
		return c.JSON(http.StatusOK, res)
	})

	// Namespace 관련 rest server
	// 나중에 전부 지울것
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

	// 웹툴에서 처리할 Connection
	e.GET("/Connection/list", controller.ConnectionListForm)
	e.GET("/Connection/reg", controller.ConnectionRegForm)
	e.POST("/Connection/reg/proc", controller.NsRegController)

	// 웹툴에서 처리할 Region
	e.GET("/Region/list", controller.RegionListForm)
	e.GET("/Region/reg", controller.RegionRegForm)
	e.POST("/Region/reg/proc", controller.NsRegController)

	// 웹툴에서 처리할 Credential
	e.GET("/Credential/list", controller.CredertialListForm)
	e.GET("/Credential/reg", controller.CredertialRegForm)

	// 웹툴에서 처리할 Driver
	e.GET("/Driver/list", controller.DriverListForm)
	e.GET("/Driver/reg", controller.DriverRegForm)

	e.Logger.Fatal(e.Start(":1234"))

}

type myStruct struct {
	Name   string
	Age    int
	Height int
}
