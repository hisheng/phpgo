package phpgo

import (
	"testing"
)


//目前只测试了 int ，bool，string
func TestIn(t *testing.T) {
	var array  = Array{"a",1,true}
	if !array.In("a") {
		t.Fatal("in_array failed ", array.In("a"))
	}
	if !array.In(1) {
		t.Fatal("in_array failed 1")
	}
	if array.In("1") {
		t.Fatal("in_array failed 1")
	}
	if !array.In(true) {
		t.Fatal("in_array failed true")
	}
}



func TestPop(t *testing.T) {
	var array  = Array{"a",1,true}

	if array.Pop() != "a" {
		t.Fatal("pop failed a")
	}
	if array.Pop() != 1 {
		t.Fatal("in_array failed 1")
	}
	if array.In("1") {
		t.Fatal("in_array failed 1")
	}
	if array.Pop() != true {
		t.Fatal("in_array failed true")
	}
}
