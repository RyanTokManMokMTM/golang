package database

import (
	"fmt"
	"gopkg.in/ini.v1"
	"music_api_server/Tool"
)

var dbConfig *database

const (
	DbHost string = "127.0.0.1"
	DbPort int    = 3306
	DbUser  string = "root"
	DbPassword string = "admin"
	DbTable    string = "musicDB"
)

type database struct{
	Host string
	Port int
	User string
	Password string
	DbName string
	source *ini.File
}

func (db *database)Load(path string)  *database{
	exists, err := Tool.PathExists(path)
	if !exists {
		fmt.Println("ini file is not exist")
		return db
	}

	db.source, err = ini.Load(path)
	if err != nil{
		panic(err)
	}
	return db
}

func (db *database)Init()  *database{
	if db.source == nil{
		return db
	}

	db.Host = db.source.Section("database").Key("address").MustString("127.0.0.1")
	db.Port = db.source.Section("database").Key("port").MustInt(3306)
	db.User = db.source.Section("database").Key("user_service").MustString("root")
	db.Password = db.source.Section("database").Key("password").MustString("")
	db.DbName = db.source.Section("database").Key("database").MustString("music")

	return db

}