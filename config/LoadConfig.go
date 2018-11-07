package config

import (
	"gopkg.in/ini.v1"
)


type LevelDBConfig struct {
	RootPath string  `ini:"leveldb_rootpath"`
}

type MysqlDataSource struct {

	User string		`ini:"mysql_user"`
	Password string	`ini:"mysql_password"`
	Protocol string	`ini:"mysql_protocol"`
	Ip string		`ini:"mysql_ip"`
	Port string		`ini:"mysql_port"`
	Database string	`ini:"mysql_database"`
}

type RedisConfig struct {
	Ip string		`ini:"redis_ip"`
	Port string		`ini:"redis_port"`
}

type LogConfig struct {
	RootPath string	`ini:"log_rootpath"`
}

type HostConfig struct {
	IP string `ini:"host_ip"`
	Port string `ini:"host_port"`
}


var LevelDB LevelDBConfig
var Mysql   MysqlDataSource
var Redis   RedisConfig
var Log     LogConfig
var Host	HostConfig

func LoadConfigFromConfigINI(configfilepath string) {

	conf,err:=ini.Load(configfilepath)
	if err != nil {
		panic("读取配置文件失败")
	}

	conf.BlockMode=false
	err=conf.MapTo(&LevelDB)
	if err !=nil {
		panic("读取leveldb配置失败")
	}
	err=conf.MapTo(&Mysql)
	if err != nil {
		panic("读取mysql配置失败")
	}
	err=conf.MapTo(&Redis)
	if err != nil {
		panic("读取redis配置失败")
	}

	err=conf.MapTo(&Log)
	if err != nil {
		panic("读取log配置失败")
	}

	err=conf.MapTo(&Host)
	if err != nil {
		panic("读取host配置失败")
	}

}


