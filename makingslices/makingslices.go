package makingslices

import "fmt"

var b =[]int {1,2,3,4,5}

func Makingslices() {
	a := make([]int, 5)
	printSlice("a", a)


	//b := make([]int, 0, 5)
	//printSlice("b", b)

	//x[lo:hi] is the index of x excluding hi.!!
	c := b[:2]
	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)
}

//len: length of the array/slice
//cap: capacity of the array[which array slice in]
//0<=len(x)<=cap(x)

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
