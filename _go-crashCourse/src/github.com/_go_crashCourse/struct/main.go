// struct
// types : values receiver, pointer receiver(change values)

package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName string
	lastName  string
	age       int
	gender    string
}

//value reciever
func (p Person) greet() string {
	return "hello, my name is " + p.firstName + strconv.Itoa(p.age) + "  " + p.lastName
}

//pointer reciever
func (p *Person) hadBirthday() {
	p.age++
}
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	person1 := Person{firstName: "heechang", lastName: "kang", age: 24, gender: "f"}
	person2 := Person{"charles", "kang", 23, "m"}

	fmt.Println(person1)
	fmt.Println(person2, person2.age)

	person1.hadBirthday()
	person1.getMarried("mike")
	fmt.Println(person1.greet())
}
