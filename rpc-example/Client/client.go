package main

import (
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"os"
	pb "rpc-example/route"
	"time"
)

//define some function to test the grpc function
func getSimpleFeature(client pb.RouteGuideClient){
	feature, err := client.GetFeature(context.Background(), &pb.Point{
		//send this point to service
		Latitude:  310235000,
		Longitude: 121437403,
	})
	if err != nil {
		return
	}

	//when we got the client ,print out the result
	fmt.Println(feature)
}

func serverStreamWithRectPoint(client pb.RouteGuideClient){
	serverStream, err := client.GetListOfFeature(context.Background(), &pb.Rectangle{
		Lo: &pb.Point{
			Latitude:  313374060,
			Longitude: 121358540,
		},
		Hi: &pb.Point{
			Latitude:  311034130,
			Longitude: 121598790,
		},
	})// calling the server stub and passing the 2 point,it will return a stream server
	if err != nil {
		log.Fatalln(err)
	}

	for {
		//received the data from stream
		feature,err := serverStream.Recv()
		//here are 2 err, a error = EOF mean is end of the streaming
		if err == io.EOF{
			//if end of file just beak
			break
		}
		if err != nil{
			log.Fatalln(err)
		}
		fmt.Println(feature)
	}
}

func simpleClientStream(client pb.RouteGuideClient){
	//create a temp point
	points := []*pb.Point{ //here are 3 points
		{Latitude: 313374060, Longitude: 121358540},
		{Latitude: 311034130, Longitude: 121598790},
		{Latitude: 310235000, Longitude: 121437403},
	}
	clientStream, err := client.GetRecordRoute(context.Background())
	if err != nil {
		return
	}

	for _ ,point := range points{
		//we are sending the point
		if err := clientStream.Send(point);err != nil{
			log.Fatalln(err)
		}
		time.Sleep(time.Second) // sleep for a second
	}
	//after sending the request,close the stream channel
	summary ,err := clientStream.CloseAndRecv() //here will close the server and receive the summary
	if err != nil{
		print("???")
		log.Fatalln(err)
	}
	fmt.Println(summary)
}

func readIntFromCmd(reader *bufio.Reader,target *int32){
	_,err := fmt.Fscanf(reader,"%d\n",target)
	if err!=nil{
		log.Fatalln("can not scan",err)
	}
}

func clientStreamAndServerStream(client pb.RouteGuideClient){
	//keep streaming
	clientStream, err := client.Recommend(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	go func(){ //using a routine to received
		for {
			feature ,err:= clientStream.Recv()
			if err == io.EOF{
				break
			}

			if err != nil{
				log.Fatalln(err)
			}
			fmt.Println(feature)
		}
	}()


	reader := bufio.NewReader(os.Stdin)

	for {
		//this loop will keep tracking user input
		req := pb.RecommendationRequest{Point: new(pb.Point)} //send this message
		var mode int32 //let user input the mode
		fmt.Println("Enter the recommendation mode(0 for far 1 for near)")
		readIntFromCmd(reader, &mode)

		fmt.Print("Enter Latitude: ")
		readIntFromCmd(reader, &req.Point.Latitude)

		fmt.Print("Enter Longitude: ")
		readIntFromCmd(reader, &req.Point.Longitude)

		req.Mode = pb.RecommendationMode(mode)
		//streaming to server
		if err := clientStream.Send(&req); err != nil{
			log.Fatalln(err)
		}

		//wait for 10ms
		time.Sleep(10 * time.Millisecond)
	}
}
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
	client := pb.NewRouteGuideClient(connection)
	//getSimpleFeature(client)
	//serverStreamWithRectPoint
	//simpleClientStream(client)
	//clientStreamAndServerStream(client)

}