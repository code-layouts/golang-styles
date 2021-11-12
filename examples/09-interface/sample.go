package main

import (
	f "fmt"
	"math"
)

type simple interface { // 인터페이스 정의
}

type geometry interface {
	area() float64
	perim() float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	f.Println("geometry:", g)
	f.Println("area:", g.area())
	f.Println("perim:", g.perim())
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func main() {
	f.Println("interface -----------------------------------------")
	var inf simple // 인터페이스 선언
	f.Println("simple:", inf)

	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	f.Println("rect:", r, "circle", c)
	f.Println("geometry for rect ---------------------------------")
	measure(r)
	f.Println("geometry for circle -------------------------------")
	measure(c)
}
