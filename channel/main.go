package main

import (
	"channel/pool"
	"fmt"
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var counter int32
var waitG sync.WaitGroup

type DBConnection struct {
	id int32
}

func (D DBConnection) Close() error {
	fmt.Println("database closed , #" + fmt.Sprint(D.id))
	return nil
}

func Factory() (io.Closer,error){
	atomic.AddInt32(&counter,1)
	return DBConnection{id: counter},nil
}

func queryDB(query int,p *pool.Pool){
	defer waitG.Done()

	//get the resource
	resource, err := p.AcquireResource()
	if err != nil {
		log.Fatalln(err)
	}
	//release the resource
	defer p.ReleaseResource(resource)
	t := rand.Int()%10 + 1
	//sleep 1s and release
	time.Sleep(time.Duration(t) * time.Second)
	fmt.Println("finished query and returning back the resource "+fmt.Sprint(query))
}

func main(){
	p,err:=pool.New(Factory,5)
	if err != nil{
		log.Fatalln(err)
	}

	nums := 10
	waitG.Add(nums)
	for id := 0; id<nums;id++{
		go queryDB(id,p)
	}

	waitG.Wait()

	// after all query
	//close the pool
	p.Close()
	fmt.Println("pool model ended")
}
//func CreateTask() func(int){
//	return func(id int){
//		time.Sleep(time.Second) //sleep 1 second
//		fmt.Printf("tasks %d is completed\n",id)
//	}
//}
//
//func main(){
//	r:= runner.New(time.Second*3) //5second to run the task
//	r.AddTask(CreateTask(),CreateTask(),CreateTask())
//	err := r.Start()
//
//	//output the err
//	switch err {
//	case runner.ErrOSInterrupt:
//		fmt.Println("task is interrupted")
//	case runner.ErrTimeOut:
//		fmt.Println("tasks timeout")
//	default:
//		fmt.Print("all tasks completed")
//	}
//}
