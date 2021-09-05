package main

import (
	//pb "rpc-example/route"
	"google.golang.org/grpc"
)

func main(){
	//connect grpc server setting
	//disable the secure certificate
	//enable blocking mode
	connection, err := grpc.Dial("localhost:8080",grpc.WithInsecure(),grpc.WithBlock())
	if err != nil {
		return
	}
	defer connection.Close() //close the connection

	//create a client stub base on current connection

}