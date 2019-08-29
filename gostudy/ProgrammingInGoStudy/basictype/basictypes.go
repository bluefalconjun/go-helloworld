package basictype

import (
	"fmt"
	"math/cmplx"
)

// var is the keyword to define variable. should use as
// var $var_name $var_type
// var keyword could be used in both inside/outside func.

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func Basictypes() {
	//%T means print the type of var, %v means value of var?
	// see https://golang.org/pkg/fmt/ for details
	const f = "%T(%v)\n"
	const v = 5
	const w int = 5

	//avoid to use := just to ignore var !!?

	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
	fmt.Printf(f, v, v)
	fmt.Printf(f, w, w)

}
