package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

const (
	UserName  string = "root"
	Password  string = "admin"
	Host string = "127.0.0.1"
	Port int= 3306
	DB string = "goDB"
)

//example
type MyModel struct {
	gorm.Model
	UserName string
	Phone string
}

//field permission setting
type DBPremission struct {
	FirstName string `gorm:"<-:create"` //allow read and create(write)
	LastName string `gorm:"<-:update"` //allow read and update(write)
	Phone string `gorm:"<-"` //allow read and write(both create and update)[here no specific what one]
	IdNum string `gorm:"<-:false"` // write disable but read
	Addr string `gorm:"->;<-:create"` //allow read and create(write)
	NickName string `gorm:"->:false;<-create"` //create(write) only ,read is disabled
	IsFriend bool `gorm:"-"` //it will ignore by gorm. won't create it
}

//embed struct
type Writer struct {
	Name string
	Email string
}

type Book struct {
	ID int `gorm:"primary key"`
	Writer Writer `gorm:"embedded"`
}
/*type book struct {
	ID int `gorm:"primary key"`
	Name string
	Email string
}
*/

//with gorm before capital letter ,_ will be added -> callMe-> call_me
type BookWithEmbed struct {
	ID int `gorm:"primary key"`
	Writer Writer `gorm:"embedded;embeddedPrefix:writer_"`
}

/*type book struct {
	ID int `gorm:"primary key"`
	writerName string
	writerEmail string
}
*/


func main(){
	//strint the setting
	config := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",UserName,Password,Host,Port,DB)

	//using gorm to connect to mysql
	db, err := gorm.Open(mysql.Open(config),&gorm.Config{

	})
	if err != nil {
		log.Println(db)
		return
	}

	//creating the table using migration
	//with gorm need to follow the Protocol -> ID=>PK ,CreateTime,UpdateTime and DeleteTime
	//we can just embed gorm.Model to our struct(this model is already created the struct of the configuration)
	/*
	type Model struct {
	  ID        uint           `gorm:"primaryKey"`
	  CreatedAt time.Time
	  UpdatedAt time.Time
	  DeletedAt gorm.DeletedAt `gorm:"index"`
	}*/
	err = db.AutoMigrate(&MyModel{})
	if err != nil {
		log.Println(err)
		return
	} //create db in the mysql
	err = db.AutoMigrate(&DBPremission{})
	if err != nil {
		log.Println(err)
		return
	}

	err = db.AutoMigrate(&Book{})
	if err != nil {
		log.Println(err)
		return
	}

	err = db.AutoMigrate(&BookWithEmbed{})
	if err != nil {
		log.Println(err)
		return
	}
	//can add the permission to the field when doing CRUD
	//exa
}
