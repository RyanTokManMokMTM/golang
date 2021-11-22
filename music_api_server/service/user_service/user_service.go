package user_service

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	tool "music_api_server/Tool"
	userModel "music_api_server/repositories/user"
	req "music_api_server/request"
	"time"
)

//UserService Implement the user_service Protocol
type UserService struct {
}

func (user *UserService)Register(req *req.RegisterRequest) error {
	bcrypt := tool.BcryptDefault
	hashPassword , err := bcrypt.MakePassword([]byte(req.Password))
	if err != nil{
		return err
	}

	req.Password = string(hashPassword)
	return userModel.CreateUser(req) // the model will return the error
}

func (user *UserService)Login(req *req.LoginRequest)  (string,error) {
	info := userModel.GetUserByEmail(req.Email)
	if info.ID == 0{
		return "",errors.New("Email is not existed.")
	}

	bcrypt := tool.BcryptDefault
	err := bcrypt.ComparePassword([]byte(info.Password), []byte(req.Password))
	if err != nil {
		return "", errors.New("Email or password is incorrect.")
	}

	token := tool.NewJsonWebToken()
	jwt, err := token.CreateToken(tool.CustomClaims{
		Name:fmt.Sprintf("%s,%s",info.FirstName,info.LastName),
		Email:info.Email,
		StandardClaims : jwt.StandardClaims{
			Issuer: "jackson.tmm",
			ExpiresAt: int64(time.Now().Unix() + 3600),
			NotBefore: time.Now().Unix()-1000 , //Before the time not available
			Subject: "Authorization",
		},
	})

	if err != nil {
		return "",errors.New("The token is failed to generated")
	}

	//Genera the JWT Token
	return jwt,nil
}

