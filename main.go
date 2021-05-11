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

	// TODO : navigation Template을 만들어서 공통으로 Set하면 page redirect 할 때 편하지 않을 까??
	// login 매핑할 middleware 추가
	loginTemplate := echotemplate.NewMiddleware(echotemplate.TemplateConfig{
		Root:      "src/views",
		Extension: ".html",
		Master:    "auth/Login",
		Partials: []string{
			"auth/LoginTop",
			"auth/SelectNamespaceModal",
			"auth/LoginFooter",
		},
		DisableCache: true,
	})

	defaultTemplate := echotemplate.NewMiddleware(echotemplate.TemplateConfig{
		Root:      "src/views",
		Extension: ".html",
		// Master:    "auth/Login",
		Partials: []string{
			"templates/Top",
			"templates/TopBox",
			"templates/MenuLeft",
			"templates/Header",
			"templates/Footer",
		},
		DisableCache: true,
	})

	mainTemplate := echotemplate.NewMiddleware(echotemplate.TemplateConfig{
		Root:      "src/views",
		Extension: ".html",
		// Master:    "setting/namespaces/NameSpaceMng",
		Partials: []string{
			"templates/Top",
			"templates/TopBox",
			"templates/LNB",
			"templates/LNBPopup",
			"templates/Modal",
			"templates/Header",
			"templates/MenuLeft",
			"templates/Footer",
		}, //
		DisableCache: true,
	})

	// namespace 매핑할 middleware 추가
	namespaceTemplate := echotemplate.NewMiddleware(echotemplate.TemplateConfig{
		Root:      "src/views",
		Extension: ".html",
		// Master:    "setting/namespaces/NameSpaceMng",
		Partials: []string{
			"templates/Top",
			"templates/TopBox",
			"templates/LNB",
			"templates/LNBPopup",
			"templates/Modal",
			"templates/Header",
			"templates/MenuLeft",
			"templates/Footer",
		}, //
		DisableCache: true,
	})
	// -> Master에 껍데기 및 header, footer 놓고, Partials에 해당 페이지에 들어가는 파일을 넣으면 될까?

	// dashboard 매핑할 middleware 추가
	dashboardTemplate := echotemplate.NewMiddleware(echotemplate.TemplateConfig{
		Root:      "src/views",
		Extension: ".html",
		// Master:    "operation/dashboards/Dashboard",
		Partials: []string{
			"templates/OperationTop",
			"templates/TopBox",
			"templates/LNBPopup",
			"templates/Modal",
			"templates/Header",
			"templates/MenuLeft",
			"templates/Footer",
		},
		DisableCache: true,
	})

	monitoringTemplate := echotemplate.NewMiddleware(echotemplate.TemplateConfig{
		Root:      "src/views",
		Extension: ".html",
		// Master:    "operation/dashboards/Dashboard",
		Partials: []string{
			"templates/OperationTop",
			"templates/TopBox",
			"templates/LNBPopup",
			"templates/Modal",
			"templates/Header",
			"templates/MenuLeft",
			"templates/Footer",
		},
		DisableCache: true,
	})

	policyTemplate := echotemplate.NewMiddleware(echotemplate.TemplateConfig{
		Root:      "src/views",
		Extension: ".html",
		// Master:    "operation/dashboards/Policy",
		Partials: []string{
			"templates/OperationTop",
			"templates/TopBox",
			"templates/LNBPopup",
			"templates/Modal",
			"templates/Header",
			"templates/MenuLeft",
			"templates/Footer",
		},
		DisableCache: true,
	})

	mcisTemplate := echotemplate.NewMiddleware(echotemplate.TemplateConfig{
		Root:      "src/views",
		Extension: ".html",
		// Master:    "operation/mcis/mcismng",
		Partials: []string{
			"templates/OperationTop", // 불러오는 css, javascript 가 setting 과 다름
			"templates/TopBox",
			"templates/LNBPopup",
			"templates/Modal",
			"templates/Header",
			"templates/MenuLeft",
			"templates/Footer", // TODO : McisCreate 파일에서 가져오는 partials는 다른 경로인데 어떻게 불러오지?

			"operation/manages/mcis/McisStatus",
			"operation/manages/mcis/McisList",
			"operation/manages/mcis/McisInfo",
			"operation/manages/mcis/McisServerInfo",
			"operation/manages/mcis/McisDetailInfo",
			"operation/manages/mcis/McisDetailView",
			"operation/manages/mcis/McisConnectionView",
			"operation/manages/mcis/McisMonitoring",
			"operation/manages/mcis/McisMonitoringView",
		},
		DisableCache: true,
	})

	// MCIS 등록 form Template
	mcisRegTemplate := echotemplate.NewMiddleware(echotemplate.TemplateConfig{
		Root:      "src/views",
		Extension: ".html",
		// Master:    "operation/mcis/mcismng",
		Partials: []string{
			"templates/OperationTop", // 불러오는 css, javascript 가 setting 과 다름
			"templates/TopBox",
			"templates/LNBPopup",
			"templates/Modal",
			"templates/Header",
			"templates/MenuLeft",
			"templates/Footer", // TODO : McisCreate 파일에서 가져오는 partials는 다른 경로인데 어떻게 불러오지?

			"operation/manages/mcis/McisVmConfigureSimple",
			"operation/manages/mcis/McisVmConfigureExpert",
			"operation/manages/mcis/McisVmConfigureImport",

			"operation/manages/mcis/McisAssistPopup",

			"operation/manages/mcis/McisOsHardware",
			"operation/manages/mcis/McisNetwork",
			"operation/manages/mcis/McisSecurity",
			"operation/manages/mcis/McisOther",
		},
		DisableCache: true,
	})

	// MCKS 등록 form Template
	mcksRegTemplate := echotemplate.NewMiddleware(echotemplate.TemplateConfig{
		Root:      "src/views",
		Extension: ".html",
		// Master:    "operation/mcis/mcismng",
		Partials: []string{
			"templates/OperationTop", // 불러오는 css, javascript 가 setting 과 다름
			"templates/TopBox",
			"templates/LNBPopup",
			"templates/Modal",
			"templates/Header",
			"templates/MenuLeft",
			"templates/Footer", // TODO : McisCreate 파일에서 가져오는 partials는 다른 경로인데 어떻게 불러오지?

			"operation/manages/mcks/McksNodeConfigure",
		},
		DisableCache: true,
	})

	cloudConnectionTemplate := echotemplate.NewMiddleware(echotemplate.TemplateConfig{
		Root:      "src/views",
		Extension: ".html",
		// Master:    "setting/connections/CloudConnectionConfigMng", // master를 이용할 때는 확장자 없이. 그 외에는 확장자까지
		Partials: []string{
			"templates/Top",
			"templates/TopBox",
			"templates/LNBPopup",
			"templates/MenuLeft",
			"templates/Header",
			"templates/Modal",
			"setting/connections/cloud/RegionModal",
			"setting/connections/cloud/CredentialModal",
			"setting/connections/cloud/DriverModal",
			"templates/Footer",
		},
		DisableCache: true,
	})

	resourceTemplate := echotemplate.NewMiddleware(echotemplate.TemplateConfig{
		Root:      "src/views",
		Extension: ".html",
		// Master:    "setting/resources/NetworkMng", // master를 이용할 때는 확장자 없이. 그 외에는 확장자까지, 여러화면에서 사용할 때에는 return할 때 master지정
		Partials: []string{
			"templates/Top",
			"templates/TopBox",
			"templates/LNBPopup",
			"templates/MenuLeft",
			"templates/Header",
			"templates/Modal",
			"templates/Footer",
		},
		DisableCache: true,
	})
	// "setting/connections/CloudConnectionModal", --> Region, Credential, Driver modal로 쪼개짐

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
	// naming rule : json인 경우 Get 등을 앞에 붙임 ex) GetConnectionConfigData
	// controller : 1개일 때 객체명, List일 때 객체명 + List
	//   ex) 1개 가져오는 json : GetConnectionConfigData, List 가져오는 json : GetConnectionConfigList
	// handler : 1개일때 controller명 + Data, List일 때 controller method명 DataList

	e.GET("/", controller.Index)
	defaultGroup := e.Group("/operation/about", defaultTemplate)
	defaultGroup.GET("/about", controller.About)

	//e.GET("/apicall", controller.ApiCall)
	e.POST("/apicall", controller.ApiCall)
	e.POST("/servercall", controller.ServerCall)

	mainGroup := e.Group("/main", mainTemplate)
	mainGroup.GET("", controller.MainForm)
	mainGroup.GET("/apitestmng", controller.ApiTestMngForm)

	loginGroup := e.Group("/login", loginTemplate)

	loginGroup.GET("", controller.LoginForm)
	loginGroup.POST("/proc", controller.LoginProc)
	e.GET("/logout", controller.LogoutProc)
	// loginGroup.POST("/process", controller.LoginProcess)
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
	dashboardGroup := e.Group("/operation/dashboards", dashboardTemplate)

	dashboardGroup.GET("/dashboardnamespace/mngform", controller.DashBoardByNameSpaceMngForm)
	dashboardGroup.GET("/dashboardglobalnamespace/mngform", controller.GlobalDashBoardMngForm)

	// // Monitoring Control
	// e.GET("/Monitoring/MCIS/list", controller.MornitoringListForm)
	// e.GET("/Monitoring/mcis", controller.MornitoringListForm)
	// e.GET("/monitoring/install/agent/:mcis_id/:vm_id/:public_ip", controller.AgentRegForm)
	monitoringGroup := e.Group("/operation/monitorings", monitoringTemplate)
	monitoringGroup.GET("/mcis/mngform", controller.McisMonitoringMngForm)
	monitoringGroup.GET("/mcis/:mcisID/vm/:vmID/agent/mngform", controller.VmMonitoringAgentRegForm)
	monitoringGroup.POST("/mcis/:mcisID/vm/:vmID/agent/reg/proc", controller.VmMonitoringAgentRegProc) // namespace 등록 처리
	monitoringGroup.GET("/mcis/:mcisID/metric/:metric", controller.GetVmMonitoringInfoData)

	policyGroup := e.Group("/operation/policies", policyTemplate)
	policyGroup.GET("/policy/mngform", controller.McisPolicyMngForm)
	policyGroup.GET("/mcis/:mcisID/vm/:vmID/agent/mngform", controller.VmMonitoringAgentRegForm)
	policyGroup.POST("/mcis/:mcisID/vm/:vmID/agent/reg/proc", controller.VmMonitoringAgentRegProc) // namespace 등록 처리
	policyGroup.GET("/mcis/:mcisID/metric/:metric", controller.GetVmMonitoringInfoData)

	// MCIS
	// e.GET("/Manage/MCIS/reg", controller.McisRegForm)
	// e.GET("/Manage/MCIS/reg/:mcis_id/:mcis_name", controller.VmAddForm)
	// e.POST("/Manage/MCIS/reg/proc", controller.McisRegController)
	// e.GET("/Manage/MCIS/list", controller.McisListForm)
	// e.GET("/Manage/MCIS/list/:mcis_id/:mcis_name", controller.McisListFormWithParam)

	// mcis에 form이 2개가 되면서 group을 나눔. json return은 굳이 group이 필요없어서 전체경로로 작음.
	mcisGroup := e.Group("/operation/manages/mcis/mngform", mcisTemplate)
	// e.GET("/mcis/reg", controller.McisRegForm)
	// e.GET("/mcis/reg/:mcis_id/:mcis_name", controller.VmAddForm)
	// e.POST("/mcis/reg/proc", controller.McisRegController)

	// mcisGroup.GET("/", controller.McisMngForm)
	mcisGroup.GET("", controller.McisMngForm)

	e.GET("/operation/manages/mcis/list", controller.GetMcisList) // 등록된 namespace의 MCIS 목록 조회. Tumblebuck 호출
	e.POST("/operation/manages/mcis/reg/proc", controller.McisRegProc)

	// TODO : namespace는 서버에 저장된 것을 사용하는데... 자칫하면 namespace와 다른 mcis의 vm으로 날아갈 수 있지 않나???
	e.GET("/operation/manages/mcis/:mcisID", controller.GetMcisInfoData)

	e.POST("/operation/manages/mcis/:mcisID/vm/reg/proc", controller.VmRegProc) // vm 등록이므로 vmID없이 reg/proc
	e.GET("/operation/manages/mcis/:mcisID/vm/:vmID", controller.GetVmInfoData)

	e.POST("/operation/manages/mcis/proc/mcislifecycle", controller.McisLifeCycle)
	//var url = "/operation/manage" + "/mcis/" + mcisID + "/operation/" + type
	e.POST("/operation/manages/mcis/proc/vmlifecycle", controller.McisVmLifeCycle)
	e.POST("/operation/manages/mcis/proc/vmmonitoring", controller.GetVmMonitoring)

	// e.POST("/operation/manages/mcis/proc/vmmonitoring", controller.GetVmMonitoring)

	// e.GET("/mcis/list/:mcis_id/:mcis_name", controller.McisListFormWithParam)

	//http://54.248.3.145:1234/Manage/MCIS/reg/mz-azure-mcis/mz-azure-mcis
	mcisRegGroup := e.Group("/operation/manages/mcis/regform", mcisRegTemplate)
	// mcisRegGroup.GET("/", controller.McisRegForm)                    // MCIS 생성 + VM생성
	mcisRegGroup.GET("", controller.McisRegForm)                     // MCIS 생성 + VM생성
	mcisRegGroup.GET("/:mcisID/:mcisName", controller.McisVmRegForm) // MCIS의 VM생성

	mcksMngGroup := e.Group("/operation/manages/mcks/mngform", mcksRegTemplate)
	mcksMngGroup.GET("", controller.McisRegForm) // MCKS 생성 + Node생성

	mcksRegGroup := e.Group("/operation/manages/mcks/regform", mcksRegTemplate)
	mcksRegGroup.GET("", controller.McisRegForm) // Node생성

	// // Resource
	// e.GET("/Resource/board", controller.ResourceBoard)

	// // 웹툴에서 사용할 rest
	// e.GET("/SET/NS/:nsid", controller.SetNameSpace)

	// // 웹툴에서 처리할 NameSpace
	// // e.GET("/NameSpace/NS/list", controller.NsListForm)
	// // e.GET("/NS/reg", controller.NsRegForm)
	// // e.POST("/NS/reg/proc", controller.NsRegController)
	// // e.GET("/GET/ns", controller.GetNameSpace)
	namespaceGroup := e.Group("/setting/namespaces", namespaceTemplate)
	namespaceGroup.GET("/namespace/mngform", controller.NameSpaceMngForm)      // namespace 보여주는 form 표시. DashboardController로 이동?
	namespaceGroup.GET("/namespace/list", controller.GetNameSpaceList)         // 등록된 namespace 목록 조회. Tumblebuck 호출
	namespaceGroup.GET("/namespace/set/:nameSpaceID", controller.SetNameSpace) // default namespace set
	namespaceGroup.POST("/namespace/reg/proc", controller.NameSpaceRegProc)    // namespace 등록 처리
	// namespaceGroup.PUT("/namespace/update/proc", controller.NameSpaceUpdateProc)// namespace 수정 : TB에 해당기능없음.
	namespaceGroup.DELETE("/namespace/del/:nameSpaceID", controller.NameSpaceDelProc) // namespace 삭제 처리

	cloudConnectionGroup := e.Group("/setting/connections", cloudConnectionTemplate)
	// cloudConnectionGroup.GET("/connections/cloudos", controller.GetCloudOSList) // TODO : 사용 안하는듯. 필요없으면 제거
	// cloudConnectionGroup.GET("/connections/CloudConnection", controller.ConnectionList) // Connection 관리화면 -> naming rule변경으로사용안함.
	cloudConnectionGroup.GET("/cloudconnectionconfig/mngform", controller.CloudConnectionConfigMngForm)            // Connection 관리화면
	cloudConnectionGroup.GET("/cloudconnectionconfig/list", controller.GetCloudConnectionConfigList)               //connection 목록 조회
	cloudConnectionGroup.GET("/cloudconnectionconfig/:configName", controller.GetCloudConnectionConfigData)        //connection 정보 상세조회
	cloudConnectionGroup.POST("/cloudconnectionconfig/reg/proc", controller.CloudConnectionConfigRegProc)          // 등록
	cloudConnectionGroup.DELETE("/cloudconnectionconfig/del/:configName", controller.CloudConnectionConfigDelProc) // 삭제

	// region form은 popup으로 대체
	cloudConnectionGroup.GET("/region", controller.GetRegionList)     // Region 목록 조회
	cloudConnectionGroup.GET("/region/:region", controller.GetRegion) // Region 조회
	cloudConnectionGroup.POST("/region/reg/proc", controller.RegionRegProc)
	cloudConnectionGroup.DELETE("/region/del/:region", controller.RegionDelProc)

	// credection form은 popup으로 대체
	cloudConnectionGroup.GET("/credential", controller.GetCredentialList)
	cloudConnectionGroup.GET("/credential/:credential", controller.GetCredential) // Credential 조회
	cloudConnectionGroup.POST("/credential/reg/proc", controller.CredentialRegProc)
	cloudConnectionGroup.DELETE("/credential/del/:credential", controller.CredentialDelProc)

	// driver form은 popup으로 대체
	cloudConnectionGroup.GET("/driver", controller.GetDriverList)
	cloudConnectionGroup.GET("/driver/:driver", controller.GetDriver) // Driver 조회
	cloudConnectionGroup.POST("/driver/reg/proc", controller.DriverRegProc)
	cloudConnectionGroup.DELETE("/driver/del/:driver", controller.DriverDelProc)

	// config form은 popup으로 대체(UI 정의되지 않음.)
	cloudConnectionGroup.GET("/config", controller.GetConfigList)
	cloudConnectionGroup.GET("/config/:configID", controller.GetConfig) // Config 조회
	cloudConnectionGroup.POST("/config/reg/proc", controller.ConfigRegProc)
	cloudConnectionGroup.DELETE("/config/del/:configID", controller.ConfigDelProc)

	resourcesGroup := e.Group("/setting/resources", resourceTemplate)
	e.POST("/getinspectresources", controller.GetInspectResourceList)
	resourcesGroup.GET("/network/mngform", controller.VpcMngForm)
	resourcesGroup.GET("/network/list", controller.GetVpcList)
	resourcesGroup.GET("/network/:vNetID", controller.GetVpcData)
	resourcesGroup.POST("/network/reg", controller.VpcRegProc)
	resourcesGroup.DELETE("/network/del/:vNetID", controller.VpcDelProc)

	resourcesGroup.GET("/securitygroup/mngform", controller.SecirityGroupMngForm)
	resourcesGroup.GET("/securitygroup/list", controller.GetSecirityGroupList)
	resourcesGroup.GET("/securitygroup/:securityGroupID", controller.GetSecirityGroupData)
	resourcesGroup.POST("/securitygroup/reg", controller.SecirityGroupRegProc)
	resourcesGroup.DELETE("/securitygroup/del/:securityGroupID", controller.SecirityGroupDelProc)

	resourcesGroup.GET("/sshkey/mngform", controller.SshKeyMngForm) // Form + SshKeyMng 같이 앞으로 넘길까?
	resourcesGroup.GET("/sshkey/list", controller.GetSshKeyList)
	resourcesGroup.GET("/sshkey/:sshKeyID", controller.GetSshKeyData)
	resourcesGroup.POST("/sshkey/reg", controller.SshKeyRegProc)             // RegProc _ SshKey 같이 앞으로 넘길까
	resourcesGroup.DELETE("/sshkey/del/:sshKeyID", controller.SshKeyDelProc) // DelProc + SskKey 같이 앞으로 넘길까

	resourcesGroup.GET("/machineimage/mngform", controller.VirtualMachineImageMngForm) // Form + SshKeyMng 같이 앞으로 넘길까?
	resourcesGroup.GET("/machineimage/list", controller.GetVirtualMachineImageList)
	resourcesGroup.GET("/machineimage/:imageID", controller.GetVirtualMachineImageData)
	resourcesGroup.POST("/machineimage/reg", controller.VirtualMachineImageRegProc)            // RegProc _ SshKey 같이 앞으로 넘길까
	resourcesGroup.DELETE("/machineimage/del/:imageID", controller.VirtualMachineImageDelProc) // DelProc + SskKey 같이 앞으로 넘길까

	resourcesGroup.GET("/machineimage/lookupimages", controller.LookupCspVirtualMachineImageList) // TODO : Image 전체목록인가? 확인필요
	//resourcesGroup.GET("/machineimage/lookupimage", controller.LookupVirtualMachineImageList)          // TODO : Image 전체목록인가? 확인필요
	resourcesGroup.GET("/machineimage/lookupimage/:imageID", controller.LookupVirtualMachineImageData) // TODO : Image 상세 정보인가? 확인필요
	resourcesGroup.GET("/machineimage/fetchimage", controller.FetchVirtualMachineImageList)            // TODO : Image 정보 갱신인가? 확인필요

	resourcesGroup.GET("/machineimage/searchimage", controller.SearchVirtualMachineImageList)

	resourcesGroup.GET("/vmspec/mngform", controller.VmSpecMngForm) // Form + SshKeyMng 같이 앞으로 넘길까?
	resourcesGroup.GET("/vmspec/list", controller.GetVmSpecList)
	resourcesGroup.GET("/vmspec/:vmSpecID", controller.GetVmSpecData)
	resourcesGroup.POST("/vmspec/reg", controller.VmSpecRegProc)             // RegProc _ SshKey 같이 앞으로 넘길까
	resourcesGroup.DELETE("/vmspec/del/:vmSpecID", controller.VmSpecDelProc) // DelProc + SskKey 같이 앞으로 넘길까
	// resourcesGroup.PUT("/vmspec/put/:vmSpecID", controller.VmSpecPutProc)	// TODO : put 만들어야 함

	resourcesGroup.GET("/vmspec/lookupvmspec", controller.LookupVmSpecList)             // TODO : Image 전체목록인가? 확인필요
	resourcesGroup.GET("/vmspec/lookupvmspec/:vmSpecName", controller.LookupVmSpecData) // TODO : Image 상세 정보인가? 확인필요
	resourcesGroup.GET("/vmspec/fetchvmspec", controller.FetchVmSpecList)               // TODO : Image 정보 갱신인가? 확인필요
	// resourcesGroup.POST("/vmspec/filterspecs", controller.FilterVmSpecList)	// TODO : post방식의 filterspec 생성필요
	// resourcesGroup.POST("/vmspec/filterspecsbyrange", controller.FilterVmSpecListByRange)// TODO : post방식의 filterspec 생성필요

	// e.GET("/SecurityGroup/list", controller.SecurityGroupListForm)
	// e.GET("/SecurityGroup/reg", controller.SecurityGroupRegForm)

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
