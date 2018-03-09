package models

import (
	"time"
)

type OsInstallManifests struct {
	Id         int
	Server     *Server `orm:"rel(fk)"`
	Status     int8    `orm:"default(1)" valid:"Range(0, 10)"` // 0：finishInstall; 1: readyForRaid; 2: readyForSys
	Comment    string
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
}
