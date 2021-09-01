//Unit Testing
package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestHandler(t *testing.T){
	//testing.T mean unit test , a basic testing

	req,err := http.NewRequest(http.MethodGet,"http://localhost:8080/cal?val=3",nil)
	if err != nil{
		t.Errorf("can not create a request %v",err)
	}

	//not real ,just for testing
	res := httptest.NewRecorder() // a response recorder
	Handler(res,req)

	//get the response result
	result := res.Result()

	//check the
	if result.StatusCode != http.StatusOK{
		//response some err
		//testing failed
		t.Errorf("received status code %d,but expect %d",result.StatusCode,http.StatusOK)
	}

	//read body
	bytes,err := io.ReadAll(res.Body)
	if err != nil{
		t.Errorf(" cannot read data form response body,err %v",err)
	}

	val,err := strconv.Atoi(strings.TrimSpace(string(bytes))) // TrimSpace return all space include \t etc
	if err != nil{
		t.Errorf("can not convert to int ,err %v",err)
	}

	//check the val
	if val != 6{
		//we want 6 here
		t.Errorf("received reuslt not correct, %v",val)
	}
}

/*
   --- PASS: TestHandler/input_2 (0.00s)
   --- PASS: TestHandler/input_4 (0.00s)
   --- PASS: TestHandler/input_str (0.00s)
   --- PASS: TestHandler/input_nil (0.00s)

*/

func TestHandler2(t *testing.T){
	//Testing with group data
	//define a data structure for testCase(what status is testing expected)
	testCases := []struct {
		name string `test case name`
		inData string `request input`
		result int `response result`
		err string `http error`
		status int `http status`
	}{
		{name:"input 2",inData: "3",result: 6,err: "",status: http.StatusOK},
		{name:"input 4",inData: "4",result: 8,err: "",status: http.StatusOK},
		{name:"input str",inData: "str",err: "the value is not a int",status: http.StatusBadRequest},
		{name:"input nil",inData: "",err: "required a value",status: http.StatusBadRequest},
	}

	//run each test cash with t.Run(testCaseName,testingFunction)
	for _, testCase := range testCases{
		//for each test case
		t.Run(testCase.name, func(t *testing.T) {
			//send  request
			req, err := http.NewRequest(http.MethodGet,"http://localhost:8080/cal?val="+testCase.inData,nil)
			if err != nil{
				t.Fatalf("can not create the request %v",err)
			}

			//create a response recorder
			resRecorder := httptest.NewRecorder()

			//run testing function with req and res
			Handler(resRecorder,req)
			res := resRecorder.Result()

			//check the statusCode
			if res.StatusCode != testCase.status{
				t.Fatalf("statusCode received %d and expect %d",res.StatusCode,testCase.status)
			}

			//read body
			rBytes ,err := io.ReadAll(res.Body)
			if err  != nil{
				t.Fatalf("can not get the reponse the body %v",err)
			}
			defer res.Body.Close()

			//we may need to test bad status code
			//if it's a bad status code ,won't run successfully code

			trimStr := strings.TrimSpace(string(rBytes))
			//check status is ok
			if res.StatusCode != http.StatusOK{
				//is bad or other Status
				if trimStr != testCase.err{
					//if the converted error not equal to what we are expected
					//failed the test
					t.Errorf("received error message %v and expect %v",trimStr,testCase.err)
				}
				return //if not error pass the test
			}

			//compare the result
			val,err := strconv.Atoi(trimStr)
			if err != nil{
				t.Fatalf("can not convert to int %v",err)
			}

			if val != testCase.result{
				t.Fatalf("result %v is not expected,%v",val,testCase.result)
			}
		})
	}

}