package main

import "fmt"

func fibonacci() {
	result := 0
	for i := 0; i <= 10; i++ {
		result = fibonacci_func(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
}

func fibonacci_func(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci_func(n-1) + fibonacci_func(n-2)
	}
	return
}
