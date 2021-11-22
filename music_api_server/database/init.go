package database

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"music_api_server/model/user_model"
)

var (
 	DB *gorm.DB
 	err error
 	dbSql *sql.DB
)

func init(){
	dbConfig = (&database{}).Load("config/server.ini").Init()

	config := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DbName)
	DB , err = gorm.Open(mysql.Open(config),&gorm.Config{

	})

	if err != nil{
		panic(err)
		return
	}

	if DB.Error != nil{
		fmt.Printf("database connection apiError %v",DB.Error)
		return
	}

	dbSql, err = DB.DB()
	if err != nil {
		fmt.Println(err)
		return
	}

	dbSql.SetMaxIdleConns(10)
	dbSql.SetMaxOpenConns(100)

	migration()
}

func migration(){
	DB.AutoMigrate(&user_model.UserInfo{})
}

func Close(){
	defer dbSql.Close()
}