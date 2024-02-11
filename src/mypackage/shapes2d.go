package mypackage

import "fmt"

type Square struct {
	Base float64
}

func (sqr Square) Area() float64 {
	return sqr.Base * sqr.Base
}

type Rectangle struct {
	Base   float64
	Height float64
}

func (rec Rectangle) Area() float64 {
	return rec.Base * rec.Height
}

type Shapes2D interface {
	Area() float64
}

func ComputeArea(f Shapes2D) {
	fmt.Println("Area:", f.Area())
}
