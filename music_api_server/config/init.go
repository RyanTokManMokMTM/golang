package config

import "fmt"

func init(){
	Server = (&server{}).Load("config/server.ini").Init()
	fmt.Println(Server)
}
