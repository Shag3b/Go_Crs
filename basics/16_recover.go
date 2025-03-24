package basics

import "fmt"

func main() {
	procesas()
	println("returned from proc")
}
func procesas() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("rec:", r)
		}
		/*another way
		r := recover()
		if r != nil {
			fmt.Println("rec:", r)
		}*/

	}()
	fmt.Println("start proc")
	panic("error")
	fmt.Println("will not exe.")
}

/*
##Practical Use Cases
==>> Graceful Recovery
==>> Cleanup
==>> Logging and Reporting
>>>>>>>>
##Best Practices
==>> Use with Defer
==>> Avoid Silent Recovery
==>> Avoid Overuse*/
