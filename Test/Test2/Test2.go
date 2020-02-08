// Test2
package main

import (
	"fmt"
)

func main() {
	dogs := []int{1, 2, 3, 5, 6, 8, 9, 10, 11, 12, 15, 17}
	fmt.Println(compress(dogs))
}

func compress(list []int) string {
	var final string
	var temp []int
	counter := 0
	for i := range list {
		temp = append(temp, list[i])
		counter++
		if list[i] != counter {
			fmt.Println("Counter:", counter, "List[i]:", list[i])
			if temp[0] < temp[len(temp)-2] {
				for j := range temp {
					final = fmt.Sprint(temp[j])
					fmt.Println("J loop")
				}
			}
		}
	}
	return final
}
