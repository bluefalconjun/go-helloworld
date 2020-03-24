package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

/*
定义Sleeper接口类型, 实现Sleep()功能.
实际环境中将使用time.Sleep
测试环境中将使用spy.Sleep
*/
type Sleeper interface {
	Sleep()
}

/*
Spysleep 定义
*/
type SpySleeper struct {
	Calls int
}

/*
Spysleep 只记录调用次数
*/
func (s *SpySleeper) Sleep() {
	s.Calls++
}

/*
实际环境的sleeper需要真实延时
*/
type ConfigurableSleeper struct {
	duration time.Duration
}

/*
调用系统sleep
*/
func (o *ConfigurableSleeper) Sleep() {
	time.Sleep(o.duration)
}

/*
倒数函数需要两个接口类型
out负责输出.
sleeper负责延时.
*/
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

/*
实际使用中传入系统timer和标准输出接口
*/
func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second}
	Countdown(os.Stdout, sleeper)
}
