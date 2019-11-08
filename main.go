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
	e.GET("/MCIS/reg", controller.McisRegForm)
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

	e.GET("/ns/:nsid/resources/sshKey", func(c echo.Context) error {
		res := map[string]interface{}{
			"sshKey": []map[string]interface{}{
				{
					"id":             "1ac8c088-69cb-4c3b-b3ad-1c7e79eb5889",
					"connectionName": "aws-config01",
					"cspSshKeyName":  "shson-ssh-test1",
					"fingerprint":    "d3:0a:9b:50:1f:c7:bb:b1:fb:4c:dd:bd:3c:9c:ae:ba:e8:e8:ab:8c",
					"username":       "",
					"publicKey":      "",
					"privateKey":     "-----BEGIN RSA PRIVATE KEY-----\nMIIEpAIBAAKCAQEAyArKyFa/3QgOeHkB27oU/HbMV9CarMIvG5+H/ljn1OYyKeMFy8AS7UcLFEIG\n7lSkKrf5G8tWB5mkb2YvUemS+ltyvkE1njQFFYuwH0bwimY1FKu9YAoN/vefnbI81KaFBOLlt0VE\na12DH1cpaWObFCJDc+3e2q9XA28Dy4Z8M7T6llPEqUoOjGYRPkYhkVIa30gjQpWNqea2kXmj2Sdt\n8+dp224LW4a2vnWeQ4mKPnkK4KknJJ49XB+ZQluVqGvgvpglGb5SyAlAUSLk7zAG0WBGQUHb0zy8\n8fRa/H1FaYJO7sO5q9KwSAabOlebwdhXeeVX0ferNZ/U0s7fzy0DXQIDAQABAoIBAE6y6DpO3qaf\nf8nnmVKPb6gvOI8ns2JZ9tyAM6ld4r8AXSXbebFB/HA67eHnZACpbficbjWAFnVg/a8R3XX1VWyH\nQ6oTz3tQ7dsfHIgBiap5MVLTiepZqk3vr20D7Sif5l8YwNUMPDGdFXPj/5fFpwIRxMW8BFu2dZ4V\nM8IDQ6O3Uev8kMwx9nAqyhmzfp/ql0w9N9R5Cu9TMH36CTb0A1ww+QcTb1tUvxvKrNq7p0hg6r1d\n4jiuJfHJiovIEAIat4N7rzIdQoRMx8yI75t2/K0jqNXFYfLBiO2smP6CHFXItjdwdbKTCv7VtlKP\nMaW2qur4oJhUgv8GIR+roBy6WyUCgYEA9PmD3gvrUcn03F4Pciif/YFpaOX14wPPcg2YDth8qrtJ\nE8XvyRrdkq2pfspRyl+3Jx2DquqlJkALrKdLBk36v4Vd9CmEO5UXT4tWcMjo+lIIQfp99L2Ez9+Y\nCWJpCr+O5J7s79/l5SPHLEr0rF5uv+pZ41V4GDM0HqJll3O5Bu8CgYEA0QuV26AsJIbGxa+ONM+t\nbRQctEnPOPpJq08GFEhpckSPT70j4WrJ3NHMwwzT1/ASVV4K/SEd81YHWc1KIWrgjpwPEAuoAjGN\nVCi21jl0WoIuIFyB9PbYJhykUUDYiZSusk5DdcbfqWTg8ntmgev5HqYO8IAP0eJ3yK3OsJ5PenMC\ngYEA6NKTX1+YoLz+OMo0h9zQYZCy6/1SehVO/SiqUcGyillBFMfUIx+jYhomstf6cAoT+dr1HmWv\n2/CWp9q/VRibrZZFOx6SDEagRvs4hiyMMAvyyTIWr5nHNgFdb93V019HoUTiDwCOb/5W92Otsnx1\naXSDRaofX3CaolrZjt1vBoUCgYEAilaqY19KFxoB7MzSxOwyjp7iqAS4V7J6kh2HnmCVN4Nbe59l\nYUV0NOe6I9IXVy2OVGQZzY3e7iueTbVnO1opJPbtmOa91kXIi0suQ/Jdp4/CyrOtZNj+DaqiqwrI\nwbNdMK5OQmDLnqQdoRo8qfnpMHkgJdP5pCHEt08eGw+I9TUCgYAa+bKuhmN19ul9WB3+CFhzOWLh\niJt7U0icOMCBfiyyrL39rm/ycKPxgxc6vNHaNbVPbOjaC1+lUWGcDwGCB8sArdlhpn3BCYC3untm\nsTArFmBRJgPgQIHNki7rkXSW0fiouvj3GEy0rkamuXJLcAQ6k4P5pWVSdWKtRxub/8UMXA==\n-----END RSA PRIVATE KEY-----",
					"description":    "shson test description.",
					"keyValueList": []map[string]interface{}{
						{
							"Key":   "KeyMaterial",
							"Value": "-----BEGIN RSA PRIVATE KEY-----\nMIIEpAIBAAKCAQEAyArKyFa/3QgOeHkB27oU/HbMV9CarMIvG5+H/ljn1OYyKeMFy8AS7UcLFEIG\n7lSkKrf5G8tWB5mkb2YvUemS+ltyvkE1njQFFYuwH0bwimY1FKu9YAoN/vefnbI81KaFBOLlt0VE\na12DH1cpaWObFCJDc+3e2q9XA28Dy4Z8M7T6llPEqUoOjGYRPkYhkVIa30gjQpWNqea2kXmj2Sdt\n8+dp224LW4a2vnWeQ4mKPnkK4KknJJ49XB+ZQluVqGvgvpglGb5SyAlAUSLk7zAG0WBGQUHb0zy8\n8fRa/H1FaYJO7sO5q9KwSAabOlebwdhXeeVX0ferNZ/U0s7fzy0DXQIDAQABAoIBAE6y6DpO3qaf\nf8nnmVKPb6gvOI8ns2JZ9tyAM6ld4r8AXSXbebFB/HA67eHnZACpbficbjWAFnVg/a8R3XX1VWyH\nQ6oTz3tQ7dsfHIgBiap5MVLTiepZqk3vr20D7Sif5l8YwNUMPDGdFXPj/5fFpwIRxMW8BFu2dZ4V\nM8IDQ6O3Uev8kMwx9nAqyhmzfp/ql0w9N9R5Cu9TMH36CTb0A1ww+QcTb1tUvxvKrNq7p0hg6r1d\n4jiuJfHJiovIEAIat4N7rzIdQoRMx8yI75t2/K0jqNXFYfLBiO2smP6CHFXItjdwdbKTCv7VtlKP\nMaW2qur4oJhUgv8GIR+roBy6WyUCgYEA9PmD3gvrUcn03F4Pciif/YFpaOX14wPPcg2YDth8qrtJ\nE8XvyRrdkq2pfspRyl+3Jx2DquqlJkALrKdLBk36v4Vd9CmEO5UXT4tWcMjo+lIIQfp99L2Ez9+Y\nCWJpCr+O5J7s79/l5SPHLEr0rF5uv+pZ41V4GDM0HqJll3O5Bu8CgYEA0QuV26AsJIbGxa+ONM+t\nbRQctEnPOPpJq08GFEhpckSPT70j4WrJ3NHMwwzT1/ASVV4K/SEd81YHWc1KIWrgjpwPEAuoAjGN\nVCi21jl0WoIuIFyB9PbYJhykUUDYiZSusk5DdcbfqWTg8ntmgev5HqYO8IAP0eJ3yK3OsJ5PenMC\ngYEA6NKTX1+YoLz+OMo0h9zQYZCy6/1SehVO/SiqUcGyillBFMfUIx+jYhomstf6cAoT+dr1HmWv\n2/CWp9q/VRibrZZFOx6SDEagRvs4hiyMMAvyyTIWr5nHNgFdb93V019HoUTiDwCOb/5W92Otsnx1\naXSDRaofX3CaolrZjt1vBoUCgYEAilaqY19KFxoB7MzSxOwyjp7iqAS4V7J6kh2HnmCVN4Nbe59l\nYUV0NOe6I9IXVy2OVGQZzY3e7iueTbVnO1opJPbtmOa91kXIi0suQ/Jdp4/CyrOtZNj+DaqiqwrI\nwbNdMK5OQmDLnqQdoRo8qfnpMHkgJdP5pCHEt08eGw+I9TUCgYAa+bKuhmN19ul9WB3+CFhzOWLh\niJt7U0icOMCBfiyyrL39rm/ycKPxgxc6vNHaNbVPbOjaC1+lUWGcDwGCB8sArdlhpn3BCYC3untm\nsTArFmBRJgPgQIHNki7rkXSW0fiouvj3GEy0rkamuXJLcAQ6k4P5pWVSdWKtRxub/8UMXA==\n-----END RSA PRIVATE KEY-----",
						},
					},
				},
			},
		}
		return c.JSON(http.StatusOK, res)
	})
	// connection config
	e.GET("/connectionconfig", func(c echo.Context) error {

		res := []map[string]string{
			{
				"ConfigName":     "aws-config01",
				"ProviderName":   "AWS",
				"DriverName":     "aws-driver01",
				"CredentialName": "aws-credential01",
				"RegionName":     "aws-region01",
			},
		}

		return c.JSON(http.StatusOK, res)
	})

	e.GET("/ns/:nsid/resources/image", func(c echo.Context) error {

		res := map[string]interface{}{
			"image": []map[string]interface{}{

				{"id": "bb3cd5b6-c1ee-471e-810e-6d20eea072da",
					"connectionName": "aws-config01",
					"cspImageId":     "ami-00a54827eb7ffcd3c",
					"cspImageName":   "ubuntu/images/hvm-ssd/ubuntu-bionic-18.04-amd64-server-20190814",
					"creationDate":   "2019-08-19T18:11:28.000Z",
					"description":    "Canonical, Ubuntu, 18.04 LTS, amd64 bionic image build on 2019-08-14",
					"guestOS":        "ubuntu",
					"status":         "",
					"keyValueList": []map[string]string{
						{
							"Key":   "key1, written in Postman req",
							"Value": "value1, written in Postman req",
						},
						{
							"Key":   "key2, written in Postman req",
							"Value": "value2, written in Postman req",
						},
					},
				}},
		}

		return c.JSON(http.StatusOK, res)
	})

	e.GET("/ns/:nsid/resources/securityGroup", func(c echo.Context) error {

		res := map[string]interface{}{
			"securityGroup": []map[string]interface{}{
				{
					"id":                   "149d0be5-767d-4bbc-a943-3e5c6b824d71",
					"connectionName":       "aws-config01",
					"cspSecurityGroupId":   "sg-032cdc7495f2dd3e0",
					"cspSecurityGroupName": "shson-test2",
					"description":          "",
					"firewallRules": []map[string]interface{}{
						{
							"fromPort":   "20",
							"toPort":     "22",
							"ipProtocol": "tcp",
							"direction":  "inbound",
						},
						{
							"fromPort":   "",
							"toPort":     "",
							"ipProtocol": "-1",
							"direction":  "outbound",
						},
					},
					"keyValueList": []map[string]interface{}{
						{
							"Key":   "GroupName",
							"Value": "shson-test2",
						},
						{
							"Key":   "VpcID",
							"Value": "vpc-0ccb5c735b1dcd646",
						},
						{
							"Key":   "OwnerID",
							"Value": "635484366616",
						},
						{
							"Key":   "Description",
							"Value": "shson-test2",
						},
					},
				},
			},
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
