package main

import (
	_ "beegotest/haruch/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase(
		"default",
		"postgres",
		"user=postgres password=postgres host=127.0.0.1 port=5432 dbname=postgres sslmode=disable")

	orm.RunSyncdb("default", false, true)
}

func main() {
	beego.Run()
}
