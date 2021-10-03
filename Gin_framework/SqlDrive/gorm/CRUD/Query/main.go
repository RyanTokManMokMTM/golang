package main

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

const (
	UserName  string = "root"
	Password string = "admin"
	Host string = "127.0.0.1"
	Port int = 3306
	DBName string = "godb"
)

type user struct {
	gorm.Model
	Name string `gorm:"<-:false` //read only
	Phone string
}
//NOTE : Can only work with specify model
//if no pk is defined for relevant model -> will order by first field
//db.Model(model).First -> work =>all retrieving data will store in that model
//BUT
//res := map[string]interface{}
//db.Table("table").First(&res) -> won't work -> not specify any model -> don't know where to store the data

/*
Usage: matching record
db.First() -> get the first record, id
db.Last() -> get the last record, id
db.Take() -> get only 1 record, id
db.Find() -> get records from the range, id

Usage:condition
db.Where("name = ?",name).Find()
db.Where("name = ? AND age = ?",name,age).first()
db.Where("name like ?","%somestr%").find()
...

*/
func main(){
	config := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",UserName,Password,Host,Port,DBName)

	db , err := gorm.Open(mysql.Open(config),&gorm.Config{

	})
	if err != nil {
		log.Println(err)
		return
	}

	//retrieving data from our previous database
	//simple example
	//get the first record order by pk
	//
	//var myRecord user
	//select * from user ORDER by id LIMIT 1
	db.First(&myRecord)
	fmt.Println(myRecord)
	//
	//also ,can get one record,in any order(just one record)
	//select * from user LIMIT 1
	var myRecord2 user
	db.Take(&myRecord2)
	fmt.Println(myRecord2)

	//Get the last record order by pk
	//select * from user ORDER by id DESC LIMIT 1
	var myRecord3 user
	db.Last(&myRecord3)
	fmt.Println(myRecord3)

	//we can also check the error with errors function
	var myRecord4 user
	result := db.First(&myRecord4)
	fmt.Println(result.RowsAffected)
	fmt.Println(result.Error)
	errors.Is(result.Error,gorm.ErrRecordNotFound) // the error is not found?

	retrieving data from no pk table
	db.AutoMigrate(&Language{})
	db.Model(&Language{}).Create([]map[string]interface{}{
		{
			"Code":123,
			"Name":"go",
		},
		{
			"Code":622,
			"Name":"js",
		},
	})
	var lang Language
	//order by the first field
	//select * from `language` order by `language`.`code` LIMIT 1
	db.First(&lang)
	fmt.Println(lang)
	retrievingByStringID("9",db)
	retrievingByIntID(11,db)
	retrievingByIDGroup([]int{
		1,10,12,
	},db)
	getName("TestA",db)
	getRecordWithoutName("jackson",db)

	getRecordFromARange([]string{
		"Tom",
	},db)

	getRecordWithSimilarName("MapBatchTest%",db)
	getRecordWithMoreCondition("jackson","65828151",db)
	getRecordBetweenCondition(
		time.Date(2021,time.October,3,16,0,0,0,time.Local),
		time.Date(2021,time.October,3,16,15,0,0,time.Local),
		db,
		)
	getRecordWithTimeCondition(time.Date(2021,time.October,3,16,45,0,0,time.Local),db)

}

type Language struct{
	Code int
	Name string
}

func retrievingByStringID(id string,db *gorm.DB){
	var record user
	//select * from user where id = id
	db.First(&record,id)
	fmt.Println(record)
}

func retrievingByIntID(id int,db *gorm.DB){
	var record user
	//select * from user where id = id
	db.First(&record,id)
	fmt.Println(record)
}

func retrievingByIDGroup(ids []int,db *gorm.DB){
	var record user
	//select * from `user` where id IN (ids)
	db.Find(&record,ids)
	fmt.Println(record)
}

//Sql condition
//db.Where

//select * from user where name = ? LIMIT 1
func getName(searchName string,db *gorm.DB){
	var result user
	db.Where("Name = ?",searchName).First(&result)
	fmt.Println(result)
}

//select * from user where name <> ?
func getRecordWithoutName(filterName string, db *gorm.DB){
	var results []user
	db.Where("Name <> ?",filterName).Find(&results)
	fmt.Println(results)
}

//select * FROM user where name IN ?
func getRecordFromARange(rangeName []string,db *gorm.DB){
	var results []user
	db.Where("name IN ?",rangeName).Find(&results)
	fmt.Println(results)
}

//select * FROM user where name LIKE ?
func getRecordWithSimilarName(likeStr string,db *gorm.DB){
	var results []user
	db.Where("name LIKE ?",likeStr).Find(&results)
	fmt.Println(results)
}

//select * from user where ? AND ?
func getRecordWithMoreCondition(name ,phone string,db *gorm.DB){
	var results []user
	db.Where("name = ? AND phone = ?",name,phone).Find(&results)
	fmt.Println(results)
}

func getRecordBetweenCondition(in time.Time,out time.Time,db *gorm.DB){
	var results []user
	fmt.Println(in)
	fmt.Println(out)
	db.Where("created_at BETWEEN ? AND ?",in,out).Find(&results)
	fmt.Println(results)
	fmt.Println(len(results))
}

func getRecordWithTimeCondition(time time.Time ,db *gorm.DB){
	var results []user
	db.Where("updated_at > ?",time).Find(&results)
	fmt.Println(results)
}