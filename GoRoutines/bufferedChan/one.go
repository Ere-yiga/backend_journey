package main

import (
	"fmt"
)

func main(){
	// Unbuffered channel, one that can have one value on it at a time
	
	//Buffered channel helps 
	
	ch := make(chan bool, 2)
	ch <- true
	ch <-true
	<-ch

	ch <- true

	fmt.Println("Done")
}