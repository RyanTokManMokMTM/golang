package user_service

import req "music_api_server/request"

//UserServiceProtocol TODO - The Service for user_service
type UserServiceProtocol interface {
	Register(*req.RegisterRequest) error
	Login(*req.LoginRequest) (string,error)
	//MORE SERVICE
}