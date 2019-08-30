package programarch

import (
	"fmt"
)

//go is capital sensitive
var heapSort = 1
var Heapsort = 2

//go use 25 key word
/*
	break      default       func     interface   select
	case       defer         go       map         struct
	chan       else          goto     package     switch
	const      fallthrough   if       range       type
	continue   for           import   return      var
*/

//pre-defined words
//pre-defined words are not keyword, you can redefined sometime but NOTE!
//carefully use it.
//Guess what happens in C: 
//#define true false

/*
	//const
	true false iota nil
	//type
	int int8 int16 int32 int64
	uint uint8 uint16 uint32 uint64 uintptr
	float32 float64 complex128 complex64
	bool byte rune string error
	//func
	make len cap new append copy close delete
	complex real imag panic recover
*/

//Namerule is 
func Namerule(){
	fmt.Println("program arch!")
}