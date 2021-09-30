package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //just call init function that declared in this package
	"log"
	"time"
)

const (
	UserName  string = "root"
	Password  string = "admin"
	IP string = "127.0.0.1"
	Port int= 3306
	DB string = "goDB"
	MaxLife int = 5
	MaxConnection int = 10
	MaxIdleConnection int = 10

)

//----------------------------SQL Manipulation-------------------------------------
func createDB(db *sql.DB){
	createTable := `
	CREATE TABLE IF NOT EXISTS golangDB(
		id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
		name VARCHAR(20)
	);
`

	//exec the db string
	if _ , err := db.Exec(createTable); err != nil{
		fmt.Print(err)
		return
	}
	fmt.Println("create table succeed")
}

//insert
func insertFunc(db *sql.DB){
	result  , err := db.Exec("INSERT INTO `golangDB`(`name`) values(?)","jackson")
	if err != nil{
		fmt.Print(err)
		return
	}

	//get result data
	resultID ,err := result.LastInsertId() //get inserted id
	if err!=nil{
		fmt.Printf("insert data error : %v",err)
		return
	}
	fmt.Printf("insert data rows id :%d",resultID)

	//rowAffected -> check data record nums is same?
	rowAffected, err := result.RowsAffected() //which row is affected! and return that row num
	if err != nil{
		fmt.Printf("Get row affected failed ,err :%v",err)
		return
	}

	fmt.Println("Affected rows is ?",rowAffected)

}

//update data
func updateData(db *sql.DB){
	updateRes , err := db.Exec("UPDATE `golangDB` set `name`=? where id = ?","jacksonTest","1")
	if err != nil{
		fmt.Println(err)
		return
	}

	//get update result data
	updateId , err := updateRes.LastInsertId()
	if err != nil{
		fmt.Printf("update data failed,err: %v",err)
		return
	}
	fmt.Println("updated id ",updateId)

	//check row effect
	updateAffected,err := updateRes.RowsAffected()
	if err != nil{
		fmt.Printf("get updating rowsaffected err %v",err)
		return
	}
	fmt.Println("affected row " ,updateAffected)
}

//delete
func deleteData(db *sql.DB){
	deleteRes , err := db.Exec("DELETE FROM `golangDB` WHERE id=?","1")
	if err != nil{
		fmt.Printf("delete failed %v",err)
		return
	}

	fmt.Println("delete data succeed:",deleteRes)

	rowAffected,err := deleteRes.RowsAffected()
	if err != nil{
		fmt.Printf("delete rowAffected failed %v",err)
		return
	}
	fmt.Println("delete affected rows",rowAffected)
}
//-----------------------------------------------------------------------------------

//---------------------------------DB Query-----------------------------------------
type userData struct{
	Uid int
	Name string
}

type customers []userData

func queryCustomers(db *sql.DB){
	var user userData
	//get single data
	data := db.QueryRow("SELECT * FROM golangdb WHERE id =?","3")

	//scan the data with scan
	//store to our data model
	err := data.Scan(&user.Uid, &user.Name)
	if err != nil {
		log.Printf("db scanning err:%v",err)
		return
	}
	//defer db.Close()
	//we are now get all the data
	log.Println(user)
}

func queryMultiData(db *sql.DB){

	dbData, err := db.Query("SELECT * FROM golangdb") //this function must call db.close,otherwise won't release the connection
	if err != nil{
		fmt.Println(err)
		return
	}
	defer db.Close()
	//loop over the data
	var users customers
	for dbData.Next(){//loop over the rows
		var user userData
		err := dbData.Scan(&user.Uid, &user.Name)
		if err != nil {
			fmt.Printf("with a scanning error %v",err)
			return
		}

		fmt.Println(user)
		users = append(users,user) //add to our list

	}
	fmt.Println(users)
}


func sqlManipulate(db *sql.DB){
	createDB(db)
	insertFunc(db)
	updateData(db)
	deleteData(db)
}

func dbTransaction(db *sql.DB){
	//all sql Instruction with transaction mode either all succeed or all failed
	//if some instruction is failed to execute,then, current state will be recovered to original state
	//it won't commit and affect the
	//commit -> store all Transaction and finished
	//rollback -> failed and rollback to the non-transacted state/save point
	//while Transaction , can mark some save point and rollback to specific point if it's needed

	//try it
	print("???")
	begin, err := db.Begin() //starting transaction
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 100;i<200;i++ {
		exec, err := begin.Exec("INSERT INTO golangdb(Name) values(?)", i)
		if err != nil {
			fmt.Println(" transaction err:", err)
			//rollback to original state
			begin.Rollback()
			return
		}

		//no error is happened
		//check row affect
		affected, err := exec.RowsAffected()
		if err != nil || affected != 1 { //not just affected 1 row
			fmt.Printf("affect row err %v", err)
			begin.Rollback()
			return
		}
	}
	begin.Commit() //finished! then commit

}
func main(){
	//setting sql configuration
	//using a string to configure the db ：format： userName:Password@[protocol](addr)/databaseName
	config := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",UserName,Password,IP,Port,DB)
	fmt.Println(config)
	conn , err := sql.Open("mysql",config)
	if err != nil{
		panic(err)
	}

	log.Println("db connected!")

	//set our driver to be closed the db driver before connection is closed by sql ,OS etc
	//recommend less than 5 minute
	conn.SetConnMaxLifetime(time.Duration(MaxLife)*time.Minute)

	//used to limit the connection by application
	//only 10 connection can be connected to drive(using mux lock)
	conn.SetMaxOpenConns(MaxConnection)

	//set the value same as MaxOpenConnection
	//MaxIdle is smaller than MaxOpenConnection can be opened and closed much more frequently
	//IdleConnection can be closed by MaxLifeTime
	conn.SetMaxIdleConns(MaxIdleConnection) //is used to close the idle connection rapidly

	log.Println("Configured is Done!")

	//
	sqlManipulate(conn)
	queryCustomers(conn)
	queryMultiData(conn)
	dbTransaction(conn)

}
