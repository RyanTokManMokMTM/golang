package main

import (
	//"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"

	//"time"
)

const (
	UserName  string = "root"
	Password  string = "admin"
	Host string = "127.0.0.1"
	Port int= 3306
	DB string = "goDB"

)

type User struct {
	gorm.Model //include id and create update and delete
	Name string
	Code string
	salary uint
}

type Player struct {
	gorm.Model
	UserName string
	Adult bool
	PlayTime time.Duration
}

//simple using orm
func main(){
	//connect setting using mysql drive
	//add some param
	//handle time => parseTime=true
	config := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",UserName,Password,Host,Port,DB)

	//using gorm to apply the config
	db,err :=gorm.Open(mysql.Open(config),&gorm.Config{
		//can provide more configurations to mysql
	})

	//var precision = 2
	//db,err :=gorm.Open(mysql.New(mysql.Config{
	//	DSN:"",//connection string
	//	DefaultStringSize: 256, //set the default string size of db text field(not pk,index, default value)
	//	DisableDatetimePrecision: true,//disable datetime Precision(for some mysql version)
	//	DefaultDatetimePrecision: &precision, //default DatetimePrecision
	//	DontSupportRenameIndex: true,//Not allow to rename the index
	//	DontSupportRenameColumn: true, //Not allow to rename the column
	//	SkipInitializeWithVersion: false, //based on used version~
	//
	//}),&gorm.Config{
	//	//can provide more configurations to mysql
	//})
	if err != nil {
		log.Printf("err to init gorm and mysql with err :%v",err)
		return
	}
	log.Println("mysql is connected")

	//create scheme
	err = db.AutoMigrate(&User{}) //using Migrate to create the db table
	if err != nil{
		log.Println(err)
		return
	}


	err = db.AutoMigrate(&Player{}) //using Migrate to create the db table
	if err != nil{
		log.Println(err)
		return
	}

	//create data
	db.Create(&User{Name: "JacksonT",Code:"HD9",salary: 9851})
	db.Create(&User{Name: "JacksonA",Code:"HDas",salary: 651})
	db.Create(&User{Name: "JacksonB",Code:"A95",salary: 96551})
	db.Create(&User{Name: "JacksonC",Code:"HD9",salary: 9251})

	db.Create(&Player{UserName: "PlayA",Adult: true,PlayTime: time.Hour * 6})
	db.Create(&Player{UserName: "PlayB",Adult: false,PlayTime: time.Hour * 2})

	var userData User
	db.First(&userData,"code=?","HD9")
	fmt.Println(userData)

	var playerData Player
	db.First(&playerData,"1")
	fmt.Println(playerData)


	////creating the connection pool
	//sql,err := db.DB() //return a connection pool
	//if err != nil {
	//	log.Println("create connection pool failed",err)
	//	return
	//}
	//defer sql.Close() //close the connection to mysql
	////setting the connection pool
	//sql.SetMaxIdleConns(10) //set max number of idle connection
	//sql.SetMaxOpenConns(100) //set max number of connection to db
	//sql.SetConnMaxLifetime(time.Minute*30) //set the maximum amount of time a connection may be reused(after 30 mins a connection will be back to the pool(break the connection))
}
