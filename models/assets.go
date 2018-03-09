package models

import (
	"time"
)

type Idc struct {
	Id         int
	Name       string `orm:"unique" valid:"Required"`
	Address    string `valid:"Required"`
	Operator   string
	Linkman    string
	Phone      string `orm:"size(64)"`
	Comment    string
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	Ip         []*Ip     `orm:"reverse(many)"`
	Server     []*Server `orm:"reverse(many)"`
}

type Ip struct {
	Ip         string `orm:"pk;size(16)"`
	Idc        *Idc   `orm:"rel(fk)" valid:"Required"`
	Network    string `orm:"size(64)" valid:"Required;Match(/[0-9]+.[0-9]+.[0-9]+.[0-9]+/[0-9]+/)"`
	IpType     int8   `orm:"default(0)" valid:"Range(0, 10)"` // 0: 远程卡IP; 1: 业务网卡IP; 2: 数据网卡IP
	SubMask    string
	Gateway    string
	Dns        string
	Status     int8 `orm:"default(0)" valid:"Range(0, 10)"` // 0: 未使用; 1: 已使用
	Comment    string
	CreateTime time.Time `orm:"auto_now;type(datetime)"`
}

type Server struct {
	Sn                 string `orm:"pk" valid:"Required"`
	Idc                *Idc   `orm:"rel(fk)" valid:"Required"`
	CabinetNo          string `valid:"Required"`
	IdInsideCabinet    string `valid:"Required;Numeric"`
	RemoteCardMac      string `orm:"unique" valid:"Required"`
	RemoteCardIp       string `orm:"unique;size(16)"`
	HostName           string
	Password           string
	Eth1Ip             string `orm:"size(16)"`
	Eth2Ip             string `orm:"size(16)"`
	Eth3Ip             string `orm:"size(16)"`
	Eth4Ip             string `orm:"size(16)"`
	Eth1Mac            string
	Eth2Mac            string
	Eth3Mac            string
	Eth4Mac            string
	Nic                string
	NicDetail          string
	Cpu                string
	CpuDetail          string
	Disk               string
	DiskDetail         string
	DiskStructure      string
	Memory             string
	MemoryDetail       string
	PowerSupply        string
	PowerSupplyDetail  string
	Brand              string
	WarrantyTime       time.Time `orm:"null;type(date)"`
	RaidConfig         string    `orm:"size(1024)"`
	OsVersion          int8      `orm:"default(0)" valid:"Range(0, 10)"` // 0：CentOS-6.9-x86_64; 1：centos73-x86_64
	Status             int8      `orm:"default(0)" valid:"Range(0, 10)"` // 0：未安装OS; 1: 未使用; 2: 已使用
	Comment            string
	CreateTime         time.Time             `orm:"auto_now;type(datetime)"`
	OsInstallManifests []*OsInstallManifests `orm:"reverse(many)"`
}
