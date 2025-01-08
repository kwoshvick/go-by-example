package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func hello() func() string {
	return func() string {
		return "Hi !"
	}
}

func main() {

	nextInt := intSeq()

	fmt.Println(nextInt)   //will print address of the function 0x1xxxxxx
	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3

	newInts := intSeq()
	fmt.Println(newInts()) // 1

	greetings := hello()
	fmt.Println(greetings)   //will print address of the function 0x1xxxxxx
	fmt.Println(greetings()) // Hi !

}
