package main

import (
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"io"
	"log"
	"math"
	"net"
	"time"

	"context"

	pb "rpc-example/route"
)

//define our own struct type to implement RouteGuideServer
type RouteGuideServer struct {
	feature []*pb.Feature
	pb.UnimplementedRouteGuideServer //we need to embed this struct because of the protoc used UnimplementedRouteGuideServer to implement something
}

//useful function
//define a function to get a point inside 2 point(rectangle)
func inRange(point *pb.Point, rect *pb.Rectangle) bool {
	left := math.Min(float64(rect.Lo.Longitude), float64(rect.Hi.Longitude))
	right := math.Max(float64(rect.Lo.Longitude), float64(rect.Hi.Longitude))
	top := math.Max(float64(rect.Lo.Latitude), float64(rect.Hi.Latitude))
	bottom := math.Min(float64(rect.Lo.Latitude), float64(rect.Hi.Latitude))

	if float64(point.Longitude) >= left &&
		float64(point.Longitude) <= right &&
		float64(point.Latitude) >= bottom &&
		float64(point.Latitude) <= top {
		return true
	}
	return false
}

//change num to r
func toRadians(num float64) float64 {
	return num * math.Pi / float64(180)
}

//calculate the distance between 2 point
func calcDistance(p1 *pb.Point, p2 *pb.Point) int32 {
	const CordFactor float64 = 1e7
	const R = float64(6371000) // earth radius in metres
	lat1 := toRadians(float64(p1.Latitude) / CordFactor)
	lat2 := toRadians(float64(p2.Latitude) / CordFactor)
	lng1 := toRadians(float64(p1.Longitude) / CordFactor)
	lng2 := toRadians(float64(p2.Longitude) / CordFactor)
	dlat := lat2 - lat1
	dlng := lng2 - lng1

	a := math.Sin(dlat/2)*math.Sin(dlat/2) +
		math.Cos(lat1)*math.Cos(lat2)*
			math.Sin(dlng/2)*math.Sin(dlng/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance := R * c
	return int32(distance)
}

//return a point that near of far the point
func (r * RouteGuideServer) recommendPoint(req *pb.RecommendationRequest) (*pb.Feature,error){
	var nearest, farthest *pb.Feature
	var nDistnace,fDistance int32

	for _,feature := range r.feature{
		distance := calcDistance(feature.Location,req.Point)
		if nearest == nil || distance < nDistnace {
			nDistnace = distance
			nearest = feature
		}
		if farthest == nil || distance > fDistance {
			fDistance = distance
			farthest = feature
		}
	}

	if req.Mode == pb.RecommendationMode_GetFarthest{
		return farthest, nil
	}else{
		return nearest, nil
	}
}

//implement the interface method

func (r *RouteGuideServer) GetFeature(ctx context.Context, p *pb.Point) (*pb.Feature, error){
	//here we need to match the point and return the matching
	for _,feature := range r.feature{
		//if the localhost is at the same point then return the feature
		if proto.Equal(feature.Location,p){
			return feature,nil
		} //compare with 2 proto buffer message is the same or not
	}

	return nil,nil //nothing is found
}

func (r *RouteGuideServer) GetListOfFeature(rectangle *pb.Rectangle, featureSer pb.RouteGuide_GetListOfFeatureServer) error{
	//here we will receive a Rectangle(2 point) and return keep sending back the return if end return nil or err
	for _, feature := range r.feature{
		//check current point is inside th rectangle
		if inRange(feature.Location,rectangle){
			//if this point is inside the rectangle, send back to client using stream
			//keep sending the feature until err
			if err := featureSer.Send(feature);err != nil{ //
				return  err
			}
		}
	}
	return nil
}

func (r *RouteGuideServer) GetRecordRoute(recordSer pb.RouteGuide_GetRecordRouteServer) error{
	//here we will keep receiving from client and return the feature
	/*
	int32 pointCount = 1;
	int32 distance = 2;
	int32 totalTime = 3;
	*/
	start := time.Now()
	var pointsCount,distance int32
	var prePoint *pb.Point
	for {
		point, err := recordSer.Recv()
		if err == io.EOF{
			end := time.Now()
			//end sending point,send back the summary
			print(distance)
			summary := pb.RouteSummary{
				PointCount: pointsCount,
				Distance: distance,
				TotalTime: int32(end.Sub(start).Seconds()), // return UTC time and set to second and cast to int
			}
			return recordSer.SendAndClose(&summary)
		}
		if err != nil {
			return err
		}
		//counting the point
		pointsCount++
		if prePoint != nil{
			//calculate the distance between previous and current
			distance += calcDistance(prePoint,point)
		}
		prePoint = point // set previous point to current point
	}
	return nil
}

func (r *RouteGuideServer) Recommend(recSer pb.RouteGuide_RecommendServer) error{
	for{
		req, err := recSer.Recv()
		if err != nil {
			return err
		}
		//handle the request and send back the
		recommended,err := r.recommendPoint(req)
		if err != nil{
			return err
		}
		err = recSer.Send(recommended)
		if err != nil {
			return err
		}
	}

}


//define a function to create our service route
func newServer() *RouteGuideServer{
	//return a new server with some data
	return &RouteGuideServer{
		feature: []*pb.Feature{
			//suppose we are setting some data below
			{Name: "上海交通大学闵行校区 上海市闵行区东川路800号", Location: &pb.Point{
				Latitude:  310235000, //here we are making the float point as int
				Longitude: 121437403, //here we are making the float point as int
			}},
			{Name: "复旦大学 上海市杨浦区五角场邯郸路220号", Location: &pb.Point{
				Latitude:  312978870, //here we are making the float point as int
				Longitude: 121503457, //here we are making the float point as int
			}},
			{Name: "华东理工大学 上海市徐汇区梅陇路130号", Location: &pb.Point{
				Latitude:  311416130, //here we are making the float point as int
				Longitude: 121424904, //here we are making the float point as int
			}},
		},
	}
}

func main(){
	//server side
	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}

	//create a new grpc server
	grpcSer := grpc.NewServer()

	//here we need to register our service to grpc server
	pb.RegisterRouteGuideServer(grpcSer,newServer())
	print("server is on!")
	log.Fatalln(grpcSer.Serve(listen)) //is grpc server has some error ,then fatal
}