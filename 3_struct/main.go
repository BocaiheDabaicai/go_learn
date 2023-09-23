package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	//contact   contactInfo
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Bruce",
		//contact: contactInfo{
		//	email:   "jim@gmail.com",
		//	zipCode: 95421,
		//},
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 95421,
		},
	}

	//jim.print()
	
	jimPointer := &jim
	jimPointer.updateName("Bob")
	jim.print()
}

func (pointToPerson *person) updateName(newFirstName string) {
	(*pointToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
