//simple example using contenxt
/*
	//create a root context
	//ctx := context.Background() //a root(empty) context,can be cancel
	//ctx2 := context.TODO() //use it ,when not know which context to use

	//create sub context bash on root context
	//ctx, cancel  := context.WithCancel(ctx) //return a ctx ,and cancel function
	//ctx , cancel := context.WithTimeout(ctx,time.Duration(2)) //2s only -> cancel
	//ctx , cancel := context.WithDeadline(ctx,time) //a specific time -> cancel
*/
package main

import (
	"context"
	"fmt"
	"time"
)

func contextToDO(ctx context.Context){
	//suppose our work will work 5s long
	//context has a Done: a readonly channel, is done or not
	select {
	case <- time.After(6 * time.Second):
		fmt.Println("word is done!") //word is done after the second
	case <-ctx.Done():
		err := ctx.Err()
		if err != nil {
			fmt.Println(err.Error()) //print out the error
		}
	}
}

func main(){
	//create a root context
	ctx := context.Background()

	//create cancelable context base on root context
	//we handle cancel oursefl
	subCtx , cancel := context.WithCancel(ctx)

	//we need to call cancel function ourself
	//simulate the client with a goroutine
	go func(){
		//is main exceed or cancel first?
		time.Sleep(time.Second * 5) //wait for the time
		cancel()
	}()

	contextToDO(subCtx)
}
