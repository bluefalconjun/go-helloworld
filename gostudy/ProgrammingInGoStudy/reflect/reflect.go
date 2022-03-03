package main

import (
	"fmt"
	"reflect"
)

func reflectsetvalue1(x interface{}) {
	value := reflect.ValueOf(x)
	if value.Kind() == reflect.String {
		value.SetString("Welcome to go world")
	}
}
func reflectsetvalue2(x interface{}) {
	value := reflect.ValueOf(x)
	// 反射中使用Elem()方法获取指针所指向的值
	if value.Elem().Kind() == reflect.String {
		value.Elem().SetString("Welcome to go world")
	}
}

func main() {
	var booknum float32 = 6
	var isbook bool = true
	bookauthor := "www.go.dev"
	bookdetail := make(map[string]string)
	bookdetail["Go"] = "Welcome to go world"

	//Typeof返回接口中保存的值得类型，Typeof(nil)会返回nil
	fmt.Println(reflect.TypeOf(booknum))
	fmt.Println(reflect.TypeOf(isbook))
	fmt.Println(reflect.TypeOf(bookauthor))
	fmt.Println(reflect.TypeOf(bookdetail))

	//ValueOf返回接口中保存的值
	fmt.Println(reflect.ValueOf(booknum))
	fmt.Println(reflect.ValueOf(isbook))
	fmt.Println(reflect.ValueOf(bookauthor))
	fmt.Println(reflect.ValueOf(bookdetail))

	address := "www.go.dev"
	// reflectsetvalue1(address)
	// 反射修改值必须通过传递变量地址来修改。若函数传递的参数是值拷贝，则会发生下述错误。
	// panic: reflect: reflect.Value.SetString using unaddressable value
	reflectsetvalue2(&address)
	fmt.Println(address)

}
