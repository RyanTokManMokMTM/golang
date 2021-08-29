package downloader

import (
	"encoding/json"
	"io"
	"net/http"
)

//struct for request
type Request struct {
	Bvid []string //need a group of ids
}

//the format of the data
type VideoData struct{
	Code int
	Data struct{
		Bvid string `json:"id"`
		Title string `json."title"`
		Desc string `json."descption"`
	} `json Data`
}

//data for response
type Response struct {
	Result []VideoData
}

//to download data
func Downloader(req Request) (Response,error){
	//response massage
	var response Response //a array of video data
	for _,id := range req.Bvid{ //for each id
		var videoInfo VideoData
		resp,err := http.Get("https://api.bilibili.com/x/web-interface/view?bvid="+id) //to get json
		if err != nil{
			return Response{},err //can't find url,not have response
		}
		bodyBytes,err := io.ReadAll(resp.Body) //read all the body
		if err != nil{
			return Response{},err //can't read the res body
		}

		//change our bytes to json format
		if err = json.Unmarshal(bodyBytes,&videoInfo);err != nil{
			return Response{},err
		}
		resp.Body.Close( //close the writer form
		response.Result = append(response.Result, videoInfo)
	}
	return response,nil //send pack the request
}