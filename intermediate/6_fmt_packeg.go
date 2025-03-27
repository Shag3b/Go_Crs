package intermediate

import (
	"fmt"
)

func main() {

	// Printing Func
	fmt.Print("Hello ")
	fmt.Print("World!")
	fmt.Print(12, 456)

	fmt.Println("Hello ")
	fmt.Println("World!")
	fmt.Println(12, 456)

	name := "TAG"
	age := 21
	fmt.Printf("Name: %s, Age: %d\n", name, age)
	fmt.Printf("Binary: %b, Hex: %X, Octal: %o\n", age, age, age)

	//===>>>  Formatting Functions
	//sprint func use to concatenate string
	s := fmt.Sprint("Hello ", "world", 43, 345)
	fmt.Println(s)

	s = fmt.Sprintln("Hello", "World!", 123, 456)
	fmt.Print(s)

	sf := fmt.Sprintf("Name: %s, Age %d", name, age)
	fmt.Println(sf)
	fmt.Println(sf)

	// Scanning Functions
	var nameScan string
	var ageScan int
	fmt.Print("Enter your name and age:")
	//fmt.Scan(&nameScan, &ageScan)      //must u give it to var
	//fmt.Scanln(&nameScan, &ageScan)  //can be nil
	fmt.Scanf("%s %d", &nameScan, &ageScan)
	fmt.Printf("Name: %s, Age: %d\n", nameScan, ageScan)

	// Error Formatting Functions
	err := checkAge(ageScan)
	if err != nil {
		fmt.Println("error: ", err)
	}
}

func checkAge(age int) error {
	if age < 18 {
		return fmt.Errorf("Age %d is too young to drive.", age)
	}
	return nil
}
