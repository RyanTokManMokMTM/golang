package main

import (
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"os"
	"simple_protoBuff/src/complexMessagepb"
	"simple_protoBuff/tutorial"
	"time"

	"simple_protoBuff/src/simplepb"
	"simple_protoBuff/src/DataTimepb"
)

func main() {
	//pb file -> generate
	simpleBPTest()
	simpleEnumTest()
	complexBPTest()
	test()
}

func simpleBPTest(){
	message := doSimple()
	//to disk file
	err := toDisk("simple.bin", message)
	if err != nil {
		log.Println(err)
		return
	}
	err = fromDisk("simple.bin")
	if err != nil {
		log.Println(err)
		return
	}

	//to json file
	jsonData, err := toJSON(message)
	if err != nil {
		return
	}

	fmt.Println(jsonData)

	err = formJSON(jsonData)
	if err != nil {
		log.Println(err)
		return
	}
}

func simpleEnumTest(){
	//all enum value defined in proto buffer -> pb -> global value
	//also will 2 map string->id and id->string
	//and a message struct
	message := DataTimepb.DateTime{
		Id: 1,
		Day : DataTimepb.WeekDay_DAY_ONE_MONDAY,
	}
	fmt.Println(message)
}

func complexBPTest(){
	//the proto buffer with 2 same level message
	//and this 2 message will convert to 2 strut in pb file

	complexMsg := complexMessagepb.ComplexMessage{
		Dummy_A: &complexMessagepb.DummyMsg{
			Id: 1,
			Name: "Jackson",
		},
		Dummy_B: []*complexMessagepb.DummyMsg{
			{
				Id: 2,
				Name: "Tom",
			},
			{
				Id: 3,
				Name: "John",
			},
		},
	}
	fmt.Println(complexMsg)
}

func test(){
	person := &tutorial.Person{
		Id: 1,
		Name: "jackson",
		Email: "RyanTokManMokMTM@hotmail.com",
		Phones: []*tutorial.Person_PhoneNumber{
			{
				Number: "0968200176",
				Type:tutorial.Person_MOBILE,
			},
			{
				Number: "65828151",
				Type: tutorial.Person_WORK,
			},
			{
				Number: "24087557",
				Type: tutorial.Person_HOME,
			},
		},
		LastUpdated: timestamppb.New(time.Now()),
	}

	//out put to the disk
	messageToDisk("person.bin",person)

	messagePersonDisk := messageFromDisk("person.bin")
	fmt.Println(messagePersonDisk)

	messageJSONStr := messageToJSON(person)
	fmt.Println(messageJSONStr)

	messageJSON := messageFromJson(messageJSONStr)
	fmt.Println(messageJSON)

}

func messageToDisk(fileName string,pb proto.Message){
	bytesData, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln(err)
		return
	}

	err = os.WriteFile(fileName, bytesData, 0666)
	if err != nil {
		log.Fatalln(err)
		return
	}

	log.Printf("saved %s successfully!",fileName)
}

func messageFromDisk(fileName string) proto.Message {
	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	person := tutorial.Person{}
	err = proto.Unmarshal(file, &person)
	if err != nil {
		log.Fatalln(err)
	}

	return &person

}

func messageToJSON(pb proto.Message) string{
	jsonMarshaler := jsonpb.Marshaler{}
	messageStr, err := jsonMarshaler.MarshalToString(pb)
	if err != nil {
		log.Fatalln(err)
		return ""
	}

	return messageStr
}

func messageFromJson(jsonStr string) proto.Message{
	person := tutorial.Person{}
	err := jsonpb.UnmarshalString(jsonStr, &person)
	if err != nil {
		log.Fatalln(err)
	}

	return &person
}
//Proto package provided a bunch of function allow dev to convert the message to other format
func doSimple() *simplepb.SimpleMessage{
	message := simplepb.SimpleMessage{
		Id:         123,
		Name:       "Jackson",
		IsSimple:   true,
		SimpleList: []int32{1, 2, 3, 4},
	}
	return &message
}
//
func toDisk(fileName string,pb proto.Message) error {
	//marshal a message to bytes
	data, err := proto.Marshal(pb)
	if err != nil {
		log.Println(err)
		return err
	}
	//os perm 0-7 0:No Access,
	/*-rwx(own)rwx(group)rwx(other)
		0:No Access,
		1:executable
		2:writeable
		3:writable and executable
		4:Readable
		5:Readable and executable
		6:writable and readable
		7:writable, readable and executable
	*/
	err = os.WriteFile(fileName, data, 0666)
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println("message is out to the file.")
	return nil
}

func fromDisk(fileName string)error{
	message :=simplepb.SimpleMessage{}
	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Println(err)
		return nil
	}

	err = proto.Unmarshal(file, &message)
	if err != nil {
		log.Println(err)
		return err
	}

	fmt.Println(message)
	return nil
}

func toJSON(pb proto.Message) (string,error){
	jsonMarshler := jsonpb.Marshaler{} //empty setting of the Marshaler
	jsonStr, err := jsonMarshler.MarshalToString(pb)
	if err != nil {
		return "",err
	} //convert to
	return jsonStr,nil
}

func formJSON(jsonStr string)error {
	message := simplepb.SimpleMessage{}
	err := jsonpb.UnmarshalString(jsonStr, &message)
	if err != nil {
		return err
	}
	fmt.Println(message)
	return nil
}