package intermediate

func main() {
	//factorial func call
	println(factorial(5))
	//sum of dg call
	println(sumOfDigit(7512))
	println(sumOfDigit(12))

}
func factorial(num int) int {
	if num == 0 {
		return 1
	} else {
		return num * factorial(num-1)
	}
}

func sumOfDigit(num int) int {
	if num <= 9 {
		return num
	}
	return num%10 + sumOfDigit(num/10)

}

/*
• Practical Use Cases
==>> Mathematical Algorithms
==>> Tree and Graph Traversal
==>> Divide and Conquer Algorithms
• Benefits of Recursion
==>> Simplicity
==>> Clarity
==>> Flexibility
• Considerations
==>> Performance
==>> Base Case
• Best Practices
==>> Testing
==>> Optimization
==>> Recursive case
*/
