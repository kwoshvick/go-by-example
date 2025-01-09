package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {

	fmt.Println(person{name: "Bob", age: 20})   //{Bob 20}
	fmt.Println(person{name: "Alice", age: 30}) //{Alice 30}
	fmt.Println(person{name: "Fred"})           //{Fred 0}
	fmt.Println(&person{"Victor", 31})          //&{Victor 31}

	fmt.Println(newPerson("Jon")) //&{Jon 42}

	s := person{"sean", 30}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age) // 30

	sp.age = 51
	fmt.Println(sp.age) //51
	fmt.Println(s.age)  // 51

	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog) //{Rex true}

}
