package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Printf("%c", s[0])
	for i := 1; i < len(s); i++ {
		fmt.Print("*")
		fmt.Printf("%c", s[i])
	}
	fmt.Print("\n")
}
