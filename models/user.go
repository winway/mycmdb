package models

import (
	"time"
)

type User struct {
	Email         string `orm:"pk;size(64)" valid:"Required"`
	Name          string
	Password      string
	Phone         string `orm:"size(64)"`
	IsAdmin       int    `orm:"default(0)"`
	LastLoginTime time.Time
	CreateTime    time.Time `orm:"auto_now_add;type(datetime)"`
}
