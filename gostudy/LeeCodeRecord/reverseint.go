package reverseint

import (
	"fmt"
	"math"
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

func reverse(x int) int {
	var minusnum = 0

	if x == 0 {
		return x
	}

	if x < 0 {
		if x == math.MinInt32 {
			return 0
		} else {
			tmpx := float64(x)
			tmpx = math.Abs(tmpx)
			x = int(tmpx)
			minusnum = 1
		}
	}

	var reverseint = 0
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

	retreverse := make([]int, numofint)

	for j, _ := range retreverse {

		retreverse[j] = ret[9-j]
		//ret[9-j] = v
		fmt.Println("j, retreverse[j]:", j, retreverse[j])
	}

	reverseint += retreverse[numofint-1]
	for i := 0; i < numofint; i++ {
		if i == numofint-1 {
			break
		}
		reverseint += retreverse[i] * int10list[10-numofint+i]
		fmt.Println("reverseint, add :", reverseint, retreverse[i]*int10list[10-numofint+i])
	}

	fmt.Println("reverseint:", reverseint)

	if reverseint > math.MaxInt32 {
		return 0
	}

	if minusnum == 1 {
		return (0 - reverseint)
	}

	return reverseint

}
