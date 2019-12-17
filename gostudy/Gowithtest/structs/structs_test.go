package structs

import "testing"


func TestArea(t *testing.T) {

	ShapeTests := []struct {
        name    string
        shape   Shape
        hasPerimeter float64
        hasArea float64
    }{
		//当测试用例不是一系列操作，而是事实的断言时，测试才清晰明了.
        {name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasPerimeter:36, hasArea: 72.0},
        {name: "Circle", shape: Circle{Radius: 10}, hasPerimeter:62.83185307179586231995926937088370, 
                                                    hasArea: 314.1592653589793},
        {name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasPerimeter:36, hasArea: 36.0},
	}

    for _, tt := range ShapeTests {
        // using tt.name from the case to use it as the `t.Run` test name
        t.Run(tt.name, func(t *testing.T) {
            got := tt.shape.Area()
            if got != tt.hasArea {
                t.Errorf("%#v got %.2f want %.2f", tt.shape, got, tt.hasArea)
            }
        })

        t.Run(tt.name, func(t *testing.T) {
            got := tt.shape.Perimeter()
            if got != tt.hasPerimeter {
                t.Errorf("%#v got %.2f want %.2f", tt.shape, got, tt.hasPerimeter)
            }
        })
    }
}