package phpgo

import (
	"fmt"
	"testing"
)


//目前只测试了 int ，bool，string
func TestIn(t *testing.T) {
	var array  = Array{"a",1,true,}
	var arrmap = make(ArrayMap,3)
	arrmap["name"] = "hisheng"
	arrmap["age"] = 31
	arrmap["height"] = 173.5
	array = append(array,arrmap)
	fmt.Println(array)

	array = array.Push("sss")
	fmt.Println(array)

	f := ArrayMap{"girlfriend":"xiaomei"}
	array = append(array,f)
	fmt.Println(array)

	array = array.Push(ArrayMap{"dad":"gang"})
	fmt.Println(array)

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

	if !array.In("hisheng") {
		t.Fatal("hisheng failed true")
	}
	if !array.In("xiaomei") {
		t.Fatal("xiaomei failed true")
	}
	if !array.In("gang") {
		t.Fatal("gang failed true")
	}

	if !array.In("sss") {
		t.Fatal("sss failed true")
	}


}



func TestShift(t *testing.T) {
	var array  = Array{"a",1,true}

	if array.Pop() != true {
		t.Fatal("pop failed true")
	}
	if array.Pop() != 1 {
		t.Fatal("in_array failed 1")
	}
	if array.Pop() != "a"{
		t.Fatal("in_array failed a")
	}

}

func TestPop(t *testing.T) {
	var array  = Array{"a",1,true,"2"}

	if array.Pop() != "2" {
		t.Fatal("pop failed 2")
	}
	if array.Pop() != true {
		t.Fatal("in_array failed true")
	}
	if array.Pop() !=  1 {
		t.Fatal("in_array failed 1")
	}
	if array.Pop() != "a" {
		t.Fatal("in_array failed a")
	}
	if array.Pop() != nil{
		t.Fatal("in_array failed nil")
	}
}

func TestCount(t *testing.T) {
	var array  = Array{"a",1,true}

	if array.Count() != 3 {
		t.Fatal("pop failed 3")
	}
	var array1  = Array{}

	if array1.Count() != 0 {
		t.Fatal("in_array failed 0")
	}
}

func TestPush(t *testing.T) {
	var array  = Array{"a",1,true}
	array.Push("aa")
	if array.Pop() != "aa"{
		t.Fatal("push failed aa")
	}

	var array1  = Array{}
	array1.Push("ss",1)
	if array1.Pop() != 1 {
		t.Fatal("push failed 1")
	}
	if array1.Pop() != "ss" {
		t.Fatal("push failed ss")
	}
}


