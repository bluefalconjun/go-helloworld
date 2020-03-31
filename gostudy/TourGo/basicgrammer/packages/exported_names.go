package packages

import (
	"fmt"
	"math"
)

/*
不能使用引用包中未被导出的变量/函数(非大写字母开头的)
打印在不指定格式时, 按照默认格式输出?
*/

func Exported() {
	//fmt.Println(math.pi)
	fmt.Println("Pi is ", math.Pi)
}
