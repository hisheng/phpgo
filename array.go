package phpgo

import (
	"fmt"
	"reflect"
)

//对应 php的 array 数组
type Array []interface{}



// in_array() 检查元素是否在 数组内
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

// array_shift 弹出数组的第一个
func (array *Array) Shift () interface{}{
	if len(*array) > 0{
		e := (*array)[0]
		*array = (*array)[1:]
		return e
	}

	return nil
}

// array_pop 弹出数组的最后一个
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


// array_push  将一个或多个单元压入数组的末尾
func (array *Array) Push(values ...interface{}) Array {
	*array = append(values)
	return *array
}



