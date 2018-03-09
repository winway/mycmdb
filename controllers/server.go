package controllers

import (
	"bytes"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"mycmdb/models"
	"mycmdb/utils"
)

type ServerController struct {
	BaseController
}

func (this *ServerController) ServerIndexPage() {
	beego.ReadFromRequest(&this.Controller) // for flash

	o := orm.NewOrm()

	serverCnt, err := o.QueryTable(new(models.Server)).Count()
	if err != nil {
		beego.Error(err)
	}

	usedServerCnt, err := o.QueryTable(new(models.Server)).Filter("Status", 2).Count()
	if err != nil {
		beego.Error(err)
	}

	var servers []*models.Server
	if _, err := o.QueryTable(new(models.Server)).RelatedSel("Idc").Distinct().All(&servers); err != nil {
		beego.Error(err)
	}

	nameSet := make(map[string]struct{})
	for _, v := range servers {
		nameSet[v.Idc.Name] = struct{}{}
	}

	brands := []orm.Params{}
	if _, err := o.QueryTable(new(models.Server)).Distinct().Values(&brands, "Brand"); err != nil {
		beego.Error(err)
	}

	this.Data["serverCnt"] = &serverCnt
	this.Data["usedServerCnt"] = &usedServerCnt
	this.Data["nameSet"] = &nameSet
	this.Data["brands"] = &brands
	this.TplName = "assets/server/serverindex.tpl"
}

func (this *ServerController) ServerAddPage() {
	o := orm.NewOrm()
	names := []orm.Params{}
	if _, err := o.QueryTable(new(models.Idc)).Distinct().Values(&names, "Id", "Name"); err != nil {
		beego.Error(err)
	}

	brands := []orm.Params{}
	if _, err := o.QueryTable(new(models.Server)).Distinct().Values(&brands, "Brand"); err != nil {
		beego.Error(err)
	}

	this.Data["names"] = &names
	this.Data["brands"] = &brands
	this.Data["title"] = "添加服务器"
	this.TplName = "assets/server/serveredit.tpl"
}

func (this *ServerController) ServerEditPage() {
	id := this.GetString(":id")
	if id == "" {
		this.Abort("404")
	}

	o := orm.NewOrm()
	server := models.Server{Sn: id}
	if err := o.Read(&server); err != nil {
		beego.Error(err)
		this.Abort("404")
	}

	names := []orm.Params{}
	if _, err := o.QueryTable(new(models.Idc)).Distinct().Values(&names, "Id", "Name"); err != nil {
		beego.Error(err)
	}

	this.Data["names"] = &names
	this.Data["server"] = &server
	this.Data["title"] = "修改服务器"
	this.TplName = "assets/server/serveredit.tpl"
}

func (this *ServerController) ServerDetailPage() {
	id := this.GetString(":id")
	if id == "" {
		this.Abort("404")
	}

	o := orm.NewOrm()
	var server models.Server
	if err := o.QueryTable(new(models.Server)).Filter("Sn", id).RelatedSel("Idc").One(&server); err != nil {
		beego.Error(err)
		this.Abort("404")
	}

	names := []orm.Params{}
	if _, err := o.QueryTable(new(models.Idc)).Distinct().Values(&names, "Id", "Name"); err != nil {
		beego.Error(err)
	}

	this.Data["names"] = &names

	this.Data["server"] = &server
	this.TplName = "assets/server/serverdetail.tpl"
}

func (this *ServerController) ListServer() {
	secho, start, length, sort_th, sort_type := utils.ParseDataTableParams(this)
	order_key := []string{
		"",
		"Sn",
		"Idc",
		"HostName",
		"RemoteCardIp",
		"Sn",
		"Sn",
		"OsVersion",
		"Brand",
		"Status",
		"CreateTime",
	}[sort_th]

	o := orm.NewOrm()
	qs := o.QueryTable(new(models.Server))
	totalRecords, err := qs.Count()
	if err != nil {
		beego.Error(err)
		totalRecords = 0
	}

	if keyword := this.GetString("filter_keyword"); keyword != "" {
		cond := orm.NewCondition()
		condKeyword := cond.Or("Sn__contains", keyword).Or("Hostname__contains", keyword).
			Or("RemoteCardIp__contains", keyword).Or("Eth1Ip__contains", keyword).
			Or("Eth2Ip__contains", keyword).Or("Eth3Ip__contains", keyword).
			Or("Eth4Ip__contains", keyword).Or("Cpu__contains", keyword).
			Or("Disk__contains", keyword).Or("Memory__contains", keyword)
		qs = qs.SetCond(cond.AndCond(condKeyword))
	}

	if idc := this.GetStrings("filter_idc[]"); idc != nil {
		qs = qs.Filter("Idc__Name__in", idc)
	}
	if brand := this.GetStrings("filter_brand[]"); brand != nil {
		qs = qs.Filter("Brand__in", brand)
	}
	if osVersion := this.GetStrings("filter_osversion[]"); osVersion != nil {
		qs = qs.Filter("OsVersion__in", osVersion)
	}
	if status := this.GetStrings("filter_status[]"); status != nil {
		qs = qs.Filter("Status__in", status)
	}

	totalDisplayRecords, err := qs.Count()
	if err != nil {
		beego.Error(err)
		totalDisplayRecords = 0
	}

	qs = qs.RelatedSel("Idc")
	qs = qs.OrderBy(sort_type + order_key)
	qs = qs.Offset(start).Limit(length)

	var maps []*models.Server
	num, err := qs.All(&maps)
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

func (this *ServerController) SaveServer() {
	server := models.Server{}

	result := map[string]interface{}{"code": "0", "msg": "ok"}

	if err := this.ParseForm(&server); err != nil {
		beego.Error(err)
		result["code"] = "1"
		result["msg"] = "参数解析错误: " + err.Error()
		this.Data["json"] = &result
		this.ServeJSON()
		this.StopRun()
	}

	valid := validation.Validation{}
	if isValid, err := valid.Valid(&server); err != nil {
		beego.Error(err)
		result["code"] = "1"
		result["msg"] = "输入验证错误: " + err.Error()
		this.Data["json"] = &result
		this.ServeJSON()
		this.StopRun()
	} else if !isValid {
		// validation does not pass
		var b bytes.Buffer
		for _, err := range valid.Errors {
			beego.Error(err.Key, err.Message)
			b.WriteString(err.Key + " " + err.Message + ".")
		}
		result["code"] = "1"
		result["msg"] = "输入非法:" + b.String()
		this.Data["json"] = &result
		this.ServeJSON()
		this.StopRun()
	}

	o := orm.NewOrm()
	idc := models.Idc{}
	idcId, err := this.GetInt("Idc")
	if err != nil {
		beego.Error(err)
		result["code"] = "1"
		result["msg"] = "获取机房ID错误: " + err.Error()
		this.Data["json"] = &result
		this.ServeJSON()
		this.StopRun()
	} else {
		idc = models.Idc{Id: idcId}
		if err := o.Read(&idc); err != nil {
			beego.Error(err)
			result["code"] = "1"
			result["msg"] = "获取机房ID错误: " + err.Error()
			this.Data["json"] = &result
			this.ServeJSON()
			this.StopRun()
		}
	}

	server.Idc = &idc
	if o.Read(&models.Server{Sn: server.Sn}) != nil {
		// auto allocate RemoteCardIp
		ips, err := utils.GetAvailIp(idcId, 0, 1)
		if err != nil {
			beego.Error(err)
			result["code"] = "1"
			result["msg"] = "保存失败: " + err.Error()
			this.Data["json"] = &result
			this.ServeJSON()
			this.StopRun()
		}

		server.RemoteCardIp = ips[0]

		// Insert new record
		_, err = o.Insert(&server)
		if err != nil {
			beego.Error(err)
			result["code"] = "1"
			result["msg"] = "保存失败: " + err.Error()
			this.Data["json"] = &result
			this.ServeJSON()
			this.StopRun()
		}

		// osInstall := models.OsInstallManidests{Server: &server, OsVersion: server.OsVersion, Status: 0, Comment: "新机器首次安装"}
		// _, err = o.Insert(&osInstall)
		// if err != nil {
		// 	beego.Error(err)
		// 	flash := beego.NewFlash()
		// 	flash.Error("OS安装申请失败: " + err.Error())
		// 	flash.Store(&this.Controller)
		// }
	} else {
		// Update exist record
		if _, err := o.Update(&server); err != nil {
			beego.Error(err)
			result["code"] = "1"
			result["msg"] = "更新失败: " + err.Error()
			this.Data["json"] = &result
			this.ServeJSON()
			this.StopRun()
		}
	}

	this.Data["json"] = &result
	this.ServeJSON()
}

func (this *ServerController) DeleteServer() {
	ids := this.GetStrings("ids[]")
	result := map[string]string{}

	o := orm.NewOrm()
	for _, id := range ids {
		if _, err := o.Delete(&models.Server{Sn: id}); err != nil {
			beego.Error(err)
			result["info"] = "操作失败: " + err.Error()
		} else {
			result["info"] = "操作成功!"
		}
	}

	this.Data["json"] = &result
	this.ServeJSON()
}
