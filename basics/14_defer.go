package basics

import "fmt"

func main() {
	//A defer statement invokes a function whose execution is deferred to the moment the surrounding function
	//first defer last out
	process(10)
}

//defer can be like finally in another lang

func process(i int) {
	defer fmt.Println("defer", i) //last out but it calc
	defer fmt.Println("defer1")
	defer fmt.Println("defer2")
	defer fmt.Println("defer3")
	i++
	fmt.Println("without defer")
	fmt.Println("value of i without defer", i)
}

/*
1)Practical Use Cases
=>Resource Cleanup
=>Unlocking Mutexes
=>Logging and Tracing
2)Best Practices
=>Keep Deferred Actions Short
=>Understand Evaluation Timing
=>Avoid Complex Control Flow
*/
