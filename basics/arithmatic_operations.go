package basics

import "fmt"

func main() {
	//decl variable
	var a, b int = 10, 3
	var result int

	//sum
	result = a + b
	fmt.Println("sum:", result)

	//sub
	result = a - b
	fmt.Println("sub:", result)

	//mul *
	result = a * b
	fmt.Println("mul:", result)

	//div
	result = a / b
	fmt.Println("div:", result)

	//mod
	result = a % b
	fmt.Println("mod:", result)
}
