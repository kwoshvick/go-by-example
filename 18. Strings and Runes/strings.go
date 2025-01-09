package main

import "fmt"

func main() {

	const s = "สวัสดี"

	fmt.Printf("%T", s)
	fmt.Println()
	fmt.Println("Len: ", len(s))

	// produces hex value for all bytes -> 18 bytes
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}

	fmt.Println()

	// produces unicode for each byte // the rune value
	for i := 0; i < len(s); i++ {
		fmt.Print(" ", s[i])
	}

	fmt.Println()
	// gives the string characters
	for i := 0; i < len(s); i++ {
		fmt.Print(" ", string(s[i]))
	}

	// unicode / rune value
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

}
