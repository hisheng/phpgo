// strval method of php implement by Golang
// Copyright (C) 2015 zhang haisheng <465908774@qq.com>
// License can be found in the LICENSE file

package phpgo

import (
	"fmt"
	"reflect"
	"strconv"
)

func Strval(data interface{}) (s string) {
	v := reflect.ValueOf(data)
	typ := reflect.TypeOf(data)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		v = v.Elem()
	}

	switch typ.Kind() {

	case reflect.String:
		s = v.String()

	case reflect.Int,reflect.Int8,reflect.Int16,reflect.Int32,reflect.Int64:
		s = strconv.FormatInt(v.Int(),10)
	case reflect.Bool:
		if v.Bool() {
			s = "1"
		} else {
			s = ""
		}
	case reflect.Float32,reflect.Float64:
		s = strconv.FormatFloat(v.Float(),'f',12,64)
	case reflect.Uint,reflect.Uint8,reflect.Uint16,reflect.Uint32,reflect.Uint64:
		s = strconv.FormatUint(v.Uint(),10)
	case reflect.Ptr:
		v := v.Elem()
		s = Strval(v)
	case reflect.Map:
		s = "map"
	case reflect.Array:
		s = "array"
	case reflect.Chan:
		s = "chan"
	case reflect.Slice:
		s = "slice"
	case reflect.Struct:
		s = "struct"
	default:
		s = fmt.Sprintf("%qv", data)
	}
	return s
}


