package main

import "fmt"

func main() {
	var a, s, r, c int
	r = 1
	s = 0
	fmt.Scan(&a)
	for a > 0 {
		c = a % 10
		s = r*c*c + s
		if c > 3 {
			r = r * 100
		} else {
			r = r * 10
		}
		a = a / 10
	}
	fmt.Println(s)
}
