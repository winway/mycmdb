package main

import (
	"github.com/astaxie/beego"
	_ "mycmdb/routers"
)

func main() {
	beego.Run()
}
