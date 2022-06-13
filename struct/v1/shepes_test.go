package main

import "testing"

// https://studygolang.gitbook.io/learn-go-with-tests/go-ji-chu/structs-methods-and-interfaces

// TestPerimeter 测试求周长方法
//  @param t
func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.00

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

// TestArea 测试面积方法
//  @param t
func TestArea(t *testing.T) {
	got := Area(10.0, 10.0)
	want := 100.00

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

// Perimeter 周长方法
//  @param width 宽
//  @param height 长
//  @return float64 周长
func Perimeter(width float64, height float64) float64 {
	return width*2 + height*2
}

// Area 面积
//  @param width 宽
//  @param height 高
//  @return float64 面积
func Area(width, height float64) float64 {
	return width * height
}
