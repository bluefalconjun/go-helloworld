package runtime

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func RuntimeDemo() {
	go say("world") //开一个新的Goroutines执行
	say("hello")    //当前Goroutines执行
}
