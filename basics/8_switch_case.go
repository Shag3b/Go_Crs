package basics

import "fmt"

func main() {
	//first switch in go don't have break like another lang
	//auto break
	//but if u need to make it check another case's u should // write "fallthrough"

	fruit := "apple"
	switch fruit {
	case "apple":
		fmt.Println("An apple")
	case "banana":
		fmt.Println("A banana")
	default:
		fmt.Println("Unknown")
	}

	day := "friday"
	switch day {
	case "sunday", "monday", "tuesday", "wednesday", "thursday":
		fmt.Println("weekday")
	case "friday", "saturday":
		fmt.Println("weekend")
	default:
		fmt.Println("invalid")
	}

	//failthrough
	num := 4
	switch {
	case 0 < num: //print
		fmt.Println("num is positive")
		fallthrough //to check after case
	case num%2 == 0: //print
		fmt.Println("num is even")
	case num < 10: //will nut print becouse has break
		fmt.Println("less than 10")
	default:
		fmt.Println("non")
	}

	checkType(10)
	checkType(3.14)
	checkType("Hello")
	checkType(true)
}

func checkType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("It's an integer")
	case float64:
		fmt.Println("It's float")
	case string:
		fmt.Println("It's a string")
	default:
		fmt.Println("Unknown Type")
	}
}
