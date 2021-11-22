package sample

import "testing"

func Test_add(t *testing.T){
	if add(200,200) != 400{
		t.Errorf("Wrong ans")
	}
}

func Test_add_2(t *testing.T){
	if add(10,10) != 20 {
		t.Errorf("Wrong ans")
	}
}

func Test_negative_add(t *testing.T){
	if add(-1,-20) < 0 {
		t.Errorf("Can't be negative value")
	}
}