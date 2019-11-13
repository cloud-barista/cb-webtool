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

	// // MCIS LifeCycle
	// e.GET("/ns/:nsid/mcis/:mcis_id", func(c echo.Context) error {
	// 	action := c.QueryParam("action")
	// 	return c.JSON(http.StatusOK, map[string]interface{}{
	// 		"message": "MCIS : " + action,
	// 	})
	// })

	// // VM LifeCycle
	// e.GET("/ns/:nsid/mcis/:mcis_id/vm/:vm_id", func(c echo.Context) error {
	// 	action := c.QueryParam("action")
	// 	return c.JSON(http.StatusOK, map[string]interface{}{
	// 		"message": "VM : " + action,
	// 	})
	// })

	// MCIS
	e.GET("/MCIS/reg", controller.McisRegForm)
	e.POST("/MCIS/reg/proc", controller.McisRegController)
	e.GET("/MCIS/list", controller.McisListForm)
	// MCIS지울것
	//예가 리스트 전부
	e.GET("/ns/:nsid/mcis", func(c echo.Context) error {
		res := map[string]interface{}{

			"mcis": []map[string]interface{}{
				{
					"id":             "6c07184d-7307-468b-8a14-68ae69307968",
					"name":           "mcis-t01",
					"status":         "Running",
					"placement_algo": "",
					"description":    "Test description",
					"vm": []map[string]interface{}{
						{
							"id":          "0a21106f-e420-4152-b549-8eeb14e3a453",
							"name":        "aws-shson-vm02",
							"config_name": "aws-config01",
							"region": map[string]string{
								"Region": "ap-southeast",
								"Zone":   "ap-southeast-2a",
							},
							"publicIP":  "52.64.97.175",
							"publicDNS": "",
							"status":    "Running",
						},
						{
							"id":          "f276a1fe-ae5c-4dba-a88a-31747a19615a",
							"name":        "aws-shson-vm01",
							"config_name": "aws-config01",
							"region": map[string]string{
								"Region": "ap-southeast",
								"Zone":   "ap-southeast-2a",
							},
							"publicIP":  "",
							"publicDNS": "",
							"status":    "Running",
						},
						{
							"id":          "f276a1fe-ae5c-4dba-a88a-31747a19615a",
							"name":        "aws-shson-vm03",
							"config_name": "aws-config01",
							"region": map[string]string{
								"Region": "ap-southeast",
								"Zone":   "ap-southeast-2a",
							},
							"publicIP":  "",
							"publicDNS": "",
							"status":    "Running",
						},
						{
							"id":          "f276a1fe-ae5c-4dba-a88a-31747a19615a",
							"name":        "aws-shson-vm04",
							"config_name": "aws-config01",
							"region": map[string]string{
								"Region": "ap-southeast",
								"Zone":   "ap-southeast-2a",
							},
							"publicIP":  "",
							"publicDNS": "",
							"status":    "Running",
						},
						{
							"id":          "f276a1fe-ae5c-4dba-a88a-31747a19615a",
							"name":        "aws-shson-vm05",
							"config_name": "aws-config01",
							"region": map[string]string{
								"Region": "ap-southeast",
								"Zone":   "ap-southeast-2a",
							},
							"publicIP":  "",
							"publicDNS": "",
							"status":    "Running",
						},
					},
				},
				{
					"id":             "12c07184d-7307-468b-8a14-68ae69307968",
					"name":           "mcis-t02",
					"status":         "Running",
					"placement_algo": "",
					"description":    "Test description",
					"vm": []map[string]interface{}{
						{
							"id":          "0a21106f-e420-4152-b549-8eeb14e3a453",
							"name":        "aws-shson-vm02",
							"config_name": "aws-config01",
							"region": map[string]string{
								"Region": "ap-southeast",
								"Zone":   "ap-southeast-2a",
							},
							"publicIP":  "52.64.97.175",
							"publicDNS": "",
							"status":    "Running",
						},
						{
							"id":          "f276a1fe-ae5c-4dba-a88a-31747a19615a",
							"name":        "aws-shson-vm01",
							"config_name": "aws-config01",
							"region": map[string]string{
								"Region": "ap-southeast",
								"Zone":   "ap-southeast-2a",
							},
							"publicIP":  "",
							"publicDNS": "",
							"status":    "Running",
						},
						{
							"id":          "f276a1fe-ae5c-4dba-a88a-31747a19615a",
							"name":        "aws-shson-vm03",
							"config_name": "aws-config01",
							"region": map[string]string{
								"Region": "ap-southeast",
								"Zone":   "ap-southeast-2a",
							},
							"publicIP":  "",
							"publicDNS": "",
							"status":    "Running",
						},
						{
							"id":          "f276a1fe-ae5c-4dba-a88a-31747a19615a",
							"name":        "aws-shson-vm04",
							"config_name": "aws-config01",
							"region": map[string]string{
								"Region": "ap-southeast",
								"Zone":   "ap-southeast-2a",
							},
							"publicIP":  "",
							"publicDNS": "",
							"status":    "Running",
						},
						{
							"id":          "f276a1fe-ae5c-4dba-a88a-31747a19615a",
							"name":        "aws-shson-vm05",
							"config_name": "aws-config01",
							"region": map[string]string{
								"Region": "ap-southeast",
								"Zone":   "ap-southeast-2a",
							},
							"publicIP":  "",
							"publicDNS": "",
							"status":    "Running",
						},
					},
				},
				{
					"id":             "42357184d-7307-468b-8a14-68ae69307968",
					"name":           "mcis-t03",
					"status":         "Running",
					"placement_algo": "ap-southeast",
					"description":    "Test description",
					"vm": []map[string]interface{}{
						{
							"id":          "0a21106f-e420-4152-b549-8eeb14e3a453",
							"name":        "aws-shson-vm02",
							"config_name": "aws-config01",
							"region": map[string]string{
								"Region": "ap-southeast",
								"Zone":   "ap-southeast-2a",
							},
							"publicIP":  "52.64.97.175",
							"publicDNS": "",
							"status":    "Running",
						},
						{
							"id":          "f276a1fe-ae5c-4dba-a88a-31747a19615a",
							"name":        "aws-shson-vm01",
							"config_name": "aws-config01",
							"region": map[string]string{
								"Region": "ap-southeast",
								"Zone":   "ap-southeast-2a",
							},
							"publicIP":  "",
							"publicDNS": "",
							"status":    "Running",
						},
						{
							"id":          "f276a1fe-ae5c-4dba-a88a-31747a19615a",
							"name":        "aws-shson-vm03",
							"config_name": "aws-config01",
							"region": map[string]string{
								"Region": "ap-southeast",
								"Zone":   "ap-southeast-2a",
							},
							"publicIP":  "",
							"publicDNS": "",
							"status":    "Running",
						},
						{
							"id":          "f276a1fe-ae5c-4dba-a88a-31747a19615a",
							"name":        "aws-shson-vm04",
							"config_name": "aws-config01",
							"region": map[string]string{
								"Region": "ap-southeast",
								"Zone":   "ap-southeast-2a",
							},
							"publicIP":  "",
							"publicDNS": "",
							"status":    "Running",
						},
						{
							"id":          "f276a1fe-ae5c-4dba-a88a-31747a19615a",
							"name":        "aws-shson-vm05",
							"config_name": "aws-config01",
							"region": map[string]string{
								"Region": "ap-southeast",
								"Zone":   "ap-southeast-2a",
							},
							"publicIP":  "",
							"publicDNS": "",
							"status":    "Running",
						},
					},
				},
			},
		}
		return c.JSON(http.StatusOK, res)
	})
	e.GET("/ns/:nsid/mcis/:mcis_id", func(c echo.Context) error {
		action := c.QueryParam("action")
		if action != "" {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"message": "MCIS : " + action,
			})
		}
		res := map[string]interface{}{
			"id":     "5d910c76-364f-484e-a2b2-90ea4feabe3a",
			"name":   "mcis-t01",
			"status": "Running",
			"vm": []map[string]interface{}{
				{
					"id":           "09177a33-63d7-477c-a81f-91a258255450",
					"name":         "aws-shson-vm02",
					"config_name":  "aws-config01",
					"spec_id":      "d3959c21-af25-46b0-9316-ab7f08934371",
					"image_id":     "bc352bf1-93d6-47f1-a558-485f1dff695b",
					"vnet_id":      "08b5de73-fcd4-4fd9-a074-7071796aec03",
					"vnic_id":      "",
					"public_ip_id": "af4633ac-0beb-4f9f-a40e-f5d33ba3b6c2",
					"security_group_ids": []string{
						"149d0be5-767d-4bbc-a943-3e5c6b824d71",
					},
					"ssh_key_id":       "1ac8c088-69cb-4c3b-b3ad-1c7e79eb5889",
					"description":      "description",
					"vm_access_id":     "",
					"vm_access_passwd": "",
					"vmUserId":         "",
					"vmUserPasswd":     "",
					"region": map[string]string{
						"Region": "ap-southeast",
						"Zone":   "ap-southeast-2a",
					},
					"publicIP":    "52.64.97.175",
					"publicDNS":   "",
					"privateIP":   "192.168.135.182",
					"privateDNS":  "ip-192-168-135-182.ap-southeast-2.compute.internal",
					"vmBootDisk":  "/dev/sda1",
					"vmBlockDisk": "/dev/sda1",
					"status":      "Running",
					"cspViewVmDetail": map[string]interface{}{
						"Name":      "aws-shson-vm02",
						"Id":        "i-0249226ec5e613be5",
						"StartTime": "0001-01-01T00:00:00Z",
						"Region": map[string]string{
							"Region": "ap-southeast",
							"Zone":   "ap-southeast-2a",
						},
						"ImageId":          "ami-00a54827eb7ffcd3c",
						"VMSpecId":         "t2.micro",
						"VirtualNetworkId": "vpc-0ccb5c735b1dcd646",
						"SecurityGroupIds": []string{
							"sg-032cdc7495f2dd3e0",
						},
						"NetworkInterfaceId": "eni-attach-049edd0dd3219b6b8",
						"PublicIP":           "52.64.97.175",
						"PublicDNS":          "",
						"PrivateIP":          "192.168.135.182",
						"PrivateDNS":         "ip-192-168-135-182.ap-southeast-2.compute.internal",
						"KeyPairName":        "shson-ssh-test1",
						"VMUserId":           "",
						"VMUserPasswd":       "",
						"VMBootDisk":         "/dev/sda1",
						"VMBlockDisk":        "/dev/sda1",
						"KeyValueList": []map[string]string{
							{
								"Key":   "State",
								"Value": "running",
							},
							{
								"Key":   "Architecture",
								"Value": "x86_64",
							},
							{
								"Key":   "VpcId",
								"Value": "vpc-0ccb5c735b1dcd646",
							},
							{
								"Key":   "SubnetId",
								"Value": "subnet-0a6d2e9a1c2052703",
							},
							{
								"Key":   "KeyName",
								"Value": "shson-ssh-test1",
							},
						},
					},
				},
				{
					"id":           "cf1f5704-7612-41c3-9235-14c519a8c0a5",
					"name":         "gcp-shson-vm01",
					"config_name":  "gcp-config01",
					"spec_id":      "d3959c21-af25-46b0-9316-ab7f08934371",
					"image_id":     "bc352bf1-93d6-47f1-a558-485f1dff695b",
					"vnet_id":      "08b5de73-fcd4-4fd9-a074-7071796aec03",
					"vnic_id":      "",
					"public_ip_id": "af4633ac-0beb-4f9f-a40e-f5d33ba3b6c2",
					"security_group_ids": []string{
						"149d0be5-767d-4bbc-a943-3e5c6b824d71",
					},
					"ssh_key_id":       "1ac8c088-69cb-4c3b-b3ad-1c7e79eb5889",
					"description":      "description",
					"vm_access_id":     "",
					"vm_access_passwd": "",
					"vmUserId":         "",
					"vmUserPasswd":     "",
					"region": map[string]string{
						"Region": "ap-southeast",
						"Zone":   "ap-southeast-2a",
					},
					"publicIP":    "52.64.97.175",
					"publicDNS":   "",
					"privateIP":   "192.168.25.205",
					"privateDNS":  "ip-192-168-25-205.ap-southeast-2.compute.internal",
					"vmBootDisk":  "/dev/sda1",
					"vmBlockDisk": "/dev/sda1",
					"status":      "Running",
					"cspViewVmDetail": map[string]interface{}{
						"Name":      "gcp-shson-vm01",
						"Id":        "i-06af16714219adbb3",
						"StartTime": "0001-01-01T00:00:00Z",
						"Region": map[string]string{
							"Region": "ap-southeast",
							"Zone":   "ap-southeast-2a",
						},
						"ImageId":          "ami-00a54827eb7ffcd3c",
						"VMSpecId":         "t2.micro",
						"VirtualNetworkId": "vpc-0ccb5c735b1dcd646",
						"SecurityGroupIds": []string{
							"sg-032cdc7495f2dd3e0",
						},
						"NetworkInterfaceId": "eni-attach-0c539abe156868d3e",
						"PublicIP":           "52.64.97.175",
						"PublicDNS":          "",
						"PrivateIP":          "192.168.25.205",
						"PrivateDNS":         "ip-192-168-25-205.ap-southeast-2.compute.internal",
						"KeyPairName":        "shson-ssh-test1",
						"VMUserId":           "",
						"VMUserPasswd":       "",
						"VMBootDisk":         "/dev/sda1",
						"VMBlockDisk":        "/dev/sda1",
						"KeyValueList": []map[string]string{
							{
								"Key":   "State",
								"Value": "running",
							},
							{
								"Key":   "Architecture",
								"Value": "x86_64",
							},
							{
								"Key":   "VpcId",
								"Value": "vpc-0ccb5c735b1dcd646",
							},
							{
								"Key":   "SubnetId",
								"Value": "subnet-0a6d2e9a1c2052703",
							},
							{
								"Key":   "KeyName",
								"Value": "shson-ssh-test1",
							},
						},
					},
				},
			},
			"placement_algo": "",
			"description":    "Test description",
		}

		return c.JSON(http.StatusOK, res)
	})
	e.GET("/ns/:nsid/mcis/:mcis_id/vm/:vm_id", func(c echo.Context) error {
		action := c.QueryParam("action")
		if action != "" {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"message": "VM : " + action,
			})
		}
		res := map[string]interface{}{

			"id":           "09177a33-63d7-477c-a81f-91a258255450",
			"name":         "aws-shson-vm02",
			"config_name":  "aws-config01",
			"spec_id":      "d3959c21-af25-46b0-9316-ab7f08934371",
			"image_id":     "bc352bf1-93d6-47f1-a558-485f1dff695b",
			"vnet_id":      "08b5de73-fcd4-4fd9-a074-7071796aec03",
			"vnic_id":      "",
			"public_ip_id": "af4633ac-0beb-4f9f-a40e-f5d33ba3b6c2",
			"security_group_ids": []string{
				"149d0be5-767d-4bbc-a943-3e5c6b824d71",
			},
			"ssh_key_id":       "1ac8c088-69cb-4c3b-b3ad-1c7e79eb5889",
			"description":      "description",
			"vm_access_id":     "",
			"vm_access_passwd": "",
			"vmUserId":         "",
			"vmUserPasswd":     "",
			"region": map[string]string{
				"Region": "ap-southeast",
				"Zone":   "ap-southeast-2a",
			},
			"publicIP":    "52.64.97.175",
			"publicDNS":   "",
			"privateIP":   "192.168.135.182",
			"privateDNS":  "ip-192-168-135-182.ap-southeast-2.compute.internal",
			"vmBootDisk":  "/dev/sda1",
			"vmBlockDisk": "/dev/sda1",
			"status":      "Running",
			"cspViewVmDetail": map[string]interface{}{
				"Name":      "aws-shson-vm02",
				"Id":        "i-0249226ec5e613be5",
				"StartTime": "0001-01-01T00:00:00Z",
				"Region": map[string]string{
					"Region": "ap-southeast",
					"Zone":   "ap-southeast-2a",
				},
				"ImageId":          "ami-00a54827eb7ffcd3c",
				"VMSpecId":         "t2.micro",
				"VirtualNetworkId": "vpc-0ccb5c735b1dcd646",
				"SecurityGroupIds": []string{
					"sg-032cdc7495f2dd3e0",
				},
				"NetworkInterfaceId": "eni-attach-049edd0dd3219b6b8",
				"PublicIP":           "52.64.97.175",
				"PublicDNS":          "",
				"PrivateIP":          "192.168.135.182",
				"PrivateDNS":         "ip-192-168-135-182.ap-southeast-2.compute.internal",
				"KeyPairName":        "shson-ssh-test1",
				"VMUserId":           "",
				"VMUserPasswd":       "",
				"VMBootDisk":         "/dev/sda1",
				"VMBlockDisk":        "/dev/sda1",
				"KeyValueList": []map[string]string{
					{
						"Key":   "State",
						"Value": "running",
					},
					{
						"Key":   "Architecture",
						"Value": "x86_64",
					},
					{
						"Key":   "VpcId",
						"Value": "vpc-0ccb5c735b1dcd646",
					},
					{
						"Key":   "SubnetId",
						"Value": "subnet-0a6d2e9a1c2052703",
					},
					{
						"Key":   "KeyName",
						"Value": "shson-ssh-test1",
					},
				},
			},
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
			{
				"ConfigName":     "gcp-config01",
				"ProviderName":   "GCP",
				"DriverName":     "gcp-driver01",
				"CredentialName": "gcp-credential01",
				"RegionName":     "gcp-region01",
			},
		}

		return c.JSON(http.StatusOK, res)
	})

	e.GET("/ns/:nsid/resources/image", func(c echo.Context) error {

		res := map[string]interface{}{
			"image": []map[string]interface{}{
				{
					"id":             "bb3cd5b6-c1ee-471e-810e-6d20eea072da",
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
				},
				{
					"id":             "12345-c1ee-471e-810e-6d20eea072da",
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
				},
				{
					"id":             "543212-c1ee-471e-810e-6d20eea072da",
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
				},
			},
		}

		return c.JSON(http.StatusOK, res)
	})

	e.GET("/ns/:nsid/resources/network", func(c echo.Context) error {

		res := map[string]interface{}{
			"network": []map[string]interface{}{
				{
					"id":             "08b5de73-fcd4-4fd9-a074-7071796aec03",
					"connectionName": "aws-config01",
					"cspNetworkId":   "subnet-0a6d2e9a1c2052703",
					"cspNetworkName": "CB-VNet-Subnet",
					"cidrBlock":      "192.168.0.0/16",
					"description":    "",
					"status":         "pending",
					"keyValueList": []map[string]string{
						{
							"Key":   "VpcId",
							"Value": "vpc-0ccb5c735b1dcd646",
						},
						{
							"Key":   "MapPublicIpOnLaunch",
							"Value": "false",
						},
						{
							"Key":   "AvailableIpAddressCount",
							"Value": "65531",
						},
						{
							"Key":   "AvailabilityZone",
							"Value": "ap-southeast-2a",
						},
					},
				},
			},
		}

		return c.JSON(http.StatusOK, res)
	})

	e.GET("/ns/:nsid/resources/publicIp", func(c echo.Context) error {

		res := map[string]interface{}{
			"publicIp": []map[string]interface{}{
				{
					"id":              "af4633ac-0beb-4f9f-a40e-f5d33ba3b6c2",
					"connectionName":  "aws-config01",
					"cspPublicIpId":   "eipalloc-03ed95943b2adafab",
					"cspPublicIpName": "ShsonTest",
					"publicIp":        "52.64.97.175",
					"ownedVmId":       "",
					"description":     "Shson test description",
					"status":          "",
					"keyValueList": []map[string]string{
						{
							"Key":   "Domain",
							"Value": "vpc",
						},
						{
							"Key":   "PublicIpv4Pool",
							"Value": "amazon",
						},
						{
							"Key":   "AllocationId",
							"Value": "eipalloc-03ed95943b2adafab",
						},
						{
							"Key":   "Name",
							"Value": "ShsonTest",
						},
					},
				},
			},
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

	// Select Pop
	e.GET("/Pop/spec", controller.PopSpec)
	e.GET("/ns/:nsid/resources/spec", func(c echo.Context) error {

		res := map[string]interface{}{
			"spec": []map[string]interface{}{
				{
					"id":                    "041c71da-c024-4e30-9b6e-092bfcca6e25",
					"name":                  "t2.medium-02",
					"connectionName":        "aws-config01",
					"os_type":               "ubuntu",
					"num_vCPU":              "2",
					"num_core":              "",
					"mem_GiB":               "2",
					"storage_GiB":           "2",
					"description":           "",
					"cost_per_hour":         "6",
					"num_storage":           "",
					"max_num_storage":       "",
					"max_total_storage_TiB": "",
					"net_bw_Gbps":           "",
					"ebs_bw_Mbps":           "",
					"gpu_model":             "",
					"num_gpu":               "",
					"gpumem_GiB":            "",
					"gpu_p2p":               "",
				},
				{
					"id":                    "0cd007b3-d2c4-4913-a773-77bc43b94eaf",
					"name":                  "t2.micro-01",
					"connectionName":        "aws-config01",
					"os_type":               "ubuntu",
					"num_vCPU":              "1",
					"num_core":              "",
					"mem_GiB":               "1",
					"storage_GiB":           "1",
					"description":           "",
					"cost_per_hour":         "1",
					"num_storage":           "",
					"max_num_storage":       "",
					"max_total_storage_TiB": "",
					"net_bw_Gbps":           "",
					"ebs_bw_Mbps":           "",
					"gpu_model":             "",
					"num_gpu":               "",
					"gpumem_GiB":            "",
					"gpu_p2p":               "",
				},
				{
					"id":                    "0ee6c54f-43f5-479f-818d-eb18af42c02f",
					"name":                  "t2.micro-04",
					"connectionName":        "aws-config01",
					"os_type":               "ubuntu",
					"num_vCPU":              "1",
					"num_core":              "",
					"mem_GiB":               "1",
					"storage_GiB":           "1",
					"description":           "",
					"cost_per_hour":         "4",
					"num_storage":           "",
					"max_num_storage":       "",
					"max_total_storage_TiB": "",
					"net_bw_Gbps":           "",
					"ebs_bw_Mbps":           "",
					"gpu_model":             "",
					"num_gpu":               "",
					"gpumem_GiB":            "",
					"gpu_p2p":               "",
				},
				{
					"id":                    "58ade7fd-d108-44a3-99db-3a018c961e9a",
					"name":                  "t2.micro-03",
					"connectionName":        "aws-config01",
					"os_type":               "ubuntu",
					"num_vCPU":              "1",
					"num_core":              "",
					"mem_GiB":               "1",
					"storage_GiB":           "1",
					"description":           "",
					"cost_per_hour":         "3",
					"num_storage":           "",
					"max_num_storage":       "",
					"max_total_storage_TiB": "",
					"net_bw_Gbps":           "",
					"ebs_bw_Mbps":           "",
					"gpu_model":             "",
					"num_gpu":               "",
					"gpumem_GiB":            "",
					"gpu_p2p":               "",
				},
				{
					"id":                    "6111d39a-d676-4c87-b247-01dce72f3292",
					"name":                  "t2.2xlarge",
					"connectionName":        "aws-config01",
					"os_type":               "ubuntu",
					"num_vCPU":              "8",
					"num_core":              "",
					"mem_GiB":               "20",
					"storage_GiB":           "100",
					"description":           "",
					"cost_per_hour":         "29",
					"num_storage":           "",
					"max_num_storage":       "",
					"max_total_storage_TiB": "",
					"net_bw_Gbps":           "",
					"ebs_bw_Mbps":           "",
					"gpu_model":             "",
					"num_gpu":               "",
					"gpumem_GiB":            "",
					"gpu_p2p":               "",
				},
				{
					"id":                    "69c573cf-0341-43c1-80a3-426835684e42",
					"name":                  "t2.micro-02",
					"connectionName":        "aws-config01",
					"os_type":               "ubuntu",
					"num_vCPU":              "1",
					"num_core":              "",
					"mem_GiB":               "1",
					"storage_GiB":           "1",
					"description":           "",
					"cost_per_hour":         "2",
					"num_storage":           "",
					"max_num_storage":       "",
					"max_total_storage_TiB": "",
					"net_bw_Gbps":           "",
					"ebs_bw_Mbps":           "",
					"gpu_model":             "",
					"num_gpu":               "",
					"gpumem_GiB":            "",
					"gpu_p2p":               "",
				},
				{
					"id":                    "9c744220-b28a-4636-a6c1-078f05c38ec9",
					"name":                  "t2.medium-01",
					"connectionName":        "aws-config01",
					"os_type":               "ubuntu",
					"num_vCPU":              "2",
					"num_core":              "",
					"mem_GiB":               "2",
					"storage_GiB":           "2",
					"description":           "",
					"cost_per_hour":         "5",
					"num_storage":           "",
					"max_num_storage":       "",
					"max_total_storage_TiB": "",
					"net_bw_Gbps":           "",
					"ebs_bw_Mbps":           "",
					"gpu_model":             "",
					"num_gpu":               "",
					"gpumem_GiB":            "",
					"gpu_p2p":               "",
				},
				{
					"id":                    "c477f9f7-8e78-4678-8e42-b7f266658888",
					"name":                  "t2.large",
					"connectionName":        "aws-config01",
					"os_type":               "ubuntu",
					"num_vCPU":              "4",
					"num_core":              "",
					"mem_GiB":               "20",
					"storage_GiB":           "80",
					"description":           "",
					"cost_per_hour":         "29",
					"num_storage":           "",
					"max_num_storage":       "",
					"max_total_storage_TiB": "",
					"net_bw_Gbps":           "",
					"ebs_bw_Mbps":           "",
					"gpu_model":             "",
					"num_gpu":               "",
					"gpumem_GiB":            "",
					"gpu_p2p":               "",
				},
				{
					"id":                    "d3959c21-af25-46b0-9316-ab7f08934371",
					"name":                  "t2.micro",
					"connectionName":        "aws-config01",
					"os_type":               "ubuntu",
					"num_vCPU":              "1",
					"num_core":              "",
					"mem_GiB":               "1",
					"storage_GiB":           "1",
					"description":           "",
					"cost_per_hour":         "1",
					"num_storage":           "",
					"max_num_storage":       "",
					"max_total_storage_TiB": "",
					"net_bw_Gbps":           "",
					"ebs_bw_Mbps":           "",
					"gpu_model":             "",
					"num_gpu":               "",
					"gpumem_GiB":            "",
					"gpu_p2p":               "",
				},
				{
					"id":                    "ffb7657a-0f4a-40cd-b84e-8147744c001d",
					"name":                  "t2.medium-03",
					"connectionName":        "aws-config01",
					"os_type":               "ubuntu",
					"num_vCPU":              "2",
					"num_core":              "",
					"mem_GiB":               "2",
					"storage_GiB":           "2",
					"description":           "",
					"cost_per_hour":         "7",
					"num_storage":           "",
					"max_num_storage":       "",
					"max_total_storage_TiB": "",
					"net_bw_Gbps":           "",
					"ebs_bw_Mbps":           "",
					"gpu_model":             "",
					"num_gpu":               "",
					"gpumem_GiB":            "",
					"gpu_p2p":               "",
				},
			},
		}
		return c.JSON(http.StatusOK, res)
	})

	e.Logger.Fatal(e.Start(":1234"))

}

type myStruct struct {
	Name   string
	Age    int
	Height int
}
