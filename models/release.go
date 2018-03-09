package models

import (
	"time"
)

type ReleaseApply struct {
	Id         int
	Subject    string
	ApplyType  int // 0: 新功能上线; 1: Bug修复; 2: 功能点优化
	Status     int // 0: 未完成; 1: 已完成
	Creator    string
	CreateTime time.Time      `orm:"auto_now_add;type(datetime)"`
	Step       []*ReleaseStep `orm:"reverse(many)"`
}

type ReleaseStep struct {
	StepId      string        `orm:"pk"`
	Apply       *ReleaseApply `orm:"rel(fk)"`
	Description string
	Content     string
	File        string
	Owner       string
	IsLast      int
	Status      string
	OperateTime string
}
