package basics

import "fmt"

func main() {
	//simple iteration
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//iterative over collection
	numbers := []int{1, 2, 3, 4, 5, 6}
	for index, value := range numbers {
		fmt.Printf("index:%d , value:%d\n", index, value)
	}

	//continue vs break
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			continue //rebeat the loop with i++
		}
		println("odd:", i)
		if i == 5 {
			break //exit the loop
		}

	}

	rows := 5
	//outer loop
	for i := 1; i <= rows; i++ {
		for k := 1; k <= rows-i; k++ {
			print(" ")
		} //make space rows - i

		for j := 1; j <= 2*i-1; j++ {
			print("*")
		} //print * depends in doub i -1

		fmt.Println() //to move to next line

	}

	//the update go  ForLoop
	for i := range 10 {
		fmt.Println(10 - i)
	}
	//or
	for i := range 5 {
		i++
		println(i)
	}

}
