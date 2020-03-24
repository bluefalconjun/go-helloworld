package main

import (
	"bytes"
	"reflect"
	"testing"
)

const write = "write"
const sleep = "sleep"

/*

 */
type CountdownOperationsSpy struct {
	Calls []string
}

/*

 */
func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

/*

 */
func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

/*
关于测试驱动开发和Mock的建议:

重构的定义是代码更改，但行为保持不变。
如果您已经决定在理论上进行一些重构，那么你应该能够在没有任何测试更改的情况下进行提交。

所以，在写测试的时候问问自己。
我是在测试我想要的行为还是实现细节？
如果我要重构这段代码，我需要对测试做很多修改吗？

虽然 Go 允许你测试私有函数，但我将避免它作为私有函数与实现有关。

我觉得如果一个测试 超过 3 个模拟，那么它就是警告 —— 是时候重新考虑设计。

小心使用监视器。监视器让你看到你正在编写的算法的内部细节，这是非常有用的，但是这意味着你的测试代码和实现之间的耦合更紧密。
如果你要监视这些细节，请确保你真的在乎这些细节。

没有对代码中重要的区域进行 mock 将会导致难以测试。

在我们的例子中，我们不能测试我们的代码在每个打印之间暂停，但是还有无数其他的例子。
调用一个 可能 失败的服务？想要在一个特定的状态测试您的系统？在不使用 mocking 的情况下测试这些场景是非常困难的。

如果没有 mock，你可能需要设置数据库和其他第三方的东西来测试简单的业务规则。
你可能会进行缓慢的测试，从而导致 缓慢的反馈循环。

当不得不启用一个数据库或者 webservice 去测试某个功能时，由于这种服务的不可靠性，
你将会得到的是一个 脆弱的测试。
*/
func TestCountdown(t *testing.T) {

	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &CountdownOperationsSpy{})

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})

	t.Run("sleep after every print", func(t *testing.T) {
		spySleepPrinter := &CountdownOperationsSpy{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})
}
