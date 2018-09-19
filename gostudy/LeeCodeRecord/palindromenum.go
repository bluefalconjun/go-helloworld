package palindromenum

import (
	"fmt"
)

//maxint32 2147483647
//minint32 -2147483648
const int10_9 = 1000000000
const int10_8 = 100000000
const int10_7 = 10000000
const int10_6 = 1000000
const int10_5 = 100000
const int10_4 = 10000
const int10_3 = 1000
const int10_2 = 100
const int10_1 = 10
const int10_0 = 1

var int10list = []int{int10_9, int10_8, int10_7, int10_6, int10_5, int10_4, int10_3, int10_2, int10_1}

func isPalindrome(x int) bool {
	if x == 0 {
		return true
	}
	if x < 0 || x%10 == 0 {
		return false
	}

	return check(x)
}

func check(x int) bool {

	ret := make([]int, 10)
	var numofint = 0
	//var reverseint = 0

	for i, _ := range ret {
		if i == 9 {
			ret[i] = x % int10_1
			fmt.Println("i, ret[i]:", i, ret[i])
			numofint++

			break
		}
		var v = x / (int10list[i])
		ret[i] = v % int10_1
		if ret[i] != 0 && numofint == 0 {
			numofint = 9 - i
		}
		fmt.Println("i, ret[i]:", i, ret[i])
	}
	fmt.Println("numofint:", numofint)

	ret = ret[9-numofint+1:]
	fmt.Println("ret:", ret)
	retreverse := make([]int, numofint)

	for j, _ := range retreverse {

		retreverse[j] = ret[numofint-j-1]
		//ret[9-j] = v
		fmt.Println("j, retreverse[j]:", j, retreverse[j])
	}

	for j, _ := range retreverse {
		if retreverse[j] != ret[j] {
			return false
		}

	}
	return true

}
