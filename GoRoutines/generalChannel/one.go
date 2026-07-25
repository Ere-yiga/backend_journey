// Channel is a special way to pass values between different go routines.

package main

import (
	"fmt"
	//"time"
)

func add(x int, y int, ch chan int){
	//time.Sleep(5 * time.Second)
	fmt.Println(x + y)
	ch <- x + y
}

func main(){
	/* ch := make(chan int)
	go add(5, 10, ch)
	go add(5, 11, ch) */
	/* go add(5, 0, ch)
	go add(5, 5, ch) */

/* 	x := <- ch
	x = <- ch
	x = <- ch
	x = <- ch */
	// fmt.Println(x)

	// If I were to run multiple add
	// If you notice, we don't have control over what prints out...

	// handling several channels

	ch := make(chan int)
	ch2 := make(chan int)
	
	go add(5, 10, ch)
	go add(5, 11, ch)
	
	for i:= 0; i < 2; i++{
		select{
		case x := <- ch:
			fmt.Println(x)
		
		case y := <- ch2:
			fmt.Println(y)
		}
	}
}

