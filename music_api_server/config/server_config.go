package config

import (
	"fmt"
	"gopkg.in/ini.v1"
	"music_api_server/Tool"
)


var Server *server

type server struct {
	Address string
	Port int
	source *ini.File
}

func (ser *server)Load(path string) *server{
	exists, err := Tool.PathExists(path)
	if !exists{
		fmt.Println("NOT EXIST")
		return ser
	}
	ser.source, err = ini.Load(path)
	if err != nil {
		panic(err)
	}
	fmt.Println("Path Loaded!!!!")
	return ser
}

func (ser *server)Init() *server{
	if ser.source == nil{
		return ser
	}
	ser.Address = ser.source.Section("server").Key("address").MustString("0.0.0.0")
	ser.Port = ser.source.Section("server").Key("port").MustInt(8080)
	return ser
}
