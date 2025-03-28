package intermediate

import "fmt"

type person struct {
	name string
	age  int
}
type emp struct {
	person //embedded struct
	//employeeInfo person // Embedded struct Named field
	empId  string
	salary float64
}

func (p person) introduce() {
	fmt.Printf("Hi, I'm %s and I'm %d years old.\n", p.name, p.age)
}

//over ride obove method
func (e emp) introduce() {
	fmt.Printf("Hi, I'm %s, employee ID: %s, and I earn %.2f.\n", e.name, e.empId, e.salary)
}
func main() {
	e1 := emp{
		person: person{name: "Khaled", age: 32},
		salary: 3451.21,
		empId:  "es123",
	}
	fmt.Println("Name:", e1.name) // Accessing the embedded struct field emp.person.name
	fmt.Println("Age:", e1.age)   // Same as above
	fmt.Println("Emp ID:", e1.empId)
	fmt.Println("Salary:", e1.salary)
	e1.introduce()
}

/*
• Best Practices and Considerations
==>> Composition over Inheritance
==>> Avoid Diamond Problem
==>> Clarity and Readability
==>> Initialization
*/
