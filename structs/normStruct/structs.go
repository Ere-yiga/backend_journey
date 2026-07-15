//structs are similar to classes
package main

import "fmt"

//how to create a struct
type Person struct{
	//field
	name string
	age uint
}

//you can also store function as a field value

//creating different methods in the struct

func(p Person) getName() string{
	return p.name
}
func(p Person) getAge() uint{
	return p.age
}

//how to use the struct
func main(){
	//writing in the order the field is written in

	//to initialize

	p1 := Person{name: "Tim", age: 33}
	fmt.Println(p1)

	value := p1.getName()
	fmt.Println(value)

	val := p1.getAge()
	fmt.Println(val)
}

//write the function, and what struct you want it to act on

//NB
//for you to successfully export a portion of your code, you have to make it a capital letter.

//lowercase is private
//uppercase is public - can be accessed

//that's the difference between exported and non exported names