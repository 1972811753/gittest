package main

import (
	"fmt"
)

func Receiver(v interface{}) {
	switch v.(type){
	case string:
		fmt.Println("这个是string")
	case int:
		fmt.Println("这个是int")
	case bool:
		fmt.Println("这个是bool")
	default:
		fmt.Println("这个是未知类型")
	}

}
func main(){
	Receiver("你好吗")
	Receiver(32)
	Receiver(true)
}
