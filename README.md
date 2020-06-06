# phpgo
用 go 实现 php 的array数组以及常用函数，给php转到go来的人用    


## install
go get github.com/hisheng/phpgo
    
    
        取值
        php 的 strval() 对应 phpgo.Strval()
        isset() ---> 待做
        
        时间函数
        php的 date("Y-m-d H:i:s") 对应 phpgo.Date()
    
        
        
        array
        php的array 数组结构 对应 phpgo.Array
        in_array()   phpgo.Array.In()
        array_pop()  phpgo.Array.Pop()    弹出数组最后一个
        array_shift()  phpgo.Array.Shift()    弹出数组第一个
        count()      phpgo.Array.Count()
        array_push()  phpgo.Array.Push()    数组尾部增加一个或者多个值
        json_encode  phpgo.Array.JsonEncode()   array 变成json
        json_decode  phpgo.JsonDecode()   json 变 array

        
        
        还需要实现 in_array() 直接传一个 []int,[]string,之类的看查找是否在
        
        
        
        
        
        
        
        
