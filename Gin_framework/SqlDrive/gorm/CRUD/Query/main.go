package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	//db.First(&myRecord)
	//fmt.Println(myRecord)
	////
	////also ,can get one record,in any order(just one record)
	////select * from user LIMIT 1
	//var myRecord2 user
	//db.Take(&myRecord2)
	//fmt.Println(myRecord2)
	//
	////Get the last record order by pk
	////select * from user ORDER by id DESC LIMIT 1
	//var myRecord3 user
	//db.Last(&myRecord3)
	//fmt.Println(myRecord3)
	//
	////we can also check the error with errors function
	//var myRecord4 user
	//result := db.First(&myRecord4)
	//fmt.Println(result.RowsAffected)
	//fmt.Println(result.Error)
	//errors.Is(result.Error,gorm.ErrRecordNotFound) // the error is not found?
	//
	//retrieving data from no pk table
	//db.AutoMigrate(&Language{})
	//db.Model(&Language{}).Create([]map[string]interface{}{
	//	{
	//		"Code":123,
	//		"Name":"go",
	//	},
	//	{
	//		"Code":622,
	//		"Name":"js",
	//	},
	//})
	//var lang Language
	////order by the first field
	////select * from `language` order by `language`.`code` LIMIT 1
	//db.First(&lang)
	//fmt.Println(lang)
	//retrievingByStringID("9",db)
	//retrievingByIntID(11,db)
	//retrievingByIDGroup([]int{
	//	1,10,12,
	//},db)
	//getName("TestA",db)
	//getRecordWithoutName("jackson",db)
	//
	//getRecordFromARange([]string{
	//	"Tom",
	//},db)
	//
	//getRecordWithSimilarName("MapBatchTest%",db)
	//getRecordWithMoreCondition("jackson","65828151",db)
	//getRecordBetweenCondition(
	//	time.Date(2021,time.October,3,16,0,0,0,time.Local),
	//	time.Date(2021,time.October,3,16,15,0,0,time.Local),
	//	db,
	//	)
	//getRecordWithTimeCondition(time.Date(2021,time.October,3,16,45,0,0,time.Local),db)
	//getRecordBySliceId(db)
	//inlineCondition(db)
	//withNotCondition(db)
	//withORCondition(db)

	//withORCondition(db)()
	//withSelectFields(db)
	withOrderField(db);
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

//Sql query using map anc struct

//only work with non-zero value
func getRecordByStrut(db *gorm.DB){
	var userData user
	//select * from user where name =? ane phone = ? Limit 1
	db.Where(&user{Name:"jackson",Phone: "65828151"}).First(&userData)
	fmt.Println(userData)
}

//suppose non-zero function
func getRecordByMap(db *gorm.DB){
	var userData user
	//select * from user where name =? ane phone = ? Limit 1
	db.Where(map[string]interface{}{"Name":"TestA","Phone":"12345678"}).Find(&userData)
	fmt.Println(userData)
}

//query with slice of pk
func getRecordBySliceId(db *gorm.DB){
	var userDatas []user
	//select * FROM user where id IN (10,9,1)
	db.Where([]int64{10,9,1}).Find(&userDatas)
	fmt.Println(userDatas)
}

// inline condition
func inlineCondition(db *gorm.DB){
	var myUser user
	//get condition with name
	db.First(&myUser,"name = ?","jackson")
	fmt.Println("my user is ",myUser)

	var users []user
	db.Find(&users,"id IN ?" ,[]int64{1,2,3})
	fmt.Println("all users are",users)

	//more condition
	var otherUSer []user
	db.Find(&otherUSer,"Name <> ? AND Phone <> ? ","jackson","65828151")
	fmt.Println("other users :",otherUSer)
	//
	//with struct or map
	var structUser user
	db.First(&structUser,user{Name:"jackson"})
	fmt.Println("struct users :",structUser)

	var mapUser user
	db.First(&mapUser,map[string]interface{}{"phone":"65828151"})
	fmt.Println("map users :",mapUser)
}

//other than match the condition
//opposite of `where`
func withNotCondition(db *gorm.DB){
	//build not condition
	//like where
	var users user
	//select * FROM user WHERE NOT name = jackson
	db.Not("name = ?","jackson").First(&users)
	//fmt.Println("user name not jackson:",users)

	var users2 []user
	//Not IN a group
	//select * from user Where name NOT IN ("jackson","Tom")
	db.Not(map[string]interface{}{"name":[]string{
		"jackson",
		"Tom",
	}}).Find(&users2)
	//fmt.Println("user name not in tom and jackson",users2)

	var user3 []user
	//select * From user WHere Name <> jackson AND Phone <> 65828151
	db.Not(user{Name: "jackson",Phone: "65828151"}).Find(&user3)
	//fmt.Println("user name and age not jackson and 65828151",user3)

	var user4 user
	db.Not([]int32{1,2,3,4}).First(&user4)
	fmt.Println("user not in (1,2,3,4)",user4)
}

//with OR condition -> WHERE AND NOT ... method are all AND Condition
func withORCondition(db *gorm.DB){
	var userOR user
	//select * from user where name = jackson or name = tom LIMIT 1

	db.Where("Name = ?","jackson").Or("Name = ?" ,"Tom").Last(&userOR)
	fmt.Println("user not jackson or Tom:",userOR)

	var userStruct user
	//using map or struct
	db.Where(user{Name :"jackson"}).Or(user{Name: "Tom"}).Last(&userStruct)
	fmt.Println("user not jackson or Tom(struct):",userStruct)

	var userMap user
	db.Where(map[string]interface{}{"Name":"jackson"}).Or(map[string]interface{}{"Name":"Tom"}).Last(&userMap)
	fmt.Println("user not jackson or Tom(map):",userMap)
}

//Select specify Field with `SELECT`
func withSelectFields(db *gorm.DB){
	//select name field wit string
	var users user
	//select `name` from user where Name = "Tom"
	db.Select("Name").Where(user{Name: "Tom"}).First(&users)
	fmt.Println(users.CreatedAt)

	//select with string
	var user2 user
	//select name , create_at from user where not name = tom limit 1
	db.Select([]string{"Name","CreatedAt"}).Not("Name = ?","Tom").First(&user2)
	fmt.Println(user2.CreatedAt)

	//var user3 user
	////select from a table
	//rows, err := db.Table("users").Select("COALESCE(phone,?)", "65828151").Rows()
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//rows.Scan(user3)
	//fmt.Println(user3)
}

//With Specify order
func withOrderField(db *gorm.DB){
	//both are same condition
	//var users []user
	////select * FROM users ORDER BY name desc,phone desc
	//db.Order("name desc,phone desc").Find(&users)
	//fmt.Println(users)
	//
	////select * FROM users ORDER BY name desc,phone desc
	//var users2 []user
	//db.Order("name desc").Order("phone desc").Find(&users2)
	//fmt.Println(users2)

	//both are same condition

	var users3 []user
	//using clauses
	//select * from user
	db.Clauses(clause.OrderBy{
		Expression: clause.Expr{SQL:"FIELD(id,?)",Vars: []interface{}{[]int{1,2,3}},WithoutParentheses: true}}).Find(&users3)

	fmt.Printf("%v",users3)
}