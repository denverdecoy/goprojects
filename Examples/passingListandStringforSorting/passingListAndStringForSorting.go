// Daily Challenge
package main

import (
	"fmt"
)

/* This problem was asked by Microsoft.
Given a dictionary of words and a string made up of those words (no spaces),
return the original sentence in a list. If there is more than one possible reconstruction,
return any of them. If there is no possible reconstruction, then return null.
For example, given the set of words 'quick', 'brown', 'the', 'fox', and the string "thequickbrownfox",
you should return ['the', 'quick', 'brown', 'fox'].
Given the set of words 'bed', 'bath', 'bedbath', 'and', 'beyond',
and the string "bedbathandbeyond", return either ['bed', 'bath', 'and', 'beyond] or ['bedbath', 'and', 'beyond'].
*/

func main() {
	x := []string{"brown", "over", "jumped", "the", "lazy", "dog", "fox", "sleeping", "quick"}
	y := "thequickbrownfoxjumpedoverthelazysleepingdog"
	a := []string{"bath", "beyond", "bed", "bedbath", "and"}
	b := "bedbathandbeyond"
	fmt.Println(checkMatch(x, y))
	fmt.Println(checkMatch(a, b))
}

func checkMatch(stringList []string, inputString string) []string {
	var tempString string
	var sortedList []string
	bookmark := 0
	for i := range inputString {
		tempString = inputString[bookmark : i+1]
		for j := range stringList {
			fmt.Printf("I is: %v\t J is: %v\t tempString is: %-6s \t\t sortedList is: %v\t Bookmark is: %v\n", i, j, tempString, sortedList, bookmark)
			if tempString == stringList[j] {
				sortedList = append(sortedList, tempString)
				bookmark = i + 1
				tempString = ""
			}
		}
	}
	return sortedList
}
