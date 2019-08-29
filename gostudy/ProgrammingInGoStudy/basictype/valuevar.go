//package main

package basictype

import (
	"fmt"
	"math"
	"math/big"
)

/*
use math.MaxInt32 to check the max value of define.
*/

func Uint8FromInt(x int) (uint8, error) {
	if 0 <= x && x <= math.MaxUint8 {
		return uint8(x), nil
	}
	return 0, fmt.Errorf("%d is out of uint8 rage", x)
}

/*
type: value range

byte	=uint8
int		=based on platform int32/int64
int8	=[-128,127]
...

rune	=uint32
uint	=based on platform uint32/uint64
uintptr	=unsigned int pointer as uint32/uint64
*/

/*
use big package to handle unlimited values(based on memory on platform)

big.Int //integer
big.Rat //有睆数
*/

//func main() {
func countPi() {
	places := 1000
	scalePi := fmt.Sprint(Pi(places))
	fmt.Printf("3.%s\n", scalePi[1:])
}

//Pi returns Pi
func Pi(place int) *big.Int {
	gigits := big.NewInt(int64(place))
	unity := big.NewInt(0)
	ten := big.NewInt(10)
	exponent := big.NewInt(0)
	unity.Exp(ten, exponent.Add(gigits, ten), nil)
	pi := big.NewInt(4)
	return pi.Div(pi, big.NewInt(0).Exp(ten, ten, nil))
}
