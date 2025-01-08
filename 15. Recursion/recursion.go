package main

import "fmt"

func fact(n int) int {
	if n == 1 {
		return 1
	}
	return n * fact(n-1)
}

func main() {

	factorialThree := fact(3) // 3x2x1 -> 6
	fmt.Println("3 factorial is : ", factorialThree)
	fmt.Println("7 factorial is : ", fact(7)) //5040

	//anonymous func -> function associated with a variable, doesnt have a name, inline fuunction
	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println(fib(7)) // 13

	//or
	var fib2 = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println(fib2(7)) // 13

}
