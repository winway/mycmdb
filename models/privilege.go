package models

import (
	"time"
)

type ApiWhiteList struct {
	Id         int
	Ip         string `orm:"unique;size(16)"`
	Token      string
	Applicant  string
	Comment    string
	CreateTime time.Time `orm:"auto_now;type(datetime)"`
}
