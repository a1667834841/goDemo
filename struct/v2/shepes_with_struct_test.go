package main

import "testing"

// 参数变为struct 结构体

// Rectangle 长方形
type Rectangle struct {
	width, height float64
}

// TestPerimeter 测试求周长方法
//  @param t
func TestPerimeter(t *testing.T) {
	got := Perimeter(Rectangle{10.0, 10.0})
	want := 40.00

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

// TestArea 测试面积方法
//  @param t
func TestArea(t *testing.T) {
	got := Area(Rectangle{2.0, 6.0})
	want := 12.00

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

// Perimeter 周长方法
//  @param rectangle 长方形
//  @return float64 周长
func Perimeter(rectangle Rectangle) float64 {
	return (rectangle.height + rectangle.width) * 2

}

// Area 面积
//  @param rectangle 长方形
//  @return float64 面积
func Area(rectangle Rectangle) float64 {
	return rectangle.width * rectangle.height
}
