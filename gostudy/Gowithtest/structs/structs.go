package structs

import (
	"math"
)

/*

声明结构体以创建我们自己的类型，让我们把数据集合在一起并达到简化代码的目地
声明接口，这样我们可以定义适合不同参数类型的函数（参数多态）
在自己的数据类型中添加方法以实现接口
列表驱动测试让断言更清晰，这样可以使测试文件更易于扩展和维护

*/

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Shape interface {
	Perimeter() float64
	Area() float64
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (t Triangle) Perimeter() float64 {
	return 2 * (t.Base + t.Height)
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height / 2
}
