package dbconnect

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"tianqiserver/logger"
	"tianqiserver/config"
	"fmt"
)

var MysqlCon  *sql.DB


func InitMysql()  {

	defer func() {
		if p := recover(); p != nil {
			logger.Error.Fatalln("数据库连接失败:",p)
		}
	}()
	datasoure:=fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=utf8",config.Mysql.User,config.Mysql.Password,config.Mysql.Protocol,
		config.Mysql.Ip,config.Mysql.Port,config.Mysql.Database)
	MysqlCon,_=sql.Open("mysql",datasoure)
	MysqlCon.SetMaxOpenConns(256)
	MysqlCon.SetMaxIdleConns(30)

}