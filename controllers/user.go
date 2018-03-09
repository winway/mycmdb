package controllers

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"io"
	"mycmdb/models"
	"time"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Prepare() {
	name := this.GetSession("name")
	if name == nil && this.Ctx.Request.RequestURI != this.URLFor("UserController.LoginPage") && this.Ctx.Request.RequestURI != this.URLFor("UserController.RegisterPage") {
		this.Redirect(this.URLFor("UserController.LoginPage"), 302)
		return
	}

	if name != nil && this.Ctx.Request.RequestURI == this.URLFor("UserController.LoginPage") && this.Ctx.Request.RequestURI != this.URLFor("UserController.RegisterPage") {
		this.Redirect(this.URLFor("MainController.Get"), 302)
	}
}

func (this *UserController) LoginPage() {
	flash := beego.ReadFromRequest(&this.Controller) // for flash
	this.Data["referer"] = flash.Data["notice"]
	this.TplName = "user/login.tpl"
}

func (this *UserController) RegisterPage() {
	beego.ReadFromRequest(&this.Controller) // for flash
	this.TplName = "user/register.tpl"
}

// 登陆处理
func (this *UserController) Login() {
	flash := beego.NewFlash()

	rf := this.GetString("rf")
	if rf == "" {
		rf = this.URLFor("MainController.Get")
	}

	beego.Error(rf)

	email := this.GetString("email")
	if email == "" {
		beego.Error("Email is null")
		flash.Error("Email is null")
		flash.Store(&this.Controller)
		this.Redirect(this.URLFor("UserController.LoginPage"), 302)
		this.StopRun()
	}

	password := this.GetString("password")
	if password == "" {
		beego.Error("Password is null")
		flash.Error("Password is null")
		flash.Store(&this.Controller)
		this.Redirect(this.URLFor("UserController.LoginPage"), 302)
		this.StopRun()
	}

	md5Password := md5.New()
	io.WriteString(md5Password, password)
	buffer := bytes.NewBuffer(nil)
	fmt.Fprintf(buffer, "%x", md5Password.Sum(nil))
	newPass := buffer.String()

	beego.Error(password, newPass)

	o := orm.NewOrm()
	u := models.User{Email: email}
	if err := o.Read(&u); err != nil {
		beego.Error(err)
		flash.Error("登陆失败: " + err.Error())
		flash.Store(&this.Controller)
		this.Redirect(this.URLFor("UserController.LoginPage"), 302)
		this.StopRun()
	}

	if u.Password == newPass {
		u.LastLoginTime = time.Now()

		if _, err := o.Update(&u, "LastLoginTime"); err != nil {
			beego.Error(err)
			flash.Error("登陆失败: " + err.Error())
			flash.Store(&this.Controller)
			this.Redirect(this.URLFor("UserController.LoginPage"), 302)
			this.StopRun()
		}

		this.SetSession("name", u.Email)
		this.SetSession("username", u.Name)
		this.SetSession("isadmin", u.IsAdmin)

		this.Redirect(rf, 302)
	} else {
		flash.Error("邮箱或密码错误")
		flash.Store(&this.Controller)
		this.Redirect(this.URLFor("UserController.LoginPage"), 302)
	}
}

func (this *UserController) Register() {
	result := map[string]interface{}{"code": "0", "msg": "ok"}

	email := this.GetString("email")
	name := this.GetString("name")
	password := this.GetString("password")
	if email == "" || name == "" || password == "" {
		result["code"] = "1"
		result["msg"] = "参数不全"
		this.Data["json"] = &result
		this.ServeJSON()
		this.StopRun()
	}

	md5Password := md5.New()
	io.WriteString(md5Password, password)
	buffer := bytes.NewBuffer(nil)
	fmt.Fprintf(buffer, "%x", md5Password.Sum(nil))
	newPass := buffer.String()

	o := orm.NewOrm()
	user := models.User{Email: email, Name: name, Password: newPass, LastLoginTime: time.Now()}

	if o.Read(&user) != nil {
		// Insert new record
		_, err := o.Insert(&user)
		if err != nil {
			beego.Error(err)
			result["code"] = "1"
			result["msg"] = "申请失败: " + err.Error()
			this.Data["json"] = &result
			this.ServeJSON()
			this.StopRun()
		}
	} else {
		result["code"] = "1"
		result["msg"] = "用户已存在"
		this.Data["json"] = &result
		this.ServeJSON()
		this.StopRun()
	}

	this.Data["json"] = &result
	this.ServeJSON()
}

func (this *UserController) Logout() {
	this.DelSession("name")
	this.Redirect(this.URLFor("UserController.LoginPage"), 302)
}

func (this *UserController) ToggleMenuStatus() {
	this.EnableRender = false

	menuType := this.GetString("type")

	switch menuType {
	case "sb":
		sbStatus := this.GetSession("sbStatus")
		beego.Error(sbStatus)
		if sbStatus == nil {
			sbStatus = false
		} else {
			sbStatus = !sbStatus.(bool)
		}
		this.SetSession("sbStatus", sbStatus)
		beego.Error(sbStatus)
	case "um":
		umStatus := this.GetSession("umStatus")
		beego.Error(umStatus)
		if umStatus == nil {
			umStatus = false
		} else {
			umStatus = !umStatus.(bool)
		}
		this.SetSession("umStatus", umStatus)
		beego.Error(umStatus)
	case "sm":
		smStatus := this.GetSession("smStatus")
		beego.Error(smStatus)
		if smStatus == nil {
			smStatus = false
		} else {
			smStatus = !smStatus.(bool)
		}
		this.SetSession("smStatus", smStatus)
		beego.Error(smStatus)
	}
}
