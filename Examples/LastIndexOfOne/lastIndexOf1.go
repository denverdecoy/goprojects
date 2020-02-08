// Test
package main

import (
	"fmt"
	//"time"
)

func main() {
	x := []int{0, 0, 0, 0, 1, 0, 0, 1, 0}
	y := []int{0, 0, 0, 0, 0}
	z := []int{0, 0, 0, 0, 0, 0, 1}
	fmt.Println(indexOfOne(x))
	fmt.Println(indexOfOne(y))
	fmt.Println(indexOfOne(z))
}

func indexOfOne(xi []int) int {
	var counter int
	for i := range xi {
		if xi[i] == 1 {
			counter = i
		}
	}
	if counter < 1 {
		return -1
	}
	return counter
}
