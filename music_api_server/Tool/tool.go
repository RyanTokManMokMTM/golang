package Tool

import (
	"crypto/rand"
	"encoding/base64"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
)

var BcryptDefault *Bcrypt

func init(){
	BcryptDefault = &Bcrypt{Cost: 10}
}

type Bcrypt struct {
	Cost int
}

func (b *Bcrypt)MakePassword(password []byte)  ([]byte,error)  {
	hashPassword, err := bcrypt.GenerateFromPassword(password,b.Cost)
	if err != nil {
		return nil, err
	}

	return hashPassword,nil
}

func (b* Bcrypt)ComparePassword(hashedPassword,password []byte) error{
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	if err != nil {
		return err
	}
	return nil
}


func Base64KeysGenerator(size int) string {
	bytes := make([]byte, size)
	_, err := rand.Read(bytes)
	if err != nil {
		log.Println(err)
		return ""
	}
	return base64.URLEncoding.EncodeToString(bytes)
}

func PathExists(path string) (bool,error){
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err){
		return false,nil
	}
	return false,err
}
