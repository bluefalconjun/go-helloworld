package function

import (
	"fmt"
	"math"
	//"strings"
)

type Vertex struct {
	X, Y float64
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func Funclosures() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	v := &Vertex{3, 4}
	fmt.Println(v.Abs())
}

//this is a method defined on Vertex struct.
//method can be defined on any type(including struct) in current package.
//method can not be defined on other package type or basic types
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
