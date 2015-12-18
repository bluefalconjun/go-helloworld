package main

import "fmt"

func ifelse() {
	var i bool
	i = false
	k()
	if i {
		fmt.Printf("i is true\n")
		return
	} else {
		fmt.Printf("i is false\n")
		return
	}
}

func k() {
	k := 6
	switch k {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:
		fmt.Println("was <= 5")
		fallthrough
	case 6:
		fmt.Println("was <= 6")
		fallthrough
	case 7:
		fmt.Println("was <= 7")
		fallthrough
	case 8:
		fmt.Println("was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}
}
