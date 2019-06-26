package phpgo

import (
	"testing"
)


func TestInt(t *testing.T) {
	int_slice := *&[]interface{}{int(1), int8(1), int16(1), int32(1), int64(1)}
	for _, num := range int_slice {
		if r := Strval(num); r != "1" {
			t.Fatal("Int conv failed with:", r, num)
		}
	}
	minus_int_slice := *&[]interface{}{int(-1), int8(-1), int16(-1), int32(-1), int64(-1)}
	for _, num := range minus_int_slice {
		if r := Strval(num); r != "-1" {
			t.Fatal("Int conv failed with:", r, num)
		}
	}
}

func TestString(t *testing.T) {
	int_slice := *&[]interface{}{"1as"}
	for _, num := range int_slice {
		if r := Strval(num); r != "1as" {
			t.Fatal("string conv failed with:", r, num)
		}
	}
	minus_int_slice := *&[]interface{}{"0"}
	for _, num := range minus_int_slice {
		if r := Strval(num); r != "0" {
			t.Fatal("string conv failed with:", r, num)
		}
	}
}

func TestBool(t *testing.T) {
	int_slice := *&[]interface{}{true}
	for _, num := range int_slice {
		if r := Strval(num); r != "1" {
			t.Fatal("true conv failed with:", r, num)
		}
	}
	minus_int_slice := *&[]interface{}{false}
	for _, num := range minus_int_slice {
		if r := Strval(num); r != "" {
			t.Fatal("string conv failed with:", r, num)
		}
	}
}


