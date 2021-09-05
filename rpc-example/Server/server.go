package main

import (
	"google.golang.org/grpc"
	"log"
	"net"

	"context"

	pb "rpc-example/route"
)

//define our own struct type to implement RouteGuideServer
type RouteGuideServer struct {
	pb.UnimplementedRouteGuideServer //we need to embed this struct because of the protoc used UnimplementedRouteGuideServer to implement something
}

//implement the interface method

func (r *route) GetFeature(ctx context.Context, p *pb.Point) (*Feature, error){
	return nil,nil
}

func (r *route) GetListOfFeature(rectangle *pb.Rectangle, featureSer pb.RouteGuide_GetListOfFeatureServer) error{
	return nil
}

func (r *route) GetRecordRoute(recordSer pb.RouteGuide_GetRecordRouteServer) error{
	return nil
}
func (r *route) Recommend(recSer pb.RouteGuide_RecommendServer) error{
	return nil
}


//define a function to create our service route
func newServer() *RouteGuideServer{
	return &RouteGuideServer{}
}

func main(){
	//servere side
	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}

	//create a new grpc server
	grpcSer := grpc.NewServer()

	//here we need to register our service to grpc server
	pb.RegisterRouteGuideServer(grpcSer,newServer())
}