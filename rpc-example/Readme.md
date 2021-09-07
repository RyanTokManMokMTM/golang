## gRPC Example
* Need a Protocol Buffers 
  * Is Used to generate message and service code
  * `Message` and `serverice`
  * ```proto
    //message data struct
    message SomeMessage{
        //type name = tag(a number)
        string name = 1; //field name
        int32 id = 2; //field id
        bool has_book = 3;//field has_book
    } 
    
    message SomeMessageReply{
        string msg = 1; //reply message
    } 
    ```
  * ```proto
    service SomeService{
        //what to do with the data
        //define service
        rpc SayHello(SomeMessage) returns (SomeMessageReply){}
        rpc Working(SomeMessage) returns (SomeMessageReply) {}
        
        //etc
    }
    ```
---
### Use Command to generate the client `stub` and server `stub`
```
//Command
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative route.proto
```
> Command above will generate 2 files

``protofilename.pb.go`` 
> Contain all message types: populate,serialize and retrieve request and response

``protofilename_grpc.pb.go``
> Contain all the service for client to call with

> Contain all interface type for servers to implements

---
## Process  explained

```go
//Client
func main(){
	connection, err := grpc.Dial("localhost:8080",grpc.WithInsecure(),grpc.WithBlock())
	if err != nil {
		return
	}
	defer connection.Close() //close the connection
	client := pb.NewRouteGuideClient(connection)

}
```

> 1.Connect to `gRpc` server using `grpc.Dial`
>
> 2.Create a new `Guide Client` base on the `connection`
>
> ```go
> type RouteGuideClient interface {
> 	//unary => send a request and response a request
> 	GetFeature(ctx context.Context, in *Point, opts ...grpc.CallOption) (*Feature, error)
> 	//streaming
> 	//server,keep streaming Feature from request points
> 	GetListOfFeature(ctx context.Context, in *Rectangle, opts ...grpc.CallOption) (RouteGuide_GetListOfFeatureClient, error)
> 	//client,keep sending point and return a summary
> 	GetRecordRoute(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_GetRecordRouteClient, error)
> 	//2 side streaming
> 	Recommend(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_RecommendClient, error)
> }
> ```
>
> 

```go
//Server
func main(){
	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	grpcSer := grpc.NewServer()
	pb.RegisterRouteGuideServer(grpcSer,newServer())
	log.Fatalln(grpcSer.Serve(listen)) //is grpc server has some error ,then fatal
}
```

> 1. Open a net listening `port` base on `tcp connection`,let client to connect
> 2. Create a new `gRpc` Server
> 3. Regis the `service`(Implement the xxxGuideServer)  to `gRpc` Server
>
> ```go
> type RouteGuideServer interface {
> 	//unary => send a request and response a request
> 	GetFeature(context.Context, *Point) (*Feature, error)
> 	//streaming
> 	//server,keep streaming Feature from request points
> 	GetListOfFeature(*Rectangle, RouteGuide_GetListOfFeatureServer) error
> 	//client,keep sending point and return a summary
> 	GetRecordRoute(RouteGuide_GetRecordRouteServer) error
> 	//2 side streaming
> 	Recommend(RouteGuide_RecommendServer) error
> 	mustEmbedUnimplementedRouteGuideServer()
> }
> ```

---

## Client Stub and Server Stub

![RPC](https://upload.cc/i1/2021/09/07/xTBLku.jpg)

>Progress(Client and Server step 1-9)
>
>**Client Side**
>
>1.Client call local service
>
>2.Client Stub received and packing data as networking message
>
>3.Client Stub find service location(server ip) and send the message to server side
>
>**Server Side**
>
>4.Server Stub received and parsing the data
>
>5.Server Stub according to parsing message and calling correspond servrice
>
>6.Server call the service and return the result to server stub
>
>7.Server Stub received and packing data as networking message and send to client via network
>
>**Client Side**
>
>8.Client Stub received the message and parsing the message
>
>9.Client get the final result

---

