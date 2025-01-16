package main

import "fmt"

func printSlice[T any](s []T) {
	for i := range s {
		fmt.Println(s[i])
	}

}

func find[T comparable](s []T, v T) int {
	for i, item := range s {
		if item == v {
			return i
		}
	}
	return -1
}

type triplet[A, B, C any] struct {
	Ap A
	Bp B
	Cp C
}

func Map[K comparable, V any](keys []K, values []V) map[K]V {

	m := make(map[K]V)
	for i, key := range keys {
		m[key] = values[i]
	}

	return m

}

func main() {
	fruits := []string{"apple", "banana", "cherry"}
	printSlice(fruits)

	fmt.Println(find(fruits, "mangoes"))
	fmt.Println(find(fruits, "cherry"))

	myStruct := triplet[string, int, []int]{"Hey", 123, []int{1, 2, 3, 4, 5, 6, 7, 8}}
	fmt.Println(myStruct.Ap, myStruct.Bp, myStruct.Cp)

	myhashMap := Map([]int{1, 2, 3}, fruits)
	fmt.Println(myhashMap)

}
