package basics

import "fmt"

func main() {
	//panic(interface{})
	//we break out func after panic
	//valid
	process(11)
	//invalid
	process(-22)
}
func process(input int) {
	defer fmt.Println("defer1")
	defer fmt.Println("defer2")
	if input < 0 {
		fmt.Println("before banic")
		panic("can't take negative")
	}
	println("processing input:", input)
}
