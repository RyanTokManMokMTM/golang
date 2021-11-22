package database

import (
	"fmt"
	"testing"
)

func TestDatabase(t *testing.T){
	dbConfig = (&database{}).Load("../config/server.ini").Init()
	fmt.Println(dbConfig)
	if dbConfig == nil{
		t.Fail()
	}
}