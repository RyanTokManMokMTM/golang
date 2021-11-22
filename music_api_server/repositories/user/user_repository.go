package user

import (
	"errors"
	orm "music_api_server/database"
	"music_api_server/model/user_model"
	req "music_api_server/request"
	"strings"
)

func CreateUser(data *req.RegisterRequest) error{
	//TODO - Assign Request data to DB data
	info := user_model.UserInfo{}
	info.FirstName = data.FirstName
	info.LastName = data.LastName
	info.Email = strings.ToLower(data.Email)
	info.Password = data.Password
	info.Icon = data.Icons[0].Filename
	var existEmail int
	orm.DB.Raw("SELECT COUNT(1) as exist FROM user_infos where email = ?",info.Email).Scan(&existEmail)
	if existEmail == 1{
		//Exist
		return errors.New("email is already registered")
	}

	err := orm.DB.Create(&info).Error
	return err
}

func GetUserByEmail(email string) *user_model.UserInfo{
	info := user_model.UserInfo{}
	orm.DB.Where("email = ?",strings.ToLower(email)).First(&info)
	return &info
}

func DeleteUserByEmail() error{
	return nil
}

func UpdateUser(){

}

func UpdateUserById(id int){

}


