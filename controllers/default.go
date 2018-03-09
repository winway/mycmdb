package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Prepare() {
	name := this.GetSession("name")
	if name == nil && this.Ctx.Request.RequestURI != this.URLFor("UserController.LoginPage") && this.Ctx.Request.RequestURI != this.URLFor("UserController.RegisterPage") {
		flash := beego.NewFlash()
		flash.Notice(this.Ctx.Request.RequestURI)
		flash.Store(&this.Controller)
		this.Redirect(this.URLFor("UserController.LoginPage"), 302)
		return
	}

	this.SetSession("url", this.Ctx.Request.RequestURI)
	controllerName, actionName := this.GetControllerAndAction()
	this.SetSession("action", controllerName+"."+actionName)

	sbStatus := this.GetSession("sbStatus")
	if sbStatus == nil {
		sbStatus = false
		this.SetSession("sbStatus", sbStatus)
	}
	smStatus := this.GetSession("smStatus")
	if smStatus == nil {
		smStatus = false
		this.SetSession("smStatus", smStatus)
	}
	umStatus := this.GetSession("umStatus")
	if umStatus == nil {
		umStatus = false
		this.SetSession("umStatus", umStatus)
	}

	this.Data["uname"] = name
	this.Data["username"] = this.GetSession("username")
	this.Data["isadmin"] = this.GetSession("isadmin")
	this.Data["sbStatus"] = sbStatus
	this.Data["smStatus"] = smStatus
	this.Data["umStatus"] = umStatus
}

type MainController struct {
	BaseController
}

func (this *MainController) Get() {
	this.TplName = "dashbord.tpl"
}
