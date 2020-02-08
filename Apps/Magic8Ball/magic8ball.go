// Tower of Hanoi in Golang
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type words string

func main() {
	for {
		dramaType(`Ask the magic 8 ball a question ("quit" to exit): `)
		input := bufio.NewReader(os.Stdin)
		a, err := input.ReadString(10)
		if err != nil {
			fmt.Println("Error", err)
		}
		a = strings.ToLower(a)
		a = a[:len(a)-1]
		fmt.Printf("\n")
		if a == "quit" {
			dramaType("Goodbye.")
			time.Sleep(1 * time.Second)
			os.Exit(3)
		}
		dramaType("That's a good question. . . . .     ")
		fmt.Printf("\n")
		dramaType("Let me think. . . . .")
		fmt.Printf("\n")
		dramaType("     - ")
		rand.Seed(time.Now().UnixNano())
		randNum := rand.Intn(20)
		switch randNum {
		case 0:
			dramaType("It is certain.")
		case 1:
			dramaType("It is decidedly so.")
		case 2:
			dramaType("Without a doubt.")
		case 3:
			dramaType("Yes - definitely.")
		case 4:
			dramaType("You may rely on it.")
		case 5:
			dramaType("As I see it, yes.")
		case 6:
			dramaType("Most Likely.")
		case 7:
			dramaType("Outlook Good.")
		case 8:
			dramaType("Yes.")
		case 9:
			dramaType("Signs point to Yes.")
		case 10:
			dramaType("Reply hazy, try again.")
		case 11:
			dramaType("Ask me again later.")
		case 12:
			dramaType("Better not tell you now.")
		case 13:
			dramaType("Cannot predict now.")
		case 14:
			dramaType("Concentrate and ask again.")
		case 15:
			dramaType("Don't count on it.")
		case 16:
			dramaType("My reply is no.")
		case 17:
			dramaType("My sources say no.")
		case 18:
			dramaType("Outlook not so good.")
		case 19:
			dramaType("Very doubtful")
		}
		fmt.Printf("\n")
		time.Sleep(3 * time.Second)
	}
}

func dramaType(s string) {
	for i := range s {
		time.Sleep(25 * time.Millisecond)
		fmt.Print(string(s[i]))
	}

}
