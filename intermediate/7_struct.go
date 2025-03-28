package intermediate

import "fmt"

type person struct {
	fName   string
	lNAme   string
	age     int
	address Address
	PhoneHomeCell
}
type PhoneHomeCell struct {
	home string
	cell string
}

type Address struct {
	city    string
	country string
}

func main() {

	p := person{
		fName: "Ali",
		lNAme: "deeb",
		age:   30,
		address: Address{
			city:    "London",
			country: "U.K.",
		},
		PhoneHomeCell: PhoneHomeCell{
			home: "465456454",
			cell: "45456464544",
		},
	}
	p1 := person{
		fName: "Ali",

		age: 30,
	}
	p2 := person{
		fName: "Jane",
		age:   25,
	}
	// p2.address.city = "New York"
	// p2.address.country = "USA"
	p3 := person{
		fName: "Jane",
		age:   25,
	}

	fmt.Println("p: ", p.fName)
	fmt.Println("p1:", p1.fName)
	fmt.Println(p.address)
	fmt.Println(p.address.country)
	fmt.Println(p.cell)
	fmt.Println(p.address.city)
	fmt.Println("Are p1 and p2 equal:", p1 == p2)
	fmt.Println("Are p3 and p2 equal:", p3 == p2)

	//func full name call
	fmt.Println("Full Name:", p.fullName())

	//Anonymous struct
	user := struct {
		username string
		email    string
	}{
		username: "Ahmed",
		email:    "ex@fake.com",
	}
	fmt.Println(user.username + " " + user.email)

	//increment func
	fmt.Println("Before increment", p1.age)
	p1.incrementAgeByOne()
	fmt.Println("After increment", p1.age)
}
func (p person) fullName() string {
	return p.fName + " " + p.lNAme
}
func (p *person) incrementAgeByOne() {
	p.age++
}

/*
# Structs are defined using the `type` and `struct` keywords followed by curly braces `{}` containing a list of fields.

# Fields are defined with a name and a type.
==>> Anonymous Structs
==>> Anonymous Fields

# Methods
>>func (value/pointer receiver) methodName(arguments, if any...) <return type, if any> { // Method implementation }

#• Method Declaration
==>> Value receiver method
====>>>>func (t Type) methodName() { // Method implementation }

• Pointer receiver method
===>>>func (t *Type) methodName() { // Method implementation }
Comparing Structs
*/
