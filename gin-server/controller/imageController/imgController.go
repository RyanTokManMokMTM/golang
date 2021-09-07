package imageController

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"sync"
)

type ImageController interface {
	//Define get all
	GetAll(ctx *gin.Context)

	//Define create
	Create(ctx *gin.Context)

	//Define update
	Update(ctx *gin.Context)

	//Define delete
	Delete(ctx *gin.Context)
}

//create a new ImageController
func NewImageController() ImageController {
	return &imageController{
		imageDB: make([]imgModel,0),
	}
}

type generateIdCounter struct{
	counter int
	mux sync.Mutex
}

func (g *generateIdCounter) generator() int{
	g.mux.Lock()
	defer g.mux.Unlock()
	g.counter++
	return g.counter
}

//create a global instance
var g *generateIdCounter = &generateIdCounter{}

type imgModel struct {
	Id int `uri:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
}

type imageController struct{
	//such a db
	imageDB []imgModel
}

func (c *imageController)GetAll(ctx *gin.Context){
	ctx.JSON(200,c.imageDB)
}

func (c *imageController)Create(ctx *gin.Context){
	newImg := imgModel{Id: g.generator()}
	if err := ctx.Bind(&newImg); err != nil{
		ctx.JSON(400,gin.H{
			"message":"something is wrong",
		})
		return
	}
	c.imageDB = append(c.imageDB,newImg)
	ctx.String(200,"create img success")
}

func (c *imageController)Update(ctx *gin.Context){
	var updateImgData imgModel
	//fmt.Println(ctx.Params.Get("id"))
	//our uri is id
	if err := ctx.ShouldBindUri(&updateImgData); err != nil{ //just bind the uri
		ctx.JSON(400,"bad request and update failed")
		return
	}

	if err := ctx.ShouldBindJSON(&updateImgData); err != nil{
		ctx.JSON(400,"bad request and update failed")
		return
	}
	fmt.Println(updateImgData)
	for idx,data := range c.imageDB{
		if data.Id == updateImgData.Id{
			//is matching ,replace
			c.imageDB[idx] = updateImgData
			ctx.JSON(200,"updated data succeed!")
			return
		}
	}
	ctx.String(400,"bad request,can't update img")
}

func (c *imageController)Delete(ctx *gin.Context){
	//we need to find the id first
	id,_ := ctx.Params.Get("id")
	s,err:=strconv.Atoi(id)
	if err != nil{
		ctx.String(400,"bad request,id not int")
		return
	}
	for idx,data := range c.imageDB{
		if data.Id == s{
			c.imageDB = append(c.imageDB[0:idx],c.imageDB[idx+1:len(c.imageDB)]...)
			ctx.JSON(200,"Delete data succeed!")
			return
		}
	}
	ctx.String(400,"bad request,can't delete image")
}