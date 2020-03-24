package interfaces

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func Interfaces() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat has method Abser
	a = &v // a *Vertex has method Abser

	// this is wrong because v is type of Vertex（but not *Vertex）
	// Vertex has no Abser。
	//a = v

	fmt.Println(a.Abs())
	SayHiSample()
}

type MyFloat float64

//this is a method
//"A method is a function with an implicit first argument, called a receiver."
//MyFloat is the receiver, Abs is the func name.
//method is bind to receiver
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
