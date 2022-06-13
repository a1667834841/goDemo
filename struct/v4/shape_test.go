package v4

import (
	"math"
	"testing"
)

/********************************
这种定义 interface 的方式与大部分其他编程语言不同。通常接口定义需要这样的代码 My type Foo implements interface Bar。
但是在我们的例子里，
Rectangle 有一个返回值类型为 float64 的方法 Area，所以它满足接口 Shape
Circle 有一个返回值类型为 float64 的方法 Area，所以它满足接口 Shape
string 没有这种方法，所以它不满足这个接口
等等
在 Go 语言中 interface resolution 是隐式的。如果传入的类型匹配接口需要的，则编译正确。
**/

type Shape interface {
	Area() float64
}

// Rectangle 长方形
type Rectangle struct {
	width, height float64
}

func (rec Rectangle) Area() float64 {
	return rec.width * rec.height
}

type Circle struct {
	radius float64
}

func (cir Circle) Area() float64 {
	return cir.radius * cir.radius * math.Pi
}

type Triangle struct {
	Base   float64
	Height float64
}

func (tri Triangle) Area() float64 {
	return tri.Base * tri.Height / 2
}

// func TestArea(t *testing.T) {

// 	areaTests := []struct {
// 		shape Shape
// 		want  float64
// 	}{
// 		{Rectangle{12, 6}, 72.0},
// 		{Circle{10}, 314.1592653589793},
// 	}

// 	for _, tt := range areaTests {
// 		got := tt.shape.Area()
// 		if got != tt.want {
// 			t.Errorf("got %.2f want %.2f", got, tt.want)
// 		}
// 	}

// }

// func TestArea(t *testing.T) {

// 	areaTests := []struct {
// 		shape Shape
// 		want  float64
// 	}{
// 		{Rectangle{12, 6}, 72.0},
// 		{Circle{10}, 314.1592653589793},
// 		{Triangle{12, 6}, 36.0},
// 	}

// 	for _, tt := range areaTests {
// 		got := tt.shape.Area()
// 		if got != tt.want {
// 			t.Errorf("got %.2f want %.2f", got, tt.want)
// 		}
// 	}

// }

func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{width: 12, height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		// using tt.name from the case to use it as the `t.Run` test name
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %.2f want %.2f", tt.shape, got, tt.hasArea)
			}
		})

	}

}
