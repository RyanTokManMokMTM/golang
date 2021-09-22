package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

//gin using binding to instead of validate:required

type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

func main(){
	server := gin.Default()
	// need to bind the custom validator
	//calling validate and register  validation function
	//there are some built-in tag
	//the tags for validation
	if value,ok := binding.Validator.Engine().(*validator.Validate); ok {
		value.RegisterValidation("bookabledate",bookingValidator)
	}

	server.GET("/booking",customTimerValidatorHandle)
	server.Run(":8080")
}

func customTimerValidatorHandle(c *gin.Context){
	//binding with shouldBindWith
	var bookingData Booking
	if err := c.ShouldBindWith(&bookingData,binding.Query);err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"message":err.Error(),
		})
		return
	}
	c.String(http.StatusOK,"Booking dates are valid")

}


var bookingValidator validator.Func = func(fl validator.FieldLevel) bool{
	fmt.Println(fl.FieldName())
	date,ok := fl.Field().Interface().(time.Time)
	if ok{
		now := time.Now()
		if now.After(date){ //check the date
			return false
		}
	}
	return true
}