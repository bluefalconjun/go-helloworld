package main

import "testing"

//注意引用本地包的路径, 是在$(GO_PATH/src) 路径下搜寻.
//使用包的函数, 需要增加$packagename.PackFun()的方式.
import (
	"github.com/bluefalconjun/lang-study/gostudy/TourGo/basicgrammer/functions"
	"github.com/bluefalconjun/lang-study/gostudy/TourGo/basicgrammer/packages"
)

func TestHello(t *testing.T) {

	t.Run("Simple Test", func(t *testing.T) {
		got := Hello()
		want := "Hello, 世界"
		assertCorrectMessage(t, got, want)
	})
}

func TestSandbox(t *testing.T) {

	t.Run("Simple Test", func(t *testing.T) {
		Sandbox()
	})
}

func TestPackages(t *testing.T) {

	t.Run("Package Test", func(t *testing.T) {
		packages.Packages()
		packages.Imports()
		packages.Exported()
	})
}

func TestFunctions(t *testing.T) {

	t.Run("Package Test", func(t *testing.T) {
		functions.Functions()
	})
}

func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
