package phpgo

import (
	"fmt"
	"reflect"
)

//对应 php的 array 数组
type Array []interface{}   //1 列表格式
type ArrayMap map[interface{}]interface{} //2 map格式




// in_array() 检查元素是否在 数组内
func (array Array) In (data interface{}) bool{
	for _,v := range array{
		typ := reflect.TypeOf(v)
		va := reflect.ValueOf(v)

		if reflect.TypeOf(v) == reflect.TypeOf(data) {
			if fmt.Sprintf("%qv", v) == fmt.Sprintf("%qv", data) {
				return true
			}
		}

		//支持 混合数组，map形式的
		if typ.Kind() == reflect.Map {
			for _,key := range va.MapKeys(){
				if fmt.Sprintf("%qv", va.MapIndex(key)) == fmt.Sprintf("%qv", data) {
					return true
				}

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
	for _,v := range values{
		*array = append(*array,v)
	}
	return *array
}



