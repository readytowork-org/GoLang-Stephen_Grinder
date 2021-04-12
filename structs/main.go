package main

import "fmt"

type ContactInfo struct {
	email   string
	zipCode int
}

type Person struct {
	firstName string
	lastName  string
	contact   ContactInfo
}

// reciever function
func (p Person) getInfo() {
	fmt.Printf("%+v", p)
}

// to update the actual data, always pass the pointer . not necessarily for just reading the value

func (p *Person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func main() {
	// ONE OF THE WAY TO DECLARE STRUCT
	p1 := &Person{
		firstName: "Alex",
		lastName:  "Roy",
		contact: ContactInfo{
			email:   "alex@gmail.com",
			zipCode: 1234,
		},
	}
	p1.updateName("Jim")
	p1.getInfo()

	// ANOTHER WAY TO DECLARE STUCT
	// var p1 Person
	// p1.firstName = "Alex"
	// p1.lastName = "Roy"
	// fmt.Println(p1)
}
