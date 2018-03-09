package routers

import (
	"github.com/astaxie/beego"
	"mycmdb/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/dashbord/", &controllers.MainController{})

	beego.Router("/user/login", &controllers.UserController{}, "Get:LoginPage;Post:Login")
	beego.Router("/user/logout", &controllers.UserController{}, "Get:Logout")
	beego.Router("/user/register", &controllers.UserController{}, "Get:RegisterPage;Post:Register")
	beego.Router("/user/menu", &controllers.UserController{}, "Post:ToggleMenuStatus")

	beego.Router("/assets/idcindex/", &controllers.IdcController{}, "Get:IdcIndexPage")
	beego.Router("/assets/idcadd/", &controllers.IdcController{}, "Get:IdcAddPage")
	beego.Router("/assets/idcedit/:id", &controllers.IdcController{}, "Get:IdcEditPage")
	beego.Router("/assets/idcdetail/:id", &controllers.IdcController{}, "Get:IdcDetailPage")
	beego.Router("/assets/idc/", &controllers.IdcController{}, "Get:ListIdc;Post:SaveIdc")
	beego.Router("/assets/idc_delete/", &controllers.IdcController{}, "Post:DeleteIdc")

	beego.Router("/assets/ipindex/", &controllers.IpController{}, "Get:IpIndexPage")
	beego.Router("/assets/ipadd/", &controllers.IpController{}, "Get:IpAddPage")
	beego.Router("/assets/ipapply/", &controllers.IpController{}, "Get:IpApplyPage")
	beego.Router("/assets/ipdetail/:id", &controllers.IpController{}, "Get:IpDetailPage")
	beego.Router("/assets/ip/", &controllers.IpController{}, "Get:ListIp;Post:SaveIp")
	beego.Router("/assets/ip_delete/", &controllers.IpController{}, "Post:DeleteIp")
	beego.Router("/assets/ip_preapply/", &controllers.IpController{}, "Post:PreApplyIp")
	beego.Router("/assets/ip_apply/", &controllers.IpController{}, "Post:ApplyIp")

	beego.Router("/assets/serverindex/", &controllers.ServerController{}, "Get:ServerIndexPage")
	beego.Router("/assets/serveradd/", &controllers.ServerController{}, "Get:ServerAddPage")
	beego.Router("/assets/serveredit/:id", &controllers.ServerController{}, "Get:ServerEditPage")
	beego.Router("/assets/serverdetail/:id", &controllers.ServerController{}, "Get:ServerDetailPage")
	beego.Router("/assets/server/", &controllers.ServerController{}, "Get:ListServer;Post:SaveServer")
	beego.Router("/assets/server_delete/", &controllers.ServerController{}, "Post:DeleteServer")

	beego.Router("/osinstall/index/", &controllers.OsInstallController{}, "Get:IndexPage")
	beego.Router("/osinstall/apply/", &controllers.OsInstallController{}, "Get:ApplyPage")
	beego.Router("/osinstall/manifest", &controllers.OsInstallController{}, "Get:List;Post:Save")
	beego.Router("/osinstall/serverinfo", &controllers.OsInstallController{}, "Get:GetServerInfo")
	beego.Router("/osinstall/cancel/", &controllers.OsInstallController{}, "Post:Cancel")

	beego.Router("/release/index/", &controllers.ReleaseController{}, "Get:IndexPage")
	beego.Router("/release/apply/", &controllers.ReleaseController{}, "Get:ApplyPage")
	beego.Router("/release/applylist", &controllers.ReleaseController{}, "Get:List;Post:Save")
	beego.Router("/release/detail/:id", &controllers.ReleaseController{}, "Get:ApplyDetailPage;Post:Handle")
	beego.Router("/release/view/:id", &controllers.ReleaseController{}, "Get:ApplyViewPage")
	beego.Router("/release/cancel/", &controllers.ReleaseController{}, "Post:Cancel")

	beego.Router("/api/osinstall/bare", &controllers.ApiController{}, "Get:GetBareRemoteCardIp")
	beego.Router("/api/osinstall/update/hw/", &controllers.ApiController{}, "Post:UpdateHWInfo")
	beego.Router("/api/osinstall/manifests", &controllers.ApiController{}, "Get:GetManifests")
	beego.Router("/api/osinstall/update/status/:id", &controllers.ApiController{}, "Post:UpdateStatus")
	// beego.Router("/api/osinstall/mac", &controllers.ApiController{}, "Get:GetRemoteCardMac;Post:ReportRemoteCardIp")
	// beego.Router("/api/osinstall/availip", &controllers.ApiController{}, "Get:GetAvailIp")
}
