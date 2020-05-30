package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	s, _ := reader.ReadString('\n')
	s = strings.ToLower(s)
	if len(s) > 0 {
		if s[0] == 'i' {
			if s[len(s)-2] == 'n' {
				if strings.Contains(s, "a") {
					fmt.Println("Found!")
					return
				}
			}
		}
	}
	fmt.Println("Not Found!")

}
