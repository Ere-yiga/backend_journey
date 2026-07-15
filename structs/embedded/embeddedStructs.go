package main

import "fmt"

type Sport struct{
	name string
	position string
}

type Person struct{
	name string
	age uint

	favSport Sport
	//embedding a struct inside another struct, use slice

	favSports [] Sport
}

func main(){
	p1 := Person(age: 23, name: "Tim", favSport: []Sport{"soccer", "D"}) 
	fmt.Println(p1.favSport.position)
}