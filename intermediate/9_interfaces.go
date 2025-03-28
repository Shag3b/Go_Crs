package intermediate

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2 * (r.height + r.width)
}
func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}
func (c circle) diameter() float64 {
	return 2 * c.radius
}

func calc(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3.0, height: 4.0}
	c := circle{radius: 5.0}
	calc(c)
	calc(r)

	myPrinter(1, "ds", 23.32, 123435, 10101, true)
}
func myPrinter(a ...interface{}) {
	for _, v := range a {
		fmt.Println(v)
	}
}
