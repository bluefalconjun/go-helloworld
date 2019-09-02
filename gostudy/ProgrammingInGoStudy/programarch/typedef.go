package programarch

import (
	"fmt"
)

//typedef grammer
//type $typename $lowerstru_type
//For all type T, a determined T(x) is used to transfer x to type T.
//only x and T have the same $lowerstru_type

//Celsius is C
type Celsius float64    //C
//Fahrenheit is F
type Fahrenheit float64 //F

const (
	//AbsoluteZeroC is zero C.
	AbsoluteZeroC Celsius = -273.15
	//FreezingC is zero
	FreezingC     Celsius = 0
	//BoilingC is 100
    BoilingC      Celsius = 100
)

//CToF transfer C to F.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

//FToC transfer F to C.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

//Typedef is
func Typedef(){

	AbsoluteZeroF := CToF(AbsoluteZeroC)

	fmt.Println(AbsoluteZeroF)
}