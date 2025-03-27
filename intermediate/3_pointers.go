package intermediate

func main() {
	//• A pointer is a variable that stores the memory address of another variable.
	var ptr *int
	a := 10
	ptr = &a      //referensing
	println(a)    //10
	println(ptr)  //address
	println(&a)   //address
	println(*ptr) //10/dereferensing
	modifyV(ptr)  //call func
	println(a)    //11

}
func modifyV(ptr *int) {
	*ptr++
}

/*
•# Use Cases
==>>. Modify the value of a variable indirectly
==>>. Pass large data structures efficiently between functions
==>>. Manage memory directly for performance reasons

•# Pointer Declaration and Initialization
=>>.Declaration Syntax:
>>>>>>>>>  var ptr *int
>>>>>>>>> `ptr is a pointer to an integer (*int')
=>>. Initialization:
>>>>>             .  var a int = 10
>>>>>             .  ptr = &a // ptr now points to a's memory address

• Pointer Operations: Limited to referencing (`&`) and dereferencing (`*`)
• Nil Pointers
• Go does not support pointer arithmetic like C or C++
• Passing Pointers to Functions
• Pointers to Structs
• Use pointers when a function needs to modify an argument's value
•`unsafe.Pointer(&x)` converts the address of `x` to an `unsafe.Pointer`f
*/
