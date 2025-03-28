package intermediate

import "fmt"

type Rectangle struct {
	len   float64
	width float64
}
type Shape struct {
	Rectangle
}

// Method with pointer receiver
func (r *Rectangle) Scale(factor float64) {
	r.len *= factor
	r.width *= factor
}
func (r Rectangle) Area() float64 {
	return r.len * r.width

}

func main() {

	rect := Rectangle{len: 12, width: 10}
	fmt.Println(rect.Area())
	rect.Scale(2)
	fmt.Println(rect.Area())

	num := MyInt(-8)
	num1 := MyInt(8)
	fmt.Println("num: ", num.Ispos())
	fmt.Println("num1: ", num1.Ispos())
	fmt.Println(num1.welc())

	S := Shape{Rectangle: Rectangle{len: 10, width: 12}}
	fmt.Println(S.Area())

}

type MyInt int

//method on a user def type
func (m MyInt) Ispos() bool {
	return m > 0
}
func (MyInt) welc() string {
	return "welcome to MyInt"
}
