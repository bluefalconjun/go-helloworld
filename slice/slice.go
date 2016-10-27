package slice

import (
	"fmt"
	"strings"
)

func Slice() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("s ==", s)

	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] == %d\n", i, s[i])
	}

	//s[lo:hi] means elements in s from lo to hi-1. excluding hi.
	fmt.Println("s[1:4] ==", s[1:4])

	//ignore lo then lo=0
	fmt.Println("s[:3] ==", s[:3])

	//ignore hi then hi=len(s)-1
	fmt.Println("s[4:] ==", s[4:])
}

func Sliceslice() {
	// Create a tic-tac-toe board.
	game := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	game[0][0] = "X"
	game[2][2] = "O"
	game[2][0] = "X"
	game[1][0] = "O"
	game[0][2] = "X"

	printBoard(game)
}

func printBoard(s [][]string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%s\n", strings.Join(s[i], " "))
	}
}

func SlicesAppend() {
	var a []byte
	printSlice("a", a)

	// append works on nil slices.
	a = append(a, 0)
	printSlice("a", a)

	// the slice grows as needed.
	a = append(a, 1)
	printSlice("a", a)

	// we can add more than one element at a time.
	//in this case. cap(a) is not enough to hold 2,3,4. so new array has beened arranged for it.
	//append method in low implement is different when array type is different. !!!
	//try to change var a[]byte to a[]int and test. !!!
	a = append(a, 2, 3)
	printSlice("a", a)
}

func printSlice(s string, x []byte) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func SlicesRange() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512}
	var pownew []byte = make([]byte, 0, 20)
	pownew = append(pownew, 1)
	pownew = append(pownew, 2)
	//note i,v will be the temp copy value from the array pow. so if you want to change value in array pow.
	//you need to use array idx like : pow[i] ... !!!
	//range loop array only count to len(array)
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	for i, v := range pownew {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
