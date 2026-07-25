//Goroutine is the implementation of concurrency in Go
package main

import (
	"fmt"
	"time"
)

func run(){
	time.Sleep(2 * time.Second)
	fmt.Println("run")
}

func run2(){
	time.Sleep(2 * time.Second)
	fmt.Println("run2")
}

func run3(){
	time.Sleep(2 * time.Second)
	fmt.Println("run3")
}

func main(){
	go run()
	go run2()
	go run3()

	time.Sleep(2 * time.Second)
	fmt.Println("Done")
}
