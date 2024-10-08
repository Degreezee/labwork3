package main

import (
	"fmt"
	"math"
)

var p int = 6
var v int = 6
var k int = 1296

func M() float64 {
	return float64(p * v)
}

func W() float64 {
	return math.Sqrt((float64(k) / M()))
}

func T() float64 {
	return 6 / W()
}

func main() {
	fmt.Println(T())
}
