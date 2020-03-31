package packages

/*
按照约定，包名与导入路径的最后一个元素一致。
例如，"math/rand" 包中的源码均以 package rand 语句开始
*/

import (
	"fmt"
	"math/rand"
)

func Packages() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
