package forgrammer

import "fmt"

func Forgrammer() {
	sum := 1

    //for_init/for_check/for_loop 1st/3rd can be ignored in code.
    //for sum:=1;sum<1000;sum++

    //use this kind of code to replace while
    //for sum<1000 {...}

    //use this kind of code to do while(1)
    //for {...}

	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}
