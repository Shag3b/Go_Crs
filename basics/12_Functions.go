package basics

import "fmt"

func main() {

	fmt.Println(add(2, 6))

	//Anonymous func
	func() {
		fmt.Println("Anon func")
	}() //call it
	//or assign it to var.
	greet := func() {
		fmt.Println("Anon func")
	}
	greet() //call it
	//////////////////////////////////////////////////////
	oper := add
	result := oper(3, 5)
	fmt.Println(result)

	//call applyoper
	fmt.Println(applyoper(2, 2, add))

	//call return func
	multiby2 := funcReturn(2)
	fmt.Println("multi", multiby2(3))

}

// addition
func add(a, b int) int {
	return a + b
}

//Function takes Func as argument
func applyoper(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}

//func return func
func funcReturn(factor int) func(int) int {
	return func(x int) int {
		return factor * x
	}
}
