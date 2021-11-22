package Tool

import (
	"fmt"
	"testing"
)

func TestBcrypt(t *testing.T){
	password := "12345"
	hash := Bcrypt{Cost: 10}

	hashPassword, err := hash.MakePassword([]byte(password))
	if err != nil {
		t.Fail()
	}

	fmt.Println(hashPassword)

	err = hash.ComparePassword(hashPassword, []byte(password))
	if err != nil {
		t.Fail()
	}

	fmt.Println("Password is matched")
}
