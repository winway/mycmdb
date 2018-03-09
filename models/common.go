package models

import (
	"time"
)

type OperateLog struct {
	Id          int
	User        string
	Action      string
	Url         string
	Model       string
	PrimaryKey  string
	Detail      string
	OperateTime time.Time `orm:"auto_now;type(datetime)"`
}
