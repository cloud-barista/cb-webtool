package main

import (
	"html/template"
	"io"
	"net/http"
	//"github.com/cloud-barista/cb-webtool/src/controller"
	"github.com/cloud-barista/cb-webtool/src/controller"
	echotemplate "github.com/foolin/echo-template"
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
func init() {

}

func main() {
	e := echo.New()

	e.Use(echosession.New())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"http://210.207.104.150"},
	// 	AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	// }))

	e.Static("/assets", "./src/static/assets")

	// paresGlob 를 사용하여 모든 경로에 있는 파일을 가져올 경우 사용하면 되겠다.
	// 사용한다음에 해당 파일을 불러오면 되네.
	// 서브디렉토리에 있는 걸 확인하기가 힘드네....
	renderer := &TemplateRender{
		templates: template.Must(template.ParseGlob(`./src/views/*.html`)),
	}
	e.Renderer = renderer

	// login 매핑할 middleware 추가
	loginTemplate := echotemplate.NewMiddleware(echotemplate.TemplateConfig{
		Root:      "src/views/auth",
		Extension: ".html",
		Master:    "Login",
		Partials: []string{
			"LoginTop",
			"SelectNamespaceModal",
			"LoginFooter",			
		},
		DisableCache: true,
	})

	// namespace 매핑할 middleware 추가
	namespaceTemplate := echotemplate.NewMiddleware(echotemplate.TemplateConfig{
		Root:      "src/views",
		Extension: ".html",
		Master:    "setting/namespaces/NameSpace",
		Partials: []string{
			"templates/Top",
			"templates/Top_box",
			"templates/LNB_popup",
			"templates/Modal",
			"templates/Header",
			"templates/Menu_left",
			"templates/Footer",
		}, //
		DisableCache: true,
	})
	// -> Master에 껍데기 및 header, footer 놓고, Partials에 해당 페이지에 들어가는 파일을 넣으면 될까?

	// dashboard 매핑할 middleware 추가
	dashboardTemplate := echotemplate.NewMiddleware(echotemplate.TemplateConfig{
		Root:         "src/views",
		Extension:    ".html",
		Master:       "operation/dashboards/Dashboard",
		Partials:     []string{
			"templates/Top",
			"templates/Top_box",
			"templates/LNB_popup",
			"templates/Modal",
			"templates/Header",
			"templates/Menu_left",
			"templates/Footer",		
		},
		DisableCache: true,
	})

	settingTemplate := echotemplate.NewMiddleware(echotemplate.TemplateConfig{
		Root:         "src/views",
		Extension:    ".html",
		Master:       "setting/connections",
		Partials:     []string{
			"templates/Top",
			"templates/Top_box",
			"templates/LNB_popup",
			"templates/Modal",
			"templates/Header",
			"templates/Menu_left",
			"templates/Footer",		
		},
		DisableCache: true,
	})



	// // mcis 매핑할 middleware 추가
	// manageMCISTemplate := echotemplate.NewMiddleware(echotemplate.TemplateConfig{
	// 	Root:         "src/views/operation/manage",
	// 	Extension:    ".html",
	// 	Master:       "src/layouts/master",
	// 	Partials:     []string{},
	// 	DisableCache: true,
	// })

	/////////////////////////
	// group에 templace set
	// 해당 그룹.GET(경로, controller의 method)
	// 해당 그룹.POST(경로, controller의 method)

	e.GET("/", controller.Index)

	// //connection 관련
	e.GET("/connectionconfig", controller.GetConnectionConfigData)

	loginGroup := e.Group("/login", loginTemplate)

	loginGroup.GET("", controller.LoginForm)
	loginGroup.POST("/proc", controller.LoginProc)
	loginGroup.POST("/process", controller.LoginProcess)

	//login 관련
	// e.GET("/login", controller.LoginForm)
	// // e.POST("/login/proc", controller.LoginController)
	// e.POST("/login/proc", controller.LoginProc)
	// // e.POST("/regUser", controller.RegUserConrtoller)
	// e.POST("/regUser", controller.RegUser)
	// e.GET("/logout", controller.LogoutForm)
	// e.GET("/logout/proc", controller.LogoutProc)

	// // Dashboard
	// e.GET("/Dashboard/Global", controller.GlobalDashBoard)
	// e.GET("/Dashboard/NS", controller.NSDashBoard)
	dashboardGroup := e.Group("/Dashboard", dashboardTemplate)
	dashboardGroup.GET("/NS", controller.NSDashBoard)
	dashboardGroup.GET("/Global", controller.GlobalDashBoard)

	// // Monitoring Control
	// e.GET("/Monitoring/MCIS/list", controller.MornitoringListForm)
	// e.GET("/Monitoring/mcis", controller.MornitoringListForm)
	// e.GET("/monitoring/install/agent/:mcis_id/:vm_id/:public_ip", controller.AgentRegForm)

	// // MCIS
	// e.GET("/Manage/MCIS/reg", controller.McisRegForm)
	// e.GET("/Manage/MCIS/reg/:mcis_id/:mcis_name", controller.VMAddForm)
	// e.POST("/Manage/MCIS/reg/proc", controller.McisRegController)
	// e.GET("/Manage/MCIS/list", controller.McisListForm)
	// e.GET("/Manage/MCIS/list/:mcis_id/:mcis_name", controller.McisListFormWithParam)

	// // Resource
	// e.GET("/Resource/board", controller.ResourceBoard)

	// // 웹툴에서 사용할 rest
	// e.GET("/SET/NS/:nsid", controller.SetNameSpace)

	// // 웹툴에서 처리할 NameSpace
	// // e.GET("/NameSpace/NS/list", controller.NsListForm)
	// // e.GET("/NS/reg", controller.NsRegForm)
	// // e.POST("/NS/reg/proc", controller.NsRegController)
	// // e.GET("/GET/ns", controller.GetNameSpace)
	namespaceGroup := e.Group("/NameSpace", namespaceTemplate)
	// namespaceGroup.GET("/NS/list", controller.NsListForm)          // namespace 보여주는 form 표시. DashboardController로 이동?
	// namespaceGroup.GET("/GET/ns", controller.GetNameSpace)         // 선택된 namespace 정보조회. Tumblebuck 호출
	// namespaceGroup.GET("/GET/nsList", controller.GetNameSpaceList) // 등록된 namespace 목록 조회. Tumblebuck 호출

	namespaceGroup.GET("/SET/NS/:nsid", controller.SetNameSpace) // default namespace set
	// namespaceGroup.GET("/NS/reg", controller.NsRegForm)          // namespace 등록 form 표시
	namespaceGroup.POST("/reg/proc", controller.NameSpaceRegProc)    // namespace 등록 처리


	settingGroup := e.Group("/setting", settingTemplate)
	settingGroup.GET("/connections/CloudConnection", controller.ConnectionList)
	// // 웹툴에서 처리할 Connection
	// e.GET("/Cloud/Connection/list", controller.ConnectionListForm)
	// e.GET("/Cloud/Connection/reg", controller.ConnectionRegForm)
	// e.POST("/Cloud/Connection/reg/proc", controller.NsRegController)

	// // 웹툴에서 처리할 Region
	// e.GET("/Region/list", controller.RegionListForm)
	// e.GET("/Region/reg", controller.RegionRegForm)
	// e.POST("/Region/reg/proc", controller.NsRegController)

	// // 웹툴에서 처리할 Credential
	// e.GET("/Credential/list", controller.CredertialListForm)
	// e.GET("/Credential/reg", controller.CredertialRegForm)

	// // 웹툴에서 처리할 Image
	// e.GET("/Image/list", controller.ImageListForm)
	// e.GET("/Image/reg", controller.ImageRegForm)

	// // 웹툴에서 처리할 VPC
	// e.GET("/Vpc/list", controller.VpcListForm)
	// e.GET("/Vpc/reg", controller.VpcRegForm)

	// // 웹툴에서 처리할 SecurityGroup
	// e.GET("/SecurityGroup/list", controller.SecurityGroupListForm)
	// e.GET("/SecurityGroup/reg", controller.SecurityGroupRegForm)

	// // 웹툴에서 처리할 Driver
	// e.GET("/Driver/list", controller.DriverListForm)
	// e.GET("/Driver/reg", controller.DriverRegForm)

	// // 웹툴에서 처리할 sshkey
	// e.GET("/SSH/list", controller.SSHListForm)
	// e.GET("/SSH/reg", controller.SSHRegForm)

	// // 웹툴에서 처리할 spec
	// e.GET("/Spec/list", controller.SpecListForm)
	// e.GET("/Spec/reg", controller.SpecRegForm)

	// // 웹툴에서 Select Pop
	// e.GET("/Pop/spec", controller.PopSpec)

	// // MAP Test
	// e.GET("/map", controller.Map)
	// e.GET("/map/geo/:mcis_id", controller.GeoInfo)

	e.Logger.Fatal(e.Start(":1234"))

}
