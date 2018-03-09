package controllers

import (
	"bytes"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"mycmdb/models"
	"mycmdb/utils"
	"strconv"
)

type IdcController struct {
	BaseController
}

func (this *IdcController) IdcIndexPage() {
	beego.ReadFromRequest(&this.Controller) // for flash

	o := orm.NewOrm()

	cnt, err := o.QueryTable(new(models.Idc)).Count()
	if err != nil {
		beego.Error(err)
	}

	names := []orm.Params{}
	if _, err := o.QueryTable(new(models.Idc)).Distinct().Values(&names, "Name"); err != nil {
		beego.Error(err)
	}

	operators := []orm.Params{}
	if _, err := o.QueryTable(new(models.Idc)).Distinct().Values(&operators, "Operator"); err != nil {
		beego.Error(err)
	}

	linkmans := []orm.Params{}
	if _, err := o.QueryTable(new(models.Idc)).Distinct().Values(&linkmans, "Linkman"); err != nil {
		beego.Error(err)
	}

	this.Data["cnt"] = cnt
	this.Data["names"] = &names
	this.Data["operators"] = &operators
	this.Data["linkmans"] = &linkmans
	this.TplName = "assets/idc/idcindex.tpl"
}

func (this *IdcController) IdcAddPage() {
	this.Data["title"] = "添加机房"
	this.TplName = "assets/idc/idcedit.tpl"
}

func (this *IdcController) IdcEditPage() {
	id, err := this.GetInt(":id")
	if err != nil {
		this.Abort("404")
	}

	o := orm.NewOrm()
	idc := models.Idc{Id: id}
	if err := o.Read(&idc); err != nil {
		beego.Error(err)
	}

	this.Data["idc"] = &idc
	this.Data["title"] = "修改机房"
	this.TplName = "assets/idc/idcedit.tpl"
}

func (this *IdcController) IdcDetailPage() {
	id, err := this.GetInt(":id")
	if err != nil {
		this.Abort("404")
	}

	o := orm.NewOrm()
	idc := models.Idc{Id: id}
	if err := o.Read(&idc); err != nil {
		beego.Error(err)
	}

	this.Data["idc"] = &idc
	this.TplName = "assets/idc/idcdetail.tpl"
}

func (this *IdcController) ListIdc() {
	secho, start, length, sort_th, sort_type := utils.ParseDataTableParams(this)

	order_key := []string{
		"",
		"Id",
		"Name",
		"Address",
		"Operator",
		"Linkman",
		"Phone",
		"CreateTime",
	}[sort_th]

	o := orm.NewOrm()
	qs := o.QueryTable(new(models.Idc))
	totalRecords, err := qs.Count()
	if err != nil {
		beego.Error(err)
		totalRecords = 0
	}

	if name := this.GetStrings("filter_name[]"); name != nil {
		qs = qs.Filter("Name__in", name)
	}
	if operator := this.GetStrings("filter_operator[]"); operator != nil {
		qs = qs.Filter("Operator__in", operator)
	}
	if linkman := this.GetStrings("filter_linkman[]"); linkman != nil {
		qs = qs.Filter("Linkman__in", linkman)
	}

	totalDisplayRecords, err := qs.Count()
	if err != nil {
		beego.Error(err)
		totalDisplayRecords = 0
	}

	qs = qs.OrderBy(sort_type + order_key)
	qs = qs.Offset(start).Limit(length)

	var maps []orm.Params
	num, err := qs.Values(&maps)
	if err != nil {
		beego.Error(err)
	}

	rs := make(map[string]interface{})
	rs["sEcho"] = secho
	rs["iTotalDisplayRecords"] = totalDisplayRecords
	rs["iTotalRecords"] = totalRecords
	if num == 0 {
		rs["aaData"] = []string{}
	} else {
		rs["aaData"] = maps
	}

	this.Data["json"] = &rs
	this.ServeJSON()
}

func (this *IdcController) SaveIdc() {
	idc := models.Idc{}
	flash := beego.NewFlash()

	if err := this.ParseForm(&idc); err != nil {
		beego.Error(err)
		flash.Error("参数解析错误: " + err.Error())
		flash.Store(&this.Controller)
		this.Redirect(this.URLFor("IdcController.IdcIndexPage"), 302)
		this.StopRun()
	}

	valid := validation.Validation{}
	if isValid, err := valid.Valid(&idc); err != nil {
		beego.Error(err)
		flash.Error("输入验证错误: " + err.Error())
		flash.Store(&this.Controller)
		this.Redirect(this.URLFor("IdcController.IdcIndexPage"), 302)
		this.StopRun()
	} else if !isValid {
		// validation does not pass
		var b bytes.Buffer
		for _, err := range valid.Errors {
			beego.Error(err.Key, err.Message)
			b.WriteString(err.Key + " " + err.Message + ".")
		}
		flash.Error("输入非法:" + b.String())
		flash.Store(&this.Controller)
		this.Redirect(this.URLFor("IdcController.IdcIndexPage"), 302)
		this.StopRun()
	}

	o := utils.MyNewOrm()
	if o.Read(&models.Idc{Id: idc.Id}) != nil {
		// Insert new record
		if _, err := o.Insert(&idc); err != nil {
			beego.Error(err)
			flash.Error("保存失败: " + err.Error())
			flash.Store(&this.Controller)
			this.Redirect(this.URLFor("IdcController.IdcIndexPage"), 302)
			this.StopRun()
		} else {
			flash.Notice("保存成功!")
			flash.Store(&this.Controller)
		}
	} else {
		// Update exist record
		if _, err := o.Update(this, &idc); err != nil {
			beego.Error(err)
			flash.Error("更新失败: " + err.Error())
			flash.Store(&this.Controller)
			this.Redirect(this.URLFor("IdcController.IdcIndexPage"), 302)
			this.StopRun()
		} else {
			flash.Notice("更新成功!")
			flash.Store(&this.Controller)
		}
	}

	this.Redirect(this.URLFor("IdcController.IdcIndexPage"), 302)
}

func (this *IdcController) DeleteIdc() {
	ids := this.GetStrings("ids[]")
	result := map[string]string{}

	beego.Debug(ids)

	o := orm.NewOrm()
	for _, id := range ids {
		if i, err := strconv.Atoi(id); err == nil {
			if _, err := o.Delete(&models.Idc{Id: i}); err != nil {
				beego.Error(err)
				result["info"] = "操作失败: " + err.Error()
			} else {
				result["info"] = "操作成功!"
			}
		}
	}

	this.Data["json"] = &result
	this.ServeJSON()
}
