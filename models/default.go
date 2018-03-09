package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.Debug = true

	/*
	   // For version 1.6
	   orm.DRMySQL
	   orm.DRSqlite
	   orm.DRPostgres

	   // < 1.6
	   orm.DR_MySQL
	   orm.DR_Sqlite
	   orm.DR_Postgres

	   mysql / sqlite3 / postgres 这三种是默认已经注册过的，所以可以无需设置
	*/
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// ORM 必须注册一个别名为 default 的数据库，作为默认使用
	// set default database
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("mysqluser")+":"+beego.AppConfig.String("mysqlpass")+
		"@tcp("+beego.AppConfig.String("mysqlurl")+")/"+beego.AppConfig.String("mysqldb")+"?charset=utf8&loc=Asia%2FShanghai", 30, 30)

	// 设置为 UTC 时间
	// orm.DefaultTimeLoc = time.UTC

	// 需要在init中注册定义的model
	// 如果使用 orm.QuerySeter 进行高级查询的话，这个是必须的。
	// 反之，如果只使用 Raw 查询和 map struct，是无需这一步的
	orm.RegisterModel(
		new(User),
		new(OperateLog),
		new(ApiWhiteList),
		new(Idc),
		new(Ip),
		new(Server),
		new(OsInstallManifests),
		new(ReleaseStep),
		new(ReleaseApply),
	)

	// create table
	orm.RunSyncdb("default", false, true)
}
