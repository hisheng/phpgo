package phpgo

import (
	"fmt"
	"reflect"
)

//对应 php的 array 数组
type Array []interface{}



// in_array()
func (array Array) In (data interface{}) bool{
	for _,v := range array{
		if reflect.TypeOf(v) == reflect.TypeOf(data) {
			if fmt.Sprintf("%qv", v) == fmt.Sprintf("%qv", data) {
				return true
			}
		}
	}
	return false
}

// array_shift
func (array *Array) Shift () interface{}{
	if len(*array) > 0{
		e := (*array)[0]
		*array = (*array)[1:]
		return e
	}

	return nil
}

// array_shift
func (array *Array) Pop () interface{}{
	if len(*array) > 0{
		last := len(*array) - 1
		e := (*array)[last]
		*array = (*array)[:last]
		return e
	}
	return nil
}

// count
func (array Array) Count () int{
	return len(array)
}



