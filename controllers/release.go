package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"mycmdb/models"
	"mycmdb/utils"
	"strconv"
	"strings"
	"time"
)

type ReleaseController struct {
	BaseController
}

func (this *ReleaseController) IndexPage() {
	o := orm.NewOrm()
	doneCnt, err := o.QueryTable(new(models.ReleaseApply)).Filter("Status", 1).Count()
	if err != nil {
		beego.Error(err)
	}

	doingCnt, err := o.QueryTable(new(models.ReleaseApply)).Filter("Status", 0).Count()
	if err != nil {
		beego.Error(err)
	}

	this.Data["doneCnt"] = doneCnt
	this.Data["doingCnt"] = doingCnt
	this.TplName = "release/index.tpl"
}

func (this *ReleaseController) ApplyPage() {
	this.TplName = "release/apply.tpl"
}

func (this *ReleaseController) ApplyDetailPage() {
	id, err := this.GetInt(":id")
	if err != nil {
		beego.Error(err)
		this.Abort("404")
	}

	o := orm.NewOrm()
	emails := []orm.Params{}
	if _, err := o.QueryTable(new(models.User)).Distinct().Values(&emails, "Email"); err != nil {
		beego.Error(err)
	}

	beego.Error(emails)

	apply := models.ReleaseApply{Id: id}
	if err := o.Read(&apply); err != nil {
		beego.Error(err)
		this.Abort("404")
	}

	beego.Error(apply)

	// wtf ?
	var steps []models.ReleaseStep
	_, err = o.QueryTable(new(models.ReleaseStep)).RelatedSel("Apply").Filter("Apply__Id", id).OrderBy("StepId").All(&steps)
	if err != nil {
		beego.Error(err)
	}

	beego.Error(steps)

	var currentStep models.ReleaseStep
	err = o.QueryTable(new(models.ReleaseStep)).RelatedSel("Apply").Filter("Apply__Id", id).Filter("Status__ne", "已完成").OrderBy("StepId").Limit(1).One(&currentStep)
	if err != nil {
		beego.Error(err)
	}

	beego.Error(currentStep)

	this.Data["emails"] = &emails
	this.Data["apply"] = &apply
	this.Data["steps"] = &steps
	this.Data["currentStep"] = &currentStep
	this.Data["email"] = this.GetSession("name")
	this.TplName = "release/applydetail.tpl"
}

func (this *ReleaseController) ApplyViewPage() {
	id := this.GetString(":id")
	if id == "" {
		this.Abort("404")
	}

	o := orm.NewOrm()
	step := models.ReleaseStep{}
	if _, err := o.QueryTable(new(models.ReleaseStep)).Filter("StepId", id).All(&step); err != nil {
		beego.Error(err)
	}

	beego.Error(step)

	this.Data["step"] = &step
	this.TplName = "release/applyview.tpl"
}

func (this *ReleaseController) List() {
	secho, start, length, sort_th, sort_type := utils.ParseDataTableParams(this)
	order_key := []string{
		"Id",
		"Subject",
		"Creator",
		"Status",
		"CreateTime",
	}[sort_th]

	o := orm.NewOrm()
	qs := o.QueryTable(new(models.ReleaseApply))
	totalRecords, err := qs.Count()
	if err != nil {
		beego.Error(err)
		totalRecords = 0
	}

	if key := this.GetString("filter_key"); key != "" {
		cond := orm.NewCondition()
		condkey := cond.Or("Subject__contains", key).Or("Creator__contains", key)
		qs = qs.SetCond(cond.AndCond(condkey))
	}
	if status := this.GetStrings("filter_status[]"); status != nil {
		qs = qs.Filter("Status__in", status)
	}

	totalDisplayRecords, err := qs.Count()
	if err != nil {
		beego.Error(err)
		totalDisplayRecords = 0
	}

	qs = qs.OrderBy(sort_type + order_key)
	qs = qs.Offset(start).Limit(length)

	var maps []*models.ReleaseApply
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

func (this *ReleaseController) Save() {
	result := map[string]interface{}{"code": "0", "msg": "ok"}

	subject := this.GetString("Subject")
	if subject == "" {
		beego.Error("subject is null")
		result["code"] = "1"
		result["msg"] = "subject is null"
		this.Data["json"] = &result
		this.ServeJSON()
		this.StopRun()
	}

	applyType, err := this.GetInt("applyType")
	if err != nil {
		beego.Error(err)
		result["code"] = "1"
		result["msg"] = err.Error()
		this.Data["json"] = &result
		this.ServeJSON()
		this.StopRun()
	}

	o := orm.NewOrm()
	apply := models.ReleaseApply{}
	apply.Subject = subject
	apply.Creator = this.GetSession("name").(string)
	apply.ApplyType = applyType
	id, err := o.Insert(&apply)
	if err != nil {
		beego.Error(err)
		result["code"] = "1"
		result["msg"] = "保存失败: " + err.Error()
		this.Data["json"] = &result
		this.ServeJSON()
		this.StopRun()
	}

	apply = models.ReleaseApply{Id: int(id)}
	if err := o.Read(&apply); err != nil {
		beego.Error(err)
		result["code"] = "1"
		result["msg"] = "保存失败: " + err.Error()
		this.Data["json"] = &result
		this.ServeJSON()
		this.StopRun()
	}

	// 根据类型，创建step
	// 开始-产品-开发-开发负责人-测试-运维-测试-完结
	// 开始-开发-开发负责人-测试-运维-测试-完结

	steps := map[int][]string{
		0: {"【产品】提交上线功能说明", "【开发】确认并提供相关服务信息",
			"【开发负责人】二次审批", "【测试】确认线下测试通过并提交测试报告",
			"【运维】进行上线", "【测试】确认线上验证通过"},
		1: {"【开发】提供相关服务信息", "【开发负责人】二次审批",
			"【测试】确认线下测试通过并提交测试报告", "【运维】进行上线",
			"【测试】确认线上验证通过"},
		2: {"【开发】提供相关服务信息", "【开发负责人】二次审批",
			"【测试】确认线下测试通过并提交测试报告", "【运维】进行上线",
			"【测试】确认线上验证通过"},
	}

	for i, j := range steps[applyType] {
		step := models.ReleaseStep{}

		step.StepId = strconv.Itoa(int(id)) + "-" + strconv.Itoa(i)
		step.Apply = &apply
		step.Description = j
		if i == 0 {
			step.Owner = this.GetSession("name").(string)
		}
		if i == len(steps[applyType])-1 {
			step.IsLast = 1
		}
		step.Status = "进行中"
		if _, err = o.Insert(&step); err != nil {
			beego.Error(err)
			result["code"] = "1"
			result["msg"] = "保存失败: " + err.Error()
			this.Data["json"] = &result
			this.ServeJSON()
			this.StopRun()
		}
	}

	applyTypeMap := map[int]string{
		0: "新功能上线",
		1: "Bug修复",
		2: "功能点优化",
	}

	go utils.SendMail(this.GetSession("name").(string), applyTypeMap[applyType], "名称："+subject+"<br/>详情见：<a href='http://172.16.103.9:8866/release/index'>http://172.16.103.9:8866/release/index</a>", "html")

	this.Data["json"] = &result
	this.ServeJSON()
}

func (this *ReleaseController) Handle() {
	result := map[string]interface{}{"code": "0", "msg": "ok"}

	applyId, err := this.GetInt("applyId")
	if err != nil {
		beego.Error("applyId is null")
		result["code"] = "1"
		result["msg"] = "applyId is null"
		this.Data["json"] = &result
		this.ServeJSON()
		this.StopRun()
	}

	currentStep := this.GetString("currentStep")
	if currentStep == "" {
		beego.Error("currentStep is null")
		result["code"] = "1"
		result["msg"] = "currentStep is null"
		this.Data["json"] = &result
		this.ServeJSON()
		this.StopRun()
	}

	content := this.GetString("Content")

	nextOwner := this.GetString("nextOwner")
	if nextOwner == "" {
		beego.Error("nextOwner is null")
		result["code"] = "1"
		result["msg"] = "nextOwner is null"
		this.Data["json"] = &result
		this.ServeJSON()
		this.StopRun()
	}

	f, h, err := this.GetFile("myfile")
	var filename string
	if f != nil {
		if err != nil {
			beego.Error(err)
			result["code"] = "1"
			result["msg"] = err.Error()
			this.Data["json"] = &result
			this.ServeJSON()
			this.StopRun()
		}
		filename = h.Filename + "." + strconv.Itoa(int(time.Now().Unix()))
		f.Close()
		this.SaveToFile("myfile", "/data/go/src/mycmdb/static/upload/"+filename)
	}

	o := orm.NewOrm()
	step := models.ReleaseStep{StepId: currentStep}
	step.Content = content
	step.File = filename
	step.Status = "已完成"
	step.OperateTime = time.Now().Local().Format("2006-01-02 15:04:05.000")
	if _, err := o.Update(&step, "Content", "File", "Status", "OperateTime"); err != nil {
		beego.Error(err)
		result["code"] = "1"
		result["msg"] = "保存失败: " + err.Error()
		this.Data["json"] = &result
		this.ServeJSON()
		this.StopRun()
	}

	idx, _ := strconv.Atoi(strings.Split(currentStep, "-")[1])
	nextStep := strings.Split(currentStep, "-")[0] + "-" + strconv.Itoa(idx+1)

	step = models.ReleaseStep{StepId: nextStep}
	if err := o.Read(&step); err != nil {
		if err == orm.ErrNoRows {
			apply := models.ReleaseApply{Id: applyId}
			apply.Status = 1
			if _, err := o.Update(&apply, "Status"); err != nil {
				beego.Error(err)
				result["code"] = "1"
				result["msg"] = "保存失败: " + err.Error()
				this.Data["json"] = &result
				this.ServeJSON()
				this.StopRun()
			}
		} else {
			beego.Error(err)
			result["code"] = "1"
			result["msg"] = "查询失败: " + err.Error()
			this.Data["json"] = &result
			this.ServeJSON()
			this.StopRun()
		}
	} else {
		step.Owner = nextOwner
		if _, err := o.Update(&step, "Owner"); err != nil {
			beego.Error(err)
			result["code"] = "1"
			result["msg"] = "保存失败: " + err.Error()
			this.Data["json"] = &result
			this.ServeJSON()
			this.StopRun()
		}
	}

	apply := models.ReleaseApply{Id: applyId}
	o.Read(&apply)
	subject := apply.Subject

	go utils.SendMail(nextOwner, step.Description, "名称："+subject+"<br/>详情见：<a href='http://172.16.103.9:8866/release/index'>http://172.16.103.9:8866/release/index</a>", "html")

	this.Data["json"] = &result
	this.ServeJSON()
}

func (this *ReleaseController) Cancel() {
	ids := this.GetStrings("ids[]")
	result := map[string]string{}

	o := orm.NewOrm()
	for _, id := range ids {
		if i, err := strconv.Atoi(id); err == nil {
			m := models.ReleaseApply{Id: i}
			if err := o.Read(&m); err != nil {
				beego.Error(err)
				result["info"] = "获取申请信息错误: " + err.Error()
				this.Data["json"] = &result
				this.ServeJSON()
				this.StopRun()
			}
			if m.Status != 1 {
				if _, err := o.Delete(&models.ReleaseApply{Id: i}); err != nil {
					beego.Error(err)
					result["info"] = "操作失败: " + err.Error()
				} else {
					result["info"] = "操作成功!"

					var steps []models.ReleaseStep
					_, err = o.QueryTable(new(models.ReleaseStep)).Filter("apply_id", id).All(&steps)
					if err == nil {
						for s := range steps {
							o.Delete(&s)
						}
					}
				}
			} else {
				result["info"] = "无法取消!"
			}
		}
	}

	this.Data["json"] = &result
	this.ServeJSON()
}
