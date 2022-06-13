package main

import (
	"math"
	"testing"
)

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

func TestArea(t *testing.T) {

	t.Run("rectangle", func(t *testing.T) {
		rec := Rectangle{
			10.0, 10.0,
		}
		got := rec.Area()
		want := 100.0

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("circle", func(t *testing.T) {
		cir := Circle{
			10.0,
		}
		got := cir.Area()
		want := 314.16

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func Area() {

}
