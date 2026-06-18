/* package main

import "fmt"

func main(){
	vowels := 0

	sentence := "the quick brown fox jumps over the lazy dog"

	for _, char :=range sentence{
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u'{
			vowels++
		}
	}
	fmt.Println(vowels)
 }*/

package main

import "fmt"

func main(){
	vowels := 0
	for _, char :=range "the quick brown fox jumps over the lazy dog"{
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u'{
			vowels++
		}
	}
	fmt.Println(vowels)
}

//I noticed there are two ways to do this....declare a variable, or input thw string in the for statement