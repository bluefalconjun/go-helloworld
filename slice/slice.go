package slice

import "fmt"

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
