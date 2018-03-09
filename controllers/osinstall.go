package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"mycmdb/models"
	"mycmdb/utils"
	"regexp"
	"strconv"
	"strings"
)

type OsInstallController struct {
	BaseController
}

func (this *OsInstallController) IndexPage() {
	// beego.ReadFromRequest(&this.Controller) // for flash

	o := orm.NewOrm()
	doneCnt, err := o.QueryTable(new(models.OsInstallManifests)).Filter("Status", 2).Count()
	if err != nil {
		beego.Error(err)
	}

	doingCnt, err := o.QueryTable(new(models.OsInstallManifests)).Filter("Status", 1).Count()
	if err != nil {
		beego.Error(err)
	}

	this.Data["doneCnt"] = doneCnt
	this.Data["doingCnt"] = doingCnt
	this.TplName = "osinstall/index.tpl"
}

func (this *OsInstallController) ApplyPage() {
	o := orm.NewOrm()
	// m := []orm.Params{}
	// if _, err := o.QueryTable(new(models.Server)).RelatedSel("OsInstallManifests").Filter("OsInstallManifests__Status__ne", 2).Values(&m, "Sn"); err != nil {
	// 	beego.Error(err)
	// }

	var servers []models.Server
	_, err := o.Raw("SELECT T0.* FROM `server` T0 LEFT JOIN `os_install_manifests` T1 ON T1.`server_id` = T0.`sn` WHERE T1.`status` IS NULL OR (T1.`status` != ? AND T1.`status` != ?)", 0, 1).QueryRows(&servers)
	if err != nil {
		beego.Error(err)
	}

	this.Data["snSet"] = &servers
	this.TplName = "osinstall/apply.tpl"
}

func (this *OsInstallController) List() {
	secho, start, length, sort_th, sort_type := utils.ParseDataTableParams(this)
	order_key := []string{
		"",
		"Id",
		"Sn",
		"Id",
		"Id",
		"OsVersion",
		"Status",
		"CreateTime",
	}[sort_th]

	o := orm.NewOrm()
	qs := o.QueryTable(new(models.OsInstallManifests))
	totalRecords, err := qs.Count()
	if err != nil {
		beego.Error(err)
		totalRecords = 0
	}

	if ip := this.GetString("filter_ip"); ip != "" {
		cond := orm.NewCondition()
		condip := cond.Or("Server__RemoteCardIp__contains", ip).Or("Server__Eth1Ip__contains", ip).
			Or("Server__Eth2Ip__contains", ip).Or("Server__Eth3Ip__contains", ip).
			Or("Server__Eth4Ip__contains", ip)
		qs = qs.SetCond(cond.AndCond(condip))
	}
	if osVersion := this.GetStrings("filter_osversion[]"); osVersion != nil {
		qs = qs.Filter("Server__OsVersion__in", osVersion)
	}
	if status := this.GetStrings("filter_status[]"); status != nil {
		qs = qs.Filter("Status__in", status)
	}

	totalDisplayRecords, err := qs.Count()
	if err != nil {
		beego.Error(err)
		totalDisplayRecords = 0
	}

	qs = qs.RelatedSel("Server")
	qs = qs.OrderBy(sort_type + order_key)
	qs = qs.Offset(start).Limit(length)

	var maps []*models.OsInstallManifests
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

func (this *OsInstallController) Save() {
	result := map[string]interface{}{"code": "0", "msg": "ok"}

	sn := this.GetString("Sn")
	if sn == "" {
		beego.Error("Sn is null")
		result["code"] = "1"
		result["msg"] = "Sn is null"
		this.Data["json"] = &result
		this.ServeJSON()
		this.StopRun()
	}

	osVersion, err := this.GetInt("OsVersion")
	if err != nil {
		beego.Error(err)
		result["code"] = "1"
		result["msg"] = err.Error()
		this.Data["json"] = &result
		this.ServeJSON()
		this.StopRun()
	}

	DiskNum, err := this.GetInt("Disknum")
	if err != nil {
		beego.Error(err)
		result["code"] = "1"
		result["msg"] = err.Error()
		this.Data["json"] = &result
		this.ServeJSON()
		this.StopRun()
	}

	raidConfigTmp := map[string]map[string]interface{}{}

	digit, err := regexp.Compile("^[0-9]+$")
	if err != nil {
		beego.Error(err)
		result["code"] = "1"
		result["msg"] = err
		this.Data["json"] = &result
		this.ServeJSON()
		this.StopRun()
	}
	digit_digit, err := regexp.Compile("^[0-9]+_[0-9]+$")
	if err != nil {
		beego.Error(err)
		result["code"] = "1"
		result["msg"] = err
		this.Data["json"] = &result
		this.ServeJSON()
		this.StopRun()
	}

	for k, v := range this.Ctx.Request.Form {
		if digit.MatchString(k) {
			raidConfigTmp[k] = map[string]interface{}{"raid": v[0], "disk": []string{}}
		}
	}

	for k, v := range this.Ctx.Request.Form {
		if digit_digit.MatchString(k) {
			raidConfigTmp[strings.Split(k, "_")[0]]["disk"] = append(raidConfigTmp[strings.Split(k, "_")[0]]["disk"].([]string), v[0])
		}
	}

	raidConfig := []map[string]interface{}{}
	for _, v := range raidConfigTmp {
		if len(v["disk"].([]string)) > 0 {
			raidConfig = append(raidConfig, v)
		}
	}

	beego.Error(DiskNum)
	beego.Error(raidConfig)

	cnt := 0
	isValid := true
	errMessage := ""
	for _, i := range raidConfig {
		cnt += len(i["disk"].([]string))

		switch i["raid"] {
		case "raid1":
			if len(i["disk"].([]string)) < 2 || len(i["disk"].([]string))%2 == 1 {
				isValid = false
				errMessage = "raid1必须大于2块磁盘且必须是偶数块"
			}
		case "raid10":
			if len(i["disk"].([]string)) < 4 || len(i["disk"].([]string))%2 == 1 {
				isValid = false
				errMessage = "raid10必须大于4块硬盘且必须是偶数块磁盘"
			}
		case "raid5":
			if len(i["disk"].([]string)) < 4 {
				isValid = false
				errMessage = "raid5 必须大于4块硬盘"
			}
		default:
			isValid = false
			errMessage = "未知的raid类型"
		}

		diskType := ""
		for _, j := range i["disk"].([]string) {
			if diskType != "" && diskType != strings.Split(j, "*")[1] {
				isValid = false
				errMessage = "做raid磁盘组 必须是同类型的磁盘"
			}
			diskType = strings.Split(j, "*")[1]
		}

		if !isValid {
			beego.Error("raid config invalid")
			result["code"] = "1"
			result["msg"] = errMessage
			this.Data["json"] = &result
			this.ServeJSON()
			this.StopRun()
		}
	}

	if cnt > DiskNum {
		beego.Error("duplicate error")
		result["code"] = "1"
		result["msg"] = "磁盘重复选择"
		this.Data["json"] = &result
		this.ServeJSON()
		this.StopRun()
	}

	if cnt < DiskNum {
		beego.Error("free disk error")
		result["code"] = "1"
		result["msg"] = "有磁盘未选择"
		this.Data["json"] = &result
		this.ServeJSON()
		this.StopRun()
	}

	raidConfigString, err := json.Marshal(raidConfig)
	if err != nil {
		beego.Error(err)
		result["code"] = "1"
		result["msg"] = err
		this.Data["json"] = &result
		this.ServeJSON()
		this.StopRun()
	}

	// var b bytes.Buffer
	// var s string
	// for _, v := range raidConfig {
	// 	if len(v["disk"].([]string)) > 0 {
	// 		b.WriteString(s + v["raid"].(string) + ":" + strings.Join(v["disk"].([]string), ","))
	// 		if s == "" {
	// 			s = ";"
	// 		}
	// 	}
	// }
	// beego.Error(b.String())

	o := orm.NewOrm()
	server := models.Server{Sn: sn}
	if err := o.Read(&server); err != nil {
		beego.Error(err)
		result["code"] = "1"
		result["msg"] = "获取服务器信息错误: " + err.Error()
		this.Data["json"] = &result
		this.ServeJSON()
		this.StopRun()
	} else {
		server.OsVersion = int8(osVersion)
		server.RaidConfig = string(raidConfigString)
		if _, err := o.Update(&server, "OsVersion", "RaidConfig"); err != nil {
			beego.Error(err)
			result["code"] = "1"
			result["msg"] = err.Error()
			this.Data["json"] = &result
			this.ServeJSON()
			this.StopRun()
		}
	}

	m := models.OsInstallManifests{}
	m.Server = &server
	m.Comment = this.GetString("Comment")
	_, err = o.Insert(&m)
	if err != nil {
		beego.Error(err)
		result["code"] = "1"
		result["msg"] = "保存失败: " + err.Error()
		this.Data["json"] = &result
		this.ServeJSON()
		this.StopRun()
	}

	this.Data["json"] = &result
	this.ServeJSON()
}

func (this *OsInstallController) GetServerInfo() {
	result := map[string]interface{}{"code": "0", "msg": "ok"}

	sn := this.GetString("Sn")
	if sn == "" {
		beego.Error("Sn is null")
		result["code"] = "1"
		result["msg"] = "Sn is null"
		this.Data["json"] = &result
		this.ServeJSON()
		this.StopRun()
	}

	o := orm.NewOrm()
	var maps models.Server
	qs := o.QueryTable(new(models.Server)).Filter("Sn", sn).RelatedSel("Idc")
	if err := qs.One(&maps); err != nil {
		beego.Error(err)
		result["code"] = "1"
		result["msg"] = err.Error()
		this.Data["json"] = &result
		this.ServeJSON()
		this.StopRun()
	}

	result["data"] = maps

	this.Data["json"] = &result
	this.ServeJSON()
}

func (this *OsInstallController) Cancel() {
	ids := this.GetStrings("ids[]")
	result := map[string]string{}

	o := orm.NewOrm()
	for _, id := range ids {
		if i, err := strconv.Atoi(id); err == nil {
			m := models.OsInstallManifests{Id: i}
			if err := o.Read(&m); err != nil {
				beego.Error(err)
				result["info"] = "获取安装列表信息错误: " + err.Error()
				this.Data["json"] = &result
				this.ServeJSON()
				this.StopRun()
			}
			if m.Status == 1 {
				if _, err := o.Delete(&models.OsInstallManifests{Id: i}); err != nil {
					beego.Error(err)
					result["info"] = "操作失败: " + err.Error()
				} else {
					result["info"] = "操作成功!"
				}
			} else {
				result["info"] = "无法取消!"
			}
		}
	}

	this.Data["json"] = &result
	this.ServeJSON()
}
