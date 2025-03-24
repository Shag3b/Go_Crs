package basics

func init() {
	println("initializing package") //first out
}
func init() {
	println("initializing package2") //second out
}
func main() {
	//init func hasn't parameters and return
	//exe before main func
	//name should be init

	println("inside main") //last out
} /*
##Practical Use Cases
==>> Setup Tasks
==>> Configuration
==>> Registering Components
==>> Database Initialization
##Best Practices
==>> Avoid Side Effects
==>> Initialization Order
==>> Documentation*/
