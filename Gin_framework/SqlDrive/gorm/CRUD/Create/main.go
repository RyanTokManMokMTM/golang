package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

const (
	UserName string = "root"
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

func main(){
	config := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",UserName,Password,Host,Port,DBName)
	db,err := gorm.Open(mysql.Open(config),&gorm.Config{
		//set the option here
		//CreateBatchSize: 5,// every insert will respect this option-> each time just insert 5 batches
	})
	//db.Session(&gorm.Session{CreateBatchSize: 5})

	if err != nil{
		fmt.Println(err)
		return
	}

	//migration model
	db.AutoMigrate(&user{})

	//createWithDefault(db)
	//createWithSpecificField(db)
	//createWithOmitField(db)
	//createWithBatchData(db)
	//createWithBatchDataB(db)
	//createWithMap(db)
	//createWithMapBatchData(db)
}

func (u *user)CreateRecord(name ,phone string){
	u.Name = name
	u.Phone = phone
	return
}

func createWithDefault(db *gorm.DB){
	record := &user{}
	record.CreateRecord("jackson","65828151")

	result := db.Create(record)

	if result.Error != nil{
		log.Println(result.Error)
	}
	log.Println(result.RowsAffected)
	log.Printf("insert data %v",record)
}

//select specific field to create the record
func createWithSpecificField(db *gorm.DB){
	//Insert INTO (some field) values (...)
	record := &user{}
	record.CreateRecord("Tom","66714585")

	//Insert INTO (`Name`) VALUES ("x","xxx")
	result := db.Select("Name").Create(record)
	if result.Error != nil {
		log.Println(result.Error)
	}

	log.Println(result.RowsAffected)
	log.Printf("created record %v",record)
}

//omit some field (no going to insert the value)
func createWithOmitField(db *gorm.DB){
	//Insert INTO (some field) values (...)
	record := &user{}
	record.CreateRecord("Jack.Chan","99621485")

	//ignoring the filed that passed in,other field will fill with the data
	//ignore "CreateAt","UpdatedAt","phone" fields
	result := db.Omit("CreateAt","UpdatedAt","phone").Create(&record)
	if result.Error != nil {
		log.Println(result.Error)
	}

	log.Println(result.RowsAffected)
	log.Printf("created record %v",record)
}

func createWithBatchData(db *gorm.DB){
	var users = []user{
		{Name: "TestA",Phone: "12345678"},
		{Name: "TestB",Phone: "56984216"},
		{Name: "TestC",Phone: "98765211"},
		{Name: "TestD",Phone: "39514568"},
		{Name: "TestE",Phone: "69513215"},
	}

	//using create method -> genera a single sql statement
	db.Create(&users)

	//check all data pk id
	for _,user := range users{
		fmt.Println(user.ID)
	}
}

func createWithBatchDataB(db *gorm.DB) {
	var users = []user{
		{Name: "Test1", Phone: "14512678"},
		{Name: "Test2", Phone: "99624216"},
		{Name: "Test3", Phone: "12345211"},
		{Name: "Test4", Phone: "35681268"},
		{Name: "Test5", Phone: "67663115"},
		{Name: "Test6", Phone: "55683115"},
		{Name: "Test7", Phone: "66153115"},
		{Name: "Test8", Phone: "69523115"},
		{Name: "Test9", Phone: "77853115"},
	}

	// can specify batch size to instead of inserting all data once
	//batchSize is 5 ,it will generate 2 sql statement to insert ,each time insert 5 records
	db.CreateInBatches(&users,5)
}

func createWithMap(db *gorm.DB){
	//gorm support map[string]interface
	db.Model(&user{}).Create(map[string]interface{}{
		"Name":"MapTestA",
		"Phone":"96512456",
	})
	log.Println("Map data Inserted succeed")
}

func createWithMapBatchData(db *gorm.DB){
	db.Model(&user{}).Create([]map[string]interface{}{
		{
			"Name":"MapBatchTestA",
			"Phone":"12345678",
		},
		{
			"Name":"MapBatchTestB",
			"Phone":"12345678",
		},
		{
			"Name":"MapBatchTestC",
			"Phone":"12345678",
		},
	})

	log.Println("map batches insert succeed")
}