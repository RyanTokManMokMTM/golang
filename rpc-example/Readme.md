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

