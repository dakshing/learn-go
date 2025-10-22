package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

// Also valid:
// type person struct {
// 	firstName string
// 	lastName  string
//    contactInfo
// }

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	// The below is valid, but not recommended because the order of fields can be mixed up
	// alex := person{"Alex", "Anderson"}

	// Declaring a struct with zero values
	// var alex person

	alex := person{firstName: "Alex", lastName: "Anderson"} // recommended way
	fmt.Println(alex)                                       // {Alex Anderson}

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim.com",
			zipCode: 12345, // trailing comma is necessary here
		},
	}

	jim.updateName("Jimmy")
	jim.print()

	jimPointer := &jim
	jimPointer.updateName("Jimbo")
	jim.print()
}

// receivers are copy of the original struct, not reference.
// therefore, the original struct won't be updated.
// func (p person) updateName(newFirstName string) {
// 	p.firstName = newFirstName
// }

// Now, to update the original struct, we use pointer receivers
func (p *person) updateName(newFirstName string) {
	// (*p).firstName = newFirstName
	// shorthand for the above line
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
