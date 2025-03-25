package intermediate

func main() {
	sequence := adder() //adder will call
	println(sequence()) //call the inside func only not adder
	println(sequence()) //call the inside func only not adder
	println(sequence()) //call the inside func only not adder

	subtraction := func() func(int) int {
		countdown := 99

		return func(x int) int {
			countdown -= x
			return countdown
		}
	}()
	//using the closuer subtracter
	println(subtraction(1))
	println(subtraction(1))
	println(subtraction(2))
	println(subtraction(3))

}
func adder() func() int {
	i := 0
	println("i before inc:", i)
	return func() int {
		i++
		println("i after inc:")
		return i
	}
}

//when assign func retern func to var will be run every time call var without run the main func for this return
/*
•# Practical Use Cases
==>>. Stateful Functions
==>>. Encapsulation
==>>. Callbacks
•# Usefulness of Closures
==>>.Encapsulation
==>>.Flexibility
==>>• Readability
==>>• Considerations
#Memory Usage
•==>> Concurrency
==>>• Best Practices
==>>• Limit Scope
==>>• Avoid Overuse
*/
