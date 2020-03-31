package packages

/*
导入多个包时, 使用()园括号进行多个包的整合
*/

import (
	"fmt"
	"math"
)

func Imports() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}
