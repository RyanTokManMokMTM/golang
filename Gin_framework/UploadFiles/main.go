package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main(){
	server := gin.Default()

	//set the maximum multipart memory for uploading the file size

	//multipart/form-data -> allow user to send Multi data
	server.MaxMultipartMemory = 8 << 20
	server.POST("/singleUpload",singleFileHandler)
	server.POST("/multiFilesUploads",multipleFilesHandler)
	server.Run(":8080")
}

//upload a single file
func singleFileHandler(c *gin.Context){

	//read data form header of multipart/form-data and the map key is file
	file,_ := c.FormFile("file")
	log.Println(file.Filename)

	c.SaveUploadedFile(file,"a.jpg")
	c.String(http.StatusOK,fmt.Sprintf("%s uploaded!",file.Filename))

}

func multipleFilesHandler(c *gin.Context){
	form,_ := c.MultipartForm() //get the form of multipart/form-data header field
	//get the file form
	log.Println(form)
	files  := form.File["file"] //get the value from the map,return a file header of a list
	//File header is a struct,store the info of the file
	//loop over the list
	for _,file := range files{
		//log.Println(file)
		log.Println(file.Filename)
		//store the data  to dst name
		c.SaveUploadedFile(file,file.Filename)
	}
	c.String(http.StatusOK,fmt.Sprintf("%d files uploaded",len(files)))
}

/*
How multipart/form-data works?
Content-type : multipart/form-data -> 1次傳送多個文件

data format of multipart/form-data
由boundary + Content-Disposition

boundary :一個key 知道多個資料的多個界(怎麼分開) -> 看到這個就知道為什麼時候是當前資料的結束
boundary 格式的定義 由 2個hypen + 長度<= 70 ASCII

Content-Disposition : 描述當前資料
有一個field 記錄名字 如果後面跟著filename 會在下一行表該文件的content-type
空一行後 才會檔案內容(如果是圖片 是已binary表示)
Content-Disposition: form-data; name="??"(name field is??); filename="a.txt"
content-type:text/plain
/r/n
txt content
// image example

Content-Disposition: form-data; name="??"(name field is??); filename="png"
content-type:img/png
/r/n

png

binary string....

例子:
POST /upload HTTP/1.1
Host: localhost:3000

Content-Type: multipart/form-data; boundary=----WebKitFormBoundaryFYGn56LlBDLnAkfd
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.66 Safari/537.36

------WebKitFormBoundaryFYGn56LlBDLnAkfd
Content-Disposition: form-data; name="name"

Test
------WebKitFormBoundaryFYGn56LlBDLnAkfd
Content-Disposition: form-data; name="file"; filename="text.txt"
Content-Type: text/plain

Hello World
------WebKitFormBoundaryFYGn56LlBDLnAkfd--

*/