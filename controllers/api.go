package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"mycmdb/models"
	"mycmdb/utils"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

type SafeApiWhiteList struct {
	sync.Mutex
	M map[string]string
}

var ApiWhiteList SafeApiWhiteList

type ApiController struct {
	beego.Controller
}

func init() {
	go func() {
		o := orm.NewOrm()
		for {
			list := []orm.Params{}
			if _, err := o.QueryTable(new(models.ApiWhiteList)).Distinct().Values(&list, "Id", "Ip", "Token"); err != nil {
				beego.Error(err)
				time.Sleep(10 * time.Minute)
				continue
			}

			m := make(map[string]string)

			for _, item := range list {
				m[item["Ip"].(string)] = item["Token"].(string)
			}

			ApiWhiteList.Lock()
			ApiWhiteList.M = m
			ApiWhiteList.Unlock()

			time.Sleep(10 * time.Minute)
		}
	}()
}

func (this *ApiController) Prepare() {
	ip := this.Ctx.Input.IP()
	token := this.Ctx.Request.Header.Get("X-Mycmdb-Auth-Token")

	beego.Error(ip, token)

	if ApiWhiteList.M[ip] != token {
		beego.Error("Api unauthorized")
		result := map[string]interface{}{"code": "1", "msg": "Api unauthorized"}
		this.Data["json"] = &result
		this.ServeJSON()
	}
}

type UpdateStatus struct {
	Status int8
}

type machineInfo struct {
	Sn     string       `json:"machine_sn"`
	Disk   diskInfoArr  `json:"disk"`
	Memory MemInfo      `json:"memory"`
	Cpu    cpuInfoArr   `json:"cpu"`
	Power  powerInfoArr `json:"power"`
	Nic    nicArr       `json:"nic"`
}

type diskInfoArr struct {
	DiskNum int        `json:"disk_num"`
	Data    []diskInfo `json:"data"`
}

type diskInfo struct {
	Sn       string `json:"sn"`
	Locate   string `json:"locate"`
	Status   string `json:"status"`
	Size     string `json:"size"`
	DiskType string `json:"disk_type"`
}

type MemInfo struct {
	MemSize  string `json:"memory_size"`
	MemType  string `json:"memory_type"`
	MemSpeed string `json:"memory_speed"`
}

type cpuInfoArr struct {
	CpuNum int       `json:"cpu_num"`
	Data   []cpuInfo `json:"data"`
}

type cpuInfo struct {
	CpuProcess string `json:"cpu_process"`
	CpuType    string `json:"cpu_type"`
}

type powerInfoArr struct {
	PowerNum int         `json:"power_num"`
	Data     []powerInfo `json:"data"`
}

type powerInfo struct {
	PowerInput  string `json:power_name`
	PowerOutput string `json:power_output`
	PowerStatus string `json:"power_status"`
}

type nicArr struct {
	NicNum int       `json:"nic_num"`
	Data   []nicInfo `json:"data"`
}

type nicInfo struct {
	NicName string `json:"nic_name"`
	NicMac  string `json:"nic_mac_address"`
}

/*
获取还没有录入基本硬件信息（如disk信息）的机器的运程卡IP
*/
func (this *ApiController) GetBareRemoteCardIp() {
	result := map[string]interface{}{"code": "0", "msg": "ok"}

	limit, err := this.GetInt("limit")
	if err != nil || limit < 0 {
		limit = 1
	}

	o := orm.NewOrm()
	m := []orm.Params{}
	qs := o.QueryTable(new(models.Server))
	cond := orm.NewCondition()
	conddisk := cond.Or("Disk", "").Or("DiskDetail", "")
	if _, err := qs.SetCond(cond.AndCond(conddisk)).Exclude("RemoteCardIp", "").Limit(limit).Values(&m, "Sn", "RemoteCardMac", "RemoteCardIp"); err != nil {
		beego.Error(err)
		result["code"] = "1"
		result["msg"] = err.Error()
		this.Data["json"] = &result
	} else {
		result["data"] = m
	}

	this.Data["json"] = &result
	this.ServeJSON()
}

/*
更新待安装机器的硬件信息，包括cpu、disk、memory、nic、电源，RemoteCardIp是GET /api/osinstall/bare返回的
*/
func (this *ApiController) UpdateHWInfo() {
	result := map[string]interface{}{"code": "0", "msg": "ok"}

	var ob machineInfo
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &ob); err != nil {
		beego.Error(err)
		result["code"] = "1"
		result["msg"] = err.Error()
		this.Data["json"] = &result
		this.ServeJSON()
		this.StopRun()
	}

	o := orm.NewOrm()
	server := models.Server{Sn: ob.Sn}
	if err := o.Read(&server); err != nil {
		beego.Error(err)
		result["code"] = "1"
		result["msg"] = err.Error()
		this.Data["json"] = &result
		this.ServeJSON()
		this.StopRun()
	}

	if ob.Disk.DiskNum > 0 {
		server.Disk, _ = DiskInfoExtract(ob.Disk, "summary")
		server.DiskDetail, _ = DiskInfoExtract(ob.Disk, "detail")
		server.DiskStructure, _ = DiskInfoExtract(ob.Disk, "structure")
	}

	beego.Error(server.Disk)
	beego.Error(server.DiskDetail)
	beego.Error(server.DiskStructure)

	server.Memory, _ = MemInfoExtract(ob.Memory, "summary")
	server.MemoryDetail, _ = MemInfoExtract(ob.Memory, "detail")

	beego.Error(server.Memory)
	beego.Error(server.MemoryDetail)

	if ob.Cpu.CpuNum > 0 {
		server.Cpu, _ = CpuInfoExtract(ob.Cpu, "summary")
		server.CpuDetail, _ = CpuInfoExtract(ob.Cpu, "detail")
	}

	beego.Error(server.Cpu)
	beego.Error(server.CpuDetail)

	if ob.Power.PowerNum > 0 {
		server.PowerSupply, _ = PowerInfoExtract(ob.Power, "summary")
		server.PowerSupplyDetail, _ = PowerInfoExtract(ob.Power, "detail")
	}

	beego.Error(server.PowerSupply)
	beego.Error(server.PowerSupplyDetail)

	if ob.Nic.NicNum > 0 {
		server.Nic, _ = NicInfoExtract(ob.Nic, "summary")
		server.NicDetail, _ = NicInfoExtract(ob.Nic, "detail")
		server.Eth1Mac, server.Eth2Mac, server.Eth3Mac, server.Eth4Mac, _ = NicMacInfoExtract(ob.Nic)

		if server.Eth1Ip == "" {
			ips, err := utils.GetAvailIp(server.Idc.Id, 1, 1)
			if err != nil {
				beego.Error(err)
				result["code"] = "1"
				result["msg"] = "自动分配网卡IP失败: " + err.Error()
				this.Data["json"] = &result
				this.ServeJSON()
				this.StopRun()
			}

			server.Eth1Ip = ips[0]

			beego.Error(server.Eth1Ip)
		}

		if server.Eth2Ip == "" {
			ips, err := utils.GetAvailIp(server.Idc.Id, 2, 1)
			if err != nil {
				beego.Error(err)
				result["code"] = "1"
				result["msg"] = "自动分配网卡IP失败: " + err.Error()
				this.Data["json"] = &result
				this.ServeJSON()
				this.StopRun()
			}

			server.Eth2Ip = ips[0]

			beego.Error(server.Eth2Ip)
		}
	}

	beego.Error(server.Nic)
	beego.Error(server.NicDetail)
	beego.Error(server.Eth1Mac, server.Eth2Mac, server.Eth3Mac, server.Eth4Mac)

	if _, err := o.Update(&server); err != nil {
		beego.Error(err)
		result["code"] = "1"
		result["msg"] = err.Error()
	}

	this.Data["json"] = &result
	this.ServeJSON()
}

/*
获取待安装信息，包括Id、序列号、机房Id、远程卡IP、OS版本等
*/
func (this *ApiController) GetManifests() {
	result := map[string]interface{}{"code": "0", "msg": "ok"}

	limit, err := this.GetInt("limit")
	if err != nil || limit < 0 {
		limit = 1
	}

	o := orm.NewOrm()
	qs := o.QueryTable(new(models.OsInstallManifests)).
		RelatedSel("Server").
		Filter("Server__RemoteCardIp__ne", "").
		Filter("Server__Disk__ne", "").
		Filter("Server__DiskDetail__ne", "").
		Filter("Server__RaidConfig__ne", "").
		Filter("Status", 1).
		Limit(limit)

	var maps []*models.OsInstallManifests
	_, err = qs.All(&maps)
	if err != nil {
		beego.Error(err)
		result["code"] = "1"
		result["msg"] = err.Error()
		this.Data["json"] = &result
		this.ServeJSON()
		this.StopRun()
	}

	var manifests []map[string]interface{}
	for _, i := range maps {
		raidConfig := []map[string]interface{}{}

		beego.Error(err, json.Unmarshal([]byte(i.Server.RaidConfig), &raidConfig))

		beego.Error(raidConfig)

		ip1 := models.Ip{Ip: i.Server.Eth1Ip}
		o.Read(&ip1)
		ip2 := models.Ip{Ip: i.Server.Eth2Ip}
		o.Read(&ip2)

		manifest := map[string]interface{}{"Id": i.Id,
			"Sn":            i.Server.Sn,
			"Idc":           i.Server.Idc.Id,
			"RemoteCardMac": i.Server.RemoteCardMac,
			"RemoteCardIp":  i.Server.RemoteCardIp,
			"Eth1Mac":       i.Server.Eth1Mac,
			"Eth1Ip": map[string]string{
				"Ip":      ip1.Ip,
				"SubMask": ip1.SubMask,
				"Gateway": ip1.Gateway,
				"Dns":     ip1.Dns,
			},
			"Eth2Mac": i.Server.Eth2Mac,
			"Eth2Ip": map[string]string{
				"Ip":      ip2.Ip,
				"SubMask": ip2.SubMask,
				"Gateway": ip2.Gateway,
				"Dns":     ip2.Dns,
			},
			"RaidConfig": raidConfig,
			"OsVersion":  i.Server.OsVersion}
		manifests = append(manifests, manifest)
	}

	result["data"] = manifests

	this.Data["json"] = &result
	this.ServeJSON()
}

/*
更新操作系统安装进度
*/
func (this *ApiController) UpdateStatus() {
	result := map[string]interface{}{"code": "0", "msg": "ok"}

	id, err := this.GetInt(":id")
	if err != nil {
		beego.Error(err)
		result["code"] = "1"
		result["msg"] = err.Error()
		this.Data["json"] = &result
		this.ServeJSON()
		this.StopRun()
	}

	var ob UpdateStatus
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &ob); err != nil {
		beego.Error(err)
		result["code"] = "1"
		result["msg"] = err.Error()
		this.Data["json"] = &result
		this.ServeJSON()
		this.StopRun()
	} else if ob.Status == 0 {
		result["code"] = "1"
		result["msg"] = "args incomplete"
		this.Data["json"] = &result
		this.ServeJSON()
		this.StopRun()
	}

	o := orm.NewOrm()
	m := models.OsInstallManifests{Id: id}
	if err := o.Read(&m); err != nil {
		beego.Error(err)
		result["code"] = "1"
		result["msg"] = err.Error()
		this.Data["json"] = &result
	} else {
		m.Status = ob.Status
		if _, err := o.Update(&m, "Status"); err != nil {
			beego.Error(err)
			result["code"] = "1"
			result["msg"] = err.Error()
		}
	}

	this.Data["json"] = &result
	this.ServeJSON()
}

func DiskInfoExtract(data diskInfoArr, operate string) (string, error) {
	if operate == "summary" {
		rs := map[string]int{}
		for _, i := range data.Data {
			rs[strings.TrimSpace(i.Size)+"*"+strings.TrimSpace(i.DiskType)] += 1
		}

		var b bytes.Buffer
		var sep string
		for k, v := range rs {
			b.WriteString(sep + k + "*" + strconv.Itoa(v))
			if sep == "" {
				sep = ";"
			}
		}

		return b.String(), nil
	} else if operate == "detail" {
		rs, err := json.Marshal(data)
		if err != nil {
			beego.Error(err)
			return "", nil
		}
		return string(rs), nil
	} else if operate == "structure" {
		rs := map[int]string{}
		r := regexp.MustCompile(`[a-zA-z ]+(\d+:\d+:(\d+))`)
		for _, i := range data.Data {
			k := r.FindStringSubmatch(i.Locate)
			j, _ := strconv.Atoi(k[2])
			rs[j] = k[1] + "-" + strings.TrimSpace(i.Size) + "*" + strings.TrimSpace(i.DiskType)
		}

		rss, err := json.Marshal(rs)
		if err != nil {
			beego.Error(err)
			return "", nil
		}
		return string(rss), nil
	}

	return "", errors.New("Unknown operate")
}

func MemInfoExtract(data MemInfo, operate string) (string, error) {
	if operate == "summary" {
		return data.MemSize, nil
	} else if operate == "detail" {
		rs, err := json.Marshal(data)
		if err != nil {
			beego.Error(err)
			return "", nil
		}
		return string(rs), nil
	}

	return "", errors.New("Unknown operate")
}

func CpuInfoExtract(data cpuInfoArr, operate string) (string, error) {
	if operate == "summary" {
		rs := map[string]int{}
		for _, i := range data.Data {
			rs[strings.TrimSpace(i.CpuProcess)] += 1
		}

		var b bytes.Buffer
		var sep string
		for k, v := range rs {
			b.WriteString(sep + k + "*" + strconv.Itoa(v))
			if sep == "" {
				sep = ";"
			}
		}

		return b.String(), nil
	} else if operate == "detail" {
		rs, err := json.Marshal(data)
		if err != nil {
			beego.Error(err)
			return "", nil
		}
		return string(rs), nil
	}

	return "", errors.New("Unknown operate")
}

func PowerInfoExtract(data powerInfoArr, operate string) (string, error) {
	if operate == "summary" {
		return strconv.Itoa(data.PowerNum), nil
	} else if operate == "detail" {
		rs, err := json.Marshal(data)
		if err != nil {
			beego.Error(err)
			return "", nil
		}
		return string(rs), nil
	}

	return "", errors.New("Unknown operate")
}

func NicInfoExtract(data nicArr, operate string) (string, error) {
	if operate == "summary" {
		return strconv.Itoa(data.NicNum), nil
	} else if operate == "detail" {
		rs, err := json.Marshal(data)
		if err != nil {
			beego.Error(err)
			return "", nil
		}
		return string(rs), nil
	}

	return "", errors.New("Unknown operate")
}

func NicMacInfoExtract(data nicArr) (string, string, string, string, error) {
	mac := map[string]string{"1": "", "2": "", "3": "", "4": ""}
	var r = regexp.MustCompile(`NIC.Integrated.1-(\d)-1`)
	for _, i := range data.Data {
		k := r.FindStringSubmatch(i.NicName)
		mac[k[1]] = i.NicMac
	}

	return mac["1"], mac["2"], mac["3"], mac["4"], nil
}

// /*
// 获取还没有配置运程卡IP的远程卡MAC
// */
// func (this *ApiController) GetRemoteCardMac() {
// 	result := map[string]interface{}{"code": "0", "msg": "ok"}

// 	limit, err := this.GetInt("limit")
// 	if err != nil || limit < 0 {
// 		limit = 1
// 	}

// 	o := orm.NewOrm()
// 	m := []orm.Params{}
// 	if _, err := o.QueryTable(new(models.Server)).Filter("RemoteCardIp", "").Limit(limit).Values(&m, "RemoteCardMac"); err != nil {
// 		beego.Error(err)
// 		result["code"] = "1"
// 		result["msg"] = err.Error()
// 		this.Data["json"] = &result
// 	} else {
// 		result["data"] = m
// 	}

// 	this.Data["json"] = &result
// 	this.ServeJSON()
// }

// /*
// 回写远程卡Mac对应的远程卡IP，只有配置了远程卡IP，才能进行操作系统安装
// */
// func (this *ApiController) ReportRemoteCardIp() {
// 	result := map[string]interface{}{"code": "0", "msg": "ok"}

// 	var ob RemoteCard
// 	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &ob); err != nil {
// 		beego.Error(err)
// 		result["code"] = "1"
// 		result["msg"] = err.Error()
// 		this.Data["json"] = &result
// 		this.ServeJSON()
// 		this.StopRun()
// 	}

// 	if ob.RemoteCardMac == "" || ob.RemoteCardIp == "" {
// 		result["code"] = "1"
// 		result["msg"] = "args incomplete"
// 		this.Data["json"] = &result
// 		this.ServeJSON()
// 		this.StopRun()
// 	}

// 	o := orm.NewOrm()
// 	server := models.Server{}
// 	if err := o.QueryTable(new(models.Server)).Filter("RemoteCardMac", ob.RemoteCardMac).One(&server); err != nil {
// 		beego.Error(err)
// 		result["code"] = "1"
// 		result["msg"] = err.Error()
// 		this.Data["json"] = &result
// 		this.ServeJSON()
// 		this.StopRun()
// 	}

// 	server.RemoteCardIp = ob.RemoteCardIp
// 	if _, err := o.Update(&server, "RemoteCardIp"); err != nil {
// 		beego.Error(err)
// 		result["code"] = "1"
// 		result["msg"] = err.Error()
// 	}

// 	this.Data["json"] = &result
// 	this.ServeJSON()
// }
//
// /*
// 获取可用远程卡、网卡IP
// */
// func (this *ApiController) GetAvailIp() {
// 	result := map[string]interface{}{"code": "0", "msg": "ok"}

// 	idc, err := this.GetInt("idc")
// 	if err != nil {
// 		beego.Error(err)
// 		result["code"] = "1"
// 		result["msg"] = err.Error()
// 		this.Data["json"] = &result
// 		this.ServeJSON()
// 		this.StopRun()
// 	}

// 	ipType, err := this.GetInt("iptype")
// 	if err != nil {
// 		beego.Error(err)
// 		result["code"] = "1"
// 		result["msg"] = err.Error()
// 		this.Data["json"] = &result
// 		this.ServeJSON()
// 		this.StopRun()
// 	}

// 	num, err := this.GetInt("num")
// 	if err != nil || num < 0 {
// 		num = 1
// 	}

// 	utils.AvailIpLock.Lock()

// 	o := orm.NewOrm()
// 	qs := o.QueryTable(new(models.Ip)).Filter("IpType", ipType).Filter("Status__ne", 1).RelatedSel("Idc").Filter("Idc__Id", idc).Offset(0).Limit(num)

// 	var maps []*models.Ip
// 	if n, err := qs.All(&maps, "Ip"); err != nil {
// 		beego.Error(err)
// 		result["code"] = "1"
// 		result["msg"] = err.Error()
// 		this.Data["json"] = &result
// 		this.ServeJSON()
// 		this.StopRun()
// 	} else if int(n) < num {
// 		result["code"] = "1"
// 		result["msg"] = "Not Enough IP"
// 		this.Data["json"] = &result
// 		this.ServeJSON()
// 		this.StopRun()
// 	}

// 	o.Begin()
// 	for _, i := range maps {
// 		i.Status = 1
// 		if _, err := o.Update(i, "Status"); err != nil {
// 			beego.Error(err)
// 			o.Rollback()
// 			result["code"] = "1"
// 			result["msg"] = err.Error()
// 			this.Data["json"] = &result
// 			this.ServeJSON()
// 			this.StopRun()
// 		}
// 	}
// 	o.Commit()

// 	utils.AvailIpLock.Unlock()

// 	ips := []string{}
// 	for _, v := range maps {
// 		ips = append(ips, v.Ip)
// 	}

// 	result["data"] = map[string]interface{}{"Num": num, "ipType": ipType, "Ip": ips, "Idc": idc}

// 	this.Data["json"] = &result
// 	this.ServeJSON()
// }
