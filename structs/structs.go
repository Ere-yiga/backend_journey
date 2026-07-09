//structs are similar to classes
package main

import "fmt"

//how to create a struct
type Person struct{
	//field
	name string
	age uint
}

//how to use the struct
func main(){
	//writing in the order the field is written in
	p1 := Person{name: "Tim", age: 33}
	fmt.Println(p1)
}

//creating different methods in the struct