package main

import (
	_ "bapi/routers"

	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/logs"
	// "github.com/astaxie/beego/orm"
	// _ "github.com/go-sql-driver/mysql"
)

func init() {
	// orm.RegisterDriver("mysql", orm.DRMySQL)
	// dbType := beego.AppConfig.String("dbType")
	// sqlConn := beego.AppConfig.String("sqlconn")
	// maxIdle, _ := beego.AppConfig.Int("maxIdle")
	// maxConn, _ := beego.AppConfig.Int("maxConn")
	// logs.Info("Connect to [%v] database with conn: [%v] \n", dbType, sqlConn)
	// err := orm.RegisterDataBase("default", dbType, sqlConn, maxIdle, maxConn)
	// if err != nil {
	// 	panic(err.Error())
	// }
}

func main() {

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
