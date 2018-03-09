package controllers

import (
	"bytes"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"mycmdb/models"
	"mycmdb/utils"
	"strconv"
	"strings"
)

type IpController struct {
	BaseController
}

func (this *IpController) IpIndexPage() {
	beego.ReadFromRequest(&this.Controller) // for flash

	o := orm.NewOrm()

	ipCnt, err := o.QueryTable(new(models.Ip)).Count()
	if err != nil {
		beego.Error(err)
	}

	usedIpCnt, err := o.QueryTable(new(models.Ip)).Filter("Status", 1).Count()
	if err != nil {
		beego.Error(err)
	}

	var ips []*models.Ip
	if _, err := o.QueryTable(new(models.Ip)).RelatedSel("Idc").Distinct().All(&ips); err != nil {
		beego.Error(err)
	}

	nameSet := make(map[string]struct{})
	for _, v := range ips {
		nameSet[v.Idc.Name] = struct{}{}
	}

	networks := []orm.Params{}
	if _, err := o.QueryTable(new(models.Ip)).Distinct().Values(&networks, "Network"); err != nil {
		beego.Error(err)
	}

	this.Data["ipCnt"] = ipCnt
	this.Data["usedIpCnt"] = usedIpCnt
	this.Data["nameSet"] = &nameSet
	this.Data["networks"] = &networks
	this.Data["networksCnt"] = len(networks)
	this.TplName = "assets/ip/ipindex.tpl"
}

func (this *IpController) IpAddPage() {
	o := orm.NewOrm()
	names := []orm.Params{}
	if _, err := o.QueryTable(new(models.Idc)).Distinct().Values(&names, "Id", "Name"); err != nil {
		beego.Error(err)
	}

	this.Data["names"] = &names
	this.Data["title"] = "添加IP"
	this.TplName = "assets/ip/ipedit.tpl"
}

func (this *IpController) IpApplyPage() {
	o := orm.NewOrm()
	names := []orm.Params{}
	if _, err := o.QueryTable(new(models.Idc)).Distinct().Values(&names, "Id", "Name"); err != nil {
		beego.Error(err)
	}

	this.Data["names"] = &names
	this.TplName = "assets/ip/ipapply.tpl"
}

func (this *IpController) IpDetailPage() {
	id := this.GetString(":id")
	if id == "" {
		this.Abort("404")
	}

	o := orm.NewOrm()
	var ip models.Ip
	if err := o.QueryTable(new(models.Ip)).Filter("Ip", id).RelatedSel("Idc").One(&ip); err != nil {
		beego.Error(err)
	}

	beego.Error(ip)

	this.Data["ip"] = &ip
	this.Data["prefix"] = strings.Split(ip.Network, "/")[1]
	this.TplName = "assets/ip/ipdetail.tpl"
}

func (this *IpController) ListIp() {
	secho, start, length, sort_th, sort_type := utils.ParseDataTableParams(this)

	order_key := []string{
		"",
		"Ip",
		"Idc",
		"Network",
		"IpType",
		"Status",
		"CreateTime",
	}[sort_th]

	o := orm.NewOrm()
	qs := o.QueryTable(new(models.Ip))
	totalRecords, err := qs.Count()
	if err != nil {
		beego.Error(err)
		totalRecords = 0
	}

	if ip := this.GetString("filter_ip"); ip != "" {
		qs = qs.Filter("Ip__contains", ip)
	}
	if idc := this.GetStrings("filter_idc[]"); idc != nil {
		qs = qs.Filter("Idc__Name__in", idc)
	}
	if network := this.GetStrings("filter_network[]"); network != nil {
		qs = qs.Filter("Network__in", network)
	}
	if ipType := this.GetStrings("filter_iptype[]"); ipType != nil {
		qs = qs.Filter("IpType__in", ipType)
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

	var maps []*models.Ip
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

func (this *IpController) SaveIp() {
	ip := models.Ip{}
	flash := beego.NewFlash()

	if err := this.ParseForm(&ip); err != nil {
		beego.Error(err)
		flash.Error("参数解析错误: " + err.Error())
		flash.Store(&this.Controller)
		this.Redirect(this.URLFor("IpController.IpIndexPage"), 302)
		this.StopRun()
	}

	valid := validation.Validation{}
	if isValid, err := valid.Valid(&ip); err != nil {
		beego.Error(err)
		flash.Error("输入验证错误: " + err.Error())
		flash.Store(&this.Controller)
		this.Redirect(this.URLFor("IpController.IpIndexPage"), 302)
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
		this.Redirect(this.URLFor("IpController.IpIndexPage"), 302)
		this.StopRun()
	}

	o := orm.NewOrm()
	idc := models.Idc{}
	if idcId, err := this.GetInt("Idc"); err != nil {
		beego.Error(err)
		flash.Error("获取机房ID错误: " + err.Error())
		flash.Store(&this.Controller)
		this.Redirect(this.URLFor("IpController.IpIndexPage"), 302)
		this.StopRun()
	} else {
		idc = models.Idc{Id: idcId}
		if err := o.Read(&idc); err != nil {
			beego.Error(err)
			flash.Error("获取机房ID错误: " + err.Error())
			flash.Store(&this.Controller)
			this.Redirect(this.URLFor("IpController.IpIndexPage"), 302)
			this.StopRun()
		}
	}

	ipSlice, network, err := utils.GetIPFromCIDR(ip.Network)
	if err != nil {
		beego.Error(err)
		flash.Error("解析CIDR错误: " + err.Error())
		flash.Store(&this.Controller)
		this.Redirect(this.URLFor("IpController.IpIndexPage"), 302)
		this.StopRun()
	}

	beego.Error(ipSlice, network, err)

	ips := []models.Ip{}
	for _, i := range ipSlice {
		ip.Ip = i
		ip.Idc = &idc
		ip.Network = network
		ips = append(ips, ip)
	}

	beego.Error(ips)

	successNums, err := o.InsertMulti(1000, ips)
	if err != nil {
		beego.Error(err)
		flash.Error("导入IP错误: " + err.Error() + "需导入" + strconv.Itoa(len(ips)) + ", 成功导入" + strconv.FormatInt(successNums, 10))
		flash.Store(&this.Controller)
		this.Redirect(this.URLFor("IpController.IpIndexPage"), 302)
		this.StopRun()
	} else {
		flash.Notice("导入成功!需导入" + strconv.Itoa(len(ips)) + ", 成功导入" + strconv.FormatInt(successNums, 10))
		flash.Store(&this.Controller)
	}

	this.Redirect(this.URLFor("IpController.IpIndexPage"), 302)
}

func (this *IpController) DeleteIp() {
	ids := this.GetStrings("ids[]")
	result := map[string]string{}

	beego.Debug(ids)

	o := orm.NewOrm()
	for _, id := range ids {
		if _, err := o.Delete(&models.Ip{Ip: id}); err != nil {
			beego.Error(err)
			result["info"] = "操作失败: " + err.Error()
		} else {
			result["info"] = "操作成功!"
		}
	}

	this.Data["json"] = &result
	this.ServeJSON()
}

func (this *IpController) PreApplyIp() {
	result := map[string]interface{}{"code": "0", "msg": "ok"}

	idc, err := this.GetInt("idc")
	if err != nil {
		beego.Error(err)
		result["code"] = "1"
		result["msg"] = err.Error()
		this.Data["json"] = &result
		this.ServeJSON()
		this.StopRun()
	}

	ipType, err := this.GetInt("ipType")
	if err != nil {
		beego.Error(err)
		result["code"] = "1"
		result["msg"] = err.Error()
		this.Data["json"] = &result
		this.ServeJSON()
		this.StopRun()
	}

	num, err := this.GetInt("num")
	if err != nil {
		beego.Error(err)
		num = 1
	} else if num < 0 {
		num = 1
	}

	o := orm.NewOrm()
	qs := o.QueryTable(new(models.Ip)).Filter("IpType", ipType).Filter("Status__ne", 1).RelatedSel("Idc").Filter("Idc__Id", idc).Offset(0).Limit(num)

	var maps []*models.Ip
	if n, err := qs.All(&maps, "Ip"); err != nil {
		beego.Error(err)
		result["code"] = "1"
		result["msg"] = err.Error()
		this.Data["json"] = &result
		this.ServeJSON()
		this.StopRun()
	} else if int(n) < num {
		result["code"] = "1"
		result["msg"] = "Not Enough IP"
		this.Data["json"] = &result
		this.ServeJSON()
		this.StopRun()
	}

	ips := []string{}
	for _, v := range maps {
		ips = append(ips, v.Ip)
	}
	result["data"] = ips

	this.Data["json"] = &result
	this.ServeJSON()
}

func (this *IpController) ApplyIp() {
	result := map[string]interface{}{"code": "0", "msg": "ok"}

	ips := this.GetStrings("ips[]")
	mail := this.GetString("mail")

	valid := validation.Validation{}
	valid.Required(ips, "ips")
	valid.Email(mail, "mail")

	if valid.HasErrors() {
		var b bytes.Buffer
		for _, err := range valid.Errors {
			beego.Error(err.Key, err.Message)
			b.WriteString(err.Key + " " + err.Message + ".")
		}
		result["code"] = "1"
		result["msg"] = b.String()
		this.Data["json"] = &result
		this.ServeJSON()
		this.StopRun()
	}

	o := orm.NewOrm()
	errflag := false
	o.Begin()
	for _, i := range ips {
		ip := models.Ip{Ip: i}
		if err := o.Read(&ip); err != nil {
			beego.Error(err)
			errflag = true
			o.Rollback()
			result["code"] = "1"
			result["msg"] = err.Error()
			this.Data["json"] = &result
			this.ServeJSON()
			this.StopRun()
		} else {
			ip.Status = 1
			if _, err := o.Update(&ip, "Status"); err != nil {
				beego.Error(err)
				errflag = true
				o.Rollback()
				result["code"] = "1"
				result["msg"] = err.Error()
				this.Data["json"] = &result
				this.ServeJSON()
				this.StopRun()
			}
		}
	}

	if err := utils.SendMail("winway1988@163.com", "IP申请", "<p>您已成功申请如下IP，请放心使用！</p><br />"+strings.Join(ips, "<br />"), "html"); err != nil {
		beego.Error(err)
		errflag = true
		o.Rollback()
		result["code"] = "1"
		result["msg"] = err.Error()
		this.Data["json"] = &result
		this.ServeJSON()
		this.StopRun()
	}

	if !errflag {
		o.Commit()
		flash := beego.NewFlash()
		flash.Notice("IP申请成功!" + strings.Join(ips, ","))
		flash.Store(&this.Controller)
		this.Data["json"] = &result
		this.ServeJSON()
	}
}
