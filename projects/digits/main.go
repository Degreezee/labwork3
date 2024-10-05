package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	var c byte
	c = '0'
	for i := 0; i < len(s); i++ {
		if s[i] > c {
			c = s[i]
		}
	}
	fmt.Printf("%c \n", c)
}
