package models

import (
	"database/sql"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type DBInit struct {
	Db *sql.DB
}

func DBInitNew() *DBInit {
	db, _ := orm.GetDB()
	dbInit := DBInit{Db: db}
	return &dbInit
}

func (dbInit *DBInit) InitDatabase() {
	if dbInit.Db == nil {
		beego.LoadAppConfig("ini", "..\\conf\\app.conf")
		orm.RegisterDriver("mysql", orm.DRMySQL)
		dbType := beego.AppConfig.String("dbType")
		sqlConn := beego.AppConfig.String("sqlconn")
		maxIdle, _ := beego.AppConfig.Int("maxIdle")
		maxConn, _ := beego.AppConfig.Int("maxConn")
		logs.Info("Connect to [%v] database with conn: [%v] \n", dbType, sqlConn)
		orm.RegisterDataBase("default", dbType, sqlConn, maxIdle, maxConn)
		orm.Debug = true
	}
}
