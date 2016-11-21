package arraytype

import (
	"fmt"
	//"strings"
)

var a = []string{"1", "2", "3"}

//const only can be used to define Char/String/Bool/Int
const b = "this is a string"
const c = false

func ArrayType() {
	//const b=[3]string{"1","2","3"}
	fmt.Println("a is %T, %v, len(a)%v, cap(a)%v", a, a, len(a), cap(a))
	fmt.Println("b is %T, %v, len(b)%v, cap(b)%v", b, b, len(b), len(b))
	fmt.Println(c)
}
