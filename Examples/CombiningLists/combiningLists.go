package main

import (
	"fmt"
	//"time"
)

func main() {
	x := []int{1, 2, 3, 4, 5}
	y := []string{"a", "b", "c", "d", "e"}
	fmt.Println(combineList(x, y))
}

func combineList(a []int, b []string) []string {
	retList := []string{}
	for i := range a {
		retList = append(retList, fmt.Sprint(a[i]))
		retList = append(retList, b[i])
	}
	return retList
}
